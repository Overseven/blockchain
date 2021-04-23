package connections

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/overseven/blockchain/network"
	pcoord "github.com/overseven/blockchain/network/protocol/coordinator"
	pnode "github.com/overseven/blockchain/network/protocol/node"
	"github.com/overseven/blockchain/utility"
	"google.golang.org/grpc"
)

func copyNodes(input *network.NodesContainer) *network.NodesContainer {
	res := new(network.NodesContainer)
	input.Mutex.Lock()
	defer input.Mutex.Unlock()

	res.Coordinator = input.Coordinator
	res.Nodes = make(map[string]interface{})
	for key := range input.Nodes {
		res.Nodes[key] = struct{}{}
	}
	return res
}

func UpdateNodesList(nodes *network.NodesContainer, maxCount int, params network.ServerParams) (*network.NodesContainer, error) {
	tmpNodes := copyNodes(nodes)
	return updateList(tmpNodes, maxCount, params)
}

func NewCoordinatorClient(address string) (pcoord.CoordinatorClient, *grpc.ClientConn, error) {
	con, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, nil, err
	}
	coordClient := pcoord.NewCoordinatorClient(con)

	return coordClient, con, nil
}

func NewNodeClient(address string) (pnode.NoderClient, *grpc.ClientConn, error) {
	con, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, nil, err
	}
	nodeClient := pnode.NewNoderClient(con)

	return nodeClient, con, nil
}

func getNodesFromCoordinator(address string) (map[string]interface{}, error) {
	fmt.Println("getNodesFromCoordinator()")
	if address == "" {
		return nil, errors.New("empty coordinator address")
	}
	coordClient, _, err := NewCoordinatorClient(address)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	reply, err := coordClient.GetListOfNodes(ctx, &pcoord.ListOfNodesRequest{})

	if err != nil {
		return nil, err
	}
	// fmt.Println("From coord:", reply.Address)
	nodes := map[string]interface{}{}
	for _, n := range reply.Address {
		nodes[n] = struct{}{}
	}

	return nodes, nil
}

func getNodesFromNode(address string) (result map[string]interface{}, err error) {
	fmt.Println("getNodesFromNode()")
	if address == "" {
		return nil, errors.New("empty node address")
	}
	nodeClient, _, err := NewNodeClient(address)
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

func getNodesFromNodes(nodes map[string]interface{}) (result, remove map[string]interface{}, err error) {
	fmt.Println("getNodesFromNodes()")
	// TODO: finish
	if len(nodes) == 0 {
		return nil, nil, errors.New("empty input params")
	}
	var nodesForConnect []string
	result = map[string]interface{}{}
	remove = map[string]interface{}{}
	for n := range nodes {
		nodesForConnect = append(nodesForConnect, n)
	}

	wg := sync.WaitGroup{}

	wg.Add(len(nodesForConnect))

	for _, n := range nodesForConnect {
		addr := n
		go func() {
			ns, err := getNodesFromNode(addr)
			if err != nil {
				remove[addr] = struct{}{}
			} else {
				for nFromN := range ns {
					result[nFromN] = struct{}{}
				}
			}
			wg.Done()
		}()
	}

	wg.Wait()

	return result, remove, nil
}

func fractalNodeFinder(nodes map[string]interface{}, max int, ownAddress string) error {
	fmt.Println("fractalNodeFinder()")
	used := map[string]interface{}{}

	for count := len(nodes); count < max; count = len(nodes) {
		// fmt.Println("Nodes: ", nodes)
		// fmt.Println("Used: ", used)
		diff := utility.MapDifference(nodes, used)
		// fmt.Println("diff before: ", diff)
		if _, ok := diff[ownAddress]; ok {
			delete(diff, ownAddress)
		}
		// fmt.Println("diff after: ", diff)

		if len(diff) == 0 {
			return nil
		}
		// fmt.Printf("fractalNodeFinder. Nodes: ")
		// fmt.Println(diff)
		res, remove, err := getNodesFromNodes(diff) // FIXME: infinity cycle
		for key := range remove {
			if _, ok := nodes[key]; ok {
				// fmt.Println("remove:", key)
				delete(nodes, key)
			}
		}
		for key := range diff {
			used[key] = struct{}{}
		}
		if err != nil {
			continue
		}
		for key := range res {
			nodes[key] = struct{}{}
		}
	}

	return nil
}

func updateList(nds *network.NodesContainer, maxCount int, params network.ServerParams) (*network.NodesContainer, error) {
	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()
	fmt.Println("updateListOfNodes()")
	nodes := map[string]interface{}{}

	nds.Mutex.Lock()
	defer nds.Mutex.Unlock()

	for key := range nds.Nodes {
		// fmt.Println("key: ", key)
		nodes[key] = struct{}{}
	}

	nCoord, err := getNodesFromCoordinator(nds.Coordinator)
	if err == nil {
		for n := range nCoord {
			nodes[n] = struct{}{}
		}
	} else {
		fmt.Println("Err nodes from coord! ", err.Error())
	}
	fmt.Println("Nodes from coord: ", nodes)
	err = fractalNodeFinder(nodes, maxCount, params.OwnAddress.String())
	if err != nil {
		fmt.Println("Err  fractalNodeFinder! ", err.Error())
	}

	if len(nodes) == 0 {
		return nil, errors.New("empty list of nodes")
	}

	if _, ok := nodes[params.OwnAddress.String()]; ok {
		delete(nodes, params.OwnAddress.String())
	}

	// fmt.Println(nodes)
	nds.Nodes = nodes

	return nds, nil
}
