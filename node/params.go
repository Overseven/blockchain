package main

import (
	"encoding/base64"
	"errors"
	"flag"

	cr "github.com/ethereum/go-ethereum/crypto"
	"github.com/overseven/blockchain/utility/config"
)

func flagParse() error {
	flagPubKey := flag.String("pubKey", "", "public node key")
	flagPrivKey := flag.String("privKey", "", "private node key")
	flagPort := flag.Uint("port", 9000, "listening port")
	flagCoordinatorIP := flag.String("coordinator", "", "coordinator address. \n"+
		"Format: ip:port \n"+
		"Example: 127.0.0.1:5000")
	flagNodeToConnectIP := flag.String("nodeToConnect", "", "another node address. \n"+
		"Format: ip:port \n"+
		"Example: 127.0.0.1:5001")
	flagCfgFile := flag.String("config", "", "config filename")
	flag.Parse()

	if *flagCfgFile == "" {
		if *flagPubKey == "" {
			return errors.New("node public key must be presented with flag '-pubKey' or in config file")
		}
		res, err := base64.StdEncoding.DecodeString(*flagPubKey)
		if err != nil {
			return nil
		}
		node.PubKey = res

		if *flagPrivKey == "" {
			return errors.New("node private key must be presented with flag '-pubKey' or in config file")
		}

		res, err = base64.StdEncoding.DecodeString(*flagPrivKey)
		if err != nil {
			return nil
		}
		privKey, err := cr.ToECDSA(res[:32])
		if err != nil {
			panic(err)
		}
		node.PrivKey = privKey

		if *flagPort != 9000 {
			node.ListeningPort = uint64(*flagPort)
		}
		if *flagCoordinatorIP == "" && *flagNodeToConnectIP == "" {
			return errors.New("coordinator or nodeToConnect must be presented")
		}
		if *flagCoordinatorIP != "" {
			// TODO: add check
			// ip := net.ParseIP(*flagCoordinatorIP)
			// if ip == nil {
			// 	return errors.New("incorrect coordinator ip")
			// }
			node.coordinator = *flagCoordinatorIP
		} else {
			// TODO: add check
			// ip := net.ParseIP(*flagNodeToConnectIP)
			// if ip == nil {
			// 	return errors.New("incorrect nodeToConnect ip")
			// }
			node.Nodes[*flagNodeToConnectIP] = struct{}{}
		}

	} else {
		params, err := config.LoadFromFile(*flagCfgFile)
		if err != nil {
			return err
		}

		node.PrivKey = params.PrivKey
		node.PubKey = params.PubKey
		node.ListeningPort = params.ListeningPort
		node.coordinator = params.Coordinator
		if params.NodeToConnect != "" {
			node.Nodes[params.NodeToConnect] = struct{}{}
		}
	}
	return nil
}
