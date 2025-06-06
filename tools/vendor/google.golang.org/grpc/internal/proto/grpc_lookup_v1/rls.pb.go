// Copyright 2020 The gRPC Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.27.1
// source: grpc/lookup/v1/rls.proto

package grpc_lookup_v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Possible reasons for making a request.
type RouteLookupRequest_Reason int32

const (
	RouteLookupRequest_REASON_UNKNOWN RouteLookupRequest_Reason = 0 // Unused
	RouteLookupRequest_REASON_MISS    RouteLookupRequest_Reason = 1 // No data available in local cache
	RouteLookupRequest_REASON_STALE   RouteLookupRequest_Reason = 2 // Data in local cache is stale
)

// Enum value maps for RouteLookupRequest_Reason.
var (
	RouteLookupRequest_Reason_name = map[int32]string{
		0: "REASON_UNKNOWN",
		1: "REASON_MISS",
		2: "REASON_STALE",
	}
	RouteLookupRequest_Reason_value = map[string]int32{
		"REASON_UNKNOWN": 0,
		"REASON_MISS":    1,
		"REASON_STALE":   2,
	}
)

func (x RouteLookupRequest_Reason) Enum() *RouteLookupRequest_Reason {
	p := new(RouteLookupRequest_Reason)
	*p = x
	return p
}

func (x RouteLookupRequest_Reason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RouteLookupRequest_Reason) Descriptor() protoreflect.EnumDescriptor {
	return file_grpc_lookup_v1_rls_proto_enumTypes[0].Descriptor()
}

func (RouteLookupRequest_Reason) Type() protoreflect.EnumType {
	return &file_grpc_lookup_v1_rls_proto_enumTypes[0]
}

func (x RouteLookupRequest_Reason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RouteLookupRequest_Reason.Descriptor instead.
func (RouteLookupRequest_Reason) EnumDescriptor() ([]byte, []int) {
	return file_grpc_lookup_v1_rls_proto_rawDescGZIP(), []int{0, 0}
}

type RouteLookupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Target type allows the client to specify what kind of target format it
	// would like from RLS to allow it to find the regional server, e.g. "grpc".
	TargetType string `protobuf:"bytes,3,opt,name=target_type,json=targetType,proto3" json:"target_type,omitempty"`
	// Reason for making this request.
	Reason RouteLookupRequest_Reason `protobuf:"varint,5,opt,name=reason,proto3,enum=grpc.lookup.v1.RouteLookupRequest_Reason" json:"reason,omitempty"`
	// For REASON_STALE, the header_data from the stale response, if any.
	StaleHeaderData string `protobuf:"bytes,6,opt,name=stale_header_data,json=staleHeaderData,proto3" json:"stale_header_data,omitempty"`
	// Map of key values extracted via key builders for the gRPC or HTTP request.
	KeyMap map[string]string `protobuf:"bytes,4,rep,name=key_map,json=keyMap,proto3" json:"key_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Application-specific optional extensions.
	Extensions []*anypb.Any `protobuf:"bytes,7,rep,name=extensions,proto3" json:"extensions,omitempty"`
}

func (x *RouteLookupRequest) Reset() {
	*x = RouteLookupRequest{}
	mi := &file_grpc_lookup_v1_rls_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RouteLookupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteLookupRequest) ProtoMessage() {}

func (x *RouteLookupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_lookup_v1_rls_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteLookupRequest.ProtoReflect.Descriptor instead.
func (*RouteLookupRequest) Descriptor() ([]byte, []int) {
	return file_grpc_lookup_v1_rls_proto_rawDescGZIP(), []int{0}
}

func (x *RouteLookupRequest) GetTargetType() string {
	if x != nil {
		return x.TargetType
	}
	return ""
}

func (x *RouteLookupRequest) GetReason() RouteLookupRequest_Reason {
	if x != nil {
		return x.Reason
	}
	return RouteLookupRequest_REASON_UNKNOWN
}

func (x *RouteLookupRequest) GetStaleHeaderData() string {
	if x != nil {
		return x.StaleHeaderData
	}
	return ""
}

func (x *RouteLookupRequest) GetKeyMap() map[string]string {
	if x != nil {
		return x.KeyMap
	}
	return nil
}

func (x *RouteLookupRequest) GetExtensions() []*anypb.Any {
	if x != nil {
		return x.Extensions
	}
	return nil
}

type RouteLookupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Prioritized list (best one first) of addressable entities to use
	// for routing, using syntax requested by the request target_type.
	// The targets will be tried in order until a healthy one is found.
	Targets []string `protobuf:"bytes,3,rep,name=targets,proto3" json:"targets,omitempty"`
	// Optional header value to pass along to AFE in the X-Google-RLS-Data header.
	// Cached with "target" and sent with all requests that match the request key.
	// Allows the RLS to pass its work product to the eventual target.
	HeaderData string `protobuf:"bytes,2,opt,name=header_data,json=headerData,proto3" json:"header_data,omitempty"`
	// Application-specific optional extensions.
	Extensions []*anypb.Any `protobuf:"bytes,4,rep,name=extensions,proto3" json:"extensions,omitempty"`
}

func (x *RouteLookupResponse) Reset() {
	*x = RouteLookupResponse{}
	mi := &file_grpc_lookup_v1_rls_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RouteLookupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteLookupResponse) ProtoMessage() {}

func (x *RouteLookupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_lookup_v1_rls_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteLookupResponse.ProtoReflect.Descriptor instead.
func (*RouteLookupResponse) Descriptor() ([]byte, []int) {
	return file_grpc_lookup_v1_rls_proto_rawDescGZIP(), []int{1}
}

func (x *RouteLookupResponse) GetTargets() []string {
	if x != nil {
		return x.Targets
	}
	return nil
}

func (x *RouteLookupResponse) GetHeaderData() string {
	if x != nil {
		return x.HeaderData
	}
	return ""
}

func (x *RouteLookupResponse) GetExtensions() []*anypb.Any {
	if x != nil {
		return x.Extensions
	}
	return nil
}

var File_grpc_lookup_v1_rls_proto protoreflect.FileDescriptor

var file_grpc_lookup_v1_rls_proto_rawDesc = []byte{
	0x0a, 0x18, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x6c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x2f, 0x76, 0x31,
	0x2f, 0x72, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x6c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb9, 0x03, 0x0a, 0x12, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x4c,
	0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x41, 0x0a,
	0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x29, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x6c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x12, 0x2a, 0x0a, 0x11, 0x73, 0x74, 0x61, 0x6c, 0x65, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x74, 0x61,
	0x6c, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x47, 0x0a, 0x07,
	0x6b, 0x65, 0x79, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x6c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x4b, 0x65, 0x79, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6b,
	0x65, 0x79, 0x4d, 0x61, 0x70, 0x12, 0x34, 0x0a, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52,
	0x0a, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4b,
	0x65, 0x79, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x3f, 0x0a, 0x06, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x12, 0x12, 0x0a, 0x0e, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x4d,
	0x49, 0x53, 0x53, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f,
	0x53, 0x54, 0x41, 0x4c, 0x45, 0x10, 0x02, 0x4a, 0x04, 0x08, 0x01, 0x10, 0x02, 0x4a, 0x04, 0x08,
	0x02, 0x10, 0x03, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x22, 0x94, 0x01, 0x0a, 0x13, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x4c, 0x6f, 0x6f, 0x6b, 0x75,
	0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x34, 0x0a, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x0a,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x4a, 0x04, 0x08, 0x01, 0x10, 0x02,
	0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x32, 0x6e, 0x0a, 0x12, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x58,
	0x0a, 0x0b, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x12, 0x22, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x6c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x23, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x6c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x4d, 0x0a, 0x11, 0x69, 0x6f, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x6c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x42, 0x08, 0x52,
	0x6c, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x6c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x6c, 0x6f,
	0x6f, 0x6b, 0x75, 0x70, 0x5f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_lookup_v1_rls_proto_rawDescOnce sync.Once
	file_grpc_lookup_v1_rls_proto_rawDescData = file_grpc_lookup_v1_rls_proto_rawDesc
)

func file_grpc_lookup_v1_rls_proto_rawDescGZIP() []byte {
	file_grpc_lookup_v1_rls_proto_rawDescOnce.Do(func() {
		file_grpc_lookup_v1_rls_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_lookup_v1_rls_proto_rawDescData)
	})
	return file_grpc_lookup_v1_rls_proto_rawDescData
}

var file_grpc_lookup_v1_rls_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_grpc_lookup_v1_rls_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_grpc_lookup_v1_rls_proto_goTypes = []any{
	(RouteLookupRequest_Reason)(0), // 0: grpc.lookup.v1.RouteLookupRequest.Reason
	(*RouteLookupRequest)(nil),     // 1: grpc.lookup.v1.RouteLookupRequest
	(*RouteLookupResponse)(nil),    // 2: grpc.lookup.v1.RouteLookupResponse
	nil,                            // 3: grpc.lookup.v1.RouteLookupRequest.KeyMapEntry
	(*anypb.Any)(nil),              // 4: google.protobuf.Any
}
var file_grpc_lookup_v1_rls_proto_depIdxs = []int32{
	0, // 0: grpc.lookup.v1.RouteLookupRequest.reason:type_name -> grpc.lookup.v1.RouteLookupRequest.Reason
	3, // 1: grpc.lookup.v1.RouteLookupRequest.key_map:type_name -> grpc.lookup.v1.RouteLookupRequest.KeyMapEntry
	4, // 2: grpc.lookup.v1.RouteLookupRequest.extensions:type_name -> google.protobuf.Any
	4, // 3: grpc.lookup.v1.RouteLookupResponse.extensions:type_name -> google.protobuf.Any
	1, // 4: grpc.lookup.v1.RouteLookupService.RouteLookup:input_type -> grpc.lookup.v1.RouteLookupRequest
	2, // 5: grpc.lookup.v1.RouteLookupService.RouteLookup:output_type -> grpc.lookup.v1.RouteLookupResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_grpc_lookup_v1_rls_proto_init() }
func file_grpc_lookup_v1_rls_proto_init() {
	if File_grpc_lookup_v1_rls_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_grpc_lookup_v1_rls_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_lookup_v1_rls_proto_goTypes,
		DependencyIndexes: file_grpc_lookup_v1_rls_proto_depIdxs,
		EnumInfos:         file_grpc_lookup_v1_rls_proto_enumTypes,
		MessageInfos:      file_grpc_lookup_v1_rls_proto_msgTypes,
	}.Build()
	File_grpc_lookup_v1_rls_proto = out.File
	file_grpc_lookup_v1_rls_proto_rawDesc = nil
	file_grpc_lookup_v1_rls_proto_goTypes = nil
	file_grpc_lookup_v1_rls_proto_depIdxs = nil
}
