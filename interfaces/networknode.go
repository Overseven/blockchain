package interfaces

import (
	"crypto/ecdsa"
	//pb "github.com/overseven/blockchain/protocol"
)

type NetworkNode interface {
	Init()
	SetPrivateKey(*ecdsa.PrivateKey)
	SetPort(uint32)
	GetPort() uint32
	StartListening() error
	GetWaitingTrans() []BlockElement
	// SendTransactionToAllNodes(element BlockElement) ([]pb.AddTransactionReply_Code, error)
	// SendTransaction(element BlockElement, nodeAddress string) (pb.AddTransactionReply_Code, error)
}
