package transaction

import (
	"encoding/json"
	"fmt"
	"time"

	cr "github.com/ethereum/go-ethereum/crypto"
)

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

func (tr *Transaction)Verify() bool {
	hash := tr.GetHash()
	valid := cr.VerifySignature(tr.Pubkey, hash, tr.Sign[0:64])
	// TODO: need to insert additional check
	return valid
}

func (tr *Transaction)IsEqual(tr2 *Transaction) bool {
	//TODO: finish!
	return false
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
