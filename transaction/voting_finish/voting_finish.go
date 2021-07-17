package voting_finish

import (
    "bytes"
    "crypto/ecdsa"
    "encoding/json"
    "errors"
    cr "github.com/ethereum/go-ethereum/crypto"
    "strconv"
    "time"

    "github.com/overseven/try-network/transaction"
    "github.com/overseven/try-network/utility"
)

type VotingFinish struct {
    Sender    	 []byte
    TransCounter uint32
    VotingId  	 uint64
    Fee       	 float64
    Timestamp 	 time.Time
    Node      	 []byte
    Signature 	 []byte
}

func NewVotingFinish(votingId uint64, fee float64) (*VotingFinish, error) {
    tr := VotingFinish{
        VotingId: votingId,
        Fee:      fee,
    }

    tr.Timestamp = utility.NewTimestamp()
    return &tr, nil
}

func (v *VotingFinish) IsEqual(tr transaction.Transaction, flags map[transaction.TransFlag]bool) bool {
    if flags == nil {
        return false
    }

    switch v2 := tr.(type) {
    case *VotingFinish:
        if !bytes.Equal(v.Sender, v2.Sender) {
            return false
        }

        if v.VotingId != v2.VotingId {
            return false
        }
        flagTimestamp, ok := flags[transaction.FlagTimestamp]
        if !ok || (ok && flagTimestamp) {
            if v.Timestamp != v2.Timestamp {
                return false
            }
        }

        flagNode, ok := flags[transaction.FlagNode]
        if !ok || (ok && flagNode) {
            if !bytes.Equal(v.Node, v2.Node) {
                return false
            }
        }

        if err := v.Verify(); err != nil {
            return false
        }
        if err := v2.Verify(); err != nil {
            return false
        }

        if v.Fee != v2.Fee {
            return false
        }
        return true
    default:
        return false
    }
}

func (v *VotingFinish) String() (string, error) {
    tmp, err := json.Marshal(v)
    if err != nil {
        return "", err
    }

    return string(tmp), nil
}

func (v *VotingFinish) Bytes() ([]byte, error) {
    var res []byte

    res = append(res, byte(transaction.TypeVotingFinish))

    if len(v.Sender) != transaction.ByteLenPubKey {
        return nil, errors.New("incorrect sender field size")
    }
    res = append(res, v.Sender...)
    res = append(res, utility.UInt32Bytes(v.TransCounter)...)
    res = append(res, utility.UInt64Bytes(v.VotingId)...)

    res = append(res, utility.Float64Bytes(v.Fee)...)

    timestamp, err := utility.TimestampToBytes(v.Timestamp)
    if err != nil {
        return nil, err
    }
    res = append(res, timestamp...)

    if len(v.Node) != transaction.ByteLenPubKey {
        return nil, errors.New("incorrect node field size")
    }
    res = append(res, v.Node...)

    if len(v.Signature) != transaction.ByteLenSign {
        return nil, errors.New("incorrect sign field size")
    }

    res = append(res, v.Signature...)

    return res, nil
}

func FromBytes(b []byte) (*VotingFinish, error) {
    if len(b) < 64 { // TODO: define min size
        return nil, errors.New("incorrect input data len")
    }
    v := new(VotingFinish)
    idx := int64(0)
    typeTr := transaction.Type(b[idx])
    if typeTr != transaction.TypeVotingFinish {
        return nil, errors.New("incorrect transaction type")
    }
    idx += 1
    v.Sender = b[idx : idx+transaction.ByteLenPubKey]
    idx += transaction.ByteLenPubKey
    v.TransCounter = utility.UInt32FromBytes(b[idx : idx+4])
    idx += 4
    v.VotingId = utility.UInt64FromBytes(b[idx : idx+8])
    idx += 8

    v.Fee = utility.Float64FromBytes(b[idx : idx+8])
    idx += 8

    timestamp, timestampLen, err := utility.TimestampFromBytes(b[idx:])
    if err != nil {
        return nil, err
    }
    v.Timestamp = timestamp
    idx += int64(timestampLen)

    v.Node = b[idx : idx+transaction.ByteLenPubKey]
    idx += transaction.ByteLenPubKey
    v.Signature = b[idx : idx+transaction.ByteLenSign]
    idx += transaction.ByteLenSign
    return v, nil
}

func (v *VotingFinish) Hash(flags map[transaction.TransFlag]bool) ([]byte, error) {
    if flags == nil {
        return nil, errors.New("empty flags")
    }
    var temp []byte
    temp = append(temp, v.Sender...)
    temp = append(temp, strconv.FormatUint(uint64(v.TransCounter), 10)...)
    temp = append(temp, utility.UInt64Bytes(v.VotingId)...)
    temp = append(temp, strconv.FormatFloat(v.Fee, 'e', 8, 64)...)
    flagTimestamp, ok := flags[transaction.FlagTimestamp]
    if !ok || (ok && flagTimestamp) {
        temp = append(temp, v.Timestamp.Format(utility.TimestampFormat)...)
    }

    flagNode, ok := flags[transaction.FlagNode]
    if !ok || (ok && flagNode) {
        temp = append(temp, v.Node...)
    }
    return cr.Keccak256(temp), nil
}

func (v *VotingFinish) Verify() error {
    hash, err := v.Hash(map[transaction.TransFlag]bool{})
    if err != nil {
        return err
    }
    if len(v.Signature) < 64 {
        return errors.New("incorrect signature len: " + strconv.FormatInt(int64(len(v.Signature)), 10))
    }
    if !cr.VerifySignature(v.Sender, hash, v.Signature[0:64]) {
        return errors.New("incorrect signature")
    }
    return nil
}

func Copy(v *VotingFinish) *VotingFinish {
    res := new(VotingFinish)
    res.Sender = v.Sender
    res.TransCounter = v.TransCounter
    res.VotingId = v.VotingId
    res.Fee = v.Fee
    res.Timestamp = v.Timestamp
    res.Node = v.Node
    res.Signature = v.Signature
    return res
}

func (v *VotingFinish) SetNode(nodePubKey []byte) transaction.Transaction {
    res := Copy(v)
    res.Node = nodePubKey
    return res
}

func (v *VotingFinish) Sign(privKey *ecdsa.PrivateKey, transCounter uint32) error {
    senderPubKey := utility.PrivToPubKey(privKey)
    v.Sender = senderPubKey
    v.TransCounter = transCounter
    hashed, err := v.Hash(map[transaction.TransFlag]bool{})
    if err != nil {
        return err
    }

    sign, err := cr.Sign(hashed, privKey)
    if err != nil {
        return err
    }

    v.Signature = sign
    return nil
}

func (v *VotingFinish) Counter() uint32 {
    return v.TransCounter
}