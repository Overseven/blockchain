package converter

import (
	"errors"
	"strconv"

	"github.com/golang/protobuf/ptypes"
	"github.com/overseven/blockchain/block"
	"github.com/overseven/blockchain/interfaces"
	pb "github.com/overseven/blockchain/protocol"
	"github.com/overseven/blockchain/transaction"
)

func BlockProto2Local(b *pb.Block) (interfaces.Blockable, error) {
	var bl interfaces.Blockable = new(block.Block)

	bl.SetId(b.GetBlockId())

	// transactions
	var tr []interfaces.Transferable
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

func BlockLocal2Proto(b interfaces.Blockable) (*pb.Block, error) {
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

func TransactionProto2Local(t *pb.Transaction) (interfaces.Transferable, error) {
	data := interfaces.Data{}
	data.Pubkey = t.Sender
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

func TransactionLocal2Proto(trans interfaces.Transferable) (*pb.Transaction, error) {
	data := trans.GetData()
	tr := new(pb.Transaction)
	tr.Type = pb.Transaction_Type(data.Type)
	tr.Sender = data.Pubkey
	tr.Receiver = data.Receiver
	tr.Message = data.Message

	var err error
	tr.Timestamp, err = ptypes.TimestampProto(data.Timestamp)
	if err != nil {
		return nil, err
	}
	tr.Pay = data.Pay
	tr.Fee = data.Fee
	tr.SenderSign = data.Sign
	return tr, nil
}
