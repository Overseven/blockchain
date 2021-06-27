package trlists

import (
	"sync"

	"github.com/overseven/blockchain/transaction"
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
		first.trans[string(t.Hash())] = t
	}
}

func FirstToSecond(sended []transaction.Transaction) {
	first.lock.Lock()
	second.lock.Lock()
	defer first.lock.Unlock()
	defer second.lock.Unlock()

	for _, t := range sended {
		hash := string(t.Hash())
		_, ok := first.trans[hash]
		if ok {
			delete(first.trans, hash)
		}
		second.trans[string(t.Hash())] = t
	}
}

func SecondToFirst(canceled []transaction.Transaction) {
	first.lock.Lock()
	second.lock.Lock()
	defer first.lock.Unlock()
	defer second.lock.Unlock()

	for _, t := range canceled {
		hash := string(t.Hash())
		_, ok := second.trans[hash]
		if ok {
			delete(second.trans, hash)
		}
		first.trans[string(t.Hash())] = t
	}
}

func Remove(trs []transaction.Transaction) {
	first.lock.Lock()
	second.lock.Lock()
	defer first.lock.Unlock()
	defer second.lock.Unlock()

	for _, t := range trs {
		hash := string(t.Hash())
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
