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

type UserPoolDomainObservation struct {
	AwsAccountID *string `json:"awsAccountId,omitempty" tf:"aws_account_id,omitempty"`

	CloudfrontDistributionArn *string `json:"cloudfrontDistributionArn,omitempty" tf:"cloudfront_distribution_arn,omitempty"`

	S3Bucket *string `json:"s3Bucket,omitempty" tf:"s3_bucket,omitempty"`

	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type UserPoolDomainParameters struct {

	// +kubebuilder:validation:Optional
	CertificateArn *string `json:"certificateArn,omitempty" tf:"certificate_arn,omitempty"`

	// +kubebuilder:validation:Required
	Domain *string `json:"domain" tf:"domain,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	UserPoolID *string `json:"userPoolId" tf:"user_pool_id,omitempty"`
}

// UserPoolDomainSpec defines the desired state of UserPoolDomain
type UserPoolDomainSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UserPoolDomainParameters `json:"forProvider"`
}

// UserPoolDomainStatus defines the observed state of UserPoolDomain.
type UserPoolDomainStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UserPoolDomainObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// UserPoolDomain is the Schema for the UserPoolDomains API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type UserPoolDomain struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              UserPoolDomainSpec   `json:"spec"`
	Status            UserPoolDomainStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UserPoolDomainList contains a list of UserPoolDomains
type UserPoolDomainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UserPoolDomain `json:"items"`
}

// Repository type metadata.
var (
	UserPoolDomainKind             = "UserPoolDomain"
	UserPoolDomainGroupKind        = schema.GroupKind{Group: Group, Kind: UserPoolDomainKind}.String()
	UserPoolDomainKindAPIVersion   = UserPoolDomainKind + "." + GroupVersion.String()
	UserPoolDomainGroupVersionKind = GroupVersion.WithKind(UserPoolDomainKind)
)

func init() {
	SchemeBuilder.Register(&UserPoolDomain{}, &UserPoolDomainList{})
}
