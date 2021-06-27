package db

import (
	"errors"

	"github.com/overseven/try-network/block"
	"github.com/overseven/try-network/transaction"
	"github.com/overseven/try-network/utility"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

var (
	db *leveldb.DB
)

func OpenFile(filepath string, o *opt.Options) error {
	db_temp, err := leveldb.OpenFile(filepath, o)
	if err != nil {
		return err
	}
	db = db_temp

	return nil
}

func IsOpen() error {
	if db == nil {
		return errors.New("db must be initialized")
	}
	return nil
}

func Put(prefix string, key, data []byte) error {
	bKey := append([]byte(prefix), key...)
	err := db.Put(bKey, data, nil)
	return err
}

func Get(prefix string, key []byte) ([]byte, error) {
	bKey := append([]byte(prefix), key...)
	return db.Get(bKey, nil)
}

func PutTransaction(tr transaction.Transaction) error {
	bKey := []byte("t")
	bKey = append(bKey, tr.Hash()...)
	bData, err := tr.Bytes()
	if err != nil {
		return err
	}
	return db.Put(bKey, bData, nil)
}

func GetTransaction(hash []byte) (transaction.Transaction, error) {
	if len(hash) != 32 {
		return nil, errors.New("incorrect hash len")
	}

	bKey := []byte("t")
	bKey = append(bKey, hash...)
	bValue, err := db.Get(bKey, nil)
	if err != nil {
		return nil, err
	}

	tr, err := TransactionFromBytes(bValue)

	if err != nil {
		return nil, err
	}

	return tr, nil
}

func PutBlock(b block.Block) {

}

func GetBlock(id uint64) (*block.Block, error) {
	bKey := []byte("b")
	bKey = append(bKey, utility.UInt64Bytes(id)...)
	bValue, err := db.Get(bKey, nil)
	if err != nil {
		return nil, err
	}

	b, err := BlockFromBytes(bValue)
	return &b, nil
}

func CreateSnapshot(id uint8) {

}

func GetLastBlock() block.Block {

	return block.Block{}
}
