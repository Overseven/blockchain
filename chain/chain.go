package chain

import "github.com/overseven/blockchain/interfaces"

var blockchain Chain

type Chain struct {
	Blocks []interfaces.Blockable
}

func (c *Chain) IsValid(startIndx, endIndx uint64) (bool, uint64) {

	// TODO: finish
	return true, 0
}

func (c *Chain) GetBlocks() []interfaces.Blockable {
	return c.Blocks
}

func (c *Chain) SetBlocks([]interfaces.Blockable) {

}

func GetBlockchain() *Chain {
	return &blockchain
}
