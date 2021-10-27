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

type AttachmentsSourceObservation struct {
}

type AttachmentsSourceParameters struct {

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type DocumentObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	CreatedDate *string `json:"createdDate,omitempty" tf:"created_date,omitempty"`

	DefaultVersion *string `json:"defaultVersion,omitempty" tf:"default_version,omitempty"`

	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	DocumentVersion *string `json:"documentVersion,omitempty" tf:"document_version,omitempty"`

	Hash *string `json:"hash,omitempty" tf:"hash,omitempty"`

	HashType *string `json:"hashType,omitempty" tf:"hash_type,omitempty"`

	LatestVersion *string `json:"latestVersion,omitempty" tf:"latest_version,omitempty"`

	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	Parameter []ParameterObservation `json:"parameter,omitempty" tf:"parameter,omitempty"`

	PlatformTypes []*string `json:"platformTypes,omitempty" tf:"platform_types,omitempty"`

	SchemaVersion *string `json:"schemaVersion,omitempty" tf:"schema_version,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type DocumentParameters struct {

	// +kubebuilder:validation:Optional
	AttachmentsSource []AttachmentsSourceParameters `json:"attachmentsSource,omitempty" tf:"attachments_source,omitempty"`

	// +kubebuilder:validation:Required
	Content *string `json:"content" tf:"content,omitempty"`

	// +kubebuilder:validation:Optional
	DocumentFormat *string `json:"documentFormat,omitempty" tf:"document_format,omitempty"`

	// +kubebuilder:validation:Required
	DocumentType *string `json:"documentType" tf:"document_type,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Permissions map[string]*string `json:"permissions,omitempty" tf:"permissions,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	TargetType *string `json:"targetType,omitempty" tf:"target_type,omitempty"`

	// +kubebuilder:validation:Optional
	VersionName *string `json:"versionName,omitempty" tf:"version_name,omitempty"`
}

type ParameterObservation struct {
}

type ParameterParameters struct {

	// +kubebuilder:validation:Optional
	DefaultValue *string `json:"defaultValue,omitempty" tf:"default_value,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// DocumentSpec defines the desired state of Document
type DocumentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DocumentParameters `json:"forProvider"`
}

// DocumentStatus defines the observed state of Document.
type DocumentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DocumentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Document is the Schema for the Documents API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type Document struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DocumentSpec   `json:"spec"`
	Status            DocumentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DocumentList contains a list of Documents
type DocumentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Document `json:"items"`
}

// Repository type metadata.
var (
	DocumentKind             = "Document"
	DocumentGroupKind        = schema.GroupKind{Group: Group, Kind: DocumentKind}.String()
	DocumentKindAPIVersion   = DocumentKind + "." + GroupVersion.String()
	DocumentGroupVersionKind = GroupVersion.WithKind(DocumentKind)
)

func init() {
	SchemeBuilder.Register(&Document{}, &DocumentList{})
}
