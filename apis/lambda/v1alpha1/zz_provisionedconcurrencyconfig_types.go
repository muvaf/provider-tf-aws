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

type ProvisionedConcurrencyConfigObservation struct {
}

type ProvisionedConcurrencyConfigParameters struct {

	// +kubebuilder:validation:Required
	FunctionName *string `json:"functionName" tf:"function_name,omitempty"`

	// +kubebuilder:validation:Required
	ProvisionedConcurrentExecutions *int64 `json:"provisionedConcurrentExecutions" tf:"provisioned_concurrent_executions,omitempty"`

	// +kubebuilder:validation:Required
	Qualifier *string `json:"qualifier" tf:"qualifier,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// ProvisionedConcurrencyConfigSpec defines the desired state of ProvisionedConcurrencyConfig
type ProvisionedConcurrencyConfigSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProvisionedConcurrencyConfigParameters `json:"forProvider"`
}

// ProvisionedConcurrencyConfigStatus defines the observed state of ProvisionedConcurrencyConfig.
type ProvisionedConcurrencyConfigStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProvisionedConcurrencyConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ProvisionedConcurrencyConfig is the Schema for the ProvisionedConcurrencyConfigs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type ProvisionedConcurrencyConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProvisionedConcurrencyConfigSpec   `json:"spec"`
	Status            ProvisionedConcurrencyConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProvisionedConcurrencyConfigList contains a list of ProvisionedConcurrencyConfigs
type ProvisionedConcurrencyConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProvisionedConcurrencyConfig `json:"items"`
}

// Repository type metadata.
var (
	ProvisionedConcurrencyConfigKind             = "ProvisionedConcurrencyConfig"
	ProvisionedConcurrencyConfigGroupKind        = schema.GroupKind{Group: Group, Kind: ProvisionedConcurrencyConfigKind}.String()
	ProvisionedConcurrencyConfigKindAPIVersion   = ProvisionedConcurrencyConfigKind + "." + GroupVersion.String()
	ProvisionedConcurrencyConfigGroupVersionKind = GroupVersion.WithKind(ProvisionedConcurrencyConfigKind)
)

func init() {
	SchemeBuilder.Register(&ProvisionedConcurrencyConfig{}, &ProvisionedConcurrencyConfigList{})
}
