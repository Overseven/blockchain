package network

import (
	"crypto/ecdsa"
	"net"
	"sync"
)

type Wallet struct {
	PrivKey *ecdsa.PrivateKey
	PubKey  []byte
}

type NodesContainer struct {
	Coordinator string

	Nodes map[string]interface{}
	Mutex sync.Mutex
}

type ServerParams struct {
	ListeningPort uint64
	OwnAddress    net.Addr
}

func (n *NodesContainer) ToStrings() []string {
	n.Mutex.Lock()
	defer n.Mutex.Unlock()

	var res []string

	for key := range n.Nodes {
		res = append(res, key)
	}
	return res
}
