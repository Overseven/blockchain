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
	airdropModeratorConfigFile = "..\\wallet.cfg"
)

func createClient() {

}

func TestNodeClientCommunication(t *testing.T) {
	clientPrKey, _, err := generateWallet()
	if err != nil {
		t.Error(err)
	}

	cl := client.Client{}
	cl.Init()
	cl.SetPrivateKey(clientPrKey)

	nd := node.Node{}
	nd.Init()

	_, rcvrPubKey, err := generateWallet()
	if err != nil {
		t.Error(err)
	}

	var (
		clientPort uint32 = 505
		nodePort   uint32 = 506
	)

	value := 152.313
	fee := 0.0004
	message := "network test"

	trans, err := transaction.NewTransfer(clientPrKey, rcvrPubKey, value, fee, message)
	if err != nil {
		t.Error(err)
	}

	cl.SetPort(clientPort)
	nd.SetPort(nodePort)

	stop := make(chan bool)
	err = nd.StartListening(stop)
	if err != nil {
		t.Error(err)
	}

	replyCode, err := cl.SendTransaction(trans, "localhost:"+strconv.FormatUint(uint64(nodePort), 10))
	if err != nil {
		t.Error(err)
	}

	if replyCode != pb.AddTransactionReply_TR_Ok {
		t.Error("replyCode != Ok")
	}
}
