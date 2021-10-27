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

type CreateRuleObservation struct {
}

type CreateRuleParameters struct {

	// +kubebuilder:validation:Required
	Interval *int64 `json:"interval" tf:"interval,omitempty"`

	// +kubebuilder:validation:Optional
	IntervalUnit *string `json:"intervalUnit,omitempty" tf:"interval_unit,omitempty"`

	// +kubebuilder:validation:Optional
	Times []*string `json:"times,omitempty" tf:"times,omitempty"`
}

type LifecyclePolicyObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type LifecyclePolicyParameters struct {

	// +kubebuilder:validation:Required
	Description *string `json:"description" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	ExecutionRoleArn *string `json:"executionRoleArn" tf:"execution_role_arn,omitempty"`

	// +kubebuilder:validation:Required
	PolicyDetails []PolicyDetailsParameters `json:"policyDetails" tf:"policy_details,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type PolicyDetailsObservation struct {
}

type PolicyDetailsParameters struct {

	// +kubebuilder:validation:Required
	ResourceTypes []*string `json:"resourceTypes" tf:"resource_types,omitempty"`

	// +kubebuilder:validation:Required
	Schedule []ScheduleParameters `json:"schedule" tf:"schedule,omitempty"`

	// +kubebuilder:validation:Required
	TargetTags map[string]*string `json:"targetTags" tf:"target_tags,omitempty"`
}

type RetainRuleObservation struct {
}

type RetainRuleParameters struct {

	// +kubebuilder:validation:Required
	Count *int64 `json:"count" tf:"count,omitempty"`
}

type ScheduleObservation struct {
}

type ScheduleParameters struct {

	// +kubebuilder:validation:Optional
	CopyTags *bool `json:"copyTags,omitempty" tf:"copy_tags,omitempty"`

	// +kubebuilder:validation:Required
	CreateRule []CreateRuleParameters `json:"createRule" tf:"create_rule,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	RetainRule []RetainRuleParameters `json:"retainRule" tf:"retain_rule,omitempty"`

	// +kubebuilder:validation:Optional
	TagsToAdd map[string]*string `json:"tagsToAdd,omitempty" tf:"tags_to_add,omitempty"`
}

// LifecyclePolicySpec defines the desired state of LifecyclePolicy
type LifecyclePolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LifecyclePolicyParameters `json:"forProvider"`
}

// LifecyclePolicyStatus defines the observed state of LifecyclePolicy.
type LifecyclePolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LifecyclePolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LifecyclePolicy is the Schema for the LifecyclePolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type LifecyclePolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LifecyclePolicySpec   `json:"spec"`
	Status            LifecyclePolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LifecyclePolicyList contains a list of LifecyclePolicys
type LifecyclePolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LifecyclePolicy `json:"items"`
}

// Repository type metadata.
var (
	LifecyclePolicyKind             = "LifecyclePolicy"
	LifecyclePolicyGroupKind        = schema.GroupKind{Group: Group, Kind: LifecyclePolicyKind}.String()
	LifecyclePolicyKindAPIVersion   = LifecyclePolicyKind + "." + GroupVersion.String()
	LifecyclePolicyGroupVersionKind = GroupVersion.WithKind(LifecyclePolicyKind)
)

func init() {
	SchemeBuilder.Register(&LifecyclePolicy{}, &LifecyclePolicyList{})
}
