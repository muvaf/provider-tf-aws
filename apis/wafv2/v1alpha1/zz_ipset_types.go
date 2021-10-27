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

type IpSetObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	LockToken *string `json:"lockToken,omitempty" tf:"lock_token,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type IpSetParameters struct {

	// +kubebuilder:validation:Optional
	Addresses []*string `json:"addresses,omitempty" tf:"addresses,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	IPAddressVersion *string `json:"ipAddressVersion" tf:"ip_address_version,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	Scope *string `json:"scope" tf:"scope,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// IpSetSpec defines the desired state of IpSet
type IpSetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IpSetParameters `json:"forProvider"`
}

// IpSetStatus defines the observed state of IpSet.
type IpSetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IpSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IpSet is the Schema for the IpSets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type IpSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IpSetSpec   `json:"spec"`
	Status            IpSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IpSetList contains a list of IpSets
type IpSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IpSet `json:"items"`
}

// Repository type metadata.
var (
	IpSetKind             = "IpSet"
	IpSetGroupKind        = schema.GroupKind{Group: Group, Kind: IpSetKind}.String()
	IpSetKindAPIVersion   = IpSetKind + "." + GroupVersion.String()
	IpSetGroupVersionKind = GroupVersion.WithKind(IpSetKind)
)

func init() {
	SchemeBuilder.Register(&IpSet{}, &IpSetList{})
}
