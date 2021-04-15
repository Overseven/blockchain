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
	String() (string, error)
	Bytes() ([]byte, error)
	FromBytes([]byte) error
	Hash() []byte
	Verify() error
}
