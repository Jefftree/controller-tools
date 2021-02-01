// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	batchv1beta1 "k8s.io/api/batch/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	batchv1 "sigs.k8s.io/controller-tools/pkg/applyconfigurations/testdata/ac/batch/v1"
	corev1 "sigs.k8s.io/controller-tools/pkg/applyconfigurations/testdata/ac/core/v1"
	metav1 "sigs.k8s.io/controller-tools/pkg/applyconfigurations/testdata/ac/meta/v1"
)

// CronJobApplyConfiguration represents a declarative configuration of the CronJob type for use
// with apply.
type CronJobApplyConfiguration struct {
	v1.TypeMeta   `json:",inline"`
	v1.ObjectMeta `json:"metadata,omitempty"`
	Spec          *CronJobSpecApplyConfiguration   `json:"spec,omitempty"`
	Status        *CronJobStatusApplyConfiguration `json:"status,omitempty"`
}

// CronJobApplyConfiguration represents a declarative configuration of the CronJob type for use
// with apply.
func CronJob() *CronJobApplyConfiguration {
	return &CronJobApplyConfiguration{}
}

// CronJobListApplyConfiguration represents a declarative configuration of the CronJobList type for use
// with apply.
type CronJobListApplyConfiguration struct {
	v1.TypeMeta                       `json:",inline"`
	metav1.ListMetaApplyConfiguration `json:"metadata,omitempty"`
	Items                             *[]CronJobApplyConfiguration `json:"items,omitempty"`
}

// CronJobListApplyConfiguration represents a declarative configuration of the CronJobList type for use
// with apply.
func CronJobList() *CronJobListApplyConfiguration {
	return &CronJobListApplyConfiguration{}
}

// CronJobSpecApplyConfiguration represents a declarative configuration of the CronJobSpec type for use
// with apply.
type CronJobSpecApplyConfiguration struct {
	Schedule                   *string                            `json:"schedule,omitempty"`
	StartingDeadlineSeconds    *int64                             `json:"startingDeadlineSeconds,omitempty"`
	ConcurrencyPolicy          *batchv1beta1.ConcurrencyPolicy    `json:"concurrencyPolicy,omitempty"`
	Suspend                    *bool                              `json:"suspend,omitempty"`
	JobTemplate                *JobTemplateSpecApplyConfiguration `json:"jobTemplate,omitempty"`
	SuccessfulJobsHistoryLimit *int32                             `json:"successfulJobsHistoryLimit,omitempty"`
	FailedJobsHistoryLimit     *int32                             `json:"failedJobsHistoryLimit,omitempty"`
}

// CronJobSpecApplyConfiguration represents a declarative configuration of the CronJobSpec type for use
// with apply.
func CronJobSpec() *CronJobSpecApplyConfiguration {
	return &CronJobSpecApplyConfiguration{}
}

// CronJobStatusApplyConfiguration represents a declarative configuration of the CronJobStatus type for use
// with apply.
type CronJobStatusApplyConfiguration struct {
	Active           *[]corev1.ObjectReferenceApplyConfiguration `json:"active,omitempty"`
	LastScheduleTime *v1.Time                                    `json:"lastScheduleTime,omitempty"`
}

// CronJobStatusApplyConfiguration represents a declarative configuration of the CronJobStatus type for use
// with apply.
func CronJobStatus() *CronJobStatusApplyConfiguration {
	return &CronJobStatusApplyConfiguration{}
}

// JobTemplateApplyConfiguration represents a declarative configuration of the JobTemplate type for use
// with apply.
type JobTemplateApplyConfiguration struct {
	v1.TypeMeta   `json:",inline"`
	v1.ObjectMeta `json:"metadata,omitempty"`
	Template      *JobTemplateSpecApplyConfiguration `json:"template,omitempty"`
}

// JobTemplateApplyConfiguration represents a declarative configuration of the JobTemplate type for use
// with apply.
func JobTemplate() *JobTemplateApplyConfiguration {
	return &JobTemplateApplyConfiguration{}
}

// JobTemplateSpecApplyConfiguration represents a declarative configuration of the JobTemplateSpec type for use
// with apply.
type JobTemplateSpecApplyConfiguration struct {
	v1.ObjectMeta `json:"metadata,omitempty"`
	Spec          *batchv1.JobSpecApplyConfiguration `json:"spec,omitempty"`
}

// JobTemplateSpecApplyConfiguration represents a declarative configuration of the JobTemplateSpec type for use
// with apply.
func JobTemplateSpec() *JobTemplateSpecApplyConfiguration {
	return &JobTemplateSpecApplyConfiguration{}
}