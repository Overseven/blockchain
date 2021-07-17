package converter

import (
	"errors"
	"github.com/overseven/try-network/transaction/vote"
	vf "github.com/overseven/try-network/transaction/voting_finish"
	"strconv"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/overseven/try-network/block"
	"github.com/overseven/try-network/network/protocol/proto"
	"github.com/overseven/try-network/transaction"
	"github.com/overseven/try-network/transaction/airdrop"
	"github.com/overseven/try-network/transaction/transfer"
	vi "github.com/overseven/try-network/transaction/voting_init"
)

const (
	protocolVersion = 1
)

func BlockProto2Local(b *proto.Block) (*block.Block, error) {
	bl := new(block.Block)

	bl.Id = b.GetBlockId()

	// transactions
	tr := map[string]transaction.Transaction{}
	for i, t := range b.GetTrans() {
		tran, err := TransactionProto2Local(t)
		if err != nil {
			return nil, errors.New("incorrect type at " + strconv.FormatInt(int64(i), 10) + " transaction")
		}
		tHash, err := tran.Hash(map[transaction.TransFlag]bool{})
		if err != nil {
			return nil, err
		}
		tr[string(tHash)] = tran

	}
	bl.Transactions = tr

	bl.PrevHash = b.PrevBlockHash
	bl.Difficulty = uint64(b.Difficulty)
	bl.Miner = b.Miner
	_, err := bl.GetHash()
	if err != nil {
		return nil, err
	}

	bl.Nonce = b.Nonce

	return bl, nil
}

func BlockLocal2Proto(b block.Block) (*proto.Block, error) {
	bl := proto.Block{}
	bl.BlockId = b.Id

	var tr []*proto.Transaction
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
	var err error
	bl.BlockHash, err = b.GetHash()
	if err != nil {
		return nil, errors.New("incorrect transaction type")
	}
	bl.Nonce = b.Nonce
	return &bl, nil
}

func AirdropProto2Local(a *proto.TransAirDrop) (*airdrop.Airdrop, error) {
	localA := new(airdrop.Airdrop)
	localA.TransCounter = a.TransactionCounter
	localA.Receiver = a.Receiver
	localA.Timestamp = a.GetTimestamp().AsTime()
	localA.Pay = a.Pay
	localA.Fee = a.Fee
	localA.Message = a.Message
	localA.Node = a.Node
	localA.Signature = a.Sign
	return localA, nil
}

func TransferProto2Local(t *proto.TransTransfer) (*transfer.Transfer, error) {
	localTr := new(transfer.Transfer)
	localTr.Sender = t.Sender
	localTr.TransCounter = t.TransactionCounter
	localTr.Receiver = t.Receiver
	localTr.Message = t.Message
	localTr.Timestamp = t.GetTimestamp().AsTime()
	localTr.Pay = t.Pay
	localTr.Fee = t.Fee
	localTr.Node = t.Node
	localTr.Signature = t.Sign
	return localTr, nil
}

func VotingInitProto2Local(t *proto.TransVotingInit) (*vi.VotingInit, error) {
	localVi := new(vi.VotingInit)
	localVi.Sender = t.Sender
	localVi.TransCounter = t.TransactionCounter
	localVi.VotingId = t.VotingId
	localVi.Parameter = uint16(t.Parameter)
	localVi.Value = t.Value
	localVi.Timestamp = t.GetTimestamp().AsTime()
	localVi.Fee = t.Fee
	localVi.Node = t.Node
	localVi.Signature = t.Sign
	return localVi, nil
}

func VoteProto2Local(t *proto.TransVote) (*vote.Vote, error) {
	localV := new(vote.Vote)
	localV.Sender = t.Sender
	localV.TransCounter = t.TransactionCounter
	localV.VotingId = t.VotingId
	localV.Opinion = t.Opinion
	localV.Timestamp = t.GetTimestamp().AsTime()
	localV.Fee = t.Fee
	localV.Node = t.Node
	localV.Signature = t.Sign
	return localV, nil
}

func VotingFinishProto2Local(t *proto.TransVotingFinish) (*vf.VotingFinish, error) {
	localVf := new(vf.VotingFinish)
	localVf.Sender = t.Sender
	localVf.TransCounter = t.TransactionCounter
	localVf.VotingId = t.VotingId
	localVf.Timestamp = t.GetTimestamp().AsTime()
	localVf.Fee = t.Fee
	localVf.Node = t.Node
	localVf.Signature = t.Sign
	return localVf, nil
}

func TransactionProto2Local(t *proto.Transaction) (transaction.Transaction, error) {
	switch tmp := t.Trans.(type) {
	case *proto.Transaction_Airdrop:
		return AirdropProto2Local(tmp.Airdrop)
	case *proto.Transaction_Transfer:
		return TransferProto2Local(tmp.Transfer)
	case *proto.Transaction_VotingInit:
		return VotingInitProto2Local(tmp.VotingInit)
	case *proto.Transaction_Vote:
		return VoteProto2Local(tmp.Vote)
	case *proto.Transaction_VotingFinish:
		return VotingFinishProto2Local(tmp.VotingFinish)
	default:
		return nil, errors.New("incorrect trans type")
	}
}

func TransferLocal2Proto(tr *transfer.Transfer) (*proto.Transaction, error) {
	t := new(proto.TransTransfer)
	t.Sender = tr.Sender
	t.TransactionCounter = tr.TransCounter
	t.Receiver = tr.Receiver
	t.Pay = tr.Pay
	t.Fee = tr.Fee
	t.Message = tr.Message
	t.Timestamp = timestamppb.New(tr.Timestamp)
	t.Node = tr.Node
	t.Sign = tr.Signature
	return &proto.Transaction{ProtocolVersion: protocolVersion, Trans: &proto.Transaction_Transfer{Transfer: t}}, nil
}

func AirdropLocal2Proto(tr *airdrop.Airdrop) (*proto.Transaction, error) {
	t := new(proto.TransAirDrop)
	t.TransactionCounter = tr.TransCounter
	t.Receiver = tr.Receiver
	t.Pay = tr.Pay
	t.Fee = tr.Fee
	t.Message = tr.Message
	t.Timestamp = timestamppb.New(tr.Timestamp)
	t.Node = tr.Node
	t.Sign = tr.Signature
	return &proto.Transaction{ProtocolVersion: protocolVersion, Trans: &proto.Transaction_Airdrop{Airdrop: t}}, nil
}

func VotingInitLocal2Proto(tr *vi.VotingInit) (*proto.Transaction, error) {
	t := new(proto.TransVotingInit)
	t.Sender = tr.Sender
	t.TransactionCounter = tr.TransCounter
	t.VotingId = tr.VotingId
	t.Parameter = uint32(tr.Parameter)
	t.Value = tr.Value
	t.Fee = tr.Fee
	t.Timestamp = timestamppb.New(tr.Timestamp)
	t.Node = tr.Node
	t.Sign = tr.Signature
	return &proto.Transaction{ProtocolVersion: protocolVersion, Trans: &proto.Transaction_VotingInit{VotingInit: t}}, nil
}

func VoteLocal2Proto(tr *vote.Vote) (*proto.Transaction, error) {
	t := new(proto.TransVote)
	t.Sender = tr.Sender
	t.TransactionCounter = tr.TransCounter
	t.VotingId = tr.VotingId
	t.Opinion = tr.Opinion
	t.Fee = tr.Fee
	t.Timestamp = timestamppb.New(tr.Timestamp)
	t.Node = tr.Node
	t.Sign = tr.Signature
	return &proto.Transaction{ProtocolVersion: protocolVersion, Trans: &proto.Transaction_Vote{Vote: t}}, nil
}

func VotingFinishLocal2Proto(tr *vf.VotingFinish) (*proto.Transaction, error) {
	t := new(proto.TransVotingFinish)
	t.Sender = tr.Sender
	t.TransactionCounter = tr.TransCounter
	t.VotingId = tr.VotingId
	t.Fee = tr.Fee
	t.Timestamp = timestamppb.New(tr.Timestamp)
	t.Node = tr.Node
	t.Sign = tr.Signature
	return &proto.Transaction{ProtocolVersion: protocolVersion, Trans: &proto.Transaction_VotingFinish{VotingFinish: t}}, nil
}

func TransactionLocal2Proto(tr transaction.Transaction) (*proto.Transaction, error) {
	switch tmp := tr.(type) {
	case *airdrop.Airdrop:
		return AirdropLocal2Proto(tmp)
	case *transfer.Transfer:
		return TransferLocal2Proto(tmp)
	case *vi.VotingInit:
		return VotingInitLocal2Proto(tmp)
	case *vote.Vote:
		return VoteLocal2Proto(tmp)
	case *vf.VotingFinish:
		return VotingFinishLocal2Proto(tmp)
	default:
		return nil, errors.New("incorrect trans type")
	}
}
