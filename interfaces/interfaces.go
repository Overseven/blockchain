package interfaces

import "time"

type Chainable interface {
	IsValid(startIndx, endIndx uint64) (bool, uint64)
	GetBlocks() []Blockable
	SetBlocks([]Blockable)
	NewBlock() Blockable
	AppendBlock(Blockable)
}

type Blockable interface {
	GetId() uint64
	SetId(uint64)
	GetBatchHash() (hash []byte)
	GetHash() (hash []byte)
	GetPrevHash() []byte
	SetPrevHash([]byte)
	IsValid(Chainable, Balancer) (bool, error)
	Mining(minerPubKey []byte, stop chan bool) []byte
	GetTransactions() []Transferable
	SetTransactions([]Transferable)
	HasTransaction(Transferable) (index int, has bool)
	AddTransaction(Transferable) error
	GetDifficulty() uint64
	SetDifficulty(uint64)
	GetMiner() []byte
	SetMiner([]byte)
	GetNonce() []byte
	SetNonce([]byte)
}

type Type int64

type Transferable interface {
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
	FullCalc(Chainable) error
	CountOfWallets() int
}

type BalanceStat struct {
	Pubkey         []byte
	LastTransBlock uint64
	CurrentBalance float64
}
