package utility

import (
	"crypto/ecdsa"
	"encoding/binary"
	"errors"
	"math"
	"strings"
	"time"

	cr "github.com/ethereum/go-ethereum/crypto"
	"github.com/overseven/blockchain/transaction"
)

const (
	TimestampFormat = "02 Jan 06 15:04 MST"
)

func UInt64FromBytes(bytes []byte) uint64 {
	return binary.LittleEndian.Uint64(bytes)
}

func UInt64Bytes(value uint64) []byte {
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, value)
	return bytes
}

func UInt32FromBytes(bytes []byte) uint32 {
	return binary.LittleEndian.Uint32(bytes)
}

func UInt32Bytes(value uint32) []byte {
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, value)
	return bytes
}
func UInt16FromBytes(bytes []byte) uint16 {
	return binary.LittleEndian.Uint16(bytes)
}

func UInt16Bytes(value uint16) []byte {
	bytes := make([]byte, 2)
	binary.LittleEndian.PutUint16(bytes, value)
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

func StringFromBytes(b []byte) (string, uint32, error) {
	if len(b) < 4 {
		return "", 0, errors.New("incorrect input data")
	}

	messageLen := UInt32FromBytes(b[:4])
	if messageLen == 0 {
		return "", 0, nil
	}
	if messageLen > transaction.MaxByteLenMessage {
		return "", 0, errors.New("message is too large")
	}
	res := string(b[4 : 4+messageLen])

	return res, 4 + messageLen , nil
}

func NewTimestamp() time.Time {
	t := time.Now()
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.UTC)
}

func TimestampToBytes(t time.Time) ([]byte, error) {
	ts := []byte(t.Format(TimestampFormat))
	res := []byte{uint8(len(ts))}
	res = append(res, []byte(t.Format(TimestampFormat))...)
	return res, nil
}

func TimestampFromBytes(b []byte) (time.Time, uint8, error) {
	if len(b) == 0 {
		return time.Time{}, 0, errors.New("empty bytes")
	}
	length := b[0]

	if length == 0 {
		return time.Time{}, 0, errors.New("empty timestamp")
	}
	if len(b)-1 < int(length) {
		return time.Time{}, 0, errors.New("length timestamp > len slice")
	}
	t, err := time.Parse(TimestampFormat, string(b[1:length+1]))

	return t, length + 1, err
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
