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

type ThingObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	DefaultClientID *string `json:"defaultClientId,omitempty" tf:"default_client_id,omitempty"`

	Version *int64 `json:"version,omitempty" tf:"version,omitempty"`
}

type ThingParameters struct {

	// +kubebuilder:validation:Optional
	Attributes map[string]*string `json:"attributes,omitempty" tf:"attributes,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	ThingTypeName *string `json:"thingTypeName,omitempty" tf:"thing_type_name,omitempty"`
}

// ThingSpec defines the desired state of Thing
type ThingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ThingParameters `json:"forProvider"`
}

// ThingStatus defines the observed state of Thing.
type ThingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ThingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Thing is the Schema for the Things API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type Thing struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ThingSpec   `json:"spec"`
	Status            ThingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ThingList contains a list of Things
type ThingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Thing `json:"items"`
}

// Repository type metadata.
var (
	ThingKind             = "Thing"
	ThingGroupKind        = schema.GroupKind{Group: Group, Kind: ThingKind}.String()
	ThingKindAPIVersion   = ThingKind + "." + GroupVersion.String()
	ThingGroupVersionKind = GroupVersion.WithKind(ThingKind)
)

func init() {
	SchemeBuilder.Register(&Thing{}, &ThingList{})
}
