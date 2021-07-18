package db

import (
	"github.com/overseven/try-network/block"
	"github.com/overseven/try-network/transaction"
)

type Database interface {
	FileProcessor
	BlockProcessor
	AccountStats
	VotingStats
}

type FileProcessor interface {
	OpenFile(filepath string) error
}

type BlockProcessor interface {
	TryAddBlock(block.Block) error
	Synchronize(blockBegin uint64, input chan block.Block) error
	GetBlock(blockId uint64) (block.Block, error)
	GetLastBlock() (block.Block, error)
}

type AccountStats interface {
	GetTransaction(address []byte, transCounter transaction.TransCounter) (transaction.Transaction, error)
	GetTransactions(address []byte, transCounterBegin, transCounterEnd transaction.TransCounter) ([]transaction.Transaction, error)
	GetTransactionCounter(address []byte) (transaction.TransCounter, error)
	GetBalance(address []byte) (transaction.Balance, error)
}

type VotingStats interface {
	GetVoting(id transaction.VotingId) (transaction.Transaction, error)
	GetLastVoting() (transaction.Transaction, error)
}
