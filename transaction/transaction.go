package transaction

import (
	"crypto/ecdsa"
)

const (
	ByteLenType         = 1
	ByteLenPubKey       = 33
	ByteLenSign         = 65
	ByteLenBalance      = 8
	ByteLenVotingId     = 4
	ByteLenTransCounter = 8
	ByteLenParameter    = 2
	MaxByteLenMessage   = 64
)

// Type is type of transaction
type Type byte

const (
	TypeAirdrop    Type = iota
	TypeTransfer        // transfer from one wallet to another
	TypeVote            // wallet vote
	TypeVotingInit      // voting initial transaction
	TypeVotingFinish
)

// TransFlag default value = true
type TransFlag byte

const (
	FlagNode      TransFlag = iota // use Node public key to calc hash
	FlagTimestamp                  // use Timestamp to calc hash
)

type Transaction interface {
	Counter() TransCounter
	IsEqual(Transaction, map[TransFlag]bool) bool
	String() (string, error)
	Bytes() ([]byte, error)
	Hash(map[TransFlag]bool) ([]byte, error)
	SetNode([]byte) Transaction
	Sign(*ecdsa.PrivateKey, TransCounter) error
	Verify() error
}

// Balance - amount of tokens. Update utility functions if the type is changed
type Balance float64

// VotingId - unique id of voting. Update utility functions if the type is changed
type VotingId uint32

// TransCounter - unique value of each transaction from address. Update utility functions if the type is changed
type TransCounter uint64
