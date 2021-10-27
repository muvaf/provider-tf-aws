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

type CognitoMemberDefinitionObservation struct {
}

type CognitoMemberDefinitionParameters struct {

	// +kubebuilder:validation:Required
	ClientID *string `json:"clientId" tf:"client_id,omitempty"`

	// +kubebuilder:validation:Required
	UserGroup *string `json:"userGroup" tf:"user_group,omitempty"`

	// +kubebuilder:validation:Required
	UserPool *string `json:"userPool" tf:"user_pool,omitempty"`
}

type MemberDefinitionObservation struct {
}

type MemberDefinitionParameters struct {

	// +kubebuilder:validation:Optional
	CognitoMemberDefinition []CognitoMemberDefinitionParameters `json:"cognitoMemberDefinition,omitempty" tf:"cognito_member_definition,omitempty"`

	// +kubebuilder:validation:Optional
	OidcMemberDefinition []OidcMemberDefinitionParameters `json:"oidcMemberDefinition,omitempty" tf:"oidc_member_definition,omitempty"`
}

type NotificationConfigurationObservation struct {
}

type NotificationConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	NotificationTopicArn *string `json:"notificationTopicArn,omitempty" tf:"notification_topic_arn,omitempty"`
}

type OidcMemberDefinitionObservation struct {
}

type OidcMemberDefinitionParameters struct {

	// +kubebuilder:validation:Required
	Groups []*string `json:"groups" tf:"groups,omitempty"`
}

type WorkteamObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	Subdomain *string `json:"subdomain,omitempty" tf:"subdomain,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type WorkteamParameters struct {

	// +kubebuilder:validation:Required
	Description *string `json:"description" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	MemberDefinition []MemberDefinitionParameters `json:"memberDefinition" tf:"member_definition,omitempty"`

	// +kubebuilder:validation:Optional
	NotificationConfiguration []NotificationConfigurationParameters `json:"notificationConfiguration,omitempty" tf:"notification_configuration,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Required
	WorkforceName *string `json:"workforceName" tf:"workforce_name,omitempty"`

	// +kubebuilder:validation:Required
	WorkteamName *string `json:"workteamName" tf:"workteam_name,omitempty"`
}

// WorkteamSpec defines the desired state of Workteam
type WorkteamSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WorkteamParameters `json:"forProvider"`
}

// WorkteamStatus defines the observed state of Workteam.
type WorkteamStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WorkteamObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Workteam is the Schema for the Workteams API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type Workteam struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WorkteamSpec   `json:"spec"`
	Status            WorkteamStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WorkteamList contains a list of Workteams
type WorkteamList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Workteam `json:"items"`
}

// Repository type metadata.
var (
	WorkteamKind             = "Workteam"
	WorkteamGroupKind        = schema.GroupKind{Group: Group, Kind: WorkteamKind}.String()
	WorkteamKindAPIVersion   = WorkteamKind + "." + GroupVersion.String()
	WorkteamGroupVersionKind = GroupVersion.WithKind(WorkteamKind)
)

func init() {
	SchemeBuilder.Register(&Workteam{}, &WorkteamList{})
}
