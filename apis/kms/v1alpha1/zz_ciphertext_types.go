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

type CiphertextObservation struct {
	CiphertextBlob *string `json:"ciphertextBlob,omitempty" tf:"ciphertext_blob,omitempty"`
}

type CiphertextParameters struct {

	// +kubebuilder:validation:Optional
	Context map[string]*string `json:"context,omitempty" tf:"context,omitempty"`

	// +kubebuilder:validation:Required
	KeyID *string `json:"keyId" tf:"key_id,omitempty"`

	// +kubebuilder:validation:Required
	PlaintextSecretRef v1.SecretKeySelector `json:"plaintextSecretRef" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// CiphertextSpec defines the desired state of Ciphertext
type CiphertextSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CiphertextParameters `json:"forProvider"`
}

// CiphertextStatus defines the observed state of Ciphertext.
type CiphertextStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CiphertextObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Ciphertext is the Schema for the Ciphertexts API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type Ciphertext struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CiphertextSpec   `json:"spec"`
	Status            CiphertextStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CiphertextList contains a list of Ciphertexts
type CiphertextList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Ciphertext `json:"items"`
}

// Repository type metadata.
var (
	CiphertextKind             = "Ciphertext"
	CiphertextGroupKind        = schema.GroupKind{Group: Group, Kind: CiphertextKind}.String()
	CiphertextKindAPIVersion   = CiphertextKind + "." + GroupVersion.String()
	CiphertextGroupVersionKind = GroupVersion.WithKind(CiphertextKind)
)

func init() {
	SchemeBuilder.Register(&Ciphertext{}, &CiphertextList{})
}
