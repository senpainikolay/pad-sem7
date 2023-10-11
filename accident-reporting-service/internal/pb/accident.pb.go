// protoc --go_out=internal/pb --go_opt=paths=source_relative --go-grpc_out=internal/pb --go-grpc_opt=paths=source_relative internal/proto/accident.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.3
// source: internal/proto/accident.proto

package accident_reporting_service

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

type GetAccidentUserEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserLong  float64 `protobuf:"fixed64,1,opt,name=user_long,json=userLong,proto3" json:"user_long,omitempty"`
	UserLat   float64 `protobuf:"fixed64,2,opt,name=user_lat,json=userLat,proto3" json:"user_lat,omitempty"`
	ZoomIndex int64   `protobuf:"varint,3,opt,name=zoom_index,json=zoomIndex,proto3" json:"zoom_index,omitempty"`
	City      string  `protobuf:"bytes,4,opt,name=city,proto3" json:"city,omitempty"`
}

func (x *GetAccidentUserEntry) Reset() {
	*x = GetAccidentUserEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_accident_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccidentUserEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccidentUserEntry) ProtoMessage() {}

func (x *GetAccidentUserEntry) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_accident_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccidentUserEntry.ProtoReflect.Descriptor instead.
func (*GetAccidentUserEntry) Descriptor() ([]byte, []int) {
	return file_internal_proto_accident_proto_rawDescGZIP(), []int{0}
}

func (x *GetAccidentUserEntry) GetUserLong() float64 {
	if x != nil {
		return x.UserLong
	}
	return 0
}

func (x *GetAccidentUserEntry) GetUserLat() float64 {
	if x != nil {
		return x.UserLat
	}
	return 0
}

func (x *GetAccidentUserEntry) GetZoomIndex() int64 {
	if x != nil {
		return x.ZoomIndex
	}
	return 0
}

func (x *GetAccidentUserEntry) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

type AccidentInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccidentLong                     float64 `protobuf:"fixed64,1,opt,name=accident_long,json=accidentLong,proto3" json:"accident_long,omitempty"`
	AccidentLat                      float64 `protobuf:"fixed64,2,opt,name=accident_lat,json=accidentLat,proto3" json:"accident_lat,omitempty"`
	ConfirmationAccidentNotification bool    `protobuf:"varint,3,opt,name=confirmation_accident_notification,json=confirmationAccidentNotification,proto3" json:"confirmation_accident_notification,omitempty"`
	ConfirmationPoliceNotification   bool    `protobuf:"varint,4,opt,name=confirmation_police_notification,json=confirmationPoliceNotification,proto3" json:"confirmation_police_notification,omitempty"`
	ConfirmedBy                      int64   `protobuf:"varint,5,opt,name=confirmedBy,proto3" json:"confirmedBy,omitempty"`
}

func (x *AccidentInfo) Reset() {
	*x = AccidentInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_accident_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccidentInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccidentInfo) ProtoMessage() {}

func (x *AccidentInfo) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_accident_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccidentInfo.ProtoReflect.Descriptor instead.
func (*AccidentInfo) Descriptor() ([]byte, []int) {
	return file_internal_proto_accident_proto_rawDescGZIP(), []int{1}
}

func (x *AccidentInfo) GetAccidentLong() float64 {
	if x != nil {
		return x.AccidentLong
	}
	return 0
}

func (x *AccidentInfo) GetAccidentLat() float64 {
	if x != nil {
		return x.AccidentLat
	}
	return 0
}

func (x *AccidentInfo) GetConfirmationAccidentNotification() bool {
	if x != nil {
		return x.ConfirmationAccidentNotification
	}
	return false
}

func (x *AccidentInfo) GetConfirmationPoliceNotification() bool {
	if x != nil {
		return x.ConfirmationPoliceNotification
	}
	return false
}

func (x *AccidentInfo) GetConfirmedBy() int64 {
	if x != nil {
		return x.ConfirmedBy
	}
	return 0
}

type GetAccidentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccidentInfo []*AccidentInfo `protobuf:"bytes,1,rep,name=accident_info,json=accidentInfo,proto3" json:"accident_info,omitempty"`
}

func (x *GetAccidentResponse) Reset() {
	*x = GetAccidentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_accident_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccidentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccidentResponse) ProtoMessage() {}

func (x *GetAccidentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_accident_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccidentResponse.ProtoReflect.Descriptor instead.
func (*GetAccidentResponse) Descriptor() ([]byte, []int) {
	return file_internal_proto_accident_proto_rawDescGZIP(), []int{2}
}

func (x *GetAccidentResponse) GetAccidentInfo() []*AccidentInfo {
	if x != nil {
		return x.AccidentInfo
	}
	return nil
}

type PostAccidentEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccidentLong float64 `protobuf:"fixed64,1,opt,name=accident_long,json=accidentLong,proto3" json:"accident_long,omitempty"`
	AccidentLat  float64 `protobuf:"fixed64,2,opt,name=accident_lat,json=accidentLat,proto3" json:"accident_lat,omitempty"`
	City         string  `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`
	StreetName   string  `protobuf:"bytes,4,opt,name=street_name,json=streetName,proto3" json:"street_name,omitempty"`
	CarsInvolved int64   `protobuf:"varint,5,opt,name=cars_involved,json=carsInvolved,proto3" json:"cars_involved,omitempty"`
}

func (x *PostAccidentEntry) Reset() {
	*x = PostAccidentEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_accident_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostAccidentEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostAccidentEntry) ProtoMessage() {}

func (x *PostAccidentEntry) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_accident_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostAccidentEntry.ProtoReflect.Descriptor instead.
func (*PostAccidentEntry) Descriptor() ([]byte, []int) {
	return file_internal_proto_accident_proto_rawDescGZIP(), []int{3}
}

func (x *PostAccidentEntry) GetAccidentLong() float64 {
	if x != nil {
		return x.AccidentLong
	}
	return 0
}

func (x *PostAccidentEntry) GetAccidentLat() float64 {
	if x != nil {
		return x.AccidentLat
	}
	return 0
}

func (x *PostAccidentEntry) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *PostAccidentEntry) GetStreetName() string {
	if x != nil {
		return x.StreetName
	}
	return ""
}

func (x *PostAccidentEntry) GetCarsInvolved() int64 {
	if x != nil {
		return x.CarsInvolved
	}
	return 0
}

type ConfirmAccidentEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccidentLong         float64 `protobuf:"fixed64,1,opt,name=accident_long,json=accidentLong,proto3" json:"accident_long,omitempty"`
	AccidentLat          float64 `protobuf:"fixed64,2,opt,name=accident_lat,json=accidentLat,proto3" json:"accident_lat,omitempty"`
	PoliceConfirmation   bool    `protobuf:"varint,3,opt,name=police_confirmation,json=policeConfirmation,proto3" json:"police_confirmation,omitempty"`
	AccidentConfirmation bool    `protobuf:"varint,4,opt,name=accident_confirmation,json=accidentConfirmation,proto3" json:"accident_confirmation,omitempty"`
}

func (x *ConfirmAccidentEntry) Reset() {
	*x = ConfirmAccidentEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_accident_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfirmAccidentEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfirmAccidentEntry) ProtoMessage() {}

func (x *ConfirmAccidentEntry) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_accident_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfirmAccidentEntry.ProtoReflect.Descriptor instead.
func (*ConfirmAccidentEntry) Descriptor() ([]byte, []int) {
	return file_internal_proto_accident_proto_rawDescGZIP(), []int{4}
}

func (x *ConfirmAccidentEntry) GetAccidentLong() float64 {
	if x != nil {
		return x.AccidentLong
	}
	return 0
}

func (x *ConfirmAccidentEntry) GetAccidentLat() float64 {
	if x != nil {
		return x.AccidentLat
	}
	return 0
}

func (x *ConfirmAccidentEntry) GetPoliceConfirmation() bool {
	if x != nil {
		return x.PoliceConfirmation
	}
	return false
}

func (x *ConfirmAccidentEntry) GetAccidentConfirmation() bool {
	if x != nil {
		return x.AccidentConfirmation
	}
	return false
}

type GenericResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error bool   `protobuf:"varint,1,opt,name=error,proto3" json:"error,omitempty"`
	Msg   string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *GenericResponse) Reset() {
	*x = GenericResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_accident_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenericResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenericResponse) ProtoMessage() {}

func (x *GenericResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_accident_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenericResponse.ProtoReflect.Descriptor instead.
func (*GenericResponse) Descriptor() ([]byte, []int) {
	return file_internal_proto_accident_proto_rawDescGZIP(), []int{5}
}

func (x *GenericResponse) GetError() bool {
	if x != nil {
		return x.Error
	}
	return false
}

func (x *GenericResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type HealthResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ready    bool   `protobuf:"varint,1,opt,name=ready,proto3" json:"ready,omitempty"`
	Database string `protobuf:"bytes,2,opt,name=database,proto3" json:"database,omitempty"`
	Load     bool   `protobuf:"varint,3,opt,name=load,proto3" json:"load,omitempty"`
}

func (x *HealthResponse) Reset() {
	*x = HealthResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_accident_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthResponse) ProtoMessage() {}

func (x *HealthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_accident_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthResponse.ProtoReflect.Descriptor instead.
func (*HealthResponse) Descriptor() ([]byte, []int) {
	return file_internal_proto_accident_proto_rawDescGZIP(), []int{6}
}

func (x *HealthResponse) GetReady() bool {
	if x != nil {
		return x.Ready
	}
	return false
}

func (x *HealthResponse) GetDatabase() string {
	if x != nil {
		return x.Database
	}
	return ""
}

func (x *HealthResponse) GetLoad() bool {
	if x != nil {
		return x.Load
	}
	return false
}

type HealthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HealthRequest) Reset() {
	*x = HealthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_accident_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthRequest) ProtoMessage() {}

func (x *HealthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_accident_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthRequest.ProtoReflect.Descriptor instead.
func (*HealthRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_accident_proto_rawDescGZIP(), []int{7}
}

type FetchAccidentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserInfo *GetAccidentUserEntry `protobuf:"bytes,1,opt,name=user_info,json=userInfo,proto3" json:"user_info,omitempty"`
}

func (x *FetchAccidentRequest) Reset() {
	*x = FetchAccidentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_accident_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchAccidentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchAccidentRequest) ProtoMessage() {}

func (x *FetchAccidentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_accident_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchAccidentRequest.ProtoReflect.Descriptor instead.
func (*FetchAccidentRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_accident_proto_rawDescGZIP(), []int{8}
}

func (x *FetchAccidentRequest) GetUserInfo() *GetAccidentUserEntry {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

type PostAccidentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccidentInfo *PostAccidentEntry `protobuf:"bytes,1,opt,name=accident_info,json=accidentInfo,proto3" json:"accident_info,omitempty"`
}

func (x *PostAccidentRequest) Reset() {
	*x = PostAccidentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_accident_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostAccidentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostAccidentRequest) ProtoMessage() {}

func (x *PostAccidentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_accident_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostAccidentRequest.ProtoReflect.Descriptor instead.
func (*PostAccidentRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_accident_proto_rawDescGZIP(), []int{9}
}

func (x *PostAccidentRequest) GetAccidentInfo() *PostAccidentEntry {
	if x != nil {
		return x.AccidentInfo
	}
	return nil
}

type ConfirmAccidentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *ConfirmAccidentEntry `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *ConfirmAccidentRequest) Reset() {
	*x = ConfirmAccidentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_accident_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfirmAccidentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfirmAccidentRequest) ProtoMessage() {}

func (x *ConfirmAccidentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_accident_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfirmAccidentRequest.ProtoReflect.Descriptor instead.
func (*ConfirmAccidentRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_accident_proto_rawDescGZIP(), []int{10}
}

func (x *ConfirmAccidentRequest) GetInfo() *ConfirmAccidentEntry {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_internal_proto_accident_proto protoreflect.FileDescriptor

var file_internal_proto_accident_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x61, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x81, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x63,
	0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6c, 0x6f, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x6e, 0x67, 0x12, 0x19, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x6c, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x4c, 0x61, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x7a, 0x6f, 0x6f, 0x6d, 0x5f,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x7a, 0x6f, 0x6f,
	0x6d, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x22, 0x90, 0x02, 0x0a, 0x0c, 0x41,
	0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x23, 0x0a, 0x0d, 0x61,
	0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0c, 0x61, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x6e, 0x67,
	0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x61, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x4c, 0x61, 0x74, 0x12, 0x4c, 0x0a, 0x22, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x20, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x63,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x48, 0x0a, 0x20, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x65, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x42, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x42, 0x79, 0x22, 0x4f, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0d, 0x61, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x0c, 0x61, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0xb5,
	0x01, 0x0a, 0x11, 0x50, 0x6f, 0x73, 0x74, 0x41, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x5f, 0x6c, 0x6f, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x61, 0x63, 0x63,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x6e, 0x67, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x0b, 0x61, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x4c, 0x61, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79,
	0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x61, 0x72, 0x73, 0x5f, 0x69, 0x6e, 0x76, 0x6f, 0x6c, 0x76,
	0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x63, 0x61, 0x72, 0x73, 0x49, 0x6e,
	0x76, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x22, 0xc4, 0x01, 0x0a, 0x14, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x72, 0x6d, 0x41, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x23, 0x0a, 0x0d, 0x61, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x6e, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x61, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x4c, 0x6f, 0x6e, 0x67, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x5f, 0x6c, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x4c, 0x61, 0x74, 0x12, 0x2f, 0x0a, 0x13, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x33, 0x0a, 0x15, 0x61, 0x63, 0x63, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x14, 0x61, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x39, 0x0a,
	0x0f, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x56, 0x0a, 0x0e, 0x48, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65,
	0x61, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x72, 0x65, 0x61, 0x64, 0x79,
	0x12, 0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6c, 0x6f, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x6c, 0x6f, 0x61, 0x64,
	0x22, 0x0f, 0x0a, 0x0d, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x50, 0x0a, 0x14, 0x46, 0x65, 0x74, 0x63, 0x68, 0x41, 0x63, 0x63, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x22, 0x54, 0x0a, 0x13, 0x50, 0x6f, 0x73, 0x74, 0x41, 0x63, 0x63, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x0d, 0x61, 0x63,
	0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x41, 0x63,
	0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x61, 0x63, 0x63,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x49, 0x0a, 0x16, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x72, 0x6d, 0x41, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72,
	0x6d, 0x41, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04,
	0x69, 0x6e, 0x66, 0x6f, 0x32, 0xb7, 0x02, 0x0a, 0x18, 0x41, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x4b, 0x0a, 0x0e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x41, 0x63, 0x63, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x73, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x65, 0x74, 0x63,
	0x68, 0x41, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44,
	0x0a, 0x0c, 0x50, 0x6f, 0x73, 0x74, 0x41, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x1a,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x41, 0x63, 0x63, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x41,
	0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x41, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x3c, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12,
	0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x3e,
	0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x65, 0x6e,
	0x70, 0x61, 0x69, 0x6e, 0x69, 0x6b, 0x6f, 0x6c, 0x61, 0x79, 0x2f, 0x70, 0x61, 0x64, 0x2d, 0x73,
	0x65, 0x6d, 0x37, 0x2f, 0x61, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x2d, 0x72, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_proto_accident_proto_rawDescOnce sync.Once
	file_internal_proto_accident_proto_rawDescData = file_internal_proto_accident_proto_rawDesc
)

func file_internal_proto_accident_proto_rawDescGZIP() []byte {
	file_internal_proto_accident_proto_rawDescOnce.Do(func() {
		file_internal_proto_accident_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_proto_accident_proto_rawDescData)
	})
	return file_internal_proto_accident_proto_rawDescData
}

var file_internal_proto_accident_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_internal_proto_accident_proto_goTypes = []interface{}{
	(*GetAccidentUserEntry)(nil),   // 0: proto.GetAccidentUserEntry
	(*AccidentInfo)(nil),           // 1: proto.AccidentInfo
	(*GetAccidentResponse)(nil),    // 2: proto.GetAccidentResponse
	(*PostAccidentEntry)(nil),      // 3: proto.PostAccidentEntry
	(*ConfirmAccidentEntry)(nil),   // 4: proto.ConfirmAccidentEntry
	(*GenericResponse)(nil),        // 5: proto.GenericResponse
	(*HealthResponse)(nil),         // 6: proto.HealthResponse
	(*HealthRequest)(nil),          // 7: proto.HealthRequest
	(*FetchAccidentRequest)(nil),   // 8: proto.FetchAccidentRequest
	(*PostAccidentRequest)(nil),    // 9: proto.PostAccidentRequest
	(*ConfirmAccidentRequest)(nil), // 10: proto.ConfirmAccidentRequest
}
var file_internal_proto_accident_proto_depIdxs = []int32{
	1,  // 0: proto.GetAccidentResponse.accident_info:type_name -> proto.AccidentInfo
	0,  // 1: proto.FetchAccidentRequest.user_info:type_name -> proto.GetAccidentUserEntry
	3,  // 2: proto.PostAccidentRequest.accident_info:type_name -> proto.PostAccidentEntry
	4,  // 3: proto.ConfirmAccidentRequest.info:type_name -> proto.ConfirmAccidentEntry
	8,  // 4: proto.AccidentReportingService.FetchAccidents:input_type -> proto.FetchAccidentRequest
	9,  // 5: proto.AccidentReportingService.PostAccident:input_type -> proto.PostAccidentRequest
	10, // 6: proto.AccidentReportingService.ConfirmAccident:input_type -> proto.ConfirmAccidentRequest
	7,  // 7: proto.AccidentReportingService.HealthCheck:input_type -> proto.HealthRequest
	2,  // 8: proto.AccidentReportingService.FetchAccidents:output_type -> proto.GetAccidentResponse
	5,  // 9: proto.AccidentReportingService.PostAccident:output_type -> proto.GenericResponse
	5,  // 10: proto.AccidentReportingService.ConfirmAccident:output_type -> proto.GenericResponse
	6,  // 11: proto.AccidentReportingService.HealthCheck:output_type -> proto.HealthResponse
	8,  // [8:12] is the sub-list for method output_type
	4,  // [4:8] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_internal_proto_accident_proto_init() }
func file_internal_proto_accident_proto_init() {
	if File_internal_proto_accident_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_proto_accident_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccidentUserEntry); i {
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
		file_internal_proto_accident_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccidentInfo); i {
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
		file_internal_proto_accident_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccidentResponse); i {
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
		file_internal_proto_accident_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostAccidentEntry); i {
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
		file_internal_proto_accident_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfirmAccidentEntry); i {
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
		file_internal_proto_accident_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenericResponse); i {
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
		file_internal_proto_accident_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthResponse); i {
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
		file_internal_proto_accident_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthRequest); i {
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
		file_internal_proto_accident_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchAccidentRequest); i {
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
		file_internal_proto_accident_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostAccidentRequest); i {
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
		file_internal_proto_accident_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfirmAccidentRequest); i {
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
			RawDescriptor: file_internal_proto_accident_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_proto_accident_proto_goTypes,
		DependencyIndexes: file_internal_proto_accident_proto_depIdxs,
		MessageInfos:      file_internal_proto_accident_proto_msgTypes,
	}.Build()
	File_internal_proto_accident_proto = out.File
	file_internal_proto_accident_proto_rawDesc = nil
	file_internal_proto_accident_proto_goTypes = nil
	file_internal_proto_accident_proto_depIdxs = nil
}
