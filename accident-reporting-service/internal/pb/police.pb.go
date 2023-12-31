// protoc --go_out=internal/pb --go_opt=paths=source_relative --go-grpc_out=internal/pb --go-grpc_opt=paths=source_relative internal/proto/police.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.3
// source: internal/proto/police.proto

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

type GetPoliceUserEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserLong  float64 `protobuf:"fixed64,1,opt,name=user_long,json=userLong,proto3" json:"user_long,omitempty"`
	UserLat   float64 `protobuf:"fixed64,2,opt,name=user_lat,json=userLat,proto3" json:"user_lat,omitempty"`
	ZoomIndex int64   `protobuf:"varint,3,opt,name=zoom_index,json=zoomIndex,proto3" json:"zoom_index,omitempty"`
	City      string  `protobuf:"bytes,4,opt,name=city,proto3" json:"city,omitempty"`
}

func (x *GetPoliceUserEntry) Reset() {
	*x = GetPoliceUserEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_police_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPoliceUserEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPoliceUserEntry) ProtoMessage() {}

func (x *GetPoliceUserEntry) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_police_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPoliceUserEntry.ProtoReflect.Descriptor instead.
func (*GetPoliceUserEntry) Descriptor() ([]byte, []int) {
	return file_internal_proto_police_proto_rawDescGZIP(), []int{0}
}

func (x *GetPoliceUserEntry) GetUserLong() float64 {
	if x != nil {
		return x.UserLong
	}
	return 0
}

func (x *GetPoliceUserEntry) GetUserLat() float64 {
	if x != nil {
		return x.UserLat
	}
	return 0
}

func (x *GetPoliceUserEntry) GetZoomIndex() int64 {
	if x != nil {
		return x.ZoomIndex
	}
	return 0
}

func (x *GetPoliceUserEntry) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

type PostPoliceEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PolLong float64 `protobuf:"fixed64,1,opt,name=pol_long,json=polLong,proto3" json:"pol_long,omitempty"`
	PolLat  float64 `protobuf:"fixed64,2,opt,name=pol_lat,json=polLat,proto3" json:"pol_lat,omitempty"`
	City    string  `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`
}

func (x *PostPoliceEntry) Reset() {
	*x = PostPoliceEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_police_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostPoliceEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostPoliceEntry) ProtoMessage() {}

func (x *PostPoliceEntry) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_police_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostPoliceEntry.ProtoReflect.Descriptor instead.
func (*PostPoliceEntry) Descriptor() ([]byte, []int) {
	return file_internal_proto_police_proto_rawDescGZIP(), []int{1}
}

func (x *PostPoliceEntry) GetPolLong() float64 {
	if x != nil {
		return x.PolLong
	}
	return 0
}

func (x *PostPoliceEntry) GetPolLat() float64 {
	if x != nil {
		return x.PolLat
	}
	return 0
}

func (x *PostPoliceEntry) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

type ConfirmPoliceEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PolLong      float64 `protobuf:"fixed64,1,opt,name=pol_long,json=polLong,proto3" json:"pol_long,omitempty"`
	PolLat       float64 `protobuf:"fixed64,2,opt,name=pol_lat,json=polLat,proto3" json:"pol_lat,omitempty"`
	City         string  `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`
	Confirmation bool    `protobuf:"varint,4,opt,name=confirmation,proto3" json:"confirmation,omitempty"`
}

func (x *ConfirmPoliceEntry) Reset() {
	*x = ConfirmPoliceEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_police_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfirmPoliceEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfirmPoliceEntry) ProtoMessage() {}

func (x *ConfirmPoliceEntry) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_police_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfirmPoliceEntry.ProtoReflect.Descriptor instead.
func (*ConfirmPoliceEntry) Descriptor() ([]byte, []int) {
	return file_internal_proto_police_proto_rawDescGZIP(), []int{2}
}

func (x *ConfirmPoliceEntry) GetPolLong() float64 {
	if x != nil {
		return x.PolLong
	}
	return 0
}

func (x *ConfirmPoliceEntry) GetPolLat() float64 {
	if x != nil {
		return x.PolLat
	}
	return 0
}

func (x *ConfirmPoliceEntry) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *ConfirmPoliceEntry) GetConfirmation() bool {
	if x != nil {
		return x.Confirmation
	}
	return false
}

type PoliceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error bool   `protobuf:"varint,1,opt,name=error,proto3" json:"error,omitempty"`
	Msg   string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *PoliceResponse) Reset() {
	*x = PoliceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_police_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PoliceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PoliceResponse) ProtoMessage() {}

func (x *PoliceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_police_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PoliceResponse.ProtoReflect.Descriptor instead.
func (*PoliceResponse) Descriptor() ([]byte, []int) {
	return file_internal_proto_police_proto_rawDescGZIP(), []int{3}
}

func (x *PoliceResponse) GetError() bool {
	if x != nil {
		return x.Error
	}
	return false
}

func (x *PoliceResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type PoliceInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PolLong                  float64 `protobuf:"fixed64,1,opt,name=pol_long,json=polLong,proto3" json:"pol_long,omitempty"`
	PolLat                   float64 `protobuf:"fixed64,2,opt,name=pol_lat,json=polLat,proto3" json:"pol_lat,omitempty"`
	ConfirmationNotification bool    `protobuf:"varint,3,opt,name=confirmation_notification,json=confirmationNotification,proto3" json:"confirmation_notification,omitempty"`
	ConfirmedBy              int64   `protobuf:"varint,4,opt,name=confirmedBy,proto3" json:"confirmedBy,omitempty"`
}

func (x *PoliceInfo) Reset() {
	*x = PoliceInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_police_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PoliceInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PoliceInfo) ProtoMessage() {}

func (x *PoliceInfo) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_police_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PoliceInfo.ProtoReflect.Descriptor instead.
func (*PoliceInfo) Descriptor() ([]byte, []int) {
	return file_internal_proto_police_proto_rawDescGZIP(), []int{4}
}

func (x *PoliceInfo) GetPolLong() float64 {
	if x != nil {
		return x.PolLong
	}
	return 0
}

func (x *PoliceInfo) GetPolLat() float64 {
	if x != nil {
		return x.PolLat
	}
	return 0
}

func (x *PoliceInfo) GetConfirmationNotification() bool {
	if x != nil {
		return x.ConfirmationNotification
	}
	return false
}

func (x *PoliceInfo) GetConfirmedBy() int64 {
	if x != nil {
		return x.ConfirmedBy
	}
	return 0
}

type GetPoliceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PoliceInfo []*PoliceInfo `protobuf:"bytes,1,rep,name=police_info,json=policeInfo,proto3" json:"police_info,omitempty"`
}

func (x *GetPoliceResponse) Reset() {
	*x = GetPoliceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_police_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPoliceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPoliceResponse) ProtoMessage() {}

func (x *GetPoliceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_police_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPoliceResponse.ProtoReflect.Descriptor instead.
func (*GetPoliceResponse) Descriptor() ([]byte, []int) {
	return file_internal_proto_police_proto_rawDescGZIP(), []int{5}
}

func (x *GetPoliceResponse) GetPoliceInfo() []*PoliceInfo {
	if x != nil {
		return x.PoliceInfo
	}
	return nil
}

type FetchPoliceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserInfo *GetPoliceUserEntry `protobuf:"bytes,1,opt,name=user_info,json=userInfo,proto3" json:"user_info,omitempty"`
}

func (x *FetchPoliceRequest) Reset() {
	*x = FetchPoliceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_police_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchPoliceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchPoliceRequest) ProtoMessage() {}

func (x *FetchPoliceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_police_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchPoliceRequest.ProtoReflect.Descriptor instead.
func (*FetchPoliceRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_police_proto_rawDescGZIP(), []int{6}
}

func (x *FetchPoliceRequest) GetUserInfo() *GetPoliceUserEntry {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

type PostPoliceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PoliceInfo *PostPoliceEntry `protobuf:"bytes,1,opt,name=police_info,json=policeInfo,proto3" json:"police_info,omitempty"`
}

func (x *PostPoliceRequest) Reset() {
	*x = PostPoliceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_police_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostPoliceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostPoliceRequest) ProtoMessage() {}

func (x *PostPoliceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_police_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostPoliceRequest.ProtoReflect.Descriptor instead.
func (*PostPoliceRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_police_proto_rawDescGZIP(), []int{7}
}

func (x *PostPoliceRequest) GetPoliceInfo() *PostPoliceEntry {
	if x != nil {
		return x.PoliceInfo
	}
	return nil
}

type ConfirmPoliceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PoliceInfo *ConfirmPoliceEntry `protobuf:"bytes,1,opt,name=police_info,json=policeInfo,proto3" json:"police_info,omitempty"`
}

func (x *ConfirmPoliceRequest) Reset() {
	*x = ConfirmPoliceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_police_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfirmPoliceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfirmPoliceRequest) ProtoMessage() {}

func (x *ConfirmPoliceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_police_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfirmPoliceRequest.ProtoReflect.Descriptor instead.
func (*ConfirmPoliceRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_police_proto_rawDescGZIP(), []int{8}
}

func (x *ConfirmPoliceRequest) GetPoliceInfo() *ConfirmPoliceEntry {
	if x != nil {
		return x.PoliceInfo
	}
	return nil
}

var File_internal_proto_police_proto protoreflect.FileDescriptor

var file_internal_proto_police_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7f, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x6c, 0x6f, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x4c, 0x6f, 0x6e, 0x67, 0x12, 0x19, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x6c, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x4c,
	0x61, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x7a, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x7a, 0x6f, 0x6f, 0x6d, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x63, 0x69, 0x74, 0x79, 0x22, 0x59, 0x0a, 0x0f, 0x50, 0x6f, 0x73, 0x74, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x6f, 0x6c, 0x5f,
	0x6c, 0x6f, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x70, 0x6f, 0x6c, 0x4c,
	0x6f, 0x6e, 0x67, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f, 0x6c, 0x5f, 0x6c, 0x61, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x70, 0x6f, 0x6c, 0x4c, 0x61, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79,
	0x22, 0x80, 0x01, 0x0a, 0x12, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x6f, 0x6c, 0x5f, 0x6c,
	0x6f, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x70, 0x6f, 0x6c, 0x4c, 0x6f,
	0x6e, 0x67, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f, 0x6c, 0x5f, 0x6c, 0x61, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x06, 0x70, 0x6f, 0x6c, 0x4c, 0x61, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12,
	0x22, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x38, 0x0a, 0x0e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x9f, 0x01,
	0x0a, 0x0a, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19, 0x0a, 0x08,
	0x70, 0x6f, 0x6c, 0x5f, 0x6c, 0x6f, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07,
	0x70, 0x6f, 0x6c, 0x4c, 0x6f, 0x6e, 0x67, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f, 0x6c, 0x5f, 0x6c,
	0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x70, 0x6f, 0x6c, 0x4c, 0x61, 0x74,
	0x12, 0x3b, 0x0a, 0x19, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x18, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a,
	0x0b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x42, 0x79, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x42, 0x79, 0x22,
	0x47, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x0b, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x65, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x4c, 0x0a, 0x12, 0x46, 0x65, 0x74, 0x63,
	0x68, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36,
	0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x65, 0x55, 0x73, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x4c, 0x0a, 0x11, 0x50, 0x6f, 0x73, 0x74, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x0b, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x22, 0x52, 0x0a, 0x14, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x0b,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72,
	0x6d, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x32, 0xe6, 0x01, 0x0a, 0x16, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x0b, 0x46, 0x65, 0x74, 0x63, 0x68, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x65, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x0a, 0x50, 0x6f, 0x73,
	0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x65, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x50, 0x6f, 0x73, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0d, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x72, 0x6d, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x65, 0x6e, 0x70, 0x61, 0x69, 0x6e, 0x69, 0x6b, 0x6f, 0x6c, 0x61, 0x79, 0x2f, 0x70, 0x61,
	0x64, 0x2d, 0x73, 0x65, 0x6d, 0x37, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_proto_police_proto_rawDescOnce sync.Once
	file_internal_proto_police_proto_rawDescData = file_internal_proto_police_proto_rawDesc
)

func file_internal_proto_police_proto_rawDescGZIP() []byte {
	file_internal_proto_police_proto_rawDescOnce.Do(func() {
		file_internal_proto_police_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_proto_police_proto_rawDescData)
	})
	return file_internal_proto_police_proto_rawDescData
}

var file_internal_proto_police_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_internal_proto_police_proto_goTypes = []interface{}{
	(*GetPoliceUserEntry)(nil),   // 0: proto.GetPoliceUserEntry
	(*PostPoliceEntry)(nil),      // 1: proto.PostPoliceEntry
	(*ConfirmPoliceEntry)(nil),   // 2: proto.ConfirmPoliceEntry
	(*PoliceResponse)(nil),       // 3: proto.PoliceResponse
	(*PoliceInfo)(nil),           // 4: proto.PoliceInfo
	(*GetPoliceResponse)(nil),    // 5: proto.GetPoliceResponse
	(*FetchPoliceRequest)(nil),   // 6: proto.FetchPoliceRequest
	(*PostPoliceRequest)(nil),    // 7: proto.PostPoliceRequest
	(*ConfirmPoliceRequest)(nil), // 8: proto.ConfirmPoliceRequest
}
var file_internal_proto_police_proto_depIdxs = []int32{
	4, // 0: proto.GetPoliceResponse.police_info:type_name -> proto.PoliceInfo
	0, // 1: proto.FetchPoliceRequest.user_info:type_name -> proto.GetPoliceUserEntry
	1, // 2: proto.PostPoliceRequest.police_info:type_name -> proto.PostPoliceEntry
	2, // 3: proto.ConfirmPoliceRequest.police_info:type_name -> proto.ConfirmPoliceEntry
	6, // 4: proto.PoliceReportingService.FetchPolice:input_type -> proto.FetchPoliceRequest
	7, // 5: proto.PoliceReportingService.PostPolice:input_type -> proto.PostPoliceRequest
	8, // 6: proto.PoliceReportingService.ConfirmPolice:input_type -> proto.ConfirmPoliceRequest
	5, // 7: proto.PoliceReportingService.FetchPolice:output_type -> proto.GetPoliceResponse
	3, // 8: proto.PoliceReportingService.PostPolice:output_type -> proto.PoliceResponse
	3, // 9: proto.PoliceReportingService.ConfirmPolice:output_type -> proto.PoliceResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_internal_proto_police_proto_init() }
func file_internal_proto_police_proto_init() {
	if File_internal_proto_police_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_proto_police_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPoliceUserEntry); i {
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
		file_internal_proto_police_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostPoliceEntry); i {
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
		file_internal_proto_police_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfirmPoliceEntry); i {
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
		file_internal_proto_police_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PoliceResponse); i {
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
		file_internal_proto_police_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PoliceInfo); i {
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
		file_internal_proto_police_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPoliceResponse); i {
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
		file_internal_proto_police_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchPoliceRequest); i {
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
		file_internal_proto_police_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostPoliceRequest); i {
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
		file_internal_proto_police_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfirmPoliceRequest); i {
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
			RawDescriptor: file_internal_proto_police_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_proto_police_proto_goTypes,
		DependencyIndexes: file_internal_proto_police_proto_depIdxs,
		MessageInfos:      file_internal_proto_police_proto_msgTypes,
	}.Build()
	File_internal_proto_police_proto = out.File
	file_internal_proto_police_proto_rawDesc = nil
	file_internal_proto_police_proto_goTypes = nil
	file_internal_proto_police_proto_depIdxs = nil
}
