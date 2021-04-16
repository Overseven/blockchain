package converter

import (
	"errors"
	"strconv"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/overseven/blockchain/block"
	"github.com/overseven/blockchain/transaction"
	"github.com/overseven/blockchain/transaction/airdrop"
	"github.com/overseven/blockchain/transaction/transfer"
)

const (
	protocolVersion = 1
)

func BlockProto2Local(b *pb.Block) (*block.Block, error) {
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

func BlockLocal2Proto(b block.Block) (*pb.Block, error) {
	bl := pb.Block{}
	bl.BlockId = b.Id

	var tr []*pb.Transaction
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

func AirdropProto2Local(a *pb.TransAirDrop) (*airdrop.Airdrop, error) {
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

func TransferProto2Local(t *pb.TransTransfer) (*transfer.Transfer, error) {
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

func TransactionProto2Local(t *pb.Transaction) (transaction.Transaction, error) {
	switch tmp := t.Trans.(type) {
	case *pb.Transaction_Drop:
		return AirdropProto2Local(tmp.Drop)
	case *pb.Transaction_Transfer:
		return TransferProto2Local(tmp.Transfer)
	default:
		return nil, errors.New("incorrect trans type")
	}
}

func TransferLocal2Proto(tr *transfer.Transfer) (*pb.Transaction, error) {
	t := new(pb.TransTransfer)
	t.Sender = tr.Sender
	t.Receiver = tr.Receiver
	t.Pay = tr.Pay
	t.Fee = tr.Fee
	t.Message = tr.Message
	t.Timestamp = timestamppb.New(tr.Timestamp)
	t.Node = tr.Node
	t.Sign = tr.Sign
	return &pb.Transaction{ProtocolVersion: protocolVersion, Trans: &pb.Transaction_Transfer{Transfer: t}}, nil
}

func AirdropLocal2Proto(tr *airdrop.Airdrop) (*pb.Transaction, error) {
	t := new(pb.TransAirDrop)
	t.Receiver = tr.Receiver
	t.Pay = tr.Pay
	t.Fee = tr.Fee
	t.Message = tr.Message
	t.Timestamp = timestamppb.New(tr.Timestamp)
	t.Node = tr.Node
	t.Sign = tr.Sign
	return &pb.Transaction{ProtocolVersion: protocolVersion, Trans: &pb.Transaction_Drop{Drop: t}}, nil
}
func TransactionLocal2Proto(tr transaction.Transaction) (*pb.Transaction, error) {
	switch tmp := tr.(type) {
	case *airdrop.Airdrop:
		return AirdropLocal2Proto(tmp)
	case *transfer.Transfer:
		return TransferLocal2Proto(tmp)
	default:
		return nil, errors.New("incorrect trans type")
	}
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
