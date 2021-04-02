package transaction_test

import (
	"bytes"
	cr "github.com/ethereum/go-ethereum/crypto"
	"github.com/overseven/blockchain/transaction"
	"github.com/overseven/blockchain/utility"
	"strconv"
	"testing"
)

func TestIsEqual(t *testing.T) {
	pr1, pub1, err := utility.GenerateWallet()
	if err != nil {
		t.Error(err)
	}

	pr2, pub2, err := utility.GenerateWallet()
	if err != nil {
		t.Error(err)
	}
	value := 65.32
	fee := 42.2222
	message := "test"
	tr1, err := transaction.NewTransfer(pr1, pub2, value, fee, message)
	if err != nil {
		t.Error(err)
	}

	tr2, err := transaction.NewTransfer(pr1, pub2, value, fee, message)
	if err != nil {
		t.Error(err)
	}

	equal := transaction.IsEqual(tr1.GetData(), tr2.GetData(), false)
	if !equal {
		t.Error("error with compare tr1 and tr2")
	}

	tr3, err := transaction.NewTransfer(pr2, pub2, value, fee, message)
	if err != nil {
		t.Error(err)
	}

	equal = transaction.IsEqual(tr1.GetData(), tr3.GetData(), false)
	if equal {
		t.Error("error with compare tr1 and tr3")
	}

	tr4, err := transaction.NewTransfer(pr1, pub1, value, fee, message)
	if err != nil {
		t.Error(err)
	}

	equal = transaction.IsEqual(tr1.GetData(), tr4.GetData(), false)
	if equal {
		t.Error("error with compare tr1 and tr4")
	}

	tr5, err := transaction.NewTransfer(pr1, pub2, value+0.0001, fee, message)
	if err != nil {
		t.Error(err)
	}

	equal = transaction.IsEqual(tr1.GetData(), tr5.GetData(), false)
	if equal {
		t.Error("error with compare tr1 and tr5")
	}

	tr6, err := transaction.NewTransfer(pr1, pub2, value, fee+0.0001, message)
	if err != nil {
		t.Error(err)
	}

	equal = transaction.IsEqual(tr1.GetData(), tr6.GetData(), false)
	if equal {
		t.Error("error with compare tr1 and tr6")
	}

	tr7, err := transaction.NewTransfer(pr1, pub2, value, fee, message+"a")
	if err != nil {
		t.Error(err)
	}

	equal = transaction.IsEqual(tr1.GetData(), tr7.GetData(), false)
	if equal {
		t.Error("error with compare tr1 and tr7")
	}

	tr2.Data = *tr1.GetData()
	tr1.Data.Timestamp = utility.NewTimestamp().Add(1)
	if err != nil {
		t.Error(err)
	}

	equal = transaction.IsEqual(tr1.GetData(), tr2.GetData(), true)
	if equal {
		t.Error("error with compare tr1 and tr2 with diff timestamp")
	}

	tr1.Data = *tr2.GetData()
	tr1.Data.Sign = []byte("Siign!")

	equal = transaction.IsEqual(tr1.GetData(), tr2.GetData(), true)
	if equal {
		t.Error("error with compare tr1 and tr2 with diff sign")
	}
}

func TestGetHash(t *testing.T) {
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

	tr1 := transaction.Transfer{}
	data := tr1.GetData()
	data.Type = transaction.TypeTransfer
	data.Sender = pr1Bytes
	data.Receiver = pub2
	data.Pay = value
	data.Fee = fee
	data.Message = message
	data.Timestamp = timestamp

	temp := []byte(strconv.FormatInt(int64(transaction.TypeTransfer), 10))
	temp = append(temp, pr1Bytes...)
	temp = append(temp, pub2...)
	temp = append(temp, message...)
	temp = append(temp, strconv.FormatFloat(value, 'e', 8, 64)...)
	temp = append(temp, strconv.FormatFloat(fee, 'e', 8, 64)...)
	temp = append(temp, timestamp.String()...)
	hash1 := cr.Keccak256(temp)

	hash2 := transaction.GetHash(tr1.GetData())

	if !bytes.Equal(hash1, hash2) {
		t.Error("err: hash1 and hash2 are not equal")
	}
}