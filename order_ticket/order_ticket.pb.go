// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.6
// source: order_ticket.proto

package __

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

type AddOrderTicketReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId  int64 `protobuf:"varint,1,opt,name=OrderId,proto3" json:"OrderId,omitempty"`
	TicketId int64 `protobuf:"varint,2,opt,name=TicketId,proto3" json:"TicketId,omitempty"`
}

func (x *AddOrderTicketReq) Reset() {
	*x = AddOrderTicketReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_ticket_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddOrderTicketReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddOrderTicketReq) ProtoMessage() {}

func (x *AddOrderTicketReq) ProtoReflect() protoreflect.Message {
	mi := &file_order_ticket_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddOrderTicketReq.ProtoReflect.Descriptor instead.
func (*AddOrderTicketReq) Descriptor() ([]byte, []int) {
	return file_order_ticket_proto_rawDescGZIP(), []int{0}
}

func (x *AddOrderTicketReq) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *AddOrderTicketReq) GetTicketId() int64 {
	if x != nil {
		return x.TicketId
	}
	return 0
}

type AddOrderTicketResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success int64 `protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty"`
}

func (x *AddOrderTicketResp) Reset() {
	*x = AddOrderTicketResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_ticket_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddOrderTicketResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddOrderTicketResp) ProtoMessage() {}

func (x *AddOrderTicketResp) ProtoReflect() protoreflect.Message {
	mi := &file_order_ticket_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddOrderTicketResp.ProtoReflect.Descriptor instead.
func (*AddOrderTicketResp) Descriptor() ([]byte, []int) {
	return file_order_ticket_proto_rawDescGZIP(), []int{1}
}

func (x *AddOrderTicketResp) GetSuccess() int64 {
	if x != nil {
		return x.Success
	}
	return 0
}

type DelOrderTicketReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId  int64 `protobuf:"varint,1,opt,name=OrderId,proto3" json:"OrderId,omitempty"`
	TicketId int64 `protobuf:"varint,2,opt,name=TicketId,proto3" json:"TicketId,omitempty"`
}

func (x *DelOrderTicketReq) Reset() {
	*x = DelOrderTicketReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_ticket_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelOrderTicketReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelOrderTicketReq) ProtoMessage() {}

func (x *DelOrderTicketReq) ProtoReflect() protoreflect.Message {
	mi := &file_order_ticket_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelOrderTicketReq.ProtoReflect.Descriptor instead.
func (*DelOrderTicketReq) Descriptor() ([]byte, []int) {
	return file_order_ticket_proto_rawDescGZIP(), []int{2}
}

func (x *DelOrderTicketReq) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *DelOrderTicketReq) GetTicketId() int64 {
	if x != nil {
		return x.TicketId
	}
	return 0
}

type DelOrderTicketResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success int64 `protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty"`
}

func (x *DelOrderTicketResp) Reset() {
	*x = DelOrderTicketResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_ticket_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelOrderTicketResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelOrderTicketResp) ProtoMessage() {}

func (x *DelOrderTicketResp) ProtoReflect() protoreflect.Message {
	mi := &file_order_ticket_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelOrderTicketResp.ProtoReflect.Descriptor instead.
func (*DelOrderTicketResp) Descriptor() ([]byte, []int) {
	return file_order_ticket_proto_rawDescGZIP(), []int{3}
}

func (x *DelOrderTicketResp) GetSuccess() int64 {
	if x != nil {
		return x.Success
	}
	return 0
}

type GetOTDetailReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId int64 `protobuf:"varint,1,opt,name=OrderId,proto3" json:"OrderId,omitempty"`
}

func (x *GetOTDetailReq) Reset() {
	*x = GetOTDetailReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_ticket_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOTDetailReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOTDetailReq) ProtoMessage() {}

func (x *GetOTDetailReq) ProtoReflect() protoreflect.Message {
	mi := &file_order_ticket_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOTDetailReq.ProtoReflect.Descriptor instead.
func (*GetOTDetailReq) Descriptor() ([]byte, []int) {
	return file_order_ticket_proto_rawDescGZIP(), []int{4}
}

func (x *GetOTDetailReq) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

type OTInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId    int64  `protobuf:"varint,1,opt,name=OrderId,proto3" json:"OrderId,omitempty"`
	TicketId   int64  `protobuf:"varint,2,opt,name=TicketId,proto3" json:"TicketId,omitempty"`
	CreateTime string `protobuf:"bytes,3,opt,name=CreateTime,proto3" json:"CreateTime,omitempty"`
	UpdateTime string `protobuf:"bytes,4,opt,name=UpdateTime,proto3" json:"UpdateTime,omitempty"`
}

func (x *OTInfo) Reset() {
	*x = OTInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_ticket_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OTInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OTInfo) ProtoMessage() {}

func (x *OTInfo) ProtoReflect() protoreflect.Message {
	mi := &file_order_ticket_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OTInfo.ProtoReflect.Descriptor instead.
func (*OTInfo) Descriptor() ([]byte, []int) {
	return file_order_ticket_proto_rawDescGZIP(), []int{5}
}

func (x *OTInfo) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *OTInfo) GetTicketId() int64 {
	if x != nil {
		return x.TicketId
	}
	return 0
}

func (x *OTInfo) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *OTInfo) GetUpdateTime() string {
	if x != nil {
		return x.UpdateTime
	}
	return ""
}

type GetOTDetailResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OTList []*OTInfo `protobuf:"bytes,1,rep,name=OTList,proto3" json:"OTList,omitempty"`
}

func (x *GetOTDetailResp) Reset() {
	*x = GetOTDetailResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_ticket_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOTDetailResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOTDetailResp) ProtoMessage() {}

func (x *GetOTDetailResp) ProtoReflect() protoreflect.Message {
	mi := &file_order_ticket_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOTDetailResp.ProtoReflect.Descriptor instead.
func (*GetOTDetailResp) Descriptor() ([]byte, []int) {
	return file_order_ticket_proto_rawDescGZIP(), []int{6}
}

func (x *GetOTDetailResp) GetOTList() []*OTInfo {
	if x != nil {
		return x.OTList
	}
	return nil
}

var File_order_ticket_proto protoreflect.FileDescriptor

var file_order_ticket_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x49, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x22,
	0x2e, 0x0a, 0x12, 0x41, 0x64, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22,
	0x49, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x22, 0x2e, 0x0a, 0x12, 0x44, 0x65,
	0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x2a, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x4f, 0x54, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x7e, 0x0a, 0x06, 0x4f, 0x54, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x18, 0x0a, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x54, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x32, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4f, 0x54, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1f, 0x0a, 0x06, 0x4f, 0x54, 0x4c,
	0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x4f, 0x54, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x06, 0x4f, 0x54, 0x4c, 0x69, 0x73, 0x74, 0x32, 0xbc, 0x01, 0x0a, 0x12, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x39, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x12, 0x12, 0x2e, 0x41, 0x64, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x41, 0x64, 0x64, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x39, 0x0a, 0x0e,
	0x44, 0x65, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x12,
	0x2e, 0x44, 0x65, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x1a, 0x13, 0x2e, 0x44, 0x65, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x30, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4f, 0x54,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x0f, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x54, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x54, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_order_ticket_proto_rawDescOnce sync.Once
	file_order_ticket_proto_rawDescData = file_order_ticket_proto_rawDesc
)

func file_order_ticket_proto_rawDescGZIP() []byte {
	file_order_ticket_proto_rawDescOnce.Do(func() {
		file_order_ticket_proto_rawDescData = protoimpl.X.CompressGZIP(file_order_ticket_proto_rawDescData)
	})
	return file_order_ticket_proto_rawDescData
}

var file_order_ticket_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_order_ticket_proto_goTypes = []interface{}{
	(*AddOrderTicketReq)(nil),  // 0: AddOrderTicketReq
	(*AddOrderTicketResp)(nil), // 1: AddOrderTicketResp
	(*DelOrderTicketReq)(nil),  // 2: DelOrderTicketReq
	(*DelOrderTicketResp)(nil), // 3: DelOrderTicketResp
	(*GetOTDetailReq)(nil),     // 4: GetOTDetailReq
	(*OTInfo)(nil),             // 5: OTInfo
	(*GetOTDetailResp)(nil),    // 6: GetOTDetailResp
}
var file_order_ticket_proto_depIdxs = []int32{
	5, // 0: GetOTDetailResp.OTList:type_name -> OTInfo
	0, // 1: OrderTicketService.AddOrderTicket:input_type -> AddOrderTicketReq
	2, // 2: OrderTicketService.DelOrderTicket:input_type -> DelOrderTicketReq
	4, // 3: OrderTicketService.GetOTDetail:input_type -> GetOTDetailReq
	1, // 4: OrderTicketService.AddOrderTicket:output_type -> AddOrderTicketResp
	3, // 5: OrderTicketService.DelOrderTicket:output_type -> DelOrderTicketResp
	6, // 6: OrderTicketService.GetOTDetail:output_type -> GetOTDetailResp
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_order_ticket_proto_init() }
func file_order_ticket_proto_init() {
	if File_order_ticket_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_order_ticket_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddOrderTicketReq); i {
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
		file_order_ticket_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddOrderTicketResp); i {
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
		file_order_ticket_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelOrderTicketReq); i {
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
		file_order_ticket_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelOrderTicketResp); i {
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
		file_order_ticket_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOTDetailReq); i {
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
		file_order_ticket_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OTInfo); i {
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
		file_order_ticket_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOTDetailResp); i {
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
			RawDescriptor: file_order_ticket_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_order_ticket_proto_goTypes,
		DependencyIndexes: file_order_ticket_proto_depIdxs,
		MessageInfos:      file_order_ticket_proto_msgTypes,
	}.Build()
	File_order_ticket_proto = out.File
	file_order_ticket_proto_rawDesc = nil
	file_order_ticket_proto_goTypes = nil
	file_order_ticket_proto_depIdxs = nil
}
