// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: sync.proto

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

type CreateSyncRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Module  string `protobuf:"bytes,1,opt,name=Module,proto3" json:"Module,omitempty"`   //vf_account
	Channel string `protobuf:"bytes,2,opt,name=Channel,proto3" json:"Channel,omitempty"` //sync_account_loyalty
}

func (x *CreateSyncRequest) Reset() {
	*x = CreateSyncRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sync_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSyncRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSyncRequest) ProtoMessage() {}

func (x *CreateSyncRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sync_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSyncRequest.ProtoReflect.Descriptor instead.
func (*CreateSyncRequest) Descriptor() ([]byte, []int) {
	return file_sync_proto_rawDescGZIP(), []int{0}
}

func (x *CreateSyncRequest) GetModule() string {
	if x != nil {
		return x.Module
	}
	return ""
}

func (x *CreateSyncRequest) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

type DropSyncRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Module  string `protobuf:"bytes,1,opt,name=Module,proto3" json:"Module,omitempty"`
	Channel string `protobuf:"bytes,2,opt,name=Channel,proto3" json:"Channel,omitempty"`
}

func (x *DropSyncRequest) Reset() {
	*x = DropSyncRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sync_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DropSyncRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DropSyncRequest) ProtoMessage() {}

func (x *DropSyncRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sync_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DropSyncRequest.ProtoReflect.Descriptor instead.
func (*DropSyncRequest) Descriptor() ([]byte, []int) {
	return file_sync_proto_rawDescGZIP(), []int{1}
}

func (x *DropSyncRequest) GetModule() string {
	if x != nil {
		return x.Module
	}
	return ""
}

func (x *DropSyncRequest) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

var File_sync_proto protoreflect.FileDescriptor

var file_sync_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x45, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x79, 0x6e,
	0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x4d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x22, 0x43, 0x0a, 0x0f, 0x44, 0x72,
	0x6f, 0x70, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x42,
	0x0c, 0x48, 0x02, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sync_proto_rawDescOnce sync.Once
	file_sync_proto_rawDescData = file_sync_proto_rawDesc
)

func file_sync_proto_rawDescGZIP() []byte {
	file_sync_proto_rawDescOnce.Do(func() {
		file_sync_proto_rawDescData = protoimpl.X.CompressGZIP(file_sync_proto_rawDescData)
	})
	return file_sync_proto_rawDescData
}

var file_sync_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_sync_proto_goTypes = []interface{}{
	(*CreateSyncRequest)(nil), // 0: proto.CreateSyncRequest
	(*DropSyncRequest)(nil),   // 1: proto.DropSyncRequest
}
var file_sync_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sync_proto_init() }
func file_sync_proto_init() {
	if File_sync_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sync_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSyncRequest); i {
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
		file_sync_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DropSyncRequest); i {
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
			RawDescriptor: file_sync_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sync_proto_goTypes,
		DependencyIndexes: file_sync_proto_depIdxs,
		MessageInfos:      file_sync_proto_msgTypes,
	}.Build()
	File_sync_proto = out.File
	file_sync_proto_rawDesc = nil
	file_sync_proto_goTypes = nil
	file_sync_proto_depIdxs = nil
}
