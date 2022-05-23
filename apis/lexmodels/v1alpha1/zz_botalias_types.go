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

type BotAliasObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	Checksum *string `json:"checksum,omitempty" tf:"checksum,omitempty"`

	ConversationLogs []ConversationLogsObservation `json:"conversationLogs,omitempty" tf:"conversation_logs,omitempty"`

	CreatedDate *string `json:"createdDate,omitempty" tf:"created_date,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	LastUpdatedDate *string `json:"lastUpdatedDate,omitempty" tf:"last_updated_date,omitempty"`
}

type BotAliasParameters struct {

	// +kubebuilder:validation:Required
	BotName *string `json:"botName" tf:"bot_name,omitempty"`

	// +kubebuilder:validation:Required
	BotVersion *string `json:"botVersion" tf:"bot_version,omitempty"`

	// +kubebuilder:validation:Optional
	ConversationLogs []ConversationLogsParameters `json:"conversationLogs,omitempty" tf:"conversation_logs,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

type ConversationLogsObservation struct {
	LogSettings []LogSettingsObservation `json:"logSettings,omitempty" tf:"log_settings,omitempty"`
}

type ConversationLogsParameters struct {

	// +kubebuilder:validation:Required
	IAMRoleArn *string `json:"iamRoleArn" tf:"iam_role_arn,omitempty"`

	// +kubebuilder:validation:Optional
	LogSettings []LogSettingsParameters `json:"logSettings,omitempty" tf:"log_settings,omitempty"`
}

type LogSettingsObservation struct {
	ResourcePrefix *string `json:"resourcePrefix,omitempty" tf:"resource_prefix,omitempty"`
}

type LogSettingsParameters struct {

	// +kubebuilder:validation:Required
	Destination *string `json:"destination" tf:"destination,omitempty"`

	// +kubebuilder:validation:Optional
	KMSKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`

	// +kubebuilder:validation:Required
	LogType *string `json:"logType" tf:"log_type,omitempty"`

	// +kubebuilder:validation:Required
	ResourceArn *string `json:"resourceArn" tf:"resource_arn,omitempty"`
}

// BotAliasSpec defines the desired state of BotAlias
type BotAliasSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BotAliasParameters `json:"forProvider"`
}

// BotAliasStatus defines the observed state of BotAlias.
type BotAliasStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BotAliasObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BotAlias is the Schema for the BotAliass API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type BotAlias struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BotAliasSpec   `json:"spec"`
	Status            BotAliasStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BotAliasList contains a list of BotAliass
type BotAliasList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BotAlias `json:"items"`
}

// Repository type metadata.
var (
	BotAlias_Kind             = "BotAlias"
	BotAlias_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BotAlias_Kind}.String()
	BotAlias_KindAPIVersion   = BotAlias_Kind + "." + CRDGroupVersion.String()
	BotAlias_GroupVersionKind = CRDGroupVersion.WithKind(BotAlias_Kind)
)

func init() {
	SchemeBuilder.Register(&BotAlias{}, &BotAliasList{})
}
