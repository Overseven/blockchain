package transaction

import (
	"bytes"
	"strconv"
	"time"

	cr "github.com/ethereum/go-ethereum/crypto"
)

// TODO: calc pubkey len and set const size for all

type Type int64

const (
	Transfer = iota
	Airdrop
)

type Transaction interface {
	Verify() bool
	SetData(*Data)
	// GetHash() []byte

	// IsEqual(tr2 *Transaction) bool

}

type Data struct {
	Type      Type
	Pubkey    []byte
	Receiver  []byte
	Message   string
	Timestamp time.Time
	Pay       float64
	Fee       float64
	Sign      []byte
}

func IsEqual(t1, t2 *Data) bool {
	if !bytes.Equal(t1.Pubkey, t2.Pubkey) {
		return false
	}
	if !bytes.Equal(t1.Receiver, t2.Receiver) {
		return false
	}
	if !bytes.Equal(t1.Sign, t2.Sign) {
		return false
	}
	if t1.Message != t2.Message {
		return false
	}
	if t1.Timestamp != t2.Timestamp {
		return false
	}
	if t1.Pay != t2.Pay {
		return false
	}
	if t1.Fee != t2.Fee {
		return false
	}

	return true
}

func GetHash(t *Data) []byte {
	temp := []byte(strconv.FormatInt(int64(t.Type), 10))
	temp = append(temp, t.Pubkey...)
	temp = append(temp, t.Receiver...)
	temp = append(temp, t.Message...)
	temp = append(temp, strconv.FormatFloat(t.Pay, 'e', 8, 64)...)
	temp = append(temp, strconv.FormatFloat(t.Fee, 'e', 8, 64)...)
	temp = append(temp, t.Timestamp.String()...)
	return cr.Keccak256(temp)
}
