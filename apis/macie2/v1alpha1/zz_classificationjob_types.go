/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AndObservation struct {
}

type AndParameters struct {

	// +kubebuilder:validation:Optional
	SimpleScopeTerm []SimpleScopeTermParameters `json:"simpleScopeTerm,omitempty" tf:"simple_scope_term,omitempty"`

	// +kubebuilder:validation:Optional
	TagScopeTerm []TagScopeTermParameters `json:"tagScopeTerm,omitempty" tf:"tag_scope_term,omitempty"`
}

type AndSimpleScopeTermObservation struct {
}

type AndSimpleScopeTermParameters struct {

	// +kubebuilder:validation:Optional
	Comparator *string `json:"comparator,omitempty" tf:"comparator,omitempty"`

	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// +kubebuilder:validation:Optional
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type AndTagScopeTermObservation struct {
}

type AndTagScopeTermParameters struct {

	// +kubebuilder:validation:Optional
	Comparator *string `json:"comparator,omitempty" tf:"comparator,omitempty"`

	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// +kubebuilder:validation:Optional
	TagValues []TagScopeTermTagValuesParameters `json:"tagValues,omitempty" tf:"tag_values,omitempty"`

	// +kubebuilder:validation:Optional
	Target *string `json:"target,omitempty" tf:"target,omitempty"`
}

type BucketDefinitionsObservation struct {
}

type BucketDefinitionsParameters struct {

	// +kubebuilder:validation:Required
	AccountID *string `json:"accountId" tf:"account_id,omitempty"`

	// +kubebuilder:validation:Required
	Buckets []*string `json:"buckets" tf:"buckets,omitempty"`
}

type ClassificationJobObservation struct {
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	JobArn *string `json:"jobArn,omitempty" tf:"job_arn,omitempty"`

	JobID *string `json:"jobId,omitempty" tf:"job_id,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	UserPausedDetails []UserPausedDetailsObservation `json:"userPausedDetails,omitempty" tf:"user_paused_details,omitempty"`
}

type ClassificationJobParameters struct {

	// +kubebuilder:validation:Optional
	CustomDataIdentifierIds []*string `json:"customDataIdentifierIds,omitempty" tf:"custom_data_identifier_ids,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	InitialRun *bool `json:"initialRun,omitempty" tf:"initial_run,omitempty"`

	// +kubebuilder:validation:Optional
	JobStatus *string `json:"jobStatus,omitempty" tf:"job_status,omitempty"`

	// +kubebuilder:validation:Required
	JobType *string `json:"jobType" tf:"job_type,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	NamePrefix *string `json:"namePrefix,omitempty" tf:"name_prefix,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	S3JobDefinition []S3JobDefinitionParameters `json:"s3JobDefinition" tf:"s3_job_definition,omitempty"`

	// +kubebuilder:validation:Optional
	SamplingPercentage *int64 `json:"samplingPercentage,omitempty" tf:"sampling_percentage,omitempty"`

	// +kubebuilder:validation:Optional
	ScheduleFrequency []ScheduleFrequencyParameters `json:"scheduleFrequency,omitempty" tf:"schedule_frequency,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ExcludesObservation struct {
}

type ExcludesParameters struct {

	// +kubebuilder:validation:Optional
	And []AndParameters `json:"and,omitempty" tf:"and,omitempty"`
}

type IncludesAndObservation struct {
}

type IncludesAndParameters struct {

	// +kubebuilder:validation:Optional
	SimpleScopeTerm []AndSimpleScopeTermParameters `json:"simpleScopeTerm,omitempty" tf:"simple_scope_term,omitempty"`

	// +kubebuilder:validation:Optional
	TagScopeTerm []AndTagScopeTermParameters `json:"tagScopeTerm,omitempty" tf:"tag_scope_term,omitempty"`
}

type IncludesObservation struct {
}

type IncludesParameters struct {

	// +kubebuilder:validation:Optional
	And []IncludesAndParameters `json:"and,omitempty" tf:"and,omitempty"`
}

type S3JobDefinitionObservation struct {
}

type S3JobDefinitionParameters struct {

	// +kubebuilder:validation:Optional
	BucketDefinitions []BucketDefinitionsParameters `json:"bucketDefinitions,omitempty" tf:"bucket_definitions,omitempty"`

	// +kubebuilder:validation:Optional
	Scoping []ScopingParameters `json:"scoping,omitempty" tf:"scoping,omitempty"`
}

type ScheduleFrequencyObservation struct {
}

type ScheduleFrequencyParameters struct {

	// +kubebuilder:validation:Optional
	DailySchedule *bool `json:"dailySchedule,omitempty" tf:"daily_schedule,omitempty"`

	// +kubebuilder:validation:Optional
	MonthlySchedule *int64 `json:"monthlySchedule,omitempty" tf:"monthly_schedule,omitempty"`

	// +kubebuilder:validation:Optional
	WeeklySchedule *string `json:"weeklySchedule,omitempty" tf:"weekly_schedule,omitempty"`
}

type ScopingObservation struct {
}

type ScopingParameters struct {

	// +kubebuilder:validation:Optional
	Excludes []ExcludesParameters `json:"excludes,omitempty" tf:"excludes,omitempty"`

	// +kubebuilder:validation:Optional
	Includes []IncludesParameters `json:"includes,omitempty" tf:"includes,omitempty"`
}

type SimpleScopeTermObservation struct {
}

type SimpleScopeTermParameters struct {

	// +kubebuilder:validation:Optional
	Comparator *string `json:"comparator,omitempty" tf:"comparator,omitempty"`

	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// +kubebuilder:validation:Optional
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type TagScopeTermObservation struct {
}

type TagScopeTermParameters struct {

	// +kubebuilder:validation:Optional
	Comparator *string `json:"comparator,omitempty" tf:"comparator,omitempty"`

	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// +kubebuilder:validation:Optional
	TagValues []TagValuesParameters `json:"tagValues,omitempty" tf:"tag_values,omitempty"`

	// +kubebuilder:validation:Optional
	Target *string `json:"target,omitempty" tf:"target,omitempty"`
}

type TagScopeTermTagValuesObservation struct {
}

type TagScopeTermTagValuesParameters struct {

	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type TagValuesObservation struct {
}

type TagValuesParameters struct {

	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type UserPausedDetailsObservation struct {
	JobExpiresAt *string `json:"jobExpiresAt,omitempty" tf:"job_expires_at,omitempty"`

	JobImminentExpirationHealthEventArn *string `json:"jobImminentExpirationHealthEventArn,omitempty" tf:"job_imminent_expiration_health_event_arn,omitempty"`

	JobPausedAt *string `json:"jobPausedAt,omitempty" tf:"job_paused_at,omitempty"`
}

type UserPausedDetailsParameters struct {
}

// ClassificationJobSpec defines the desired state of ClassificationJob
type ClassificationJobSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClassificationJobParameters `json:"forProvider"`
}

// ClassificationJobStatus defines the observed state of ClassificationJob.
type ClassificationJobStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClassificationJobObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ClassificationJob is the Schema for the ClassificationJobs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type ClassificationJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClassificationJobSpec   `json:"spec"`
	Status            ClassificationJobStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClassificationJobList contains a list of ClassificationJobs
type ClassificationJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClassificationJob `json:"items"`
}

// Repository type metadata.
var (
	ClassificationJobKind             = "ClassificationJob"
	ClassificationJobGroupKind        = schema.GroupKind{Group: Group, Kind: ClassificationJobKind}.String()
	ClassificationJobKindAPIVersion   = ClassificationJobKind + "." + GroupVersion.String()
	ClassificationJobGroupVersionKind = GroupVersion.WithKind(ClassificationJobKind)
)

func init() {
	SchemeBuilder.Register(&ClassificationJob{}, &ClassificationJobList{})
}
