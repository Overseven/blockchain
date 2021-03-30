package interfaces

import (
	pb "github.com/overseven/blockchain/protocol"
)

type ClientMode int32

const (
	ModeFull = iota
	ModeLight
)

type NetworkClient interface {
	Networker
	WalletOwner
	ChainHolder
	Init()
	SetMode(mode ClientMode)
	GetMode() ClientMode
	SendTransactions(element BlockElement, nodesAddress []string) ([]pb.AddTransactionReply_Code, error)
}
