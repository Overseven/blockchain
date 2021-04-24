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

func NewTimestamp() time.Time {
	t := time.Now()
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.UTC)
}

func MapDifference(first, second map[string]interface{}) map[string]interface{} {
	if first == nil || second == nil {
		return nil
	}

	diff := map[string]interface{}{}
	for k := range first {
		if _, ok := second[k]; !ok {
			diff[k] = struct{}{}
		}
	}

	return diff
}

// Find takes a slice and looks for an element in it. If found it will
// return it's key, otherwise it will return -1 and a bool of false.
func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}
