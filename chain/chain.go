package chain

import (
	"github.com/overseven/blockchain/chain/ichain"
)

var blockchain Chain

type Chain struct {
	Blocks []ichain.IBlock
}

func (c *Chain) IsValid(startIndx, endIndx uint64) (bool, uint64) {

	// TODO: finish
	return true, 0
}

func (c *Chain) GetBlocks() []ichain.IBlock {
	return c.Blocks
}

func (c *Chain) SetBlocks([]ichain.IBlock)  {

}

func GetBlockchain() Chain{
	return blockchain
}