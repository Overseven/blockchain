package interfaces

import (
	"crypto/ecdsa"
)

// type BlockConnecter interface {
// 	IsValid(startIndx, endIndx uint64) (bool, uint64)
// 	GetBlocks() []TransactionsContainer
// 	SetBlocks([]TransactionsContainer)
// 	NewBlock() TransactionsContainer
// 	AppendBlock(TransactionsContainer)
// }

type Balancer interface {
	Init()
	IsBeing(pubkey []byte) bool
	Info(pubkey []byte) (BalanceStat, error)
	Update(pubkey []byte, lastTransBlock uint64, sum float64) (isNew bool)
	Clear()
	FullCalc(BlockConnecter) error
	CountOfWallets() int
}

type BalanceStat struct {
	Pubkey         []byte
	LastTransBlock uint64
	CurrentBalance float64
}

type WalletOwner interface {
	GetPrivateKey() *ecdsa.PrivateKey
	SetPrivateKey(*ecdsa.PrivateKey)
	GetPublicKey() []byte
	SetPublicKey([]byte)
}

type ChainHolder interface {
	GetChain() BlockConnecter
	SetChain(BlockConnecter)
}

type Networker interface {
	SetPort(uint32)
	GetPort() uint32
}
