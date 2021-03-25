package test

import (
	"testing"

	cr "github.com/ethereum/go-ethereum/crypto"
	"github.com/overseven/blockchain/balance"
	"github.com/overseven/blockchain/interfaces"
	"github.com/overseven/blockchain/transaction"
	"github.com/overseven/blockchain/wallet"
)

func TestAirdropVerify(t *testing.T) {
	var usersBalance interfaces.Balancer = &balance.Balance{}
	usersBalance.Init()

	_, receiver1, err := generateWallet(0.0, usersBalance)
	if err != nil {
		t.Error(err)
	}

	airdropPubKey, airdropPrKey, err := wallet.LoadFromFile(airdropModeratorConfigFile)
	if err != nil {
		t.Error(err)
	}

	transaction.AirDropModeratorPubKey = airdropPubKey

	airPrKey, err := cr.ToECDSA(airdropPrKey[:32])
	if err != nil {
		t.Error(err)
	}

	airdrop, err := transaction.NewAirdrop(receiver1, airPrKey, 110.1, 11.1, usersBalance)
	if err != nil {
		t.Error(err)
	}

	if err := airdrop.Verify(usersBalance); err != nil {
		t.Error(err)
	}

}
