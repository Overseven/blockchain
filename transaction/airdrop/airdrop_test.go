package airdrop_test

import (
	"bytes"
	"errors"
	"fmt"
	"testing"

	"github.com/overseven/blockchain/transaction"
	"github.com/overseven/blockchain/transaction/airdrop"
	"github.com/overseven/blockchain/utility"
	"github.com/overseven/blockchain/utility/config"
)

const airdropModeratorConfigFile = "D:\\go\\src\\github.com\\overseven\\blockchain\\wallet.cfg"

func TestHash(t *testing.T) {
	// TODO: finish
}

func TestNewAirdrop(t *testing.T) {
	params, err := config.LoadFromFile(airdropModeratorConfigFile)
	if err != nil {
		t.Error(err)
	}

	airdrop.AirDropModeratorPubKey = params.PubKey

	_, nodePubKey, err := utility.GenerateWallet()
	if err != nil {
		t.Error(err)
	}
	// pr1Bytes := utility.PrivToPubKey(pr1)

	_, pub2, err := utility.GenerateWallet()
	if err != nil {
		t.Error(err)
	}
	value := 65.32
	fee := 42.2222
	message := "test"

	tr1, err := airdrop.NewAirdrop(pub2, value, fee, message)
	if err != nil {
		t.Error(err)
	}

	tr1 = tr1.SetNode(nodePubKey).(*airdrop.Airdrop)
	if err != nil {
		t.Error(err)
	}

	tr1.Sign(params.PrivKey)
	if err != nil {
		t.Error(err)
	}

	// fmt.Println(tr1.String())

	err = tr1.Verify()
	if err != nil {
		t.Error(err)
	}
}

func TestNewAirdrop2(t *testing.T) {
	params, err := config.LoadFromFile(airdropModeratorConfigFile)
	if err != nil {
		t.Error(err)
	}

	airdrop.AirDropModeratorPubKey = params.PubKey

	_, nodePubKey, err := utility.GenerateWallet()
	if err != nil {
		t.Error(err)
	}
	// pr1Bytes := utility.PrivToPubKey(pr1)

	_, pub2, err := utility.GenerateWallet()
	if err != nil {
		t.Error(err)
	}
	value := 65235.32
	fee := 421.22422
	message := "test 2"

	tr1, err := airdrop.NewAirdrop(pub2, value, fee, message)
	if err != nil {
		t.Error(err)
	}

	tr1 = tr1.SetNode(nodePubKey).(*airdrop.Airdrop)
	if err != nil {
		t.Error(err)
	}

	tr1.Sign(params.PrivKey)
	if err != nil {
		t.Error(err)
	}

	err = tr1.Verify()
	if err != nil {
		t.Error(err)
	}

}

func TestByteConversation(t *testing.T) {
	params, err := config.LoadFromFile(airdropModeratorConfigFile)
	if err != nil {
		t.Error(err)
	}

	// adminPK, err := cr.DecompressPubkey(params.PubKey)

	airdrop.AirDropModeratorPubKey = params.PubKey

	_, nodePubKey, err := utility.GenerateWallet()
	if err != nil {
		t.Error(err)
	}
	// pr1Bytes := utility.PrivToPubKey(pr1)

	_, pub2, err := utility.GenerateWallet()
	if err != nil {
		t.Error(err)
	}
	value := 65235.32
	fee := 421.22422
	message := "test 2"

	tr1, err := airdrop.NewAirdrop(pub2, value, fee, message)
	if err != nil {
		t.Error(err)
	}

	tr1 = tr1.SetNode(nodePubKey).(*airdrop.Airdrop)
	if err != nil {
		t.Error(err)
	}

	tr1.Sign(params.PrivKey)
	if err != nil {
		t.Error(err)
	}

	err = tr1.Verify()
	if err != nil {
		t.Error(err)
	}

	b, err := tr1.Bytes()
	if err != nil {
		t.Error(err)
	}

	aReverse, err := airdrop.FromBytes(b)
	if err != nil {
		t.Error(err)
	}

	s1, err := tr1.String()
	if err != nil {
		t.Error(err)
	}

	s2, err := aReverse.String()
	if err != nil {
		t.Error(err)
	}

	fmt.Println("orig Airdrop:", s1)
	fmt.Println("Reverse Airdrop:", s2)
	h1, err := tr1.Hash(map[transaction.TransFlag]bool{})
	if err != nil {
		t.Error(err)
	}
	h2, err := aReverse.Hash(map[transaction.TransFlag]bool{})
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(h1, h2) {
		t.Error(errors.New("incorrect hash"))
	}
}
