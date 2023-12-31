//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023.

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

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectionProperties) DeepCopyInto(out *ConnectionProperties) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionProperties.
func (in *ConnectionProperties) DeepCopy() *ConnectionProperties {
	if in == nil {
		return nil
	}
	out := new(ConnectionProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DremioRestServer) DeepCopyInto(out *DremioRestServer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DremioRestServer.
func (in *DremioRestServer) DeepCopy() *DremioRestServer {
	if in == nil {
		return nil
	}
	out := new(DremioRestServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DremioRestServer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DremioRestServerList) DeepCopyInto(out *DremioRestServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DremioRestServer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DremioRestServerList.
func (in *DremioRestServerList) DeepCopy() *DremioRestServerList {
	if in == nil {
		return nil
	}
	out := new(DremioRestServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DremioRestServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DremioRestServerSpec) DeepCopyInto(out *DremioRestServerSpec) {
	*out = *in
	out.Connection = in.Connection
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DremioRestServerSpec.
func (in *DremioRestServerSpec) DeepCopy() *DremioRestServerSpec {
	if in == nil {
		return nil
	}
	out := new(DremioRestServerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DremioRestServerStatus) DeepCopyInto(out *DremioRestServerStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DremioRestServerStatus.
func (in *DremioRestServerStatus) DeepCopy() *DremioRestServerStatus {
	if in == nil {
		return nil
	}
	out := new(DremioRestServerStatus)
	in.DeepCopyInto(out)
	return out
}
