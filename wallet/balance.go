package wallet

import "errors"

type WalletBalance struct {
    Pubkey    []byte
    LastTransBlock uint64
    CurrentBalance float64
}

var usersBalances map[string]WalletBalance

func WalletIsBeing(pubkey []byte) bool{
    _, ok := usersBalances[string(pubkey)]
    return ok
}

func WalletInfo(pubkey []byte) (WalletBalance, error){
    value, ok := usersBalances[string(pubkey)]
    if !ok {
        return WalletBalance{}, errors.New("wallet information is not found")
    }
    return value, nil
}

