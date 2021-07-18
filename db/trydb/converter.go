package trydb

import (
	"errors"
	"github.com/overseven/try-network/block/tryblock"
	"github.com/overseven/try-network/utility"

	"github.com/overseven/try-network/transaction"
	"github.com/overseven/try-network/transaction/airdrop"
	"github.com/overseven/try-network/transaction/transfer"
)

func BlockFromBytes(b []byte) (Block, error) {
	if len(b) < 9 {
		return Block{}, errors.New("incorrect input data len")
	}

	block :=  Block{}
	idx := int64(0)
	block.Id = utility.UInt64FromBytes(b[idx:idx+8])
	idx += 8
	block.NumOfTrans = utility.UInt8FromBytes(b[idx:idx+1])
	idx += 1
	if len(b) < int(32 * block.NumOfTrans + 8 + 32 + 32 + 8) {
		return Block{}, errors.New("incorrect input data len! ")
	}
	block.TransHashes = make([][]byte, block.NumOfTrans)
	var i uint8
	for ; i < block.NumOfTrans; i++ {
		copy(block.TransHashes[i], b[idx:idx+32])
		idx += 32
	}

	block.Difficulty = utility.UInt64FromBytes(b[idx:idx+8])
	idx += 8

	copy(block.MinerPubKey, b[idx:idx+32])
	idx += 32

	copy(block.Hash, b[idx:idx+32])
	idx += 32

	block.Nonce = utility.UInt64FromBytes(b[idx:idx+8])
	idx += 8

	return block, nil
}

func BlockBytes (block tryblock.TryBlock) ([]byte, error){
	var b []byte
	b = append(b, utility.UInt64Bytes(block.Id)...)
	if block.Transactions != nil {
		b = append(b, utility.UInt32Bytes(uint32(len(block.Transactions)))...)
		for _, t := range block.Transactions {
			tHash, err := t.Hash(map[transaction.TransFlag]bool{})
			if err != nil {
				return nil, err
			}
			b = append(b, tHash...)
		}
	}
	b = append(b, block.PrevHash...)
	b = append(b, utility.UInt64Bytes(block.Difficulty)...)
	b = append(b, block.Miner...)
	timeHash, err := utility.TimestampToBytes(block.Timestamp)
	if err != nil {
		return nil, err
	}
	b = append(b, timeHash...)
	b = append(b, block.Hash...)
	b = append(b, utility.UInt64Bytes(block.Nonce)...)

	return b, nil
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
	//err := tr.FromBytes(b) // TODO: create byte to Transfer conversation in transaction/transfer/transfer.go
	return &tr, nil
}

func airdropFromBytes(b []byte) (*airdrop.Airdrop, error) {
	a := airdrop.Airdrop{}
	//err := a.FromBytes(b)  // TODO: create byte to Transfer conversation in transaction/transfer/transfer.go
	return &a, nil
}
