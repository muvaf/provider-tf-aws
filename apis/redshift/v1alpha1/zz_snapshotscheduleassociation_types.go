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

type SnapshotScheduleAssociationObservation struct {
}

type SnapshotScheduleAssociationParameters struct {

	// +kubebuilder:validation:Required
	ClusterIdentifier *string `json:"clusterIdentifier" tf:"cluster_identifier,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	ScheduleIdentifier *string `json:"scheduleIdentifier" tf:"schedule_identifier,omitempty"`
}

// SnapshotScheduleAssociationSpec defines the desired state of SnapshotScheduleAssociation
type SnapshotScheduleAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SnapshotScheduleAssociationParameters `json:"forProvider"`
}

// SnapshotScheduleAssociationStatus defines the observed state of SnapshotScheduleAssociation.
type SnapshotScheduleAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SnapshotScheduleAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SnapshotScheduleAssociation is the Schema for the SnapshotScheduleAssociations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type SnapshotScheduleAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SnapshotScheduleAssociationSpec   `json:"spec"`
	Status            SnapshotScheduleAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SnapshotScheduleAssociationList contains a list of SnapshotScheduleAssociations
type SnapshotScheduleAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SnapshotScheduleAssociation `json:"items"`
}

// Repository type metadata.
var (
	SnapshotScheduleAssociationKind             = "SnapshotScheduleAssociation"
	SnapshotScheduleAssociationGroupKind        = schema.GroupKind{Group: Group, Kind: SnapshotScheduleAssociationKind}.String()
	SnapshotScheduleAssociationKindAPIVersion   = SnapshotScheduleAssociationKind + "." + GroupVersion.String()
	SnapshotScheduleAssociationGroupVersionKind = GroupVersion.WithKind(SnapshotScheduleAssociationKind)
)

func init() {
	SchemeBuilder.Register(&SnapshotScheduleAssociation{}, &SnapshotScheduleAssociationList{})
}
