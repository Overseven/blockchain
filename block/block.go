package block

import (
	"bytes"
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/overseven/blockchain/transaction"

	cr "github.com/ethereum/go-ethereum/crypto"
	blUtility "github.com/overseven/blockchain/utility"
)

type Block struct {
	Id           uint64
	Transactions map[string]transaction.Transaction
	PrevHash     []byte
	//WalletsStats map[string]WalletStats

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
		toHashBytes = append(toHashBytes, tran.Hash()...)
	}
	hash = cr.Keccak256(toHashBytes)
	return
}

func (block *Block) GetHash() (hash []byte) {
	hash = blUtility.UInt64Bytes(block.Id)
	hash = append(hash, block.GetBatchHash()...)
	hash = append(hash, block.PrevHash...)
	//hash = append(hash, block.GetWalletStatsHash()...)
	hash = append(hash, blUtility.UInt64Bytes(block.Difficulty)...)
	hash = append(hash, block.Miner...)
	hash = cr.Keccak256(hash)
	block.Hash = hash
	return hash
}

func (block *Block) IsValid() error {
	// check hash
	hash := cr.Keccak256(append(block.GetHash(), block.Nonce...))
	mask := make([]byte, block.Difficulty)
	if !bytes.HasPrefix(hash, mask) { // TODO: add hash difficult logic
		return errors.New("incorrect resulting hash, not have required zeroes")
	}
	return nil
}

func (block *Block) Mining(minerPubKey []byte, stop chan bool) []byte {
	block.Miner = minerPubKey
	mask := make([]byte, block.Difficulty)
	for i := uint64(0); i < math.MaxUint64; i++ {
		select {
		case <-stop:
			fmt.Println("Interrupt of mining.")
			return []byte{}
		default:

		}
		nonce := []byte(strconv.FormatUint(i, 10))
		tryData := append(block.GetHash(), nonce...)
		hash := cr.Keccak256(tryData)

		if bytes.HasPrefix(hash, mask) {
			// YEA!! Finish work
			block.Nonce = nonce
			return nonce
		}
	}

	// Not found :(
	return []byte{}
}

func (block *Block) HasTransaction(tr transaction.Transaction) bool {
	_, ok := block.Transactions[string(tr.Hash())]
	return ok
}

func (block *Block) AddTransaction(tr transaction.Transaction) error {
	block.Transactions[string(tr.Hash())] = tr

	return nil
}
