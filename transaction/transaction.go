package transaction

import (
	"bytes"
	"strconv"

	cr "github.com/ethereum/go-ethereum/crypto"
	"github.com/overseven/blockchain/interfaces"
)

// TODO: calc pubkey len and set const size for all

type Type int64

const (
	TypeAirdrop = iota
	TypeTransfer
)

func IsEqual(t1, t2 *interfaces.Data) bool {
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

func GetHash(t *interfaces.Data) []byte {
	temp := []byte(strconv.FormatInt(int64(t.Type), 10))
	temp = append(temp, t.Pubkey...)
	temp = append(temp, t.Receiver...)
	temp = append(temp, t.Message...)
	temp = append(temp, strconv.FormatFloat(t.Pay, 'e', 8, 64)...)
	temp = append(temp, strconv.FormatFloat(t.Fee, 'e', 8, 64)...)
	temp = append(temp, t.Timestamp.String()...)
	return cr.Keccak256(temp)
}
