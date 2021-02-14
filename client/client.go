package client

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	tr "github.com/overseven/blockchain/transaction"
	wallet "github.com/overseven/blockchain/wallet"
	"encoding/base64"
	"bytes"
	"io/ioutil"
	cr "github.com/ethereum/go-ethereum/crypto"
)

func Run(configFile string) {

	test2(configFile)

}

func test2(configFile string)  {
	_, privkeyBytes := wallet.LoadFromFile(configFile)
	//fmt.Println("Len privKey:", len(privkeyBytes))
	privkey, err := cr.ToECDSA(privkeyBytes[:32])
	if err != nil {
		panic(err)
	}
	encodedStr := base64.StdEncoding.EncodeToString(privkey.D.Bytes())
	pubkey := cr.CompressPubkey(&privkey.PublicKey)
	fmt.Println("Public key:", base64.StdEncoding.EncodeToString(pubkey))
	fmt.Println("Private key:", encodedStr)
	fmt.Println("Test transaction:")
	valid := testTransaction(privkey)
	if valid {
		fmt.Println("Valid!")
	} else{
		fmt.Println("Invalid!")
	}

}

func test(configFile string){
	transaction := tr.Transaction{}
	privkey, err := cr.GenerateKey()
	if err != nil {
		panic(err)
	}
	privkey2, err := cr.GenerateKey()
	if err != nil {
		panic(err)
	}

	pubkey := cr.CompressPubkey(&privkey.PublicKey)
	pubkey2 := cr.CompressPubkey(&privkey2.PublicKey)

	fmt.Println("Public key:", base64.StdEncoding.EncodeToString(pubkey))
	encodedStr := base64.StdEncoding.EncodeToString(privkey.D.Bytes())
	fmt.Println("Private key:", encodedStr)
	//fmt.Println(encodedStr)

	transaction.Pubkey = pubkey
	transaction.Pay = 14
	transaction.Fee = 0.123
	transaction.Receiver = pubkey2
	transaction.Message = "memsage"
	{
		t := time.Now()
		transaction.Timestamp = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.UTC)
	}
	hashed := transaction.GetHash()

	//fmt.Println("Public key:", pubkey)

	sign, err := cr.Sign(hashed, privkey)
	if err != nil {
		panic(err)
	}

	transaction.Sign = sign
	//spew.Dump(transaction)

	jsonTr, err := json.Marshal(transaction)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonTr))

	reverseTr := tr.FromJSON(jsonTr)
	valid := reverseTr.Verify
	fmt.Println("Valid:", valid)
	url := "http://127.0.0.1:8090/test"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonTr))
	if err != nil {
		panic(err)
	}
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

func testTransaction(privkey *ecdsa.PrivateKey) bool {
	privkey2, err := cr.GenerateKey()
	if err != nil {
		panic(err)
	}

	pubkey := cr.CompressPubkey(&privkey.PublicKey)
	pubkey2 := cr.CompressPubkey(&privkey2.PublicKey)

	transaction := tr.Transaction{}
	transaction.Pubkey = pubkey
	transaction.Pay = 14
	transaction.Fee = 0.123
	transaction.Receiver = pubkey2
	transaction.Message = "memsage"
	{
		t := time.Now()
		transaction.Timestamp = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.UTC)
	}
	hashed := transaction.GetHash()

	//fmt.Println("Public key:", pubkey)

	sign, err := cr.Sign(hashed, privkey)
	if err != nil {
		panic(err)
	}

	transaction.Sign = sign
	//spew.Dump(transaction)

	jsonTr, err := json.Marshal(transaction)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonTr))

	reverseTr := tr.FromJSON(jsonTr)
	valid := reverseTr.Verify()
	return valid
}