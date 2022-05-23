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

type NetworkACLEgressObservation struct {
}

type NetworkACLEgressParameters struct {

	// +kubebuilder:validation:Optional
	Action *string `json:"action,omitempty" tf:"action"`

	// +kubebuilder:validation:Optional
	CidrBlock *string `json:"cidrBlock,omitempty" tf:"cidr_block"`

	// +kubebuilder:validation:Optional
	FromPort *float64 `json:"fromPort,omitempty" tf:"from_port"`

	// +kubebuilder:validation:Optional
	IPv6CidrBlock *string `json:"ipv6CidrBlock,omitempty" tf:"ipv6_cidr_block"`

	// +kubebuilder:validation:Optional
	IcmpCode *float64 `json:"icmpCode,omitempty" tf:"icmp_code"`

	// +kubebuilder:validation:Optional
	IcmpType *float64 `json:"icmpType,omitempty" tf:"icmp_type"`

	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol"`

	// +kubebuilder:validation:Optional
	RuleNo *float64 `json:"ruleNo,omitempty" tf:"rule_no"`

	// +kubebuilder:validation:Optional
	ToPort *float64 `json:"toPort,omitempty" tf:"to_port"`
}

type NetworkACLIngressObservation struct {
}

type NetworkACLIngressParameters struct {

	// +kubebuilder:validation:Optional
	Action *string `json:"action,omitempty" tf:"action"`

	// +kubebuilder:validation:Optional
	CidrBlock *string `json:"cidrBlock,omitempty" tf:"cidr_block"`

	// +kubebuilder:validation:Optional
	FromPort *float64 `json:"fromPort,omitempty" tf:"from_port"`

	// +kubebuilder:validation:Optional
	IPv6CidrBlock *string `json:"ipv6CidrBlock,omitempty" tf:"ipv6_cidr_block"`

	// +kubebuilder:validation:Optional
	IcmpCode *float64 `json:"icmpCode,omitempty" tf:"icmp_code"`

	// +kubebuilder:validation:Optional
	IcmpType *float64 `json:"icmpType,omitempty" tf:"icmp_type"`

	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol"`

	// +kubebuilder:validation:Optional
	RuleNo *float64 `json:"ruleNo,omitempty" tf:"rule_no"`

	// +kubebuilder:validation:Optional
	ToPort *float64 `json:"toPort,omitempty" tf:"to_port"`
}

type NetworkACLObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	OwnerID *string `json:"ownerId,omitempty" tf:"owner_id,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type NetworkACLParameters struct {

	// +kubebuilder:validation:Optional
	Egress []NetworkACLEgressParameters `json:"egress,omitempty" tf:"egress,omitempty"`

	// +kubebuilder:validation:Optional
	Ingress []NetworkACLIngressParameters `json:"ingress,omitempty" tf:"ingress,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	SubnetIdRefs []v1.Reference `json:"subnetIdRefs,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SubnetIdSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-aws/apis/ec2/v1alpha2.Subnet
	// +crossplane:generate:reference:refFieldName=SubnetIdRefs
	// +crossplane:generate:reference:selectorFieldName=SubnetIdSelector
	// +kubebuilder:validation:Optional
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-aws/apis/ec2/v1alpha2.VPC
	// +crossplane:generate:reference:refFieldName=VpcIdRef
	// +crossplane:generate:reference:selectorFieldName=VpcIdSelector
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// +kubebuilder:validation:Optional
	VpcIdRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	VpcIdSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

// NetworkACLSpec defines the desired state of NetworkACL
type NetworkACLSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NetworkACLParameters `json:"forProvider"`
}

// NetworkACLStatus defines the observed state of NetworkACL.
type NetworkACLStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NetworkACLObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkACL is the Schema for the NetworkACLs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type NetworkACL struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkACLSpec   `json:"spec"`
	Status            NetworkACLStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkACLList contains a list of NetworkACLs
type NetworkACLList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkACL `json:"items"`
}

// Repository type metadata.
var (
	NetworkACL_Kind             = "NetworkACL"
	NetworkACL_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NetworkACL_Kind}.String()
	NetworkACL_KindAPIVersion   = NetworkACL_Kind + "." + CRDGroupVersion.String()
	NetworkACL_GroupVersionKind = CRDGroupVersion.WithKind(NetworkACL_Kind)
)

func init() {
	SchemeBuilder.Register(&NetworkACL{}, &NetworkACLList{})
}
