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

type ApnsVoipSandboxChannelObservation struct {
}

type ApnsVoipSandboxChannelParameters struct {

	// +kubebuilder:validation:Required
	ApplicationID *string `json:"applicationId" tf:"application_id,omitempty"`

	// +kubebuilder:validation:Optional
	BundleIDSecretRef v1.SecretKeySelector `json:"bundleIdSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	CertificateSecretRef v1.SecretKeySelector `json:"certificateSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	DefaultAuthenticationMethod *string `json:"defaultAuthenticationMethod,omitempty" tf:"default_authentication_method,omitempty"`

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	PrivateKeySecretRef v1.SecretKeySelector `json:"privateKeySecretRef,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	TeamIDSecretRef v1.SecretKeySelector `json:"teamIdSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	TokenKeyIDSecretRef v1.SecretKeySelector `json:"tokenKeyIdSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	TokenKeySecretRef v1.SecretKeySelector `json:"tokenKeySecretRef,omitempty" tf:"-"`
}

// ApnsVoipSandboxChannelSpec defines the desired state of ApnsVoipSandboxChannel
type ApnsVoipSandboxChannelSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ApnsVoipSandboxChannelParameters `json:"forProvider"`
}

// ApnsVoipSandboxChannelStatus defines the observed state of ApnsVoipSandboxChannel.
type ApnsVoipSandboxChannelStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ApnsVoipSandboxChannelObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ApnsVoipSandboxChannel is the Schema for the ApnsVoipSandboxChannels API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type ApnsVoipSandboxChannel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApnsVoipSandboxChannelSpec   `json:"spec"`
	Status            ApnsVoipSandboxChannelStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApnsVoipSandboxChannelList contains a list of ApnsVoipSandboxChannels
type ApnsVoipSandboxChannelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApnsVoipSandboxChannel `json:"items"`
}

// Repository type metadata.
var (
	ApnsVoipSandboxChannelKind             = "ApnsVoipSandboxChannel"
	ApnsVoipSandboxChannelGroupKind        = schema.GroupKind{Group: Group, Kind: ApnsVoipSandboxChannelKind}.String()
	ApnsVoipSandboxChannelKindAPIVersion   = ApnsVoipSandboxChannelKind + "." + GroupVersion.String()
	ApnsVoipSandboxChannelGroupVersionKind = GroupVersion.WithKind(ApnsVoipSandboxChannelKind)
)

func init() {
	SchemeBuilder.Register(&ApnsVoipSandboxChannel{}, &ApnsVoipSandboxChannelList{})
}
