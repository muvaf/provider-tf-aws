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

type CustomizedMetricSpecificationObservation struct {
}

type CustomizedMetricSpecificationParameters struct {

	// +kubebuilder:validation:Optional
	Dimensions []DimensionsParameters `json:"dimensions,omitempty" tf:"dimensions,omitempty"`

	// +kubebuilder:validation:Required
	MetricName *string `json:"metricName" tf:"metric_name,omitempty"`

	// +kubebuilder:validation:Required
	Namespace *string `json:"namespace" tf:"namespace,omitempty"`

	// +kubebuilder:validation:Required
	Statistic *string `json:"statistic" tf:"statistic,omitempty"`

	// +kubebuilder:validation:Optional
	Unit *string `json:"unit,omitempty" tf:"unit,omitempty"`
}

type DimensionsObservation struct {
}

type DimensionsParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type PolicyObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`
}

type PolicyParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	PolicyType *string `json:"policyType,omitempty" tf:"policy_type,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	ResourceID *string `json:"resourceId" tf:"resource_id,omitempty"`

	// +kubebuilder:validation:Required
	ScalableDimension *string `json:"scalableDimension" tf:"scalable_dimension,omitempty"`

	// +kubebuilder:validation:Required
	ServiceNamespace *string `json:"serviceNamespace" tf:"service_namespace,omitempty"`

	// +kubebuilder:validation:Optional
	StepScalingPolicyConfiguration []StepScalingPolicyConfigurationParameters `json:"stepScalingPolicyConfiguration,omitempty" tf:"step_scaling_policy_configuration,omitempty"`

	// +kubebuilder:validation:Optional
	TargetTrackingScalingPolicyConfiguration []TargetTrackingScalingPolicyConfigurationParameters `json:"targetTrackingScalingPolicyConfiguration,omitempty" tf:"target_tracking_scaling_policy_configuration,omitempty"`
}

type PredefinedMetricSpecificationObservation struct {
}

type PredefinedMetricSpecificationParameters struct {

	// +kubebuilder:validation:Required
	PredefinedMetricType *string `json:"predefinedMetricType" tf:"predefined_metric_type,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceLabel *string `json:"resourceLabel,omitempty" tf:"resource_label,omitempty"`
}

type StepAdjustmentObservation struct {
}

type StepAdjustmentParameters struct {

	// +kubebuilder:validation:Optional
	MetricIntervalLowerBound *string `json:"metricIntervalLowerBound,omitempty" tf:"metric_interval_lower_bound,omitempty"`

	// +kubebuilder:validation:Optional
	MetricIntervalUpperBound *string `json:"metricIntervalUpperBound,omitempty" tf:"metric_interval_upper_bound,omitempty"`

	// +kubebuilder:validation:Required
	ScalingAdjustment *int64 `json:"scalingAdjustment" tf:"scaling_adjustment,omitempty"`
}

type StepScalingPolicyConfigurationObservation struct {
}

type StepScalingPolicyConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	AdjustmentType *string `json:"adjustmentType,omitempty" tf:"adjustment_type,omitempty"`

	// +kubebuilder:validation:Optional
	Cooldown *int64 `json:"cooldown,omitempty" tf:"cooldown,omitempty"`

	// +kubebuilder:validation:Optional
	MetricAggregationType *string `json:"metricAggregationType,omitempty" tf:"metric_aggregation_type,omitempty"`

	// +kubebuilder:validation:Optional
	MinAdjustmentMagnitude *int64 `json:"minAdjustmentMagnitude,omitempty" tf:"min_adjustment_magnitude,omitempty"`

	// +kubebuilder:validation:Optional
	StepAdjustment []StepAdjustmentParameters `json:"stepAdjustment,omitempty" tf:"step_adjustment,omitempty"`
}

type TargetTrackingScalingPolicyConfigurationObservation struct {
}

type TargetTrackingScalingPolicyConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	CustomizedMetricSpecification []CustomizedMetricSpecificationParameters `json:"customizedMetricSpecification,omitempty" tf:"customized_metric_specification,omitempty"`

	// +kubebuilder:validation:Optional
	DisableScaleIn *bool `json:"disableScaleIn,omitempty" tf:"disable_scale_in,omitempty"`

	// +kubebuilder:validation:Optional
	PredefinedMetricSpecification []PredefinedMetricSpecificationParameters `json:"predefinedMetricSpecification,omitempty" tf:"predefined_metric_specification,omitempty"`

	// +kubebuilder:validation:Optional
	ScaleInCooldown *int64 `json:"scaleInCooldown,omitempty" tf:"scale_in_cooldown,omitempty"`

	// +kubebuilder:validation:Optional
	ScaleOutCooldown *int64 `json:"scaleOutCooldown,omitempty" tf:"scale_out_cooldown,omitempty"`

	// +kubebuilder:validation:Required
	TargetValue *float64 `json:"targetValue" tf:"target_value,omitempty"`
}

// PolicySpec defines the desired state of Policy
type PolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PolicyParameters `json:"forProvider"`
}

// PolicyStatus defines the observed state of Policy.
type PolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Policy is the Schema for the Policys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type Policy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PolicySpec   `json:"spec"`
	Status            PolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PolicyList contains a list of Policys
type PolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Policy `json:"items"`
}

// Repository type metadata.
var (
	PolicyKind             = "Policy"
	PolicyGroupKind        = schema.GroupKind{Group: Group, Kind: PolicyKind}.String()
	PolicyKindAPIVersion   = PolicyKind + "." + GroupVersion.String()
	PolicyGroupVersionKind = GroupVersion.WithKind(PolicyKind)
)

func init() {
	SchemeBuilder.Register(&Policy{}, &PolicyList{})
}
