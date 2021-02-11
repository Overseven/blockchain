package blockchain

import (
	"github.com/overseven/blockchain/transaction"
)

type Block struct{
	Id uint64
	Transactions []transaction.Transaction
	PrevHash []byte
	Hash []byte
}