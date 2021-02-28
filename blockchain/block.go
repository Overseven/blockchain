package blockchain

import (
	"bytes"
	"errors"
	"fmt"
	cr "github.com/ethereum/go-ethereum/crypto"
	"github.com/overseven/blockchain/transaction"
	"github.com/overseven/blockchain/utility"
	"math"
	"strconv"
)

type Block struct{
	Id uint64
	Transactions []transaction.Transaction
	PrevHash     []byte
	WalletsStats map[string]WalletStats

	Difficulty uint64
	Miner []byte
	Hash  []byte
	Nonce []byte
}

type WalletStats struct{
	Address []byte
	BalanceBefore float64
	BalanceAfter float64
	PrevTransBlockId uint64
}

func (block *Block) GetBatchHash() (hash []byte){
	var toHashBytes []byte
	for _, tran := range block.Transactions {
		toHashBytes = append(toHashBytes, tran.GetHash()...)
	}
	hash = cr.Keccak256(toHashBytes)
	return
}

func (block *Block) GetWalletStatsHash() (hash []byte){
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

func (block *Block) GetHash() (hash []byte){
	hash = blUtility.UInt64Bytes(block.Id)
	hash = append(hash, block.GetBatchHash()...)
	hash = append(hash, block.PrevHash...)
	hash = append(hash, block.GetWalletStatsHash()...)
	hash = append(hash, blUtility.UInt64Bytes(block.Difficulty)...)
	hash = append(hash, block.Miner...)
	return hash
}

func (block *Block) IsValid(blockchain *Blockchain) (bool, error) {
	// TODO: finish him!!
	if uint64(len(blockchain.Blocks)) + 1 != block.Id {
		return false, errors.New("incorrect block ID")
	}
	//block.WalletsStats
	return true, nil
}

func (block *Block) Mining(stop chan bool) []byte {
	mask := make([]byte, block.Difficulty)
	for  i := uint64(0); i< math.MaxUint64; i++ {
		select{
		case <-stop:
			fmt.Println("Interrupt of mining.")
			return []byte{}
		default:

		}
		nonce := []byte(strconv.FormatUint(i, 10))
		tryData := append(block.Hash, nonce...)
		hash := cr.Keccak256(tryData)

		if bytes.HasPrefix(hash, mask){
			// YEA!! Finish work
			return nonce
		}
	}

	// Not found :(
	return []byte{}
}

func (block *Block)HasTransaction(tr *transaction.Transaction) (index int, has bool){
	for i, tran := range block.Transactions{
		if tran.IsEqual(tr){
			return i, true
		}
	}
	return 0,false
}

func (block *Block)AddTransaction(tr *transaction.Transaction) error{
	// TODO: Finish him!!
	return nil
}
