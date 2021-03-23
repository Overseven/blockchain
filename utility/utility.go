package utility

import (
	"crypto/ecdsa"
	"encoding/binary"
	"math"

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
