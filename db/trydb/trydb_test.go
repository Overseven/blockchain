package trydb_test

import (
	"github.com/overseven/try-network/db/trydb"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/syndtr/goleveldb/leveldb/util"
	"testing"
)

func TestName(t *testing.T) {
	err := trydb.OpenFile("tmp_bibus", &opt.Options{})
	if err != nil {
		t.Error(err)
	}

	prefix, key := "aaa", []byte("ff")
	err = trydb.Put(prefix, key, []byte("asf"))
	if err != nil {
		t.Error(err)
	}

	key2 := []byte("ff423")
	err = trydb.Put(prefix, key2, []byte("11223344"))
	if err != nil {
		t.Error(err)
	}

	prefix2, key3 := "a4a", []byte("GG31")
	err = trydb.Put(prefix2, key3, []byte("ASMR"))
	if err != nil {
		t.Error(err)
	}

	dd, err := trydb.Get(prefix, key)
	if err != nil {
		t.Error(err)
	}
	t.Log("d: ", dd)

	batch := new(leveldb.Batch)
	batch.Put([]byte("aaaff"), []byte("value"))
	//err = db.Write(batch, nil)

	iter := trydb.NewIterator(&util.Range{Start: []byte("aaa"), Limit: []byte("aab")}, nil)
	for iter.Next() {
		// Use key/value.
		key := iter.Key()
		value := iter.Value()
		t.Log("key: ", key, " value:", value)
	}
	iter.Release()
	err = iter.Error()
}
