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

type IpGroupObservation struct {
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type IpGroupParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Rules []RulesParameters `json:"rules,omitempty" tf:"rules,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type RulesObservation struct {
}

type RulesParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Source *string `json:"source" tf:"source,omitempty"`
}

// IpGroupSpec defines the desired state of IpGroup
type IpGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IpGroupParameters `json:"forProvider"`
}

// IpGroupStatus defines the observed state of IpGroup.
type IpGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IpGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IpGroup is the Schema for the IpGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type IpGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IpGroupSpec   `json:"spec"`
	Status            IpGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IpGroupList contains a list of IpGroups
type IpGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IpGroup `json:"items"`
}

// Repository type metadata.
var (
	IpGroupKind             = "IpGroup"
	IpGroupGroupKind        = schema.GroupKind{Group: Group, Kind: IpGroupKind}.String()
	IpGroupKindAPIVersion   = IpGroupKind + "." + GroupVersion.String()
	IpGroupGroupVersionKind = GroupVersion.WithKind(IpGroupKind)
)

func init() {
	SchemeBuilder.Register(&IpGroup{}, &IpGroupList{})
}
