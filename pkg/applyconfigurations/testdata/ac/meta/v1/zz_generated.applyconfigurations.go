// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

// APIGroupApplyConfiguration represents a declarative configuration of the APIGroup type for use
// with apply.
type APIGroupApplyConfiguration struct {
	TypeMetaApplyConfiguration `json:",inline"`
	Name                       *string                                        `json:"name,omitempty"`
	Versions                   *[]GroupVersionForDiscoveryApplyConfiguration  `json:"versions,omitempty"`
	PreferredVersion           *GroupVersionForDiscoveryApplyConfiguration    `json:"preferredVersion,omitempty"`
	ServerAddressByClientCIDRs *[]ServerAddressByClientCIDRApplyConfiguration `json:"serverAddressByClientCIDRs,omitempty"`
}

// APIGroupApplyConfiguration represents a declarative configuration of the APIGroup type for use
// with apply.
func APIGroup() *APIGroupApplyConfiguration {
	return &APIGroupApplyConfiguration{}
}

// APIGroupListApplyConfiguration represents a declarative configuration of the APIGroupList type for use
// with apply.
type APIGroupListApplyConfiguration struct {
	TypeMetaApplyConfiguration `json:",inline"`
	Groups                     *[]APIGroupApplyConfiguration `json:"groups,omitempty"`
}

// APIGroupListApplyConfiguration represents a declarative configuration of the APIGroupList type for use
// with apply.
func APIGroupList() *APIGroupListApplyConfiguration {
	return &APIGroupListApplyConfiguration{}
}

// APIResourceApplyConfiguration represents a declarative configuration of the APIResource type for use
// with apply.
type APIResourceApplyConfiguration struct {
	Name               *string       `json:"name,omitempty"`
	SingularName       *string       `json:"singularName,omitempty"`
	Namespaced         *bool         `json:"namespaced,omitempty"`
	Group              *string       `json:"group,omitempty"`
	Version            *string       `json:"version,omitempty"`
	Kind               *string       `json:"kind,omitempty"`
	Verbs              *metav1.Verbs `json:"verbs,omitempty"`
	ShortNames         *[]string     `json:"shortNames,omitempty"`
	Categories         *[]string     `json:"categories,omitempty"`
	StorageVersionHash *string       `json:"storageVersionHash,omitempty"`
}

// APIResourceApplyConfiguration represents a declarative configuration of the APIResource type for use
// with apply.
func APIResource() *APIResourceApplyConfiguration {
	return &APIResourceApplyConfiguration{}
}

// APIResourceListApplyConfiguration represents a declarative configuration of the APIResourceList type for use
// with apply.
type APIResourceListApplyConfiguration struct {
	TypeMetaApplyConfiguration `json:",inline"`
	GroupVersion               *string                          `json:"groupVersion,omitempty"`
	APIResources               *[]APIResourceApplyConfiguration `json:"resources,omitempty"`
}

// APIResourceListApplyConfiguration represents a declarative configuration of the APIResourceList type for use
// with apply.
func APIResourceList() *APIResourceListApplyConfiguration {
	return &APIResourceListApplyConfiguration{}
}

// APIVersionsApplyConfiguration represents a declarative configuration of the APIVersions type for use
// with apply.
type APIVersionsApplyConfiguration struct {
	TypeMetaApplyConfiguration `json:",inline"`
	Versions                   *[]string                                      `json:"versions,omitempty"`
	ServerAddressByClientCIDRs *[]ServerAddressByClientCIDRApplyConfiguration `json:"serverAddressByClientCIDRs,omitempty"`
}

// APIVersionsApplyConfiguration represents a declarative configuration of the APIVersions type for use
// with apply.
func APIVersions() *APIVersionsApplyConfiguration {
	return &APIVersionsApplyConfiguration{}
}

// ConditionApplyConfiguration represents a declarative configuration of the Condition type for use
// with apply.
type ConditionApplyConfiguration struct {
	Type               *string                 `json:"type,omitempty"`
	Status             *metav1.ConditionStatus `json:"status,omitempty"`
	ObservedGeneration *int64                  `json:"observedGeneration,omitempty"`
	LastTransitionTime *metav1.Time            `json:"lastTransitionTime,omitempty"`
	Reason             *string                 `json:"reason,omitempty"`
	Message            *string                 `json:"message,omitempty"`
}

// ConditionApplyConfiguration represents a declarative configuration of the Condition type for use
// with apply.
func Condition() *ConditionApplyConfiguration {
	return &ConditionApplyConfiguration{}
}

// CreateOptionsApplyConfiguration represents a declarative configuration of the CreateOptions type for use
// with apply.
type CreateOptionsApplyConfiguration struct {
	TypeMetaApplyConfiguration `json:",inline"`
	DryRun                     *[]string `json:"dryRun,omitempty"`
	FieldManager               *string   `json:"fieldManager,omitempty"`
}

// CreateOptionsApplyConfiguration represents a declarative configuration of the CreateOptions type for use
// with apply.
func CreateOptions() *CreateOptionsApplyConfiguration {
	return &CreateOptionsApplyConfiguration{}
}

// DeleteOptionsApplyConfiguration represents a declarative configuration of the DeleteOptions type for use
// with apply.
type DeleteOptionsApplyConfiguration struct {
	TypeMetaApplyConfiguration `json:",inline"`
	GracePeriodSeconds         *int64                           `json:"gracePeriodSeconds,omitempty"`
	Preconditions              *PreconditionsApplyConfiguration `json:"preconditions,omitempty"`
	OrphanDependents           *bool                            `json:"orphanDependents,omitempty"`
	PropagationPolicy          *metav1.DeletionPropagation      `json:"propagationPolicy,omitempty"`
	DryRun                     *[]string                        `json:"dryRun,omitempty"`
}

// DeleteOptionsApplyConfiguration represents a declarative configuration of the DeleteOptions type for use
// with apply.
func DeleteOptions() *DeleteOptionsApplyConfiguration {
	return &DeleteOptionsApplyConfiguration{}
}

// ExportOptionsApplyConfiguration represents a declarative configuration of the ExportOptions type for use
// with apply.
type ExportOptionsApplyConfiguration struct {
	TypeMetaApplyConfiguration `json:",inline"`
	Export                     *bool `json:"export,omitempty"`
	Exact                      *bool `json:"exact,omitempty"`
}

// ExportOptionsApplyConfiguration represents a declarative configuration of the ExportOptions type for use
// with apply.
func ExportOptions() *ExportOptionsApplyConfiguration {
	return &ExportOptionsApplyConfiguration{}
}

// GetOptionsApplyConfiguration represents a declarative configuration of the GetOptions type for use
// with apply.
type GetOptionsApplyConfiguration struct {
	TypeMetaApplyConfiguration `json:",inline"`
	ResourceVersion            *string `json:"resourceVersion,omitempty"`
}

// GetOptionsApplyConfiguration represents a declarative configuration of the GetOptions type for use
// with apply.
func GetOptions() *GetOptionsApplyConfiguration {
	return &GetOptionsApplyConfiguration{}
}

// GroupKindApplyConfiguration represents a declarative configuration of the GroupKind type for use
// with apply.
type GroupKindApplyConfiguration struct {
	Group *string `json:"group,omitempty"`
	Kind  *string `json:"kind,omitempty"`
}

// GroupKindApplyConfiguration represents a declarative configuration of the GroupKind type for use
// with apply.
func GroupKind() *GroupKindApplyConfiguration {
	return &GroupKindApplyConfiguration{}
}

// GroupResourceApplyConfiguration represents a declarative configuration of the GroupResource type for use
// with apply.
type GroupResourceApplyConfiguration struct {
	Group    *string `json:"group,omitempty"`
	Resource *string `json:"resource,omitempty"`
}

// GroupResourceApplyConfiguration represents a declarative configuration of the GroupResource type for use
// with apply.
func GroupResource() *GroupResourceApplyConfiguration {
	return &GroupResourceApplyConfiguration{}
}

// GroupVersionApplyConfiguration represents a declarative configuration of the GroupVersion type for use
// with apply.
type GroupVersionApplyConfiguration struct {
	Group   *string `json:"group,omitempty"`
	Version *string `json:"version,omitempty"`
}

// GroupVersionApplyConfiguration represents a declarative configuration of the GroupVersion type for use
// with apply.
func GroupVersion() *GroupVersionApplyConfiguration {
	return &GroupVersionApplyConfiguration{}
}

// GroupVersionForDiscoveryApplyConfiguration represents a declarative configuration of the GroupVersionForDiscovery type for use
// with apply.
type GroupVersionForDiscoveryApplyConfiguration struct {
	GroupVersion *string `json:"groupVersion,omitempty"`
	Version      *string `json:"version,omitempty"`
}

// GroupVersionForDiscoveryApplyConfiguration represents a declarative configuration of the GroupVersionForDiscovery type for use
// with apply.
func GroupVersionForDiscovery() *GroupVersionForDiscoveryApplyConfiguration {
	return &GroupVersionForDiscoveryApplyConfiguration{}
}

// GroupVersionKindApplyConfiguration represents a declarative configuration of the GroupVersionKind type for use
// with apply.
type GroupVersionKindApplyConfiguration struct {
	Group   *string `json:"group,omitempty"`
	Version *string `json:"version,omitempty"`
	Kind    *string `json:"kind,omitempty"`
}

// GroupVersionKindApplyConfiguration represents a declarative configuration of the GroupVersionKind type for use
// with apply.
func GroupVersionKind() *GroupVersionKindApplyConfiguration {
	return &GroupVersionKindApplyConfiguration{}
}

// GroupVersionResourceApplyConfiguration represents a declarative configuration of the GroupVersionResource type for use
// with apply.
type GroupVersionResourceApplyConfiguration struct {
	Group    *string `json:"group,omitempty"`
	Version  *string `json:"version,omitempty"`
	Resource *string `json:"resource,omitempty"`
}

// GroupVersionResourceApplyConfiguration represents a declarative configuration of the GroupVersionResource type for use
// with apply.
func GroupVersionResource() *GroupVersionResourceApplyConfiguration {
	return &GroupVersionResourceApplyConfiguration{}
}

// LabelSelectorApplyConfiguration represents a declarative configuration of the LabelSelector type for use
// with apply.
type LabelSelectorApplyConfiguration struct {
	MatchLabels      *map[string]string                            `json:"matchLabels,omitempty"`
	MatchExpressions *[]LabelSelectorRequirementApplyConfiguration `json:"matchExpressions,omitempty"`
}

// LabelSelectorApplyConfiguration represents a declarative configuration of the LabelSelector type for use
// with apply.
func LabelSelector() *LabelSelectorApplyConfiguration {
	return &LabelSelectorApplyConfiguration{}
}

// LabelSelectorRequirementApplyConfiguration represents a declarative configuration of the LabelSelectorRequirement type for use
// with apply.
type LabelSelectorRequirementApplyConfiguration struct {
	Key      *string                       `json:"key,omitempty"`
	Operator *metav1.LabelSelectorOperator `json:"operator,omitempty"`
	Values   *[]string                     `json:"values,omitempty"`
}

// LabelSelectorRequirementApplyConfiguration represents a declarative configuration of the LabelSelectorRequirement type for use
// with apply.
func LabelSelectorRequirement() *LabelSelectorRequirementApplyConfiguration {
	return &LabelSelectorRequirementApplyConfiguration{}
}

// ListApplyConfiguration represents a declarative configuration of the List type for use
// with apply.
type ListApplyConfiguration struct {
	TypeMetaApplyConfiguration  `json:",inline"`
	*ListMetaApplyConfiguration `json:"metadata,omitempty"`
	Items                       *[]runtime.RawExtension `json:"items,omitempty"`
}

// ListApplyConfiguration represents a declarative configuration of the List type for use
// with apply.
func List() *ListApplyConfiguration {
	return &ListApplyConfiguration{}
}

// ListMetaApplyConfiguration represents a declarative configuration of the ListMeta type for use
// with apply.
type ListMetaApplyConfiguration struct {
	SelfLink           *string `json:"selfLink,omitempty"`
	ResourceVersion    *string `json:"resourceVersion,omitempty"`
	Continue           *string `json:"continue,omitempty"`
	RemainingItemCount *int64  `json:"remainingItemCount,omitempty"`
}

// ListMetaApplyConfiguration represents a declarative configuration of the ListMeta type for use
// with apply.
func ListMeta() *ListMetaApplyConfiguration {
	return &ListMetaApplyConfiguration{}
}

// ListOptionsApplyConfiguration represents a declarative configuration of the ListOptions type for use
// with apply.
type ListOptionsApplyConfiguration struct {
	TypeMetaApplyConfiguration `json:",inline"`
	LabelSelector              *string                      `json:"labelSelector,omitempty"`
	FieldSelector              *string                      `json:"fieldSelector,omitempty"`
	Watch                      *bool                        `json:"watch,omitempty"`
	AllowWatchBookmarks        *bool                        `json:"allowWatchBookmarks,omitempty"`
	ResourceVersion            *string                      `json:"resourceVersion,omitempty"`
	ResourceVersionMatch       *metav1.ResourceVersionMatch `json:"resourceVersionMatch,omitempty"`
	TimeoutSeconds             *int64                       `json:"timeoutSeconds,omitempty"`
	Limit                      *int64                       `json:"limit,omitempty"`
	Continue                   *string                      `json:"continue,omitempty"`
}

// ListOptionsApplyConfiguration represents a declarative configuration of the ListOptions type for use
// with apply.
func ListOptions() *ListOptionsApplyConfiguration {
	return &ListOptionsApplyConfiguration{}
}

// ManagedFieldsEntryApplyConfiguration represents a declarative configuration of the ManagedFieldsEntry type for use
// with apply.
type ManagedFieldsEntryApplyConfiguration struct {
	Manager    *string                            `json:"manager,omitempty"`
	Operation  *metav1.ManagedFieldsOperationType `json:"operation,omitempty"`
	APIVersion *string                            `json:"apiVersion,omitempty"`
	Time       *metav1.Time                       `json:"time,omitempty"`
	FieldsType *string                            `json:"fieldsType,omitempty"`
	FieldsV1   *metav1.FieldsV1                   `json:"fieldsV1,omitempty"`
}

// ManagedFieldsEntryApplyConfiguration represents a declarative configuration of the ManagedFieldsEntry type for use
// with apply.
func ManagedFieldsEntry() *ManagedFieldsEntryApplyConfiguration {
	return &ManagedFieldsEntryApplyConfiguration{}
}

// ObjectMetaApplyConfiguration represents a declarative configuration of the ObjectMeta type for use
// with apply.
type ObjectMetaApplyConfiguration struct {
	Name                       *string                                 `json:"name,omitempty"`
	GenerateName               *string                                 `json:"generateName,omitempty"`
	Namespace                  *string                                 `json:"namespace,omitempty"`
	SelfLink                   *string                                 `json:"selfLink,omitempty"`
	UID                        *types.UID                              `json:"uid,omitempty"`
	ResourceVersion            *string                                 `json:"resourceVersion,omitempty"`
	Generation                 *int64                                  `json:"generation,omitempty"`
	CreationTimestamp          *metav1.Time                            `json:"creationTimestamp,omitempty"`
	DeletionTimestamp          *metav1.Time                            `json:"deletionTimestamp,omitempty"`
	DeletionGracePeriodSeconds *int64                                  `json:"deletionGracePeriodSeconds,omitempty"`
	Labels                     *map[string]string                      `json:"labels,omitempty"`
	Annotations                *map[string]string                      `json:"annotations,omitempty"`
	OwnerReferences            *[]OwnerReferenceApplyConfiguration     `json:"ownerReferences,omitempty"`
	Finalizers                 *[]string                               `json:"finalizers,omitempty"`
	ClusterName                *string                                 `json:"clusterName,omitempty"`
	ManagedFields              *[]ManagedFieldsEntryApplyConfiguration `json:"managedFields,omitempty"`
}

// ObjectMetaApplyConfiguration represents a declarative configuration of the ObjectMeta type for use
// with apply.
func ObjectMeta() *ObjectMetaApplyConfiguration {
	return &ObjectMetaApplyConfiguration{}
}

// OwnerReferenceApplyConfiguration represents a declarative configuration of the OwnerReference type for use
// with apply.
type OwnerReferenceApplyConfiguration struct {
	APIVersion         *string    `json:"apiVersion,omitempty"`
	Kind               *string    `json:"kind,omitempty"`
	Name               *string    `json:"name,omitempty"`
	UID                *types.UID `json:"uid,omitempty"`
	Controller         *bool      `json:"controller,omitempty"`
	BlockOwnerDeletion *bool      `json:"blockOwnerDeletion,omitempty"`
}

// OwnerReferenceApplyConfiguration represents a declarative configuration of the OwnerReference type for use
// with apply.
func OwnerReference() *OwnerReferenceApplyConfiguration {
	return &OwnerReferenceApplyConfiguration{}
}

// PartialObjectMetadataApplyConfiguration represents a declarative configuration of the PartialObjectMetadata type for use
// with apply.
type PartialObjectMetadataApplyConfiguration struct {
	TypeMetaApplyConfiguration    `json:",inline"`
	*ObjectMetaApplyConfiguration `json:"metadata,omitempty"`
}

// PartialObjectMetadataApplyConfiguration represents a declarative configuration of the PartialObjectMetadata type for use
// with apply.
func PartialObjectMetadata() *PartialObjectMetadataApplyConfiguration {
	return &PartialObjectMetadataApplyConfiguration{}
}

// PartialObjectMetadataListApplyConfiguration represents a declarative configuration of the PartialObjectMetadataList type for use
// with apply.
type PartialObjectMetadataListApplyConfiguration struct {
	TypeMetaApplyConfiguration  `json:",inline"`
	*ListMetaApplyConfiguration `json:"metadata,omitempty"`
	Items                       *[]PartialObjectMetadataApplyConfiguration `json:"items,omitempty"`
}

// PartialObjectMetadataListApplyConfiguration represents a declarative configuration of the PartialObjectMetadataList type for use
// with apply.
func PartialObjectMetadataList() *PartialObjectMetadataListApplyConfiguration {
	return &PartialObjectMetadataListApplyConfiguration{}
}

// PatchOptionsApplyConfiguration represents a declarative configuration of the PatchOptions type for use
// with apply.
type PatchOptionsApplyConfiguration struct {
	TypeMetaApplyConfiguration `json:",inline"`
	DryRun                     *[]string `json:"dryRun,omitempty"`
	Force                      *bool     `json:"force,omitempty"`
	FieldManager               *string   `json:"fieldManager,omitempty"`
}

// PatchOptionsApplyConfiguration represents a declarative configuration of the PatchOptions type for use
// with apply.
func PatchOptions() *PatchOptionsApplyConfiguration {
	return &PatchOptionsApplyConfiguration{}
}

// PreconditionsApplyConfiguration represents a declarative configuration of the Preconditions type for use
// with apply.
type PreconditionsApplyConfiguration struct {
	UID             *types.UID `json:"uid,omitempty"`
	ResourceVersion *string    `json:"resourceVersion,omitempty"`
}

// PreconditionsApplyConfiguration represents a declarative configuration of the Preconditions type for use
// with apply.
func Preconditions() *PreconditionsApplyConfiguration {
	return &PreconditionsApplyConfiguration{}
}

// RootPathsApplyConfiguration represents a declarative configuration of the RootPaths type for use
// with apply.
type RootPathsApplyConfiguration struct {
	Paths *[]string `json:"paths,omitempty"`
}

// RootPathsApplyConfiguration represents a declarative configuration of the RootPaths type for use
// with apply.
func RootPaths() *RootPathsApplyConfiguration {
	return &RootPathsApplyConfiguration{}
}

// ServerAddressByClientCIDRApplyConfiguration represents a declarative configuration of the ServerAddressByClientCIDR type for use
// with apply.
type ServerAddressByClientCIDRApplyConfiguration struct {
	ClientCIDR    *string `json:"clientCIDR,omitempty"`
	ServerAddress *string `json:"serverAddress,omitempty"`
}

// ServerAddressByClientCIDRApplyConfiguration represents a declarative configuration of the ServerAddressByClientCIDR type for use
// with apply.
func ServerAddressByClientCIDR() *ServerAddressByClientCIDRApplyConfiguration {
	return &ServerAddressByClientCIDRApplyConfiguration{}
}

// StatusApplyConfiguration represents a declarative configuration of the Status type for use
// with apply.
type StatusApplyConfiguration struct {
	TypeMetaApplyConfiguration  `json:",inline"`
	*ListMetaApplyConfiguration `json:"metadata,omitempty"`
	Status                      *string                          `json:"status,omitempty"`
	Message                     *string                          `json:"message,omitempty"`
	Reason                      *metav1.StatusReason             `json:"reason,omitempty"`
	Details                     *StatusDetailsApplyConfiguration `json:"details,omitempty"`
	Code                        *int32                           `json:"code,omitempty"`
}

// StatusApplyConfiguration represents a declarative configuration of the Status type for use
// with apply.
func Status() *StatusApplyConfiguration {
	return &StatusApplyConfiguration{}
}

// StatusCauseApplyConfiguration represents a declarative configuration of the StatusCause type for use
// with apply.
type StatusCauseApplyConfiguration struct {
	Type    *metav1.CauseType `json:"reason,omitempty"`
	Message *string           `json:"message,omitempty"`
	Field   *string           `json:"field,omitempty"`
}

// StatusCauseApplyConfiguration represents a declarative configuration of the StatusCause type for use
// with apply.
func StatusCause() *StatusCauseApplyConfiguration {
	return &StatusCauseApplyConfiguration{}
}

// StatusDetailsApplyConfiguration represents a declarative configuration of the StatusDetails type for use
// with apply.
type StatusDetailsApplyConfiguration struct {
	Name              *string                          `json:"name,omitempty"`
	Group             *string                          `json:"group,omitempty"`
	Kind              *string                          `json:"kind,omitempty"`
	UID               *types.UID                       `json:"uid,omitempty"`
	Causes            *[]StatusCauseApplyConfiguration `json:"causes,omitempty"`
	RetryAfterSeconds *int32                           `json:"retryAfterSeconds,omitempty"`
}

// StatusDetailsApplyConfiguration represents a declarative configuration of the StatusDetails type for use
// with apply.
func StatusDetails() *StatusDetailsApplyConfiguration {
	return &StatusDetailsApplyConfiguration{}
}

// TableApplyConfiguration represents a declarative configuration of the Table type for use
// with apply.
type TableApplyConfiguration struct {
	TypeMetaApplyConfiguration  `json:",inline"`
	*ListMetaApplyConfiguration `json:"metadata,omitempty"`
	ColumnDefinitions           *[]TableColumnDefinitionApplyConfiguration `json:"columnDefinitions,omitempty"`
	Rows                        *[]TableRowApplyConfiguration              `json:"rows,omitempty"`
}

// TableApplyConfiguration represents a declarative configuration of the Table type for use
// with apply.
func Table() *TableApplyConfiguration {
	return &TableApplyConfiguration{}
}

// TableColumnDefinitionApplyConfiguration represents a declarative configuration of the TableColumnDefinition type for use
// with apply.
type TableColumnDefinitionApplyConfiguration struct {
	Name        *string `json:"name,omitempty"`
	Type        *string `json:"type,omitempty"`
	Format      *string `json:"format,omitempty"`
	Description *string `json:"description,omitempty"`
	Priority    *int32  `json:"priority,omitempty"`
}

// TableColumnDefinitionApplyConfiguration represents a declarative configuration of the TableColumnDefinition type for use
// with apply.
func TableColumnDefinition() *TableColumnDefinitionApplyConfiguration {
	return &TableColumnDefinitionApplyConfiguration{}
}

// TableOptionsApplyConfiguration represents a declarative configuration of the TableOptions type for use
// with apply.
type TableOptionsApplyConfiguration struct {
	TypeMetaApplyConfiguration `json:",inline"`
	IncludeObject              *metav1.IncludeObjectPolicy `json:"includeObject,omitempty"`
}

// TableOptionsApplyConfiguration represents a declarative configuration of the TableOptions type for use
// with apply.
func TableOptions() *TableOptionsApplyConfiguration {
	return &TableOptionsApplyConfiguration{}
}

// TableRowApplyConfiguration represents a declarative configuration of the TableRow type for use
// with apply.
type TableRowApplyConfiguration struct {
	Cells      *[]interface{}                         `json:"cells,omitempty"`
	Conditions *[]TableRowConditionApplyConfiguration `json:"conditions,omitempty"`
	Object     *runtime.RawExtension                  `json:"object,omitempty"`
}

// TableRowApplyConfiguration represents a declarative configuration of the TableRow type for use
// with apply.
func TableRow() *TableRowApplyConfiguration {
	return &TableRowApplyConfiguration{}
}

// TableRowConditionApplyConfiguration represents a declarative configuration of the TableRowCondition type for use
// with apply.
type TableRowConditionApplyConfiguration struct {
	Type    *metav1.RowConditionType `json:"type,omitempty"`
	Status  *metav1.ConditionStatus  `json:"status,omitempty"`
	Reason  *string                  `json:"reason,omitempty"`
	Message *string                  `json:"message,omitempty"`
}

// TableRowConditionApplyConfiguration represents a declarative configuration of the TableRowCondition type for use
// with apply.
func TableRowCondition() *TableRowConditionApplyConfiguration {
	return &TableRowConditionApplyConfiguration{}
}

// TimestampApplyConfiguration represents a declarative configuration of the Timestamp type for use
// with apply.
type TimestampApplyConfiguration struct {
	Seconds *int64 `json:"seconds,omitempty"`
	Nanos   *int32 `json:"nanos,omitempty"`
}

// TimestampApplyConfiguration represents a declarative configuration of the Timestamp type for use
// with apply.
func Timestamp() *TimestampApplyConfiguration {
	return &TimestampApplyConfiguration{}
}

// TypeMetaApplyConfiguration represents a declarative configuration of the TypeMeta type for use
// with apply.
type TypeMetaApplyConfiguration struct {
	Kind       *string `json:"kind,omitempty"`
	APIVersion *string `json:"apiVersion,omitempty"`
}

// TypeMetaApplyConfiguration represents a declarative configuration of the TypeMeta type for use
// with apply.
func TypeMeta() *TypeMetaApplyConfiguration {
	return &TypeMetaApplyConfiguration{}
}

// UpdateOptionsApplyConfiguration represents a declarative configuration of the UpdateOptions type for use
// with apply.
type UpdateOptionsApplyConfiguration struct {
	TypeMetaApplyConfiguration `json:",inline"`
	DryRun                     *[]string `json:"dryRun,omitempty"`
	FieldManager               *string   `json:"fieldManager,omitempty"`
}

// UpdateOptionsApplyConfiguration represents a declarative configuration of the UpdateOptions type for use
// with apply.
func UpdateOptions() *UpdateOptionsApplyConfiguration {
	return &UpdateOptionsApplyConfiguration{}
}

// WatchEventApplyConfiguration represents a declarative configuration of the WatchEvent type for use
// with apply.
type WatchEventApplyConfiguration struct {
	Type   *string               `json:"type,omitempty"`
	Object *runtime.RawExtension `json:"object,omitempty"`
}

// WatchEventApplyConfiguration represents a declarative configuration of the WatchEvent type for use
// with apply.
func WatchEvent() *WatchEventApplyConfiguration {
	return &WatchEventApplyConfiguration{}
}
