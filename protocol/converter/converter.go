package converter

import (
	"errors"
	"strconv"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/overseven/blockchain/block"
	"github.com/overseven/blockchain/interfaces"
	pb "github.com/overseven/blockchain/protocol"
	"github.com/overseven/blockchain/transaction"
	"github.com/overseven/blockchain/transaction/airdrop"
	"github.com/overseven/blockchain/transaction/transfer"
)

func BlockProto2Local(b *pb.Block) (interfaces.TransactionsContainer, error) {
	var bl interfaces.TransactionsContainer = new(block.Block)

	bl.SetId(b.GetBlockId())

	// transactions
	var tr []interfaces.BlockElement
	for i, t := range b.GetTrans() {
		tran, err := TransactionProto2Local(t)
		if err != nil {
			return nil, errors.New("incorrect type at " + strconv.FormatInt(int64(i), 10) + " transaction")
		}
		tr = append(tr, tran)

	}
	bl.SetTransactions(tr)

	bl.SetPrevHash(b.PrevBlockHash)
	bl.SetDifficulty(uint64(b.Difficulty))
	bl.SetMiner(b.Miner)
	bl.GetHash()
	bl.SetNonce(b.Nonce)

	return bl, nil
}

func BlockLocal2Proto(b interfaces.TransactionsContainer) (*pb.Block, error) {
	bl := pb.Block{}
	bl.BlockId = b.GetId()

	var tr []*pb.Transaction
	for i, t := range b.GetTransactions() {
		tran, err := TransactionLocal2Proto(t)
		if err != nil {
			return nil, errors.New("incorrect type at " + strconv.FormatInt(int64(i), 10) + " transaction")
		}
		tr = append(tr, tran)

	}
	bl.Trans = tr
	bl.PrevBlockHash = b.GetPrevHash()
	bl.Difficulty = uint32(b.GetDifficulty())
	bl.Miner = b.GetMiner()
	bl.BlockHash = b.GetHash()
	bl.Nonce = b.GetNonce()
	return &bl, nil
}

func TransactionProto2Local(t *pb.Transaction) (interfaces.BlockElement, error) {
	data := interfaces.Data{}
	data.Sender = t.Sender
	data.Receiver = t.Receiver
	data.Message = t.Message
	data.Timestamp = t.GetTimestamp().AsTime()
	data.Pay = t.Pay
	data.Fee = t.Fee
	data.Sign = t.GetSenderSign()

	if t.GetType() == transaction.TypeAirdrop {
		return &transaction.Airdrop{Data: data}, nil
	} else if t.GetType() == transaction.TypeTransfer {
		return &transaction.Transfer{Data: data}, nil
	} else {
		return nil, errors.New("incorrect type of transaction")
	}
}

func TransactionLocal2Proto(trans interfaces.BlockElement) (*pb.Transaction, error) {
	data := trans.GetData()
	tr := new(pb.Transaction)
	tr.Type = pb.Transaction_Type(data.Type)
	tr.Sender = data.Sender
	tr.Receiver = data.Receiver
	tr.Message = data.Message

	var err error
	tr.Timestamp = timestamppb.New(data.Timestamp)
	if err != nil {
		return nil, err
	}
	tr.Pay = data.Pay
	tr.Fee = data.Fee
	tr.SenderSign = data.Sign
	return tr, nil
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
