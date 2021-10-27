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

type ArtifactsObservation struct {
}

type ArtifactsParameters struct {

	// +kubebuilder:validation:Optional
	ArtifactIdentifier *string `json:"artifactIdentifier,omitempty" tf:"artifact_identifier,omitempty"`

	// +kubebuilder:validation:Optional
	EncryptionDisabled *bool `json:"encryptionDisabled,omitempty" tf:"encryption_disabled,omitempty"`

	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	NamespaceType *string `json:"namespaceType,omitempty" tf:"namespace_type,omitempty"`

	// +kubebuilder:validation:Optional
	OverrideArtifactName *bool `json:"overrideArtifactName,omitempty" tf:"override_artifact_name,omitempty"`

	// +kubebuilder:validation:Optional
	Packaging *string `json:"packaging,omitempty" tf:"packaging,omitempty"`

	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type AuthObservation struct {
}

type AuthParameters struct {

	// +kubebuilder:validation:Optional
	ResourceSecretRef v1.SecretKeySelector `json:"resourceSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type BuildBatchConfigObservation struct {
}

type BuildBatchConfigParameters struct {

	// +kubebuilder:validation:Optional
	CombineArtifacts *bool `json:"combineArtifacts,omitempty" tf:"combine_artifacts,omitempty"`

	// +kubebuilder:validation:Optional
	Restrictions []RestrictionsParameters `json:"restrictions,omitempty" tf:"restrictions,omitempty"`

	// +kubebuilder:validation:Required
	ServiceRole *string `json:"serviceRole" tf:"service_role,omitempty"`

	// +kubebuilder:validation:Optional
	TimeoutInMins *int64 `json:"timeoutInMins,omitempty" tf:"timeout_in_mins,omitempty"`
}

type BuildStatusConfigObservation struct {
}

type BuildStatusConfigParameters struct {

	// +kubebuilder:validation:Optional
	Context *string `json:"context,omitempty" tf:"context,omitempty"`

	// +kubebuilder:validation:Optional
	TargetURL *string `json:"targetUrl,omitempty" tf:"target_url,omitempty"`
}

type CacheObservation struct {
}

type CacheParameters struct {

	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// +kubebuilder:validation:Optional
	Modes []*string `json:"modes,omitempty" tf:"modes,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type CloudwatchLogsObservation struct {
}

type CloudwatchLogsParameters struct {

	// +kubebuilder:validation:Optional
	GroupName *string `json:"groupName,omitempty" tf:"group_name,omitempty"`

	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// +kubebuilder:validation:Optional
	StreamName *string `json:"streamName,omitempty" tf:"stream_name,omitempty"`
}

type EnvironmentObservation struct {
}

type EnvironmentParameters struct {

	// +kubebuilder:validation:Optional
	Certificate *string `json:"certificate,omitempty" tf:"certificate,omitempty"`

	// +kubebuilder:validation:Required
	ComputeType *string `json:"computeType" tf:"compute_type,omitempty"`

	// +kubebuilder:validation:Optional
	EnvironmentVariable []EnvironmentVariableParameters `json:"environmentVariable,omitempty" tf:"environment_variable,omitempty"`

	// +kubebuilder:validation:Required
	Image *string `json:"image" tf:"image,omitempty"`

	// +kubebuilder:validation:Optional
	ImagePullCredentialsType *string `json:"imagePullCredentialsType,omitempty" tf:"image_pull_credentials_type,omitempty"`

	// +kubebuilder:validation:Optional
	PrivilegedMode *bool `json:"privilegedMode,omitempty" tf:"privileged_mode,omitempty"`

	// +kubebuilder:validation:Optional
	RegistryCredential []RegistryCredentialParameters `json:"registryCredential,omitempty" tf:"registry_credential,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type EnvironmentVariableObservation struct {
}

type EnvironmentVariableParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type FileSystemLocationsObservation struct {
}

type FileSystemLocationsParameters struct {

	// +kubebuilder:validation:Optional
	Identifier *string `json:"identifier,omitempty" tf:"identifier,omitempty"`

	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// +kubebuilder:validation:Optional
	MountOptions *string `json:"mountOptions,omitempty" tf:"mount_options,omitempty"`

	// +kubebuilder:validation:Optional
	MountPoint *string `json:"mountPoint,omitempty" tf:"mount_point,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type GitSubmodulesConfigObservation struct {
}

type GitSubmodulesConfigParameters struct {

	// +kubebuilder:validation:Required
	FetchSubmodules *bool `json:"fetchSubmodules" tf:"fetch_submodules,omitempty"`
}

type LogsConfigObservation struct {
}

type LogsConfigParameters struct {

	// +kubebuilder:validation:Optional
	CloudwatchLogs []CloudwatchLogsParameters `json:"cloudwatchLogs,omitempty" tf:"cloudwatch_logs,omitempty"`

	// +kubebuilder:validation:Optional
	S3Logs []S3LogsParameters `json:"s3Logs,omitempty" tf:"s3_logs,omitempty"`
}

type ProjectObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	BadgeURL *string `json:"badgeUrl,omitempty" tf:"badge_url,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type ProjectParameters struct {

	// +kubebuilder:validation:Required
	Artifacts []ArtifactsParameters `json:"artifacts" tf:"artifacts,omitempty"`

	// +kubebuilder:validation:Optional
	BadgeEnabled *bool `json:"badgeEnabled,omitempty" tf:"badge_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	BuildBatchConfig []BuildBatchConfigParameters `json:"buildBatchConfig,omitempty" tf:"build_batch_config,omitempty"`

	// +kubebuilder:validation:Optional
	BuildTimeout *int64 `json:"buildTimeout,omitempty" tf:"build_timeout,omitempty"`

	// +kubebuilder:validation:Optional
	Cache []CacheParameters `json:"cache,omitempty" tf:"cache,omitempty"`

	// +kubebuilder:validation:Optional
	ConcurrentBuildLimit *int64 `json:"concurrentBuildLimit,omitempty" tf:"concurrent_build_limit,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	EncryptionKey *string `json:"encryptionKey,omitempty" tf:"encryption_key,omitempty"`

	// +kubebuilder:validation:Required
	Environment []EnvironmentParameters `json:"environment" tf:"environment,omitempty"`

	// +kubebuilder:validation:Optional
	FileSystemLocations []FileSystemLocationsParameters `json:"fileSystemLocations,omitempty" tf:"file_system_locations,omitempty"`

	// +kubebuilder:validation:Optional
	LogsConfig []LogsConfigParameters `json:"logsConfig,omitempty" tf:"logs_config,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	QueuedTimeout *int64 `json:"queuedTimeout,omitempty" tf:"queued_timeout,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	SecondaryArtifacts []SecondaryArtifactsParameters `json:"secondaryArtifacts,omitempty" tf:"secondary_artifacts,omitempty"`

	// +kubebuilder:validation:Optional
	SecondarySources []SecondarySourcesParameters `json:"secondarySources,omitempty" tf:"secondary_sources,omitempty"`

	// +kubebuilder:validation:Required
	ServiceRole *string `json:"serviceRole" tf:"service_role,omitempty"`

	// +kubebuilder:validation:Required
	Source []SourceParameters `json:"source" tf:"source,omitempty"`

	// +kubebuilder:validation:Optional
	SourceVersion *string `json:"sourceVersion,omitempty" tf:"source_version,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	VpcConfig []VpcConfigParameters `json:"vpcConfig,omitempty" tf:"vpc_config,omitempty"`
}

type RegistryCredentialObservation struct {
}

type RegistryCredentialParameters struct {

	// +kubebuilder:validation:Required
	Credential *string `json:"credential" tf:"credential,omitempty"`

	// +kubebuilder:validation:Required
	CredentialProvider *string `json:"credentialProvider" tf:"credential_provider,omitempty"`
}

type RestrictionsObservation struct {
}

type RestrictionsParameters struct {

	// +kubebuilder:validation:Optional
	ComputeTypesAllowed []*string `json:"computeTypesAllowed,omitempty" tf:"compute_types_allowed,omitempty"`

	// +kubebuilder:validation:Optional
	MaximumBuildsAllowed *int64 `json:"maximumBuildsAllowed,omitempty" tf:"maximum_builds_allowed,omitempty"`
}

type S3LogsObservation struct {
}

type S3LogsParameters struct {

	// +kubebuilder:validation:Optional
	EncryptionDisabled *bool `json:"encryptionDisabled,omitempty" tf:"encryption_disabled,omitempty"`

	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type SecondaryArtifactsObservation struct {
}

type SecondaryArtifactsParameters struct {

	// +kubebuilder:validation:Required
	ArtifactIdentifier *string `json:"artifactIdentifier" tf:"artifact_identifier,omitempty"`

	// +kubebuilder:validation:Optional
	EncryptionDisabled *bool `json:"encryptionDisabled,omitempty" tf:"encryption_disabled,omitempty"`

	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	NamespaceType *string `json:"namespaceType,omitempty" tf:"namespace_type,omitempty"`

	// +kubebuilder:validation:Optional
	OverrideArtifactName *bool `json:"overrideArtifactName,omitempty" tf:"override_artifact_name,omitempty"`

	// +kubebuilder:validation:Optional
	Packaging *string `json:"packaging,omitempty" tf:"packaging,omitempty"`

	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type SecondarySourcesObservation struct {
}

type SecondarySourcesParameters struct {

	// +kubebuilder:validation:Optional
	Auth []AuthParameters `json:"auth,omitempty" tf:"auth,omitempty"`

	// +kubebuilder:validation:Optional
	BuildStatusConfig []BuildStatusConfigParameters `json:"buildStatusConfig,omitempty" tf:"build_status_config,omitempty"`

	// +kubebuilder:validation:Optional
	Buildspec *string `json:"buildspec,omitempty" tf:"buildspec,omitempty"`

	// +kubebuilder:validation:Optional
	GitCloneDepth *int64 `json:"gitCloneDepth,omitempty" tf:"git_clone_depth,omitempty"`

	// +kubebuilder:validation:Optional
	GitSubmodulesConfig []GitSubmodulesConfigParameters `json:"gitSubmodulesConfig,omitempty" tf:"git_submodules_config,omitempty"`

	// +kubebuilder:validation:Optional
	InsecureSsl *bool `json:"insecureSsl,omitempty" tf:"insecure_ssl,omitempty"`

	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// +kubebuilder:validation:Optional
	ReportBuildStatus *bool `json:"reportBuildStatus,omitempty" tf:"report_build_status,omitempty"`

	// +kubebuilder:validation:Required
	SourceIdentifier *string `json:"sourceIdentifier" tf:"source_identifier,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type SourceAuthObservation struct {
}

type SourceAuthParameters struct {

	// +kubebuilder:validation:Optional
	ResourceSecretRef v1.SecretKeySelector `json:"resourceSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type SourceBuildStatusConfigObservation struct {
}

type SourceBuildStatusConfigParameters struct {

	// +kubebuilder:validation:Optional
	Context *string `json:"context,omitempty" tf:"context,omitempty"`

	// +kubebuilder:validation:Optional
	TargetURL *string `json:"targetUrl,omitempty" tf:"target_url,omitempty"`
}

type SourceGitSubmodulesConfigObservation struct {
}

type SourceGitSubmodulesConfigParameters struct {

	// +kubebuilder:validation:Required
	FetchSubmodules *bool `json:"fetchSubmodules" tf:"fetch_submodules,omitempty"`
}

type SourceObservation struct {
}

type SourceParameters struct {

	// +kubebuilder:validation:Optional
	Auth []SourceAuthParameters `json:"auth,omitempty" tf:"auth,omitempty"`

	// +kubebuilder:validation:Optional
	BuildStatusConfig []SourceBuildStatusConfigParameters `json:"buildStatusConfig,omitempty" tf:"build_status_config,omitempty"`

	// +kubebuilder:validation:Optional
	Buildspec *string `json:"buildspec,omitempty" tf:"buildspec,omitempty"`

	// +kubebuilder:validation:Optional
	GitCloneDepth *int64 `json:"gitCloneDepth,omitempty" tf:"git_clone_depth,omitempty"`

	// +kubebuilder:validation:Optional
	GitSubmodulesConfig []SourceGitSubmodulesConfigParameters `json:"gitSubmodulesConfig,omitempty" tf:"git_submodules_config,omitempty"`

	// +kubebuilder:validation:Optional
	InsecureSsl *bool `json:"insecureSsl,omitempty" tf:"insecure_ssl,omitempty"`

	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// +kubebuilder:validation:Optional
	ReportBuildStatus *bool `json:"reportBuildStatus,omitempty" tf:"report_build_status,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type VpcConfigObservation struct {
}

type VpcConfigParameters struct {

	// +kubebuilder:validation:Required
	SecurityGroupIds []*string `json:"securityGroupIds" tf:"security_group_ids,omitempty"`

	// +kubebuilder:validation:Required
	Subnets []*string `json:"subnets" tf:"subnets,omitempty"`

	// +kubebuilder:validation:Required
	VpcID *string `json:"vpcId" tf:"vpc_id,omitempty"`
}

// ProjectSpec defines the desired state of Project
type ProjectSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProjectParameters `json:"forProvider"`
}

// ProjectStatus defines the observed state of Project.
type ProjectStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProjectObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Project is the Schema for the Projects API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type Project struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProjectSpec   `json:"spec"`
	Status            ProjectStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProjectList contains a list of Projects
type ProjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Project `json:"items"`
}

// Repository type metadata.
var (
	ProjectKind             = "Project"
	ProjectGroupKind        = schema.GroupKind{Group: Group, Kind: ProjectKind}.String()
	ProjectKindAPIVersion   = ProjectKind + "." + GroupVersion.String()
	ProjectGroupVersionKind = GroupVersion.WithKind(ProjectKind)
)

func init() {
	SchemeBuilder.Register(&Project{}, &ProjectList{})
}
