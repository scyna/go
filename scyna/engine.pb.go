// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: engine.proto

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

type Configuration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NatsUrl      string `protobuf:"bytes,1,opt,name=NatsUrl,proto3" json:"NatsUrl,omitempty"`
	NatsUsername string `protobuf:"bytes,2,opt,name=NatsUsername,proto3" json:"NatsUsername,omitempty"`
	NatsPassword string `protobuf:"bytes,3,opt,name=NatsPassword,proto3" json:"NatsPassword,omitempty"`
	DBHost       string `protobuf:"bytes,4,opt,name=DBHost,proto3" json:"DBHost,omitempty"`
	DBUsername   string `protobuf:"bytes,5,opt,name=DBUsername,proto3" json:"DBUsername,omitempty"`
	DBPassword   string `protobuf:"bytes,6,opt,name=DBPassword,proto3" json:"DBPassword,omitempty"`
	DBLocation   string `protobuf:"bytes,7,opt,name=DBLocation,proto3" json:"DBLocation,omitempty"`
}

func (x *Configuration) Reset() {
	*x = Configuration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_engine_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Configuration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Configuration) ProtoMessage() {}

func (x *Configuration) ProtoReflect() protoreflect.Message {
	mi := &file_engine_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Configuration.ProtoReflect.Descriptor instead.
func (*Configuration) Descriptor() ([]byte, []int) {
	return file_engine_proto_rawDescGZIP(), []int{0}
}

func (x *Configuration) GetNatsUrl() string {
	if x != nil {
		return x.NatsUrl
	}
	return ""
}

func (x *Configuration) GetNatsUsername() string {
	if x != nil {
		return x.NatsUsername
	}
	return ""
}

func (x *Configuration) GetNatsPassword() string {
	if x != nil {
		return x.NatsPassword
	}
	return ""
}

func (x *Configuration) GetDBHost() string {
	if x != nil {
		return x.DBHost
	}
	return ""
}

func (x *Configuration) GetDBUsername() string {
	if x != nil {
		return x.DBUsername
	}
	return ""
}

func (x *Configuration) GetDBPassword() string {
	if x != nil {
		return x.DBPassword
	}
	return ""
}

func (x *Configuration) GetDBLocation() string {
	if x != nil {
		return x.DBLocation
	}
	return ""
}

//session
type CreateSessionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Context string `protobuf:"bytes,1,opt,name=Context,proto3" json:"Context,omitempty"`
	Secret  string `protobuf:"bytes,2,opt,name=Secret,proto3" json:"Secret,omitempty"`
}

func (x *CreateSessionRequest) Reset() {
	*x = CreateSessionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_engine_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSessionRequest) ProtoMessage() {}

func (x *CreateSessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_engine_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSessionRequest.ProtoReflect.Descriptor instead.
func (*CreateSessionRequest) Descriptor() ([]byte, []int) {
	return file_engine_proto_rawDescGZIP(), []int{1}
}

func (x *CreateSessionRequest) GetContext() string {
	if x != nil {
		return x.Context
	}
	return ""
}

func (x *CreateSessionRequest) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

type CreateSessionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionID uint64         `protobuf:"varint,1,opt,name=SessionID,proto3" json:"SessionID,omitempty"`
	Config    *Configuration `protobuf:"bytes,5,opt,name=Config,proto3" json:"Config,omitempty"`
}

func (x *CreateSessionResponse) Reset() {
	*x = CreateSessionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_engine_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSessionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSessionResponse) ProtoMessage() {}

func (x *CreateSessionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_engine_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSessionResponse.ProtoReflect.Descriptor instead.
func (*CreateSessionResponse) Descriptor() ([]byte, []int) {
	return file_engine_proto_rawDescGZIP(), []int{2}
}

func (x *CreateSessionResponse) GetSessionID() uint64 {
	if x != nil {
		return x.SessionID
	}
	return 0
}

func (x *CreateSessionResponse) GetConfig() *Configuration {
	if x != nil {
		return x.Config
	}
	return nil
}

type EndSessionSignal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID     uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Code   string `protobuf:"bytes,2,opt,name=Code,proto3" json:"Code,omitempty"`
	Module string `protobuf:"bytes,3,opt,name=Module,proto3" json:"Module,omitempty"`
}

func (x *EndSessionSignal) Reset() {
	*x = EndSessionSignal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_engine_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndSessionSignal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndSessionSignal) ProtoMessage() {}

func (x *EndSessionSignal) ProtoReflect() protoreflect.Message {
	mi := &file_engine_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndSessionSignal.ProtoReflect.Descriptor instead.
func (*EndSessionSignal) Descriptor() ([]byte, []int) {
	return file_engine_proto_rawDescGZIP(), []int{3}
}

func (x *EndSessionSignal) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *EndSessionSignal) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *EndSessionSignal) GetModule() string {
	if x != nil {
		return x.Module
	}
	return ""
}

type UpdateSessionSignal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID     uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Module string `protobuf:"bytes,2,opt,name=Module,proto3" json:"Module,omitempty"`
}

func (x *UpdateSessionSignal) Reset() {
	*x = UpdateSessionSignal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_engine_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSessionSignal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSessionSignal) ProtoMessage() {}

func (x *UpdateSessionSignal) ProtoReflect() protoreflect.Message {
	mi := &file_engine_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSessionSignal.ProtoReflect.Descriptor instead.
func (*UpdateSessionSignal) Descriptor() ([]byte, []int) {
	return file_engine_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateSessionSignal) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *UpdateSessionSignal) GetModule() string {
	if x != nil {
		return x.Module
	}
	return ""
}

//GENERATOR
type GetIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetIDRequest) Reset() {
	*x = GetIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_engine_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIDRequest) ProtoMessage() {}

func (x *GetIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_engine_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIDRequest.ProtoReflect.Descriptor instead.
func (*GetIDRequest) Descriptor() ([]byte, []int) {
	return file_engine_proto_rawDescGZIP(), []int{5}
}

type GetIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Prefix uint32 `protobuf:"varint,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
	Start  uint64 `protobuf:"varint,2,opt,name=start,proto3" json:"start,omitempty"`
	End    uint64 `protobuf:"varint,3,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *GetIDResponse) Reset() {
	*x = GetIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_engine_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIDResponse) ProtoMessage() {}

func (x *GetIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_engine_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIDResponse.ProtoReflect.Descriptor instead.
func (*GetIDResponse) Descriptor() ([]byte, []int) {
	return file_engine_proto_rawDescGZIP(), []int{6}
}

func (x *GetIDResponse) GetPrefix() uint32 {
	if x != nil {
		return x.Prefix
	}
	return 0
}

func (x *GetIDResponse) GetStart() uint64 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *GetIDResponse) GetEnd() uint64 {
	if x != nil {
		return x.End
	}
	return 0
}

type GetSNRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *GetSNRequest) Reset() {
	*x = GetSNRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_engine_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSNRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSNRequest) ProtoMessage() {}

func (x *GetSNRequest) ProtoReflect() protoreflect.Message {
	mi := &file_engine_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSNRequest.ProtoReflect.Descriptor instead.
func (*GetSNRequest) Descriptor() ([]byte, []int) {
	return file_engine_proto_rawDescGZIP(), []int{7}
}

func (x *GetSNRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type GetSNResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Prefix uint32 `protobuf:"varint,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
	Start  uint64 `protobuf:"varint,2,opt,name=start,proto3" json:"start,omitempty"`
	End    uint64 `protobuf:"varint,3,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *GetSNResponse) Reset() {
	*x = GetSNResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_engine_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSNResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSNResponse) ProtoMessage() {}

func (x *GetSNResponse) ProtoReflect() protoreflect.Message {
	mi := &file_engine_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSNResponse.ProtoReflect.Descriptor instead.
func (*GetSNResponse) Descriptor() ([]byte, []int) {
	return file_engine_proto_rawDescGZIP(), []int{8}
}

func (x *GetSNResponse) GetPrefix() uint32 {
	if x != nil {
		return x.Prefix
	}
	return 0
}

func (x *GetSNResponse) GetStart() uint64 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *GetSNResponse) GetEnd() uint64 {
	if x != nil {
		return x.End
	}
	return 0
}

//LOG
type LogCreatedSignal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time    uint64 `protobuf:"varint,1,opt,name=Time,proto3" json:"Time,omitempty"`
	Level   uint32 `protobuf:"varint,2,opt,name=Level,proto3" json:"Level,omitempty"`
	Text    string `protobuf:"bytes,3,opt,name=Text,proto3" json:"Text,omitempty"`
	ID      uint64 `protobuf:"varint,4,opt,name=ID,proto3" json:"ID,omitempty"`
	SEQ     uint64 `protobuf:"varint,5,opt,name=SEQ,proto3" json:"SEQ,omitempty"`
	Session bool   `protobuf:"varint,6,opt,name=Session,proto3" json:"Session,omitempty"`
}

func (x *LogCreatedSignal) Reset() {
	*x = LogCreatedSignal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_engine_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogCreatedSignal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogCreatedSignal) ProtoMessage() {}

func (x *LogCreatedSignal) ProtoReflect() protoreflect.Message {
	mi := &file_engine_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogCreatedSignal.ProtoReflect.Descriptor instead.
func (*LogCreatedSignal) Descriptor() ([]byte, []int) {
	return file_engine_proto_rawDescGZIP(), []int{9}
}

func (x *LogCreatedSignal) GetTime() uint64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *LogCreatedSignal) GetLevel() uint32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *LogCreatedSignal) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *LogCreatedSignal) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *LogCreatedSignal) GetSEQ() uint64 {
	if x != nil {
		return x.SEQ
	}
	return 0
}

func (x *LogCreatedSignal) GetSession() bool {
	if x != nil {
		return x.Session
	}
	return false
}

type TraceCreatedSignal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	ParentID  uint64 `protobuf:"varint,2,opt,name=ParentID,proto3" json:"ParentID,omitempty"`
	Type      uint32 `protobuf:"varint,3,opt,name=Type,proto3" json:"Type,omitempty"`
	Time      uint64 `protobuf:"varint,4,opt,name=Time,proto3" json:"Time,omitempty"`
	Duration  uint64 `protobuf:"varint,5,opt,name=Duration,proto3" json:"Duration,omitempty"`
	Path      string `protobuf:"bytes,6,opt,name=Path,proto3" json:"Path,omitempty"`
	Source    string `protobuf:"bytes,7,opt,name=Source,proto3" json:"Source,omitempty"`
	SessionID uint64 `protobuf:"varint,8,opt,name=SessionID,proto3" json:"SessionID,omitempty"`
	Status    int32  `protobuf:"varint,9,opt,name=Status,proto3" json:"Status,omitempty"`
}

func (x *TraceCreatedSignal) Reset() {
	*x = TraceCreatedSignal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_engine_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TraceCreatedSignal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TraceCreatedSignal) ProtoMessage() {}

func (x *TraceCreatedSignal) ProtoReflect() protoreflect.Message {
	mi := &file_engine_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TraceCreatedSignal.ProtoReflect.Descriptor instead.
func (*TraceCreatedSignal) Descriptor() ([]byte, []int) {
	return file_engine_proto_rawDescGZIP(), []int{10}
}

func (x *TraceCreatedSignal) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *TraceCreatedSignal) GetParentID() uint64 {
	if x != nil {
		return x.ParentID
	}
	return 0
}

func (x *TraceCreatedSignal) GetType() uint32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *TraceCreatedSignal) GetTime() uint64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *TraceCreatedSignal) GetDuration() uint64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *TraceCreatedSignal) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *TraceCreatedSignal) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *TraceCreatedSignal) GetSessionID() uint64 {
	if x != nil {
		return x.SessionID
	}
	return 0
}

func (x *TraceCreatedSignal) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type TagCreatedSignal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TraceID uint64 `protobuf:"varint,1,opt,name=TraceID,proto3" json:"TraceID,omitempty"`
	Key     string `protobuf:"bytes,2,opt,name=Key,proto3" json:"Key,omitempty"`
	Value   string `protobuf:"bytes,3,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *TagCreatedSignal) Reset() {
	*x = TagCreatedSignal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_engine_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TagCreatedSignal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagCreatedSignal) ProtoMessage() {}

func (x *TagCreatedSignal) ProtoReflect() protoreflect.Message {
	mi := &file_engine_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagCreatedSignal.ProtoReflect.Descriptor instead.
func (*TagCreatedSignal) Descriptor() ([]byte, []int) {
	return file_engine_proto_rawDescGZIP(), []int{11}
}

func (x *TagCreatedSignal) GetTraceID() uint64 {
	if x != nil {
		return x.TraceID
	}
	return 0
}

func (x *TagCreatedSignal) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *TagCreatedSignal) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type ServiceDoneSignal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TraceID  uint64 `protobuf:"varint,1,opt,name=TraceID,proto3" json:"TraceID,omitempty"`
	Response string `protobuf:"bytes,2,opt,name=Response,proto3" json:"Response,omitempty"`
}

func (x *ServiceDoneSignal) Reset() {
	*x = ServiceDoneSignal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_engine_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceDoneSignal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceDoneSignal) ProtoMessage() {}

func (x *ServiceDoneSignal) ProtoReflect() protoreflect.Message {
	mi := &file_engine_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceDoneSignal.ProtoReflect.Descriptor instead.
func (*ServiceDoneSignal) Descriptor() ([]byte, []int) {
	return file_engine_proto_rawDescGZIP(), []int{12}
}

func (x *ServiceDoneSignal) GetTraceID() uint64 {
	if x != nil {
		return x.TraceID
	}
	return 0
}

func (x *ServiceDoneSignal) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

var File_engine_proto protoreflect.FileDescriptor

var file_engine_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x73, 0x63, 0x79, 0x6e, 0x61, 0x22, 0xe9, 0x01, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x4e, 0x61, 0x74, 0x73, 0x55,
	0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4e, 0x61, 0x74, 0x73, 0x55, 0x72,
	0x6c, 0x12, 0x22, 0x0a, 0x0c, 0x4e, 0x61, 0x74, 0x73, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x4e, 0x61, 0x74, 0x73, 0x55, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x4e, 0x61, 0x74, 0x73, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x4e, 0x61, 0x74,
	0x73, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x44, 0x42, 0x48,
	0x6f, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x44, 0x42, 0x48, 0x6f, 0x73,
	0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x44, 0x42, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x44, 0x42, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x44, 0x42, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x44, 0x42, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x44, 0x42, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x44, 0x42, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x48, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x22, 0x63, 0x0a, 0x15, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x49, 0x44, 0x12, 0x2c, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x63, 0x79, 0x6e, 0x61, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x22, 0x4e, 0x0a, 0x10, 0x45, 0x6e, 0x64, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x69,
	0x67, 0x6e, 0x61, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x4d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x22, 0x3d, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x4d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x22,
	0x0e, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x4f, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x65, 0x6e, 0x64,
	0x22, 0x20, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x53, 0x4e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x22, 0x4f, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x53, 0x4e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03,
	0x65, 0x6e, 0x64, 0x22, 0x8c, 0x01, 0x0a, 0x10, 0x4c, 0x6f, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x54, 0x65, 0x78, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x12, 0x10, 0x0a, 0x03, 0x53, 0x45, 0x51, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x03, 0x53, 0x45, 0x51, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x22, 0xe6, 0x01, 0x0a, 0x12, 0x54, 0x72, 0x61, 0x63, 0x65, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x50, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x08, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x61, 0x74,
	0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x50, 0x61, 0x74, 0x68, 0x12, 0x16, 0x0a,
	0x06, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x49, 0x44, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x54, 0x0a, 0x10, 0x54,
	0x61, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x12,
	0x18, 0x0a, 0x07, 0x54, 0x72, 0x61, 0x63, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x07, 0x54, 0x72, 0x61, 0x63, 0x65, 0x49, 0x44, 0x12, 0x10, 0x0a, 0x03, 0x4b, 0x65, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x49, 0x0a, 0x11, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x6f, 0x6e, 0x65,
	0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x54, 0x72, 0x61, 0x63, 0x65, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x54, 0x72, 0x61, 0x63, 0x65, 0x49, 0x44,
	0x12, 0x1a, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2c, 0x0a, 0x0e,
	0x69, 0x6f, 0x2e, 0x73, 0x63, 0x79, 0x6e, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x48, 0x02,
	0x50, 0x01, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x73, 0x63, 0x79, 0x6e, 0x61, 0xaa, 0x02, 0x0b, 0x73,
	0x63, 0x79, 0x6e, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_engine_proto_rawDescOnce sync.Once
	file_engine_proto_rawDescData = file_engine_proto_rawDesc
)

func file_engine_proto_rawDescGZIP() []byte {
	file_engine_proto_rawDescOnce.Do(func() {
		file_engine_proto_rawDescData = protoimpl.X.CompressGZIP(file_engine_proto_rawDescData)
	})
	return file_engine_proto_rawDescData
}

var file_engine_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_engine_proto_goTypes = []interface{}{
	(*Configuration)(nil),         // 0: scyna.Configuration
	(*CreateSessionRequest)(nil),  // 1: scyna.CreateSessionRequest
	(*CreateSessionResponse)(nil), // 2: scyna.CreateSessionResponse
	(*EndSessionSignal)(nil),      // 3: scyna.EndSessionSignal
	(*UpdateSessionSignal)(nil),   // 4: scyna.UpdateSessionSignal
	(*GetIDRequest)(nil),          // 5: scyna.GetIDRequest
	(*GetIDResponse)(nil),         // 6: scyna.GetIDResponse
	(*GetSNRequest)(nil),          // 7: scyna.GetSNRequest
	(*GetSNResponse)(nil),         // 8: scyna.GetSNResponse
	(*LogCreatedSignal)(nil),      // 9: scyna.LogCreatedSignal
	(*TraceCreatedSignal)(nil),    // 10: scyna.TraceCreatedSignal
	(*TagCreatedSignal)(nil),      // 11: scyna.TagCreatedSignal
	(*ServiceDoneSignal)(nil),     // 12: scyna.ServiceDoneSignal
}
var file_engine_proto_depIdxs = []int32{
	0, // 0: scyna.CreateSessionResponse.Config:type_name -> scyna.Configuration
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_engine_proto_init() }
func file_engine_proto_init() {
	if File_engine_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_engine_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Configuration); i {
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
		file_engine_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSessionRequest); i {
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
		file_engine_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSessionResponse); i {
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
		file_engine_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndSessionSignal); i {
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
		file_engine_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSessionSignal); i {
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
		file_engine_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetIDRequest); i {
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
		file_engine_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetIDResponse); i {
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
		file_engine_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSNRequest); i {
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
		file_engine_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSNResponse); i {
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
		file_engine_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogCreatedSignal); i {
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
		file_engine_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TraceCreatedSignal); i {
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
		file_engine_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TagCreatedSignal); i {
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
		file_engine_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceDoneSignal); i {
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
			RawDescriptor: file_engine_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_engine_proto_goTypes,
		DependencyIndexes: file_engine_proto_depIdxs,
		MessageInfos:      file_engine_proto_msgTypes,
	}.Build()
	File_engine_proto = out.File
	file_engine_proto_rawDesc = nil
	file_engine_proto_goTypes = nil
	file_engine_proto_depIdxs = nil
}
