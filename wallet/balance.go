package wallet

import (
    "errors"
    "sync"
)

type Balance struct {
    Pubkey    []byte
    LastTransBlock uint64
    CurrentBalance float64
}

var (
	mutex sync.Mutex
    usersBalances map[string]Balance
)

func Init(){
    if len(usersBalances) < 1 {
        usersBalances = make(map[string]Balance)
    }
}

func IsBeing(pubkey []byte) bool{
    _, ok := usersBalances[string(pubkey)]
    return ok
}

func Info(pubkey []byte) (Balance, error){
    value, ok := usersBalances[string(pubkey)]
    if !ok {
        return Balance{}, errors.New("wallet information is not found")
    }
    return value, nil
}

func Update(pubkey []byte, lastTransBlock uint64, sum float64) (isNew bool){
    mutex.Lock()
    defer mutex.Unlock()

    // TODO: add pubkey check

    if w, ok := usersBalances[string(pubkey)]; ok {
        w.LastTransBlock = lastTransBlock
        w.CurrentBalance = sum
        return false
    }

    newWallet := Balance{Pubkey: pubkey, LastTransBlock: lastTransBlock, CurrentBalance: sum}
    usersBalances[string(pubkey)] = newWallet

    return true
}