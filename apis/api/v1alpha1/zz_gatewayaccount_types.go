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

type GatewayAccountObservation struct {
	ThrottleSettings []ThrottleSettingsObservation `json:"throttleSettings,omitempty" tf:"throttle_settings,omitempty"`
}

type GatewayAccountParameters struct {

	// +kubebuilder:validation:Optional
	CloudwatchRoleArn *string `json:"cloudwatchRoleArn,omitempty" tf:"cloudwatch_role_arn,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

type ThrottleSettingsObservation struct {
	BurstLimit *int64 `json:"burstLimit,omitempty" tf:"burst_limit,omitempty"`

	RateLimit *float64 `json:"rateLimit,omitempty" tf:"rate_limit,omitempty"`
}

type ThrottleSettingsParameters struct {
}

// GatewayAccountSpec defines the desired state of GatewayAccount
type GatewayAccountSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GatewayAccountParameters `json:"forProvider"`
}

// GatewayAccountStatus defines the observed state of GatewayAccount.
type GatewayAccountStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GatewayAccountObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GatewayAccount is the Schema for the GatewayAccounts API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type GatewayAccount struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GatewayAccountSpec   `json:"spec"`
	Status            GatewayAccountStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GatewayAccountList contains a list of GatewayAccounts
type GatewayAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GatewayAccount `json:"items"`
}

// Repository type metadata.
var (
	GatewayAccountKind             = "GatewayAccount"
	GatewayAccountGroupKind        = schema.GroupKind{Group: Group, Kind: GatewayAccountKind}.String()
	GatewayAccountKindAPIVersion   = GatewayAccountKind + "." + GroupVersion.String()
	GatewayAccountGroupVersionKind = GroupVersion.WithKind(GatewayAccountKind)
)

func init() {
	SchemeBuilder.Register(&GatewayAccount{}, &GatewayAccountList{})
}
