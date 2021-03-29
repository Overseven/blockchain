package interfaces

import (
	"crypto/ecdsa"

	pb "github.com/overseven/blockchain/protocol"
)

type ClientMode int32

const (
	ModeFull = iota
	ModeLight
)

type NetworkClient interface {
	Init()
	SetMode(mode ClientMode)
	GetMode() ClientMode
	SetPrivateKey(*ecdsa.PrivateKey)
	SetPort(uint32)
	GetPort() uint32
	SendTransactionToAllNodes(element BlockElement) ([]pb.AddTransactionReply_Code, error)
	SendTransaction(element BlockElement, nodeAddress string) (pb.AddTransactionReply_Code, error)
}
