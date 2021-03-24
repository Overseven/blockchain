package ichain

import (
	"github.com/overseven/blockchain/transaction/itransaction"
)

type IChain interface {
	IsValid(startIndx, endIndx uint64) (bool, uint64)
	GetBlocks() []IBlock
	SetBlocks([]IBlock)
}

//var B17 Blockchain
type IBlock interface {
	GetId() uint64
	GetBatchHash() (hash []byte)
	GetWalletStatsHash() (hash []byte)
	GetHash() (hash []byte)
	IsValid(blockchain IChain) (bool, error)
	Mining(stop chan bool) []byte
	GetTransaction() []itransaction.ITransaction
	HasTransaction(transact *itransaction.ITransaction) (index int, has bool)
	AddTransaction(tr *itransaction.ITransaction) error
}
