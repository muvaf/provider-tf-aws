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

type CaptureContentTypeHeaderObservation struct {
}

type CaptureContentTypeHeaderParameters struct {

	// +kubebuilder:validation:Optional
	CsvContentTypes []*string `json:"csvContentTypes,omitempty" tf:"csv_content_types,omitempty"`

	// +kubebuilder:validation:Optional
	JSONContentTypes []*string `json:"jsonContentTypes,omitempty" tf:"json_content_types,omitempty"`
}

type CaptureOptionsObservation struct {
}

type CaptureOptionsParameters struct {

	// +kubebuilder:validation:Required
	CaptureMode *string `json:"captureMode" tf:"capture_mode,omitempty"`
}

type DataCaptureConfigObservation struct {
}

type DataCaptureConfigParameters struct {

	// +kubebuilder:validation:Optional
	CaptureContentTypeHeader []CaptureContentTypeHeaderParameters `json:"captureContentTypeHeader,omitempty" tf:"capture_content_type_header,omitempty"`

	// +kubebuilder:validation:Required
	CaptureOptions []CaptureOptionsParameters `json:"captureOptions" tf:"capture_options,omitempty"`

	// +kubebuilder:validation:Required
	DestinationS3URI *string `json:"destinationS3Uri" tf:"destination_s3_uri,omitempty"`

	// +kubebuilder:validation:Optional
	EnableCapture *bool `json:"enableCapture,omitempty" tf:"enable_capture,omitempty"`

	// +kubebuilder:validation:Required
	InitialSamplingPercentage *float64 `json:"initialSamplingPercentage" tf:"initial_sampling_percentage,omitempty"`

	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`
}

type EndpointConfigurationObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type EndpointConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	DataCaptureConfig []DataCaptureConfigParameters `json:"dataCaptureConfig,omitempty" tf:"data_capture_config,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-aws/apis/kms/v1alpha2.Key
	// +kubebuilder:validation:Optional
	KMSKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`

	// +kubebuilder:validation:Optional
	KMSKeyArnRef *v1.Reference `json:"kmsKeyArnRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	KMSKeyArnSelector *v1.Selector `json:"kmsKeyArnSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	ProductionVariants []ProductionVariantsParameters `json:"productionVariants" tf:"production_variants,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ProductionVariantsObservation struct {
}

type ProductionVariantsParameters struct {

	// +kubebuilder:validation:Optional
	AcceleratorType *string `json:"acceleratorType,omitempty" tf:"accelerator_type,omitempty"`

	// +kubebuilder:validation:Required
	InitialInstanceCount *float64 `json:"initialInstanceCount" tf:"initial_instance_count,omitempty"`

	// +kubebuilder:validation:Optional
	InitialVariantWeight *float64 `json:"initialVariantWeight,omitempty" tf:"initial_variant_weight,omitempty"`

	// +kubebuilder:validation:Required
	InstanceType *string `json:"instanceType" tf:"instance_type,omitempty"`

	// +kubebuilder:validation:Required
	ModelName *string `json:"modelName" tf:"model_name,omitempty"`

	// +kubebuilder:validation:Optional
	VariantName *string `json:"variantName,omitempty" tf:"variant_name,omitempty"`
}

// EndpointConfigurationSpec defines the desired state of EndpointConfiguration
type EndpointConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EndpointConfigurationParameters `json:"forProvider"`
}

// EndpointConfigurationStatus defines the observed state of EndpointConfiguration.
type EndpointConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EndpointConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EndpointConfiguration is the Schema for the EndpointConfigurations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type EndpointConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EndpointConfigurationSpec   `json:"spec"`
	Status            EndpointConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EndpointConfigurationList contains a list of EndpointConfigurations
type EndpointConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EndpointConfiguration `json:"items"`
}

// Repository type metadata.
var (
	EndpointConfiguration_Kind             = "EndpointConfiguration"
	EndpointConfiguration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EndpointConfiguration_Kind}.String()
	EndpointConfiguration_KindAPIVersion   = EndpointConfiguration_Kind + "." + CRDGroupVersion.String()
	EndpointConfiguration_GroupVersionKind = CRDGroupVersion.WithKind(EndpointConfiguration_Kind)
)

func init() {
	SchemeBuilder.Register(&EndpointConfiguration{}, &EndpointConfigurationList{})
}
