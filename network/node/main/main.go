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

	newNodes, remNodes, err := connections.UpdateNodesList(&node.ActiveNodes, 50, node.ServParams)
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
		for _, nd := range newNodes {
			node.ActiveNodes.Nodes[nd] = struct{}{}
		}
		for _, nd := range remNodes {
			if _, ok := node.ActiveNodes.Nodes[nd]; ok {
				delete(node.ActiveNodes.Nodes, nd)
			}
		}
	}()

	go func() {
		for true {
			time.Sleep(5 * time.Second)
			func() {
				newNodes, removeNodes, err := connections.UpdateNodesList(&node.ActiveNodes, 50, node.ServParams)
				if err != nil {
					fmt.Println("Update list of nodes error:", err.Error())
				}else {

					fmt.Println("Cycle. Nodes:")
					node.ActiveNodes.Mutex.Lock()
					defer node.ActiveNodes.Mutex.Unlock()
					for _, nd := range newNodes {
						node.ActiveNodes.Nodes[nd] = struct{}{}
					}
					for _, nd := range removeNodes {
						if _, ok := node.ActiveNodes.Nodes[nd]; ok {
							delete(node.ActiveNodes.Nodes, nd)
						}
					}

					fmt.Println(node.ActiveNodes.Nodes)
					//for key := range node.Nodes {
					//	fmt.Printf("'%s'\n", key)
					//}
					fmt.Printf("(%d elems)\n\n", len(node.ActiveNodes.Nodes))
				}
				go func() {
					err := chnode.RegisterNodeOnNodes(node.ActiveNodes.ToStrings(), node.ServParams.OwnAddress.String())
					if err != nil {
						fmt.Println("Register node on nodes err:", err.Error())
					}
				}()

			}()

		}
	}()
	for true {
		time.Sleep(time.Second)
	}

	b := 5
	b = 23
	fmt.Println(b)
}
