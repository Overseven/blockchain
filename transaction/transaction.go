package transaction

import (
	"crypto/ecdsa"
)

const (
	ByteLenPubKey     = 33
	ByteLenSign       = 65
	MaxByteLenMessage = 64
)

type Type byte

const (
	TypeAirdrop Type = iota
	TypeTransfer
	TypeVote
	TypeVotingInit
)

// default value = true
type TransFlag byte

const (
	FlagNode TransFlag = iota
	FlagTimestamp
)

type Transaction interface {
	IsEqual(Transaction, map[TransFlag]bool) bool
	String() (string, error)
	Bytes() ([]byte, error)
	Hash(map[TransFlag]bool) ([]byte, error)
	SetNode([]byte) Transaction
	Sign(*ecdsa.PrivateKey) error
	Verify() error
}
