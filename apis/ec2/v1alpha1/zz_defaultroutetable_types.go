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

type DefaultRouteTableObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	OwnerID *string `json:"ownerId,omitempty" tf:"owner_id,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type DefaultRouteTableParameters struct {

	// +kubebuilder:validation:Required
	DefaultRouteTableID *string `json:"defaultRouteTableId" tf:"default_route_table_id,omitempty"`

	// +kubebuilder:validation:Optional
	PropagatingVgws []*string `json:"propagatingVgws,omitempty" tf:"propagating_vgws,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Route []RouteParameters `json:"route,omitempty" tf:"route,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type RouteObservation struct {
}

type RouteParameters struct {

	// +kubebuilder:validation:Optional
	CidrBlock *string `json:"cidrBlock,omitempty" tf:"cidr_block"`

	// +kubebuilder:validation:Optional
	DestinationPrefixListID *string `json:"destinationPrefixListId,omitempty" tf:"destination_prefix_list_id"`

	// +kubebuilder:validation:Optional
	EgressOnlyGatewayID *string `json:"egressOnlyGatewayId,omitempty" tf:"egress_only_gateway_id"`

	// +kubebuilder:validation:Optional
	GatewayID *string `json:"gatewayId,omitempty" tf:"gateway_id"`

	// +kubebuilder:validation:Optional
	IPv6CidrBlock *string `json:"ipv6CidrBlock,omitempty" tf:"ipv6_cidr_block"`

	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id"`

	// +kubebuilder:validation:Optional
	NATGatewayID *string `json:"natGatewayId,omitempty" tf:"nat_gateway_id"`

	// +kubebuilder:validation:Optional
	NetworkInterfaceID *string `json:"networkInterfaceId,omitempty" tf:"network_interface_id"`

	// +kubebuilder:validation:Optional
	TransitGatewayID *string `json:"transitGatewayId,omitempty" tf:"transit_gateway_id"`

	// +kubebuilder:validation:Optional
	VPCEndpointID *string `json:"vpcEndpointId,omitempty" tf:"vpc_endpoint_id"`

	// +kubebuilder:validation:Optional
	VPCPeeringConnectionID *string `json:"vpcPeeringConnectionId,omitempty" tf:"vpc_peering_connection_id"`
}

// DefaultRouteTableSpec defines the desired state of DefaultRouteTable
type DefaultRouteTableSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DefaultRouteTableParameters `json:"forProvider"`
}

// DefaultRouteTableStatus defines the observed state of DefaultRouteTable.
type DefaultRouteTableStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DefaultRouteTableObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DefaultRouteTable is the Schema for the DefaultRouteTables API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type DefaultRouteTable struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DefaultRouteTableSpec   `json:"spec"`
	Status            DefaultRouteTableStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DefaultRouteTableList contains a list of DefaultRouteTables
type DefaultRouteTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DefaultRouteTable `json:"items"`
}

// Repository type metadata.
var (
	DefaultRouteTable_Kind             = "DefaultRouteTable"
	DefaultRouteTable_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DefaultRouteTable_Kind}.String()
	DefaultRouteTable_KindAPIVersion   = DefaultRouteTable_Kind + "." + CRDGroupVersion.String()
	DefaultRouteTable_GroupVersionKind = CRDGroupVersion.WithKind(DefaultRouteTable_Kind)
)

func init() {
	SchemeBuilder.Register(&DefaultRouteTable{}, &DefaultRouteTableList{})
}
