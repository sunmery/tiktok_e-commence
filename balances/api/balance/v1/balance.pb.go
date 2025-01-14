// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: balance/v1/balance.proto

package v1

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

type BalanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Owner   string  `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Name    string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Balance float64 `protobuf:"fixed64,3,opt,name=balance,proto3" json:"balance,omitempty"`
}

func (x *BalanceRequest) Reset() {
	*x = BalanceRequest{}
	mi := &file_balance_v1_balance_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BalanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BalanceRequest) ProtoMessage() {}

func (x *BalanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_balance_v1_balance_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BalanceRequest.ProtoReflect.Descriptor instead.
func (*BalanceRequest) Descriptor() ([]byte, []int) {
	return file_balance_v1_balance_proto_rawDescGZIP(), []int{0}
}

func (x *BalanceRequest) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *BalanceRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BalanceRequest) GetBalance() float64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

type BalanceReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string  `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Code    uint32  `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Owner   string  `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
	Name    string  `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Balance float64 `protobuf:"fixed64,5,opt,name=balance,proto3" json:"balance,omitempty"`
}

func (x *BalanceReply) Reset() {
	*x = BalanceReply{}
	mi := &file_balance_v1_balance_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BalanceReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BalanceReply) ProtoMessage() {}

func (x *BalanceReply) ProtoReflect() protoreflect.Message {
	mi := &file_balance_v1_balance_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BalanceReply.ProtoReflect.Descriptor instead.
func (*BalanceReply) Descriptor() ([]byte, []int) {
	return file_balance_v1_balance_proto_rawDescGZIP(), []int{1}
}

func (x *BalanceReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *BalanceReply) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *BalanceReply) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *BalanceReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BalanceReply) GetBalance() float64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

type GetBalanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetBalanceRequest) Reset() {
	*x = GetBalanceRequest{}
	mi := &file_balance_v1_balance_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBalanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBalanceRequest) ProtoMessage() {}

func (x *GetBalanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_balance_v1_balance_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBalanceRequest.ProtoReflect.Descriptor instead.
func (*GetBalanceRequest) Descriptor() ([]byte, []int) {
	return file_balance_v1_balance_proto_rawDescGZIP(), []int{2}
}

func (x *GetBalanceRequest) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *GetBalanceRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_balance_v1_balance_proto protoreflect.FileDescriptor

var file_balance_v1_balance_proto_rawDesc = []byte{
	0x0a, 0x18, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x61, 0x70, 0x69, 0x2e,
	0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x54, 0x0a, 0x0e, 0x42, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x80,
	0x01, 0x0a, 0x0c, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x22, 0x3d, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x32, 0xbe, 0x02, 0x0a, 0x07, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x66, 0x0a, 0x0d,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1e, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x17, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x11, 0x3a, 0x01, 0x2a, 0x22, 0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x73, 0x12, 0x66, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x3a, 0x01, 0x2a, 0x32, 0x0c,
	0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x63, 0x0a, 0x0a,
	0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x21, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x14, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x73, 0x42, 0x2e, 0x0a, 0x0e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x1a, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_balance_v1_balance_proto_rawDescOnce sync.Once
	file_balance_v1_balance_proto_rawDescData = file_balance_v1_balance_proto_rawDesc
)

func file_balance_v1_balance_proto_rawDescGZIP() []byte {
	file_balance_v1_balance_proto_rawDescOnce.Do(func() {
		file_balance_v1_balance_proto_rawDescData = protoimpl.X.CompressGZIP(file_balance_v1_balance_proto_rawDescData)
	})
	return file_balance_v1_balance_proto_rawDescData
}

var file_balance_v1_balance_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_balance_v1_balance_proto_goTypes = []any{
	(*BalanceRequest)(nil),    // 0: api.balance.v1.BalanceRequest
	(*BalanceReply)(nil),      // 1: api.balance.v1.BalanceReply
	(*GetBalanceRequest)(nil), // 2: api.balance.v1.GetBalanceRequest
}
var file_balance_v1_balance_proto_depIdxs = []int32{
	0, // 0: api.balance.v1.Balance.CreateBalance:input_type -> api.balance.v1.BalanceRequest
	0, // 1: api.balance.v1.Balance.UpdateBalance:input_type -> api.balance.v1.BalanceRequest
	2, // 2: api.balance.v1.Balance.GetBalance:input_type -> api.balance.v1.GetBalanceRequest
	1, // 3: api.balance.v1.Balance.CreateBalance:output_type -> api.balance.v1.BalanceReply
	1, // 4: api.balance.v1.Balance.UpdateBalance:output_type -> api.balance.v1.BalanceReply
	1, // 5: api.balance.v1.Balance.GetBalance:output_type -> api.balance.v1.BalanceReply
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_balance_v1_balance_proto_init() }
func file_balance_v1_balance_proto_init() {
	if File_balance_v1_balance_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_balance_v1_balance_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_balance_v1_balance_proto_goTypes,
		DependencyIndexes: file_balance_v1_balance_proto_depIdxs,
		MessageInfos:      file_balance_v1_balance_proto_msgTypes,
	}.Build()
	File_balance_v1_balance_proto = out.File
	file_balance_v1_balance_proto_rawDesc = nil
	file_balance_v1_balance_proto_goTypes = nil
	file_balance_v1_balance_proto_depIdxs = nil
}
