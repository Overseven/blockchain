package airdrop

import (
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
	AirDropModeratorPubKey []byte
)

type Airdrop struct {
	Receiver  []byte
	Pay       float64
	Fee       float64
	Message   string
	Timestamp time.Time
	Sign      []byte
}

func (t *Airdrop) String() (string, error) {
	tmp, err := json.Marshal(t)
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
	res = append(res, utility.StringToBytes(a.Message)...)

	if len(a.Sign) != transaction.ByteLenSign {
		return nil, errors.New("incorrect sign field size")
	}

	res = append(res, a.Sign...)

	return res, nil
}

func (a *Airdrop) FromBytes([]byte) error {
	// TODO: finish
	return nil
}

func (a *Airdrop) Hash() []byte {
	temp := make([]byte, 64)
	temp = append(temp, a.Receiver...)
	temp = append(temp, a.Message...)
	temp = append(temp, strconv.FormatFloat(a.Pay, 'e', 8, 64)...)
	temp = append(temp, strconv.FormatFloat(a.Fee, 'e', 8, 64)...)
	temp = append(temp, a.Timestamp.String()...)
	return cr.Keccak256(temp)
}

func (a *Airdrop) Verify() error {

	if len(AirDropModeratorPubKey) == 0 {
		return errors.New("empty AirDrop moderator public key")
	}

	hash := a.Hash()
	if !cr.VerifySignature(AirDropModeratorPubKey, hash, a.Sign[0:64]) {
		return errors.New("incorrect AirDrop moderator signature")
	}

	// if bytes.Compare(a.Sender, AirDropModeratorPubKey) != 0 {
	// 	return errors.New("incorrect AirDrop moderator public key")
	// }

	return nil
}

// Airdrop is sending value from unlimited admin wallet to user wallet
func NewAirdrop(receiver []byte, adminPrivKey *ecdsa.PrivateKey, payment, fee float64) (*Airdrop, error) {
	// TODO: add check below

	a := new(Airdrop)
	a.Receiver = receiver

	a.Pay = payment
	a.Fee = fee
	a.Message = "Airdrop"
	{
		t := time.Now()
		a.Timestamp = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.UTC)
	}
	hashed := a.Hash()

	sign, err := cr.Sign(hashed, adminPrivKey)
	if err != nil {
		return nil, err
	}

	a.Sign = sign

	return a, nil
}
