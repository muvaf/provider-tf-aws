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

type ProtectionObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type ProtectionParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	ResourceArn *string `json:"resourceArn" tf:"resource_arn,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// ProtectionSpec defines the desired state of Protection
type ProtectionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProtectionParameters `json:"forProvider"`
}

// ProtectionStatus defines the observed state of Protection.
type ProtectionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProtectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Protection is the Schema for the Protections API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type Protection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProtectionSpec   `json:"spec"`
	Status            ProtectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProtectionList contains a list of Protections
type ProtectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Protection `json:"items"`
}

// Repository type metadata.
var (
	Protection_Kind             = "Protection"
	Protection_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Protection_Kind}.String()
	Protection_KindAPIVersion   = Protection_Kind + "." + CRDGroupVersion.String()
	Protection_GroupVersionKind = CRDGroupVersion.WithKind(Protection_Kind)
)

func init() {
	SchemeBuilder.Register(&Protection{}, &ProtectionList{})
}
