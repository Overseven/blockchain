package transfer_test

import (
	"bytes"
	"github.com/overseven/try-network/transaction"
	"github.com/overseven/try-network/transaction/transfer"
	"github.com/overseven/try-network/utility"
	"testing"
)

func TestNewTransfer(t *testing.T) {
	senderPrivKey, _, err := utility.GenerateWallet()
	if err != nil {
		t.Error(err)
	}
	// pr1Bytes := utility.PrivToPubKey(pr1)

	_, receiverPubKey, err := utility.GenerateWallet()
	if err != nil {
		t.Error(err)
	}

	_, nodePubKey, err := utility.GenerateWallet()
	if err != nil {
		t.Error(err)
	}

	value := transaction.Balance(65.32)
	fee := transaction.Balance(42.2222)
	message := "test"

	tr1, err := transfer.NewTransfer(receiverPubKey, value, fee, message)
	if err != nil {
		t.Error(err)
	}

	tr1 = tr1.SetNode(nodePubKey).(*transfer.Transfer)
	if err != nil {
		t.Error(err)
	}

	err = tr1.Sign(senderPrivKey, 0)
	if err != nil {
		t.Error(err)
	}

	// fmt.Println(tr1.String())

	err = tr1.Verify()
	if err != nil {
		t.Error(err)
	}
}

func TestByteConversation(t *testing.T) {
	senderPrivKey, _, err := utility.GenerateWallet()
	if err != nil {
		t.Error(err)
	}
	// pr1Bytes := utility.PrivToPubKey(pr1)

	_, receiverPubKey, err := utility.GenerateWallet()
	if err != nil {
		t.Error(err)
	}

	_, nodePubKey, err := utility.GenerateWallet()
	if err != nil {
		t.Error(err)
	}

	value := transaction.Balance(65.32)
	fee := transaction.Balance(42.2222)
	message := "test"

	tr1, err := transfer.NewTransfer(receiverPubKey, value, fee, message)
	if err != nil {
		t.Error(err)
	}

	tr1 = tr1.SetNode(nodePubKey).(*transfer.Transfer)
	if err != nil {
		t.Error(err)
	}

	err = tr1.Sign(senderPrivKey, 0)
	if err != nil {
		t.Error(err)
	}

	// fmt.Println(tr1.String())

	err = tr1.Verify()
	if err != nil {
		t.Error(err)
	}

	b, err := tr1.Bytes()
	if err != nil {
		t.Error(err)
	}

	aReverse, err := transfer.FromBytes(b)
	if err != nil {
		t.Error(err)
	}

	//s1, err := tr1.String()
	//if err != nil {
	//	t.Error(err)
	//}
	//
	//s2, err := aReverse.String()
	//if err != nil {
	//	t.Error(err)
	//}

	//fmt.Println("orig Airdrop:", s1)
	//fmt.Println("Reverse Airdrop:", s2)
	h1, err := tr1.Hash(map[transaction.TransFlag]bool{})
	if err != nil {
		t.Error(err)
	}
	h2, err := aReverse.Hash(map[transaction.TransFlag]bool{})
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(h1, h2) {
		t.Error("incorrect hash")
	}
}
