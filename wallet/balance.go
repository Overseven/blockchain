package wallet

import (
	"errors"
	"github.com/overseven/blockchain/chain"
	"github.com/overseven/blockchain/chain/ichain"
	"github.com/overseven/blockchain/transaction/itransaction"
	"strconv"
	"sync"
)

type Balance struct {
	Pubkey         []byte
	LastTransBlock uint64
	CurrentBalance float64
}

var (
	mutex         sync.Mutex
	usersBalances map[string]Balance
)

func Init() {
	if len(usersBalances) < 1 {
		usersBalances = make(map[string]Balance)
	}
}

func IsBeing(pubkey []byte) bool {
	_, ok := usersBalances[string(pubkey)]
	return ok
}

func Info(pubkey []byte) (Balance, error) {
	value, ok := usersBalances[string(pubkey)]
	if !ok {
		return Balance{}, errors.New("wallet information is not found")
	}
	return value, nil
}

func Update(pubkey []byte, lastTransBlock uint64, sum float64) (isNew bool) {
	mutex.Lock()
	defer mutex.Unlock()

	// TODO: add pubkey check

	if w, ok := usersBalances[string(pubkey)]; ok {
		w.LastTransBlock = lastTransBlock
		w.CurrentBalance = sum
		return false
	}

	newWallet := Balance{Pubkey: pubkey, LastTransBlock: lastTransBlock, CurrentBalance: sum}
	usersBalances[string(pubkey)] = newWallet

	return true
}

func Clear(){
	usersBalances = make(map[string]Balance)
}

func FullCalc() error{
	// TODO : finish
	blockchain := chain.GetBlockchain()
	Clear()

	for _, block := range blockchain.GetBlocks(){
		//b := block.(chain.Block)
		//c := ichain.IChain(*chain)
		if _, err := block.IsValid(blockchain); err != nil {
			Clear()
			return errors.New("incorrect block with number: " + strconv.FormatUint(block.GetId(), 10))
		}
		for _, trans := range block.GetTransaction(){
			data := trans.GetData()
			// sender
			sndrBalance := usersBalances[string(data.Pubkey)]
			sndrBalance.LastTransBlock = block.GetId()
			sndrBalance.CurrentBalance -= data.Pay + data.Fee

			// receiver
			rcvrBalance := usersBalances[string(data.Pubkey)]
			rcvrBalance.LastTransBlock = block.GetId()
			rcvrBalance.CurrentBalance += data.Pay

			// miner fee
			if data.Type != itransaction.Airdrop {
				sndrData := trans.GetData()
				balance := usersBalances[string(sndrData.Pubkey)]
				balance.LastTransBlock = block.GetId()
				balance.CurrentBalance += sndrData.Fee
			}
		}
	}

	return nil
}
