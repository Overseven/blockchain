// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.15.6
// source: description.proto

package protocol

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Transaction_Type int32

const (
	Transaction_AirDrop  Transaction_Type = 0
	Transaction_Transfer Transaction_Type = 1
)

// Enum value maps for Transaction_Type.
var (
	Transaction_Type_name = map[int32]string{
		0: "AirDrop",
		1: "Transfer",
	}
	Transaction_Type_value = map[string]int32{
		"AirDrop":  0,
		"Transfer": 1,
	}
)

func (x Transaction_Type) Enum() *Transaction_Type {
	p := new(Transaction_Type)
	*p = x
	return p
}

func (x Transaction_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Transaction_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_description_proto_enumTypes[0].Descriptor()
}

func (Transaction_Type) Type() protoreflect.EnumType {
	return &file_description_proto_enumTypes[0]
}

func (x Transaction_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Transaction_Type.Descriptor instead.
func (Transaction_Type) EnumDescriptor() ([]byte, []int) {
	return file_description_proto_rawDescGZIP(), []int{0, 0}
}

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProtocolVersion uint32                 `protobuf:"varint,1,opt,name=protocolVersion,proto3" json:"protocolVersion,omitempty"`
	Type            Transaction_Type       `protobuf:"varint,2,opt,name=type,proto3,enum=protocol.Transaction_Type" json:"type,omitempty"`
	Sender          []byte                 `protobuf:"bytes,3,opt,name=sender,proto3" json:"sender,omitempty"`
	Receiver        []byte                 `protobuf:"bytes,4,opt,name=receiver,proto3" json:"receiver,omitempty"`
	Message         string                 `protobuf:"bytes,5,opt,name=message,proto3" json:"message,omitempty"`
	Timestamp       *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Pay             float64                `protobuf:"fixed64,7,opt,name=pay,proto3" json:"pay,omitempty"`
	Fee             float64                `protobuf:"fixed64,8,opt,name=fee,proto3" json:"fee,omitempty"`
	SenderSign      []byte                 `protobuf:"bytes,9,opt,name=senderSign,proto3" json:"senderSign,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_description_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_description_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_description_proto_rawDescGZIP(), []int{0}
}

func (x *Transaction) GetProtocolVersion() uint32 {
	if x != nil {
		return x.ProtocolVersion
	}
	return 0
}

func (x *Transaction) GetType() Transaction_Type {
	if x != nil {
		return x.Type
	}
	return Transaction_AirDrop
}

func (x *Transaction) GetSender() []byte {
	if x != nil {
		return x.Sender
	}
	return nil
}

func (x *Transaction) GetReceiver() []byte {
	if x != nil {
		return x.Receiver
	}
	return nil
}

func (x *Transaction) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Transaction) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *Transaction) GetPay() float64 {
	if x != nil {
		return x.Pay
	}
	return 0
}

func (x *Transaction) GetFee() float64 {
	if x != nil {
		return x.Fee
	}
	return 0
}

func (x *Transaction) GetSenderSign() []byte {
	if x != nil {
		return x.SenderSign
	}
	return nil
}

type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProtocolVersion uint32         `protobuf:"varint,1,opt,name=protocolVersion,proto3" json:"protocolVersion,omitempty"`
	BlockId         uint64         `protobuf:"varint,2,opt,name=blockId,proto3" json:"blockId,omitempty"`
	Trans           []*Transaction `protobuf:"bytes,3,rep,name=trans,proto3" json:"trans,omitempty"`
	PrevBlockHash   []byte         `protobuf:"bytes,4,opt,name=prevBlockHash,proto3" json:"prevBlockHash,omitempty"`
	Difficulty      uint32         `protobuf:"varint,5,opt,name=difficulty,proto3" json:"difficulty,omitempty"`
	Miner           []byte         `protobuf:"bytes,6,opt,name=miner,proto3" json:"miner,omitempty"`
	BlockHash       []byte         `protobuf:"bytes,7,opt,name=blockHash,proto3" json:"blockHash,omitempty"`
	Nonce           []byte         `protobuf:"bytes,8,opt,name=nonce,proto3" json:"nonce,omitempty"`
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_description_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_description_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_description_proto_rawDescGZIP(), []int{1}
}

func (x *Block) GetProtocolVersion() uint32 {
	if x != nil {
		return x.ProtocolVersion
	}
	return 0
}

func (x *Block) GetBlockId() uint64 {
	if x != nil {
		return x.BlockId
	}
	return 0
}

func (x *Block) GetTrans() []*Transaction {
	if x != nil {
		return x.Trans
	}
	return nil
}

func (x *Block) GetPrevBlockHash() []byte {
	if x != nil {
		return x.PrevBlockHash
	}
	return nil
}

func (x *Block) GetDifficulty() uint32 {
	if x != nil {
		return x.Difficulty
	}
	return 0
}

func (x *Block) GetMiner() []byte {
	if x != nil {
		return x.Miner
	}
	return nil
}

func (x *Block) GetBlockHash() []byte {
	if x != nil {
		return x.BlockHash
	}
	return nil
}

func (x *Block) GetNonce() []byte {
	if x != nil {
		return x.Nonce
	}
	return nil
}

type AddTransactionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age  int32  `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
}

func (x *AddTransactionRequest) Reset() {
	*x = AddTransactionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_description_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddTransactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTransactionRequest) ProtoMessage() {}

func (x *AddTransactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_description_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddTransactionRequest.ProtoReflect.Descriptor instead.
func (*AddTransactionRequest) Descriptor() ([]byte, []int) {
	return file_description_proto_rawDescGZIP(), []int{2}
}

func (x *AddTransactionRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddTransactionRequest) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

type AddTransactionReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message    string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Additional string `protobuf:"bytes,2,opt,name=additional,proto3" json:"additional,omitempty"`
}

func (x *AddTransactionReply) Reset() {
	*x = AddTransactionReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_description_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddTransactionReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTransactionReply) ProtoMessage() {}

func (x *AddTransactionReply) ProtoReflect() protoreflect.Message {
	mi := &file_description_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddTransactionReply.ProtoReflect.Descriptor instead.
func (*AddTransactionReply) Descriptor() ([]byte, []int) {
	return file_description_proto_rawDescGZIP(), []int{3}
}

func (x *AddTransactionReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *AddTransactionReply) GetAdditional() string {
	if x != nil {
		return x.Additional
	}
	return ""
}

type GetBlocksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockIdBegin uint64 `protobuf:"varint,1,opt,name=blockIdBegin,proto3" json:"blockIdBegin,omitempty"`
	BlockIdEnd   uint64 `protobuf:"varint,2,opt,name=blockIdEnd,proto3" json:"blockIdEnd,omitempty"`
}

func (x *GetBlocksRequest) Reset() {
	*x = GetBlocksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_description_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlocksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlocksRequest) ProtoMessage() {}

func (x *GetBlocksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_description_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlocksRequest.ProtoReflect.Descriptor instead.
func (*GetBlocksRequest) Descriptor() ([]byte, []int) {
	return file_description_proto_rawDescGZIP(), []int{4}
}

func (x *GetBlocksRequest) GetBlockIdBegin() uint64 {
	if x != nil {
		return x.BlockIdBegin
	}
	return 0
}

func (x *GetBlocksRequest) GetBlockIdEnd() uint64 {
	if x != nil {
		return x.BlockIdEnd
	}
	return 0
}

type GetBlocksReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blocks []*Block `protobuf:"bytes,1,rep,name=blocks,proto3" json:"blocks,omitempty"`
}

func (x *GetBlocksReply) Reset() {
	*x = GetBlocksReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_description_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlocksReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlocksReply) ProtoMessage() {}

func (x *GetBlocksReply) ProtoReflect() protoreflect.Message {
	mi := &file_description_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlocksReply.ProtoReflect.Descriptor instead.
func (*GetBlocksReply) Descriptor() ([]byte, []int) {
	return file_description_proto_rawDescGZIP(), []int{5}
}

func (x *GetBlocksReply) GetBlocks() []*Block {
	if x != nil {
		return x.Blocks
	}
	return nil
}

var File_description_proto protoreflect.FileDescriptor

var file_description_proto_rawDesc = []byte{
	0x0a, 0x11, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd6,
	0x02, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28,
	0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x12, 0x10, 0x0a, 0x03, 0x70, 0x61, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x70,
	0x61, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x65, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x03, 0x66, 0x65, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x53, 0x69,
	0x67, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x53, 0x69, 0x67, 0x6e, 0x22, 0x21, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07,
	0x41, 0x69, 0x72, 0x44, 0x72, 0x6f, 0x70, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x10, 0x01, 0x22, 0x88, 0x02, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x12, 0x28, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x05, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x72, 0x65, 0x76, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48,
	0x61, 0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x70, 0x72, 0x65, 0x76, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x69, 0x66, 0x66,
	0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x64, 0x69,
	0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x69, 0x6e, 0x65,
	0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x6d, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x1c,
	0x0a, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x14, 0x0a, 0x05,
	0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x6e, 0x6f, 0x6e,
	0x63, 0x65, 0x22, 0x3d, 0x0a, 0x15, 0x41, 0x64, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67,
	0x65, 0x22, 0x4f, 0x0a, 0x13, 0x41, 0x64, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x6c, 0x22, 0x56, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49,
	0x64, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x49, 0x64, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x49, 0x64, 0x45, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x45, 0x6e, 0x64, 0x22, 0x39, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x27, 0x0a, 0x06,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x06, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x32, 0xa0, 0x01, 0x0a, 0x05, 0x4e, 0x6f, 0x64, 0x65, 0x72, 0x12,
	0x52, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x41, 0x64, 0x64,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x41, 0x64,
	0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73,
	0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x76, 0x65, 0x72, 0x73, 0x65, 0x76, 0x65, 0x6e,
	0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_description_proto_rawDescOnce sync.Once
	file_description_proto_rawDescData = file_description_proto_rawDesc
)

func file_description_proto_rawDescGZIP() []byte {
	file_description_proto_rawDescOnce.Do(func() {
		file_description_proto_rawDescData = protoimpl.X.CompressGZIP(file_description_proto_rawDescData)
	})
	return file_description_proto_rawDescData
}

var file_description_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_description_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_description_proto_goTypes = []interface{}{
	(Transaction_Type)(0),         // 0: protocol.Transaction.Type
	(*Transaction)(nil),           // 1: protocol.Transaction
	(*Block)(nil),                 // 2: protocol.Block
	(*AddTransactionRequest)(nil), // 3: protocol.AddTransactionRequest
	(*AddTransactionReply)(nil),   // 4: protocol.AddTransactionReply
	(*GetBlocksRequest)(nil),      // 5: protocol.GetBlocksRequest
	(*GetBlocksReply)(nil),        // 6: protocol.GetBlocksReply
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_description_proto_depIdxs = []int32{
	0, // 0: protocol.Transaction.type:type_name -> protocol.Transaction.Type
	7, // 1: protocol.Transaction.timestamp:type_name -> google.protobuf.Timestamp
	1, // 2: protocol.Block.trans:type_name -> protocol.Transaction
	2, // 3: protocol.GetBlocksReply.blocks:type_name -> protocol.Block
	3, // 4: protocol.Noder.AddTransaction:input_type -> protocol.AddTransactionRequest
	5, // 5: protocol.Noder.GetBlocks:input_type -> protocol.GetBlocksRequest
	4, // 6: protocol.Noder.AddTransaction:output_type -> protocol.AddTransactionReply
	6, // 7: protocol.Noder.GetBlocks:output_type -> protocol.GetBlocksReply
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_description_proto_init() }
func file_description_proto_init() {
	if File_description_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_description_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
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
		file_description_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Block); i {
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
		file_description_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddTransactionRequest); i {
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
		file_description_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddTransactionReply); i {
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
		file_description_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlocksRequest); i {
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
		file_description_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlocksReply); i {
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
			RawDescriptor: file_description_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_description_proto_goTypes,
		DependencyIndexes: file_description_proto_depIdxs,
		EnumInfos:         file_description_proto_enumTypes,
		MessageInfos:      file_description_proto_msgTypes,
	}.Build()
	File_description_proto = out.File
	file_description_proto_rawDesc = nil
	file_description_proto_goTypes = nil
	file_description_proto_depIdxs = nil
}
