package transaction_test

import (
	"bytes"
	"strconv"
	"testing"

	cr "github.com/ethereum/go-ethereum/crypto"
	"github.com/overseven/try-network/transaction"
	"github.com/overseven/try-network/transaction/transfer"
	"github.com/overseven/try-network/utility"
)

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

	tr1, err := transfer.NewTransfer(pr1, pub2, value, fee, message)
	if err != nil {
		t.Error(err)
	}

	temp := []byte(strconv.FormatInt(int64(transaction.TypeTransfer), 10))
	temp = append(temp, pr1Bytes...)
	temp = append(temp, pub2...)
	temp = append(temp, message...)
	temp = append(temp, strconv.FormatFloat(value, 'e', 8, 64)...)
	temp = append(temp, strconv.FormatFloat(fee, 'e', 8, 64)...)
	temp = append(temp, tr1.Timestamp.Format(utility.TimestampFormat)...)
	hash1 := cr.Keccak256(temp)

	hash2, err := tr1.Hash(map[transaction.TransFlag]bool{})
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(hash1, hash2) {
		t.Error("err: hash1 and hash2 are not equal")
	}
}
