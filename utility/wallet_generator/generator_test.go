package main_test

import (
	"encoding/hex"
	"fmt"
	"testing"

	cr "github.com/ethereum/go-ethereum/crypto"
	"github.com/overseven/try-network/utility"
)

func TestGeneration(t *testing.T) {
	privHex, pubHex, err := utility.Generate()
	if err != nil {
		t.Error(err)
	}

	priv, err := utility.ParseKeys(privHex, pubHex)
	if err != nil {
		t.Error(err)
	}
	data := "I'm data!"
	dataHash := cr.Keccak256Hash([]byte(data))

	//pubAver := cr.FromECDSAPub(&priv.PublicKey)
	pubComp := cr.CompressPubkey(&priv.PublicKey)
	//pubComp = pubComp[:len(pubComp)-1]
	fmt.Println("PubHexComp:", hex.EncodeToString(pubComp))

	sign, err := cr.Sign(dataHash.Bytes(), priv)
	if err != nil {
		t.Error(err)
	}

	//sigPublicKey, err := cr.Ecrecover(dataHash.Bytes(), sign)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//matches := bytes.Equal(sigPublicKey, pubComp)
	//fmt.Println(matches) // true
	//
	//sigPublicKeyECDSA, err := cr.SigToPub(dataHash.Bytes(), sign)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//sigPublicKeyBytes := cr.FromECDSAPub(sigPublicKeyECDSA)
	//matches = bytes.Equal(sigPublicKeyBytes, pubComp)
	//fmt.Println(matches) // true

	signatureNoRecoverID := sign[:len(sign)-1] // remove recovery id
	verified := cr.VerifySignature(pubComp, dataHash.Bytes(), signatureNoRecoverID)
	fmt.Println(verified) // true

}

func TestParsing(t *testing.T) {
	privHex := "eb70805d4fa2ebe2ba199b13abe89f85015a3418932be4abf348218a703dba0c"
	pubHex := "024c3ba55861c27a3ac4edefd40ec49a7d66f03eb1c4bddfc67776bf1df3bdfeb0"

	_, err := utility.ParseKeys(privHex, pubHex)
	if err != nil {
		t.Error(err)
	}
}
