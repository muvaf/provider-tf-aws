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

type NamedQueryObservation struct {
}

type NamedQueryParameters struct {

	// +kubebuilder:validation:Required
	Database *string `json:"database" tf:"database,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Query *string `json:"query" tf:"query,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Workgroup *string `json:"workgroup,omitempty" tf:"workgroup,omitempty"`
}

// NamedQuerySpec defines the desired state of NamedQuery
type NamedQuerySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NamedQueryParameters `json:"forProvider"`
}

// NamedQueryStatus defines the observed state of NamedQuery.
type NamedQueryStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NamedQueryObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NamedQuery is the Schema for the NamedQuerys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type NamedQuery struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NamedQuerySpec   `json:"spec"`
	Status            NamedQueryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NamedQueryList contains a list of NamedQuerys
type NamedQueryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NamedQuery `json:"items"`
}

// Repository type metadata.
var (
	NamedQueryKind             = "NamedQuery"
	NamedQueryGroupKind        = schema.GroupKind{Group: Group, Kind: NamedQueryKind}.String()
	NamedQueryKindAPIVersion   = NamedQueryKind + "." + GroupVersion.String()
	NamedQueryGroupVersionKind = GroupVersion.WithKind(NamedQueryKind)
)

func init() {
	SchemeBuilder.Register(&NamedQuery{}, &NamedQueryList{})
}
