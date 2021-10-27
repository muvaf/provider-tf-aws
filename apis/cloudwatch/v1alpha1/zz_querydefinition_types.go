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

type QueryDefinitionObservation struct {
	QueryDefinitionID *string `json:"queryDefinitionId,omitempty" tf:"query_definition_id,omitempty"`
}

type QueryDefinitionParameters struct {

	// +kubebuilder:validation:Optional
	LogGroupNames []*string `json:"logGroupNames,omitempty" tf:"log_group_names,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	QueryString *string `json:"queryString" tf:"query_string,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// QueryDefinitionSpec defines the desired state of QueryDefinition
type QueryDefinitionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     QueryDefinitionParameters `json:"forProvider"`
}

// QueryDefinitionStatus defines the observed state of QueryDefinition.
type QueryDefinitionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        QueryDefinitionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// QueryDefinition is the Schema for the QueryDefinitions API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type QueryDefinition struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              QueryDefinitionSpec   `json:"spec"`
	Status            QueryDefinitionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// QueryDefinitionList contains a list of QueryDefinitions
type QueryDefinitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []QueryDefinition `json:"items"`
}

// Repository type metadata.
var (
	QueryDefinitionKind             = "QueryDefinition"
	QueryDefinitionGroupKind        = schema.GroupKind{Group: Group, Kind: QueryDefinitionKind}.String()
	QueryDefinitionKindAPIVersion   = QueryDefinitionKind + "." + GroupVersion.String()
	QueryDefinitionGroupVersionKind = GroupVersion.WithKind(QueryDefinitionKind)
)

func init() {
	SchemeBuilder.Register(&QueryDefinition{}, &QueryDefinitionList{})
}
