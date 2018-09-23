// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaServerMongoDB) DeepCopyInto(out *PerconaServerMongoDB) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaServerMongoDB.
func (in *PerconaServerMongoDB) DeepCopy() *PerconaServerMongoDB {
	if in == nil {
		return nil
	}
	out := new(PerconaServerMongoDB)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PerconaServerMongoDB) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaServerMongoDBList) DeepCopyInto(out *PerconaServerMongoDBList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PerconaServerMongoDB, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaServerMongoDBList.
func (in *PerconaServerMongoDBList) DeepCopy() *PerconaServerMongoDBList {
	if in == nil {
		return nil
	}
	out := new(PerconaServerMongoDBList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PerconaServerMongoDBList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaServerMongoDBSpec) DeepCopyInto(out *PerconaServerMongoDBSpec) {
	*out = *in
	if in.MongoDB != nil {
		in, out := &in.MongoDB, &out.MongoDB
		*out = new(PerconaServerMongoDBSpecMongoDB)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaServerMongoDBSpec.
func (in *PerconaServerMongoDBSpec) DeepCopy() *PerconaServerMongoDBSpec {
	if in == nil {
		return nil
	}
	out := new(PerconaServerMongoDBSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaServerMongoDBSpecMongoDB) DeepCopyInto(out *PerconaServerMongoDBSpecMongoDB) {
	*out = *in
	if in.MMAPv1 != nil {
		in, out := &in.MMAPv1, &out.MMAPv1
		*out = new(PerconaServerMongoDBSpecMongoDBMMAPv1)
		**out = **in
	}
	if in.WiredTiger != nil {
		in, out := &in.WiredTiger, &out.WiredTiger
		*out = new(PerconaServerMongoDBSpecMongoDBWiredTiger)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaServerMongoDBSpecMongoDB.
func (in *PerconaServerMongoDBSpecMongoDB) DeepCopy() *PerconaServerMongoDBSpecMongoDB {
	if in == nil {
		return nil
	}
	out := new(PerconaServerMongoDBSpecMongoDB)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaServerMongoDBSpecMongoDBMMAPv1) DeepCopyInto(out *PerconaServerMongoDBSpecMongoDBMMAPv1) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaServerMongoDBSpecMongoDBMMAPv1.
func (in *PerconaServerMongoDBSpecMongoDBMMAPv1) DeepCopy() *PerconaServerMongoDBSpecMongoDBMMAPv1 {
	if in == nil {
		return nil
	}
	out := new(PerconaServerMongoDBSpecMongoDBMMAPv1)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaServerMongoDBSpecMongoDBWiredTiger) DeepCopyInto(out *PerconaServerMongoDBSpecMongoDBWiredTiger) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaServerMongoDBSpecMongoDBWiredTiger.
func (in *PerconaServerMongoDBSpecMongoDBWiredTiger) DeepCopy() *PerconaServerMongoDBSpecMongoDBWiredTiger {
	if in == nil {
		return nil
	}
	out := new(PerconaServerMongoDBSpecMongoDBWiredTiger)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaServerMongoDBStatus) DeepCopyInto(out *PerconaServerMongoDBStatus) {
	*out = *in
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaServerMongoDBStatus.
func (in *PerconaServerMongoDBStatus) DeepCopy() *PerconaServerMongoDBStatus {
	if in == nil {
		return nil
	}
	out := new(PerconaServerMongoDBStatus)
	in.DeepCopyInto(out)
	return out
}
