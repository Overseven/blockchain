package test

import (
	"crypto/ecdsa"
	"testing"

	cr "github.com/ethereum/go-ethereum/crypto"
	"github.com/overseven/balance"
	"github.com/overseven/blockchain/interfaces"
	"github.com/overseven/blockchain/transaction"

	"github.com/overseven/blockchain/utility"
)

func generateWallet(value float64) (privKey *ecdsa.PrivateKey, pubKey []byte, err error) {
	privKey, err = cr.GenerateKey()
	if err != nil {
		return nil, nil, err
	}

	pubKey = utility.PrivToPubKey(privKey)

	UsersBalance.Update(pubKey, 0, value)
	return
}

func TestBlockIsValid(t *testing.T) {
	// TODO: finish test
	var balance interfaces.Balancer = &balance.Balance{}
	balance.Init()

	sndrPrivKey, sndrPubKey, err := generateWallet(15.0)
	if err != nil {
		t.Error(err)
	}
	_, rcvrPubKey, err := generateWallet(0.0)
	if err != nil {
		t.Error(err)
	}

	//trans := generateTransaction(sndrPrivKey, rcvrPubKey, 14, 0.5, "trans1")
	trans, err := transaction.NewTransfer(sndrPrivKey, rcvrPubKey, 14, 0.5, "trans1")
	if err != nil {
		t.Error(err)
	}
	err = balance.FullCalc()
	if err != nil {
		t.Error(err)
	}
	var bl Block
	transBase := interfaces.Transferable(trans)

	err = bl.AddTransaction(&transBase)
	if err != nil {
		t.Error(err)
	}

	sndrWal, err := UsersBalance.Info(sndrPubKey)
	if err != nil {
		panic(err)
	}
	rcvrWal, err := UsersBalance.Info(rcvrPubKey)
	if err != nil {
		panic(err)
	}

	if sndrWal.CurrentBalance != 0.5 || sndrWal.LastTransBlock != 0 {
		t.Errorf("Error. sender wallet: %f  last trans. block: %d", sndrWal.CurrentBalance, sndrWal.LastTransBlock)
	} else {
		t.Logf("Sender wallet: %f  last trans. block: %d\n", sndrWal.CurrentBalance, sndrWal.LastTransBlock)
	}

	if rcvrWal.CurrentBalance != 0.5 || rcvrWal.LastTransBlock != 0 {
		t.Errorf("Error. sender wallet: %f  last trans. block: %d", rcvrWal.CurrentBalance, rcvrWal.LastTransBlock)
	} else {
		t.Logf("Receiver wallet: %f  last trans. block: %d\n", rcvrWal.CurrentBalance, rcvrWal.LastTransBlock)
	}
}
