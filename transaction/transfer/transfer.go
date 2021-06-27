package transfer

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"github.com/overseven/try-network/transaction"

	cr "github.com/ethereum/go-ethereum/crypto"
	// "github.com/overseven/try-network/chain"
	"github.com/overseven/try-network/utility"
)

type Transfer struct {
	Sender    []byte // compressed public key
	Receiver  []byte // compressed public key
	Pay       float64
	Fee       float64
	Message   string
	Timestamp time.Time
	Node      []byte
	Signature []byte
}

func (t *Transfer) String() (string, error) {
	tmp, err := json.Marshal(t)
	if err != nil {
		return "", err
	}

	return string(tmp), nil
}

func (t *Transfer) Bytes() ([]byte, error) {
	var res []byte

	res = append(res, byte(transaction.TypeTransfer))

	if len(t.Sender) != transaction.ByteLenPubKey {
		return nil, errors.New("incorrect sender field size")
	}
	res = append(res, t.Sender...)

	if len(t.Receiver) != transaction.ByteLenPubKey {
		return nil, errors.New("incorrect receiver field size")
	}
	res = append(res, t.Receiver...)

	res = append(res, utility.Float64Bytes(t.Pay)...)
	res = append(res, utility.Float64Bytes(t.Fee)...)
	res = append(res, utility.StringToBytes(t.Message)...)

	timestamp, err := utility.TimestampToBytes(t.Timestamp)
	if err != nil {
		return nil, err
	}

	res = append(res, timestamp...)

	if len(t.Node) != transaction.ByteLenPubKey {
		return nil, errors.New("incorrect node field size")
	}
	res = append(res, t.Node...)

	if len(t.Signature) != transaction.ByteLenSign {
		return nil, errors.New("incorrect sign field size")
	}

	res = append(res, t.Signature...)

	return res, nil
}

func FromBytes(b []byte) (*Transfer, error) {
	if len(b) < 64 { // TODO: define min size
		return nil, errors.New("incorrect input data len")
	}
	t := new(Transfer)

	idx := int64(0)
	typeTr := transaction.Type(b[idx])
	if typeTr != transaction.TypeTransfer {
		return nil, errors.New("incorrect transaction type")
	}
	idx += 1
	t.Sender = b[idx : idx+transaction.ByteLenPubKey]
	idx += transaction.ByteLenPubKey
	t.Receiver = b[idx : idx+transaction.ByteLenPubKey]
	idx += transaction.ByteLenPubKey
	t.Pay = utility.Float64FromBytes(b[idx : idx+8])
	idx += 8
	t.Fee = utility.Float64FromBytes(b[idx : idx+8])
	idx += 8
	message, messageLen, err := utility.StringFromBytes(b[idx:])
	if err != nil {
		return nil, err
	}
	t.Message = message
	idx += int64(messageLen)

	timestamp, timestampLen, err := utility.TimestampFromBytes(b[idx:])
	if err != nil {
		return nil, err
	}
	t.Timestamp = timestamp
	idx += int64(timestampLen)

	t.Node = b[idx : idx+transaction.ByteLenPubKey]
	idx += transaction.ByteLenPubKey
	t.Signature = b[idx : idx+transaction.ByteLenSign]
	idx += transaction.ByteLenSign
	return t, nil
}

func (t *Transfer) Hash(flags map[transaction.TransFlag]bool) ([]byte, error) {
	if flags == nil {
		return nil, errors.New("empty flags")
	}
	var temp []byte
	temp = append(temp, t.Sender...)
	temp = append(temp, t.Receiver...)
	temp = append(temp, t.Message...)
	temp = append(temp, strconv.FormatFloat(t.Pay, 'e', 8, 64)...)
	temp = append(temp, strconv.FormatFloat(t.Fee, 'e', 8, 64)...)
	flagTimestamp, ok := flags[transaction.FlagTimestamp]
	if !ok || (ok && flagTimestamp) {
		temp = append(temp, t.Timestamp.Format(utility.TimestampFormat)...)
	}

	flagNode, ok := flags[transaction.FlagNode]
	if !ok || (ok && flagNode) {
		temp = append(temp, t.Node...)
	}
	return cr.Keccak256(temp), nil
}

func (t *Transfer) Verify() error {
	hash, err := t.Hash(map[transaction.TransFlag]bool{})
	if err != nil {
		return err
	}
	if len(t.Signature) < 64 {
		return errors.New("incorrect signature len: " + strconv.FormatInt(int64(len(t.Signature)), 10))
	}
	if !cr.VerifySignature(t.Sender, hash, t.Signature[0:64]) {
		return errors.New("incorrect signature")
	}

	// senderWallet, err := balance.Info(t.Sender)
	// if err != nil {
	// 	return err
	// }
	// if senderWallet.CurrentBalance < (t.Pay + t.Fee) {
	// 	return errors.New("sender wallet not enough tokens")
	// }

	return nil
}

func NewTransfer(rcvrPubKey []byte, value, fee float64, message string) (*Transfer, error) {
	tr := Transfer{
		Receiver: rcvrPubKey,
		Pay:      value,
		Fee:      fee,
		Message:  message, // TODO: add fix size message
	}

	tr.Timestamp = utility.NewTimestamp()

	//hashed, err := tr.Hash(map[transaction.TransFlag]bool{})
	//if err != nil {
	//	return nil, err
	//}
	//
	//sign, err := cr.Sign(hashed, sndrPrivKey)
	//if err != nil {
	//	return nil, err
	//}
	//
	//tr.Signature = sign

	return &tr, nil
}

func (t *Transfer) IsEqual(tr transaction.Transaction, flags map[transaction.TransFlag]bool) bool {
	if flags == nil {
		return false
	}

	switch t2 := tr.(type) {
	case *Transfer:
		if !bytes.Equal(t.Sender, t2.Sender) {
			return false
		}

		if !bytes.Equal(t.Receiver, t2.Receiver) {
			return false
		}

		flagTimestamp, ok := flags[transaction.FlagTimestamp]
		if !ok || (ok && flagTimestamp) {
			if t.Timestamp != t2.Timestamp {
				return false
			}
		}

		flagNode, ok := flags[transaction.FlagNode]
		if !ok || (ok && flagNode) {
			if !bytes.Equal(t.Node, t2.Node) {
				return false
			}
		}

		if err := t.Verify(); err != nil {
			return false
		}
		if err := t2.Verify(); err != nil {
			return false
		}

		if t.Message != t2.Message {
			return false
		}
		if t.Pay != t2.Pay {
			return false
		}
		if t.Fee != t2.Fee {
			return false
		}
		return true
	default:
		return false
	}
}

func Copy(t *Transfer) *Transfer {
	res := new(Transfer)
	res.Sender = t.Sender
	res.Receiver = t.Receiver
	res.Pay = t.Pay
	res.Fee = t.Fee
	res.Message = t.Message
	res.Timestamp = t.Timestamp
	res.Node = t.Node
	res.Signature = t.Signature
	return res
}

func (t *Transfer) SetNode(nodePubKey []byte) transaction.Transaction {
	res := Copy(t)
	res.Node = nodePubKey
	return res
}

func (t *Transfer) Sign(privKey *ecdsa.PrivateKey) error {
	senderPubKey := utility.PrivToPubKey(privKey)
	t.Sender = senderPubKey

	hashed, err := t.Hash(map[transaction.TransFlag]bool{})
	if err != nil {
		return err
	}

	sign, err := cr.Sign(hashed, privKey)
	if err != nil {
		return err
	}

	t.Signature = sign
	return nil
}
