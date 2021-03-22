package blockchain

import (
	"crypto/ecdsa"
	"fmt"
	"testing"
	"time"

	tr "github.com/Overseven/blockchain/transaction"
	"github.com/Overseven/blockchain/wallet"
	cr "github.com/ethereum/go-ethereum/crypto"
	"github.com/overseven/blockchain/blockchain"
)

func generateWallet(value float64) (privKey *ecdsa.PrivateKey, pubKey []byte) {
	privKey, err := cr.GenerateKey()
	if err != nil {
		fmt.Errorf(err.Error())
	}

	pubKey = cr.CompressPubkey(&privKey.PublicKey)

	wallet.Update(pubKey, 0, value)
	return
}

func generateTransaction(sndrPrivKey *ecdsa.PrivateKey, rcvrPubKey []byte, value, fee float64, message string) tr.Transaction {
	transaction := tr.Transaction{}
	transaction.Pubkey = cr.CompressPubkey(&sndrPrivKey.PublicKey)
	transaction.Pay = value
	transaction.Fee = fee
	transaction.Receiver = rcvrPubKey
	transaction.Message = message
	{
		t := time.Now()
		transaction.Timestamp = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.UTC)
	}
	hashed := transaction.GetHash()

	//fmt.Println("Public key:", pubkey)

	sign, err := cr.Sign(hashed, sndrPrivKey)
	if err != nil {
		panic(err)
	}

	transaction.Sign = sign

	return transaction
}

func TestBlockIsValid(t *testing.T) {
	// TODO: finish test
	sndrPrivKey, sndrPubKey := generateWallet(15.0)
	_, rcvrPubKey := generateWallet(0.0)

	trans := generateTransaction(sndrPrivKey, rcvrPubKey, 14, 0.5, "trans1")
	bl := blockchain.Block{}
	bl.AddTransaction(&trans)

	sndrWal, err := wallet.Info(sndrPubKey)
	if err != nil {
		panic(err)
	}
	rcvrWal, err := wallet.Info(rcvrPubKey)
	if err != nil {
		panic(err)
	}

	if sndrWal.CurrentBalance != 0.5 || sndrWal.LastTransBlock != 0 {
		fmt.Errorf("Error. sender wallet: %f  last trans. block: %d", sndrWal.CurrentBalance, sndrWal.LastTransBlock)
	} else {
		fmt.Printf("Sender wallet: %f  last trans. block: %d\n", sndrWal.CurrentBalance, sndrWal.LastTransBlock)
	}

	if sndrWal.CurrentBalance != 0.5 || sndrWal.LastTransBlock != 0 {
		fmt.Errorf("Error. sender wallet: %f  last trans. block: %d", sndrWal.CurrentBalance, sndrWal.LastTransBlock)
	} else {
		fmt.Printf("Receiver wallet: %f  last trans. block: %d\n", rcvrWal.CurrentBalance, rcvrWal.LastTransBlock)
	}
}
