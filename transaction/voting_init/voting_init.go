package voting_init

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	cr "github.com/ethereum/go-ethereum/crypto"
	"strconv"
	"time"

	"github.com/overseven/try-network/transaction"
	"github.com/overseven/try-network/utility"
)

type VotingInit struct {
	Sender    []byte
	VotingId  uint64
	Parameter uint16
	Value     string
	Fee       float64
	Timestamp time.Time
	Node      []byte
	Signature []byte
}

func NewVoting(votingId uint64, parameter uint16, value string, fee float64) (*VotingInit, error) {
	tr := VotingInit{
		VotingId:  votingId,
		Parameter: parameter,
		Value:     value,
		Fee:       fee,
	}

	tr.Timestamp = utility.NewTimestamp()
	return &tr, nil
}

func (v *VotingInit) IsEqual(tr transaction.Transaction, flags map[transaction.TransFlag]bool) bool {
	if flags == nil {
		return false
	}

	switch v2 := tr.(type) {
	case *VotingInit:
		if !bytes.Equal(v.Sender, v2.Sender) {
			return false
		}

		if v.VotingId != v2.VotingId {
			return false
		}

		if v.Parameter != v2.Parameter {
			return false
		}

		if v.Value != v2.Value {
			return false
		}

		flagTimestamp, ok := flags[transaction.FlagTimestamp]
		if !ok || (ok && flagTimestamp) {
			if v.Timestamp != v2.Timestamp {
				return false
			}
		}

		flagNode, ok := flags[transaction.FlagNode]
		if !ok || (ok && flagNode) {
			if !bytes.Equal(v.Node, v2.Node) {
				return false
			}
		}

		if err := v.Verify(); err != nil {
			return false
		}
		if err := v2.Verify(); err != nil {
			return false
		}

		if v.Fee != v2.Fee {
			return false
		}
		return true
	default:
		return false
	}
}

func (v *VotingInit) String() (string, error) {
	tmp, err := json.Marshal(v)
	if err != nil {
		return "", err
	}

	return string(tmp), nil
}

func (v *VotingInit) Bytes() ([]byte, error) {
	var res []byte

	res = append(res, byte(transaction.TypeVotingInit))

	if len(v.Sender) != transaction.ByteLenPubKey {
		return nil, errors.New("incorrect sender field size")
	}
	res = append(res, v.Sender...)

	res = append(res, utility.UInt64Bytes(v.VotingId)...)
	res = append(res, utility.UInt16Bytes(v.Parameter)...)
	res = append(res, utility.StringToBytes(v.Value)...)
	res = append(res, utility.Float64Bytes(v.Fee)...)

	timestamp, err := utility.TimestampToBytes(v.Timestamp)
	if err != nil {
		return nil, err
	}
	res = append(res, timestamp...)

	if len(v.Node) != transaction.ByteLenPubKey {
		return nil, errors.New("incorrect node field size")
	}
	res = append(res, v.Node...)

	if len(v.Signature) != transaction.ByteLenSign {
		return nil, errors.New("incorrect sign field size")
	}

	res = append(res, v.Signature...)

	return res, nil
}

func FromBytes(b []byte) (*VotingInit, error) {
	if len(b) < 64 { // TODO: define min size
		return nil, errors.New("incorrect input data len")
	}
	v := new(VotingInit)
	idx := int64(0)
	typeTr := transaction.Type(b[idx])
	if typeTr != transaction.TypeVotingInit {
		return nil, errors.New("incorrect transaction type")
	}
	idx += 1
	v.Sender = b[idx : idx+transaction.ByteLenPubKey]
	idx += transaction.ByteLenPubKey

	v.VotingId = utility.UInt64FromBytes(b[idx : idx+8])
	idx += 8
	v.Parameter = utility.UInt16FromBytes(b[idx : idx+2])
	idx += 2
	value, valueLen, err := utility.StringFromBytes(b[idx:])
	if err != nil {
		return nil, err
	}
	v.Value = value
	idx += int64(valueLen)

	v.Fee = utility.Float64FromBytes(b[idx : idx+8])
	idx += 8

	timestamp, timestampLen, err := utility.TimestampFromBytes(b[idx:])
	if err != nil {
		return nil, err
	}
	v.Timestamp = timestamp
	idx += int64(timestampLen)

	v.Node = b[idx : idx+transaction.ByteLenPubKey]
	idx += transaction.ByteLenPubKey
	v.Signature = b[idx : idx+transaction.ByteLenSign]
	idx += transaction.ByteLenSign
	return v, nil
}

func (v *VotingInit) Hash(flags map[transaction.TransFlag]bool) ([]byte, error) {
	if flags == nil {
		return nil, errors.New("empty flags")
	}
	var temp []byte
	temp = append(temp, v.Sender...)
	temp = append(temp, utility.UInt64Bytes(v.VotingId)...)
	temp = append(temp, utility.UInt16Bytes(v.Parameter)...)
	temp = append(temp, v.Value...)
	temp = append(temp, strconv.FormatFloat(v.Fee, 'e', 8, 64)...)
	flagTimestamp, ok := flags[transaction.FlagTimestamp]
	if !ok || (ok && flagTimestamp) {
		temp = append(temp, v.Timestamp.Format(utility.TimestampFormat)...)
	}

	flagNode, ok := flags[transaction.FlagNode]
	if !ok || (ok && flagNode) {
		temp = append(temp, v.Node...)
	}
	return cr.Keccak256(temp), nil
}

func (v *VotingInit) Verify() error {
	hash, err := v.Hash(map[transaction.TransFlag]bool{})
	if err != nil {
		return err
	}
	if len(v.Signature) < 64 {
		return errors.New("incorrect signature len: " + strconv.FormatInt(int64(len(v.Signature)), 10))
	}
	if !cr.VerifySignature(v.Sender, hash, v.Signature[0:64]) {
		return errors.New("incorrect signature")
	}
	return nil
}

func Copy(v *VotingInit) *VotingInit {
	res := new(VotingInit)
	res.Sender = v.Sender
	res.VotingId = v.VotingId
	res.Parameter = v.Parameter
	res.Value = v.Value
	res.Fee = v.Fee
	res.Timestamp = v.Timestamp
	res.Node = v.Node
	res.Signature = v.Signature
	return res
}

func (v *VotingInit) SetNode(nodePubKey []byte) transaction.Transaction {
	res := Copy(v)
	res.Node = nodePubKey
	return res
}

func (v *VotingInit) Sign(privKey *ecdsa.PrivateKey) error {
	senderPubKey := utility.PrivToPubKey(privKey)
	v.Sender = senderPubKey

	hashed, err := v.Hash(map[transaction.TransFlag]bool{})
	if err != nil {
		return err
	}

	sign, err := cr.Sign(hashed, privKey)
	if err != nil {
		return err
	}

	v.Signature = sign
	return nil
}
