package tryblock

import (
	"bytes"
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/overseven/try-network/transaction"

	cr "github.com/ethereum/go-ethereum/crypto"
	blUtility "github.com/overseven/try-network/utility"
)

type TryBlock struct {
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

func (block *TryBlock) GetBatchHash() (hash []byte, err error) {
	var toHashBytes []byte
	for _, tran := range block.Transactions {
		tHash, err := tran.Hash(map[transaction.TransFlag]bool{})
		if err != nil {
			return []byte{}, err
		}
		toHashBytes = append(toHashBytes, tHash...)
	}
	hash = cr.Keccak256(toHashBytes)
	return
}

func (block *TryBlock) GetHash() (hash []byte, err error) {
	hash = blUtility.UInt64Bytes(block.Id)
	bHash, err := block.GetBatchHash()
	if err != nil {
		return nil, err
	}
	hash = append(hash, bHash...)
	hash = append(hash, block.PrevHash...)
	//hash = append(hash, block.GetWalletStatsHash()...)
	hash = append(hash, blUtility.UInt64Bytes(block.Difficulty)...)
	hash = append(hash, block.Miner...)
	hash = cr.Keccak256(hash)
	block.Hash = hash
	return hash, nil
}

func (block *TryBlock) IsValid() error {
	// check hash
	blockHash, err := block.GetHash()
	if err != nil {
		return err
	}
	hash := cr.Keccak256(append(blockHash, block.Nonce...))
	mask := make([]byte, block.Difficulty)
	if !bytes.HasPrefix(hash, mask) { // TODO: add hash difficult logic
		return errors.New("incorrect resulting hash, not have required zeroes")
	}
	return nil
}

func (block *TryBlock) Mining(minerPubKey []byte, stop chan bool) ([]byte, error) {
	block.Miner = minerPubKey
	mask := make([]byte, block.Difficulty)
	for i := uint64(0); i < math.MaxUint64; i++ {
		select {
		case <-stop:
			fmt.Println("Interrupt of mining.")
			return []byte{}, nil
		default:

		}
		nonce := []byte(strconv.FormatUint(i, 10))
		blockHash, err := block.GetHash()
		if err != nil {
			return nil, err
		}
		tryData := append(blockHash, nonce...)
		hash := cr.Keccak256(tryData)

		if bytes.HasPrefix(hash, mask) {
			// YEA!! Finish work
			block.Nonce = nonce
			return nonce, nil
		}
	}

	// Not found :(
	return nil, errors.New("not found")
}

func (block *TryBlock) HasTransaction(tr transaction.Transaction) (bool, error) {
	tHash, err := tr.Hash(map[transaction.TransFlag]bool{})
	if err != nil {
		return false, err
	}
	_, ok := block.Transactions[string(tHash)]
	return ok, nil
}

func (block *TryBlock) AddTransaction(tr transaction.Transaction) error {
	tHash, err := tr.Hash(map[transaction.TransFlag]bool{})
	if err != nil {
		return err
	}
	block.Transactions[string(tHash)] = tr

	return nil
}
