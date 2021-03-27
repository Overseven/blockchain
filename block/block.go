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
	Transactions []interfaces.Transferable
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
		toHashBytes = append(toHashBytes, transaction.GetHash(tran.GetData())...)
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
	block.Hash = hash
	return hash
}

func (block *Block) IsValid(blockchain interfaces.Chainable, balance interfaces.Balancer) (bool, error) {
	// TODO: finish him!!

	var blocks []interfaces.Blockable = blockchain.GetBlocks()

	if uint64(len(blockchain.GetBlocks())) < block.Id+1 {
		return false, errors.New("incorrect block Id: " + strconv.FormatUint(block.Id, 10))
	}

	if block.Id != 0 {
		if blocks[block.Id-1].GetId()+1 != block.Id {
			return false, errors.New("conflicting block Id with prev. block Id: " + strconv.FormatUint(block.Id, 10))
		}
	}
	if uint64(len(blocks)) > block.Id+1 {
		if blocks[block.Id+1].GetId() != block.Id+1 {
			return false, errors.New("conflicting block Id with next block Id: " + strconv.FormatUint(block.Id, 10))
		}
	}

	// if block is the first in chain
	for _, t := range block.Transactions {
		data := t.GetData()
		if block.Id == 0 && data.Type != transaction.TypeAirdrop {
			return false, errors.New("first block must have only airdrop transactions")
		}
		if err := t.Verify(balance); err != nil {
			return false, errors.New("not valid transaction: " + err.Error())
		}
	}

	// TODO: finish for not first block

	//block.WalletsStats
	return true, nil
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
			return nonce
		}
	}

	// Not found :(
	return []byte{}
}

func (block *Block) GetTransactions() []interfaces.Transferable {
	return block.Transactions
}

func (block *Block) SetTransactions(tr []interfaces.Transferable) {
	block.Transactions = tr
}

func (block *Block) HasTransaction(transact interfaces.Transferable) (index int, has bool) {
	for i, tran := range block.Transactions {
		if transaction.IsEqual(transact.GetData(), tran.GetData()) {
			return i, true
		}
	}
	return 0, false
}

func (block *Block) AddTransaction(tr interfaces.Transferable) error {
	block.Transactions = append(block.Transactions, tr)

	return nil
}

func (block *Block) GetId() uint64 {
	return block.Id
}

func (block *Block) SetId(id uint64) {
	block.Id = id
}

func (block *Block) GetMiner() []byte {
	return block.Miner
}

func (block *Block) SetMiner(m []byte) {
	block.Miner = m
}

func (block *Block) GetPrevHash() []byte {
	return block.PrevHash
}

func (block *Block) SetPrevHash(ph []byte) {
	block.PrevHash = ph
}

func (block *Block) GetDifficulty() uint64 {
	return block.Difficulty
}

func (block *Block) SetDifficulty(d uint64) {
	block.Difficulty = d
}

func (block *Block) GetNonce() []byte {
	return block.Nonce
}

func (block *Block) SetNonce(n []byte) {
	block.Nonce = n
}