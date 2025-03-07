// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: examples/internal/proto/examplepb/openapi_merge_a.proto

// Merging Services
//
// This is an example of merging two proto files.

package examplepb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// InMessageA represents a message to ServiceA and ServiceC.
type InMessageA struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Here is the explanation about InMessageA.values
	Values []string `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *InMessageA) Reset() {
	*x = InMessageA{}
	if protoimpl.UnsafeEnabled {
		mi := &file_examples_internal_proto_examplepb_openapi_merge_a_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InMessageA) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InMessageA) ProtoMessage() {}

func (x *InMessageA) ProtoReflect() protoreflect.Message {
	mi := &file_examples_internal_proto_examplepb_openapi_merge_a_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InMessageA.ProtoReflect.Descriptor instead.
func (*InMessageA) Descriptor() ([]byte, []int) {
	return file_examples_internal_proto_examplepb_openapi_merge_a_proto_rawDescGZIP(), []int{0}
}

func (x *InMessageA) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

// OutMessageA represents a message returned from ServiceA.
type OutMessageA struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Here is the explanation about OutMessageA.value
	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *OutMessageA) Reset() {
	*x = OutMessageA{}
	if protoimpl.UnsafeEnabled {
		mi := &file_examples_internal_proto_examplepb_openapi_merge_a_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutMessageA) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutMessageA) ProtoMessage() {}

func (x *OutMessageA) ProtoReflect() protoreflect.Message {
	mi := &file_examples_internal_proto_examplepb_openapi_merge_a_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutMessageA.ProtoReflect.Descriptor instead.
func (*OutMessageA) Descriptor() ([]byte, []int) {
	return file_examples_internal_proto_examplepb_openapi_merge_a_proto_rawDescGZIP(), []int{1}
}

func (x *OutMessageA) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// OutMessageC represents a message returned from ServiceC.
type OutMessageC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Here is the explanation about OutMessageC.value
	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *OutMessageC) Reset() {
	*x = OutMessageC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_examples_internal_proto_examplepb_openapi_merge_a_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutMessageC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutMessageC) ProtoMessage() {}

func (x *OutMessageC) ProtoReflect() protoreflect.Message {
	mi := &file_examples_internal_proto_examplepb_openapi_merge_a_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutMessageC.ProtoReflect.Descriptor instead.
func (*OutMessageC) Descriptor() ([]byte, []int) {
	return file_examples_internal_proto_examplepb_openapi_merge_a_proto_rawDescGZIP(), []int{2}
}

func (x *OutMessageC) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_examples_internal_proto_examplepb_openapi_merge_a_proto protoreflect.FileDescriptor

var file_examples_internal_proto_examplepb_openapi_merge_a_proto_rawDesc = []byte{
	0x0a, 0x37, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x70, 0x62, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x5f, 0x6d, 0x65, 0x72, 0x67,
	0x65, 0x5f, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x28, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73,
	0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x70, 0x62, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x24, 0x0a, 0x0a, 0x49, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x41, 0x12,
	0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x23, 0x0a, 0x0b, 0x4f, 0x75, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x41, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x23, 0x0a, 0x0b,
	0x4f, 0x75, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x32, 0xb8, 0x02, 0x0a, 0x08, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x12, 0x94,
	0x01, 0x0a, 0x09, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x6e, 0x65, 0x12, 0x34, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x41, 0x1a, 0x35, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x4f, 0x75,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x41, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x14, 0x3a, 0x01, 0x2a, 0x22, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2f, 0x61, 0x2f, 0x31, 0x12, 0x94, 0x01, 0x0a, 0x09, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x54, 0x77, 0x6f, 0x12, 0x35, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x4f,
	0x75, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x41, 0x1a, 0x34, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x41,
	0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x3a, 0x01, 0x2a, 0x22, 0x0f, 0x2f, 0x76, 0x31,
	0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x61, 0x2f, 0x32, 0x32, 0xb8, 0x02, 0x0a,
	0x08, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x12, 0x94, 0x01, 0x0a, 0x09, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x4f, 0x6e, 0x65, 0x12, 0x34, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x70, 0x62, 0x2e, 0x49, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x41, 0x1a, 0x35, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x4f, 0x75, 0x74, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x43, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x3a, 0x01, 0x2a, 0x22,
	0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x63, 0x2f, 0x31,
	0x12, 0x94, 0x01, 0x0a, 0x09, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x54, 0x77, 0x6f, 0x12, 0x35,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x4f, 0x75, 0x74, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x41, 0x1a, 0x34, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62,
	0x2e, 0x49, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x41, 0x22, 0x1a, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x14, 0x3a, 0x01, 0x2a, 0x22, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2f, 0x63, 0x2f, 0x32, 0x42, 0x49, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x76, 0x69, 0x64, 0x68, 0x6f, 0x6e, 0x67, 0x31,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x76, 0x32,
	0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_examples_internal_proto_examplepb_openapi_merge_a_proto_rawDescOnce sync.Once
	file_examples_internal_proto_examplepb_openapi_merge_a_proto_rawDescData = file_examples_internal_proto_examplepb_openapi_merge_a_proto_rawDesc
)

func file_examples_internal_proto_examplepb_openapi_merge_a_proto_rawDescGZIP() []byte {
	file_examples_internal_proto_examplepb_openapi_merge_a_proto_rawDescOnce.Do(func() {
		file_examples_internal_proto_examplepb_openapi_merge_a_proto_rawDescData = protoimpl.X.CompressGZIP(file_examples_internal_proto_examplepb_openapi_merge_a_proto_rawDescData)
	})
	return file_examples_internal_proto_examplepb_openapi_merge_a_proto_rawDescData
}

var file_examples_internal_proto_examplepb_openapi_merge_a_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_examples_internal_proto_examplepb_openapi_merge_a_proto_goTypes = []interface{}{
	(*InMessageA)(nil),  // 0: grpc.gateway.examples.internal.examplepb.InMessageA
	(*OutMessageA)(nil), // 1: grpc.gateway.examples.internal.examplepb.OutMessageA
	(*OutMessageC)(nil), // 2: grpc.gateway.examples.internal.examplepb.OutMessageC
}
var file_examples_internal_proto_examplepb_openapi_merge_a_proto_depIdxs = []int32{
	0, // 0: grpc.gateway.examples.internal.examplepb.ServiceA.MethodOne:input_type -> grpc.gateway.examples.internal.examplepb.InMessageA
	1, // 1: grpc.gateway.examples.internal.examplepb.ServiceA.MethodTwo:input_type -> grpc.gateway.examples.internal.examplepb.OutMessageA
	0, // 2: grpc.gateway.examples.internal.examplepb.ServiceC.MethodOne:input_type -> grpc.gateway.examples.internal.examplepb.InMessageA
	1, // 3: grpc.gateway.examples.internal.examplepb.ServiceC.MethodTwo:input_type -> grpc.gateway.examples.internal.examplepb.OutMessageA
	1, // 4: grpc.gateway.examples.internal.examplepb.ServiceA.MethodOne:output_type -> grpc.gateway.examples.internal.examplepb.OutMessageA
	0, // 5: grpc.gateway.examples.internal.examplepb.ServiceA.MethodTwo:output_type -> grpc.gateway.examples.internal.examplepb.InMessageA
	2, // 6: grpc.gateway.examples.internal.examplepb.ServiceC.MethodOne:output_type -> grpc.gateway.examples.internal.examplepb.OutMessageC
	0, // 7: grpc.gateway.examples.internal.examplepb.ServiceC.MethodTwo:output_type -> grpc.gateway.examples.internal.examplepb.InMessageA
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_examples_internal_proto_examplepb_openapi_merge_a_proto_init() }
func file_examples_internal_proto_examplepb_openapi_merge_a_proto_init() {
	if File_examples_internal_proto_examplepb_openapi_merge_a_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_examples_internal_proto_examplepb_openapi_merge_a_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InMessageA); i {
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
		file_examples_internal_proto_examplepb_openapi_merge_a_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutMessageA); i {
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
		file_examples_internal_proto_examplepb_openapi_merge_a_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutMessageC); i {
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
			RawDescriptor: file_examples_internal_proto_examplepb_openapi_merge_a_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_examples_internal_proto_examplepb_openapi_merge_a_proto_goTypes,
		DependencyIndexes: file_examples_internal_proto_examplepb_openapi_merge_a_proto_depIdxs,
		MessageInfos:      file_examples_internal_proto_examplepb_openapi_merge_a_proto_msgTypes,
	}.Build()
	File_examples_internal_proto_examplepb_openapi_merge_a_proto = out.File
	file_examples_internal_proto_examplepb_openapi_merge_a_proto_rawDesc = nil
	file_examples_internal_proto_examplepb_openapi_merge_a_proto_goTypes = nil
	file_examples_internal_proto_examplepb_openapi_merge_a_proto_depIdxs = nil
}
