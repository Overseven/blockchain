package transaction

import (
	"crypto/ecdsa"
)

const (
	ByteLenPubKey     = 33
	ByteLenSign       = 65
	MaxByteLenMessage = 64
)

// Type is type of transaction
type Type byte

const (
	TypeAirdrop Type = iota
	TypeTransfer // transfer from one wallet to another
	TypeVote // wallet vote
	TypeVotingInit // voting initial transaction
)

// TransFlag default value = true
type TransFlag byte

const (
	FlagNode TransFlag = iota // use Node public key to calc hash
	FlagTimestamp // use Timestamp to calc hash
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
