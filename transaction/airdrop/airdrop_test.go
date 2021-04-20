package airdrop_test

import (
	"testing"

	cr "github.com/ethereum/go-ethereum/crypto"
	"github.com/overseven/blockchain/balance"
	"github.com/overseven/blockchain/interfaces"
	"github.com/overseven/blockchain/transaction"
	"github.com/overseven/blockchain/transaction/transfer"
	"github.com/overseven/blockchain/utility"
	"github.com/overseven/blockchain/wallet"
)

const airdropModeratorConfigFile = "..\\wallet.cfg"

func TestAirdrop_GetData(t *testing.T) {
	data := interfaces.Data{}
	pr1, _, err := utility.GenerateWallet()
	if err != nil {
		t.Error(err)
	}
	pr1Bytes := utility.PrivToPubKey(pr1)

	_, pub2, err := utility.GenerateWallet()
	if err != nil {
		t.Error(err)
	}
	value := 65.32
	fee := 42.2222
	message := "test"
	timestamp := utility.NewTimestamp()

	data.Type = transaction.TypeTransfer
	data.Sender = pr1Bytes
	data.Receiver = pub2
	data.Pay = value
	data.Fee = fee
	data.Message = message
	data.Timestamp = timestamp
	data.Sign = []byte("eeee sign!")

	tr1 := transfer.Transfer{Data: data}

	if !transaction.IsEqual(&data, tr1.GetData(), true) {
		t.Error("err: data and tr1.data are not equal!")
	}

	data.Message = "Press F"

	if transaction.IsEqual(&data, tr1.GetData(), true) {
		t.Error("err: modified data and tr1.data are false equal!")
	}
}
func TestNewAirdrop(t *testing.T) {

}

func TestAirdropVerify(t *testing.T) {
	var usersBalance interfaces.Balancer = &balance.Balance{}
	usersBalance.Init()

	_, receiver1, err := utility.GenerateWallet()
	if err != nil {
		t.Error(err)
	}

	airdropPubKey, airdropPrKey, err := config.LoadFromFile(airdropModeratorConfigFile)
	if err != nil {
		t.Error(err)
	}

	transaction.AirDropModeratorPubKey = airdropPubKey

	airPrKey, err := cr.ToECDSA(airdropPrKey[:32])
	if err != nil {
		t.Error(err)
	}

	airdrop, err := transaction.NewAirdrop(receiver1, airPrKey, 110.1, 11.1)
	if err != nil {
		t.Error(err)
	}

	if err := airdrop.Verify(usersBalance); err != nil {
		t.Error(err)
	}

}
