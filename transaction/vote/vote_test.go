package vote_test

import (
    "bytes"
    "github.com/overseven/blockchain/transaction"
    "github.com/overseven/blockchain/transaction/vote"
    "github.com/overseven/blockchain/utility"
    "testing"
)

func TestNewVote(t *testing.T) {
    senderPrivKey, _, err := utility.GenerateWallet()
    if err != nil {
        t.Error(err)
    }

    _, nodePubKey, err := utility.GenerateWallet()
    if err != nil {
        t.Error(err)
    }

    var votingId uint64 = 115
    opinion:= "no"
    fee := 42.2222

    v, err := vote.NewVote(votingId, opinion, fee)
    if err != nil {
        t.Error(err)
    }

    v = v.SetNode(nodePubKey).(*vote.Vote)
    if err != nil {
        t.Error(err)
    }

    err = v.Sign(senderPrivKey)
    if err != nil {
        t.Error(err)
    }

    // fmt.Println(tr1.String())

    err = v.Verify()
    if err != nil {
        t.Error(err)
    }
}

func TestByteConversation(t *testing.T){
    senderPrivKey, _, err := utility.GenerateWallet()
    if err != nil {
        t.Error(err)
    }

    _, nodePubKey, err := utility.GenerateWallet()
    if err != nil {
        t.Error(err)
    }

    var votingId uint64 = 115
    opinion := "yes"
    fee := 42.2222

    v, err := vote.NewVote(votingId, opinion, fee)
    if err != nil {
        t.Error(err)
    }

    v = v.SetNode(nodePubKey).(*vote.Vote)
    if err != nil {
        t.Error(err)
    }

    err = v.Sign(senderPrivKey)
    if err != nil {
        t.Error(err)
    }

    // fmt.Println(tr1.String())

    err = v.Verify()
    if err != nil {
        t.Error(err)
    }

    b, err := v.Bytes()
    if err != nil {
        t.Error(err)
    }

    vReverse, err := vote.FromBytes(b)
    if err != nil {
        t.Error(err)
    }

    //s1, err := v.String()
    //if err != nil {
    //	t.Error(err)
    //}
    //
    //s2, err := vReverse.String()
    //if err != nil {
    //	t.Error(err)
    //}
    //
    //fmt.Println("orig Vote:", s1)
    //fmt.Println("Reverse Vote:", s2)

    h1, err := v.Hash(map[transaction.TransFlag]bool{})
    if err != nil {
        t.Error(err)
    }
    h2, err := vReverse.Hash(map[transaction.TransFlag]bool{})
    if err != nil {
        t.Error(err)
    }

    if !bytes.Equal(h1, h2) {
        t.Error("incorrect hash")
    }
}