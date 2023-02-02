// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: nats.proto

package proto

import (
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

type AddStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *AddStreamRequest) Reset() {
	*x = AddStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nats_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddStreamRequest) ProtoMessage() {}

func (x *AddStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nats_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddStreamRequest.ProtoReflect.Descriptor instead.
func (*AddStreamRequest) Descriptor() ([]byte, []int) {
	return file_nats_proto_rawDescGZIP(), []int{0}
}

func (x *AddStreamRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DeleteStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeleteStreamRequest) Reset() {
	*x = DeleteStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nats_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteStreamRequest) ProtoMessage() {}

func (x *DeleteStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nats_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteStreamRequest.ProtoReflect.Descriptor instead.
func (*DeleteStreamRequest) Descriptor() ([]byte, []int) {
	return file_nats_proto_rawDescGZIP(), []int{1}
}

func (x *DeleteStreamRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetStreamRequest) Reset() {
	*x = GetStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nats_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStreamRequest) ProtoMessage() {}

func (x *GetStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nats_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStreamRequest.ProtoReflect.Descriptor instead.
func (*GetStreamRequest) Descriptor() ([]byte, []int) {
	return file_nats_proto_rawDescGZIP(), []int{2}
}

func (x *GetStreamRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ListStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListStreamRequest) Reset() {
	*x = ListStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nats_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListStreamRequest) ProtoMessage() {}

func (x *ListStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nats_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListStreamRequest.ProtoReflect.Descriptor instead.
func (*ListStreamRequest) Descriptor() ([]byte, []int) {
	return file_nats_proto_rawDescGZIP(), []int{3}
}

type StreamInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Subject []string `protobuf:"bytes,2,rep,name=subject,proto3" json:"subject,omitempty"`
	Created string   `protobuf:"bytes,3,opt,name=created,proto3" json:"created,omitempty"`
}

func (x *StreamInfo) Reset() {
	*x = StreamInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nats_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamInfo) ProtoMessage() {}

func (x *StreamInfo) ProtoReflect() protoreflect.Message {
	mi := &file_nats_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamInfo.ProtoReflect.Descriptor instead.
func (*StreamInfo) Descriptor() ([]byte, []int) {
	return file_nats_proto_rawDescGZIP(), []int{4}
}

func (x *StreamInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StreamInfo) GetSubject() []string {
	if x != nil {
		return x.Subject
	}
	return nil
}

func (x *StreamInfo) GetCreated() string {
	if x != nil {
		return x.Created
	}
	return ""
}

type ListStreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Streams []*StreamInfo `protobuf:"bytes,1,rep,name=streams,proto3" json:"streams,omitempty"`
}

func (x *ListStreamResponse) Reset() {
	*x = ListStreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nats_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListStreamResponse) ProtoMessage() {}

func (x *ListStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nats_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListStreamResponse.ProtoReflect.Descriptor instead.
func (*ListStreamResponse) Descriptor() ([]byte, []int) {
	return file_nats_proto_rawDescGZIP(), []int{5}
}

func (x *ListStreamResponse) GetStreams() []*StreamInfo {
	if x != nil {
		return x.Streams
	}
	return nil
}

type UpdateStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *UpdateStreamRequest) Reset() {
	*x = UpdateStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nats_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStreamRequest) ProtoMessage() {}

func (x *UpdateStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nats_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStreamRequest.ProtoReflect.Descriptor instead.
func (*UpdateStreamRequest) Descriptor() ([]byte, []int) {
	return file_nats_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateStreamRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_nats_proto protoreflect.FileDescriptor

var file_nats_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6e, 0x61, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x26, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x29, 0x0a, 0x13, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x26, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x13,
	0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x54, 0x0a, 0x0a, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x22, 0x41, 0x0a, 0x12, 0x4c, 0x69, 0x73,
	0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2b, 0x0a, 0x07, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x07, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x22, 0x29, 0x0a, 0x13,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0c, 0x48, 0x02, 0x5a, 0x08, 0x2e, 0x2f, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nats_proto_rawDescOnce sync.Once
	file_nats_proto_rawDescData = file_nats_proto_rawDesc
)

func file_nats_proto_rawDescGZIP() []byte {
	file_nats_proto_rawDescOnce.Do(func() {
		file_nats_proto_rawDescData = protoimpl.X.CompressGZIP(file_nats_proto_rawDescData)
	})
	return file_nats_proto_rawDescData
}

var file_nats_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_nats_proto_goTypes = []interface{}{
	(*AddStreamRequest)(nil),    // 0: proto.AddStreamRequest
	(*DeleteStreamRequest)(nil), // 1: proto.DeleteStreamRequest
	(*GetStreamRequest)(nil),    // 2: proto.GetStreamRequest
	(*ListStreamRequest)(nil),   // 3: proto.ListStreamRequest
	(*StreamInfo)(nil),          // 4: proto.StreamInfo
	(*ListStreamResponse)(nil),  // 5: proto.ListStreamResponse
	(*UpdateStreamRequest)(nil), // 6: proto.UpdateStreamRequest
}
var file_nats_proto_depIdxs = []int32{
	4, // 0: proto.ListStreamResponse.streams:type_name -> proto.StreamInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_nats_proto_init() }
func file_nats_proto_init() {
	if File_nats_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_nats_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddStreamRequest); i {
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
		file_nats_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteStreamRequest); i {
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
		file_nats_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStreamRequest); i {
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
		file_nats_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListStreamRequest); i {
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
		file_nats_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamInfo); i {
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
		file_nats_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListStreamResponse); i {
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
		file_nats_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateStreamRequest); i {
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
			RawDescriptor: file_nats_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_nats_proto_goTypes,
		DependencyIndexes: file_nats_proto_depIdxs,
		MessageInfos:      file_nats_proto_msgTypes,
	}.Build()
	File_nats_proto = out.File
	file_nats_proto_rawDesc = nil
	file_nats_proto_goTypes = nil
	file_nats_proto_depIdxs = nil
}
