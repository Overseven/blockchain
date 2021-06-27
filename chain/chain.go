package chain

/*
import (
	"sync"

	"github.com/overseven/try-network/block"
)

type Chain struct {
	Blocks []interfaces.TransactionsContainer
	mutex  sync.Mutex
}

func (c *Chain) IsValid(startIndx, endIndx uint64) (bool, uint64) {

	// TODO: finish
	return true, 0
}

func (c *Chain) GetBlocks() []interfaces.TransactionsContainer {
	return c.Blocks
}

func (c *Chain) SetBlocks([]interfaces.TransactionsContainer) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	// TODO: finish
}

func (c *Chain) NewBlock() interfaces.TransactionsContainer {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if len(c.Blocks) == 0 {
		var bl interfaces.TransactionsContainer = &block.Block{Id: 0}
		c.Blocks = append(c.Blocks, bl)
		return bl
	}

	var bl interfaces.TransactionsContainer = &block.Block{Id: c.Blocks[len(c.Blocks)-1].GetId() + 1, PrevHash: c.Blocks[len(c.Blocks)-1].GetHash()}
	c.Blocks = append(c.Blocks, bl)
	return bl
}

func (c *Chain) AppendBlock(b interfaces.TransactionsContainer) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.Blocks = append(c.Blocks, b)
}

*/
