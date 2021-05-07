package config

import (
	"crypto/ecdsa"
	"errors"
	"io/ioutil"
	"net"
	"path/filepath"
	"strconv"
	"strings"

	cr "github.com/ethereum/go-ethereum/crypto"
	"github.com/overseven/blockchain/utility"
)

const (
	fieldPubKey        = "pubKey"
	fieldPrivKey       = "privkey"
	fieldPort          = "port"
	fieldCoordinator   = "coordinator"
	fieldNodeToConnect = "nodeToConnect"
)

type Params struct {
	PubKey        []byte // compressed public key
	PrivKey       *ecdsa.PrivateKey
	ListeningPort uint64
	Coordinator   string
	NodeToConnect string
}

func LoadFromFile(file string) (*Params, error) {

	p := new(Params)
	//fmt.Println(path + "\\" + file)
	absPath, err := filepath.Abs(file)
	if err != nil {
		//fmt.Println("Error read wallet cfg file.")
		return nil, err
	}
	data, err := ioutil.ReadFile(absPath)
	if err != nil {
		//fmt.Println("Error read wallet cfg file.")
		return nil, err
	}
	lines := strings.Split(strings.ReplaceAll(string(data), " ", ""), "\r\n")

	// fmt.Println("lines: ", lines)
	var privKeyHex string
	//fmt.Println("Data in config file:")
	for _, val := range lines {
		pair := strings.SplitN(val, "=", 2)
		//fmt.Println(pair)
		if len(pair) != 2 {
			continue
		}
		switch pair[0] {

		case fieldPubKey:
			p.PubKey = []byte(pair[1])

		case fieldPrivKey:
			privKeyHex = pair[1]

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

	priv, err := utility.ParseKeys(privKeyHex, string(p.PubKey))
	if err != nil {
		return nil, err
	}
	p.PrivKey = priv

	p.PubKey = cr.CompressPubkey(&p.PrivKey.PublicKey)

	//fmt.Println("End data in config file.")
	return p, nil
}
