// +build !ignore_autogenerated

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Envoy) DeepCopyInto(out *Envoy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Envoy.
func (in *Envoy) DeepCopy() *Envoy {
	if in == nil {
		return nil
	}
	out := new(Envoy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Envoy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvoyDeploymentSpec) DeepCopyInto(out *EnvoyDeploymentSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvoyDeploymentSpec.
func (in *EnvoyDeploymentSpec) DeepCopy() *EnvoyDeploymentSpec {
	if in == nil {
		return nil
	}
	out := new(EnvoyDeploymentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvoyList) DeepCopyInto(out *EnvoyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Envoy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvoyList.
func (in *EnvoyList) DeepCopy() *EnvoyList {
	if in == nil {
		return nil
	}
	out := new(EnvoyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EnvoyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvoySpec) DeepCopyInto(out *EnvoySpec) {
	*out = *in
	if in.ImageCommand != nil {
		in, out := &in.ImageCommand, &out.ImageCommand
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ServicePorts != nil {
		in, out := &in.ServicePorts, &out.ServicePorts
		*out = make([]int32, len(*in))
		copy(*out, *in)
	}
	if in.Deployment != nil {
		in, out := &in.Deployment, &out.Deployment
		if *in == nil {
			*out = nil
		} else {
			*out = new(EnvoyDeploymentSpec)
			**out = **in
		}
	}
	if in.Injection != nil {
		in, out := &in.Injection, &out.Injection
		if *in == nil {
			*out = nil
		} else {
			*out = new(InjectionSpec)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvoySpec.
func (in *EnvoySpec) DeepCopy() *EnvoySpec {
	if in == nil {
		return nil
	}
	out := new(EnvoySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvoyStatus) DeepCopyInto(out *EnvoyStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvoyStatus.
func (in *EnvoyStatus) DeepCopy() *EnvoyStatus {
	if in == nil {
		return nil
	}
	out := new(EnvoyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InjectionSpec) DeepCopyInto(out *InjectionSpec) {
	*out = *in
	if in.Namespaceslist != nil {
		in, out := &in.Namespaceslist, &out.Namespaceslist
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InjectionSpec.
func (in *InjectionSpec) DeepCopy() *InjectionSpec {
	if in == nil {
		return nil
	}
	out := new(InjectionSpec)
	in.DeepCopyInto(out)
	return out
}
