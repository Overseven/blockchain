package transfer

import (
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"github.com/overseven/blockchain/transaction"

	cr "github.com/ethereum/go-ethereum/crypto"
	// "github.com/overseven/blockchain/chain"
	"github.com/overseven/blockchain/utility"
)

type Transfer struct {
	Sender    []byte
	Receiver  []byte
	Pay       float64
	Fee       float64
	Message   string
	Timestamp time.Time
	Node      []byte
	Sign      []byte
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

	// TODO: timestamp
	if len(t.Node) != transaction.ByteLenPubKey {
		return nil, errors.New("incorrect node field size")
	}
	res = append(res, t.Node...)

	if len(t.Sign) != transaction.ByteLenSign {
		return nil, errors.New("incorrect sign field size")
	}

	res = append(res, t.Sign...)

	return res, nil
}

func (t *Transfer) FromBytes(b []byte) error {
	if len(b) < 64 { // TODO: define min size
		return errors.New("incorrect input data len")
	}

	var err error
	idx := 0
	t.Sender = b[idx : idx+transaction.ByteLenPubKey]
	idx += transaction.ByteLenPubKey
	t.Receiver = b[idx : idx+transaction.ByteLenPubKey]
	idx += transaction.ByteLenPubKey
	t.Pay = utility.Float64FromBytes(b[idx : idx+8])
	idx += 8
	t.Fee = utility.Float64FromBytes(b[idx : idx+8])
	idx += 8
	t.Message, err = utility.StringFromBytes(b)
	if err != nil {
		return err
	}
	idx += 1 + len([]byte(t.Message))

	// TODO: timestamp

	t.Node = b[idx : idx+transaction.ByteLenPubKey]
	idx += transaction.ByteLenPubKey
	t.Sign = b[idx : idx+transaction.ByteLenSign]
	idx += transaction.ByteLenSign
	return nil
}

func (t *Transfer) Hash() []byte {
	temp := make([]byte, 64)
	temp = append(temp, t.Sender...)
	temp = append(temp, t.Receiver...)
	temp = append(temp, t.Message...)
	temp = append(temp, strconv.FormatFloat(t.Pay, 'e', 8, 64)...)
	temp = append(temp, strconv.FormatFloat(t.Fee, 'e', 8, 64)...)
	temp = append(temp, t.Timestamp.String()...)
	temp = append(temp, t.Node...)
	return cr.Keccak256(temp)
}

func (t *Transfer) Verify() error {
	hash := t.Hash()
	if len(t.Sign) < 64 {
		return errors.New("incorrect signature len: " + strconv.FormatInt(int64(len(t.Sign)), 10))
	}
	if !cr.VerifySignature(t.Sender, hash, t.Sign[0:64]) {
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

// TODO: finish
func NewTransfer(sndrPrivKey *ecdsa.PrivateKey, rcvrPubKey []byte, value, fee float64, message string) (*Transfer, error) {
	sndrPubKey := utility.PrivToPubKey(sndrPrivKey)

	tr := Transfer{
		Sender:   sndrPubKey,
		Receiver: rcvrPubKey,
		Pay:      value,
		Fee:      fee,
		Message:  message, // TODO: add fix size message
	}

	tr.Timestamp = utility.NewTimestamp()
	hashed := tr.Hash()

	sign, err := cr.Sign(hashed, sndrPrivKey)
	if err != nil {
		return nil, err
	}

	tr.Sign = sign

	return &tr, nil
}
