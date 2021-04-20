package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	pcoord "github.com/overseven/blockchain/protocol/coordinator"
	pnode "github.com/overseven/blockchain/protocol/node"
	"github.com/overseven/blockchain/utility"
	"google.golang.org/grpc"
)

const (
	maxCountOfNodes = 40
)

func connectToNodes() error {
	if node.coordinator != "" {
		err := connectToCoordinator()
		if err != nil {
			return err
		}

	} else if len(node.Nodes) == 0 {
		return errors.New("coordinator or nodeToConnect must be presented")
	}

	return updateListOfNodes()
}

func newCoordinatorClient(address string) (pcoord.CoordinatorClient, *grpc.ClientConn, error) {
	con, err := grpc.Dial(address, grpc.WithInsecure())
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

func getNodesFromCoordinator() (map[string]interface{}, error) {
	fmt.Println("getNodesFromCoordinator()")
	if node.coordinator == "" {
		return nil, errors.New("empty coordinator address")
	}
	coordClient, _, err := newCoordinatorClient(node.coordinator)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	reply, err := coordClient.GetListOfNodes(ctx, &pcoord.ListOfNodesRequest{})

	if err != nil {
		return nil, err
	}

	nodes := map[string]interface{}{}
	for _, n := range reply.Address {
		nodes[n] = struct{}{}
	}

	return nodes, nil
}

func getNodesFromNode(address string) (map[string]interface{}, error) {
	fmt.Println("getNodesFromNode()")
	if address == "" {
		return nil, errors.New("empty node address")
	}
	nodeClient, _, err := newNodeClient(address)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	reply, err := nodeClient.GetListOfNodes(ctx, &pnode.ListOfNodesRequest{})

	if err != nil {
		return nil, err
	}

	nodes := map[string]interface{}{}
	for _, n := range reply.Address {
		nodes[n] = struct{}{}
	}

	return nodes, nil
}

func getNodesFromNodes(nodes map[string]interface{}) (map[string]interface{}, error) {
	fmt.Println("getNodesFromNodes()")
	// TODO: finish
	if len(nodes) == 0 {
		return nil, errors.New("empty input params")
	}
	var nodesForConnect []string
	result := map[string]interface{}{}
	for n := range nodes {
		nodesForConnect = append(nodesForConnect, n)
	}

	wg := sync.WaitGroup{}

	wg.Add(len(nodesForConnect))

	for _, n := range nodesForConnect {
		addr := n
		go func() {
			ns, err := getNodesFromNode(addr)
			if err == nil {
				for nFromN := range ns {
					result[nFromN] = struct{}{}
				}
			}
			wg.Done()
		}()
	}

	wg.Wait()

	return result, nil
}



func fractalNodeFinder(nodes map[string]interface{}, max int) error {
	fmt.Println("fractalNodeFinder()")
	used := map[string]interface{}{}

	for count:= len(nodes); count < max; count = len(nodes) {
		diff := utility.MapDifference(nodes, used)
		if _, ok := diff[node.OwnAddress.String()]; !ok {
			delete(diff, node.OwnAddress.String())
		}

		if len(diff) == 0 {
			return nil
		}
		fmt.Printf("Nodes: ")
		fmt.Println(diff)
		res, err := getNodesFromNodes(diff) // FIXME: infinity cycle
		if err != nil {
			continue
		}
		for key := range res {
			nodes[key] = struct{}{}
		}
	}

	
	return nil
}

func updateListOfNodes() error {
	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()
	fmt.Println("updateListOfNodes()")
	nodes := map[string]interface{}{}

	node.mutex.Lock()
	defer node.mutex.Unlock()

	for key := range node.Nodes {
		nodes[key] = struct{}{}
	}

	nCoord, err := getNodesFromCoordinator()
	if err == nil {
		for n := range nCoord {
			nodes[n] = struct{}{}
		}
	}

	err = fractalNodeFinder(nodes, maxCountOfNodes)
	if err != nil {
		// TODO: finish
	}

	if len(nodes) == 0 {
		return errors.New("empty list of nodes")
	}

	if _, ok := nodes[node.OwnAddress.String()]; ok {
		delete(nodes, node.OwnAddress.String())
	}

	for key := range nodes {
		node.Nodes[key] = struct{}{}
	}

	return nil
}
