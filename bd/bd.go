package bd

import (
	"errors"

	"github.com/overseven/blockchain/interfaces"
	"github.com/overseven/blockchain/transaction"
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

func PutTransaction(tr interfaces.BlockElement) error {
	if err := IsOpen(); err != nil {
		return err
	}

	bKey := []byte("t")
	bKey = append(bKey, transaction.GetHash(tr.GetData())...)
	bData := transaction.Bytes(tr)
	return db.Put(bKey, bData, nil)
}

func PutBlock() {

}

func PutSnapshot() {

}
