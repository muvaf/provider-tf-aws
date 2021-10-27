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

type OrganizationalUnitAccountsObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type OrganizationalUnitAccountsParameters struct {
}

type OrganizationalUnitObservation struct {
	Accounts []OrganizationalUnitAccountsObservation `json:"accounts,omitempty" tf:"accounts,omitempty"`

	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type OrganizationalUnitParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	ParentID *string `json:"parentId" tf:"parent_id,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// OrganizationalUnitSpec defines the desired state of OrganizationalUnit
type OrganizationalUnitSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OrganizationalUnitParameters `json:"forProvider"`
}

// OrganizationalUnitStatus defines the observed state of OrganizationalUnit.
type OrganizationalUnitStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OrganizationalUnitObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// OrganizationalUnit is the Schema for the OrganizationalUnits API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type OrganizationalUnit struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OrganizationalUnitSpec   `json:"spec"`
	Status            OrganizationalUnitStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OrganizationalUnitList contains a list of OrganizationalUnits
type OrganizationalUnitList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OrganizationalUnit `json:"items"`
}

// Repository type metadata.
var (
	OrganizationalUnitKind             = "OrganizationalUnit"
	OrganizationalUnitGroupKind        = schema.GroupKind{Group: Group, Kind: OrganizationalUnitKind}.String()
	OrganizationalUnitKindAPIVersion   = OrganizationalUnitKind + "." + GroupVersion.String()
	OrganizationalUnitGroupVersionKind = GroupVersion.WithKind(OrganizationalUnitKind)
)

func init() {
	SchemeBuilder.Register(&OrganizationalUnit{}, &OrganizationalUnitList{})
}
