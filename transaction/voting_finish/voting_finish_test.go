package voting_finish_test

import (
	"bytes"
	"github.com/overseven/try-network/transaction"
	vf "github.com/overseven/try-network/transaction/voting_finish"
	"github.com/overseven/try-network/utility"
	"testing"
)

func TestNewVotingFinish(t *testing.T) {
	senderPrivKey, _, err := utility.GenerateWallet()
	if err != nil {
		t.Error(err)
	}

	_, nodePubKey, err := utility.GenerateWallet()
	if err != nil {
		t.Error(err)
	}

	votingId := transaction.VotingId(115)
	fee := transaction.Balance(42.2222)

	v, err := vf.NewVotingFinish(votingId, fee)
	if err != nil {
		t.Error(err)
	}

	v = v.SetNode(nodePubKey).(*vf.VotingFinish)
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

	votingId := transaction.VotingId(115)
	fee := transaction.Balance(42.2222)

	v, err := vf.NewVotingFinish(votingId, fee)
	if err != nil {
		t.Error(err)
	}

	v = v.SetNode(nodePubKey).(*vf.VotingFinish)
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

	vReverse, err := vf.FromBytes(b)
	if err != nil {
		t.Error(err)
	}

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
