package utility

import (
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	cr "github.com/ethereum/go-ethereum/crypto"
)

func ToHex(key *ecdsa.PrivateKey) (hexPriv, hexPubCompressed string, err error) {
	privStr := hex.EncodeToString(cr.FromECDSA(key))
	pubStr  := hex.EncodeToString(cr.CompressPubkey(&key.PublicKey))
	return privStr, pubStr, nil
}

func Generate() (hexPriv, hexPubCompressed string, err error){
	priv, err := cr.GenerateKey()
	if err != nil {
		return "", "", err
	}
	return ToHex(priv)
}

func ParseKeys(hexPriv, hexPubCompressed string) (privateKey *ecdsa.PrivateKey, err error){
	if hexPubCompressed == "" {
		return nil, errors.New("empty public key")
	}

	priv, err := cr.HexToECDSA(hexPriv)
	if err != nil {
		fmt.Println("Private key decode error!", err.Error())
		return nil, err
	}
	//privDecStr := hex.EncodeToString(cr.FromECDSA(privDec))

	_, pubFromPriv, err := ToHex(priv)
	if err != nil {
		fmt.Println("Error!", err.Error())
		return nil, err
	}
	if pubFromPriv != hexPubCompressed {
		fmt.Println("The public keys from function arg and from private key are different!")
		return nil, err
	}

	return priv,nil
}

func GenerateWallet() (privKey *ecdsa.PrivateKey, pubKey []byte, err error) {
	privKey, err = cr.GenerateKey()
	if err != nil {
		return nil, nil, err
	}

	pubKey = PrivToPubKey(privKey)
	return
}