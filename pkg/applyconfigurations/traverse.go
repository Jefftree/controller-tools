/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package applyconfigurations

import (
	"fmt"
	"go/ast"
	"go/types"
	"io"
	"path"
	"strings"
	"unicode"
	"unicode/utf8"

	"sigs.k8s.io/controller-tools/pkg/loader"
	"sigs.k8s.io/controller-tools/pkg/markers"
)

// TODO(jpbetz): This was copied from deepcopy-gen

// codeWriter assists in writing out Go code lines and blocks to a writer.
type codeWriter struct {
	out io.Writer
}

// Line writes a single line.
func (c *codeWriter) Line(line string) {
	if _, err := fmt.Fprintln(c.out, line); err != nil {
		panic(err)
	}
}

// Linef writes a single line with formatting (as per fmt.Sprintf).
func (c *codeWriter) Linef(line string, args ...interface{}) {
	if _, err := fmt.Fprintf(c.out, line+"\n", args...); err != nil {
		panic(err)
	}
}

// importsList keeps track of required imports, automatically assigning aliases
// to import statement.
type importsList struct {
	byPath  map[string]string
	byAlias map[string]string

	pkg *loader.Package
}

// NeedImport marks that the given package is needed in the list of imports,
// returning the ident (import alias) that should be used to reference the package.
func (l *importsList) NeedImport(importPath string) string {
	// we get an actual path from Package, which might include venddored
	// packages if running on a package in vendor.
	if ind := strings.LastIndex(importPath, "/vendor/"); ind != -1 {
		importPath = importPath[ind+8: /* len("/vendor/") */]
	}

	// check to see if we've already assigned an alias, and just return that.
	alias, exists := l.byPath[importPath]
	if exists {
		return alias
	}

	// otherwise, calculate an import alias by joining path parts till we get something unique
	restPath, nextWord := path.Split(importPath)

	for otherPath, exists := "", true; exists && otherPath != importPath; otherPath, exists = l.byAlias[alias] {
		if restPath == "" {
			// do something else to disambiguate if we're run out of parts and
			// still have duplicates, somehow
			alias += "x"
		}

		// can't have a first digit, per Go identifier rules, so just skip them
		for firstRune, runeLen := utf8.DecodeRuneInString(nextWord); unicode.IsDigit(firstRune); firstRune, runeLen = utf8.DecodeRuneInString(nextWord) {
			nextWord = nextWord[runeLen:]
		}

		// make a valid identifier by replacing "bad" characters with underscores
		nextWord = strings.Map(func(r rune) rune {
			if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' {
				return r
			}
			return '_'
		}, nextWord)

		alias = nextWord + alias
		if len(restPath) > 0 {
			restPath, nextWord = path.Split(restPath[:len(restPath)-1] /* chop off final slash */)
		}
	}

	l.byPath[importPath] = alias
	l.byAlias[alias] = importPath
	return alias
}

// ImportSpecs returns a string form of each import spec
// (i.e. `alias "path/to/import").  Aliases are only present
// when they don't match the package name.
func (l *importsList) ImportSpecs() []string {
	res := make([]string, 0, len(l.byPath))
	for importPath, alias := range l.byPath {
		pkg := l.pkg.Imports()[importPath]
		if pkg != nil && pkg.Name == alias {
			// don't print if alias is the same as package name
			// (we've already taken care of duplicates).
			res = append(res, fmt.Sprintf("%q", importPath))
		} else {
			res = append(res, fmt.Sprintf("%s %q", alias, importPath))
		}
	}
	return res
}

// namingInfo holds package and syntax for referencing a field, type,
// etc.  It's used to allow lazily marking import usage.
// You should generally retrieve the syntax using Syntax.
type namingInfo struct {
	// typeInfo is the type being named.
	typeInfo     types.Type
	nameOverride string
}

// Syntax calculates the code representation of the given type or name,
// and returns the apply representation
func (n *namingInfo) Syntax(basePkg *loader.Package, imports *importsList) string {
	if n.nameOverride != "" {
		return n.nameOverride
	}

	// NB(directxman12): typeInfo.String gets us most of the way there,
	// but fails (for us) on named imports, since it uses the full package path.
	switch typeInfo := n.typeInfo.(type) {
	case *types.Named:
		// register that we need an import for this type,
		// so we can get the appropriate alias to use.

		var lastType types.Type
		appendString := "ApplyConfiguration"
		for underlyingType := typeInfo.Underlying(); underlyingType != lastType; lastType, underlyingType = underlyingType, underlyingType.Underlying() {
			if _, isBasic := underlyingType.(*types.Basic); isBasic {
				appendString = ""
			}
		}

		typeName := typeInfo.Obj()
		otherPkg := typeName.Pkg()
		if otherPkg == basePkg.Types {
			// local import
			return typeName.Name() + appendString
		}
		alias := imports.NeedImport(loader.NonVendorPath(otherPkg.Path()))
		return alias + "." + typeName.Name()
	case *types.Basic:
		return typeInfo.String()
	case *types.Pointer:
		return "*" + (&namingInfo{typeInfo: typeInfo.Elem()}).Syntax(basePkg, imports)
	case *types.Slice:
		return "[]" + (&namingInfo{typeInfo: typeInfo.Elem()}).Syntax(basePkg, imports)
	case *types.Map:
		return fmt.Sprintf(
			"map[%s]%s",
			(&namingInfo{typeInfo: typeInfo.Key()}).Syntax(basePkg, imports),
			(&namingInfo{typeInfo: typeInfo.Elem()}).Syntax(basePkg, imports))
	case *types.Interface:
		return "interface{}"
	default:
		basePkg.AddError(fmt.Errorf("name requested for invalid type: %s", typeInfo))
		return typeInfo.String()
	}
}

// copyMethodMakers makes apply configurations for Go types,
// writing them to its codeWriter.
type applyConfigurationMaker struct {
	pkg *loader.Package
	*importsList
	*codeWriter
}

// GenerateTypesFor makes makes apply configuration types for the given type, when appropriate
func (c *applyConfigurationMaker) GenerateTypesFor(root *loader.Package, info *markers.TypeInfo) {
	typeInfo := root.TypesInfo.TypeOf(info.RawSpec.Name)
	if typeInfo == types.Typ[types.Invalid] {
		root.AddError(loader.ErrFromNode(fmt.Errorf("unknown type: %s", info.Name), info.RawSpec))
	}

	// TODO(jpbetz): Generate output here

	c.Linef("// %sApplyConfiguration represents a declarative configuration of the %s type for use", info.Name, info.Name)
	c.Linef("// with apply.")
	c.Linef("type %sApplyConfiguration struct {", info.Name)
	if len(info.Fields) > 0 {
		for _, field := range info.Fields {
			fieldName := field.Name
			fieldType := root.TypesInfo.TypeOf(field.RawField.Type)
			fieldNamingInfo := namingInfo{typeInfo: fieldType}
			fieldTypeString := fieldNamingInfo.Syntax(root, c.importsList)
			if tags, ok := lookupJsonTags(field); ok {
				if tags.inline {
					c.Linef("%s %s `json:\"%s\"`", fieldName, fieldTypeString, tags.String())
				} else {
					tags.omitempty = true
					c.Linef("%s *%s `json:\"%s\"`", fieldName, fieldTypeString, tags.String())
				}
			}
		}
	}
	c.Linef("}")
}

func generatePrivateName(s string) string {
	if len(s) == 0 {
		return s
	}
	s2 := []byte(s)
	s2[0] = s2[0] | ('a' - 'A')
	return fmt.Sprintf("%s", s2)
}

func (c *applyConfigurationMaker) GenerateStruct(root *loader.Package, info *markers.TypeInfo) {
	typeInfo := root.TypesInfo.TypeOf(info.RawSpec.Name)
	if typeInfo == types.Typ[types.Invalid] {
		root.AddError(loader.ErrFromNode(fmt.Errorf("unknown type: %s", info.Name), info.RawSpec))
	}

	c.Linef("// %sApplyConfiguration represents a declarative configuration of the %s type for use", info.Name, info.Name)
	c.Linef("// with apply.")
	c.Linef("type %sApplyConfiguration struct {", info.Name)
	for _, field := range info.Fields {
		privateFieldName := generatePrivateName(field.Name)
		fieldType := root.TypesInfo.TypeOf(field.RawField.Type)
		fieldNamingInfo := namingInfo{typeInfo: fieldType}
		fieldTypeString := fieldNamingInfo.Syntax(root, c.importsList)
		if tags, ok := lookupJsonTags(field); ok {
			if !tags.inline {
				continue
			}
			c.Linef("%s %s // inlined type", privateFieldName, fieldTypeString)
		}
	}
	c.Linef("fields %sFields", generatePrivateName(info.Name))
	c.Linef("}")
}

func (c *applyConfigurationMaker) GenerateFieldsStruct(root *loader.Package, info *markers.TypeInfo) {
	typeInfo := root.TypesInfo.TypeOf(info.RawSpec.Name)
	if typeInfo == types.Typ[types.Invalid] {
		root.AddError(loader.ErrFromNode(fmt.Errorf("unknown type: %s", info.Name), info.RawSpec))
	}

	privateTypeName := generatePrivateName(info.Name)
	c.Linef("// %s owns all fields except inlined fields", privateTypeName)
	// TODO|jefftree: Copy over rest of documentation
	c.Linef("type %sFields struct {", privateTypeName)
	if len(info.Fields) > 0 {
		for _, field := range info.Fields {
			fieldName := field.Name
			fieldType := root.TypesInfo.TypeOf(field.RawField.Type)
			fieldNamingInfo := namingInfo{typeInfo: fieldType}
			fieldTypeString := fieldNamingInfo.Syntax(root, c.importsList)
			if tags, ok := lookupJsonTags(field); ok {
				// TODO:jefftree: Handle inline objects
				if tags.inline {
					continue
				}
				tags.omitempty = true
				c.Linef("%s *%s `json:\"%s\"`", fieldName, fieldTypeString, tags.String())
			}
		}
	}
	c.Linef("}")
}

func (c *applyConfigurationMaker) GenerateStructConstructor(root *loader.Package, info *markers.TypeInfo) {
	c.Linef("// %sApplyConfiguration represents a declarative configuration of the %s type for use", info.Name, info.Name)
	c.Linef("// with apply.")
	c.Linef("func %s() *%sApplyConfiguration {", info.Name, info.Name)
	c.Linef("return &%sApplyConfiguration{}", info.Name)
	c.Linef("}")
}

func (c *applyConfigurationMaker) GenerateMemberSet(field markers.FieldInfo, root *loader.Package, info *markers.TypeInfo) {
	fieldType := root.TypesInfo.TypeOf(field.RawField.Type)
	fieldNamingInfo := namingInfo{typeInfo: fieldType}
	fieldTypeString := fieldNamingInfo.Syntax(root, c.importsList)

	c.Linef("// Set%s sets the %s field in the declarative configuration to the given value", field.Name, field.Name)
	c.Linef("func (b *%sApplyConfiguration) Set%s(value %s) *%sApplyConfiguration {", info.Name, field.Name, fieldTypeString, info.Name)

	if tags, ok := lookupJsonTags(field); ok {
		// TODO|jefftree: Do we support double pointers?
		if tags.inline {
			c.Linef("if value != nil {")
			c.Linef("b.%s = *value", field.Name)
			c.Linef("}")
		} else {
			// TODO|jefftree: Confirm with jpbetz on how to handle fields that are already pointers
			c.Linef("b.fields.%s = &value", field.Name)
		}
	}
	c.Linef("return b")
	c.Linef("}")
}

func (c *applyConfigurationMaker) GenerateMemberRemove(field markers.FieldInfo, root *loader.Package, info *markers.TypeInfo) {
	fieldType := root.TypesInfo.TypeOf(field.RawField.Type)
	fieldNamingInfo := namingInfo{typeInfo: fieldType}
	fieldTypeString := fieldNamingInfo.Syntax(root, c.importsList)

	if tags, ok := lookupJsonTags(field); ok {
		if tags.inline {
			return
		}
	}

	c.Linef("// Remove%s removes the %s field in the declarative configuration", field.Name, field.Name)
	c.Linef("func (b *%sApplyConfiguration) Remove%s(value %s) *%sApplyConfiguration {", info.Name, field.Name, fieldTypeString, info.Name)
	c.Linef("b.fields.%s = nil", field.Name)
	c.Linef("return b")
	c.Linef("}")
}

func (c *applyConfigurationMaker) GenerateMemberGet(field markers.FieldInfo, root *loader.Package, info *markers.TypeInfo) {
	fieldType := root.TypesInfo.TypeOf(field.RawField.Type)
	fieldNamingInfo := namingInfo{typeInfo: fieldType}
	fieldTypeString := fieldNamingInfo.Syntax(root, c.importsList)

	c.Linef("// Get%s gets the %s field in the declarative configuration", field.Name, field.Name)
	c.Linef("func (b *%sApplyConfiguration) Get%s() (value %s, ok bool) {", info.Name, field.Name, fieldTypeString)

	if tags, ok := lookupJsonTags(field); ok {
		if tags.inline {
			c.Linef("return b.%s, true", field.Name)
		} else {
			// TODO|jefftree: Confirm with jpbetz on how to handle fields that are already pointers

			// c.Linef("return b.fields.%s, b.fields.%s != nil", field.Name, field.Name)
			c.Linef("if v := b.fields.%s; v != nil {", field.Name)
			c.Linef("return *v, true")
			c.Linef("}")
			c.Linef("return value, false")
		}
	}
	c.Linef("}")
}

func (c *applyConfigurationMaker) GenerateToUnstructured(root *loader.Package, info *markers.TypeInfo) {
	runtime := c.NeedImport("k8s.io/apimachinery/pkg/runtime")
	var toUnstructured = `
// ToUnstructured converts %s to unstructured.
func (b *%sApplyConfiguration) ToUnstructured() interface{} {
  if b == nil {
    return nil
  }
  b.preMarshal()
  u, err := %s.DefaultUnstructuredConverter.ToUnstructured(&b.fields)
  if err != nil {
    panic(err)
  }
  return u
}
`
	c.Linef(toUnstructured, info.Name, info.Name, runtime)
}

func (c *applyConfigurationMaker) GenerateFromUnstructured(root *loader.Package, info *markers.TypeInfo) {
	runtime := c.NeedImport("k8s.io/apimachinery/pkg/runtime")
	var fromUnstructured = `
// FromUnstructured converts unstructured to %sApplyConfiguration, replacing the contents
// of %sApplyConfiguration.
func (b *%sApplyConfiguration) FromUnstructured(u map[string]interface{}) error {
  m := &%sFields{}
  err := %s.DefaultUnstructuredConverter.FromUnstructured(u, m)
  if err != nil {
    return err
  }
  b.fields = *m
  b.postUnmarshal()
  return nil
}
`
	c.Linef(fromUnstructured, info.Name, info.Name, info.Name, generatePrivateName(info.Name), runtime)
}

func (c *applyConfigurationMaker) GenerateMarshal(root *loader.Package, info *markers.TypeInfo) {
	jsonImport := c.NeedImport("encoding/json")
	var marshal = `
// MarshalJSON marshals %sApplyConfiguration to JSON.
func (b *%sApplyConfiguration) MarshalJSON() ([]byte, error) {
  b.preMarshal()
  return %s.Marshal(b.fields)
}
`
	c.Linef(marshal, info.Name, info.Name, jsonImport)

}

func (c *applyConfigurationMaker) GenerateUnmarshal(root *loader.Package, info *markers.TypeInfo) {
	jsonImport := c.NeedImport("encoding/json")
	var unmarshal = `
// UnmarshalJSON unmarshals JSON into %sApplyConfiguration, replacing the contents of
// %sApplyConfiguration.
func (b *%sApplyConfiguration) UnmarshalJSON(data []byte) error {
  if err := %s.Unmarshal(data, &b.fields); err != nil {
    return err
  }
  b.postUnmarshal()
  return nil
}
`
	c.Linef(unmarshal, info.Name, info.Name, info.Name, jsonImport)
}

func (c *applyConfigurationMaker) GeneratePrePostFunctions(root *loader.Package, info *markers.TypeInfo) {
	c.Linef("func (b *%sApplyConfiguration) preMarshal() {", info.Name)
	for _, field := range info.Fields {
		// fieldName := field.Name
		// privateName := generatePrivateName(fieldName)
		// fieldType := root.TypesInfo.TypeOf(field.RawField.Type)
		// fieldNamingInfo := namingInfo{typeInfo: fieldType}
		// fieldTypeString := fieldNamingInfo.Syntax(root, c.importsList)
		if tags, ok := lookupJsonTags(field); ok {
			if !tags.inline {
				continue
			}
			// c.Linef("if v, ok := b.%s.Get", privateName)
		}
	}
	c.Linef("}")
	c.Linef("func (b *%sApplyConfiguration) postUnmarshal() {", info.Name)
	c.Linef("}")
}

func (c *applyConfigurationMaker) GenerateListMapAlias(root *loader.Package, info *markers.TypeInfo) {
	var listAlias = `
// %[1]sList represents a listAlias of %[1]sApplyConfiguration.
type %[1]sList []*%[1]sApplyConfiguration
`
	var mapAlias = `
// %[1]sMap represents a map of %[1]sApplyConfiguration.
type %[1]sMap map[string]%[1]sApplyConfiguration
`

	c.Linef(listAlias, info.Name)
	c.Linef(mapAlias, info.Name)
}

// shouldBeApplyConfiguration checks if we're supposed to make apply configurations for the given type.
//
// TODO(jpbetz): Copy over logic for inclusion from requiresApplyConfiguration
func shouldBeApplyConfiguration(pkg *loader.Package, info *markers.TypeInfo) bool {
	if !ast.IsExported(info.Name) {
		return false
	}

	typeInfo := pkg.TypesInfo.TypeOf(info.RawSpec.Name)
	if typeInfo == types.Typ[types.Invalid] {
		pkg.AddError(loader.ErrFromNode(fmt.Errorf("unknown type: %s", info.Name), info.RawSpec))
		return false
	}

	// according to gengo, everything named is an alias, except for an alias to a pointer,
	// which is just a pointer, afaict.  Just roll with it.
	if asPtr, isPtr := typeInfo.(*types.Named).Underlying().(*types.Pointer); isPtr {
		typeInfo = asPtr
	}

	lastType := typeInfo
	if _, isNamed := typeInfo.(*types.Named); isNamed {
		for underlyingType := typeInfo.Underlying(); underlyingType != lastType; lastType, underlyingType = underlyingType, underlyingType.Underlying() {
			// aliases to other things besides basics need copy methods
			// (basics can be straight-up shallow-copied)
			if _, isBasic := underlyingType.(*types.Basic); !isBasic {
				return true
			}
		}
	}

	// structs are the only thing that's not a basic that apply configurations are generated for
	_, isStruct := lastType.(*types.Struct)
	if !isStruct {
		return false
	}
	if _, ok := excludeTypes[info.Name]; ok { // TODO(jpbetz): What to do here?
		return false
	}
	var hasJsonTaggedMembers bool
	for _, field := range info.Fields {
		if _, ok := lookupJsonTags(field); ok {
			hasJsonTaggedMembers = true
		}
	}
	return hasJsonTaggedMembers
}

var (
	rawExtension = "k8s.io/apimachinery/pkg/runtime/RawExtension"
	unknown      = "k8s.io/apimachinery/pkg/runtime/Unknown"
)

// excludeTypes contains well known types that we do not generate apply configurations for.
// Hard coding because we only have two, very specific types that serve a special purpose
// in the type system here.
var excludeTypes = map[string]struct{}{
	rawExtension: {},
	unknown:      {},
	// DO NOT ADD TO THIS LIST. If we need to exclude other types, we should consider allowing the
	// go type declarations to be annotated as excluded from this generator.
}
