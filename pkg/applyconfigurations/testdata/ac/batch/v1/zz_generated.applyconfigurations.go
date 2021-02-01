// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	batchv1 "k8s.io/api/batch/v1"
	apicorev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1 "sigs.k8s.io/controller-tools/pkg/applyconfigurations/testdata/ac/core/v1"
	acmetav1 "sigs.k8s.io/controller-tools/pkg/applyconfigurations/testdata/ac/meta/v1"
)

// JobApplyConfiguration represents a declarative configuration of the Job type for use
// with apply.
type JobApplyConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              *JobSpecApplyConfiguration   `json:"spec,omitempty"`
	Status            *JobStatusApplyConfiguration `json:"status,omitempty"`
}

// JobApplyConfiguration represents a declarative configuration of the Job type for use
// with apply.
func Job() *JobApplyConfiguration {
	return &JobApplyConfiguration{}
}

// JobConditionApplyConfiguration represents a declarative configuration of the JobCondition type for use
// with apply.
type JobConditionApplyConfiguration struct {
	Type               *batchv1.JobConditionType  `json:"type,omitempty"`
	Status             *apicorev1.ConditionStatus `json:"status,omitempty"`
	LastProbeTime      *metav1.Time               `json:"lastProbeTime,omitempty"`
	LastTransitionTime *metav1.Time               `json:"lastTransitionTime,omitempty"`
	Reason             *string                    `json:"reason,omitempty"`
	Message            *string                    `json:"message,omitempty"`
}

// JobConditionApplyConfiguration represents a declarative configuration of the JobCondition type for use
// with apply.
func JobCondition() *JobConditionApplyConfiguration {
	return &JobConditionApplyConfiguration{}
}

// JobListApplyConfiguration represents a declarative configuration of the JobList type for use
// with apply.
type JobListApplyConfiguration struct {
	metav1.TypeMeta                     `json:",inline"`
	acmetav1.ListMetaApplyConfiguration `json:"metadata,omitempty"`
	Items                               *[]JobApplyConfiguration `json:"items,omitempty"`
}

// JobListApplyConfiguration represents a declarative configuration of the JobList type for use
// with apply.
func JobList() *JobListApplyConfiguration {
	return &JobListApplyConfiguration{}
}

// JobSpecApplyConfiguration represents a declarative configuration of the JobSpec type for use
// with apply.
type JobSpecApplyConfiguration struct {
	Parallelism             *int32                                    `json:"parallelism,omitempty"`
	Completions             *int32                                    `json:"completions,omitempty"`
	ActiveDeadlineSeconds   *int64                                    `json:"activeDeadlineSeconds,omitempty"`
	BackoffLimit            *int32                                    `json:"backoffLimit,omitempty"`
	Selector                *acmetav1.LabelSelectorApplyConfiguration `json:"selector,omitempty"`
	ManualSelector          *bool                                     `json:"manualSelector,omitempty"`
	Template                *corev1.PodTemplateSpecApplyConfiguration `json:"template,omitempty"`
	TTLSecondsAfterFinished *int32                                    `json:"ttlSecondsAfterFinished,omitempty"`
}

// JobSpecApplyConfiguration represents a declarative configuration of the JobSpec type for use
// with apply.
func JobSpec() *JobSpecApplyConfiguration {
	return &JobSpecApplyConfiguration{}
}

// JobStatusApplyConfiguration represents a declarative configuration of the JobStatus type for use
// with apply.
type JobStatusApplyConfiguration struct {
	Conditions     *[]JobConditionApplyConfiguration `json:"conditions,omitempty"`
	StartTime      *metav1.Time                      `json:"startTime,omitempty"`
	CompletionTime *metav1.Time                      `json:"completionTime,omitempty"`
	Active         *int32                            `json:"active,omitempty"`
	Succeeded      *int32                            `json:"succeeded,omitempty"`
	Failed         *int32                            `json:"failed,omitempty"`
}

// JobStatusApplyConfiguration represents a declarative configuration of the JobStatus type for use
// with apply.
func JobStatus() *JobStatusApplyConfiguration {
	return &JobStatusApplyConfiguration{}
}