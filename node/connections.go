package main

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	pcoord "github.com/overseven/blockchain/protocol/coordinator"
	"google.golang.org/grpc"
)

func connectToNodes() error {
	if node.coordinatorIP != nil {
		err := connectToCoordinator()
		if err != nil {
			return err
		}

	} else if node.nodeToConnectIP != nil {
		err := connectToFirstNode()
		if err != nil {
			return err
		}
	} else {
		return errors.New("coordinator or nodeToConnect must be presented.")
	}

	return updateListOfNodes()
}

func connectToCoordinator() error {
	if node.coordinatorIP == nil {
		return nil
	}

	con, err := grpc.Dial(node.coordinatorIP.String()+":"+strconv.Itoa(int(node.coordinatorPort)), grpc.WithInsecure())
	if err != nil {
		return err
	}
	node.coordinatorClient = pcoord.NewCoordinatorClient(con)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	reply, err := node.coordinatorClient.Connect(ctx, &pcoord.ConnectRequest{RequesterAddress: node.OwnAddress.String()})
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

	if node.coordinatorClient != nil {
		reply, err := node.coordinatorClient.GetListOfNodes(ctx, &pcoord.ListOfNodesRequest{})

		if err != nil {
			return err
		}
		fmt.Println("List of nodes ip:")
		for _, n := range reply.Address {
			fmt.Println(n)
		}
		fmt.Printf("(%d elems)\n", len(reply.Address))

	} else if node.nodeToConnectClient != nil {
		// TODO: finish

	} else if len(node.Connected) > 0 {
		// TODO: finish

	} else {
		return errors.New("coordinator or nodeToConnect must be presented.")
	}
	return nil
}
