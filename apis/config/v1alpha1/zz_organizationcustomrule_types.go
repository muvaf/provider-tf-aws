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

type OrganizationCustomRuleObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`
}

type OrganizationCustomRuleParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	ExcludedAccounts []*string `json:"excludedAccounts,omitempty" tf:"excluded_accounts,omitempty"`

	// +kubebuilder:validation:Optional
	InputParameters *string `json:"inputParameters,omitempty" tf:"input_parameters,omitempty"`

	// +kubebuilder:validation:Required
	LambdaFunctionArn *string `json:"lambdaFunctionArn" tf:"lambda_function_arn,omitempty"`

	// +kubebuilder:validation:Optional
	MaximumExecutionFrequency *string `json:"maximumExecutionFrequency,omitempty" tf:"maximum_execution_frequency,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceIDScope *string `json:"resourceIdScope,omitempty" tf:"resource_id_scope,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceTypesScope []*string `json:"resourceTypesScope,omitempty" tf:"resource_types_scope,omitempty"`

	// +kubebuilder:validation:Optional
	TagKeyScope *string `json:"tagKeyScope,omitempty" tf:"tag_key_scope,omitempty"`

	// +kubebuilder:validation:Optional
	TagValueScope *string `json:"tagValueScope,omitempty" tf:"tag_value_scope,omitempty"`

	// +kubebuilder:validation:Required
	TriggerTypes []*string `json:"triggerTypes" tf:"trigger_types,omitempty"`
}

// OrganizationCustomRuleSpec defines the desired state of OrganizationCustomRule
type OrganizationCustomRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OrganizationCustomRuleParameters `json:"forProvider"`
}

// OrganizationCustomRuleStatus defines the observed state of OrganizationCustomRule.
type OrganizationCustomRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OrganizationCustomRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// OrganizationCustomRule is the Schema for the OrganizationCustomRules API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type OrganizationCustomRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OrganizationCustomRuleSpec   `json:"spec"`
	Status            OrganizationCustomRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OrganizationCustomRuleList contains a list of OrganizationCustomRules
type OrganizationCustomRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OrganizationCustomRule `json:"items"`
}

// Repository type metadata.
var (
	OrganizationCustomRuleKind             = "OrganizationCustomRule"
	OrganizationCustomRuleGroupKind        = schema.GroupKind{Group: Group, Kind: OrganizationCustomRuleKind}.String()
	OrganizationCustomRuleKindAPIVersion   = OrganizationCustomRuleKind + "." + GroupVersion.String()
	OrganizationCustomRuleGroupVersionKind = GroupVersion.WithKind(OrganizationCustomRuleKind)
)

func init() {
	SchemeBuilder.Register(&OrganizationCustomRule{}, &OrganizationCustomRuleList{})
}
