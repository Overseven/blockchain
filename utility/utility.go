package utility

import (
	"crypto/ecdsa"
	"encoding/binary"
	"math"
	"time"

	cr "github.com/ethereum/go-ethereum/crypto"
)

func Int64FromBytes(bytes []byte) uint64 {
	return uint64(binary.LittleEndian.Uint64(bytes))
}

func UInt64Bytes(value uint64) []byte {
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, value)
	return bytes
}

func Float64FromBytes(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)
	float := math.Float64frombits(bits)
	return float
}

func Float64Bytes(float float64) []byte {
	return UInt64Bytes(math.Float64bits(float))
}

func PrivToPubKey(sndrPrivKey *ecdsa.PrivateKey) []byte {
	return cr.CompressPubkey(&sndrPrivKey.PublicKey)
}

func GenerateWallet() (privKey *ecdsa.PrivateKey, pubKey []byte, err error) {
	privKey, err = cr.GenerateKey()
	if err != nil {
		return nil, nil, err
	}

	pubKey = PrivToPubKey(privKey)
	return
}

func NewTimestamp() time.Time {
	t := time.Now()
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.UTC)
}