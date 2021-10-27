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

type ActivationObservation struct {
	ActivationCode *string `json:"activationCode,omitempty" tf:"activation_code,omitempty"`

	Expired *bool `json:"expired,omitempty" tf:"expired,omitempty"`

	RegistrationCount *int64 `json:"registrationCount,omitempty" tf:"registration_count,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type ActivationParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	ExpirationDate *string `json:"expirationDate,omitempty" tf:"expiration_date,omitempty"`

	// +kubebuilder:validation:Required
	IamRole *string `json:"iamRole" tf:"iam_role,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	RegistrationLimit *int64 `json:"registrationLimit,omitempty" tf:"registration_limit,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// ActivationSpec defines the desired state of Activation
type ActivationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ActivationParameters `json:"forProvider"`
}

// ActivationStatus defines the observed state of Activation.
type ActivationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ActivationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Activation is the Schema for the Activations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type Activation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ActivationSpec   `json:"spec"`
	Status            ActivationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ActivationList contains a list of Activations
type ActivationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Activation `json:"items"`
}

// Repository type metadata.
var (
	ActivationKind             = "Activation"
	ActivationGroupKind        = schema.GroupKind{Group: Group, Kind: ActivationKind}.String()
	ActivationKindAPIVersion   = ActivationKind + "." + GroupVersion.String()
	ActivationGroupVersionKind = GroupVersion.WithKind(ActivationKind)
)

func init() {
	SchemeBuilder.Register(&Activation{}, &ActivationList{})
}
