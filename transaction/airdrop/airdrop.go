package airdrop

import (
	"bytes"
	"errors"

	"github.com/Overseven/blockchain/blockchain"
	tr "github.com/Overseven/blockchain/transaction"
	cr "github.com/ethereum/go-ethereum/crypto"
)

var (
	AirDropModeratorPubKey []byte
)

type Airdrop struct {
	data tr.Data
}

/*
	Use it only if the sender has no more than one transaction in the block
*/
func (a *Airdrop) Verify() bool {

	if len(AirDropModeratorPubKey) == 0 {
		return false
	}

	hash := tr.GetHash(&a.data)
	if !cr.VerifySignature(a.data.Pubkey, hash, a.data.Sign[0:64]) {
		return false
	}

	if bytes.Compare(a.data.Pubkey, AirDropModeratorPubKey) != 0 {
		return false
	}

	return true
}

// Airdrop is sending value from unlimited admin wallet to user wallet
func New(receiver, adminPrivKey []byte) (*Airdrop, error) {
	if len(blockchain.B17.Blocks) > 0 {
		return nil, errors.New("airdrop for not first block is not alowed")
	}
	a := new(Airdrop)

	a.data.Type = tr.Airdrop

	return a, nil
}

func (a *Airdrop) SetData(d *tr.Data) {
	a.data = *d
}
