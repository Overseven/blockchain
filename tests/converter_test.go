package test

import (
	"bytes"
	"errors"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strconv"
	"testing"
	"time"

	cr "github.com/ethereum/go-ethereum/crypto"
	"github.com/overseven/blockchain/interfaces"
	pb "github.com/overseven/blockchain/protocol"
	"github.com/overseven/blockchain/protocol/converter"
	"github.com/overseven/blockchain/transaction"
	"github.com/overseven/blockchain/wallet"
)

func compareTransactions(t1 interfaces.Transferable, t2 *pb.Transaction) error {
	data := t1.GetData()
	if bytes.Compare(data.Sender, t2.Sender) != 0 {
		return errors.New("err: diff sender")
	}
	if bytes.Compare(data.Receiver, t2.Receiver) != 0 {
		return errors.New("err: diff receiver")
	}
	if data.Message != t2.Message {
		return errors.New("err: diff message")
	}
	if data.Timestamp != t2.GetTimestamp().AsTime() {
		return errors.New("err: diff timestamp")
	}
	if data.Pay != t2.Pay {
		return errors.New("err: diff pay: " + strconv.FormatFloat(data.Pay, 'e', 7, 64) + " vs " + strconv.FormatFloat(t2.Pay, 'e', 7, 64))
	}
	if data.Fee != t2.Fee {
		return errors.New("err: diff pay: " + strconv.FormatFloat(data.Fee, 'e', 7, 64) + " vs " + strconv.FormatFloat(t2.Fee, 'e', 7, 64))
	}
	if bytes.Compare(data.Sign, t2.GetSenderSign()) != 0 {
		return errors.New("err: diff sign")
	}

	return nil
}

func compareBlocks(b1 interfaces.Blockable, b2 *pb.Block) error {

	if b1.GetId() != b2.BlockId {
		return errors.New("err: diff pay: " + strconv.FormatUint(b1.GetId(), 10) + " vs " + strconv.FormatUint(b2.BlockId, 10))
	}

	if len(b1.GetTransactions()) != len(b2.Trans) {
		return errors.New("err: diff pay: " + strconv.FormatInt(int64(len(b1.GetTransactions())), 10) + " vs " + strconv.FormatInt(int64(len(b2.Trans)), 10))
	}
	for i := 0; i < len(b2.Trans); i++ {
		err := compareTransactions(b1.GetTransactions()[i], b2.Trans[i])
		if err != nil {
			return errors.New("err: incorrect at " + strconv.FormatInt(int64(i), 10) + " transaction: " + err.Error())
		}
	}

	if bytes.Compare(b1.GetPrevHash(), b2.PrevBlockHash) != 0 {
		return errors.New("err: diff prev block hash")
	}

	if b1.GetDifficulty() != uint64(b2.Difficulty) {
		return errors.New("err: diff difficulty")
	}

	if bytes.Compare(b1.GetMiner(), b2.Miner) != 0 {
		return errors.New("err: diff Miner")
	}

	if bytes.Compare(b1.GetHash(), b2.BlockHash) != 0 {
		return errors.New("err: diff block hash")
	}

	if bytes.Compare(b1.GetNonce(), b2.Nonce) != 0 {
		return errors.New("err: diff nonce")
	}

	return nil
}

func TestTransLocal2ProtoAirdrop(t *testing.T) {
	_, rcvrPubKey, err := generateWallet()

	if err != nil {
		t.Error(err)
	}

	airdropPubKey, airdropPrKey, err := wallet.LoadFromFile(airdropModeratorConfigFile)
	if err != nil {
		t.Error(err)
	}

	transaction.AirDropModeratorPubKey = airdropPubKey

	airPrKey, err := cr.ToECDSA(airdropPrKey[:32])
	if err != nil {
		t.Error(err)
	}

	airdrop, err := transaction.NewAirdrop(rcvrPubKey, airPrKey, 178.9, 12.1)
	if err != nil {
		t.Error(err)
	}

	airdropPb, err := converter.TransactionLocal2Proto(airdrop)
	if err != nil {
		t.Error(err)
	}

	err = compareTransactions(airdrop, airdropPb)
	if err != nil {
		t.Error(err)
	}

}

func TestTransLocal2ProtoTransfer(t *testing.T) {

	sndrPrivKey, _, err := generateWallet()
	if err != nil {
		t.Error(err)
	}

	_, rcvrPubKey, err := generateWallet()
	if err != nil {
		t.Error(err)
	}

	transfer, err := transaction.NewTransfer(sndrPrivKey, rcvrPubKey, 14, 0.5, "trans1")
	if err != nil {
		t.Error(err)
	}

	transferPb, err := converter.TransactionLocal2Proto(transfer)
	if err != nil {
		t.Error(err)
	}

	err = compareTransactions(transfer, transferPb)
	if err != nil {
		t.Error(err)
	}

}

// TODO: add TestTransLocal2ProtoAirdrop

func TestTransProto2LocalTransfer(t *testing.T) {
	sndrPrivKey, sndrPubKey, err := generateWallet()
	if err != nil {
		t.Error(err)
	}

	_, rcvrPubKey, err := generateWallet()
	if err != nil {
		t.Error(err)
	}
	ti := time.Now()
	trTimestamp := time.Date(ti.Year(), ti.Month(), ti.Day(), ti.Hour(), ti.Minute(), ti.Second(), ti.Nanosecond(), time.UTC)

	data := interfaces.Data{
		Type: transaction.TypeTransfer,
		Sender: sndrPubKey,
		Receiver: rcvrPubKey,
		Pay: 156.4,
		Fee: 13.225,
		Message: "tr message",
		Timestamp: trTimestamp,
	}


	hashed := transaction.GetHash(&data)

	sign, err := cr.Sign(hashed, sndrPrivKey)
	if err != nil {
		t.Error(err)
	}

	data.Sign = sign

	tr := pb.Transaction{}
	tr.Type = pb.Transaction_Type(data.Type)
	tr.Sender = sndrPubKey
	tr.Receiver = rcvrPubKey
	tr.Message = data.Message
	tr.Pay = data.Pay
	tr.Fee = data.Fee
	tr.Timestamp = timestamppb.New(trTimestamp)
	tr.SenderSign = data.Sign

	var localTr interfaces.Transferable = &transaction.Transfer{Data: data}
	err = compareTransactions(localTr, &tr)
	if err != nil {
		t.Error(err)
	}
}

func TestBlockProto2Local(t *testing.T) {
	// TODO: finish
}

func TestBlockLocal2Proto(t *testing.T) {
	// TODO: finish
}
