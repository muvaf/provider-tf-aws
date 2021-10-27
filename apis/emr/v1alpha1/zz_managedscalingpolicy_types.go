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

type ComputeLimitsObservation struct {
}

type ComputeLimitsParameters struct {

	// +kubebuilder:validation:Required
	MaximumCapacityUnits *int64 `json:"maximumCapacityUnits" tf:"maximum_capacity_units,omitempty"`

	// +kubebuilder:validation:Optional
	MaximumCoreCapacityUnits *int64 `json:"maximumCoreCapacityUnits,omitempty" tf:"maximum_core_capacity_units,omitempty"`

	// +kubebuilder:validation:Optional
	MaximumOndemandCapacityUnits *int64 `json:"maximumOndemandCapacityUnits,omitempty" tf:"maximum_ondemand_capacity_units,omitempty"`

	// +kubebuilder:validation:Required
	MinimumCapacityUnits *int64 `json:"minimumCapacityUnits" tf:"minimum_capacity_units,omitempty"`

	// +kubebuilder:validation:Required
	UnitType *string `json:"unitType" tf:"unit_type,omitempty"`
}

type ManagedScalingPolicyObservation struct {
}

type ManagedScalingPolicyParameters struct {

	// +kubebuilder:validation:Required
	ClusterID *string `json:"clusterId" tf:"cluster_id,omitempty"`

	// +kubebuilder:validation:Required
	ComputeLimits []ComputeLimitsParameters `json:"computeLimits" tf:"compute_limits,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// ManagedScalingPolicySpec defines the desired state of ManagedScalingPolicy
type ManagedScalingPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ManagedScalingPolicyParameters `json:"forProvider"`
}

// ManagedScalingPolicyStatus defines the observed state of ManagedScalingPolicy.
type ManagedScalingPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ManagedScalingPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ManagedScalingPolicy is the Schema for the ManagedScalingPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type ManagedScalingPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ManagedScalingPolicySpec   `json:"spec"`
	Status            ManagedScalingPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ManagedScalingPolicyList contains a list of ManagedScalingPolicys
type ManagedScalingPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagedScalingPolicy `json:"items"`
}

// Repository type metadata.
var (
	ManagedScalingPolicyKind             = "ManagedScalingPolicy"
	ManagedScalingPolicyGroupKind        = schema.GroupKind{Group: Group, Kind: ManagedScalingPolicyKind}.String()
	ManagedScalingPolicyKindAPIVersion   = ManagedScalingPolicyKind + "." + GroupVersion.String()
	ManagedScalingPolicyGroupVersionKind = GroupVersion.WithKind(ManagedScalingPolicyKind)
)

func init() {
	SchemeBuilder.Register(&ManagedScalingPolicy{}, &ManagedScalingPolicyList{})
}
