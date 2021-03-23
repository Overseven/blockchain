package block

import (
	"crypto/ecdsa"
	"testing"
	"time"

	tr "github.com/Overseven/blockchain/transaction"
	"github.com/Overseven/blockchain/transaction/transfer"
	"github.com/Overseven/blockchain/utility"
	"github.com/Overseven/blockchain/wallet"
	cr "github.com/ethereum/go-ethereum/crypto"
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

func generateTransaction(sndrPrivKey *ecdsa.PrivateKey, rcvrPubKey []byte, value, fee float64, message string) tr.Transaction {
	//transaction := transfer.Transfer{}
	data := tr.Data{}
	data.Pubkey = cr.CompressPubkey(&sndrPrivKey.PublicKey)
	data.Pay = value
	data.Fee = fee
	data.Receiver = rcvrPubKey
	data.Message = message
	{
		t := time.Now()
		data.Timestamp = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.UTC)
	}
	hashed := tr.GetHash(&data)

	//fmt.Println("Public key:", pubkey)

	sign, err := cr.Sign(hashed, sndrPrivKey)
	if err != nil {
		panic(err)
	}

	data.Sign = sign

	return &transfer.Transfer{Data: data}
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
	trans, err := transfer.New(sndrPrivKey, rcvrPubKey, 14, 0.5, "trans1")
	if err != nil {
		t.Error(err)
	}
	bl := Block{}
	transBase := tr.Transaction(trans)

	bl.AddTransaction(&transBase)

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
