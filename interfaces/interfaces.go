package interfaces

import "time"

type Chainable interface {
	IsValid(startIndx, endIndx uint64) (bool, uint64)
	GetBlocks() []Blockable
	SetBlocks([]Blockable)
}

type Blockable interface {
	GetId() uint64
	GetBatchHash() (hash []byte)
	GetWalletStatsHash() (hash []byte)
	GetHash() (hash []byte)
	IsValid(blockchain Chainable) (bool, error)
	Mining(stop chan bool) []byte
	GetTransaction() []Transferable
	HasTransaction(transact *Transferable) (index int, has bool)
	AddTransaction(tr *Transferable) error
}

type Type int64

type Transferable interface {
	SetData(*Data)
	GetData() *Data
	Verify() error
}

type Data struct {
	Type      Type
	Pubkey    []byte
	Receiver  []byte
	Message   string
	Timestamp time.Time
	Pay       float64
	Fee       float64
	Sign      []byte
}
