package chainimpl

import (
	chain "github.com/Overseven/blockchain/chain"
)

type Chain struct {
	Blocks []chain.Block
}

func (c *Chain) IsValid(startIndx, endIndx uint64) (bool, uint64) {

	// TODO: finish
	return true, 0
}

func (c *Chain) GetBlocks() []chain.Block {
	return c.Blocks
}
