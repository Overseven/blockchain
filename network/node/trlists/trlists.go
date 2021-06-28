package trlists

import (
	"sync"

	"github.com/overseven/try-network/transaction"
)

var (
	first, second List
)

type List struct {
	trans map[string]transaction.Transaction
	lock  sync.Mutex
}

func init() {
	ClearAll()
}

func AddToFirst(toSend []transaction.Transaction) {
	first.lock.Lock()
	defer first.lock.Unlock()

	for _, t := range toSend {
		tHash, err := t.Hash(map[transaction.TransFlag]bool{}) // TODO: need some flags?
		hash := string(tHash)
		if err != nil {
			return
		}
		first.trans[hash] = t
	}
}

func FirstToSecond(sended []transaction.Transaction) {
	first.lock.Lock()
	second.lock.Lock()
	defer first.lock.Unlock()
	defer second.lock.Unlock()

	for _, t := range sended {
		tHash, err := t.Hash(map[transaction.TransFlag]bool{}) // TODO: need some flags?
		hash := string(tHash)
		if err != nil {
			return
		}
		_, ok := first.trans[hash]
		if ok {
			delete(first.trans, hash)
		}
		second.trans[hash] = t
	}
}

func SecondToFirst(canceled []transaction.Transaction) {
	first.lock.Lock()
	second.lock.Lock()
	defer first.lock.Unlock()
	defer second.lock.Unlock()

	for _, t := range canceled {
		tHash, err := t.Hash(map[transaction.TransFlag]bool{}) // TODO: need some flags?
		hash := string(tHash)
		if err != nil {
			return
		}
		_, ok := second.trans[hash]
		if ok {
			delete(second.trans, hash)
		}
		first.trans[hash] = t
	}
}

func Remove(trs []transaction.Transaction) {
	first.lock.Lock()
	second.lock.Lock()
	defer first.lock.Unlock()
	defer second.lock.Unlock()

	for _, t := range trs {
		tHash, err := t.Hash(map[transaction.TransFlag]bool{}) // TODO: need some flags?
		if err != nil {
			return
		}
		hash := string(tHash)
		_, ok := first.trans[hash]
		if ok {
			delete(first.trans, hash)
		}
		_, ok = second.trans[hash]
		if ok {
			delete(second.trans, hash)
		}
	}

}

func ClearAll() {
	first.lock.Lock()
	second.lock.Lock()
	defer first.lock.Unlock()
	defer second.lock.Unlock()

	first.trans = map[string]transaction.Transaction{}
	second.trans = map[string]transaction.Transaction{}
}
