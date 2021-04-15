package trlists

import (
	"sync"

	"github.com/overseven/blockchain/interfaces"
)

var (
	first, second List
)

type List struct {
	trans map[uint64]interfaces.BlockElement
	sync.Mutex
}
