package config

import (
	"crypto/ecdsa"
	"encoding/base64"
	"errors"
	"io/ioutil"
	"net"
	"strconv"
	"strings"

	cr "github.com/ethereum/go-ethereum/crypto"
)

const (
	fieldPubKey        = "pubKey"
	fieldPrivKey       = "privkey"
	fieldPort          = "port"
	fieldCoordinator   = "coordinator"
	fieldNodeToConnect = "nodeToConnect"
)

type Params struct {
	PubKey        []byte
	PrivKey       *ecdsa.PrivateKey
	ListeningPort uint64
	Coordinator   string
	NodeToConnect string
}

func LoadFromFile(file string) (*Params, error) {

	p := new(Params)
	//fmt.Println(path + "\\" + file)
	data, err := ioutil.ReadFile(file)
	if err != nil {
		//fmt.Println("Error read wallet cfg file.")
		return nil, err
	}
	lines := strings.Split(strings.ReplaceAll(string(data), " ", ""), "\r\n")

	//fmt.Println("Data in config file:")
	for _, val := range lines {
		pair := strings.SplitN(val, "=", 2)
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
				return nil, err
			}
			p.PrivKey = privKey

		case fieldPort:
			port, err := strconv.ParseUint(pair[1], 10, 64)
			if err != nil {
				return nil, err
			}
			p.ListeningPort = port

		case fieldCoordinator:
			addr := strings.Split(pair[1], ":")
			if len(addr) != 2 {
				return nil, errors.New("incorrect coordinator address format. Wanted: x.x.x.x:x")
			}
			ip := net.ParseIP(addr[0])
			if ip == nil {
				return nil, errors.New("incorrect coordinator ip")
			}

			_, err = strconv.ParseUint(addr[1], 10, 64)
			if err != nil {
				return nil, errors.New("incorrect coordinator address format. Wanted: x.x.x.x:x")
			}
			p.Coordinator = pair[1]

		case fieldNodeToConnect:
			addr := strings.Split(pair[1], ":")
			if len(addr) != 2 {
				return nil, errors.New("incorrect nodeToConnect address format. Wanted: x.x.x.x:x")
			}
			ip := net.ParseIP(addr[0])
			if ip == nil {
				return nil, errors.New("incorrect nodeToConnect ip")
			}

			_, err = strconv.ParseUint(addr[1], 10, 64)
			if err != nil {
				return nil, errors.New("incorrect nodeToConnect address format. Wanted: x.x.x.x:x")
			}
			p.NodeToConnect = pair[1]
		default:
		}

	}
	//fmt.Println("End data in config file.")
	return p, nil
}
