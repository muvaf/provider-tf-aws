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

type ProductObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	CreatedTime *string `json:"createdTime,omitempty" tf:"created_time,omitempty"`

	HasDefaultPath *bool `json:"hasDefaultPath,omitempty" tf:"has_default_path,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type ProductParameters struct {

	// +kubebuilder:validation:Optional
	AcceptLanguage *string `json:"acceptLanguage,omitempty" tf:"accept_language,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	Distributor *string `json:"distributor,omitempty" tf:"distributor,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Owner *string `json:"owner" tf:"owner,omitempty"`

	// +kubebuilder:validation:Required
	ProvisioningArtifactParameters []ProvisioningArtifactParametersParameters `json:"provisioningArtifactParameters" tf:"provisioning_artifact_parameters,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	SupportDescription *string `json:"supportDescription,omitempty" tf:"support_description,omitempty"`

	// +kubebuilder:validation:Optional
	SupportEmail *string `json:"supportEmail,omitempty" tf:"support_email,omitempty"`

	// +kubebuilder:validation:Optional
	SupportURL *string `json:"supportUrl,omitempty" tf:"support_url,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type ProvisioningArtifactParametersObservation struct {
}

type ProvisioningArtifactParametersParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	DisableTemplateValidation *bool `json:"disableTemplateValidation,omitempty" tf:"disable_template_validation,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	TemplatePhysicalID *string `json:"templatePhysicalId,omitempty" tf:"template_physical_id,omitempty"`

	// +kubebuilder:validation:Optional
	TemplateURL *string `json:"templateUrl,omitempty" tf:"template_url,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// ProductSpec defines the desired state of Product
type ProductSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProductParameters `json:"forProvider"`
}

// ProductStatus defines the observed state of Product.
type ProductStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProductObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Product is the Schema for the Products API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type Product struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProductSpec   `json:"spec"`
	Status            ProductStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProductList contains a list of Products
type ProductList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Product `json:"items"`
}

// Repository type metadata.
var (
	ProductKind             = "Product"
	ProductGroupKind        = schema.GroupKind{Group: Group, Kind: ProductKind}.String()
	ProductKindAPIVersion   = ProductKind + "." + GroupVersion.String()
	ProductGroupVersionKind = GroupVersion.WithKind(ProductKind)
)

func init() {
	SchemeBuilder.Register(&Product{}, &ProductList{})
}
