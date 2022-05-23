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

type DomainAssociationObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	CertificateVerificationDNSRecord *string `json:"certificateVerificationDnsRecord,omitempty" tf:"certificate_verification_dns_record,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	SubDomain []SubDomainObservation `json:"subDomain,omitempty" tf:"sub_domain,omitempty"`
}

type DomainAssociationParameters struct {

	// +kubebuilder:validation:Required
	AppID *string `json:"appId" tf:"app_id,omitempty"`

	// +kubebuilder:validation:Required
	DomainName *string `json:"domainName" tf:"domain_name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	SubDomain []SubDomainParameters `json:"subDomain" tf:"sub_domain,omitempty"`

	// +kubebuilder:validation:Optional
	WaitForVerification *bool `json:"waitForVerification,omitempty" tf:"wait_for_verification,omitempty"`
}

type SubDomainObservation struct {
	DNSRecord *string `json:"dnsRecord,omitempty" tf:"dns_record,omitempty"`

	Verified *bool `json:"verified,omitempty" tf:"verified,omitempty"`
}

type SubDomainParameters struct {

	// +kubebuilder:validation:Required
	BranchName *string `json:"branchName" tf:"branch_name,omitempty"`

	// +kubebuilder:validation:Required
	Prefix *string `json:"prefix" tf:"prefix,omitempty"`
}

// DomainAssociationSpec defines the desired state of DomainAssociation
type DomainAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DomainAssociationParameters `json:"forProvider"`
}

// DomainAssociationStatus defines the observed state of DomainAssociation.
type DomainAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DomainAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DomainAssociation is the Schema for the DomainAssociations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type DomainAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DomainAssociationSpec   `json:"spec"`
	Status            DomainAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DomainAssociationList contains a list of DomainAssociations
type DomainAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DomainAssociation `json:"items"`
}

// Repository type metadata.
var (
	DomainAssociation_Kind             = "DomainAssociation"
	DomainAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DomainAssociation_Kind}.String()
	DomainAssociation_KindAPIVersion   = DomainAssociation_Kind + "." + CRDGroupVersion.String()
	DomainAssociation_GroupVersionKind = CRDGroupVersion.WithKind(DomainAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&DomainAssociation{}, &DomainAssociationList{})
}
