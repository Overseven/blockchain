package interfaces

import (
	"time"
)

type BlockConnecter interface {
	IsValid(startIndx, endIndx uint64) (bool, uint64)
	GetBlocks() []TransactionsContainer
	SetBlocks([]TransactionsContainer)
	NewBlock() TransactionsContainer
	AppendBlock(TransactionsContainer)
}

type TransactionsContainer interface {
	GetId() uint64
	SetId(uint64)
	GetBatchHash() (hash []byte)
	GetHash() (hash []byte)
	GetPrevHash() []byte
	SetPrevHash([]byte)
	IsValid(BlockConnecter, Balancer) (bool, error)
	Mining(minerPubKey []byte, stop chan bool) []byte
	GetTransactions() []BlockElement
	SetTransactions([]BlockElement)
	HasTransaction(BlockElement) (index int, has bool)
	AddTransaction(BlockElement) error
	GetDifficulty() uint64
	SetDifficulty(uint64)
	GetMiner() []byte
	SetMiner([]byte)
	GetNonce() []byte
	SetNonce([]byte)
}

type Type int64

type BlockElement interface {
	SetData(*Data)
	GetData() *Data
	Verify(Balancer) error
}

type Data struct {
	Type      Type
	Sender    []byte
	Receiver  []byte
	Message   string
	Timestamp time.Time
	Pay       float64
	Fee       float64
	Sign      []byte
}

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
