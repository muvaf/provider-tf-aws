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

type PermissionObservation struct {
}

type PermissionParameters struct {

	// +kubebuilder:validation:Required
	Action *string `json:"action" tf:"action,omitempty"`

	// +kubebuilder:validation:Optional
	EventSourceToken *string `json:"eventSourceToken,omitempty" tf:"event_source_token,omitempty"`

	// +kubebuilder:validation:Required
	FunctionName *string `json:"functionName" tf:"function_name,omitempty"`

	// +kubebuilder:validation:Required
	Principal *string `json:"principal" tf:"principal,omitempty"`

	// +kubebuilder:validation:Optional
	Qualifier *string `json:"qualifier,omitempty" tf:"qualifier,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	SourceAccount *string `json:"sourceAccount,omitempty" tf:"source_account,omitempty"`

	// +kubebuilder:validation:Optional
	SourceArn *string `json:"sourceArn,omitempty" tf:"source_arn,omitempty"`

	// +kubebuilder:validation:Optional
	StatementID *string `json:"statementId,omitempty" tf:"statement_id,omitempty"`

	// +kubebuilder:validation:Optional
	StatementIDPrefix *string `json:"statementIdPrefix,omitempty" tf:"statement_id_prefix,omitempty"`
}

// PermissionSpec defines the desired state of Permission
type PermissionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PermissionParameters `json:"forProvider"`
}

// PermissionStatus defines the observed state of Permission.
type PermissionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PermissionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Permission is the Schema for the Permissions API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type Permission struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PermissionSpec   `json:"spec"`
	Status            PermissionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PermissionList contains a list of Permissions
type PermissionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Permission `json:"items"`
}

// Repository type metadata.
var (
	PermissionKind             = "Permission"
	PermissionGroupKind        = schema.GroupKind{Group: Group, Kind: PermissionKind}.String()
	PermissionKindAPIVersion   = PermissionKind + "." + GroupVersion.String()
	PermissionGroupVersionKind = GroupVersion.WithKind(PermissionKind)
)

func init() {
	SchemeBuilder.Register(&Permission{}, &PermissionList{})
}
