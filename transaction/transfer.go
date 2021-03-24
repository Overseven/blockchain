package transaction

import (
	"crypto/ecdsa"
	"errors"
	"strconv"
	"time"

	"github.com/overseven/blockchain/interfaces"

	cr "github.com/ethereum/go-ethereum/crypto"
	// "github.com/overseven/blockchain/chain"
	"github.com/overseven/blockchain/utility"
)

type Transfer struct {
	Data interfaces.Data
}

func (t *Transfer) GetData() *interfaces.Data {
	return &t.Data
}

func (t *Transfer) SetData(d *interfaces.Data) {
	t.Data = *d
}

func (t *Transfer) Verify(balance interfaces.Balancer) error {
	hash := GetHash(t.GetData())
	if len(t.Data.Sign) < 64 {
		return errors.New("incorrect signature len: " + strconv.FormatInt(int64(len(t.Data.Sign)), 10))
	}
	if !cr.VerifySignature(t.Data.Pubkey, hash, t.Data.Sign[0:64]) {
		return errors.New("incorrect signature")
	}

	senderWallet, err := balance.Info(t.Data.Pubkey)
	if err != nil {
		return err
	}
	if senderWallet.CurrentBalance < (t.Data.Pay + t.Data.Fee) {
		return errors.New("sender wallet not enough tokens")
	}

	return nil
}

// TODO: finish
func NewTransfer(sndrPrivKey *ecdsa.PrivateKey, rcvrPubKey []byte, value, fee float64, message string, balance interfaces.Balancer) (*Transfer, error) {
	sndrPubKey := utility.PrivToPubKey(sndrPrivKey)
	wall, err := balance.Info(sndrPubKey)
	if err != nil {
		return nil, err
	}
	if wall.CurrentBalance < (value + fee) {
		return nil, errors.New("not enough tokens")
	}

	data := interfaces.Data{}

	data.Type = TypeTransfer
	data.Pubkey = sndrPubKey
	data.Receiver = rcvrPubKey
	data.Pay = value
	data.Fee = fee
	data.Message = message // TODO: add fix size message

	{
		t := time.Now()
		data.Timestamp = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.UTC)
	}
	hashed := GetHash(&data)

	sign, err := cr.Sign(hashed, sndrPrivKey)
	if err != nil {
		panic(err)
	}

	data.Sign = sign

	return &Transfer{Data: data}, nil
}
