// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthenticationOperatorConfig) DeepCopyInto(out *AuthenticationOperatorConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthenticationOperatorConfig.
func (in *AuthenticationOperatorConfig) DeepCopy() *AuthenticationOperatorConfig {
	if in == nil {
		return nil
	}
	out := new(AuthenticationOperatorConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AuthenticationOperatorConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthenticationOperatorConfigList) DeepCopyInto(out *AuthenticationOperatorConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AuthenticationOperatorConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthenticationOperatorConfigList.
func (in *AuthenticationOperatorConfigList) DeepCopy() *AuthenticationOperatorConfigList {
	if in == nil {
		return nil
	}
	out := new(AuthenticationOperatorConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AuthenticationOperatorConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthenticationOperatorConfigSpec) DeepCopyInto(out *AuthenticationOperatorConfigSpec) {
	*out = *in
	in.OperatorSpec.DeepCopyInto(&out.OperatorSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthenticationOperatorConfigSpec.
func (in *AuthenticationOperatorConfigSpec) DeepCopy() *AuthenticationOperatorConfigSpec {
	if in == nil {
		return nil
	}
	out := new(AuthenticationOperatorConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthenticationOperatorConfigStatus) DeepCopyInto(out *AuthenticationOperatorConfigStatus) {
	*out = *in
	in.OperatorStatus.DeepCopyInto(&out.OperatorStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthenticationOperatorConfigStatus.
func (in *AuthenticationOperatorConfigStatus) DeepCopy() *AuthenticationOperatorConfigStatus {
	if in == nil {
		return nil
	}
	out := new(AuthenticationOperatorConfigStatus)
	in.DeepCopyInto(out)
	return out
}
