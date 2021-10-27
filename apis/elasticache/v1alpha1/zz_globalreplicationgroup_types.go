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

type GlobalReplicationGroupObservation struct {
	ActualEngineVersion *string `json:"actualEngineVersion,omitempty" tf:"actual_engine_version,omitempty"`

	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	AtRestEncryptionEnabled *bool `json:"atRestEncryptionEnabled,omitempty" tf:"at_rest_encryption_enabled,omitempty"`

	AuthTokenEnabled *bool `json:"authTokenEnabled,omitempty" tf:"auth_token_enabled,omitempty"`

	CacheNodeType *string `json:"cacheNodeType,omitempty" tf:"cache_node_type,omitempty"`

	ClusterEnabled *bool `json:"clusterEnabled,omitempty" tf:"cluster_enabled,omitempty"`

	Engine *string `json:"engine,omitempty" tf:"engine,omitempty"`

	EngineVersionActual *string `json:"engineVersionActual,omitempty" tf:"engine_version_actual,omitempty"`

	GlobalReplicationGroupID *string `json:"globalReplicationGroupId,omitempty" tf:"global_replication_group_id,omitempty"`

	TransitEncryptionEnabled *bool `json:"transitEncryptionEnabled,omitempty" tf:"transit_encryption_enabled,omitempty"`
}

type GlobalReplicationGroupParameters struct {

	// +kubebuilder:validation:Optional
	GlobalReplicationGroupDescription *string `json:"globalReplicationGroupDescription,omitempty" tf:"global_replication_group_description,omitempty"`

	// +kubebuilder:validation:Required
	GlobalReplicationGroupIDSuffix *string `json:"globalReplicationGroupIdSuffix" tf:"global_replication_group_id_suffix,omitempty"`

	// +kubebuilder:validation:Required
	PrimaryReplicationGroupID *string `json:"primaryReplicationGroupId" tf:"primary_replication_group_id,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// GlobalReplicationGroupSpec defines the desired state of GlobalReplicationGroup
type GlobalReplicationGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GlobalReplicationGroupParameters `json:"forProvider"`
}

// GlobalReplicationGroupStatus defines the observed state of GlobalReplicationGroup.
type GlobalReplicationGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GlobalReplicationGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GlobalReplicationGroup is the Schema for the GlobalReplicationGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type GlobalReplicationGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GlobalReplicationGroupSpec   `json:"spec"`
	Status            GlobalReplicationGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GlobalReplicationGroupList contains a list of GlobalReplicationGroups
type GlobalReplicationGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GlobalReplicationGroup `json:"items"`
}

// Repository type metadata.
var (
	GlobalReplicationGroupKind             = "GlobalReplicationGroup"
	GlobalReplicationGroupGroupKind        = schema.GroupKind{Group: Group, Kind: GlobalReplicationGroupKind}.String()
	GlobalReplicationGroupKindAPIVersion   = GlobalReplicationGroupKind + "." + GroupVersion.String()
	GlobalReplicationGroupGroupVersionKind = GroupVersion.WithKind(GlobalReplicationGroupKind)
)

func init() {
	SchemeBuilder.Register(&GlobalReplicationGroup{}, &GlobalReplicationGroupList{})
}
