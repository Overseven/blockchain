package transaction

import (
	"crypto/ecdsa"
	"errors"
	"github.com/overseven/blockchain/transaction/itransaction"
	"strconv"
	"time"

	"github.com/Overseven/blockchain/utility"
	"github.com/Overseven/blockchain/wallet"
	cr "github.com/ethereum/go-ethereum/crypto"
)

type Transfer struct {
	itransaction.ITransaction
	Data itransaction.Data
}

func (t *Transfer) GetData() *itransaction.Data {
	return &t.Data
}

func (t *Transfer) SetData(d *itransaction.Data) {
	t.Data = *d
}

func (t *Transfer) Verify() error {
	hash := itransaction.GetHash(t.GetData())
	if len(t.Data.Sign) < 64 {
		return errors.New("incorrect signature len: " + strconv.FormatInt(int64(len(t.Data.Sign)), 10))
	}
	if !cr.VerifySignature(t.Data.Pubkey, hash, t.Data.Sign[0:64]) {
		return errors.New("incorrect signature")
	}

	senderWallet, err := wallet.Info(t.Data.Pubkey)
	if err != nil {
		return err
	}
	if senderWallet.CurrentBalance < (t.Data.Pay + t.Data.Fee) {
		return errors.New("sender wallet not enough tokens")
	}

	return nil
}

// TODO: finish
func NewTransfer(sndrPrivKey *ecdsa.PrivateKey, rcvrPubKey []byte, value, fee float64, message string) (*Transfer, error) {
	sndrPubKey := utility.PrivToPubKey(sndrPrivKey)
	wall, err := wallet.Info(sndrPubKey)
	if err != nil {
		return nil, err
	}
	if wall.CurrentBalance < (value + fee) {
		return nil, errors.New("not enough tokens")
	}

	data := itransaction.Data{}

	data.Type = itransaction.Transfer
	data.Pubkey = sndrPubKey
	data.Receiver = rcvrPubKey
	data.Pay = value
	data.Fee = fee
	data.Message = message // TODO: add fix size message

	{
		t := time.Now()
		data.Timestamp = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.UTC)
	}
	hashed := itransaction.GetHash(&data)

	sign, err := cr.Sign(hashed, sndrPrivKey)
	if err != nil {
		panic(err)
	}

	data.Sign = sign

	return &Transfer{Data: data}, nil
}
