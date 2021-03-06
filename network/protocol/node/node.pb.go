// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.15.6
// source: network/protocol/node/node.proto

package node

import (
	pcommon "github.com/overseven/try-network/network/protocol/pcommon"
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

type AddTransactionReply_Code int32

const (
	AddTransactionReply_TR_Ok    AddTransactionReply_Code = 0
	AddTransactionReply_TR_Error AddTransactionReply_Code = 1
)

// Enum value maps for AddTransactionReply_Code.
var (
	AddTransactionReply_Code_name = map[int32]string{
		0: "TR_Ok",
		1: "TR_Error",
	}
	AddTransactionReply_Code_value = map[string]int32{
		"TR_Ok":    0,
		"TR_Error": 1,
	}
)

func (x AddTransactionReply_Code) Enum() *AddTransactionReply_Code {
	p := new(AddTransactionReply_Code)
	*p = x
	return p
}

func (x AddTransactionReply_Code) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AddTransactionReply_Code) Descriptor() protoreflect.EnumDescriptor {
	return file_network_protocol_node_node_proto_enumTypes[0].Descriptor()
}

func (AddTransactionReply_Code) Type() protoreflect.EnumType {
	return &file_network_protocol_node_node_proto_enumTypes[0]
}

func (x AddTransactionReply_Code) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AddTransactionReply_Code.Descriptor instead.
func (AddTransactionReply_Code) EnumDescriptor() ([]byte, []int) {
	return file_network_protocol_node_node_proto_rawDescGZIP(), []int{5, 0}
}

type PushBlockReply_Code int32

const (
	PushBlockReply_PBR_Ok          PushBlockReply_Code = 0
	PushBlockReply_PBR_AlreadyHave PushBlockReply_Code = 1
	PushBlockReply_PBR_TooOld      PushBlockReply_Code = 2
	PushBlockReply_PBR_Incorrect   PushBlockReply_Code = 3
)

// Enum value maps for PushBlockReply_Code.
var (
	PushBlockReply_Code_name = map[int32]string{
		0: "PBR_Ok",
		1: "PBR_AlreadyHave",
		2: "PBR_TooOld",
		3: "PBR_Incorrect",
	}
	PushBlockReply_Code_value = map[string]int32{
		"PBR_Ok":          0,
		"PBR_AlreadyHave": 1,
		"PBR_TooOld":      2,
		"PBR_Incorrect":   3,
	}
)

func (x PushBlockReply_Code) Enum() *PushBlockReply_Code {
	p := new(PushBlockReply_Code)
	*p = x
	return p
}

func (x PushBlockReply_Code) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PushBlockReply_Code) Descriptor() protoreflect.EnumDescriptor {
	return file_network_protocol_node_node_proto_enumTypes[1].Descriptor()
}

func (PushBlockReply_Code) Type() protoreflect.EnumType {
	return &file_network_protocol_node_node_proto_enumTypes[1]
}

func (x PushBlockReply_Code) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PushBlockReply_Code.Descriptor instead.
func (PushBlockReply_Code) EnumDescriptor() ([]byte, []int) {
	return file_network_protocol_node_node_proto_rawDescGZIP(), []int{7, 0}
}

type ConnectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequesterAddress string `protobuf:"bytes,1,opt,name=requesterAddress,proto3" json:"requesterAddress,omitempty"`
}

func (x *ConnectRequest) Reset() {
	*x = ConnectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_protocol_node_node_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectRequest) ProtoMessage() {}

func (x *ConnectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_network_protocol_node_node_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectRequest.ProtoReflect.Descriptor instead.
func (*ConnectRequest) Descriptor() ([]byte, []int) {
	return file_network_protocol_node_node_proto_rawDescGZIP(), []int{0}
}

func (x *ConnectRequest) GetRequesterAddress() string {
	if x != nil {
		return x.RequesterAddress
	}
	return ""
}

type ConnectReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReplyerAddress []byte `protobuf:"bytes,1,opt,name=replyerAddress,proto3" json:"replyerAddress,omitempty"`
}

func (x *ConnectReply) Reset() {
	*x = ConnectReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_protocol_node_node_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectReply) ProtoMessage() {}

func (x *ConnectReply) ProtoReflect() protoreflect.Message {
	mi := &file_network_protocol_node_node_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectReply.ProtoReflect.Descriptor instead.
func (*ConnectReply) Descriptor() ([]byte, []int) {
	return file_network_protocol_node_node_proto_rawDescGZIP(), []int{1}
}

func (x *ConnectReply) GetReplyerAddress() []byte {
	if x != nil {
		return x.ReplyerAddress
	}
	return nil
}

type ListOfNodesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListOfNodesRequest) Reset() {
	*x = ListOfNodesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_protocol_node_node_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOfNodesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOfNodesRequest) ProtoMessage() {}

func (x *ListOfNodesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_network_protocol_node_node_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOfNodesRequest.ProtoReflect.Descriptor instead.
func (*ListOfNodesRequest) Descriptor() ([]byte, []int) {
	return file_network_protocol_node_node_proto_rawDescGZIP(), []int{2}
}

type ListOfNodesReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address []string `protobuf:"bytes,1,rep,name=address,proto3" json:"address,omitempty"`
}

func (x *ListOfNodesReply) Reset() {
	*x = ListOfNodesReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_protocol_node_node_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOfNodesReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOfNodesReply) ProtoMessage() {}

func (x *ListOfNodesReply) ProtoReflect() protoreflect.Message {
	mi := &file_network_protocol_node_node_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOfNodesReply.ProtoReflect.Descriptor instead.
func (*ListOfNodesReply) Descriptor() ([]byte, []int) {
	return file_network_protocol_node_node_proto_rawDescGZIP(), []int{3}
}

func (x *ListOfNodesReply) GetAddress() []string {
	if x != nil {
		return x.Address
	}
	return nil
}

type AddTransactionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction *pcommon.Transaction `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
}

func (x *AddTransactionRequest) Reset() {
	*x = AddTransactionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_protocol_node_node_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddTransactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTransactionRequest) ProtoMessage() {}

func (x *AddTransactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_network_protocol_node_node_proto_msgTypes[4]
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
	return file_network_protocol_node_node_proto_rawDescGZIP(), []int{4}
}

func (x *AddTransactionRequest) GetTransaction() *pcommon.Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

type AddTransactionReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reply      AddTransactionReply_Code `protobuf:"varint,1,opt,name=reply,proto3,enum=pnode.AddTransactionReply_Code" json:"reply,omitempty"`
	Message    string                   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Additional string                   `protobuf:"bytes,3,opt,name=additional,proto3" json:"additional,omitempty"`
}

func (x *AddTransactionReply) Reset() {
	*x = AddTransactionReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_protocol_node_node_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddTransactionReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTransactionReply) ProtoMessage() {}

func (x *AddTransactionReply) ProtoReflect() protoreflect.Message {
	mi := &file_network_protocol_node_node_proto_msgTypes[5]
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
	return file_network_protocol_node_node_proto_rawDescGZIP(), []int{5}
}

func (x *AddTransactionReply) GetReply() AddTransactionReply_Code {
	if x != nil {
		return x.Reply
	}
	return AddTransactionReply_TR_Ok
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

type PushBlockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NewBlock *pcommon.Block `protobuf:"bytes,1,opt,name=newBlock,proto3" json:"newBlock,omitempty"`
}

func (x *PushBlockRequest) Reset() {
	*x = PushBlockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_protocol_node_node_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushBlockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushBlockRequest) ProtoMessage() {}

func (x *PushBlockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_network_protocol_node_node_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushBlockRequest.ProtoReflect.Descriptor instead.
func (*PushBlockRequest) Descriptor() ([]byte, []int) {
	return file_network_protocol_node_node_proto_rawDescGZIP(), []int{6}
}

func (x *PushBlockRequest) GetNewBlock() *pcommon.Block {
	if x != nil {
		return x.NewBlock
	}
	return nil
}

type PushBlockReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reply PushBlockReply_Code `protobuf:"varint,1,opt,name=reply,proto3,enum=pnode.PushBlockReply_Code" json:"reply,omitempty"`
}

func (x *PushBlockReply) Reset() {
	*x = PushBlockReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_protocol_node_node_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushBlockReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushBlockReply) ProtoMessage() {}

func (x *PushBlockReply) ProtoReflect() protoreflect.Message {
	mi := &file_network_protocol_node_node_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushBlockReply.ProtoReflect.Descriptor instead.
func (*PushBlockReply) Descriptor() ([]byte, []int) {
	return file_network_protocol_node_node_proto_rawDescGZIP(), []int{7}
}

func (x *PushBlockReply) GetReply() PushBlockReply_Code {
	if x != nil {
		return x.Reply
	}
	return PushBlockReply_PBR_Ok
}

type BlocksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockIdBegin uint64 `protobuf:"varint,1,opt,name=blockIdBegin,proto3" json:"blockIdBegin,omitempty"`
	BlockIdEnd   uint64 `protobuf:"varint,2,opt,name=blockIdEnd,proto3" json:"blockIdEnd,omitempty"`
}

func (x *BlocksRequest) Reset() {
	*x = BlocksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_protocol_node_node_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlocksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlocksRequest) ProtoMessage() {}

func (x *BlocksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_network_protocol_node_node_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlocksRequest.ProtoReflect.Descriptor instead.
func (*BlocksRequest) Descriptor() ([]byte, []int) {
	return file_network_protocol_node_node_proto_rawDescGZIP(), []int{8}
}

func (x *BlocksRequest) GetBlockIdBegin() uint64 {
	if x != nil {
		return x.BlockIdBegin
	}
	return 0
}

func (x *BlocksRequest) GetBlockIdEnd() uint64 {
	if x != nil {
		return x.BlockIdEnd
	}
	return 0
}

type BlocksReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blocks []*pcommon.Block `protobuf:"bytes,1,rep,name=blocks,proto3" json:"blocks,omitempty"`
}

func (x *BlocksReply) Reset() {
	*x = BlocksReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_protocol_node_node_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlocksReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlocksReply) ProtoMessage() {}

func (x *BlocksReply) ProtoReflect() protoreflect.Message {
	mi := &file_network_protocol_node_node_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlocksReply.ProtoReflect.Descriptor instead.
func (*BlocksReply) Descriptor() ([]byte, []int) {
	return file_network_protocol_node_node_proto_rawDescGZIP(), []int{9}
}

func (x *BlocksReply) GetBlocks() []*pcommon.Block {
	if x != nil {
		return x.Blocks
	}
	return nil
}

type WalletBalanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PubKey []byte `protobuf:"bytes,1,opt,name=pubKey,proto3" json:"pubKey,omitempty"`
}

func (x *WalletBalanceRequest) Reset() {
	*x = WalletBalanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_protocol_node_node_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WalletBalanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WalletBalanceRequest) ProtoMessage() {}

func (x *WalletBalanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_network_protocol_node_node_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WalletBalanceRequest.ProtoReflect.Descriptor instead.
func (*WalletBalanceRequest) Descriptor() ([]byte, []int) {
	return file_network_protocol_node_node_proto_rawDescGZIP(), []int{10}
}

func (x *WalletBalanceRequest) GetPubKey() []byte {
	if x != nil {
		return x.PubKey
	}
	return nil
}

type WalletBalanceReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PubKey []byte  `protobuf:"bytes,1,opt,name=pubKey,proto3" json:"pubKey,omitempty"`
	Value  float64 `protobuf:"fixed64,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *WalletBalanceReply) Reset() {
	*x = WalletBalanceReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_protocol_node_node_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WalletBalanceReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WalletBalanceReply) ProtoMessage() {}

func (x *WalletBalanceReply) ProtoReflect() protoreflect.Message {
	mi := &file_network_protocol_node_node_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WalletBalanceReply.ProtoReflect.Descriptor instead.
func (*WalletBalanceReply) Descriptor() ([]byte, []int) {
	return file_network_protocol_node_node_proto_rawDescGZIP(), []int{11}
}

func (x *WalletBalanceReply) GetPubKey() []byte {
	if x != nil {
		return x.PubKey
	}
	return nil
}

func (x *WalletBalanceReply) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

var File_network_protocol_node_node_proto protoreflect.FileDescriptor

var file_network_protocol_node_node_proto_rawDesc = []byte{
	0x0a, 0x20, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x6e, 0x6f, 0x64, 0x65, 0x1a, 0x26, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x70, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x3c, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x72,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22,
	0x36, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x26, 0x0a, 0x0e, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x65, 0x72,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x14, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x4f,
	0x66, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2c, 0x0a,
	0x10, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x66, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x4f, 0x0a, 0x15, 0x41,
	0x64, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xa7, 0x01, 0x0a,
	0x13, 0x41, 0x64, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x35, 0x0a, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x70, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e,
	0x43, 0x6f, 0x64, 0x65, 0x52, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x64, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x22, 0x1f, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x09, 0x0a,
	0x05, 0x54, 0x52, 0x5f, 0x4f, 0x6b, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x54, 0x52, 0x5f, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x10, 0x01, 0x22, 0x3e, 0x0a, 0x10, 0x50, 0x75, 0x73, 0x68, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x08, 0x6e, 0x65,
	0x77, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x08, 0x6e, 0x65,
	0x77, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x8e, 0x01, 0x0a, 0x0e, 0x50, 0x75, 0x73, 0x68, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x30, 0x0a, 0x05, 0x72, 0x65, 0x70,
	0x6c, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x70, 0x6e, 0x6f, 0x64, 0x65,
	0x2e, 0x50, 0x75, 0x73, 0x68, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e,
	0x43, 0x6f, 0x64, 0x65, 0x52, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x4a, 0x0a, 0x04, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x42, 0x52, 0x5f, 0x4f, 0x6b, 0x10, 0x00, 0x12,
	0x13, 0x0a, 0x0f, 0x50, 0x42, 0x52, 0x5f, 0x41, 0x6c, 0x72, 0x65, 0x61, 0x64, 0x79, 0x48, 0x61,
	0x76, 0x65, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x42, 0x52, 0x5f, 0x54, 0x6f, 0x6f, 0x4f,
	0x6c, 0x64, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x50, 0x42, 0x52, 0x5f, 0x49, 0x6e, 0x63, 0x6f,
	0x72, 0x72, 0x65, 0x63, 0x74, 0x10, 0x03, 0x22, 0x53, 0x0a, 0x0d, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x49, 0x64, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x12, 0x1e, 0x0a, 0x0a,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x45, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x45, 0x6e, 0x64, 0x22, 0x35, 0x0a, 0x0b,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x26, 0x0a, 0x06, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x06, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x73, 0x22, 0x2e, 0x0a, 0x14, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x42, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70,
	0x75, 0x62, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x70, 0x75, 0x62,
	0x4b, 0x65, 0x79, 0x22, 0x42, 0x0a, 0x12, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x42, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x75, 0x62,
	0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x70, 0x75, 0x62, 0x4b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32, 0x9c, 0x03, 0x0a, 0x05, 0x4e, 0x6f, 0x64, 0x65,
	0x72, 0x12, 0x37, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x15, 0x2e, 0x70,
	0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x66, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x19, 0x2e, 0x70,
	0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x66, 0x4e, 0x6f, 0x64, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x6e, 0x6f, 0x64, 0x65, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x4f, 0x66, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x12, 0x4c, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x2e, 0x70, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x41, 0x64, 0x64,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00,
	0x12, 0x3d, 0x0a, 0x09, 0x50, 0x75, 0x73, 0x68, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x17, 0x2e,
	0x70, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x50,
	0x75, 0x73, 0x68, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12,
	0x37, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x14, 0x2e, 0x70,
	0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x57,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1b, 0x2e, 0x70,
	0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x6e, 0x6f, 0x64,
	0x65, 0x2e, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x76, 0x65, 0x72, 0x73, 0x65, 0x76, 0x65, 0x6e, 0x2f, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_network_protocol_node_node_proto_rawDescOnce sync.Once
	file_network_protocol_node_node_proto_rawDescData = file_network_protocol_node_node_proto_rawDesc
)

func file_network_protocol_node_node_proto_rawDescGZIP() []byte {
	file_network_protocol_node_node_proto_rawDescOnce.Do(func() {
		file_network_protocol_node_node_proto_rawDescData = protoimpl.X.CompressGZIP(file_network_protocol_node_node_proto_rawDescData)
	})
	return file_network_protocol_node_node_proto_rawDescData
}

var file_network_protocol_node_node_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_network_protocol_node_node_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_network_protocol_node_node_proto_goTypes = []interface{}{
	(AddTransactionReply_Code)(0), // 0: pnode.AddTransactionReply.Code
	(PushBlockReply_Code)(0),      // 1: pnode.PushBlockReply.Code
	(*ConnectRequest)(nil),        // 2: pnode.ConnectRequest
	(*ConnectReply)(nil),          // 3: pnode.ConnectReply
	(*ListOfNodesRequest)(nil),    // 4: pnode.ListOfNodesRequest
	(*ListOfNodesReply)(nil),      // 5: pnode.ListOfNodesReply
	(*AddTransactionRequest)(nil), // 6: pnode.AddTransactionRequest
	(*AddTransactionReply)(nil),   // 7: pnode.AddTransactionReply
	(*PushBlockRequest)(nil),      // 8: pnode.PushBlockRequest
	(*PushBlockReply)(nil),        // 9: pnode.PushBlockReply
	(*BlocksRequest)(nil),         // 10: pnode.BlocksRequest
	(*BlocksReply)(nil),           // 11: pnode.BlocksReply
	(*WalletBalanceRequest)(nil),  // 12: pnode.WalletBalanceRequest
	(*WalletBalanceReply)(nil),    // 13: pnode.WalletBalanceReply
	(*pcommon.Transaction)(nil),   // 14: pcommon.Transaction
	(*pcommon.Block)(nil),         // 15: pcommon.Block
}
var file_network_protocol_node_node_proto_depIdxs = []int32{
	14, // 0: pnode.AddTransactionRequest.transaction:type_name -> pcommon.Transaction
	0,  // 1: pnode.AddTransactionReply.reply:type_name -> pnode.AddTransactionReply.Code
	15, // 2: pnode.PushBlockRequest.newBlock:type_name -> pcommon.Block
	1,  // 3: pnode.PushBlockReply.reply:type_name -> pnode.PushBlockReply.Code
	15, // 4: pnode.BlocksReply.blocks:type_name -> pcommon.Block
	2,  // 5: pnode.Noder.Connect:input_type -> pnode.ConnectRequest
	4,  // 6: pnode.Noder.GetListOfNodes:input_type -> pnode.ListOfNodesRequest
	6,  // 7: pnode.Noder.AddTransaction:input_type -> pnode.AddTransactionRequest
	8,  // 8: pnode.Noder.PushBlock:input_type -> pnode.PushBlockRequest
	10, // 9: pnode.Noder.GetBlocks:input_type -> pnode.BlocksRequest
	12, // 10: pnode.Noder.GetWalletBalance:input_type -> pnode.WalletBalanceRequest
	3,  // 11: pnode.Noder.Connect:output_type -> pnode.ConnectReply
	5,  // 12: pnode.Noder.GetListOfNodes:output_type -> pnode.ListOfNodesReply
	7,  // 13: pnode.Noder.AddTransaction:output_type -> pnode.AddTransactionReply
	9,  // 14: pnode.Noder.PushBlock:output_type -> pnode.PushBlockReply
	11, // 15: pnode.Noder.GetBlocks:output_type -> pnode.BlocksReply
	13, // 16: pnode.Noder.GetWalletBalance:output_type -> pnode.WalletBalanceReply
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_network_protocol_node_node_proto_init() }
func file_network_protocol_node_node_proto_init() {
	if File_network_protocol_node_node_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_network_protocol_node_node_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectRequest); i {
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
		file_network_protocol_node_node_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectReply); i {
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
		file_network_protocol_node_node_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListOfNodesRequest); i {
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
		file_network_protocol_node_node_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListOfNodesReply); i {
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
		file_network_protocol_node_node_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_network_protocol_node_node_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_network_protocol_node_node_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushBlockRequest); i {
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
		file_network_protocol_node_node_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushBlockReply); i {
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
		file_network_protocol_node_node_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlocksRequest); i {
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
		file_network_protocol_node_node_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlocksReply); i {
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
		file_network_protocol_node_node_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WalletBalanceRequest); i {
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
		file_network_protocol_node_node_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WalletBalanceReply); i {
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
			RawDescriptor: file_network_protocol_node_node_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_network_protocol_node_node_proto_goTypes,
		DependencyIndexes: file_network_protocol_node_node_proto_depIdxs,
		EnumInfos:         file_network_protocol_node_node_proto_enumTypes,
		MessageInfos:      file_network_protocol_node_node_proto_msgTypes,
	}.Build()
	File_network_protocol_node_node_proto = out.File
	file_network_protocol_node_node_proto_rawDesc = nil
	file_network_protocol_node_node_proto_goTypes = nil
	file_network_protocol_node_node_proto_depIdxs = nil
}
