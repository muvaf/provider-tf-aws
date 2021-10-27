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

type JupyterServerAppSettingsDefaultResourceSpecObservation struct {
}

type JupyterServerAppSettingsDefaultResourceSpecParameters struct {

	// +kubebuilder:validation:Optional
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type,omitempty"`

	// +kubebuilder:validation:Optional
	SagemakerImageArn *string `json:"sagemakerImageArn,omitempty" tf:"sagemaker_image_arn,omitempty"`
}

type KernelGatewayAppSettingsCustomImageObservation struct {
}

type KernelGatewayAppSettingsCustomImageParameters struct {

	// +kubebuilder:validation:Required
	AppImageConfigName *string `json:"appImageConfigName" tf:"app_image_config_name,omitempty"`

	// +kubebuilder:validation:Required
	ImageName *string `json:"imageName" tf:"image_name,omitempty"`

	// +kubebuilder:validation:Optional
	ImageVersionNumber *int64 `json:"imageVersionNumber,omitempty" tf:"image_version_number,omitempty"`
}

type UserProfileObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	HomeEfsFileSystemUID *string `json:"homeEfsFileSystemUid,omitempty" tf:"home_efs_file_system_uid,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type UserProfileParameters struct {

	// +kubebuilder:validation:Required
	DomainID *string `json:"domainId" tf:"domain_id,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	SingleSignOnUserIdentifier *string `json:"singleSignOnUserIdentifier,omitempty" tf:"single_sign_on_user_identifier,omitempty"`

	// +kubebuilder:validation:Optional
	SingleSignOnUserValue *string `json:"singleSignOnUserValue,omitempty" tf:"single_sign_on_user_value,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Required
	UserProfileName *string `json:"userProfileName" tf:"user_profile_name,omitempty"`

	// +kubebuilder:validation:Optional
	UserSettings []UserSettingsParameters `json:"userSettings,omitempty" tf:"user_settings,omitempty"`
}

type UserSettingsJupyterServerAppSettingsObservation struct {
}

type UserSettingsJupyterServerAppSettingsParameters struct {

	// +kubebuilder:validation:Required
	DefaultResourceSpec []JupyterServerAppSettingsDefaultResourceSpecParameters `json:"defaultResourceSpec" tf:"default_resource_spec,omitempty"`
}

type UserSettingsKernelGatewayAppSettingsDefaultResourceSpecObservation struct {
}

type UserSettingsKernelGatewayAppSettingsDefaultResourceSpecParameters struct {

	// +kubebuilder:validation:Optional
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type,omitempty"`

	// +kubebuilder:validation:Optional
	SagemakerImageArn *string `json:"sagemakerImageArn,omitempty" tf:"sagemaker_image_arn,omitempty"`
}

type UserSettingsKernelGatewayAppSettingsObservation struct {
}

type UserSettingsKernelGatewayAppSettingsParameters struct {

	// +kubebuilder:validation:Optional
	CustomImage []KernelGatewayAppSettingsCustomImageParameters `json:"customImage,omitempty" tf:"custom_image,omitempty"`

	// +kubebuilder:validation:Required
	DefaultResourceSpec []UserSettingsKernelGatewayAppSettingsDefaultResourceSpecParameters `json:"defaultResourceSpec" tf:"default_resource_spec,omitempty"`
}

type UserSettingsObservation struct {
}

type UserSettingsParameters struct {

	// +kubebuilder:validation:Required
	ExecutionRole *string `json:"executionRole" tf:"execution_role,omitempty"`

	// +kubebuilder:validation:Optional
	JupyterServerAppSettings []UserSettingsJupyterServerAppSettingsParameters `json:"jupyterServerAppSettings,omitempty" tf:"jupyter_server_app_settings,omitempty"`

	// +kubebuilder:validation:Optional
	KernelGatewayAppSettings []UserSettingsKernelGatewayAppSettingsParameters `json:"kernelGatewayAppSettings,omitempty" tf:"kernel_gateway_app_settings,omitempty"`

	// +kubebuilder:validation:Optional
	SecurityGroups []*string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`

	// +kubebuilder:validation:Optional
	SharingSettings []UserSettingsSharingSettingsParameters `json:"sharingSettings,omitempty" tf:"sharing_settings,omitempty"`

	// +kubebuilder:validation:Optional
	TensorBoardAppSettings []UserSettingsTensorBoardAppSettingsParameters `json:"tensorBoardAppSettings,omitempty" tf:"tensor_board_app_settings,omitempty"`
}

type UserSettingsSharingSettingsObservation struct {
}

type UserSettingsSharingSettingsParameters struct {

	// +kubebuilder:validation:Optional
	NotebookOutputOption *string `json:"notebookOutputOption,omitempty" tf:"notebook_output_option,omitempty"`

	// +kubebuilder:validation:Optional
	S3KmsKeyID *string `json:"s3KmsKeyId,omitempty" tf:"s3_kms_key_id,omitempty"`

	// +kubebuilder:validation:Optional
	S3OutputPath *string `json:"s3OutputPath,omitempty" tf:"s3_output_path,omitempty"`
}

type UserSettingsTensorBoardAppSettingsDefaultResourceSpecObservation struct {
}

type UserSettingsTensorBoardAppSettingsDefaultResourceSpecParameters struct {

	// +kubebuilder:validation:Optional
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type,omitempty"`

	// +kubebuilder:validation:Optional
	SagemakerImageArn *string `json:"sagemakerImageArn,omitempty" tf:"sagemaker_image_arn,omitempty"`
}

type UserSettingsTensorBoardAppSettingsObservation struct {
}

type UserSettingsTensorBoardAppSettingsParameters struct {

	// +kubebuilder:validation:Required
	DefaultResourceSpec []UserSettingsTensorBoardAppSettingsDefaultResourceSpecParameters `json:"defaultResourceSpec" tf:"default_resource_spec,omitempty"`
}

// UserProfileSpec defines the desired state of UserProfile
type UserProfileSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UserProfileParameters `json:"forProvider"`
}

// UserProfileStatus defines the observed state of UserProfile.
type UserProfileStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UserProfileObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// UserProfile is the Schema for the UserProfiles API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type UserProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              UserProfileSpec   `json:"spec"`
	Status            UserProfileStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UserProfileList contains a list of UserProfiles
type UserProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UserProfile `json:"items"`
}

// Repository type metadata.
var (
	UserProfileKind             = "UserProfile"
	UserProfileGroupKind        = schema.GroupKind{Group: Group, Kind: UserProfileKind}.String()
	UserProfileKindAPIVersion   = UserProfileKind + "." + GroupVersion.String()
	UserProfileGroupVersionKind = GroupVersion.WithKind(UserProfileKind)
)

func init() {
	SchemeBuilder.Register(&UserProfile{}, &UserProfileList{})
}
