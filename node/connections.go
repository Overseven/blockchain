package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	pcoord "github.com/overseven/blockchain/protocol/coordinator"
	pnode "github.com/overseven/blockchain/protocol/node"
	"google.golang.org/grpc"
)

const (
	maxCountOfNodes = 40
)

func connectToNodes() error {
	if len(node.coordinator) != 0 {
		err := connectToCoordinator()
		if err != nil {
			return err
		}

	} else if len(node.Nodes) != 0 {
		err := connectToFirstNode()
		if err != nil {
			return err
		}
	} else {
		return errors.New("coordinator or nodeToConnect must be presented.")
	}

	return updateListOfNodes()
}

func newCoordinatorClient(address string) (pcoord.CoordinatorClient, *grpc.ClientConn, error) {
	con, err := grpc.Dial(node.coordinator, grpc.WithInsecure())
	if err != nil {
		return nil, nil, err
	}
	coordClient := pcoord.NewCoordinatorClient(con)

	return coordClient, con, nil
}

func newNodeClient(address string) (pnode.NoderClient, *grpc.ClientConn, error) {
	con, err := grpc.Dial(node.coordinator, grpc.WithInsecure())
	if err != nil {
		return nil, nil, err
	}
	nodeClient := pnode.NewNoderClient(con)

	return nodeClient, con, nil
}

func connectToCoordinator() error {
	if len(node.coordinator) == 0 {
		return nil
	}

	coordClient, _, err := newCoordinatorClient(node.coordinator)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	reply, err := coordClient.Connect(ctx, &pcoord.ConnectRequest{RequesterAddress: node.OwnAddress.String()})
	if err != nil {
		return err
	}
	fmt.Println("Coordinator connection status:", reply.Ok)
	return nil
}

func connectToFirstNode() error {
	// TODO: finish
	return nil
}

func updateListOfNodes() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	nodes := []string{}

	if node.coordinator != "" {
		coordClient, _, err := newCoordinatorClient(node.coordinator)
		if err == nil {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			reply, err := coordClient.GetListOfNodes(ctx, &pcoord.ListOfNodesRequest{})

			if err == nil {
				nodes = append(nodes, reply.Address...)
				fmt.Println("List of nodes ip from coordinator:")
				for _, n := range reply.Address {
					fmt.Println(n)
				}
				fmt.Printf("(%d elems)\n", len(reply.Address))
			} else {
				fmt.Println("Warning: coordinator connection error")
			}
		} else {
			fmt.Println("Warning: create coordinator client error")
		}
	}

	if len(node.Nodes) != 0 {
		// wg := sync.WaitGroup{}
		// wg.Add(len(node.Nodes))
		for n := range node.Nodes {
			nodeClient, _, err := newNodeClient(n)
			if err != nil {
				// TODO: remove node
				continue
			}
			reply, err := nodeClient.GetListOfNodes(ctx, &pnode.ListOfNodesRequest{})
			if err == nil {
				fmt.Println("List of nodes ip from coordinator:")
				for _, n := range reply.Address {
					fmt.Println(n)
				}
				fmt.Printf("(%d elems)\n", len(reply.Address))
				nodes = append(nodes, reply.Address...)
			} else {
				fmt.Println("Warning: nodeToConnect connection error")
			}
		}

	}

	if len(node.Nodes) > 0 {
		if len(nodes) < maxCountOfNodes {
			// TODO: finish request other nodes ip
		}

	}

	if len(nodes) == 0 {
		return errors.New("empty list of nodes")
	}

	for _, n := range nodes {
		if n == node.OwnAddress.String() {
			continue
		}
		if _, ok := node.Nodes[n]; ok {
			continue
		}
		node.Nodes[n] = struct{}{}
	}

	return nil
}
