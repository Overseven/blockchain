package test

import (
	"crypto/ecdsa"
	"strconv"
	"testing"

	cr "github.com/ethereum/go-ethereum/crypto"
	"github.com/overseven/blockchain/balance"
	"github.com/overseven/blockchain/wallet"

	"github.com/overseven/blockchain/chain"
	"github.com/overseven/blockchain/interfaces"
	"github.com/overseven/blockchain/transaction"

	"github.com/overseven/blockchain/utility"
)

const (
	airdropModeratorConfigFile = "..\\wallet.cfg"
)

func generateWallet(value float64, balance interfaces.Balancer) (privKey *ecdsa.PrivateKey, pubKey []byte, err error) {
	privKey, err = cr.GenerateKey()
	if err != nil {
		return nil, nil, err
	}

	pubKey = utility.PrivToPubKey(privKey)

	balance.Update(pubKey, 0, value)
	return
}

func TestBlockIsValid(t *testing.T) {
	// TODO: finish test

	var bchain interfaces.Chainable = &chain.Chain{}

	var usersBalance interfaces.Balancer = &balance.Balance{}
	usersBalance.Init()

	sndrPrivKey, sndrPubKey, err := generateWallet(15.0, usersBalance)
	if err != nil {
		t.Error(err)
	}
	_, rcvrPubKey, err := generateWallet(0.1, usersBalance)
	if err != nil {
		t.Error(err)
	}

	t.Log("Count of wallets: " + strconv.FormatInt(int64(usersBalance.CountOfWallets()), 10))
	//trans := generateTransaction(sndrPrivKey, rcvrPubKey, 14, 0.5, "trans1")

	_, airdropPrKey, err := wallet.LoadFromFile(airdropModeratorConfigFile)
	if err != nil {
		t.Error(err)
	}
	airPrKey, err := cr.ToECDSA(airdropPrKey[:32])
	if err != nil {
		t.Error(err)
	}
	airdrop, err := transaction.NewAirdrop(sndrPubKey, airPrKey, bchain, 100.0, usersBalance)
	if err != nil {
		t.Error(err)
	}

	trans, err := transaction.NewTransfer(sndrPrivKey, rcvrPubKey, 14, 0.5, "trans1", usersBalance)
	if err != nil {
		t.Error(err)
	}

	block1 := bchain.NewBlock()
	err = block1.AddTransaction(airdrop)
	if err != nil {
		t.Error(err)
	}

	block2 := bchain.NewBlock()
	err = block2.AddTransaction(trans)
	if err != nil {
		t.Error(err)
	}

	err = usersBalance.FullCalc(bchain)
	if err != nil {
		t.Error(err)
	}

	t.Log("Count of wallets after FullCalc: " + strconv.FormatInt(int64(usersBalance.CountOfWallets()), 10))

	sndrWal, err := usersBalance.Info(sndrPubKey)
	if err != nil {
		t.Error(err)
	}
	rcvrWal, err := usersBalance.Info(rcvrPubKey)
	if err != nil {
		t.Error(err)
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
