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

type SnapshotCopyGrantObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type SnapshotCopyGrantParameters struct {

	// +kubebuilder:validation:Optional
	KmsKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	SnapshotCopyGrantName *string `json:"snapshotCopyGrantName" tf:"snapshot_copy_grant_name,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// SnapshotCopyGrantSpec defines the desired state of SnapshotCopyGrant
type SnapshotCopyGrantSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SnapshotCopyGrantParameters `json:"forProvider"`
}

// SnapshotCopyGrantStatus defines the observed state of SnapshotCopyGrant.
type SnapshotCopyGrantStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SnapshotCopyGrantObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SnapshotCopyGrant is the Schema for the SnapshotCopyGrants API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type SnapshotCopyGrant struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SnapshotCopyGrantSpec   `json:"spec"`
	Status            SnapshotCopyGrantStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SnapshotCopyGrantList contains a list of SnapshotCopyGrants
type SnapshotCopyGrantList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SnapshotCopyGrant `json:"items"`
}

// Repository type metadata.
var (
	SnapshotCopyGrantKind             = "SnapshotCopyGrant"
	SnapshotCopyGrantGroupKind        = schema.GroupKind{Group: Group, Kind: SnapshotCopyGrantKind}.String()
	SnapshotCopyGrantKindAPIVersion   = SnapshotCopyGrantKind + "." + GroupVersion.String()
	SnapshotCopyGrantGroupVersionKind = GroupVersion.WithKind(SnapshotCopyGrantKind)
)

func init() {
	SchemeBuilder.Register(&SnapshotCopyGrant{}, &SnapshotCopyGrantList{})
}
