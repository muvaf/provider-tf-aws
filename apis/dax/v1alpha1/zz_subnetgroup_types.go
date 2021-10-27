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

type SubnetGroupObservation struct {
	VpcID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type SubnetGroupParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	SubnetIds []*string `json:"subnetIds" tf:"subnet_ids,omitempty"`
}

// SubnetGroupSpec defines the desired state of SubnetGroup
type SubnetGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SubnetGroupParameters `json:"forProvider"`
}

// SubnetGroupStatus defines the observed state of SubnetGroup.
type SubnetGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SubnetGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SubnetGroup is the Schema for the SubnetGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type SubnetGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SubnetGroupSpec   `json:"spec"`
	Status            SubnetGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SubnetGroupList contains a list of SubnetGroups
type SubnetGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SubnetGroup `json:"items"`
}

// Repository type metadata.
var (
	SubnetGroupKind             = "SubnetGroup"
	SubnetGroupGroupKind        = schema.GroupKind{Group: Group, Kind: SubnetGroupKind}.String()
	SubnetGroupKindAPIVersion   = SubnetGroupKind + "." + GroupVersion.String()
	SubnetGroupGroupVersionKind = GroupVersion.WithKind(SubnetGroupKind)
)

func init() {
	SchemeBuilder.Register(&SubnetGroup{}, &SubnetGroupList{})
}
