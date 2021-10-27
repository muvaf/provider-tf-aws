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

type ListenerObservation struct {
}

type ListenerParameters struct {

	// +kubebuilder:validation:Required
	AcceleratorArn *string `json:"acceleratorArn" tf:"accelerator_arn,omitempty"`

	// +kubebuilder:validation:Optional
	ClientAffinity *string `json:"clientAffinity,omitempty" tf:"client_affinity,omitempty"`

	// +kubebuilder:validation:Required
	PortRange []PortRangeParameters `json:"portRange" tf:"port_range,omitempty"`

	// +kubebuilder:validation:Required
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

type PortRangeObservation struct {
}

type PortRangeParameters struct {

	// +kubebuilder:validation:Optional
	FromPort *int64 `json:"fromPort,omitempty" tf:"from_port,omitempty"`

	// +kubebuilder:validation:Optional
	ToPort *int64 `json:"toPort,omitempty" tf:"to_port,omitempty"`
}

// ListenerSpec defines the desired state of Listener
type ListenerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ListenerParameters `json:"forProvider"`
}

// ListenerStatus defines the observed state of Listener.
type ListenerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ListenerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Listener is the Schema for the Listeners API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type Listener struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ListenerSpec   `json:"spec"`
	Status            ListenerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ListenerList contains a list of Listeners
type ListenerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Listener `json:"items"`
}

// Repository type metadata.
var (
	ListenerKind             = "Listener"
	ListenerGroupKind        = schema.GroupKind{Group: Group, Kind: ListenerKind}.String()
	ListenerKindAPIVersion   = ListenerKind + "." + GroupVersion.String()
	ListenerGroupVersionKind = GroupVersion.WithKind(ListenerKind)
)

func init() {
	SchemeBuilder.Register(&Listener{}, &ListenerList{})
}
