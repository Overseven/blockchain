package test

import (
	"strconv"
	"testing"

	cr "github.com/ethereum/go-ethereum/crypto"
	"github.com/overseven/blockchain/balance"
	pb "github.com/overseven/blockchain/protocol"
	"github.com/overseven/blockchain/wallet"

	"github.com/overseven/blockchain/chain"
	"github.com/overseven/blockchain/interfaces"
	"github.com/overseven/blockchain/transaction"

	"github.com/overseven/blockchain/utility"
)


func TestBlockIsValid(t *testing.T) {
	var bchain interfaces.BlockConnecter = &chain.Chain{}

	var usersBalance interfaces.Balancer = &balance.Balance{}
	usersBalance.Init()

	sndrPrivKey, sndrPubKey, err := utility.GenerateWallet()
	if err != nil {
		t.Error(err)
	}
	_, rcvrPubKey, err := utility.GenerateWallet()
	if err != nil {
		t.Error(err)
	}

	_, minerPubKey, err := utility.GenerateWallet()
	if err != nil {
		t.Error(err)
	}

	//t.Log("Count of wallets: " + strconv.FormatInt(int64(usersBalance.CountOfWallets()), 10))
	//trans := generateTransaction(sndrPrivKey, rcvrPubKey, 14, 0.5, "trans1")

	airdropPubKey, airdropPrKey, err := config.LoadFromFile(airdropModeratorConfigFile)
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
		//t.Log("Count of wallets after FullCalc: " + strconv.FormatInt(wallets, 10))
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
		//t.Logf("Sender wallet: %f  last trans. block: %d\n", sndrWal.CurrentBalance, sndrWal.LastTransBlock)
	}

	if rcvrWal.CurrentBalance != 14.0 || rcvrWal.LastTransBlock != 1 {
		t.Errorf("Error. sender wallet: %f  last trans. block: %d", rcvrWal.CurrentBalance, rcvrWal.LastTransBlock)
	} else {
		//t.Logf("Receiver wallet: %f  last trans. block: %d\n", rcvrWal.CurrentBalance, rcvrWal.LastTransBlock)
	}

	if minerWal.CurrentBalance != 7.5 || minerWal.LastTransBlock != 1 {
		t.Errorf("Error. Miner wallet: %f  last trans. block: %d", minerWal.CurrentBalance, minerWal.LastTransBlock)
	} else {
		//t.Logf("Miner wallet: %f  last trans. block: %d\n", minerWal.CurrentBalance, minerWal.LastTransBlock)
	}
}

func Test3Airdrop1Block(t *testing.T) {
	var bchain interfaces.BlockConnecter = &chain.Chain{}

	var usersBalance interfaces.Balancer = &balance.Balance{}
	usersBalance.Init()

	_, receiver1, err := utility.GenerateWallet()
	if err != nil {
		t.Error(err)
	}

	_, receiver2, err := utility.GenerateWallet()
	if err != nil {
		t.Error(err)
	}

	_, receiver3, err := utility.GenerateWallet()
	if err != nil {
		t.Error(err)
	}

	_, minerPubKey, err := utility.GenerateWallet()
	if err != nil {
		t.Error(err)
	}

	//t.Log("Count of wallets: " + strconv.FormatInt(int64(usersBalance.CountOfWallets()), 10))
	//trans := generateTransaction(sndrPrivKey, rcvrPubKey, 14, 0.5, "trans1")

	airdropPubKey, airdropPrKey, err := config.LoadFromFile(airdropModeratorConfigFile)
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
		//t.Log("Count of wallets after FullCalc: " + strconv.FormatInt(wallets, 10))
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
		//t.Logf("Receiver1 wallet: %f  last trans. block: %d\n", receiverWal1.CurrentBalance, receiverWal1.LastTransBlock)
	}

	if receiverWal2.CurrentBalance != 120.2 || receiverWal2.LastTransBlock != 0 {
		t.Errorf("Error. receiver2 wallet: %f  last trans. block: %d", receiverWal2.CurrentBalance, receiverWal2.LastTransBlock)
	} else {
		//t.Logf("Receiver2 wallet: %f  last trans. block: %d\n", receiverWal2.CurrentBalance, receiverWal2.LastTransBlock)
	}

	if receiverWal3.CurrentBalance != 130.3 || receiverWal3.LastTransBlock != 0 {
		t.Errorf("Error. receiver3 wallet: %f  last trans. block: %d", receiverWal3.CurrentBalance, receiverWal3.LastTransBlock)
	} else {
		//t.Logf("Receiver3 wallet: %f  last trans. block: %d\n", receiverWal3.CurrentBalance, receiverWal3.LastTransBlock)
	}

	if minerWal.CurrentBalance != 66.6 || minerWal.LastTransBlock != 0 {
		t.Errorf("Error. Miner wallet: %f  last trans. block: %d", minerWal.CurrentBalance, minerWal.LastTransBlock)
	} else {
		//t.Logf("Miner wallet: %f  last trans. block: %d\n", minerWal.CurrentBalance, minerWal.LastTransBlock)
	}
}

func TestNetworkBlockGeneration(t *testing.T) {
	nd, err := createNode()
	if err != nil {
		t.Error(err)
	}

	cl, err := createClient()
	if err != nil {
		t.Error(err)
	}

	_, rcvrPubKey1, err := utility.GenerateWallet()
	if err != nil {
		t.Error(err)
	}

	_, rcvrPubKey2, err := utility.GenerateWallet()
	if err != nil {
		t.Error(err)
	}

	value1, value2 := 152.313, 535.5
	fee1, fee2 := 0.0004, 1.32
	//message := "network test"

	airdropPubKey, airdropPrKey, err := config.LoadFromFile(airdropModeratorConfigFile)
	if err != nil {
		t.Error(err)
	}

	transaction.AirDropModeratorPubKey = airdropPubKey

	airPrKey, err := cr.ToECDSA(airdropPrKey[:32])
	if err != nil {
		t.Error(err)
	}

	airdrop1, err := transaction.NewAirdrop(rcvrPubKey1, airPrKey, value1, fee1)
	if err != nil {
		t.Error(err)
	}

	airdrop2, err := transaction.NewAirdrop(rcvrPubKey2, airPrKey, value2, fee2)
	if err != nil {
		t.Error(err)
	}

	stop := make(chan bool)
	err = nd.StartListening(stop)
	if err != nil {
		t.Error(err)
	}
	defer func() { stop <- true }()

	replyCode1, err := cl.SendTransaction(airdrop1, "localhost:"+strconv.FormatUint(uint64(nodePort), 10))
	if err != nil {
		t.Error(err)
	}

	if replyCode1 != pb.AddTransactionReply_TR_Ok {
		t.Error("replyCode1 != Ok")
	}

	replyCode2, err := cl.SendTransaction(airdrop2, "localhost:"+strconv.FormatUint(uint64(nodePort), 10))
	if err != nil {
		t.Error(err)
	}

	if replyCode2 != pb.AddTransactionReply_TR_Ok {
		t.Error("replyCode2 != Ok")
	}

	nodeTrans := nd.GetWaitingTrans()

	if len(nodeTrans) != 2 {
		t.Error("error: incorrect count of node waiting trans = " + strconv.FormatInt(int64(len(nodeTrans)), 10))
		return
	}

	if !transaction.IsEqual(airdrop1.GetData(), nodeTrans[0].GetData(), true) {
		t.Error("error: client and node version of one transfer is different")
	}

	if !transaction.IsEqual(airdrop2.GetData(), nodeTrans[1].GetData(), true) {
		t.Error("error: client and node version of one transfer is different")
	}

	block := nd.CreateBlock(nodeTrans)
	var difficulty uint64 = 2
	block.SetDifficulty(difficulty)
	for _, t := range nodeTrans {
		block.AddTransaction(t)
	}

	stopMining := make(chan bool)
	block.Mining(nd.GetPublicKey(), stopMining)
	//blockHash := cr.Keccak256(append(block.GetHash(), block.GetNonce()...))
	//t.Log("Block hash: ")
	//t.Log(blockHash)

	b := createBalance()

	if err := block.IsValid(nd.GetChain(), b); err != nil {
		t.Error(err)
	}

}
