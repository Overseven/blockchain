package block

import "github.com/overseven/try-network/transaction"

type Block interface {
	GetBatchHash() (hash []byte, err error)
	GetHash() (hash []byte, err error)
	IsValid() error
	Mining(minerPubKey []byte, stop chan bool) ([]byte, error)
	HasTransaction(tr transaction.Transaction) (bool, error)
	AddTransaction(tr transaction.Transaction) error
	//RemoveTransaction(tr transaction.Transaction) error
}
