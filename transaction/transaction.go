package transaction

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/overseven/blockchain/blockchain"
	"time"

	cr "github.com/ethereum/go-ethereum/crypto"
)

// TODO: calc pubkey len and set const size for all

type Transaction struct {
	Pubkey    []byte
	Receiver  []byte
	Message   string
	Timestamp time.Time
	Pay       float64
	Fee       float64
	Sign      []byte
}

func (tr *Transaction)GetHash() []byte {
	temp := append(tr.Pubkey, tr.Receiver...)
	temp = append(temp, tr.Message...)
	temp = append(temp, fmt.Sprintf("%.4f", tr.Pay)...)
	temp = append(temp, fmt.Sprintf("%.4f", tr.Fee)...)
	temp = append(temp, tr.Timestamp.String()...)
	return cr.Keccak256(temp)
}

/*
	Use it only if the sender has no more than one transaction in the block
 */
func (tr *Transaction)Verify() bool {
	hash := tr.GetHash()
	if !cr.VerifySignature(tr.Pubkey, hash, tr.Sign[0:64]){
		return false
	}

	wallet, err := blockchain.WalletInfo(tr.Pubkey)
	if err != nil {
		return false
	}
	if wallet.CurrentBalance < (tr.Pay + tr.Fee) {
		return false
	}

	return true
}

func (tr *Transaction)IsEqual(tr2 *Transaction) bool {
	if !bytes.Equal(tr.Pubkey, tr2.Pubkey){
		return false
	}
	if !bytes.Equal(tr.Receiver, tr2.Receiver){
		return false
	}
	if !bytes.Equal(tr.Sign, tr2.Sign){
		return false
	}
	if tr.Message != tr2.Message{
		return false
	}
	if tr.Timestamp != tr2.Timestamp{
		return false
	}
	if tr.Pay != tr2.Pay{
		return false
	}
	if tr.Fee != tr2.Fee{
		return false
	}

	return true
}

func FromJSON(js []byte) Transaction {
	tr := Transaction{}
	err := json.Unmarshal(js, &tr)
	if err != nil {
		panic(err)
	}
	//spew.Dump(tr)
	return tr
}
