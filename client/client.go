package client

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	cr "github.com/ethereum/go-ethereum/crypto"

	"github.com/overseven/blockchain/balance"
	"github.com/overseven/blockchain/chain"
	tr "github.com/overseven/blockchain/transaction"
	wallet "github.com/overseven/blockchain/wallet"
)

var (
	usersBalance balance.Balance
	bchain       chain.Chain
	address      = "localhost:50051"
)

func Run(configFile string) {
	// TODO: create Balance
	usersBalance.Init()
	test2(configFile)
}

func test2(configFile string) {
	_, privkeyBytes, err := wallet.LoadFromFile(configFile)
	if err != nil {
		panic(err)
	}
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
	} else {
		fmt.Println("Invalid!")
	}

}

func test(configFile string) {
	transaction := tr.Transfer{}
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

	data := transaction.GetData()
	data.Sender = pubkey
	data.Pay = 14
	data.Fee = 0.123
	data.Receiver = pubkey2
	data.Message = "memsage"
	{
		t := time.Now()
		data.Timestamp = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.UTC)
	}
	hashed := tr.GetHash(data)

	//fmt.Println("Public key:", pubkey)

	sign, err := cr.Sign(hashed, privkey)
	if err != nil {
		panic(err)
	}

	data.Sign = sign
	//spew.Dump(transaction)

	jsonTr, err := json.Marshal(transaction)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonTr))

	reverseTr, err := tr.FromJSON(jsonTr)
	if err != nil {
		panic(err)
	}

	transf1 := tr.Transfer{Data: *reverseTr}
	valid := transf1.Verify(&usersBalance)
	fmt.Println("Valid: ", valid)
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

	usersBalance.Update(pubkey, 0, 2345.7)

	transaction, err := tr.NewTransfer(privkey, pubkey2, 14, 0.123, "memsage", &usersBalance)
	if err != nil {
		panic(err)
	}

	jsonTr, err := json.Marshal(transaction)
	if err != nil {
		panic(err)
	}

	var out2 bytes.Buffer

	err = json.Indent(&out2, jsonTr, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out2.Bytes()))

	reverseTr, err := tr.FromJSON(jsonTr)
	if err != nil {
		panic(err)
	}
	transf := tr.Transfer{Data: *reverseTr}
	err = transf.Verify(&usersBalance)
	if err != nil {
		return false
	}
	return true
}
