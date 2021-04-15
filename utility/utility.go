package utility

import (
	"crypto/ecdsa"
	"encoding/binary"
	"errors"
	"math"
	"strings"
	"time"

	cr "github.com/ethereum/go-ethereum/crypto"
)

func UInt64FromBytes(bytes []byte) uint64 {
	return uint64(binary.LittleEndian.Uint64(bytes))
}

func UInt64Bytes(value uint64) []byte {
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, value)
	return bytes
}

func UInt32FromBytes(bytes []byte) uint32 {
	return uint32(binary.LittleEndian.Uint32(bytes))
}

func UInt32Bytes(value uint32) []byte {
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint32(bytes, value)
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

func SetSuffix(str string) string {
	if len(str) == 0 {
		return str
	}

	str = strings.TrimSuffix(str, "\n")
	str += "\n"
	return str
}

func StringToBytes(s string) []byte {
	messageLen := uint32(len([]byte(s)))
	res := UInt32Bytes(messageLen)
	res = append(res, s...)
	return res
}

func StringFromBytes(b []byte) (string, error) {
	if len(b) < 4 {
		return "", errors.New("incorrect input data")
	}

	messageLen := UInt32FromBytes(b[:4])
	if messageLen == 0 {
		return "", nil
	}
	res := string(b[5 : 5+messageLen])

	return res, nil
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
