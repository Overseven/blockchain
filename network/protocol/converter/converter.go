package converter

import (
	"errors"
	"strconv"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/overseven/try-network/block"
	"github.com/overseven/try-network/network/protocol/pcommon"
	"github.com/overseven/try-network/transaction"
	"github.com/overseven/try-network/transaction/airdrop"
	"github.com/overseven/try-network/transaction/transfer"
)

const (
	protocolVersion = 1
)

func BlockProto2Local(b *pcommon.Block) (*block.Block, error) {
	bl := new(block.Block)

	bl.Id = b.GetBlockId()

	// transactions
	tr := map[string]transaction.Transaction{}
	for i, t := range b.GetTrans() {
		tran, err := TransactionProto2Local(t)
		if err != nil {
			return nil, errors.New("incorrect type at " + strconv.FormatInt(int64(i), 10) + " transaction")
		}
		tr[string(tran.Hash())] = tran

	}
	bl.Transactions = tr

	bl.PrevHash = b.PrevBlockHash
	bl.Difficulty = uint64(b.Difficulty)
	bl.Miner = b.Miner
	bl.GetHash()
	bl.Nonce = b.Nonce

	return bl, nil
}

func BlockLocal2Proto(b block.Block) (*pcommon.Block, error) {
	bl := pcommon.Block{}
	bl.BlockId = b.Id

	var tr []*pcommon.Transaction
	for _, t := range b.Transactions {
		tran, err := TransactionLocal2Proto(t)
		if err != nil {
			return nil, errors.New("incorrect transaction type")
		}
		tr = append(tr, tran)

	}
	bl.Trans = tr
	bl.PrevBlockHash = b.PrevHash
	bl.Difficulty = uint32(b.Difficulty)
	bl.Miner = b.Miner
	bl.BlockHash = b.GetHash()
	bl.Nonce = b.Nonce
	return &bl, nil
}

func AirdropProto2Local(a *pcommon.TransAirDrop) (*airdrop.Airdrop, error) {
	local_a := new(airdrop.Airdrop)
	local_a.Receiver = a.Receiver
	local_a.Timestamp = a.GetTimestamp().AsTime()
	local_a.Pay = a.Pay
	local_a.Fee = a.Fee
	local_a.Message = a.Message
	local_a.Node = a.Node
	local_a.Sign = a.Sign
	return local_a, nil
}

func TransferProto2Local(t *pcommon.TransTransfer) (*transfer.Transfer, error) {
	local_tr := new(transfer.Transfer)
	local_tr.Sender = t.Sender
	local_tr.Receiver = t.Receiver
	local_tr.Message = t.Message
	local_tr.Timestamp = t.GetTimestamp().AsTime()
	local_tr.Pay = t.Pay
	local_tr.Fee = t.Fee
	local_tr.Node = t.Node
	local_tr.Sign = t.Sign
	return local_tr, nil
}

func TransactionProto2Local(t *pcommon.Transaction) (transaction.Transaction, error) {
	switch tmp := t.Trans.(type) {
	case *pcommon.Transaction_Drop:
		return AirdropProto2Local(tmp.Drop)
	case *pcommon.Transaction_Transfer:
		return TransferProto2Local(tmp.Transfer)
	default:
		return nil, errors.New("incorrect trans type")
	}
}

func TransferLocal2Proto(tr *transfer.Transfer) (*pcommon.Transaction, error) {
	t := new(pcommon.TransTransfer)
	t.Sender = tr.Sender
	t.Receiver = tr.Receiver
	t.Pay = tr.Pay
	t.Fee = tr.Fee
	t.Message = tr.Message
	t.Timestamp = timestamppb.New(tr.Timestamp)
	t.Node = tr.Node
	t.Sign = tr.Sign
	return &pcommon.Transaction{ProtocolVersion: protocolVersion, Trans: &pcommon.Transaction_Transfer{Transfer: t}}, nil
}

func AirdropLocal2Proto(tr *airdrop.Airdrop) (*pcommon.Transaction, error) {
	t := new(pcommon.TransAirDrop)
	t.Receiver = tr.Receiver
	t.Pay = tr.Pay
	t.Fee = tr.Fee
	t.Message = tr.Message
	t.Timestamp = timestamppb.New(tr.Timestamp)
	t.Node = tr.Node
	t.Sign = tr.Sign
	return &pcommon.Transaction{ProtocolVersion: protocolVersion, Trans: &pcommon.Transaction_Drop{Drop: t}}, nil
}

func TransactionLocal2Proto(tr transaction.Transaction) (*pcommon.Transaction, error) {
	switch tmp := tr.(type) {
	case *airdrop.Airdrop:
		return AirdropLocal2Proto(tmp)
	case *transfer.Transfer:
		return TransferLocal2Proto(tmp)
	default:
		return nil, errors.New("incorrect trans type")
	}
}
