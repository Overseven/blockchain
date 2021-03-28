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

func generateWallet() (privKey *ecdsa.PrivateKey, pubKey []byte, err error) {
	privKey, err = cr.GenerateKey()
	if err != nil {
		return nil, nil, err
	}

	pubKey = utility.PrivToPubKey(privKey)
	return
}

func TestBlockIsValid(t *testing.T) {
	var bchain interfaces.BlockConnecter = &chain.Chain{}

	var usersBalance interfaces.Balancer = &balance.Balance{}
	usersBalance.Init()

	sndrPrivKey, sndrPubKey, err := generateWallet()
	if err != nil {
		t.Error(err)
	}
	_, rcvrPubKey, err := generateWallet()
	if err != nil {
		t.Error(err)
	}

	_, minerPubKey, err := generateWallet()
	if err != nil {
		t.Error(err)
	}

	t.Log("Count of wallets: " + strconv.FormatInt(int64(usersBalance.CountOfWallets()), 10))
	//trans := generateTransaction(sndrPrivKey, rcvrPubKey, 14, 0.5, "trans1")

	airdropPubKey, airdropPrKey, err := wallet.LoadFromFile(airdropModeratorConfigFile)
	if err != nil {
		t.Error(err)
	}

	transaction.AirDropModeratorPubKey = airdropPubKey

	airPrKey, err := cr.ToECDSA(airdropPrKey[:32])
	if err != nil {
		t.Error(err)
	}
	airdrop, err := transaction.NewAirdrop(sndrPubKey, airPrKey, 100.0, 7.0)
	if err != nil {
		t.Error(err)
	}

	trans, err := transaction.NewTransfer(sndrPrivKey, rcvrPubKey, 14, 0.5, "trans1")
	if err != nil {
		t.Error(err)
	}

	block1 := bchain.NewBlock()
	err = block1.AddTransaction(airdrop)
	if err != nil {
		t.Error(err)
	}
	stop := make(chan bool)
	block1.Mining(minerPubKey, stop)

	block2 := bchain.NewBlock()
	err = block2.AddTransaction(trans)
	if err != nil {
		t.Error(err)
	}

	block2.Mining(minerPubKey, stop)

	err = usersBalance.FullCalc(bchain)
	if err != nil {
		t.Error(err)
	}

	if wallets := int64(usersBalance.CountOfWallets()); wallets != 3 {
		t.Error("Count of wallets after FullCalc: " + strconv.FormatInt(wallets, 10))
	} else {
		t.Log("Count of wallets after FullCalc: " + strconv.FormatInt(wallets, 10))
	}

	sndrWal, err := usersBalance.Info(sndrPubKey)
	if err != nil {
		t.Error(err)
	}
	rcvrWal, err := usersBalance.Info(rcvrPubKey)
	if err != nil {
		t.Error(err)
	}

	minerWal, err := usersBalance.Info(minerPubKey)
	if err != nil {
		t.Error(err)
	}

	if sndrWal.CurrentBalance != 85.5 || sndrWal.LastTransBlock != 1 {
		t.Errorf("Error. sender wallet: %f  last trans. block: %d", sndrWal.CurrentBalance, sndrWal.LastTransBlock)
	} else {
		t.Logf("Sender wallet: %f  last trans. block: %d\n", sndrWal.CurrentBalance, sndrWal.LastTransBlock)
	}

	if rcvrWal.CurrentBalance != 14.0 || rcvrWal.LastTransBlock != 1 {
		t.Errorf("Error. sender wallet: %f  last trans. block: %d", rcvrWal.CurrentBalance, rcvrWal.LastTransBlock)
	} else {
		t.Logf("Receiver wallet: %f  last trans. block: %d\n", rcvrWal.CurrentBalance, rcvrWal.LastTransBlock)
	}

	if minerWal.CurrentBalance != 7.5 || minerWal.LastTransBlock != 1 {
		t.Errorf("Error. Miner wallet: %f  last trans. block: %d", minerWal.CurrentBalance, minerWal.LastTransBlock)
	} else {
		t.Logf("Miner wallet: %f  last trans. block: %d\n", minerWal.CurrentBalance, minerWal.LastTransBlock)
	}
}

func Test3Airdrop1Block(t *testing.T) {
	var bchain interfaces.BlockConnecter = &chain.Chain{}

	var usersBalance interfaces.Balancer = &balance.Balance{}
	usersBalance.Init()

	_, receiver1, err := generateWallet()
	if err != nil {
		t.Error(err)
	}

	_, receiver2, err := generateWallet()
	if err != nil {
		t.Error(err)
	}

	_, receiver3, err := generateWallet()
	if err != nil {
		t.Error(err)
	}

	_, minerPubKey, err := generateWallet()
	if err != nil {
		t.Error(err)
	}

	t.Log("Count of wallets: " + strconv.FormatInt(int64(usersBalance.CountOfWallets()), 10))
	//trans := generateTransaction(sndrPrivKey, rcvrPubKey, 14, 0.5, "trans1")

	airdropPubKey, airdropPrKey, err := wallet.LoadFromFile(airdropModeratorConfigFile)
	if err != nil {
		t.Error(err)
	}

	transaction.AirDropModeratorPubKey = airdropPubKey

	airPrKey, err := cr.ToECDSA(airdropPrKey[:32])
	if err != nil {
		t.Error(err)
	}

	airdrop1, err := transaction.NewAirdrop(receiver1, airPrKey, 110.1, 11.1)
	if err != nil {
		t.Error(err)
	}

	airdrop2, err := transaction.NewAirdrop(receiver2, airPrKey, 120.2, 22.2)
	if err != nil {
		t.Error(err)
	}

	airdrop3, err := transaction.NewAirdrop(receiver3, airPrKey, 130.3, 33.3)
	if err != nil {
		t.Error(err)
	}

	block := bchain.NewBlock()
	err = block.AddTransaction(airdrop1)
	if err != nil {
		t.Error(err)
	}

	err = block.AddTransaction(airdrop2)
	if err != nil {
		t.Error(err)
	}

	err = block.AddTransaction(airdrop3)
	if err != nil {
		t.Error(err)
	}

	stop := make(chan bool)

	block.Mining(minerPubKey, stop)

	err = usersBalance.FullCalc(bchain)
	if err != nil {
		t.Error(err)
	}

	if wallets := int64(usersBalance.CountOfWallets()); wallets != 4 {
		t.Error("Error. Count of wallets after FullCalc: " + strconv.FormatInt(wallets, 10))
	} else {
		t.Log("Count of wallets after FullCalc: " + strconv.FormatInt(wallets, 10))
	}

	receiverWal1, err := usersBalance.Info(receiver1)
	if err != nil {
		t.Error(err)
	}

	receiverWal2, err := usersBalance.Info(receiver2)
	if err != nil {
		t.Error(err)
	}

	receiverWal3, err := usersBalance.Info(receiver3)
	if err != nil {
		t.Error(err)
	}

	minerWal, err := usersBalance.Info(minerPubKey)
	if err != nil {
		t.Error(err)
	}

	if receiverWal1.CurrentBalance != 110.1 || receiverWal1.LastTransBlock != 0 {
		t.Errorf("Error. receiver1 wallet: %f  last trans. block: %d", receiverWal1.CurrentBalance, receiverWal1.LastTransBlock)
	} else {
		t.Logf("Receiver1 wallet: %f  last trans. block: %d\n", receiverWal1.CurrentBalance, receiverWal1.LastTransBlock)
	}

	if receiverWal2.CurrentBalance != 120.2 || receiverWal2.LastTransBlock != 0 {
		t.Errorf("Error. receiver2 wallet: %f  last trans. block: %d", receiverWal2.CurrentBalance, receiverWal2.LastTransBlock)
	} else {
		t.Logf("Receiver2 wallet: %f  last trans. block: %d\n", receiverWal2.CurrentBalance, receiverWal2.LastTransBlock)
	}

	if receiverWal3.CurrentBalance != 130.3 || receiverWal3.LastTransBlock != 0 {
		t.Errorf("Error. receiver3 wallet: %f  last trans. block: %d", receiverWal3.CurrentBalance, receiverWal3.LastTransBlock)
	} else {
		t.Logf("Receiver3 wallet: %f  last trans. block: %d\n", receiverWal3.CurrentBalance, receiverWal3.LastTransBlock)
	}

	if minerWal.CurrentBalance != 66.6 || minerWal.LastTransBlock != 0 {
		t.Errorf("Error. Miner wallet: %f  last trans. block: %d", minerWal.CurrentBalance, minerWal.LastTransBlock)
	} else {
		t.Logf("Miner wallet: %f  last trans. block: %d\n", minerWal.CurrentBalance, minerWal.LastTransBlock)
	}
}
