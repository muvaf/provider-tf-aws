// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cloudtrail) DeepCopyInto(out *Cloudtrail) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cloudtrail.
func (in *Cloudtrail) DeepCopy() *Cloudtrail {
	if in == nil {
		return nil
	}
	out := new(Cloudtrail)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Cloudtrail) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudtrailList) DeepCopyInto(out *CloudtrailList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Cloudtrail, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudtrailList.
func (in *CloudtrailList) DeepCopy() *CloudtrailList {
	if in == nil {
		return nil
	}
	out := new(CloudtrailList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudtrailList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudtrailObservation) DeepCopyInto(out *CloudtrailObservation) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.HomeRegion != nil {
		in, out := &in.HomeRegion, &out.HomeRegion
		*out = new(string)
		**out = **in
	}
	if in.TagsAll != nil {
		in, out := &in.TagsAll, &out.TagsAll
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudtrailObservation.
func (in *CloudtrailObservation) DeepCopy() *CloudtrailObservation {
	if in == nil {
		return nil
	}
	out := new(CloudtrailObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudtrailParameters) DeepCopyInto(out *CloudtrailParameters) {
	*out = *in
	if in.CloudWatchLogsGroupArn != nil {
		in, out := &in.CloudWatchLogsGroupArn, &out.CloudWatchLogsGroupArn
		*out = new(string)
		**out = **in
	}
	if in.CloudWatchLogsRoleArn != nil {
		in, out := &in.CloudWatchLogsRoleArn, &out.CloudWatchLogsRoleArn
		*out = new(string)
		**out = **in
	}
	if in.EnableLogFileValidation != nil {
		in, out := &in.EnableLogFileValidation, &out.EnableLogFileValidation
		*out = new(bool)
		**out = **in
	}
	if in.EnableLogging != nil {
		in, out := &in.EnableLogging, &out.EnableLogging
		*out = new(bool)
		**out = **in
	}
	if in.EventSelector != nil {
		in, out := &in.EventSelector, &out.EventSelector
		*out = make([]EventSelectorParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.IncludeGlobalServiceEvents != nil {
		in, out := &in.IncludeGlobalServiceEvents, &out.IncludeGlobalServiceEvents
		*out = new(bool)
		**out = **in
	}
	if in.InsightSelector != nil {
		in, out := &in.InsightSelector, &out.InsightSelector
		*out = make([]InsightSelectorParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.IsMultiRegionTrail != nil {
		in, out := &in.IsMultiRegionTrail, &out.IsMultiRegionTrail
		*out = new(bool)
		**out = **in
	}
	if in.IsOrganizationTrail != nil {
		in, out := &in.IsOrganizationTrail, &out.IsOrganizationTrail
		*out = new(bool)
		**out = **in
	}
	if in.KmsKeyID != nil {
		in, out := &in.KmsKeyID, &out.KmsKeyID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.S3BucketName != nil {
		in, out := &in.S3BucketName, &out.S3BucketName
		*out = new(string)
		**out = **in
	}
	if in.S3KeyPrefix != nil {
		in, out := &in.S3KeyPrefix, &out.S3KeyPrefix
		*out = new(string)
		**out = **in
	}
	if in.SnsTopicName != nil {
		in, out := &in.SnsTopicName, &out.SnsTopicName
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudtrailParameters.
func (in *CloudtrailParameters) DeepCopy() *CloudtrailParameters {
	if in == nil {
		return nil
	}
	out := new(CloudtrailParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudtrailSpec) DeepCopyInto(out *CloudtrailSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudtrailSpec.
func (in *CloudtrailSpec) DeepCopy() *CloudtrailSpec {
	if in == nil {
		return nil
	}
	out := new(CloudtrailSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudtrailStatus) DeepCopyInto(out *CloudtrailStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudtrailStatus.
func (in *CloudtrailStatus) DeepCopy() *CloudtrailStatus {
	if in == nil {
		return nil
	}
	out := new(CloudtrailStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataResourceObservation) DeepCopyInto(out *DataResourceObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataResourceObservation.
func (in *DataResourceObservation) DeepCopy() *DataResourceObservation {
	if in == nil {
		return nil
	}
	out := new(DataResourceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataResourceParameters) DeepCopyInto(out *DataResourceParameters) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataResourceParameters.
func (in *DataResourceParameters) DeepCopy() *DataResourceParameters {
	if in == nil {
		return nil
	}
	out := new(DataResourceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventSelectorObservation) DeepCopyInto(out *EventSelectorObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventSelectorObservation.
func (in *EventSelectorObservation) DeepCopy() *EventSelectorObservation {
	if in == nil {
		return nil
	}
	out := new(EventSelectorObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventSelectorParameters) DeepCopyInto(out *EventSelectorParameters) {
	*out = *in
	if in.DataResource != nil {
		in, out := &in.DataResource, &out.DataResource
		*out = make([]DataResourceParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.IncludeManagementEvents != nil {
		in, out := &in.IncludeManagementEvents, &out.IncludeManagementEvents
		*out = new(bool)
		**out = **in
	}
	if in.ReadWriteType != nil {
		in, out := &in.ReadWriteType, &out.ReadWriteType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventSelectorParameters.
func (in *EventSelectorParameters) DeepCopy() *EventSelectorParameters {
	if in == nil {
		return nil
	}
	out := new(EventSelectorParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InsightSelectorObservation) DeepCopyInto(out *InsightSelectorObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InsightSelectorObservation.
func (in *InsightSelectorObservation) DeepCopy() *InsightSelectorObservation {
	if in == nil {
		return nil
	}
	out := new(InsightSelectorObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InsightSelectorParameters) DeepCopyInto(out *InsightSelectorParameters) {
	*out = *in
	if in.InsightType != nil {
		in, out := &in.InsightType, &out.InsightType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InsightSelectorParameters.
func (in *InsightSelectorParameters) DeepCopy() *InsightSelectorParameters {
	if in == nil {
		return nil
	}
	out := new(InsightSelectorParameters)
	in.DeepCopyInto(out)
	return out
}
