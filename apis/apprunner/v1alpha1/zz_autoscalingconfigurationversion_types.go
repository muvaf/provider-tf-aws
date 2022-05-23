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

type AutoScalingConfigurationVersionObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	AutoScalingConfigurationRevision *float64 `json:"autoScalingConfigurationRevision,omitempty" tf:"auto_scaling_configuration_revision,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Latest *bool `json:"latest,omitempty" tf:"latest,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type AutoScalingConfigurationVersionParameters struct {

	// +kubebuilder:validation:Required
	AutoScalingConfigurationName *string `json:"autoScalingConfigurationName" tf:"auto_scaling_configuration_name,omitempty"`

	// +kubebuilder:validation:Optional
	MaxConcurrency *float64 `json:"maxConcurrency,omitempty" tf:"max_concurrency,omitempty"`

	// +kubebuilder:validation:Optional
	MaxSize *float64 `json:"maxSize,omitempty" tf:"max_size,omitempty"`

	// +kubebuilder:validation:Optional
	MinSize *float64 `json:"minSize,omitempty" tf:"min_size,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// AutoScalingConfigurationVersionSpec defines the desired state of AutoScalingConfigurationVersion
type AutoScalingConfigurationVersionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AutoScalingConfigurationVersionParameters `json:"forProvider"`
}

// AutoScalingConfigurationVersionStatus defines the observed state of AutoScalingConfigurationVersion.
type AutoScalingConfigurationVersionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AutoScalingConfigurationVersionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AutoScalingConfigurationVersion is the Schema for the AutoScalingConfigurationVersions API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type AutoScalingConfigurationVersion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AutoScalingConfigurationVersionSpec   `json:"spec"`
	Status            AutoScalingConfigurationVersionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AutoScalingConfigurationVersionList contains a list of AutoScalingConfigurationVersions
type AutoScalingConfigurationVersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AutoScalingConfigurationVersion `json:"items"`
}

// Repository type metadata.
var (
	AutoScalingConfigurationVersion_Kind             = "AutoScalingConfigurationVersion"
	AutoScalingConfigurationVersion_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AutoScalingConfigurationVersion_Kind}.String()
	AutoScalingConfigurationVersion_KindAPIVersion   = AutoScalingConfigurationVersion_Kind + "." + CRDGroupVersion.String()
	AutoScalingConfigurationVersion_GroupVersionKind = CRDGroupVersion.WithKind(AutoScalingConfigurationVersion_Kind)
)

func init() {
	SchemeBuilder.Register(&AutoScalingConfigurationVersion{}, &AutoScalingConfigurationVersionList{})
}
