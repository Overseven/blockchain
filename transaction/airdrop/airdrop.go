package airdrop

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"strconv"
	"time"

	// "github.com/overseven/blockchain/chain/ichain"

	"github.com/overseven/blockchain/transaction"
	"github.com/overseven/blockchain/utility"

	cr "github.com/ethereum/go-ethereum/crypto"
)

var (
	AirDropModeratorPubKey []byte // decompressed public key
)

type Airdrop struct {
	Receiver  []byte // compressed public key
	Pay       float64
	Fee       float64
	Message   string
	Timestamp time.Time
	Node      []byte
	Signature []byte
}

func (a *Airdrop) IsEqual(t transaction.Transaction, flags map[transaction.TransFlag]bool) bool {
	if flags == nil {
		return false
	}

	switch a2 := t.(type) {
	case *Airdrop:
		if !bytes.Equal(a.Receiver, a2.Receiver) {
			return false
		}

		flagTimestamp, ok := flags[transaction.FlagTimestamp]
		if !ok || (ok && flagTimestamp) {
			if a.Timestamp != a2.Timestamp {
				return false
			}
		}

		flagNode, ok := flags[transaction.FlagNode]
		if !ok || (ok && flagNode) {
			if !bytes.Equal(a.Node, a2.Node) {
				return false
			}
		}

		if err := a.Verify(); err != nil {
			return false
		}
		if err := a2.Verify(); err != nil {
			return false
		}

		if a.Message != a2.Message {
			return false
		}
		if a.Pay != a2.Pay {
			return false
		}
		if a.Fee != a2.Fee {
			return false
		}
		return true
	default:
		return false
	}
}

func (a *Airdrop) String() (string, error) {
	tmp, err := json.MarshalIndent(a, "", "")
	if err != nil {
		return "", err
	}

	return string(tmp), nil
}

func (a *Airdrop) Bytes() ([]byte, error) {
	var res []byte

	res = append(res, byte(transaction.TypeAirdrop))

	if len(a.Receiver) != transaction.ByteLenPubKey {
		return nil, errors.New("incorrect receiver field size")
	}
	res = append(res, a.Receiver...)

	res = append(res, utility.Float64Bytes(a.Pay)...)
	res = append(res, utility.Float64Bytes(a.Fee)...)
	message := utility.StringToBytes(a.Message) // TODO: FIX ME!
	//res = append(res, uint8(len(message)))
	res = append(res, message...)

	timestamp, err := utility.TimestampToBytes(a.Timestamp)
	if err != nil {
		return nil, err
	}

	res = append(res, uint8(len(timestamp)))
	res = append(res, timestamp...)

	if len(a.Signature) != transaction.ByteLenSign {
		return nil, errors.New("incorrect sign field size")
	}
	res = append(res, a.Node...)
	res = append(res, a.Signature...)

	return res, nil
}

func FromBytes(b []byte) (*Airdrop, error) {
	if len(b) < 64 { // TODO: define min size
		return nil, errors.New("incorrect input data len")
	}
	a := new(Airdrop)
	var err error
	idx := int64(0)
	typeTr := transaction.Type(b[idx])
	if typeTr != transaction.TypeAirdrop {
		return nil, errors.New("incorrect transaction type")
	}
	idx += 1

	a.Receiver = b[idx : idx+transaction.ByteLenPubKey]
	idx += transaction.ByteLenPubKey
	a.Pay = utility.Float64FromBytes(b[idx : idx+8])
	idx += 8
	a.Fee = utility.Float64FromBytes(b[idx : idx+8])
	idx += 8
	message, messageLen, err := utility.StringFromBytes(b[idx:])
	if err != nil {
		return nil, err
	}
	a.Message = message
	idx += int64(messageLen)

	timestamp, timestampLen, err := utility.TimestampFromBytes(b[idx:])
	if err != nil {
		return nil, err
	}
	a.Timestamp = timestamp
	idx += int64(timestampLen)

	a.Node = b[idx : idx+transaction.ByteLenPubKey]
	idx += transaction.ByteLenPubKey
	a.Signature = b[idx : idx+transaction.ByteLenSign]
	idx += transaction.ByteLenSign
	return a, nil
}

func (a *Airdrop) Hash(flags map[transaction.TransFlag]bool) ([]byte, error) {
	if flags == nil {
		return nil, errors.New("empty flags")
	}
	temp := []byte{}
	temp = append(temp, a.Receiver...)
	temp = append(temp, a.Message...)
	temp = append(temp, strconv.FormatFloat(a.Pay, 'e', 8, 64)...)
	temp = append(temp, strconv.FormatFloat(a.Fee, 'e', 8, 64)...)
	flagTimestamp, ok := flags[transaction.FlagTimestamp]
	if !ok || (ok && flagTimestamp) {
		temp = append(temp, a.Timestamp.Format(utility.TimestampFormat)...)
	}

	flagNode, ok := flags[transaction.FlagNode]
	if !ok || (ok && flagNode) {
		temp = append(temp, a.Node...)
	}

	return cr.Keccak256(temp), nil
}

func (a *Airdrop) Verify() error {

	if len(AirDropModeratorPubKey) == 0 {
		return errors.New("empty AirDrop moderator public key")
	}

	hash, err := a.Hash(map[transaction.TransFlag]bool{})
	if err != nil {
		return err
	}
	if !cr.VerifySignature(AirDropModeratorPubKey, hash, a.Signature[0:64]) {
		return errors.New("incorrect AirDrop moderator signature")
	}

	// if bytes.Compare(a.Sender, AirDropModeratorPubKey) != 0 {
	// 	return errors.New("incorrect AirDrop moderator public key")
	// }

	return nil
}

// NewAirdrop is sending value from unlimited admin wallet to user wallet
func NewAirdrop(receiver []byte, payment, fee float64, message string) (*Airdrop, error) {
	// TODO: add check below

	a := new(Airdrop)
	a.Receiver = receiver

	a.Pay = payment
	a.Fee = fee
	a.Message = message
	{
		t := time.Now()
		a.Timestamp = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.UTC)
	}

	return a, nil
}

func Copy(a *Airdrop) *Airdrop {
	res := new(Airdrop)
	res.Receiver = a.Receiver
	res.Pay = a.Pay
	res.Fee = a.Fee
	res.Message = a.Message
	res.Timestamp = a.Timestamp
	res.Node = a.Node
	return res
}

func (a *Airdrop) SetNode(nodePubKey []byte) transaction.Transaction {
	res := Copy(a)
	res.Node = nodePubKey
	return res
}

func (a *Airdrop) Sign(privKey *ecdsa.PrivateKey) error {
	hashed, err := a.Hash(map[transaction.TransFlag]bool{})
	if err != nil {
		return err
	}

	sign, err := cr.Sign(hashed, privKey)
	if err != nil {
		return err
	}

	a.Signature = sign
	return nil
}
