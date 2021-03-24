package chain

import (
	"sync"

	"github.com/overseven/blockchain/block"
	"github.com/overseven/blockchain/interfaces"
)

type Chain struct {
	Blocks []interfaces.Blockable
	mutex  sync.Mutex
}

func (c *Chain) IsValid(startIndx, endIndx uint64) (bool, uint64) {

	// TODO: finish
	return true, 0
}

func (c *Chain) GetBlocks() []interfaces.Blockable {
	return c.Blocks
}

func (c *Chain) SetBlocks([]interfaces.Blockable) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	// TODO: finish
}

func (c *Chain) NewBlock() interfaces.Blockable {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if len(c.Blocks) == 0 {
		var bl interfaces.Blockable = &block.Block{Id: 0}
		c.Blocks = append(c.Blocks, bl)
		return bl
	}

	var bl interfaces.Blockable = &block.Block{Id: c.Blocks[len(c.Blocks)-1].GetId() + 1, PrevHash: c.Blocks[len(c.Blocks)-1].GetHash()}
	c.Blocks = append(c.Blocks, bl)
	return bl
}

func (c *Chain) AppendBlock(b interfaces.Blockable) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.Blocks = append(c.Blocks, b)
}
