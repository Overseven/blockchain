package test

import (
	"strconv"
	"testing"

	"github.com/overseven/blockchain/client"
	"github.com/overseven/blockchain/node"
	pb "github.com/overseven/blockchain/protocol"
	"github.com/overseven/blockchain/transaction"
)

const (
	airdropModeratorConfigFile        = "..\\wallet.cfg"
	clientPort                 uint32 = 505
	nodePort                   uint32 = 506
)

func createClient() (*client.Client, error) {
	clientPrKey, _, err := generateWallet()
	if err != nil {
		return nil, err
	}

	cl := client.Client{}
	cl.Init()
	cl.SetPrivateKey(clientPrKey)
	cl.SetPort(clientPort)

	return &cl, nil
}

func createNode() (*node.Node, error) {
	nd := node.Node{}
	nd.Init()
	nd.SetPort(nodePort)

	return &nd, nil
}

func TestNodeClientCommunication(t *testing.T) {
	nd, err := createNode()
	if err != nil {
		t.Error(err)
	}

	cl, err := createClient()
	if err != nil {
		t.Error(err)
	}

	_, rcvrPubKey, err := generateWallet()
	if err != nil {
		t.Error(err)
	}

	value := 152.313
	fee := 0.0004
	message := "network test"

	trans, err := transaction.NewTransfer(cl.GetPrivateKey(), rcvrPubKey, value, fee, message)
	if err != nil {
		t.Error(err)
	}

	stop := make(chan bool)
	err = nd.StartListening(stop)
	if err != nil {
		t.Error(err)
	}
	defer func() { stop <- true }()

	replyCode, err := cl.SendTransaction(trans, "localhost:"+strconv.FormatUint(uint64(nodePort), 10))
	if err != nil {
		t.Error(err)
	}

	if replyCode != pb.AddTransactionReply_TR_Ok {
		t.Error("replyCode != Ok")
	}

	nodeTrans := nd.GetWaitingTrans()

	if len(nodeTrans) != 1 {
		t.Error("error: incorrect count of node waiting trans = " + strconv.FormatInt(int64(len(nodeTrans)), 10))
		return
	}
	ndTran := nodeTrans[0]
	if !transaction.IsEqual(trans.GetData(), ndTran.GetData()) {
		t.Error("error: client and node version of one transfer is different")
	}
}

func TestAirdrop(t *testing.T) {

}
