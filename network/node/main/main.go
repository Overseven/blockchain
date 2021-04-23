package main

import (
	"fmt"
	"time"

	"github.com/overseven/blockchain/network/connections"
	chnode "github.com/overseven/blockchain/network/node"
)

func main() {
	node := chnode.NewNode()
	err := chnode.FlagParse(node)
	if err != nil {
		fmt.Println("Error!", err)
		return
	}

	stopListeningCh := make(chan interface{})
	err = node.StartListening(stopListeningCh)

	if err != nil {
		fmt.Println("Error!", err)
		return
	}

	err = chnode.RegisterNodeOnCoordinator(node.ActiveNodes.Coordinator, node.ServParams.OwnAddress.String())
	if err != nil {
		fmt.Println("Error!", err)
		//return
	}

	newNodes, err := connections.UpdateNodesList(&node.ActiveNodes, 50, node.ServParams)
	if err != nil {
		fmt.Println("Error!", err)

	}

	//go chnode.RegisterNodeOnNodes(node.ActiveNodes.ToStrings(), node.ServParams.OwnAddress.String())
	// if err != nil {
	// 	fmt.Println("Error!", err)
	// 	//return
	// }

	func() {
		node.ActiveNodes.Mutex.Lock()
		defer node.ActiveNodes.Mutex.Unlock()
		if newNodes != nil {
			node.ActiveNodes.Nodes = newNodes.Nodes
		} else {
			node.ActiveNodes.Nodes = map[string]interface{}{}
		}
	}()

	go func() {
		for true {
			time.Sleep(5 * time.Second)
			func() {
				newNodes, err := connections.UpdateNodesList(&node.ActiveNodes, 50, node.ServParams)
				if err != nil {
					fmt.Println("Update list of nodes error:", err.Error())
				}

				fmt.Println("Cycle. Nodes:")
				node.ActiveNodes.Mutex.Lock()
				defer node.ActiveNodes.Mutex.Unlock()
				node.ActiveNodes.Nodes = newNodes.Nodes
				fmt.Println(node.ActiveNodes.Nodes)
				//for key := range node.Nodes {
				//	fmt.Printf("'%s'\n", key)
				//}
				fmt.Printf("(%d elems)\n\n", len(node.ActiveNodes.Nodes))

				//go chnode.RegisterNodeOnNodes(node.ActiveNodes.ToStrings(), node.ServParams.OwnAddress.String())

			}()

		}
	}()
	for true {
		time.Sleep(time.Second)
	}

}
