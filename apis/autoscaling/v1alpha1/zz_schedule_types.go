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

type ScheduleObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`
}

type ScheduleParameters struct {

	// +kubebuilder:validation:Required
	AutoscalingGroupName *string `json:"autoscalingGroupName" tf:"autoscaling_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	DesiredCapacity *int64 `json:"desiredCapacity,omitempty" tf:"desired_capacity,omitempty"`

	// +kubebuilder:validation:Optional
	EndTime *string `json:"endTime,omitempty" tf:"end_time,omitempty"`

	// +kubebuilder:validation:Optional
	MaxSize *int64 `json:"maxSize,omitempty" tf:"max_size,omitempty"`

	// +kubebuilder:validation:Optional
	MinSize *int64 `json:"minSize,omitempty" tf:"min_size,omitempty"`

	// +kubebuilder:validation:Optional
	Recurrence *string `json:"recurrence,omitempty" tf:"recurrence,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	ScheduledActionName *string `json:"scheduledActionName" tf:"scheduled_action_name,omitempty"`

	// +kubebuilder:validation:Optional
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`

	// +kubebuilder:validation:Optional
	TimeZone *string `json:"timeZone,omitempty" tf:"time_zone,omitempty"`
}

// ScheduleSpec defines the desired state of Schedule
type ScheduleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ScheduleParameters `json:"forProvider"`
}

// ScheduleStatus defines the observed state of Schedule.
type ScheduleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ScheduleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Schedule is the Schema for the Schedules API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type Schedule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ScheduleSpec   `json:"spec"`
	Status            ScheduleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ScheduleList contains a list of Schedules
type ScheduleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Schedule `json:"items"`
}

// Repository type metadata.
var (
	ScheduleKind             = "Schedule"
	ScheduleGroupKind        = schema.GroupKind{Group: Group, Kind: ScheduleKind}.String()
	ScheduleKindAPIVersion   = ScheduleKind + "." + GroupVersion.String()
	ScheduleGroupVersionKind = GroupVersion.WithKind(ScheduleKind)
)

func init() {
	SchemeBuilder.Register(&Schedule{}, &ScheduleList{})
}
