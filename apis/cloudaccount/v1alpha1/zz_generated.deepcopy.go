//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Aws) DeepCopyInto(out *Aws) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Aws.
func (in *Aws) DeepCopy() *Aws {
	if in == nil {
		return nil
	}
	out := new(Aws)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Aws) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AwsList) DeepCopyInto(out *AwsList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Aws, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AwsList.
func (in *AwsList) DeepCopy() *AwsList {
	if in == nil {
		return nil
	}
	out := new(AwsList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AwsList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AwsObservation) DeepCopyInto(out *AwsObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AwsObservation.
func (in *AwsObservation) DeepCopy() *AwsObservation {
	if in == nil {
		return nil
	}
	out := new(AwsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AwsParameters) DeepCopyInto(out *AwsParameters) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.AwsAccessKey != nil {
		in, out := &in.AwsAccessKey, &out.AwsAccessKey
		*out = new(string)
		**out = **in
	}
	if in.AwsSecretKeySecretRef != nil {
		in, out := &in.AwsSecretKeySecretRef, &out.AwsSecretKeySecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.Context != nil {
		in, out := &in.Context, &out.Context
		*out = new(string)
		**out = **in
	}
	if in.ExternalIDSecretRef != nil {
		in, out := &in.ExternalIDSecretRef, &out.ExternalIDSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AwsParameters.
func (in *AwsParameters) DeepCopy() *AwsParameters {
	if in == nil {
		return nil
	}
	out := new(AwsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AwsSpec) DeepCopyInto(out *AwsSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AwsSpec.
func (in *AwsSpec) DeepCopy() *AwsSpec {
	if in == nil {
		return nil
	}
	out := new(AwsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AwsStatus) DeepCopyInto(out *AwsStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AwsStatus.
func (in *AwsStatus) DeepCopy() *AwsStatus {
	if in == nil {
		return nil
	}
	out := new(AwsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Azure) DeepCopyInto(out *Azure) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Azure.
func (in *Azure) DeepCopy() *Azure {
	if in == nil {
		return nil
	}
	out := new(Azure)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Azure) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureList) DeepCopyInto(out *AzureList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Azure, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureList.
func (in *AzureList) DeepCopy() *AzureList {
	if in == nil {
		return nil
	}
	out := new(AzureList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AzureList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureObservation) DeepCopyInto(out *AzureObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureObservation.
func (in *AzureObservation) DeepCopy() *AzureObservation {
	if in == nil {
		return nil
	}
	out := new(AzureObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureParameters) DeepCopyInto(out *AzureParameters) {
	*out = *in
	if in.AzureClientID != nil {
		in, out := &in.AzureClientID, &out.AzureClientID
		*out = new(string)
		**out = **in
	}
	out.AzureClientSecretSecretRef = in.AzureClientSecretSecretRef
	if in.AzureTenantID != nil {
		in, out := &in.AzureTenantID, &out.AzureTenantID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureParameters.
func (in *AzureParameters) DeepCopy() *AzureParameters {
	if in == nil {
		return nil
	}
	out := new(AzureParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureSpec) DeepCopyInto(out *AzureSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureSpec.
func (in *AzureSpec) DeepCopy() *AzureSpec {
	if in == nil {
		return nil
	}
	out := new(AzureSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureStatus) DeepCopyInto(out *AzureStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureStatus.
func (in *AzureStatus) DeepCopy() *AzureStatus {
	if in == nil {
		return nil
	}
	out := new(AzureStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCP) DeepCopyInto(out *GCP) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCP.
func (in *GCP) DeepCopy() *GCP {
	if in == nil {
		return nil
	}
	out := new(GCP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GCP) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCPList) DeepCopyInto(out *GCPList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GCP, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCPList.
func (in *GCPList) DeepCopy() *GCPList {
	if in == nil {
		return nil
	}
	out := new(GCPList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GCPList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCPObservation) DeepCopyInto(out *GCPObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCPObservation.
func (in *GCPObservation) DeepCopy() *GCPObservation {
	if in == nil {
		return nil
	}
	out := new(GCPObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCPParameters) DeepCopyInto(out *GCPParameters) {
	*out = *in
	out.GCPJSONCredentialsSecretRef = in.GCPJSONCredentialsSecretRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCPParameters.
func (in *GCPParameters) DeepCopy() *GCPParameters {
	if in == nil {
		return nil
	}
	out := new(GCPParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCPSpec) DeepCopyInto(out *GCPSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCPSpec.
func (in *GCPSpec) DeepCopy() *GCPSpec {
	if in == nil {
		return nil
	}
	out := new(GCPSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCPStatus) DeepCopyInto(out *GCPStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCPStatus.
func (in *GCPStatus) DeepCopy() *GCPStatus {
	if in == nil {
		return nil
	}
	out := new(GCPStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Maas) DeepCopyInto(out *Maas) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Maas.
func (in *Maas) DeepCopy() *Maas {
	if in == nil {
		return nil
	}
	out := new(Maas)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Maas) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaasList) DeepCopyInto(out *MaasList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Maas, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaasList.
func (in *MaasList) DeepCopy() *MaasList {
	if in == nil {
		return nil
	}
	out := new(MaasList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MaasList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaasObservation) DeepCopyInto(out *MaasObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaasObservation.
func (in *MaasObservation) DeepCopy() *MaasObservation {
	if in == nil {
		return nil
	}
	out := new(MaasObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaasParameters) DeepCopyInto(out *MaasParameters) {
	*out = *in
	if in.MaasAPIEndpoint != nil {
		in, out := &in.MaasAPIEndpoint, &out.MaasAPIEndpoint
		*out = new(string)
		**out = **in
	}
	if in.MaasAPIKeySecretRef != nil {
		in, out := &in.MaasAPIKeySecretRef, &out.MaasAPIKeySecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.PrivateCloudGatewayID != nil {
		in, out := &in.PrivateCloudGatewayID, &out.PrivateCloudGatewayID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaasParameters.
func (in *MaasParameters) DeepCopy() *MaasParameters {
	if in == nil {
		return nil
	}
	out := new(MaasParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaasSpec) DeepCopyInto(out *MaasSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaasSpec.
func (in *MaasSpec) DeepCopy() *MaasSpec {
	if in == nil {
		return nil
	}
	out := new(MaasSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaasStatus) DeepCopyInto(out *MaasStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaasStatus.
func (in *MaasStatus) DeepCopy() *MaasStatus {
	if in == nil {
		return nil
	}
	out := new(MaasStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Openstack) DeepCopyInto(out *Openstack) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Openstack.
func (in *Openstack) DeepCopy() *Openstack {
	if in == nil {
		return nil
	}
	out := new(Openstack)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Openstack) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenstackList) DeepCopyInto(out *OpenstackList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Openstack, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenstackList.
func (in *OpenstackList) DeepCopy() *OpenstackList {
	if in == nil {
		return nil
	}
	out := new(OpenstackList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenstackList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenstackObservation) DeepCopyInto(out *OpenstackObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenstackObservation.
func (in *OpenstackObservation) DeepCopy() *OpenstackObservation {
	if in == nil {
		return nil
	}
	out := new(OpenstackObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenstackParameters) DeepCopyInto(out *OpenstackParameters) {
	*out = *in
	if in.CACertificate != nil {
		in, out := &in.CACertificate, &out.CACertificate
		*out = new(string)
		**out = **in
	}
	if in.DefaultDomain != nil {
		in, out := &in.DefaultDomain, &out.DefaultDomain
		*out = new(string)
		**out = **in
	}
	if in.DefaultProject != nil {
		in, out := &in.DefaultProject, &out.DefaultProject
		*out = new(string)
		**out = **in
	}
	if in.IdentityEndpoint != nil {
		in, out := &in.IdentityEndpoint, &out.IdentityEndpoint
		*out = new(string)
		**out = **in
	}
	if in.OpenstackAllowInsecure != nil {
		in, out := &in.OpenstackAllowInsecure, &out.OpenstackAllowInsecure
		*out = new(bool)
		**out = **in
	}
	out.OpenstackPasswordSecretRef = in.OpenstackPasswordSecretRef
	if in.OpenstackUsername != nil {
		in, out := &in.OpenstackUsername, &out.OpenstackUsername
		*out = new(string)
		**out = **in
	}
	if in.ParentRegion != nil {
		in, out := &in.ParentRegion, &out.ParentRegion
		*out = new(string)
		**out = **in
	}
	if in.PrivateCloudGatewayID != nil {
		in, out := &in.PrivateCloudGatewayID, &out.PrivateCloudGatewayID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenstackParameters.
func (in *OpenstackParameters) DeepCopy() *OpenstackParameters {
	if in == nil {
		return nil
	}
	out := new(OpenstackParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenstackSpec) DeepCopyInto(out *OpenstackSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenstackSpec.
func (in *OpenstackSpec) DeepCopy() *OpenstackSpec {
	if in == nil {
		return nil
	}
	out := new(OpenstackSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenstackStatus) DeepCopyInto(out *OpenstackStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenstackStatus.
func (in *OpenstackStatus) DeepCopy() *OpenstackStatus {
	if in == nil {
		return nil
	}
	out := new(OpenstackStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Tencent) DeepCopyInto(out *Tencent) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tencent.
func (in *Tencent) DeepCopy() *Tencent {
	if in == nil {
		return nil
	}
	out := new(Tencent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Tencent) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TencentList) DeepCopyInto(out *TencentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Tencent, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TencentList.
func (in *TencentList) DeepCopy() *TencentList {
	if in == nil {
		return nil
	}
	out := new(TencentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TencentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TencentObservation) DeepCopyInto(out *TencentObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TencentObservation.
func (in *TencentObservation) DeepCopy() *TencentObservation {
	if in == nil {
		return nil
	}
	out := new(TencentObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TencentParameters) DeepCopyInto(out *TencentParameters) {
	*out = *in
	if in.TencentSecretID != nil {
		in, out := &in.TencentSecretID, &out.TencentSecretID
		*out = new(string)
		**out = **in
	}
	if in.TencentSecretKeySecretRef != nil {
		in, out := &in.TencentSecretKeySecretRef, &out.TencentSecretKeySecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TencentParameters.
func (in *TencentParameters) DeepCopy() *TencentParameters {
	if in == nil {
		return nil
	}
	out := new(TencentParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TencentSpec) DeepCopyInto(out *TencentSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TencentSpec.
func (in *TencentSpec) DeepCopy() *TencentSpec {
	if in == nil {
		return nil
	}
	out := new(TencentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TencentStatus) DeepCopyInto(out *TencentStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TencentStatus.
func (in *TencentStatus) DeepCopy() *TencentStatus {
	if in == nil {
		return nil
	}
	out := new(TencentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Vsphere) DeepCopyInto(out *Vsphere) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Vsphere.
func (in *Vsphere) DeepCopy() *Vsphere {
	if in == nil {
		return nil
	}
	out := new(Vsphere)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Vsphere) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VsphereList) DeepCopyInto(out *VsphereList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Vsphere, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VsphereList.
func (in *VsphereList) DeepCopy() *VsphereList {
	if in == nil {
		return nil
	}
	out := new(VsphereList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VsphereList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VsphereObservation) DeepCopyInto(out *VsphereObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VsphereObservation.
func (in *VsphereObservation) DeepCopy() *VsphereObservation {
	if in == nil {
		return nil
	}
	out := new(VsphereObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VsphereParameters) DeepCopyInto(out *VsphereParameters) {
	*out = *in
	if in.Context != nil {
		in, out := &in.Context, &out.Context
		*out = new(string)
		**out = **in
	}
	if in.PrivateCloudGatewayID != nil {
		in, out := &in.PrivateCloudGatewayID, &out.PrivateCloudGatewayID
		*out = new(string)
		**out = **in
	}
	if in.VsphereIgnoreInsecureError != nil {
		in, out := &in.VsphereIgnoreInsecureError, &out.VsphereIgnoreInsecureError
		*out = new(bool)
		**out = **in
	}
	out.VspherePasswordSecretRef = in.VspherePasswordSecretRef
	if in.VsphereUsername != nil {
		in, out := &in.VsphereUsername, &out.VsphereUsername
		*out = new(string)
		**out = **in
	}
	if in.VsphereVcenter != nil {
		in, out := &in.VsphereVcenter, &out.VsphereVcenter
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VsphereParameters.
func (in *VsphereParameters) DeepCopy() *VsphereParameters {
	if in == nil {
		return nil
	}
	out := new(VsphereParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VsphereSpec) DeepCopyInto(out *VsphereSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VsphereSpec.
func (in *VsphereSpec) DeepCopy() *VsphereSpec {
	if in == nil {
		return nil
	}
	out := new(VsphereSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VsphereStatus) DeepCopyInto(out *VsphereStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VsphereStatus.
func (in *VsphereStatus) DeepCopy() *VsphereStatus {
	if in == nil {
		return nil
	}
	out := new(VsphereStatus)
	in.DeepCopyInto(out)
	return out
}
