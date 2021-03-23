package airdrop

import (
	"bytes"

	"github.com/Overseven/blockchain/chain"
	tr "github.com/Overseven/blockchain/transaction"
	cr "github.com/ethereum/go-ethereum/crypto"
)

var (
	AirDropModeratorPubKey []byte
)

type Airdrop struct {
	Data tr.Data
}

func (a *Airdrop) GetData() *tr.Data {
	return &a.Data
}

/*
	Use it only if the sender has no more than one transaction in the block
*/
func (a *Airdrop) Verify() bool {

	if len(AirDropModeratorPubKey) == 0 {
		return false
	}

	hash := tr.GetHash(a.GetData())
	if !cr.VerifySignature(a.Data.Pubkey, hash, a.Data.Sign[0:64]) {
		return false
	}

	if bytes.Compare(a.Data.Pubkey, AirDropModeratorPubKey) != 0 {
		return false
	}

	return true
}

// Airdrop is sending value from unlimited admin wallet to user wallet
func New(receiver, adminPrivKey []byte, chain chain.Chain) (*Airdrop, error) {
	// TODO: add check below

	// if len(blockchain.B17.Blocks) > 0 {
	// 	return nil, errors.New("airdrop for not first block is not alowed")
	// }
	a := new(Airdrop)

	a.Data.Type = tr.Airdrop

	return a, nil
}

func (a *Airdrop) SetData(d *tr.Data) {
	a.Data = *d
}
