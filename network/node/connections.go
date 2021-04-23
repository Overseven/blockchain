package node

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/overseven/blockchain/network/connections"
	pcoord "github.com/overseven/blockchain/network/protocol/coordinator"
	pnode "github.com/overseven/blockchain/network/protocol/node"
	"github.com/overseven/blockchain/utility"
)

func RegisterNodeOnCoordinator(coordAddress, ownAddress string) error {
	if len(coordAddress) == 0 {
		return errors.New("empty coord address")
	}
	if len(ownAddress) == 0 {
		return errors.New("empty own address")
	}

	coordClient, _, err := connections.NewCoordinatorClient(coordAddress)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	reply, err := coordClient.Connect(ctx, &pcoord.ConnectRequest{RequesterAddress: ownAddress})
	if err != nil {
		return err
	}
	fmt.Println("Coordinator connection status:", reply.Ok)
	return nil
}

// TODO: check - deep copy?
func RegisterNodeOnNodes(nodesAddress []string, ownAddress string) error {

	i, ok := utility.Find(nodesAddress, ownAddress)
	for ok {
		nodesAddress = utility.RemoveIndex(nodesAddress, i)
		i, ok = utility.Find(nodesAddress, ownAddress)
	}

	if len(nodesAddress) == 0 {
		return errors.New("empty list of nodes address")
	}

	wg := new(sync.WaitGroup)
	wg.Add(len(nodesAddress))

	for _, n := range nodesAddress {
		nodeClient, _, err := connections.NewNodeClient(n)
		if err != nil {
			return err
		}
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		reply, err := nodeClient.Connect(ctx, &pnode.ConnectRequest{RequesterAddress: ownAddress})
		if err != nil {
			return err
		}
		fmt.Println("Register. Replyer address:", reply.ReplyerAddress)
	}
	wg.Wait()
	return nil
}
