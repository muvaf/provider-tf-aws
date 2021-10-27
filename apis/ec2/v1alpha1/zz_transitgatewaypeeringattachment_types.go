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

type TransitGatewayPeeringAttachmentObservation struct {
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type TransitGatewayPeeringAttachmentParameters struct {

	// +kubebuilder:validation:Optional
	PeerAccountID *string `json:"peerAccountId,omitempty" tf:"peer_account_id,omitempty"`

	// +kubebuilder:validation:Required
	PeerRegion *string `json:"peerRegion" tf:"peer_region,omitempty"`

	// +kubebuilder:validation:Required
	PeerTransitGatewayID *string `json:"peerTransitGatewayId" tf:"peer_transit_gateway_id,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Required
	TransitGatewayID *string `json:"transitGatewayId" tf:"transit_gateway_id,omitempty"`
}

// TransitGatewayPeeringAttachmentSpec defines the desired state of TransitGatewayPeeringAttachment
type TransitGatewayPeeringAttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TransitGatewayPeeringAttachmentParameters `json:"forProvider"`
}

// TransitGatewayPeeringAttachmentStatus defines the observed state of TransitGatewayPeeringAttachment.
type TransitGatewayPeeringAttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TransitGatewayPeeringAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TransitGatewayPeeringAttachment is the Schema for the TransitGatewayPeeringAttachments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type TransitGatewayPeeringAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TransitGatewayPeeringAttachmentSpec   `json:"spec"`
	Status            TransitGatewayPeeringAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TransitGatewayPeeringAttachmentList contains a list of TransitGatewayPeeringAttachments
type TransitGatewayPeeringAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TransitGatewayPeeringAttachment `json:"items"`
}

// Repository type metadata.
var (
	TransitGatewayPeeringAttachmentKind             = "TransitGatewayPeeringAttachment"
	TransitGatewayPeeringAttachmentGroupKind        = schema.GroupKind{Group: Group, Kind: TransitGatewayPeeringAttachmentKind}.String()
	TransitGatewayPeeringAttachmentKindAPIVersion   = TransitGatewayPeeringAttachmentKind + "." + GroupVersion.String()
	TransitGatewayPeeringAttachmentGroupVersionKind = GroupVersion.WithKind(TransitGatewayPeeringAttachmentKind)
)

func init() {
	SchemeBuilder.Register(&TransitGatewayPeeringAttachment{}, &TransitGatewayPeeringAttachmentList{})
}
