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
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/types"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"sigs.k8s.io/controller-tools/pkg/genall"
	"sigs.k8s.io/controller-tools/pkg/loader"
	"sigs.k8s.io/controller-tools/pkg/markers"
)

// Based on deepcopy gen but with legacy marker support removed.

var (
	enablePkgMarker  = markers.Must(markers.MakeDefinition("kubebuilder:object:generate", markers.DescribesPackage, false))
	enableTypeMarker = markers.Must(markers.MakeDefinition("kubebuilder:object:generate", markers.DescribesType, false))
	isObjectMarker   = markers.Must(markers.MakeDefinition("kubebuilder:object:root", markers.DescribesType, false))
)

// +controllertools:marker:generateHelp

// Generator generates code containing apply configuration type implementations.
type Generator struct {
	// HeaderFile specifies the header text (e.g. license) to prepend to generated files.
	HeaderFile string `marker:",optional"`
	// Year specifies the year to substitute for " YEAR" in the header file.
	Year string `marker:",optional"`
}

func (Generator) CheckFilter() loader.NodeFilter {
	return func(node ast.Node) bool {
		// ignore interfaces
		_, isIface := node.(*ast.InterfaceType)
		return !isIface
	}
}

func (Generator) RegisterMarkers(into *markers.Registry) error {
	if err := markers.RegisterAll(into,
		enablePkgMarker, enableTypeMarker, isObjectMarker); err != nil {
		return err
	}
	into.AddHelp(enablePkgMarker,
		markers.SimpleHelp("apply", "enables or disables object apply configuration generation for this package"))
	into.AddHelp(
		enableTypeMarker, markers.SimpleHelp("apply", "overrides enabling or disabling apply configuration generation for this type"))
	into.AddHelp(isObjectMarker,
		markers.SimpleHelp("apply", "enables apply configuration generation for this type"))
	return nil
}

func enabledOnPackage(col *markers.Collector, pkg *loader.Package) (bool, error) {
	pkgMarkers, err := markers.PackageMarkers(col, pkg)
	if err != nil {
		return false, err
	}
	pkgMarker := pkgMarkers.Get(enablePkgMarker.Name)
	if pkgMarker != nil {
		return pkgMarker.(bool), nil
	}

	return false, nil
}

func enabledOnType(allTypes bool, info *markers.TypeInfo) bool {
	if typeMarker := info.Markers.Get(enableTypeMarker.Name); typeMarker != nil {
		return typeMarker.(bool)
	}
	return allTypes || genObjectInterface(info)
}

func genObjectInterface(info *markers.TypeInfo) bool {
	objectEnabled := info.Markers.Get(isObjectMarker.Name)
	if objectEnabled != nil {
		return objectEnabled.(bool)
	}
	return false
}

func groupAndPackageVersion(pkg string) string {
	parts := strings.Split(pkg, "/")
	return parts[len(parts)-2] + "/" + parts[len(parts)-1]
}

func (d Generator) Generate(ctx *genall.GenerationContext) error {
	var headerText string

	if d.HeaderFile != "" {
		headerBytes, err := ctx.ReadFile(d.HeaderFile)
		if err != nil {
			return err
		}
		headerText = string(headerBytes)
	}
	headerText = strings.ReplaceAll(headerText, " YEAR", " "+d.Year)

	objGenCtx := ObjectGenCtx{
		Collector:  ctx.Collector,
		Checker:    ctx.Checker,
		HeaderText: headerText,
	}

	var pkgList []*loader.Package
	visited := make(map[string]*loader.Package)

	//TODO|jefftree: Import variable might be better
	crdRoot := ctx.Roots[0]

	for _, root := range ctx.Roots {
		pkgList = append(pkgList, root)
	}

	for len(pkgList) != 0 {
		pkg := pkgList[0]
		pkgList = pkgList[1:]
		if _, ok := visited[pkg.PkgPath]; ok {
			continue
		}

		visited[pkg.PkgPath] = pkg

		for _, imp := range pkg.Imports() {
			// Only index k8s types
			// TODO|jefftree:There must be a better way to filter this
			if !strings.Contains(imp.PkgPath, "k8s.io") || (!strings.Contains(imp.PkgPath, "api/") && !strings.Contains(imp.PkgPath, "apis/")) {
				continue
			}
			pkgList = append(pkgList, imp)
		}
	}

	var eligibleTypes []types.Type
	for _, pkg := range visited {
		eligibleTypes = append(eligibleTypes, objGenCtx.generateEligibleTypes(pkg)...)
	}

	universe := Universe{
		eligibleTypes: eligibleTypes,
	}

	basePath := filepath.Dir(crdRoot.CompiledGoFiles[0])

	for _, pkg := range visited {
		outContents := objGenCtx.generateForPackage(universe, pkg)
		if outContents == nil {
			continue
		}

		if pkg != ctx.Roots[0] {
			desiredPath := basePath + "/ac/" + groupAndPackageVersion(pkg.PkgPath) + "/"
			if _, err := os.Stat(desiredPath); os.IsNotExist(err) {
				os.MkdirAll(desiredPath, os.ModePerm)
			}
			pkg.CompiledGoFiles[0] = desiredPath
		} else {
			pkg.CompiledGoFiles[0] = basePath + "/ac/"
			os.MkdirAll(pkg.CompiledGoFiles[0], os.ModePerm)
		}
		writeOut(ctx, pkg, outContents)

	}
	return nil
}

// ObjectGenCtx contains the common info for generating apply configuration implementations.
// It mostly exists so that generating for a package can be easily tested without
// requiring a full set of output rules, etc.
type ObjectGenCtx struct {
	Collector  *markers.Collector
	Checker    *loader.TypeChecker
	HeaderText string
}

// writeHeader writes out the build tag, package declaration, and imports
func writeHeader(pkg *loader.Package, out io.Writer, packageName string, imports *importsList, headerText string) {
	// NB(directxman12): blank line after build tags to distinguish them from comments
	_, err := fmt.Fprintf(out, `// +build !ignore_autogenerated

%[3]s

// Code generated by controller-gen. DO NOT EDIT.

package %[1]s

import (
%[2]s
)

`, packageName, strings.Join(imports.ImportSpecs(), "\n"), headerText)
	if err != nil {
		pkg.AddError(err)
	}

}

func (ctx *ObjectGenCtx) generateEligibleTypes(root *loader.Package) []types.Type {
	ctx.Checker.Check(root)
	root.NeedTypesInfo()
	var eligibleTypes []types.Type

	if err := markers.EachType(ctx.Collector, root, func(info *markers.TypeInfo) {
		if shouldBeApplyConfiguration(root, info) {
			typeInfo := root.TypesInfo.TypeOf(info.RawSpec.Name)
			eligibleTypes = append(eligibleTypes, typeInfo)
		}
	}); err != nil {
		root.AddError(err)
		return nil
	}
	return eligibleTypes
}

type Universe struct {
	eligibleTypes []types.Type
}

// generateForPackage generates apply configuration implementations for
// types in the given package, writing the formatted result to given writer.
// May return nil if source could not be generated.
func (ctx *ObjectGenCtx) generateForPackage(universe Universe, root *loader.Package) []byte {
	byType := make(map[string][]byte)
	imports := &importsList{
		byPath:  make(map[string]string),
		byAlias: make(map[string]string),
		pkg:     root,
	}
	// avoid confusing aliases by "reserving" the root package's name as an alias
	imports.byAlias[root.Name] = ""

	if err := markers.EachType(ctx.Collector, root, func(info *markers.TypeInfo) {
		outContent := new(bytes.Buffer)
		// not all types required a generate apply configuration. For example, no apply configuration
		// type is needed for Quantity, IntOrString, RawExtension or Unknown.
		if !shouldBeApplyConfiguration(root, info) {
			return
		}

		copyCtx := &applyConfigurationMaker{
			pkg:         root,
			importsList: imports,
			codeWriter:  &codeWriter{out: outContent},
		}

		// TODO|jefftree: Make this a CLI toggle between builder and go structs
		if false {
			copyCtx.GenerateStruct(&universe, root, info)
			copyCtx.GenerateFieldsStruct(&universe, root, info)
			for _, field := range info.Fields {
				if field.Name != "" {
					copyCtx.GenerateMemberSet(&universe, field, root, info)
					copyCtx.GenerateMemberRemove(&universe, field, root, info)
					copyCtx.GenerateMemberGet(&universe, field, root, info)
				}
			}
			copyCtx.GenerateToUnstructured(root, info)
			copyCtx.GenerateFromUnstructured(root, info)
			copyCtx.GenerateMarshal(root, info)
			copyCtx.GenerateUnmarshal(root, info)
			copyCtx.GeneratePrePostFunctions(root, info)

		} else {
			copyCtx.GenerateTypesFor(&universe, root, info)
			copyCtx.GenerateStructConstructor(root, info)
		}
		// copyCtx.GenerateListMapAlias(root, info)

		outBytes := outContent.Bytes()
		if len(outBytes) > 0 {
			byType[info.Name] = outBytes
		}
	}); err != nil {
		root.AddError(err)
		return nil
	}

	if len(byType) == 0 {
		return nil
	}

	outContent := new(bytes.Buffer)
	writeHeader(root, outContent, root.Name, imports, ctx.HeaderText)
	writeTypes(root, outContent, byType)

	outBytes := outContent.Bytes()
	formattedBytes, err := format.Source(outBytes)
	if err != nil {
		root.AddError(err)
		// we still write the invalid source to disk to figure out what went wrong
	} else {
		outBytes = formattedBytes
	}

	return outBytes
}

// writeTypes writes each method to the file, sorted by type name.
func writeTypes(pkg *loader.Package, out io.Writer, byType map[string][]byte) {
	sortedNames := make([]string, 0, len(byType))
	for name := range byType {
		sortedNames = append(sortedNames, name)
	}
	sort.Strings(sortedNames)

	for _, name := range sortedNames {
		_, err := out.Write(byType[name])
		if err != nil {
			pkg.AddError(err)
		}
	}
}

// writeFormatted outputs the given code, after gofmt-ing it.  If we couldn't gofmt,
// we write the unformatted code for debugging purposes.
func writeOut(ctx *genall.GenerationContext, root *loader.Package, outBytes []byte) {
	outputFile, err := ctx.Open(root, root.Name+"_zz_generated.applyconfigurations.go")
	if err != nil {
		root.AddError(err)
		return
	}
	defer outputFile.Close()
	n, err := outputFile.Write(outBytes)
	if err != nil {
		root.AddError(err)
		return
	}
	if n < len(outBytes) {
		root.AddError(io.ErrShortWrite)
	}
}
