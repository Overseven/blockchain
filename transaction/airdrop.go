package transaction

import (
	"bytes"
	"crypto/ecdsa"
	"errors"
	"time"

	// "github.com/overseven/blockchain/chain/ichain"
	"github.com/overseven/blockchain/interfaces"

	cr "github.com/ethereum/go-ethereum/crypto"
)

var (
	AirDropModeratorPubKey []byte
)

type Airdrop struct {
	Data interfaces.Data
}

func (a *Airdrop) GetData() *interfaces.Data {
	return &a.Data
}

func (a *Airdrop) Verify(balance interfaces.Balancer) error {

	if len(AirDropModeratorPubKey) == 0 {
		return errors.New("empty AirDrop moderator public key")
	}

	hash := GetHash(a.GetData())
	if !cr.VerifySignature(a.Data.Pubkey, hash, a.Data.Sign[0:64]) {
		return errors.New("incorrect AirDrop moderator signature")
	}

	if bytes.Compare(a.Data.Pubkey, AirDropModeratorPubKey) != 0 {
		return errors.New("incorrect AirDrop moderator public key")
	}

	return nil
}

// Airdrop is sending value from unlimited admin wallet to user wallet
func NewAirdrop(receiver []byte, adminPrivKey *ecdsa.PrivateKey, chain interfaces.Chainable, payment float64, balance interfaces.Balancer) (*Airdrop, error) {
	// TODO: add check below

	// if len(blockchain.B17.Blocks) > 0 {
	// 	return nil, errors.New("airdrop for not first block is not alowed")
	// }
	a := new(Airdrop)

	a.Data.Type = TypeAirdrop
	a.Data.Receiver = receiver

	a.Data.Pay = payment

	{
		t := time.Now()
		a.Data.Timestamp = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.UTC)
	}
	hashed := GetHash(&a.Data)

	sign, err := cr.Sign(hashed, adminPrivKey)
	if err != nil {
		return nil, err
	}
	a.Data.Message = "Airdrop"
	a.Data.Sign = sign

	return a, nil
}

func (a *Airdrop) SetData(d *interfaces.Data) {
	a.Data = *d
}
