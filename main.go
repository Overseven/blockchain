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

	flag.Parse()

	if *flagNode {
		fmt.Println("Node choosen!")
		node.Run()

	}else if *flagClient {
		fmt.Println("Client choosen!")
		client.Run()

	} else{
		fmt.Println("Choose mode. Use \"-h\" param.")
		return
	}
	return
}
