// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.15.6
// source: pcommon.proto

package proto

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

type TransAirDrop struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionCounter uint32                 `protobuf:"varint,1,opt,name=transactionCounter,proto3" json:"transactionCounter,omitempty"`
	Receiver           []byte                 `protobuf:"bytes,2,opt,name=receiver,proto3" json:"receiver,omitempty"`
	Pay                float64                `protobuf:"fixed64,3,opt,name=pay,proto3" json:"pay,omitempty"`
	Fee                float64                `protobuf:"fixed64,4,opt,name=fee,proto3" json:"fee,omitempty"`
	Message            string                 `protobuf:"bytes,5,opt,name=message,proto3" json:"message,omitempty"`
	Timestamp          *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Node               []byte                 `protobuf:"bytes,7,opt,name=node,proto3" json:"node,omitempty"`
	Sign               []byte                 `protobuf:"bytes,8,opt,name=sign,proto3" json:"sign,omitempty"`
}

func (x *TransAirDrop) Reset() {
	*x = TransAirDrop{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pcommon_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransAirDrop) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransAirDrop) ProtoMessage() {}

func (x *TransAirDrop) ProtoReflect() protoreflect.Message {
	mi := &file_pcommon_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransAirDrop.ProtoReflect.Descriptor instead.
func (*TransAirDrop) Descriptor() ([]byte, []int) {
	return file_pcommon_proto_rawDescGZIP(), []int{0}
}

func (x *TransAirDrop) GetTransactionCounter() uint32 {
	if x != nil {
		return x.TransactionCounter
	}
	return 0
}

func (x *TransAirDrop) GetReceiver() []byte {
	if x != nil {
		return x.Receiver
	}
	return nil
}

func (x *TransAirDrop) GetPay() float64 {
	if x != nil {
		return x.Pay
	}
	return 0
}

func (x *TransAirDrop) GetFee() float64 {
	if x != nil {
		return x.Fee
	}
	return 0
}

func (x *TransAirDrop) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *TransAirDrop) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *TransAirDrop) GetNode() []byte {
	if x != nil {
		return x.Node
	}
	return nil
}

func (x *TransAirDrop) GetSign() []byte {
	if x != nil {
		return x.Sign
	}
	return nil
}

type TransTransfer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sender             []byte                 `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	TransactionCounter uint32                 `protobuf:"varint,2,opt,name=transactionCounter,proto3" json:"transactionCounter,omitempty"`
	Receiver           []byte                 `protobuf:"bytes,3,opt,name=receiver,proto3" json:"receiver,omitempty"`
	Pay                float64                `protobuf:"fixed64,4,opt,name=pay,proto3" json:"pay,omitempty"`
	Fee                float64                `protobuf:"fixed64,5,opt,name=fee,proto3" json:"fee,omitempty"`
	Message            string                 `protobuf:"bytes,6,opt,name=message,proto3" json:"message,omitempty"`
	Timestamp          *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Node               []byte                 `protobuf:"bytes,8,opt,name=node,proto3" json:"node,omitempty"`
	Sign               []byte                 `protobuf:"bytes,9,opt,name=sign,proto3" json:"sign,omitempty"`
}

func (x *TransTransfer) Reset() {
	*x = TransTransfer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pcommon_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransTransfer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransTransfer) ProtoMessage() {}

func (x *TransTransfer) ProtoReflect() protoreflect.Message {
	mi := &file_pcommon_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransTransfer.ProtoReflect.Descriptor instead.
func (*TransTransfer) Descriptor() ([]byte, []int) {
	return file_pcommon_proto_rawDescGZIP(), []int{1}
}

func (x *TransTransfer) GetSender() []byte {
	if x != nil {
		return x.Sender
	}
	return nil
}

func (x *TransTransfer) GetTransactionCounter() uint32 {
	if x != nil {
		return x.TransactionCounter
	}
	return 0
}

func (x *TransTransfer) GetReceiver() []byte {
	if x != nil {
		return x.Receiver
	}
	return nil
}

func (x *TransTransfer) GetPay() float64 {
	if x != nil {
		return x.Pay
	}
	return 0
}

func (x *TransTransfer) GetFee() float64 {
	if x != nil {
		return x.Fee
	}
	return 0
}

func (x *TransTransfer) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *TransTransfer) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *TransTransfer) GetNode() []byte {
	if x != nil {
		return x.Node
	}
	return nil
}

func (x *TransTransfer) GetSign() []byte {
	if x != nil {
		return x.Sign
	}
	return nil
}

type TransVotingInit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sender             []byte                 `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	TransactionCounter uint32                 `protobuf:"varint,2,opt,name=transactionCounter,proto3" json:"transactionCounter,omitempty"`
	VotingId           uint64                 `protobuf:"varint,3,opt,name=votingId,proto3" json:"votingId,omitempty"`
	Parameter          uint32                 `protobuf:"varint,4,opt,name=parameter,proto3" json:"parameter,omitempty"`
	Value              string                 `protobuf:"bytes,5,opt,name=value,proto3" json:"value,omitempty"`
	Fee                float64                `protobuf:"fixed64,6,opt,name=fee,proto3" json:"fee,omitempty"`
	Timestamp          *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Node               []byte                 `protobuf:"bytes,8,opt,name=node,proto3" json:"node,omitempty"`
	Sign               []byte                 `protobuf:"bytes,9,opt,name=sign,proto3" json:"sign,omitempty"`
}

func (x *TransVotingInit) Reset() {
	*x = TransVotingInit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pcommon_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransVotingInit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransVotingInit) ProtoMessage() {}

func (x *TransVotingInit) ProtoReflect() protoreflect.Message {
	mi := &file_pcommon_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransVotingInit.ProtoReflect.Descriptor instead.
func (*TransVotingInit) Descriptor() ([]byte, []int) {
	return file_pcommon_proto_rawDescGZIP(), []int{2}
}

func (x *TransVotingInit) GetSender() []byte {
	if x != nil {
		return x.Sender
	}
	return nil
}

func (x *TransVotingInit) GetTransactionCounter() uint32 {
	if x != nil {
		return x.TransactionCounter
	}
	return 0
}

func (x *TransVotingInit) GetVotingId() uint64 {
	if x != nil {
		return x.VotingId
	}
	return 0
}

func (x *TransVotingInit) GetParameter() uint32 {
	if x != nil {
		return x.Parameter
	}
	return 0
}

func (x *TransVotingInit) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *TransVotingInit) GetFee() float64 {
	if x != nil {
		return x.Fee
	}
	return 0
}

func (x *TransVotingInit) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *TransVotingInit) GetNode() []byte {
	if x != nil {
		return x.Node
	}
	return nil
}

func (x *TransVotingInit) GetSign() []byte {
	if x != nil {
		return x.Sign
	}
	return nil
}

type TransVote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sender             []byte                 `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	TransactionCounter uint32                 `protobuf:"varint,2,opt,name=transactionCounter,proto3" json:"transactionCounter,omitempty"`
	VotingId           uint64                 `protobuf:"varint,3,opt,name=votingId,proto3" json:"votingId,omitempty"`
	Opinion            string                 `protobuf:"bytes,4,opt,name=opinion,proto3" json:"opinion,omitempty"`
	Fee                float64                `protobuf:"fixed64,5,opt,name=fee,proto3" json:"fee,omitempty"`
	Timestamp          *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Node               []byte                 `protobuf:"bytes,7,opt,name=node,proto3" json:"node,omitempty"`
	Sign               []byte                 `protobuf:"bytes,8,opt,name=sign,proto3" json:"sign,omitempty"`
}

func (x *TransVote) Reset() {
	*x = TransVote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pcommon_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransVote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransVote) ProtoMessage() {}

func (x *TransVote) ProtoReflect() protoreflect.Message {
	mi := &file_pcommon_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransVote.ProtoReflect.Descriptor instead.
func (*TransVote) Descriptor() ([]byte, []int) {
	return file_pcommon_proto_rawDescGZIP(), []int{3}
}

func (x *TransVote) GetSender() []byte {
	if x != nil {
		return x.Sender
	}
	return nil
}

func (x *TransVote) GetTransactionCounter() uint32 {
	if x != nil {
		return x.TransactionCounter
	}
	return 0
}

func (x *TransVote) GetVotingId() uint64 {
	if x != nil {
		return x.VotingId
	}
	return 0
}

func (x *TransVote) GetOpinion() string {
	if x != nil {
		return x.Opinion
	}
	return ""
}

func (x *TransVote) GetFee() float64 {
	if x != nil {
		return x.Fee
	}
	return 0
}

func (x *TransVote) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *TransVote) GetNode() []byte {
	if x != nil {
		return x.Node
	}
	return nil
}

func (x *TransVote) GetSign() []byte {
	if x != nil {
		return x.Sign
	}
	return nil
}

type TransVotingFinish struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sender             []byte                 `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	TransactionCounter uint32                 `protobuf:"varint,2,opt,name=transactionCounter,proto3" json:"transactionCounter,omitempty"`
	VotingId           uint64                 `protobuf:"varint,3,opt,name=votingId,proto3" json:"votingId,omitempty"`
	Fee                float64                `protobuf:"fixed64,4,opt,name=fee,proto3" json:"fee,omitempty"`
	Timestamp          *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Node               []byte                 `protobuf:"bytes,6,opt,name=node,proto3" json:"node,omitempty"`
	Sign               []byte                 `protobuf:"bytes,7,opt,name=sign,proto3" json:"sign,omitempty"`
}

func (x *TransVotingFinish) Reset() {
	*x = TransVotingFinish{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pcommon_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransVotingFinish) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransVotingFinish) ProtoMessage() {}

func (x *TransVotingFinish) ProtoReflect() protoreflect.Message {
	mi := &file_pcommon_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransVotingFinish.ProtoReflect.Descriptor instead.
func (*TransVotingFinish) Descriptor() ([]byte, []int) {
	return file_pcommon_proto_rawDescGZIP(), []int{4}
}

func (x *TransVotingFinish) GetSender() []byte {
	if x != nil {
		return x.Sender
	}
	return nil
}

func (x *TransVotingFinish) GetTransactionCounter() uint32 {
	if x != nil {
		return x.TransactionCounter
	}
	return 0
}

func (x *TransVotingFinish) GetVotingId() uint64 {
	if x != nil {
		return x.VotingId
	}
	return 0
}

func (x *TransVotingFinish) GetFee() float64 {
	if x != nil {
		return x.Fee
	}
	return 0
}

func (x *TransVotingFinish) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *TransVotingFinish) GetNode() []byte {
	if x != nil {
		return x.Node
	}
	return nil
}

func (x *TransVotingFinish) GetSign() []byte {
	if x != nil {
		return x.Sign
	}
	return nil
}

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProtocolVersion uint32 `protobuf:"varint,1,opt,name=protocolVersion,proto3" json:"protocolVersion,omitempty"`
	// Types that are assignable to Trans:
	//	*Transaction_Transfer
	//	*Transaction_Airdrop
	//	*Transaction_VotingInit
	//	*Transaction_Vote
	//	*Transaction_VotingFinish
	Trans isTransaction_Trans `protobuf_oneof:"trans"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pcommon_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_pcommon_proto_msgTypes[5]
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
	return file_pcommon_proto_rawDescGZIP(), []int{5}
}

func (x *Transaction) GetProtocolVersion() uint32 {
	if x != nil {
		return x.ProtocolVersion
	}
	return 0
}

func (m *Transaction) GetTrans() isTransaction_Trans {
	if m != nil {
		return m.Trans
	}
	return nil
}

func (x *Transaction) GetTransfer() *TransTransfer {
	if x, ok := x.GetTrans().(*Transaction_Transfer); ok {
		return x.Transfer
	}
	return nil
}

func (x *Transaction) GetAirdrop() *TransAirDrop {
	if x, ok := x.GetTrans().(*Transaction_Airdrop); ok {
		return x.Airdrop
	}
	return nil
}

func (x *Transaction) GetVotingInit() *TransVotingInit {
	if x, ok := x.GetTrans().(*Transaction_VotingInit); ok {
		return x.VotingInit
	}
	return nil
}

func (x *Transaction) GetVote() *TransVote {
	if x, ok := x.GetTrans().(*Transaction_Vote); ok {
		return x.Vote
	}
	return nil
}

func (x *Transaction) GetVotingFinish() *TransVotingFinish {
	if x, ok := x.GetTrans().(*Transaction_VotingFinish); ok {
		return x.VotingFinish
	}
	return nil
}

type isTransaction_Trans interface {
	isTransaction_Trans()
}

type Transaction_Transfer struct {
	Transfer *TransTransfer `protobuf:"bytes,3,opt,name=transfer,proto3,oneof"`
}

type Transaction_Airdrop struct {
	Airdrop *TransAirDrop `protobuf:"bytes,4,opt,name=airdrop,proto3,oneof"`
}

type Transaction_VotingInit struct {
	VotingInit *TransVotingInit `protobuf:"bytes,5,opt,name=votingInit,proto3,oneof"`
}

type Transaction_Vote struct {
	Vote *TransVote `protobuf:"bytes,6,opt,name=vote,proto3,oneof"`
}

type Transaction_VotingFinish struct {
	VotingFinish *TransVotingFinish `protobuf:"bytes,7,opt,name=votingFinish,proto3,oneof"`
}

func (*Transaction_Transfer) isTransaction_Trans() {}

func (*Transaction_Airdrop) isTransaction_Trans() {}

func (*Transaction_VotingInit) isTransaction_Trans() {}

func (*Transaction_Vote) isTransaction_Trans() {}

func (*Transaction_VotingFinish) isTransaction_Trans() {}

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
		mi := &file_pcommon_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_pcommon_proto_msgTypes[6]
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
	return file_pcommon_proto_rawDescGZIP(), []int{6}
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

var File_pcommon_proto protoreflect.FileDescriptor

var file_pcommon_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x70, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfa, 0x01, 0x0a, 0x0c, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x41, 0x69, 0x72, 0x44, 0x72, 0x6f, 0x70, 0x12, 0x2e, 0x0a, 0x12, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x72, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x61, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x03, 0x70, 0x61, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x65, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x66, 0x65, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x6e, 0x6f,
	0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x67, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x73, 0x69, 0x67, 0x6e, 0x22, 0x93, 0x02, 0x0a, 0x0d, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x12, 0x2e, 0x0a, 0x12, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72,
	0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03,
	0x70, 0x61, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x70, 0x61, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x66, 0x65, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x66, 0x65, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x67, 0x6e,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x73, 0x69, 0x67, 0x6e, 0x22, 0x9d, 0x02, 0x0a,
	0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x56, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x69, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x12, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x6f, 0x74, 0x69,
	0x6e, 0x67, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x76, 0x6f, 0x74, 0x69,
	0x6e, 0x67, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74,
	0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x65, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x66, 0x65, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x67, 0x6e,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x73, 0x69, 0x67, 0x6e, 0x22, 0xfd, 0x01, 0x0a,
	0x09, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x12, 0x2e, 0x0a, 0x12, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x6f, 0x70, 0x69, 0x6e, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6f, 0x70, 0x69, 0x6e, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x65, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x66, 0x65, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x67, 0x6e,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x73, 0x69, 0x67, 0x6e, 0x22, 0xeb, 0x01, 0x0a,
	0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x56, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x6e, 0x69,
	0x73, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x12, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x6f,
	0x74, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x76, 0x6f,
	0x74, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x65, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x03, 0x66, 0x65, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x67, 0x6e, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x73, 0x69, 0x67, 0x6e, 0x22, 0xd1, 0x02, 0x0a, 0x0b, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x0f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x48, 0x00,
	0x52, 0x08, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x31, 0x0a, 0x07, 0x61, 0x69,
	0x72, 0x64, 0x72, 0x6f, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x41, 0x69, 0x72, 0x44, 0x72,
	0x6f, 0x70, 0x48, 0x00, 0x52, 0x07, 0x61, 0x69, 0x72, 0x64, 0x72, 0x6f, 0x70, 0x12, 0x3a, 0x0a,
	0x0a, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x70, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x56, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x69, 0x74, 0x48, 0x00, 0x52, 0x0a, 0x76,
	0x6f, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x69, 0x74, 0x12, 0x28, 0x0a, 0x04, 0x76, 0x6f, 0x74,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x56, 0x6f, 0x74, 0x65, 0x48, 0x00, 0x52, 0x04, 0x76,
	0x6f, 0x74, 0x65, 0x12, 0x40, 0x0a, 0x0c, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x6e,
	0x69, 0x73, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x56, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x46,
	0x69, 0x6e, 0x69, 0x73, 0x68, 0x48, 0x00, 0x52, 0x0c, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x46,
	0x69, 0x6e, 0x69, 0x73, 0x68, 0x42, 0x07, 0x0a, 0x05, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x22, 0x87,
	0x02, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x28, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x05,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x05, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x72, 0x65, 0x76,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0d, 0x70, 0x72, 0x65, 0x76, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1e,
	0x0a, 0x0a, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0a, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x6d, 0x69, 0x6e, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x6d,
	0x69, 0x6e, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73,
	0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61,
	0x73, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x76, 0x65, 0x72, 0x73, 0x65, 0x76, 0x65, 0x6e,
	0x2f, 0x74, 0x72, 0x79, 0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_pcommon_proto_rawDescOnce sync.Once
	file_pcommon_proto_rawDescData = file_pcommon_proto_rawDesc
)

func file_pcommon_proto_rawDescGZIP() []byte {
	file_pcommon_proto_rawDescOnce.Do(func() {
		file_pcommon_proto_rawDescData = protoimpl.X.CompressGZIP(file_pcommon_proto_rawDescData)
	})
	return file_pcommon_proto_rawDescData
}

var file_pcommon_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_pcommon_proto_goTypes = []interface{}{
	(*TransAirDrop)(nil),          // 0: pcommon.TransAirDrop
	(*TransTransfer)(nil),         // 1: pcommon.TransTransfer
	(*TransVotingInit)(nil),       // 2: pcommon.TransVotingInit
	(*TransVote)(nil),             // 3: pcommon.TransVote
	(*TransVotingFinish)(nil),     // 4: pcommon.TransVotingFinish
	(*Transaction)(nil),           // 5: pcommon.Transaction
	(*Block)(nil),                 // 6: pcommon.Block
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_pcommon_proto_depIdxs = []int32{
	7,  // 0: pcommon.TransAirDrop.timestamp:type_name -> google.protobuf.Timestamp
	7,  // 1: pcommon.TransTransfer.timestamp:type_name -> google.protobuf.Timestamp
	7,  // 2: pcommon.TransVotingInit.timestamp:type_name -> google.protobuf.Timestamp
	7,  // 3: pcommon.TransVote.timestamp:type_name -> google.protobuf.Timestamp
	7,  // 4: pcommon.TransVotingFinish.timestamp:type_name -> google.protobuf.Timestamp
	1,  // 5: pcommon.Transaction.transfer:type_name -> pcommon.TransTransfer
	0,  // 6: pcommon.Transaction.airdrop:type_name -> pcommon.TransAirDrop
	2,  // 7: pcommon.Transaction.votingInit:type_name -> pcommon.TransVotingInit
	3,  // 8: pcommon.Transaction.vote:type_name -> pcommon.TransVote
	4,  // 9: pcommon.Transaction.votingFinish:type_name -> pcommon.TransVotingFinish
	5,  // 10: pcommon.Block.trans:type_name -> pcommon.Transaction
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_pcommon_proto_init() }
func file_pcommon_proto_init() {
	if File_pcommon_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pcommon_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransAirDrop); i {
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
		file_pcommon_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransTransfer); i {
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
		file_pcommon_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransVotingInit); i {
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
		file_pcommon_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransVote); i {
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
		file_pcommon_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransVotingFinish); i {
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
		file_pcommon_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_pcommon_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
	}
	file_pcommon_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*Transaction_Transfer)(nil),
		(*Transaction_Airdrop)(nil),
		(*Transaction_VotingInit)(nil),
		(*Transaction_Vote)(nil),
		(*Transaction_VotingFinish)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pcommon_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pcommon_proto_goTypes,
		DependencyIndexes: file_pcommon_proto_depIdxs,
		MessageInfos:      file_pcommon_proto_msgTypes,
	}.Build()
	File_pcommon_proto = out.File
	file_pcommon_proto_rawDesc = nil
	file_pcommon_proto_goTypes = nil
	file_pcommon_proto_depIdxs = nil
}
