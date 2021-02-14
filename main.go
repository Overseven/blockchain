package main

import (
	"flag"
	"fmt"
	node "github.com/overseven/blockchain/node"
	client "github.com/overseven/blockchain/client"
)

func main() {
	flagNode := flag.Bool("node", false, "start as node")
	flagClient := flag.Bool("client", false, "start as client")
	paramCfgFile := flag.String("config", "", "config filename")
	flag.Parse()

	if *flagNode {
		fmt.Println("Node choosen!")
		node.Run()

	}else if *flagClient {
		fmt.Println("Client choosen!")
		if len(*paramCfgFile) != 0{
			client.Run(*paramCfgFile)
		} else {
			panic("aaa config file?")
			return
		}

	} else{
		fmt.Println("Choose mode. Use \"-h\" param.")
		return
	}
	return
}
