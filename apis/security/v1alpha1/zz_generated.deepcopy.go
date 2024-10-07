//go:build !ignore_autogenerated

/*
Copyright 2022 The Crossplane Authors.

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
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APICredentials) DeepCopyInto(out *APICredentials) {
	*out = *in
	in.CommonCredentialSelectors.DeepCopyInto(&out.CommonCredentialSelectors)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APICredentials.
func (in *APICredentials) DeepCopy() *APICredentials {
	if in == nil {
		return nil
	}
	out := new(APICredentials)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalaccountTrustConfiguration) DeepCopyInto(out *GlobalaccountTrustConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalaccountTrustConfiguration.
func (in *GlobalaccountTrustConfiguration) DeepCopy() *GlobalaccountTrustConfiguration {
	if in == nil {
		return nil
	}
	out := new(GlobalaccountTrustConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GlobalaccountTrustConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalaccountTrustConfigurationInitParameters) DeepCopyInto(out *GlobalaccountTrustConfigurationInitParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.IdentityProvider != nil {
		in, out := &in.IdentityProvider, &out.IdentityProvider
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Origin != nil {
		in, out := &in.Origin, &out.Origin
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalaccountTrustConfigurationInitParameters.
func (in *GlobalaccountTrustConfigurationInitParameters) DeepCopy() *GlobalaccountTrustConfigurationInitParameters {
	if in == nil {
		return nil
	}
	out := new(GlobalaccountTrustConfigurationInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalaccountTrustConfigurationList) DeepCopyInto(out *GlobalaccountTrustConfigurationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GlobalaccountTrustConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalaccountTrustConfigurationList.
func (in *GlobalaccountTrustConfigurationList) DeepCopy() *GlobalaccountTrustConfigurationList {
	if in == nil {
		return nil
	}
	out := new(GlobalaccountTrustConfigurationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GlobalaccountTrustConfigurationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalaccountTrustConfigurationObservation) DeepCopyInto(out *GlobalaccountTrustConfigurationObservation) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.IdentityProvider != nil {
		in, out := &in.IdentityProvider, &out.IdentityProvider
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Origin != nil {
		in, out := &in.Origin, &out.Origin
		*out = new(string)
		**out = **in
	}
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(string)
		**out = **in
	}
	if in.ReadOnly != nil {
		in, out := &in.ReadOnly, &out.ReadOnly
		*out = new(bool)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalaccountTrustConfigurationObservation.
func (in *GlobalaccountTrustConfigurationObservation) DeepCopy() *GlobalaccountTrustConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(GlobalaccountTrustConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalaccountTrustConfigurationParameters) DeepCopyInto(out *GlobalaccountTrustConfigurationParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.IdentityProvider != nil {
		in, out := &in.IdentityProvider, &out.IdentityProvider
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Origin != nil {
		in, out := &in.Origin, &out.Origin
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalaccountTrustConfigurationParameters.
func (in *GlobalaccountTrustConfigurationParameters) DeepCopy() *GlobalaccountTrustConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(GlobalaccountTrustConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalaccountTrustConfigurationSpec) DeepCopyInto(out *GlobalaccountTrustConfigurationSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalaccountTrustConfigurationSpec.
func (in *GlobalaccountTrustConfigurationSpec) DeepCopy() *GlobalaccountTrustConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(GlobalaccountTrustConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalaccountTrustConfigurationStatus) DeepCopyInto(out *GlobalaccountTrustConfigurationStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalaccountTrustConfigurationStatus.
func (in *GlobalaccountTrustConfigurationStatus) DeepCopy() *GlobalaccountTrustConfigurationStatus {
	if in == nil {
		return nil
	}
	out := new(GlobalaccountTrustConfigurationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleCollection) DeepCopyInto(out *RoleCollection) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleCollection.
func (in *RoleCollection) DeepCopy() *RoleCollection {
	if in == nil {
		return nil
	}
	out := new(RoleCollection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RoleCollection) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleCollectionAssignment) DeepCopyInto(out *RoleCollectionAssignment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleCollectionAssignment.
func (in *RoleCollectionAssignment) DeepCopy() *RoleCollectionAssignment {
	if in == nil {
		return nil
	}
	out := new(RoleCollectionAssignment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RoleCollectionAssignment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleCollectionAssignmentList) DeepCopyInto(out *RoleCollectionAssignmentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RoleCollectionAssignment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleCollectionAssignmentList.
func (in *RoleCollectionAssignmentList) DeepCopy() *RoleCollectionAssignmentList {
	if in == nil {
		return nil
	}
	out := new(RoleCollectionAssignmentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RoleCollectionAssignmentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleCollectionAssignmentObservation) DeepCopyInto(out *RoleCollectionAssignmentObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleCollectionAssignmentObservation.
func (in *RoleCollectionAssignmentObservation) DeepCopy() *RoleCollectionAssignmentObservation {
	if in == nil {
		return nil
	}
	out := new(RoleCollectionAssignmentObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleCollectionAssignmentParameters) DeepCopyInto(out *RoleCollectionAssignmentParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleCollectionAssignmentParameters.
func (in *RoleCollectionAssignmentParameters) DeepCopy() *RoleCollectionAssignmentParameters {
	if in == nil {
		return nil
	}
	out := new(RoleCollectionAssignmentParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleCollectionAssignmentSpec) DeepCopyInto(out *RoleCollectionAssignmentSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
	in.APICredentials.DeepCopyInto(&out.APICredentials)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleCollectionAssignmentSpec.
func (in *RoleCollectionAssignmentSpec) DeepCopy() *RoleCollectionAssignmentSpec {
	if in == nil {
		return nil
	}
	out := new(RoleCollectionAssignmentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleCollectionAssignmentStatus) DeepCopyInto(out *RoleCollectionAssignmentStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleCollectionAssignmentStatus.
func (in *RoleCollectionAssignmentStatus) DeepCopy() *RoleCollectionAssignmentStatus {
	if in == nil {
		return nil
	}
	out := new(RoleCollectionAssignmentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleCollectionList) DeepCopyInto(out *RoleCollectionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RoleCollection, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleCollectionList.
func (in *RoleCollectionList) DeepCopy() *RoleCollectionList {
	if in == nil {
		return nil
	}
	out := new(RoleCollectionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RoleCollectionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleCollectionObservation) DeepCopyInto(out *RoleCollectionObservation) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.RoleReferences != nil {
		in, out := &in.RoleReferences, &out.RoleReferences
		*out = new([]RoleReference)
		if **in != nil {
			in, out := *in, *out
			*out = make([]RoleReference, len(*in))
			copy(*out, *in)
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleCollectionObservation.
func (in *RoleCollectionObservation) DeepCopy() *RoleCollectionObservation {
	if in == nil {
		return nil
	}
	out := new(RoleCollectionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleCollectionParameters) DeepCopyInto(out *RoleCollectionParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.RoleReferences != nil {
		in, out := &in.RoleReferences, &out.RoleReferences
		*out = make([]RoleReference, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleCollectionParameters.
func (in *RoleCollectionParameters) DeepCopy() *RoleCollectionParameters {
	if in == nil {
		return nil
	}
	out := new(RoleCollectionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleCollectionSpec) DeepCopyInto(out *RoleCollectionSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.APICredentials.DeepCopyInto(&out.APICredentials)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleCollectionSpec.
func (in *RoleCollectionSpec) DeepCopy() *RoleCollectionSpec {
	if in == nil {
		return nil
	}
	out := new(RoleCollectionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleCollectionStatus) DeepCopyInto(out *RoleCollectionStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleCollectionStatus.
func (in *RoleCollectionStatus) DeepCopy() *RoleCollectionStatus {
	if in == nil {
		return nil
	}
	out := new(RoleCollectionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleReference) DeepCopyInto(out *RoleReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleReference.
func (in *RoleReference) DeepCopy() *RoleReference {
	if in == nil {
		return nil
	}
	out := new(RoleReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubaccountTrustConfiguration) DeepCopyInto(out *SubaccountTrustConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubaccountTrustConfiguration.
func (in *SubaccountTrustConfiguration) DeepCopy() *SubaccountTrustConfiguration {
	if in == nil {
		return nil
	}
	out := new(SubaccountTrustConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SubaccountTrustConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubaccountTrustConfigurationInitParameters) DeepCopyInto(out *SubaccountTrustConfigurationInitParameters) {
	*out = *in
	if in.AutoCreateShadowUsers != nil {
		in, out := &in.AutoCreateShadowUsers, &out.AutoCreateShadowUsers
		*out = new(bool)
		**out = **in
	}
	if in.AvailableForUserLogon != nil {
		in, out := &in.AvailableForUserLogon, &out.AvailableForUserLogon
		*out = new(bool)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Domain != nil {
		in, out := &in.Domain, &out.Domain
		*out = new(string)
		**out = **in
	}
	if in.IdentityProvider != nil {
		in, out := &in.IdentityProvider, &out.IdentityProvider
		*out = new(string)
		**out = **in
	}
	if in.LinkText != nil {
		in, out := &in.LinkText, &out.LinkText
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.SubaccountID != nil {
		in, out := &in.SubaccountID, &out.SubaccountID
		*out = new(string)
		**out = **in
	}
	if in.SubaccountRef != nil {
		in, out := &in.SubaccountRef, &out.SubaccountRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.SubaccountSelector != nil {
		in, out := &in.SubaccountSelector, &out.SubaccountSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubaccountTrustConfigurationInitParameters.
func (in *SubaccountTrustConfigurationInitParameters) DeepCopy() *SubaccountTrustConfigurationInitParameters {
	if in == nil {
		return nil
	}
	out := new(SubaccountTrustConfigurationInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubaccountTrustConfigurationList) DeepCopyInto(out *SubaccountTrustConfigurationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SubaccountTrustConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubaccountTrustConfigurationList.
func (in *SubaccountTrustConfigurationList) DeepCopy() *SubaccountTrustConfigurationList {
	if in == nil {
		return nil
	}
	out := new(SubaccountTrustConfigurationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SubaccountTrustConfigurationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubaccountTrustConfigurationObservation) DeepCopyInto(out *SubaccountTrustConfigurationObservation) {
	*out = *in
	if in.AutoCreateShadowUsers != nil {
		in, out := &in.AutoCreateShadowUsers, &out.AutoCreateShadowUsers
		*out = new(bool)
		**out = **in
	}
	if in.AvailableForUserLogon != nil {
		in, out := &in.AvailableForUserLogon, &out.AvailableForUserLogon
		*out = new(bool)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Domain != nil {
		in, out := &in.Domain, &out.Domain
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.IdentityProvider != nil {
		in, out := &in.IdentityProvider, &out.IdentityProvider
		*out = new(string)
		**out = **in
	}
	if in.LinkText != nil {
		in, out := &in.LinkText, &out.LinkText
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Origin != nil {
		in, out := &in.Origin, &out.Origin
		*out = new(string)
		**out = **in
	}
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(string)
		**out = **in
	}
	if in.ReadOnly != nil {
		in, out := &in.ReadOnly, &out.ReadOnly
		*out = new(bool)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.SubaccountID != nil {
		in, out := &in.SubaccountID, &out.SubaccountID
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubaccountTrustConfigurationObservation.
func (in *SubaccountTrustConfigurationObservation) DeepCopy() *SubaccountTrustConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(SubaccountTrustConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubaccountTrustConfigurationParameters) DeepCopyInto(out *SubaccountTrustConfigurationParameters) {
	*out = *in
	if in.AutoCreateShadowUsers != nil {
		in, out := &in.AutoCreateShadowUsers, &out.AutoCreateShadowUsers
		*out = new(bool)
		**out = **in
	}
	if in.AvailableForUserLogon != nil {
		in, out := &in.AvailableForUserLogon, &out.AvailableForUserLogon
		*out = new(bool)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Domain != nil {
		in, out := &in.Domain, &out.Domain
		*out = new(string)
		**out = **in
	}
	if in.IdentityProvider != nil {
		in, out := &in.IdentityProvider, &out.IdentityProvider
		*out = new(string)
		**out = **in
	}
	if in.LinkText != nil {
		in, out := &in.LinkText, &out.LinkText
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.SubaccountID != nil {
		in, out := &in.SubaccountID, &out.SubaccountID
		*out = new(string)
		**out = **in
	}
	if in.SubaccountRef != nil {
		in, out := &in.SubaccountRef, &out.SubaccountRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.SubaccountSelector != nil {
		in, out := &in.SubaccountSelector, &out.SubaccountSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubaccountTrustConfigurationParameters.
func (in *SubaccountTrustConfigurationParameters) DeepCopy() *SubaccountTrustConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(SubaccountTrustConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubaccountTrustConfigurationSpec) DeepCopyInto(out *SubaccountTrustConfigurationSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubaccountTrustConfigurationSpec.
func (in *SubaccountTrustConfigurationSpec) DeepCopy() *SubaccountTrustConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(SubaccountTrustConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubaccountTrustConfigurationStatus) DeepCopyInto(out *SubaccountTrustConfigurationStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubaccountTrustConfigurationStatus.
func (in *SubaccountTrustConfigurationStatus) DeepCopy() *SubaccountTrustConfigurationStatus {
	if in == nil {
		return nil
	}
	out := new(SubaccountTrustConfigurationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *XsuaaBinding) DeepCopyInto(out *XsuaaBinding) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new XsuaaBinding.
func (in *XsuaaBinding) DeepCopy() *XsuaaBinding {
	if in == nil {
		return nil
	}
	out := new(XsuaaBinding)
	in.DeepCopyInto(out)
	return out
}