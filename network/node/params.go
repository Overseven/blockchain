package node

import (
	"errors"
	"flag"

	"github.com/overseven/try-network/utility"
	"github.com/overseven/try-network/utility/config"
)

func FlagParse(node *Node) error {
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
		if *flagPrivKey == "" {
			return errors.New("node private key must be presented with flag '-pubKey' or in config file")
		}
		res, err := utility.ParseKeys(*flagPrivKey, *flagPubKey)
		if err != nil {
			return nil
		}
		_, pub, err := utility.ToBytes(res)
		if err != nil {
			return nil
		}
		node.Wallet.PrivKey = res
		node.Wallet.PubKey = pub

		if *flagPort != 9000 {
			node.ServParams.ListeningPort = uint64(*flagPort)
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
			node.ActiveNodes.Coordinator = *flagCoordinatorIP
		} else {
			// TODO: add check
			// ip := net.ParseIP(*flagNodeToConnectIP)
			// if ip == nil {
			// 	return errors.New("incorrect nodeToConnect ip")
			// }
			node.ActiveNodes.Nodes[*flagNodeToConnectIP] = struct{}{}
		}

	} else {
		params, err := config.LoadFromFile(*flagCfgFile)
		if err != nil {
			return err
		}

		node.Wallet.PrivKey = params.PrivKey
		node.Wallet.PubKey = params.PubKey
		node.ServParams.ListeningPort = params.ListeningPort
		node.ActiveNodes.Coordinator = params.Coordinator
		if params.NodeToConnect != "" {
			node.ActiveNodes.Nodes[params.NodeToConnect] = struct{}{}
		}
	}
	return nil
}
