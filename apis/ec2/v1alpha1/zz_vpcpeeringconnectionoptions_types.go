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

type VPCPeeringConnectionOptionsAccepterObservation struct {
}

type VPCPeeringConnectionOptionsAccepterParameters struct {

	// +kubebuilder:validation:Optional
	AllowClassicLinkToRemoteVPC *bool `json:"allowClassicLinkToRemoteVpc,omitempty" tf:"allow_classic_link_to_remote_vpc,omitempty"`

	// +kubebuilder:validation:Optional
	AllowRemoteVPCDNSResolution *bool `json:"allowRemoteVpcDnsResolution,omitempty" tf:"allow_remote_vpc_dns_resolution,omitempty"`

	// +kubebuilder:validation:Optional
	AllowVPCToRemoteClassicLink *bool `json:"allowVpcToRemoteClassicLink,omitempty" tf:"allow_vpc_to_remote_classic_link,omitempty"`
}

type VPCPeeringConnectionOptionsObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type VPCPeeringConnectionOptionsParameters struct {

	// +kubebuilder:validation:Optional
	Accepter []VPCPeeringConnectionOptionsAccepterParameters `json:"accepter,omitempty" tf:"accepter,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Requester []VPCPeeringConnectionOptionsRequesterParameters `json:"requester,omitempty" tf:"requester,omitempty"`

	// +kubebuilder:validation:Required
	VPCPeeringConnectionID *string `json:"vpcPeeringConnectionId" tf:"vpc_peering_connection_id,omitempty"`
}

type VPCPeeringConnectionOptionsRequesterObservation struct {
}

type VPCPeeringConnectionOptionsRequesterParameters struct {

	// +kubebuilder:validation:Optional
	AllowClassicLinkToRemoteVPC *bool `json:"allowClassicLinkToRemoteVpc,omitempty" tf:"allow_classic_link_to_remote_vpc,omitempty"`

	// +kubebuilder:validation:Optional
	AllowRemoteVPCDNSResolution *bool `json:"allowRemoteVpcDnsResolution,omitempty" tf:"allow_remote_vpc_dns_resolution,omitempty"`

	// +kubebuilder:validation:Optional
	AllowVPCToRemoteClassicLink *bool `json:"allowVpcToRemoteClassicLink,omitempty" tf:"allow_vpc_to_remote_classic_link,omitempty"`
}

// VPCPeeringConnectionOptionsSpec defines the desired state of VPCPeeringConnectionOptions
type VPCPeeringConnectionOptionsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VPCPeeringConnectionOptionsParameters `json:"forProvider"`
}

// VPCPeeringConnectionOptionsStatus defines the observed state of VPCPeeringConnectionOptions.
type VPCPeeringConnectionOptionsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VPCPeeringConnectionOptionsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VPCPeeringConnectionOptions is the Schema for the VPCPeeringConnectionOptionss API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type VPCPeeringConnectionOptions struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VPCPeeringConnectionOptionsSpec   `json:"spec"`
	Status            VPCPeeringConnectionOptionsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VPCPeeringConnectionOptionsList contains a list of VPCPeeringConnectionOptionss
type VPCPeeringConnectionOptionsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VPCPeeringConnectionOptions `json:"items"`
}

// Repository type metadata.
var (
	VPCPeeringConnectionOptions_Kind             = "VPCPeeringConnectionOptions"
	VPCPeeringConnectionOptions_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VPCPeeringConnectionOptions_Kind}.String()
	VPCPeeringConnectionOptions_KindAPIVersion   = VPCPeeringConnectionOptions_Kind + "." + CRDGroupVersion.String()
	VPCPeeringConnectionOptions_GroupVersionKind = CRDGroupVersion.WithKind(VPCPeeringConnectionOptions_Kind)
)

func init() {
	SchemeBuilder.Register(&VPCPeeringConnectionOptions{}, &VPCPeeringConnectionOptionsList{})
}
