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

type EipAssociationObservation struct {
}

type EipAssociationParameters struct {

	// +kubebuilder:validation:Optional
	AllocationID *string `json:"allocationId,omitempty" tf:"allocation_id,omitempty"`

	// +kubebuilder:validation:Optional
	AllowReassociation *bool `json:"allowReassociation,omitempty" tf:"allow_reassociation,omitempty"`

	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// +kubebuilder:validation:Optional
	NetworkInterfaceID *string `json:"networkInterfaceId,omitempty" tf:"network_interface_id,omitempty"`

	// +kubebuilder:validation:Optional
	PrivateIPAddress *string `json:"privateIpAddress,omitempty" tf:"private_ip_address,omitempty"`

	// +kubebuilder:validation:Optional
	PublicIP *string `json:"publicIp,omitempty" tf:"public_ip,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// EipAssociationSpec defines the desired state of EipAssociation
type EipAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EipAssociationParameters `json:"forProvider"`
}

// EipAssociationStatus defines the observed state of EipAssociation.
type EipAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EipAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EipAssociation is the Schema for the EipAssociations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type EipAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EipAssociationSpec   `json:"spec"`
	Status            EipAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EipAssociationList contains a list of EipAssociations
type EipAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EipAssociation `json:"items"`
}

// Repository type metadata.
var (
	EipAssociationKind             = "EipAssociation"
	EipAssociationGroupKind        = schema.GroupKind{Group: Group, Kind: EipAssociationKind}.String()
	EipAssociationKindAPIVersion   = EipAssociationKind + "." + GroupVersion.String()
	EipAssociationGroupVersionKind = GroupVersion.WithKind(EipAssociationKind)
)

func init() {
	SchemeBuilder.Register(&EipAssociation{}, &EipAssociationList{})
}
