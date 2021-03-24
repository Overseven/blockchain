package chain

import (
	"errors"
	"strconv"
	"sync"

	"github.com/overseven/blockchain/interfaces"
	"github.com/overseven/blockchain/transaction"
)

var UsersBalance Balance

type Balance struct {
	usersBalances map[string]BalanceStat
	mutex         sync.Mutex
}

type BalanceStat struct {
	Pubkey         []byte
	LastTransBlock uint64
	CurrentBalance float64
}

func (b *Balance) Init() {
	if len(b.usersBalances) < 1 {
		b.usersBalances = make(map[string]BalanceStat)
	}
}

func (b *Balance) IsBeing(pubkey []byte) bool {
	_, ok := b.usersBalances[string(pubkey)]
	return ok
}

func (b *Balance) Info(pubkey []byte) (BalanceStat, error) {
	value, ok := b.usersBalances[string(pubkey)]
	if !ok {
		return BalanceStat{}, errors.New("wallet information is not found")
	}
	return value, nil
}

func (b *Balance) Update(pubkey []byte, lastTransBlock uint64, sum float64) (isNew bool) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	// TODO: add pubkey check

	if w, ok := b.usersBalances[string(pubkey)]; ok {
		w.LastTransBlock = lastTransBlock
		w.CurrentBalance = sum
		return false
	}

	newWallet := BalanceStat{Pubkey: pubkey, LastTransBlock: lastTransBlock, CurrentBalance: sum}
	b.usersBalances[string(pubkey)] = newWallet

	return true
}

func (b *Balance) Clear() {
	b.usersBalances = make(map[string]BalanceStat)
}

func (b *Balance) FullCalc(blockchain interfaces.Chainable) error {
	// TODO : finish

	b.Clear()

	for _, block := range blockchain.GetBlocks() {
		//b := block.(chain.Block)
		//c := ichain.IChain(*chain)
		if _, err := block.IsValid(blockchain); err != nil {
			b.Clear()
			return errors.New("incorrect block with number: " + strconv.FormatUint(block.GetId(), 10))
		}
		for _, trans := range block.GetTransaction() {
			data := trans.GetData()
			// sender
			sndrBalance := b.usersBalances[string(data.Pubkey)]
			sndrBalance.LastTransBlock = block.GetId()
			sndrBalance.CurrentBalance -= data.Pay + data.Fee

			// receiver
			rcvrBalance := b.usersBalances[string(data.Pubkey)]
			rcvrBalance.LastTransBlock = block.GetId()
			rcvrBalance.CurrentBalance += data.Pay

			// miner fee
			if data.Type != transaction.TypeAirdrop {
				sndrData := trans.GetData()
				balance := b.usersBalances[string(sndrData.Pubkey)]
				balance.LastTransBlock = block.GetId()
				balance.CurrentBalance += sndrData.Fee
			}
		}
	}

	return nil
}
