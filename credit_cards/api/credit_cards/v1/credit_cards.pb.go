// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: credit_cards/v1/credit_cards.proto

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

type CreditCards struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number          string `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	Cvv             string `protobuf:"bytes,2,opt,name=cvv,proto3" json:"cvv,omitempty"`
	ExpirationYear  string `protobuf:"bytes,3,opt,name=expiration_year,proto3" json:"expiration_year,omitempty"`
	ExpirationMonth string `protobuf:"bytes,4,opt,name=expiration_month,proto3" json:"expiration_month,omitempty"`
	Owner           string `protobuf:"bytes,5,opt,name=owner,proto3" json:"owner,omitempty"`
	Name            string `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	Id              uint32 `protobuf:"varint,7,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreditCards) Reset() {
	*x = CreditCards{}
	mi := &file_credit_cards_v1_credit_cards_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreditCards) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreditCards) ProtoMessage() {}

func (x *CreditCards) ProtoReflect() protoreflect.Message {
	mi := &file_credit_cards_v1_credit_cards_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreditCards.ProtoReflect.Descriptor instead.
func (*CreditCards) Descriptor() ([]byte, []int) {
	return file_credit_cards_v1_credit_cards_proto_rawDescGZIP(), []int{0}
}

func (x *CreditCards) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *CreditCards) GetCvv() string {
	if x != nil {
		return x.Cvv
	}
	return ""
}

func (x *CreditCards) GetExpirationYear() string {
	if x != nil {
		return x.ExpirationYear
	}
	return ""
}

func (x *CreditCards) GetExpirationMonth() string {
	if x != nil {
		return x.ExpirationMonth
	}
	return ""
}

func (x *CreditCards) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *CreditCards) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreditCards) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type CardsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Code    int32  `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *CardsReply) Reset() {
	*x = CardsReply{}
	mi := &file_credit_cards_v1_credit_cards_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CardsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CardsReply) ProtoMessage() {}

func (x *CardsReply) ProtoReflect() protoreflect.Message {
	mi := &file_credit_cards_v1_credit_cards_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CardsReply.ProtoReflect.Descriptor instead.
func (*CardsReply) Descriptor() ([]byte, []int) {
	return file_credit_cards_v1_credit_cards_proto_rawDescGZIP(), []int{1}
}

func (x *CardsReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CardsReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

type DeleteCreditCardsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteCreditCardsRequest) Reset() {
	*x = DeleteCreditCardsRequest{}
	mi := &file_credit_cards_v1_credit_cards_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteCreditCardsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCreditCardsRequest) ProtoMessage() {}

func (x *DeleteCreditCardsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_credit_cards_v1_credit_cards_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCreditCardsRequest.ProtoReflect.Descriptor instead.
func (*DeleteCreditCardsRequest) Descriptor() ([]byte, []int) {
	return file_credit_cards_v1_credit_cards_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteCreditCardsRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetCreditCardsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Owner  string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Number string `protobuf:"bytes,3,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *GetCreditCardsRequest) Reset() {
	*x = GetCreditCardsRequest{}
	mi := &file_credit_cards_v1_credit_cards_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCreditCardsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCreditCardsRequest) ProtoMessage() {}

func (x *GetCreditCardsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_credit_cards_v1_credit_cards_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCreditCardsRequest.ProtoReflect.Descriptor instead.
func (*GetCreditCardsRequest) Descriptor() ([]byte, []int) {
	return file_credit_cards_v1_credit_cards_proto_rawDescGZIP(), []int{3}
}

func (x *GetCreditCardsRequest) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *GetCreditCardsRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetCreditCardsRequest) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

type GetCreditCardsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreditCards *CreditCards `protobuf:"bytes,1,opt,name=credit_cards,proto3" json:"credit_cards,omitempty"`
}

func (x *GetCreditCardsReply) Reset() {
	*x = GetCreditCardsReply{}
	mi := &file_credit_cards_v1_credit_cards_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCreditCardsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCreditCardsReply) ProtoMessage() {}

func (x *GetCreditCardsReply) ProtoReflect() protoreflect.Message {
	mi := &file_credit_cards_v1_credit_cards_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCreditCardsReply.ProtoReflect.Descriptor instead.
func (*GetCreditCardsReply) Descriptor() ([]byte, []int) {
	return file_credit_cards_v1_credit_cards_proto_rawDescGZIP(), []int{4}
}

func (x *GetCreditCardsReply) GetCreditCards() *CreditCards {
	if x != nil {
		return x.CreditCards
	}
	return nil
}

type SearchCreditCardsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreditCards []*CreditCards `protobuf:"bytes,1,rep,name=credit_cards,proto3" json:"credit_cards,omitempty"`
}

func (x *SearchCreditCardsReply) Reset() {
	*x = SearchCreditCardsReply{}
	mi := &file_credit_cards_v1_credit_cards_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchCreditCardsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchCreditCardsReply) ProtoMessage() {}

func (x *SearchCreditCardsReply) ProtoReflect() protoreflect.Message {
	mi := &file_credit_cards_v1_credit_cards_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchCreditCardsReply.ProtoReflect.Descriptor instead.
func (*SearchCreditCardsReply) Descriptor() ([]byte, []int) {
	return file_credit_cards_v1_credit_cards_proto_rawDescGZIP(), []int{5}
}

func (x *SearchCreditCardsReply) GetCreditCards() []*CreditCards {
	if x != nil {
		return x.CreditCards
	}
	return nil
}

type ListCreditCardsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ListCreditCardsRequest) Reset() {
	*x = ListCreditCardsRequest{}
	mi := &file_credit_cards_v1_credit_cards_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListCreditCardsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCreditCardsRequest) ProtoMessage() {}

func (x *ListCreditCardsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_credit_cards_v1_credit_cards_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCreditCardsRequest.ProtoReflect.Descriptor instead.
func (*ListCreditCardsRequest) Descriptor() ([]byte, []int) {
	return file_credit_cards_v1_credit_cards_proto_rawDescGZIP(), []int{6}
}

func (x *ListCreditCardsRequest) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *ListCreditCardsRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ListCreditCardsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreditCards []*CreditCards `protobuf:"bytes,1,rep,name=credit_cards,proto3" json:"credit_cards,omitempty"`
}

func (x *ListCreditCardsReply) Reset() {
	*x = ListCreditCardsReply{}
	mi := &file_credit_cards_v1_credit_cards_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListCreditCardsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCreditCardsReply) ProtoMessage() {}

func (x *ListCreditCardsReply) ProtoReflect() protoreflect.Message {
	mi := &file_credit_cards_v1_credit_cards_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCreditCardsReply.ProtoReflect.Descriptor instead.
func (*ListCreditCardsReply) Descriptor() ([]byte, []int) {
	return file_credit_cards_v1_credit_cards_proto_rawDescGZIP(), []int{7}
}

func (x *ListCreditCardsReply) GetCreditCards() []*CreditCards {
	if x != nil {
		return x.CreditCards
	}
	return nil
}

var File_credit_cards_v1_credit_cards_proto protoreflect.FileDescriptor

var file_credit_cards_v1_credit_cards_proto_rawDesc = []byte{
	0x0a, 0x22, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74,
	0x5f, 0x63, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc7, 0x01, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x64,
	0x69, 0x74, 0x43, 0x61, 0x72, 0x64, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x10, 0x0a, 0x03, 0x63, 0x76, 0x76, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x76,
	0x76, 0x12, 0x28, 0x0a, 0x0f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x79, 0x65, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x79, 0x65, 0x61, 0x72, 0x12, 0x2a, 0x0a, 0x10, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x3a, 0x0a, 0x0a, 0x43, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x2a, 0x0a,
	0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x43, 0x61, 0x72,
	0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x59, 0x0a, 0x15, 0x47, 0x65, 0x74,
	0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x43, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x22, 0x5b, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x72, 0x65, 0x64, 0x69,
	0x74, 0x43, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x44, 0x0a, 0x0c, 0x63,
	0x72, 0x65, 0x64, 0x69, 0x74, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x5f, 0x63,
	0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x43, 0x61,
	0x72, 0x64, 0x73, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x5f, 0x63, 0x61, 0x72, 0x64,
	0x73, 0x22, 0x5e, 0x0a, 0x16, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x43, 0x72, 0x65, 0x64, 0x69,
	0x74, 0x43, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x44, 0x0a, 0x0c, 0x63,
	0x72, 0x65, 0x64, 0x69, 0x74, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x5f, 0x63,
	0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x43, 0x61,
	0x72, 0x64, 0x73, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x5f, 0x63, 0x61, 0x72, 0x64,
	0x73, 0x22, 0x42, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x43,
	0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x5c, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x72, 0x65,
	0x64, 0x69, 0x74, 0x43, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x44, 0x0a,
	0x0c, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74,
	0x5f, 0x63, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74,
	0x43, 0x61, 0x72, 0x64, 0x73, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x5f, 0x63, 0x61,
	0x72, 0x64, 0x73, 0x32, 0x95, 0x05, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x43, 0x61,
	0x72, 0x64, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x72, 0x0a, 0x10, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x43, 0x61, 0x72, 0x64, 0x12, 0x20,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x5f, 0x63, 0x61, 0x72, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x43, 0x61, 0x72, 0x64, 0x73,
	0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x5f, 0x63, 0x61,
	0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x3a, 0x01, 0x2a, 0x22, 0x10, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x73, 0x12, 0x72,
	0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x43, 0x61,
	0x72, 0x64, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x5f,
	0x63, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x43,
	0x61, 0x72, 0x64, 0x73, 0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x72, 0x65, 0x64, 0x69,
	0x74, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x73,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x3a, 0x01, 0x2a,
	0x32, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x5f, 0x63, 0x61, 0x72,
	0x64, 0x73, 0x12, 0x81, 0x01, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x72, 0x65,
	0x64, 0x69, 0x74, 0x43, 0x61, 0x72, 0x64, 0x12, 0x2d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x72,
	0x65, 0x64, 0x69, 0x74, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x43, 0x61, 0x72, 0x64, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x72, 0x65,
	0x64, 0x69, 0x74, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x72,
	0x64, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x2a,
	0x15, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x5f, 0x63, 0x61, 0x72, 0x64,
	0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x88, 0x01, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x43, 0x72,
	0x65, 0x64, 0x69, 0x74, 0x43, 0x61, 0x72, 0x64, 0x12, 0x2a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x72, 0x65, 0x64, 0x69, 0x74, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x43, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x72, 0x65, 0x64, 0x69,
	0x74, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x72,
	0x65, 0x64, 0x69, 0x74, 0x43, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x21,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x65, 0x64,
	0x69, 0x74, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x73, 0x2f, 0x7b, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x7d, 0x12, 0x87, 0x01, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74,
	0x43, 0x61, 0x72, 0x64, 0x73, 0x12, 0x2b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x72, 0x65, 0x64,
	0x69, 0x74, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x43, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x29, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x5f,
	0x63, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x72, 0x65,
	0x64, 0x69, 0x74, 0x43, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1c, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x16, 0x12, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x65, 0x64, 0x69,
	0x74, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x73, 0x2f, 0x61, 0x6c, 0x6c, 0x42, 0x3c, 0x0a, 0x13, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x50, 0x01, 0x5a, 0x23, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x5f, 0x63, 0x61, 0x72,
	0x64, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x5f, 0x63, 0x61,
	0x72, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_credit_cards_v1_credit_cards_proto_rawDescOnce sync.Once
	file_credit_cards_v1_credit_cards_proto_rawDescData = file_credit_cards_v1_credit_cards_proto_rawDesc
)

func file_credit_cards_v1_credit_cards_proto_rawDescGZIP() []byte {
	file_credit_cards_v1_credit_cards_proto_rawDescOnce.Do(func() {
		file_credit_cards_v1_credit_cards_proto_rawDescData = protoimpl.X.CompressGZIP(file_credit_cards_v1_credit_cards_proto_rawDescData)
	})
	return file_credit_cards_v1_credit_cards_proto_rawDescData
}

var file_credit_cards_v1_credit_cards_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_credit_cards_v1_credit_cards_proto_goTypes = []any{
	(*CreditCards)(nil),              // 0: api.credit_cards.v1.CreditCards
	(*CardsReply)(nil),               // 1: api.credit_cards.v1.CardsReply
	(*DeleteCreditCardsRequest)(nil), // 2: api.credit_cards.v1.DeleteCreditCardsRequest
	(*GetCreditCardsRequest)(nil),    // 3: api.credit_cards.v1.GetCreditCardsRequest
	(*GetCreditCardsReply)(nil),      // 4: api.credit_cards.v1.GetCreditCardsReply
	(*SearchCreditCardsReply)(nil),   // 5: api.credit_cards.v1.SearchCreditCardsReply
	(*ListCreditCardsRequest)(nil),   // 6: api.credit_cards.v1.ListCreditCardsRequest
	(*ListCreditCardsReply)(nil),     // 7: api.credit_cards.v1.ListCreditCardsReply
}
var file_credit_cards_v1_credit_cards_proto_depIdxs = []int32{
	0, // 0: api.credit_cards.v1.GetCreditCardsReply.credit_cards:type_name -> api.credit_cards.v1.CreditCards
	0, // 1: api.credit_cards.v1.SearchCreditCardsReply.credit_cards:type_name -> api.credit_cards.v1.CreditCards
	0, // 2: api.credit_cards.v1.ListCreditCardsReply.credit_cards:type_name -> api.credit_cards.v1.CreditCards
	0, // 3: api.credit_cards.v1.CreditCardsService.CreateCreditCard:input_type -> api.credit_cards.v1.CreditCards
	0, // 4: api.credit_cards.v1.CreditCardsService.UpdateCreditCard:input_type -> api.credit_cards.v1.CreditCards
	2, // 5: api.credit_cards.v1.CreditCardsService.DeleteCreditCard:input_type -> api.credit_cards.v1.DeleteCreditCardsRequest
	3, // 6: api.credit_cards.v1.CreditCardsService.GetCreditCard:input_type -> api.credit_cards.v1.GetCreditCardsRequest
	6, // 7: api.credit_cards.v1.CreditCardsService.ListCreditCards:input_type -> api.credit_cards.v1.ListCreditCardsRequest
	1, // 8: api.credit_cards.v1.CreditCardsService.CreateCreditCard:output_type -> api.credit_cards.v1.CardsReply
	1, // 9: api.credit_cards.v1.CreditCardsService.UpdateCreditCard:output_type -> api.credit_cards.v1.CardsReply
	1, // 10: api.credit_cards.v1.CreditCardsService.DeleteCreditCard:output_type -> api.credit_cards.v1.CardsReply
	4, // 11: api.credit_cards.v1.CreditCardsService.GetCreditCard:output_type -> api.credit_cards.v1.GetCreditCardsReply
	7, // 12: api.credit_cards.v1.CreditCardsService.ListCreditCards:output_type -> api.credit_cards.v1.ListCreditCardsReply
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_credit_cards_v1_credit_cards_proto_init() }
func file_credit_cards_v1_credit_cards_proto_init() {
	if File_credit_cards_v1_credit_cards_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_credit_cards_v1_credit_cards_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_credit_cards_v1_credit_cards_proto_goTypes,
		DependencyIndexes: file_credit_cards_v1_credit_cards_proto_depIdxs,
		MessageInfos:      file_credit_cards_v1_credit_cards_proto_msgTypes,
	}.Build()
	File_credit_cards_v1_credit_cards_proto = out.File
	file_credit_cards_v1_credit_cards_proto_rawDesc = nil
	file_credit_cards_v1_credit_cards_proto_goTypes = nil
	file_credit_cards_v1_credit_cards_proto_depIdxs = nil
}
