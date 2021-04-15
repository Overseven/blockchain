package transaction

const (
	ByteLenPubKey     = 8
	ByteLenSign       = 8
	MaxByteLenMessage = 64
)

type Type byte

const (
	TypeAirdrop Type = iota
	TypeTransfer
)

type Transaction interface {
	String() string
	Bytes() []byte
	Hash() []byte
	Verify() error
}
