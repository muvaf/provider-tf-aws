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

type DestinationObservation struct {
}

type DestinationParameters struct {

	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"region,omitempty"`

	// +kubebuilder:validation:Required
	RegistryID *string `json:"registryId" tf:"registry_id,omitempty"`
}

type ReplicationConfigurationObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	RegistryID *string `json:"registryId,omitempty" tf:"registry_id,omitempty"`
}

type ReplicationConfigurationParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	ReplicationConfiguration []ReplicationConfigurationReplicationConfigurationParameters `json:"replicationConfiguration,omitempty" tf:"replication_configuration,omitempty"`
}

type ReplicationConfigurationReplicationConfigurationObservation struct {
}

type ReplicationConfigurationReplicationConfigurationParameters struct {

	// +kubebuilder:validation:Required
	Rule []RuleParameters `json:"rule" tf:"rule,omitempty"`
}

type RuleObservation struct {
}

type RuleParameters struct {

	// +kubebuilder:validation:Required
	Destination []DestinationParameters `json:"destination" tf:"destination,omitempty"`
}

// ReplicationConfigurationSpec defines the desired state of ReplicationConfiguration
type ReplicationConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ReplicationConfigurationParameters `json:"forProvider"`
}

// ReplicationConfigurationStatus defines the observed state of ReplicationConfiguration.
type ReplicationConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ReplicationConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ReplicationConfiguration is the Schema for the ReplicationConfigurations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type ReplicationConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ReplicationConfigurationSpec   `json:"spec"`
	Status            ReplicationConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ReplicationConfigurationList contains a list of ReplicationConfigurations
type ReplicationConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ReplicationConfiguration `json:"items"`
}

// Repository type metadata.
var (
	ReplicationConfiguration_Kind             = "ReplicationConfiguration"
	ReplicationConfiguration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ReplicationConfiguration_Kind}.String()
	ReplicationConfiguration_KindAPIVersion   = ReplicationConfiguration_Kind + "." + CRDGroupVersion.String()
	ReplicationConfiguration_GroupVersionKind = CRDGroupVersion.WithKind(ReplicationConfiguration_Kind)
)

func init() {
	SchemeBuilder.Register(&ReplicationConfiguration{}, &ReplicationConfigurationList{})
}
