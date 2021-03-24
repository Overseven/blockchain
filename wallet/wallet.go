package wallet

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const (
	fieldPubKey  = "wallet"
	fieldPrivKey = "privkey"
)

func LoadFromFile(file string) (address, privkey []byte, err error) {
	path, err := os.Getwd()
	if err != nil {
		return nil, nil, err
	}

	//fmt.Println(path + "\\" + file)
	data, err := ioutil.ReadFile(path + "\\" + file)
	if err != nil {
		fmt.Println("Error read wallet cfg file.")
		return nil, nil, err
	}
	lines := strings.Split(string(data), "\r\n")

	//fmt.Println("Data in config file:")
	for _, val := range lines {
		pair := strings.Split(val, ":")
		//fmt.Println(pair)
		if len(pair) == 2 {
			if pair[0] == fieldPubKey {
				address, err = base64.StdEncoding.DecodeString(pair[1])
				if err != nil {
					return nil, nil, err
				}
			} else if pair[0] == fieldPrivKey {
				privkey, err = base64.StdEncoding.DecodeString(pair[1])
				if err != nil {
					return nil, nil, err
				}
			}
		}
	}
	//fmt.Println("End data in config file.")
	return
}
