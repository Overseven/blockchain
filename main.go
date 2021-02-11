package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	tr "github.com/overseven/blockchain/transaction"

	cr "github.com/ethereum/go-ethereum/crypto"
)

func main() {
	transaction := tr.Transaction{}
	privkey, err := cr.GenerateKey()
	if err != nil {
		panic(err)
	}
	privkey2, err := cr.GenerateKey()
	if err != nil {
		panic(err)
	}
	//fmt.Println("Private key:", privkey.D)

	pubkey := cr.CompressPubkey(&privkey.PublicKey)
	pubkey2 := cr.CompressPubkey(&privkey2.PublicKey)
	transaction.Pubkey = pubkey
	transaction.Pay = 14
	transaction.Fee = 0.123
	transaction.Receiver = pubkey2
	transaction.Message = "memsage"
	{
		t := time.Now()
		transaction.Timestamp = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.UTC)
	}
	hashed := tr.GetHash(&transaction)

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
	valid := tr.Verify(&reverseTr)
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
