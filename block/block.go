package block

import (
	"bytes"
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/overseven/blockchain/interfaces"
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

func (block *Block) IsValid(blockchain interfaces.BlockConnecter, balance interfaces.Balancer) error {
	// TODO: add test
	var blocks []interfaces.TransactionsContainer = blockchain.GetBlocks()

	if uint64(len(blockchain.GetBlocks())) < block.Id+1 {
		return errors.New("incorrect block Id: " + strconv.FormatUint(block.Id, 10))
	}

	if block.Id != 0 {
		if blocks[block.Id-1].GetId()+1 != block.Id {
			return errors.New("conflicting block Id with prev. block Id: " + strconv.FormatUint(block.Id, 10))
		}
	}
	// if uint64(len(blocks)) > block.Id+1 {
	// 	if blocks[block.Id+1].GetId() != block.Id+1 {
	// 		return errors.New("conflicting block Id with next block Id: " + strconv.FormatUint(block.Id, 10))
	// 	}
	// }

	// if block is the first in chain
	for _, t := range block.Transactions {
		data := t.GetData()
		if block.Id == 0 && data.Type != transaction.TypeAirdrop {
			return errors.New("first block must have only airdrop transactions")
		}
		if err := t.Verify(balance); err != nil {
			return errors.New("not valid transaction: " + err.Error())
		}
	}

	// check hash
	hash := cr.Keccak256(append(block.GetHash(), block.Nonce...))
	mask := make([]byte, block.Difficulty)
	if !bytes.HasPrefix(hash, mask) {
		return errors.New("incorrect resulting hash, not have required zeroes")
	}

	//block.WalletsStats
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

func (block *Block) HasTransaction(transact transaction.Transaction) (index int, has bool) {
	for i, tran := range block.Transactions {
		if transaction.IsEqual(transact.GetData(), tran.GetData(), true) {
			return i, true
		}
	}
	return 0, false
}

func (block *Block) AddTransaction(tr transaction.Transaction) error {
	block.Transactions = append(block.Transactions, tr)

	return nil
}
