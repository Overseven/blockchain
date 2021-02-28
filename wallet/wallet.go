package wallet

import (
	"fmt"
	"io/ioutil"
	"strings"
	"encoding/base64"
)

const (
	fieldPubKey = "wallet"
	fieldPrivKey = "privkey"
)

func LoadFromFile(file string) (address, privkey []byte){
	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("Error read wallet cfg file.")
		panic(err)
	}
	lines := strings.Split(string(data), "\r\n")

	//fmt.Println("Data in config file:")
	for _, val := range lines{
		pair := strings.Split(val, ":")
		//fmt.Println(pair)
		if len(pair) == 2{
			if pair[0] == fieldPubKey{
				address, err = base64.StdEncoding.DecodeString(pair[1])
				if err != nil {
					panic(err)
				}
			} else if pair[0] == fieldPrivKey {
				privkey, err = base64.StdEncoding.DecodeString(pair[1])
				if err != nil {
					panic(err)
				}
			}
		}
	}
	//fmt.Println("End data in config file.")
	return
}