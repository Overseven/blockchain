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

func GetHash(tr *Transaction) []byte {
	temp := append(tr.Pubkey, tr.Receiver...)
	temp = append(temp, tr.Message...)
	temp = append(temp, fmt.Sprintf("%.4f", tr.Pay)...)
	temp = append(temp, fmt.Sprintf("%.4f", tr.Fee)...)
	temp = append(temp, tr.Timestamp.String()...)
	return cr.Keccak256(temp)
}

func Verify(tr *Transaction) bool {
	hash := GetHash(tr)
	valid := cr.VerifySignature(tr.Pubkey, hash, tr.Sign[0:64])
	// need to insert additional check
	return valid
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
