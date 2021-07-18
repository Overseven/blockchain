package db

import (
	"github.com/overseven/try-network/block"
	"github.com/overseven/try-network/transaction"
)

type Database interface {
	FileController
	AccountStats
	VotingStats

	TryAddBlock(block.Block) error
	GetLastBlock() (block.Block, error)
	GetLastVoting() (transaction.Transaction, error)
}

type FileController interface {
	OpenFile(filepath string) error
}

type AccountStats interface {
	GetTransaction(address []byte, transCounter uint32) (transaction.Transaction, error)
	GetTransactions(address []byte, transCounterBegin, transCounterEnd uint32) ([]transaction.Transaction, error)
	GetTransactionCounter(address []byte) (uint32, error)
	GetBalance(address []byte) (float64, error) // TODO: replace float by fixed point type

}

type VotingStats interface {
	GetVoting(id transaction.VotingId) (transaction.Transaction, error)
}
