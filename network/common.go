package network

import (
	"crypto/ecdsa"
	"net"
	"sync"
)

type NetParams struct {
	PrivKey *ecdsa.PrivateKey
	PubKey  []byte

	Coordinator string

	Nodes map[string]interface{}
	Mutex sync.Mutex
}

type ServerParams struct {
	ListeningPort uint64
	OwnAddress    net.Addr
}
