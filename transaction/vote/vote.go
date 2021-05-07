package vote

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/overseven/blockchain/transaction"
	"github.com/overseven/blockchain/utility"
)

type Vote struct {
	Sender    []byte
	VotingId  uint64
	Opinion   uint8
	Fee       float64
	Timestamp time.Time
	Node      []byte
	Sign      []byte
}

func (v *Vote) IsEqual(transaction.Transaction, map[transaction.TransFlag]bool) bool {

}

func (v *Vote) String() (string, error) {
	tmp, err := json.Marshal(v)
	if err != nil {
		return "", err
	}

	return string(tmp), nil
}

func (v *Vote) Bytes() ([]byte, error) {
	var res []byte

	res = append(res, byte(transaction.TypeTransfer))

	if len(v.Sender) != transaction.ByteLenPubKey {
		return nil, errors.New("incorrect sender field size")
	}
	res = append(res, v.Sender...)

	res = append(res, utility.UInt64Bytes(v.VotingId)...)

	res = append(res, v.Opinion)
	res = append(res, utility.Float64Bytes(v.Fee)...)

	timestamp, err := utility.TimestampToBytes(v.Timestamp)
	if err != nil {
		return nil, err
	}

	res = append(res, uint8(len(timestamp)))
	res = append(res, timestamp...)

	if len(v.Node) != transaction.ByteLenPubKey {
		return nil, errors.New("incorrect node field size")
	}
	res = append(res, v.Node...)

	if len(v.Sign) != transaction.ByteLenSign {
		return nil, errors.New("incorrect sign field size")
	}

	res = append(res, v.Sign...)

	return res, nil
}

func (v *Vote) FromBytes(b []byte) error {
	if len(b) < 64 { // TODO: define min size
		return errors.New("incorrect input data len")
	}
	idx := 0
	v.Sender = b[idx : idx+transaction.ByteLenPubKey]
	idx += transaction.ByteLenPubKey

	v.VotingId = utility.UInt64FromBytes(b[idx : idx+8])
	idx += 8
	v.Opinion = b[idx]
	idx += 1
	v.Fee = utility.Float64FromBytes(b[idx : idx+8])
	idx += 8

	// TODO: timestamp

	v.Node = b[idx : idx+transaction.ByteLenPubKey]
	idx += transaction.ByteLenPubKey
	v.Sign = b[idx : idx+transaction.ByteLenSign]
	idx += transaction.ByteLenSign
	return nil
}

func (v *Vote) Hash(map[transaction.TransFlag]bool) ([]byte, error) {

}

func (v *Vote) Verify() error {

}
