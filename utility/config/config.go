package config

import (
	"crypto/ecdsa"
	"encoding/base64"
	"errors"
	cr "github.com/ethereum/go-ethereum/crypto"
	"io/ioutil"
	"net"
	"strings"
)

const (
	fieldPubKey  = "pubKey"
	fieldPrivKey = "privkey"
	fieldCoordinator = "coordinator"
	fieldNodeToConnect = "nodeToConnect"

)

type Params struct {
	PubKey []byte
	PrivKey  *ecdsa.PrivateKey
	CoordinatorIP net.IP
	NodeToConnectIP net.IP
}

func LoadFromFile(file string) (*Params, error) {

	p := new(Params)
	//fmt.Println(path + "\\" + file)
	data, err := ioutil.ReadFile(file)
	if err != nil {
		//fmt.Println("Error read wallet cfg file.")
		return nil, err
	}
	lines := strings.Split(string(data), "\r\n")

	//fmt.Println("Data in config file:")
	for _, val := range lines {
		pair := strings.Split(val, ":")
		//fmt.Println(pair)
		if len(pair) != 2 {
			continue
		}
		switch pair[0] {

		case fieldPubKey:
			p.PubKey, err = base64.StdEncoding.DecodeString(pair[1])
			if err != nil {
				return nil, err
			}
		case fieldPrivKey:
			pKey, err := base64.StdEncoding.DecodeString(pair[1])
			if err != nil {
				return nil, err
			}
			privKey, err := cr.ToECDSA(pKey[:32])
			if err != nil {
				panic(err)
			}
			p.PrivKey = privKey
		case fieldCoordinator:
			ip := net.ParseIP(pair[1])
			if ip == nil {
				return nil, errors.New("incorrect coordinator ip")
			}
			p.CoordinatorIP = ip
		case fieldNodeToConnect:
			ip := net.ParseIP(pair[1])
			if ip == nil {
				return nil, errors.New("incorrect nodeToConnect ip")
			}
			p.NodeToConnectIP = ip
		}

	}
	//fmt.Println("End data in config file.")
	return p, nil
}
