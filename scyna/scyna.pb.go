// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: scyna.proto

package scyna

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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TraceID uint64 `protobuf:"varint,1,opt,name=TraceID,proto3" json:"TraceID,omitempty"`
	Body    []byte `protobuf:"bytes,2,opt,name=Body,proto3" json:"Body,omitempty"`
	Data    string `protobuf:"bytes,4,opt,name=Data,proto3" json:"Data,omitempty"`
	JSON    bool   `protobuf:"varint,3,opt,name=JSON,proto3" json:"JSON,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scyna_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_scyna_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_scyna_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetTraceID() uint64 {
	if x != nil {
		return x.TraceID
	}
	return 0
}

func (x *Request) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *Request) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *Request) GetJSON() bool {
	if x != nil {
		return x.JSON
	}
	return false
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code      int32  `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Body      []byte `protobuf:"bytes,2,opt,name=Body,proto3" json:"Body,omitempty"`
	SessionID uint64 `protobuf:"varint,3,opt,name=SessionID,proto3" json:"SessionID,omitempty"`
	Token     string `protobuf:"bytes,4,opt,name=Token,proto3" json:"Token,omitempty"`
	Expired   uint64 `protobuf:"varint,5,opt,name=Expired,proto3" json:"Expired,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scyna_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_scyna_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_scyna_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Response) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *Response) GetSessionID() uint64 {
	if x != nil {
		return x.SessionID
	}
	return 0
}

func (x *Response) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *Response) GetExpired() uint64 {
	if x != nil {
		return x.Expired
	}
	return 0
}

type EventOrSignal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ParentID uint64   `protobuf:"varint,1,opt,name=ParentID,proto3" json:"ParentID,omitempty"`
	Body     []byte   `protobuf:"bytes,2,opt,name=Body,proto3" json:"Body,omitempty"`
	Data     string   `protobuf:"bytes,3,opt,name=Data,proto3" json:"Data,omitempty"`
	Entities []uint64 `protobuf:"varint,4,rep,packed,name=Entities,proto3" json:"Entities,omitempty"`
}

func (x *EventOrSignal) Reset() {
	*x = EventOrSignal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scyna_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventOrSignal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventOrSignal) ProtoMessage() {}

func (x *EventOrSignal) ProtoReflect() protoreflect.Message {
	mi := &file_scyna_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventOrSignal.ProtoReflect.Descriptor instead.
func (*EventOrSignal) Descriptor() ([]byte, []int) {
	return file_scyna_proto_rawDescGZIP(), []int{2}
}

func (x *EventOrSignal) GetParentID() uint64 {
	if x != nil {
		return x.ParentID
	}
	return 0
}

func (x *EventOrSignal) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *EventOrSignal) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *EventOrSignal) GetEntities() []uint64 {
	if x != nil {
		return x.Entities
	}
	return nil
}

type EmptyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyRequest) Reset() {
	*x = EmptyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scyna_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyRequest) ProtoMessage() {}

func (x *EmptyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_scyna_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyRequest.ProtoReflect.Descriptor instead.
func (*EmptyRequest) Descriptor() ([]byte, []int) {
	return file_scyna_proto_rawDescGZIP(), []int{3}
}

var File_scyna_proto protoreflect.FileDescriptor

var file_scyna_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x73, 0x63, 0x79, 0x6e, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73,
	0x63, 0x79, 0x6e, 0x61, 0x22, 0x5f, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x54, 0x72, 0x61, 0x63, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x07, 0x54, 0x72, 0x61, 0x63, 0x65, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x42, 0x6f, 0x64,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x12, 0x0a,
	0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x12, 0x0a, 0x04, 0x4a, 0x53, 0x4f, 0x4e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x04, 0x4a, 0x53, 0x4f, 0x4e, 0x22, 0x80, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x42, 0x6f, 0x64, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x53,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x18,
	0x0a, 0x07, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x07, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x22, 0x6f, 0x0a, 0x0d, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x4f, 0x72, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x50, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x42, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x04, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74,
	0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a,
	0x08, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x04, 0x52,
	0x08, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x22, 0x0e, 0x0a, 0x0c, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x2c, 0x0a, 0x0e, 0x69, 0x6f, 0x2e,
	0x73, 0x63, 0x79, 0x6e, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x48, 0x02, 0x50, 0x01, 0x5a,
	0x08, 0x2e, 0x2f, 0x3b, 0x73, 0x63, 0x79, 0x6e, 0x61, 0xaa, 0x02, 0x0b, 0x73, 0x63, 0x79, 0x6e,
	0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_scyna_proto_rawDescOnce sync.Once
	file_scyna_proto_rawDescData = file_scyna_proto_rawDesc
)

func file_scyna_proto_rawDescGZIP() []byte {
	file_scyna_proto_rawDescOnce.Do(func() {
		file_scyna_proto_rawDescData = protoimpl.X.CompressGZIP(file_scyna_proto_rawDescData)
	})
	return file_scyna_proto_rawDescData
}

var file_scyna_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_scyna_proto_goTypes = []interface{}{
	(*Request)(nil),       // 0: scyna.Request
	(*Response)(nil),      // 1: scyna.Response
	(*EventOrSignal)(nil), // 2: scyna.EventOrSignal
	(*EmptyRequest)(nil),  // 3: scyna.EmptyRequest
}
var file_scyna_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_scyna_proto_init() }
func file_scyna_proto_init() {
	if File_scyna_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_scyna_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_scyna_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_scyna_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventOrSignal); i {
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
		file_scyna_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyRequest); i {
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
			RawDescriptor: file_scyna_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_scyna_proto_goTypes,
		DependencyIndexes: file_scyna_proto_depIdxs,
		MessageInfos:      file_scyna_proto_msgTypes,
	}.Build()
	File_scyna_proto = out.File
	file_scyna_proto_rawDesc = nil
	file_scyna_proto_goTypes = nil
	file_scyna_proto_depIdxs = nil
}
