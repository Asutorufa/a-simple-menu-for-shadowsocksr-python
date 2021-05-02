// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.8
// source: internal/api/api.proto

package api

import (
	config "github.com/Asutorufa/yuhaiin/internal/config"
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type DaUaDrUr struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Download string `protobuf:"bytes,1,opt,name=Download,proto3" json:"Download,omitempty"`
	Upload   string `protobuf:"bytes,2,opt,name=Upload,proto3" json:"Upload,omitempty"`
	DownRate string `protobuf:"bytes,3,opt,name=DownRate,proto3" json:"DownRate,omitempty"`
	UpRate   string `protobuf:"bytes,4,opt,name=UpRate,proto3" json:"UpRate,omitempty"`
}

func (x *DaUaDrUr) Reset() {
	*x = DaUaDrUr{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_api_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DaUaDrUr) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DaUaDrUr) ProtoMessage() {}

func (x *DaUaDrUr) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DaUaDrUr.ProtoReflect.Descriptor instead.
func (*DaUaDrUr) Descriptor() ([]byte, []int) {
	return file_internal_api_api_proto_rawDescGZIP(), []int{0}
}

func (x *DaUaDrUr) GetDownload() string {
	if x != nil {
		return x.Download
	}
	return ""
}

func (x *DaUaDrUr) GetUpload() string {
	if x != nil {
		return x.Upload
	}
	return ""
}

func (x *DaUaDrUr) GetDownRate() string {
	if x != nil {
		return x.DownRate
	}
	return ""
}

func (x *DaUaDrUr) GetUpRate() string {
	if x != nil {
		return x.UpRate
	}
	return ""
}

type NodeMap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value map[string]string `protobuf:"bytes,1,rep,name=Value,proto3" json:"Value,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *NodeMap) Reset() {
	*x = NodeMap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_api_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeMap) ProtoMessage() {}

func (x *NodeMap) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeMap.ProtoReflect.Descriptor instead.
func (*NodeMap) Descriptor() ([]byte, []int) {
	return file_internal_api_api_proto_rawDescGZIP(), []int{1}
}

func (x *NodeMap) GetValue() map[string]string {
	if x != nil {
		return x.Value
	}
	return nil
}

type Nodes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value map[string]*AllGroupOrNode `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Nodes) Reset() {
	*x = Nodes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_api_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Nodes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Nodes) ProtoMessage() {}

func (x *Nodes) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Nodes.ProtoReflect.Descriptor instead.
func (*Nodes) Descriptor() ([]byte, []int) {
	return file_internal_api_api_proto_rawDescGZIP(), []int{2}
}

func (x *Nodes) GetValue() map[string]*AllGroupOrNode {
	if x != nil {
		return x.Value
	}
	return nil
}

type AllGroupOrNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []string `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty"`
}

func (x *AllGroupOrNode) Reset() {
	*x = AllGroupOrNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_api_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllGroupOrNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllGroupOrNode) ProtoMessage() {}

func (x *AllGroupOrNode) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllGroupOrNode.ProtoReflect.Descriptor instead.
func (*AllGroupOrNode) Descriptor() ([]byte, []int) {
	return file_internal_api_api_proto_rawDescGZIP(), []int{3}
}

func (x *AllGroupOrNode) GetValue() []string {
	if x != nil {
		return x.Value
	}
	return nil
}

type GroupAndNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Group string `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
	Node  string `protobuf:"bytes,2,opt,name=node,proto3" json:"node,omitempty"`
}

func (x *GroupAndNode) Reset() {
	*x = GroupAndNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_api_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupAndNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupAndNode) ProtoMessage() {}

func (x *GroupAndNode) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupAndNode.ProtoReflect.Descriptor instead.
func (*GroupAndNode) Descriptor() ([]byte, []int) {
	return file_internal_api_api_proto_rawDescGZIP(), []int{4}
}

func (x *GroupAndNode) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *GroupAndNode) GetNode() string {
	if x != nil {
		return x.Node
	}
	return ""
}

type Link struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Url  string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Type string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *Link) Reset() {
	*x = Link{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_api_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Link) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Link) ProtoMessage() {}

func (x *Link) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Link.ProtoReflect.Descriptor instead.
func (*Link) Descriptor() ([]byte, []int) {
	return file_internal_api_api_proto_rawDescGZIP(), []int{5}
}

func (x *Link) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Link) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Link) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type Links struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value map[string]*Link `protobuf:"bytes,1,rep,name=Value,proto3" json:"Value,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Links) Reset() {
	*x = Links{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_api_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Links) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Links) ProtoMessage() {}

func (x *Links) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Links.ProtoReflect.Descriptor instead.
func (*Links) Descriptor() ([]byte, []int) {
	return file_internal_api_api_proto_rawDescGZIP(), []int{6}
}

func (x *Links) GetValue() map[string]*Link {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_internal_api_api_proto protoreflect.FileDescriptor

var file_internal_api_api_proto_rawDesc = []byte{
	0x0a, 0x16, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x79, 0x75, 0x68, 0x61, 0x69, 0x69,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x72, 0x0a, 0x08, 0x44, 0x61, 0x55, 0x61, 0x44, 0x72, 0x55, 0x72, 0x12, 0x1a, 0x0a, 0x08,
	0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x44, 0x6f, 0x77, 0x6e, 0x52, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x44, 0x6f, 0x77, 0x6e, 0x52, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x55, 0x70, 0x52, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x70,
	0x52, 0x61, 0x74, 0x65, 0x22, 0x7a, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x4d, 0x61, 0x70, 0x12,
	0x35, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x79, 0x75, 0x68, 0x61, 0x69, 0x69, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x6f, 0x64,
	0x65, 0x4d, 0x61, 0x70, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x38, 0x0a, 0x0a, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x93, 0x01, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x33, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x79, 0x75, 0x68, 0x61,
	0x69, 0x69, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x2e, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x1a,
	0x55, 0x0a, 0x0a, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x31, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x79, 0x75, 0x68, 0x61, 0x69, 0x69, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x6c, 0x6c,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x4f, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x26, 0x0a, 0x0e, 0x61, 0x6c, 0x6c, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x4f, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x38,
	0x0a, 0x0c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x6e, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x22, 0x40, 0x0a, 0x04, 0x4c, 0x69, 0x6e, 0x6b,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x89, 0x01, 0x0a, 0x05, 0x4c,
	0x69, 0x6e, 0x6b, 0x73, 0x12, 0x33, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x79, 0x75, 0x68, 0x61, 0x69, 0x69, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x4b, 0x0a, 0x0a, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x27, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x79, 0x75, 0x68, 0x61, 0x69,
	0x69, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0xa7, 0x04, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x49, 0x6e, 0x69, 0x74, 0x12, 0x40, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4c, 0x6f, 0x63, 0x6b, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3d, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x49, 0x6e, 0x69, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x46, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x52, 0x75,
	0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x3a, 0x0a, 0x08, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4f, 0x6e, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3d, 0x0a, 0x0b, 0x50,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x45, 0x78, 0x69, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x44, 0x0a, 0x0c, 0x47, 0x65,
	0x74, 0x4b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x50, 0x69, 0x64, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x3c, 0x0a, 0x0a, 0x53, 0x74, 0x6f, 0x70, 0x4b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x12, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x50,
	0x0a, 0x0e, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x28, 0x01, 0x30, 0x01,
	0x32, 0xfa, 0x01, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x39, 0x0a, 0x09, 0x47,
	0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x14, 0x2e, 0x79, 0x75, 0x68, 0x61, 0x69, 0x69, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x39, 0x0a, 0x09, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x14, 0x2e, 0x79, 0x75, 0x68, 0x61, 0x69, 0x69, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x12, 0x3e, 0x0a, 0x0c, 0x52, 0x65, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x75, 0x6c,
	0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x12, 0x3a, 0x0a, 0x07, 0x67, 0x65, 0x74, 0x52, 0x61, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x15, 0x2e, 0x79, 0x75, 0x68, 0x61, 0x69, 0x69, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x44, 0x61, 0x55, 0x61, 0x44, 0x72, 0x55, 0x72, 0x30, 0x01, 0x32, 0xcc, 0x04,
	0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x36, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64,
	0x65, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x12, 0x2e, 0x79, 0x75, 0x68,
	0x61, 0x69, 0x69, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x3f,
	0x0a, 0x08, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x1b, 0x2e, 0x79, 0x75, 0x68, 0x61, 0x69, 0x69, 0x6e, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x61, 0x6c, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4f, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x12,
	0x44, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x1b, 0x2e, 0x79, 0x75, 0x68, 0x61, 0x69,
	0x69, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x6c, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4f,
	0x72, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x47, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x77, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x41, 0x6e, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x19, 0x2e, 0x79, 0x75, 0x68, 0x61, 0x69, 0x69, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x6e, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x42,
	0x0a, 0x0d, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4e, 0x6f, 0x77, 0x4e, 0x6f, 0x64, 0x65, 0x12,
	0x19, 0x2e, 0x79, 0x75, 0x68, 0x61, 0x69, 0x69, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x41, 0x6e, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x12, 0x37, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x2e,
	0x79, 0x75, 0x68, 0x61, 0x69, 0x69, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x6f, 0x64, 0x65,
	0x4d, 0x61, 0x70, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3a, 0x0a, 0x0a, 0x4d,
	0x6f, 0x64, 0x69, 0x66, 0x79, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x2e, 0x79, 0x75, 0x68, 0x61,
	0x69, 0x69, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x4d, 0x61, 0x70, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3f, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x19, 0x2e, 0x79, 0x75, 0x68, 0x61, 0x69, 0x69, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x6e, 0x64, 0x4e, 0x6f, 0x64, 0x65,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x42, 0x0a, 0x07, 0x4c, 0x61, 0x74, 0x65,
	0x6e, 0x63, 0x79, 0x12, 0x19, 0x2e, 0x79, 0x75, 0x68, 0x61, 0x69, 0x69, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x6e, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x1a, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x32, 0xfb, 0x01, 0x0a,
	0x09, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x3b, 0x0a, 0x09, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x39, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x53, 0x75,
	0x62, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x12,
	0x2e, 0x79, 0x75, 0x68, 0x61, 0x69, 0x69, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x6e,
	0x6b, 0x73, 0x12, 0x33, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x53, 0x75, 0x62, 0x4c, 0x69, 0x6e, 0x6b,
	0x12, 0x11, 0x2e, 0x79, 0x75, 0x68, 0x61, 0x69, 0x69, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c,
	0x69, 0x6e, 0x6b, 0x1a, 0x12, 0x2e, 0x79, 0x75, 0x68, 0x61, 0x69, 0x69, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x12, 0x41, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x53, 0x75, 0x62, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x12, 0x2e, 0x79, 0x75, 0x68, 0x61, 0x69, 0x69, 0x6e,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x42, 0x22, 0x5a, 0x20, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x73, 0x75, 0x74, 0x6f, 0x72, 0x75,
	0x66, 0x61, 0x2f, 0x79, 0x75, 0x68, 0x61, 0x69, 0x69, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_api_api_proto_rawDescOnce sync.Once
	file_internal_api_api_proto_rawDescData = file_internal_api_api_proto_rawDesc
)

func file_internal_api_api_proto_rawDescGZIP() []byte {
	file_internal_api_api_proto_rawDescOnce.Do(func() {
		file_internal_api_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_api_api_proto_rawDescData)
	})
	return file_internal_api_api_proto_rawDescData
}

var file_internal_api_api_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_internal_api_api_proto_goTypes = []interface{}{
	(*DaUaDrUr)(nil),               // 0: yuhaiin.api.DaUaDrUr
	(*NodeMap)(nil),                // 1: yuhaiin.api.nodeMap
	(*Nodes)(nil),                  // 2: yuhaiin.api.nodes
	(*AllGroupOrNode)(nil),         // 3: yuhaiin.api.allGroupOrNode
	(*GroupAndNode)(nil),           // 4: yuhaiin.api.GroupAndNode
	(*Link)(nil),                   // 5: yuhaiin.api.Link
	(*Links)(nil),                  // 6: yuhaiin.api.Links
	nil,                            // 7: yuhaiin.api.nodeMap.ValueEntry
	nil,                            // 8: yuhaiin.api.nodes.ValueEntry
	nil,                            // 9: yuhaiin.api.Links.ValueEntry
	(*emptypb.Empty)(nil),          // 10: google.protobuf.Empty
	(*wrapperspb.StringValue)(nil), // 11: google.protobuf.StringValue
	(*config.Setting)(nil),         // 12: yuhaiin.api.Setting
	(*wrapperspb.UInt32Value)(nil), // 13: google.protobuf.UInt32Value
}
var file_internal_api_api_proto_depIdxs = []int32{
	7,  // 0: yuhaiin.api.nodeMap.Value:type_name -> yuhaiin.api.nodeMap.ValueEntry
	8,  // 1: yuhaiin.api.nodes.value:type_name -> yuhaiin.api.nodes.ValueEntry
	9,  // 2: yuhaiin.api.Links.Value:type_name -> yuhaiin.api.Links.ValueEntry
	3,  // 3: yuhaiin.api.nodes.ValueEntry.value:type_name -> yuhaiin.api.allGroupOrNode
	5,  // 4: yuhaiin.api.Links.ValueEntry.value:type_name -> yuhaiin.api.Link
	10, // 5: yuhaiin.api.processInit.CreateLockFile:input_type -> google.protobuf.Empty
	10, // 6: yuhaiin.api.processInit.ProcessInit:input_type -> google.protobuf.Empty
	10, // 7: yuhaiin.api.processInit.GetRunningHost:input_type -> google.protobuf.Empty
	10, // 8: yuhaiin.api.processInit.ClientOn:input_type -> google.protobuf.Empty
	10, // 9: yuhaiin.api.processInit.ProcessExit:input_type -> google.protobuf.Empty
	10, // 10: yuhaiin.api.processInit.GetKernelPid:input_type -> google.protobuf.Empty
	10, // 11: yuhaiin.api.processInit.StopKernel:input_type -> google.protobuf.Empty
	11, // 12: yuhaiin.api.processInit.SingleInstance:input_type -> google.protobuf.StringValue
	10, // 13: yuhaiin.api.config.GetConfig:input_type -> google.protobuf.Empty
	12, // 14: yuhaiin.api.config.SetConfig:input_type -> yuhaiin.api.Setting
	10, // 15: yuhaiin.api.config.ReimportRule:input_type -> google.protobuf.Empty
	10, // 16: yuhaiin.api.config.getRate:input_type -> google.protobuf.Empty
	10, // 17: yuhaiin.api.Node.GetNodes:input_type -> google.protobuf.Empty
	10, // 18: yuhaiin.api.Node.GetGroup:input_type -> google.protobuf.Empty
	11, // 19: yuhaiin.api.Node.GetNode:input_type -> google.protobuf.StringValue
	10, // 20: yuhaiin.api.Node.GetNowGroupAndName:input_type -> google.protobuf.Empty
	4,  // 21: yuhaiin.api.Node.ChangeNowNode:input_type -> yuhaiin.api.GroupAndNode
	1,  // 22: yuhaiin.api.Node.AddNode:input_type -> yuhaiin.api.nodeMap
	1,  // 23: yuhaiin.api.Node.ModifyNode:input_type -> yuhaiin.api.nodeMap
	4,  // 24: yuhaiin.api.Node.DeleteNode:input_type -> yuhaiin.api.GroupAndNode
	4,  // 25: yuhaiin.api.Node.Latency:input_type -> yuhaiin.api.GroupAndNode
	10, // 26: yuhaiin.api.Subscribe.UpdateSub:input_type -> google.protobuf.Empty
	10, // 27: yuhaiin.api.Subscribe.GetSubLinks:input_type -> google.protobuf.Empty
	5,  // 28: yuhaiin.api.Subscribe.AddSubLink:input_type -> yuhaiin.api.Link
	11, // 29: yuhaiin.api.Subscribe.DeleteSubLink:input_type -> google.protobuf.StringValue
	10, // 30: yuhaiin.api.processInit.CreateLockFile:output_type -> google.protobuf.Empty
	10, // 31: yuhaiin.api.processInit.ProcessInit:output_type -> google.protobuf.Empty
	11, // 32: yuhaiin.api.processInit.GetRunningHost:output_type -> google.protobuf.StringValue
	10, // 33: yuhaiin.api.processInit.ClientOn:output_type -> google.protobuf.Empty
	10, // 34: yuhaiin.api.processInit.ProcessExit:output_type -> google.protobuf.Empty
	13, // 35: yuhaiin.api.processInit.GetKernelPid:output_type -> google.protobuf.UInt32Value
	10, // 36: yuhaiin.api.processInit.StopKernel:output_type -> google.protobuf.Empty
	11, // 37: yuhaiin.api.processInit.SingleInstance:output_type -> google.protobuf.StringValue
	12, // 38: yuhaiin.api.config.GetConfig:output_type -> yuhaiin.api.Setting
	10, // 39: yuhaiin.api.config.SetConfig:output_type -> google.protobuf.Empty
	10, // 40: yuhaiin.api.config.ReimportRule:output_type -> google.protobuf.Empty
	0,  // 41: yuhaiin.api.config.getRate:output_type -> yuhaiin.api.DaUaDrUr
	2,  // 42: yuhaiin.api.Node.GetNodes:output_type -> yuhaiin.api.nodes
	3,  // 43: yuhaiin.api.Node.GetGroup:output_type -> yuhaiin.api.allGroupOrNode
	3,  // 44: yuhaiin.api.Node.GetNode:output_type -> yuhaiin.api.allGroupOrNode
	4,  // 45: yuhaiin.api.Node.GetNowGroupAndName:output_type -> yuhaiin.api.GroupAndNode
	10, // 46: yuhaiin.api.Node.ChangeNowNode:output_type -> google.protobuf.Empty
	10, // 47: yuhaiin.api.Node.AddNode:output_type -> google.protobuf.Empty
	10, // 48: yuhaiin.api.Node.ModifyNode:output_type -> google.protobuf.Empty
	10, // 49: yuhaiin.api.Node.DeleteNode:output_type -> google.protobuf.Empty
	11, // 50: yuhaiin.api.Node.Latency:output_type -> google.protobuf.StringValue
	10, // 51: yuhaiin.api.Subscribe.UpdateSub:output_type -> google.protobuf.Empty
	6,  // 52: yuhaiin.api.Subscribe.GetSubLinks:output_type -> yuhaiin.api.Links
	6,  // 53: yuhaiin.api.Subscribe.AddSubLink:output_type -> yuhaiin.api.Links
	6,  // 54: yuhaiin.api.Subscribe.DeleteSubLink:output_type -> yuhaiin.api.Links
	30, // [30:55] is the sub-list for method output_type
	5,  // [5:30] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_internal_api_api_proto_init() }
func file_internal_api_api_proto_init() {
	if File_internal_api_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_api_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DaUaDrUr); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_api_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeMap); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_api_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Nodes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_api_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllGroupOrNode); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_api_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupAndNode); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_api_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Link); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_api_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Links); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_api_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   4,
		},
		GoTypes:           file_internal_api_api_proto_goTypes,
		DependencyIndexes: file_internal_api_api_proto_depIdxs,
		MessageInfos:      file_internal_api_api_proto_msgTypes,
	}.Build()
	File_internal_api_api_proto = out.File
	file_internal_api_api_proto_rawDesc = nil
	file_internal_api_api_proto_goTypes = nil
	file_internal_api_api_proto_depIdxs = nil
}
