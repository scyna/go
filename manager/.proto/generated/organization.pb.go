// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: organization.proto

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

type Organization struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code     string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *Organization) Reset() {
	*x = Organization{}
	if protoimpl.UnsafeEnabled {
		mi := &file_organization_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Organization) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Organization) ProtoMessage() {}

func (x *Organization) ProtoReflect() protoreflect.Message {
	mi := &file_organization_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Organization.ProtoReflect.Descriptor instead.
func (*Organization) Descriptor() ([]byte, []int) {
	return file_organization_proto_rawDescGZIP(), []int{0}
}

func (x *Organization) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Organization) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Organization) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type DestroyOrganizationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code     string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *DestroyOrganizationRequest) Reset() {
	*x = DestroyOrganizationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_organization_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DestroyOrganizationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DestroyOrganizationRequest) ProtoMessage() {}

func (x *DestroyOrganizationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_organization_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DestroyOrganizationRequest.ProtoReflect.Descriptor instead.
func (*DestroyOrganizationRequest) Descriptor() ([]byte, []int) {
	return file_organization_proto_rawDescGZIP(), []int{1}
}

func (x *DestroyOrganizationRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *DestroyOrganizationRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

var File_organization_proto protoreflect.FileDescriptor

var file_organization_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x52, 0x0a, 0x0c, 0x4f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22,
	0x4c, 0x0a, 0x1a, 0x44, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x42, 0x0c, 0x48,
	0x02, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_organization_proto_rawDescOnce sync.Once
	file_organization_proto_rawDescData = file_organization_proto_rawDesc
)

func file_organization_proto_rawDescGZIP() []byte {
	file_organization_proto_rawDescOnce.Do(func() {
		file_organization_proto_rawDescData = protoimpl.X.CompressGZIP(file_organization_proto_rawDescData)
	})
	return file_organization_proto_rawDescData
}

var file_organization_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_organization_proto_goTypes = []interface{}{
	(*Organization)(nil),               // 0: proto.Organization
	(*DestroyOrganizationRequest)(nil), // 1: proto.DestroyOrganizationRequest
}
var file_organization_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_organization_proto_init() }
func file_organization_proto_init() {
	if File_organization_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_organization_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Organization); i {
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
		file_organization_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DestroyOrganizationRequest); i {
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
			RawDescriptor: file_organization_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_organization_proto_goTypes,
		DependencyIndexes: file_organization_proto_depIdxs,
		MessageInfos:      file_organization_proto_msgTypes,
	}.Build()
	File_organization_proto = out.File
	file_organization_proto_rawDesc = nil
	file_organization_proto_goTypes = nil
	file_organization_proto_depIdxs = nil
}
