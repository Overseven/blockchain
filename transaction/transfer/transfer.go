package transfer

import (
	"errors"

	"github.com/Overseven/blockchain/blockchain"
	tr "github.com/Overseven/blockchain/transaction"
	"github.com/Overseven/blockchain/wallet"
	cr "github.com/ethereum/go-ethereum/crypto"
)

type Transfer struct {
	data tr.Data
}

func (t *Transfer) Verify() bool {
	hash := tr.GetHash(&t.data)
	if !cr.VerifySignature(t.data.Pubkey, hash, t.data.Sign[0:64]) {
		return false
	}

	senderWallet, err := wallet.Info(t.data.Pubkey)
	if err != nil {
		return false
	}
	if senderWallet.CurrentBalance < (t.data.Pay + t.data.Fee) {
		return false
	}

	return true
}

// TODO: finish
func New(receiver, adminPrivKey []byte) (*Transfer, error) {
	if len(blockchain.B17.Blocks) > 0 {
		return nil, errors.New("airdrop for not first block is not alowed")
	}
	a := new(Transfer)

	a.data.Type = tr.Transfer

	return a, nil
}

func (t *Transfer) SetData(d *tr.Data) {
	t.data = *d
}
