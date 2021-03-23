package chain

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/overseven/blockchain/chain/ichain"
	"github.com/overseven/blockchain/transaction/itransaction"
	"math"
	"strconv"

	blUtility "github.com/Overseven/blockchain/utility"
	cr "github.com/ethereum/go-ethereum/crypto"
)

type Block struct {
	Id           uint64
	Transactions []itransaction.ITransaction
	PrevHash     []byte
	WalletsStats map[string]WalletStats

	Difficulty uint64
	Miner      []byte
	Hash       []byte
	Nonce      []byte
}

type WalletStats struct {
	Address          []byte
	BalanceBefore    float64
	BalanceAfter     float64
	PrevTransBlockId uint64
}

func (block *Block) GetBatchHash() (hash []byte) {
	var toHashBytes []byte
	for _, tran := range block.Transactions {
		toHashBytes = append(toHashBytes, itransaction.GetHash(tran.GetData())...)
	}
	hash = cr.Keccak256(toHashBytes)
	return
}

func (block *Block) GetWalletStatsHash() (hash []byte) {
	var toHashBytes []byte
	for _, stats := range block.WalletsStats {
		toHashBytes = append(toHashBytes, stats.Address...)
		toHashBytes = append(toHashBytes, blUtility.Float64Bytes(stats.BalanceBefore)...)
		toHashBytes = append(toHashBytes, blUtility.Float64Bytes(stats.BalanceAfter)...)
		toHashBytes = append(toHashBytes, blUtility.UInt64Bytes(stats.PrevTransBlockId)...)
	}
	hash = cr.Keccak256(toHashBytes)
	return
}

func (block *Block) GetHash() (hash []byte) {
	hash = blUtility.UInt64Bytes(block.Id)
	hash = append(hash, block.GetBatchHash()...)
	hash = append(hash, block.PrevHash...)
	hash = append(hash, block.GetWalletStatsHash()...)
	hash = append(hash, blUtility.UInt64Bytes(block.Difficulty)...)
	hash = append(hash, block.Miner...)
	return hash
}

func (block *Block) IsValid(blockchain ichain.IChain) (bool, error) {
	// TODO: finish him!!
	if uint64(len(blockchain.GetBlocks()))+1 != block.Id {
		return false, errors.New("incorrect block ID")
	}

	// if block is the first in chain
	if len(blockchain.GetBlocks()) == 0 {
		for _, t := range block.Transactions {
			data := t.GetData()
			if data.Type != itransaction.Airdrop {
				return false, errors.New("first block must have only airdrop transactions")
			}
			if err := t.Verify(); err != nil {
				return false, errors.New("not valid transaction: " + err.Error())
			}
		}
		return true, nil
	}

	// TODO: finish for not first block

	//block.WalletsStats
	return true, nil
}

func (block *Block) Mining(stop chan bool) []byte {
	mask := make([]byte, block.Difficulty)
	for i := uint64(0); i < math.MaxUint64; i++ {
		select {
		case <-stop:
			fmt.Println("Interrupt of mining.")
			return []byte{}
		default:

		}
		nonce := []byte(strconv.FormatUint(i, 10))
		tryData := append(block.Hash, nonce...)
		hash := cr.Keccak256(tryData)

		if bytes.HasPrefix(hash, mask) {
			// YEA!! Finish work
			return nonce
		}
	}

	// Not found :(
	return []byte{}
}

func (block *Block) GetTransaction() []itransaction.ITransaction{
	return block.Transactions
}


func (block *Block) HasTransaction(transact *itransaction.ITransaction) (index int, has bool) {
	for i, tran := range block.Transactions {
		if itransaction.IsEqual((*transact).GetData(), tran.GetData()) {
			return i, true
		}
	}
	return 0, false
}

func (block *Block) AddTransaction(tr *itransaction.ITransaction) error {
	// TODO: Finish him!!
	return nil
}

func (block *Block) GetId() uint64{
	return block.Id
}
