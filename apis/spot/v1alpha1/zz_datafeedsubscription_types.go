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

type DatafeedSubscriptionObservation struct {
}

type DatafeedSubscriptionParameters struct {

	// +kubebuilder:validation:Required
	Bucket *string `json:"bucket" tf:"bucket,omitempty"`

	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// DatafeedSubscriptionSpec defines the desired state of DatafeedSubscription
type DatafeedSubscriptionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DatafeedSubscriptionParameters `json:"forProvider"`
}

// DatafeedSubscriptionStatus defines the observed state of DatafeedSubscription.
type DatafeedSubscriptionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DatafeedSubscriptionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DatafeedSubscription is the Schema for the DatafeedSubscriptions API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type DatafeedSubscription struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatafeedSubscriptionSpec   `json:"spec"`
	Status            DatafeedSubscriptionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DatafeedSubscriptionList contains a list of DatafeedSubscriptions
type DatafeedSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DatafeedSubscription `json:"items"`
}

// Repository type metadata.
var (
	DatafeedSubscriptionKind             = "DatafeedSubscription"
	DatafeedSubscriptionGroupKind        = schema.GroupKind{Group: Group, Kind: DatafeedSubscriptionKind}.String()
	DatafeedSubscriptionKindAPIVersion   = DatafeedSubscriptionKind + "." + GroupVersion.String()
	DatafeedSubscriptionGroupVersionKind = GroupVersion.WithKind(DatafeedSubscriptionKind)
)

func init() {
	SchemeBuilder.Register(&DatafeedSubscription{}, &DatafeedSubscriptionList{})
}
