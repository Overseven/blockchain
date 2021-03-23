package chain

import (
	tr "github.com/Overseven/blockchain/transaction"
)

type Chain interface {
	IsValid(startIndx, endIndx uint64)
	GetBlocks() []Block
}

//var B17 Blockchain
type Block interface {
	GetBatchHash() (hash []byte)
	GetWalletStatsHash() (hash []byte)
	GetHash() (hash []byte)
	IsValid(blockchain *Chain) (bool, error)
	Mining(stop chan bool) []byte
	HasTransaction(transact *tr.Transaction) (index int, has bool)
	AddTransaction(tr *tr.Transaction) error
}
