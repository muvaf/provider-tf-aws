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

type GatewayObservation struct {
	NetworkInterfaceID *string `json:"networkInterfaceId,omitempty" tf:"network_interface_id,omitempty"`

	PrivateIP *string `json:"privateIp,omitempty" tf:"private_ip,omitempty"`

	PublicIP *string `json:"publicIp,omitempty" tf:"public_ip,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type GatewayParameters struct {

	// +kubebuilder:validation:Optional
	AllocationID *string `json:"allocationId,omitempty" tf:"allocation_id,omitempty"`

	// +kubebuilder:validation:Optional
	ConnectivityType *string `json:"connectivityType,omitempty" tf:"connectivity_type,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	SubnetID *string `json:"subnetId" tf:"subnet_id,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// GatewaySpec defines the desired state of Gateway
type GatewaySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GatewayParameters `json:"forProvider"`
}

// GatewayStatus defines the observed state of Gateway.
type GatewayStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GatewayObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Gateway is the Schema for the Gateways API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type Gateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GatewaySpec   `json:"spec"`
	Status            GatewayStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GatewayList contains a list of Gateways
type GatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Gateway `json:"items"`
}

// Repository type metadata.
var (
	GatewayKind             = "Gateway"
	GatewayGroupKind        = schema.GroupKind{Group: Group, Kind: GatewayKind}.String()
	GatewayKindAPIVersion   = GatewayKind + "." + GroupVersion.String()
	GatewayGroupVersionKind = GroupVersion.WithKind(GatewayKind)
)

func init() {
	SchemeBuilder.Register(&Gateway{}, &GatewayList{})
}
