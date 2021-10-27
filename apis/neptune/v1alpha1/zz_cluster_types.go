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

type ClusterObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ClusterMembers []*string `json:"clusterMembers,omitempty" tf:"cluster_members,omitempty"`

	ClusterResourceID *string `json:"clusterResourceId,omitempty" tf:"cluster_resource_id,omitempty"`

	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	HostedZoneID *string `json:"hostedZoneId,omitempty" tf:"hosted_zone_id,omitempty"`

	ReaderEndpoint *string `json:"readerEndpoint,omitempty" tf:"reader_endpoint,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type ClusterParameters struct {

	// +kubebuilder:validation:Optional
	ApplyImmediately *bool `json:"applyImmediately,omitempty" tf:"apply_immediately,omitempty"`

	// +kubebuilder:validation:Optional
	AvailabilityZones []*string `json:"availabilityZones,omitempty" tf:"availability_zones,omitempty"`

	// +kubebuilder:validation:Optional
	BackupRetentionPeriod *int64 `json:"backupRetentionPeriod,omitempty" tf:"backup_retention_period,omitempty"`

	// +kubebuilder:validation:Optional
	ClusterIdentifier *string `json:"clusterIdentifier,omitempty" tf:"cluster_identifier,omitempty"`

	// +kubebuilder:validation:Optional
	ClusterIdentifierPrefix *string `json:"clusterIdentifierPrefix,omitempty" tf:"cluster_identifier_prefix,omitempty"`

	// +kubebuilder:validation:Optional
	CopyTagsToSnapshot *bool `json:"copyTagsToSnapshot,omitempty" tf:"copy_tags_to_snapshot,omitempty"`

	// +kubebuilder:validation:Optional
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// +kubebuilder:validation:Optional
	EnableCloudwatchLogsExports []*string `json:"enableCloudwatchLogsExports,omitempty" tf:"enable_cloudwatch_logs_exports,omitempty"`

	// +kubebuilder:validation:Optional
	Engine *string `json:"engine,omitempty" tf:"engine,omitempty"`

	// +kubebuilder:validation:Optional
	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`

	// +kubebuilder:validation:Optional
	FinalSnapshotIdentifier *string `json:"finalSnapshotIdentifier,omitempty" tf:"final_snapshot_identifier,omitempty"`

	// +kubebuilder:validation:Optional
	IamDatabaseAuthenticationEnabled *bool `json:"iamDatabaseAuthenticationEnabled,omitempty" tf:"iam_database_authentication_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	IamRoles []*string `json:"iamRoles,omitempty" tf:"iam_roles,omitempty"`

	// +kubebuilder:validation:Optional
	KmsKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`

	// +kubebuilder:validation:Optional
	NeptuneClusterParameterGroupName *string `json:"neptuneClusterParameterGroupName,omitempty" tf:"neptune_cluster_parameter_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	NeptuneSubnetGroupName *string `json:"neptuneSubnetGroupName,omitempty" tf:"neptune_subnet_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	Port *int64 `json:"port,omitempty" tf:"port,omitempty"`

	// +kubebuilder:validation:Optional
	PreferredBackupWindow *string `json:"preferredBackupWindow,omitempty" tf:"preferred_backup_window,omitempty"`

	// +kubebuilder:validation:Optional
	PreferredMaintenanceWindow *string `json:"preferredMaintenanceWindow,omitempty" tf:"preferred_maintenance_window,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	ReplicationSourceIdentifier *string `json:"replicationSourceIdentifier,omitempty" tf:"replication_source_identifier,omitempty"`

	// +kubebuilder:validation:Optional
	SkipFinalSnapshot *bool `json:"skipFinalSnapshot,omitempty" tf:"skip_final_snapshot,omitempty"`

	// +kubebuilder:validation:Optional
	SnapshotIdentifier *string `json:"snapshotIdentifier,omitempty" tf:"snapshot_identifier,omitempty"`

	// +kubebuilder:validation:Optional
	StorageEncrypted *bool `json:"storageEncrypted,omitempty" tf:"storage_encrypted,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	VpcSecurityGroupIds []*string `json:"vpcSecurityGroupIds,omitempty" tf:"vpc_security_group_ids,omitempty"`
}

// ClusterSpec defines the desired state of Cluster
type ClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClusterParameters `json:"forProvider"`
}

// ClusterStatus defines the observed state of Cluster.
type ClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Cluster is the Schema for the Clusters API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterSpec   `json:"spec"`
	Status            ClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterList contains a list of Clusters
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cluster `json:"items"`
}

// Repository type metadata.
var (
	ClusterKind             = "Cluster"
	ClusterGroupKind        = schema.GroupKind{Group: Group, Kind: ClusterKind}.String()
	ClusterKindAPIVersion   = ClusterKind + "." + GroupVersion.String()
	ClusterGroupVersionKind = GroupVersion.WithKind(ClusterKind)
)

func init() {
	SchemeBuilder.Register(&Cluster{}, &ClusterList{})
}
