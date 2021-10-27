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

type ImagePipelineImageTestsConfigurationObservation struct {
}

type ImagePipelineImageTestsConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	ImageTestsEnabled *bool `json:"imageTestsEnabled,omitempty" tf:"image_tests_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	TimeoutMinutes *int64 `json:"timeoutMinutes,omitempty" tf:"timeout_minutes,omitempty"`
}

type ImagePipelineObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	DateCreated *string `json:"dateCreated,omitempty" tf:"date_created,omitempty"`

	DateLastRun *string `json:"dateLastRun,omitempty" tf:"date_last_run,omitempty"`

	DateNextRun *string `json:"dateNextRun,omitempty" tf:"date_next_run,omitempty"`

	DateUpdated *string `json:"dateUpdated,omitempty" tf:"date_updated,omitempty"`

	Platform *string `json:"platform,omitempty" tf:"platform,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type ImagePipelineParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	DistributionConfigurationArn *string `json:"distributionConfigurationArn,omitempty" tf:"distribution_configuration_arn,omitempty"`

	// +kubebuilder:validation:Optional
	EnhancedImageMetadataEnabled *bool `json:"enhancedImageMetadataEnabled,omitempty" tf:"enhanced_image_metadata_enabled,omitempty"`

	// +kubebuilder:validation:Required
	ImageRecipeArn *string `json:"imageRecipeArn" tf:"image_recipe_arn,omitempty"`

	// +kubebuilder:validation:Optional
	ImageTestsConfiguration []ImagePipelineImageTestsConfigurationParameters `json:"imageTestsConfiguration,omitempty" tf:"image_tests_configuration,omitempty"`

	// +kubebuilder:validation:Required
	InfrastructureConfigurationArn *string `json:"infrastructureConfigurationArn" tf:"infrastructure_configuration_arn,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Schedule []ScheduleParameters `json:"schedule,omitempty" tf:"schedule,omitempty"`

	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ScheduleObservation struct {
}

type ScheduleParameters struct {

	// +kubebuilder:validation:Optional
	PipelineExecutionStartCondition *string `json:"pipelineExecutionStartCondition,omitempty" tf:"pipeline_execution_start_condition,omitempty"`

	// +kubebuilder:validation:Required
	ScheduleExpression *string `json:"scheduleExpression" tf:"schedule_expression,omitempty"`
}

// ImagePipelineSpec defines the desired state of ImagePipeline
type ImagePipelineSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ImagePipelineParameters `json:"forProvider"`
}

// ImagePipelineStatus defines the observed state of ImagePipeline.
type ImagePipelineStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ImagePipelineObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ImagePipeline is the Schema for the ImagePipelines API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type ImagePipeline struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ImagePipelineSpec   `json:"spec"`
	Status            ImagePipelineStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ImagePipelineList contains a list of ImagePipelines
type ImagePipelineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ImagePipeline `json:"items"`
}

// Repository type metadata.
var (
	ImagePipelineKind             = "ImagePipeline"
	ImagePipelineGroupKind        = schema.GroupKind{Group: Group, Kind: ImagePipelineKind}.String()
	ImagePipelineKindAPIVersion   = ImagePipelineKind + "." + GroupVersion.String()
	ImagePipelineGroupVersionKind = GroupVersion.WithKind(ImagePipelineKind)
)

func init() {
	SchemeBuilder.Register(&ImagePipeline{}, &ImagePipelineList{})
}
