package voting_init_test

import (
	"bytes"
	"github.com/overseven/try-network/transaction"
	"github.com/overseven/try-network/transaction/voting_init"
	"github.com/overseven/try-network/utility"
	"testing"
)

func TestNewVoting(t *testing.T) {
	senderPrivKey, _, err := utility.GenerateWallet()
	if err != nil {
		t.Error(err)
	}

	_, nodePubKey, err := utility.GenerateWallet()
	if err != nil {
		t.Error(err)
	}

	var votingId uint64 = 115

	var parameter uint16 = 17
	value := "125.34"
	fee := 42.2222

	v, err := voting_init.NewVotingInit(votingId, parameter, value, fee)
	if err != nil {
		t.Error(err)
	}

	v = v.SetNode(nodePubKey).(*voting_init.VotingInit)
	if err != nil {
		t.Error(err)
	}

	err = v.Sign(senderPrivKey, 0)
	if err != nil {
		t.Error(err)
	}

	// fmt.Println(tr1.String())

	err = v.Verify()
	if err != nil {
		t.Error(err)
	}
}

func TestByteConversation(t *testing.T) {
	senderPrivKey, _, err := utility.GenerateWallet()
	if err != nil {
		t.Error(err)
	}

	_, nodePubKey, err := utility.GenerateWallet()
	if err != nil {
		t.Error(err)
	}

	var votingId uint64 = 115
	var parameter uint16 = 1026
	value := "1251"
	fee := 42.2222

	v, err := voting_init.NewVotingInit(votingId, parameter, value, fee)
	if err != nil {
		t.Error(err)
	}

	v = v.SetNode(nodePubKey).(*voting_init.VotingInit)
	if err != nil {
		t.Error(err)
	}

	err = v.Sign(senderPrivKey, 0)
	if err != nil {
		t.Error(err)
	}

	// fmt.Println(tr1.String())

	err = v.Verify()
	if err != nil {
		t.Error(err)
	}

	b, err := v.Bytes()
	if err != nil {
		t.Error(err)
	}

	vReverse, err := voting_init.FromBytes(b)
	if err != nil {
		t.Error(err)
	}

	//s1, err := v.String()
	//if err != nil {
	//    t.Error(err)
	//}
	//
	//s2, err := vReverse.String()
	//if err != nil {
	//    t.Error(err)
	//}

	//fmt.Println("orig VotintInit:", s1)
	//fmt.Println("Reverse VotintInit:", s2)

	h1, err := v.Hash(map[transaction.TransFlag]bool{})
	if err != nil {
		t.Error(err)
	}
	h2, err := vReverse.Hash(map[transaction.TransFlag]bool{})
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(h1, h2) {
		t.Error("incorrect hash")
	}
}
