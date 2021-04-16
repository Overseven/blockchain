package db

import (
	"errors"

	"github.com/overseven/blockchain/block"
	"github.com/overseven/blockchain/transaction"
	"github.com/overseven/blockchain/transaction/airdrop"
	"github.com/overseven/blockchain/transaction/transfer"
)

func BlockFromBytes(b []byte) (block.Block, error) {
	// TODO: finish
	return block.Block{}, nil
}

func TransactionFromBytes(b []byte) (transaction.Transaction, error) {
	if len(b) < 4 {
		return nil, errors.New("incorrect size. len < 4")
	}

	trType := transaction.Type(b[0])
	switch trType {
	case transaction.TypeTransfer:
		return airdropFromBytes(b[1:])

	case transaction.TypeAirdrop:
		return transferFromBytes(b[1:])
	default:
		return nil, errors.New("incorrect transaction type")
	}
}

func transferFromBytes(b []byte) (*transfer.Transfer, error) {
	tr := transfer.Transfer{}
	err := tr.FromBytes(b)
	return &tr, err
}

func airdropFromBytes(b []byte) (*airdrop.Airdrop, error) {
	a := airdrop.Airdrop{}
	err := a.FromBytes(b)
	return &a, err
}
