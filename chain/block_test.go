package chain

import (
	"crypto/ecdsa"
	"github.com/Overseven/blockchain/utility"
	"github.com/Overseven/blockchain/wallet"
	cr "github.com/ethereum/go-ethereum/crypto"
	"github.com/overseven/blockchain/transaction"
	. "github.com/overseven/blockchain/transaction/itransaction"
	"testing"
)

func generateWallet(value float64) (privKey *ecdsa.PrivateKey, pubKey []byte, err error) {
	privKey, err = cr.GenerateKey()
	if err != nil {
		return nil, nil, err
	}

	pubKey = utility.PrivToPubKey(privKey)

	wallet.Update(pubKey, 0, value)
	return
}



func TestBlockIsValid(t *testing.T) {
	// TODO: finish test
	wallet.Init()
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
	err = wallet.FullCalc()
	if err != nil {
		t.Error(err)
	}
	var bl Block
	transBase := ITransaction(trans)

	err = bl.AddTransaction(&transBase)
	if err != nil {
		t.Error(err)
	}

	sndrWal, err := wallet.Info(sndrPubKey)
	if err != nil {
		panic(err)
	}
	rcvrWal, err := wallet.Info(rcvrPubKey)
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
