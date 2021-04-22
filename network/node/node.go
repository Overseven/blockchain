package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/overseven/blockchain/network"
	"github.com/overseven/blockchain/node/trlists"
	"github.com/overseven/blockchain/protocol/converter"
	pnode "github.com/overseven/blockchain/protocol/node"
	"github.com/overseven/blockchain/transaction"
	"github.com/overseven/blockchain/utility"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

var node Node

type Node struct {
	pnode.UnimplementedNoderServer
	ServParams network.ServerParams
	NetParams  network.NetParams
}

type Connection struct {
	address string
	pubKey  []byte
}

func init() {
	node.NetParams.Nodes = map[string]interface{}{}
}

func (n *Node) SetPrivateKey(key *ecdsa.PrivateKey) {
	n.NetParams.PrivKey = key
	n.NetParams.PubKey = utility.PrivToPubKey(key)
}

func (n *Node) StartListening(stop chan interface{}) error {
	// TODO: stop signal handling
	// TODO: goroutine ?
	lis, err := net.Listen("tcp", "localhost:"+strconv.FormatUint(n.ServParams.ListeningPort, 10))
	if err != nil {
		return err
	}
	s := grpc.NewServer()
	n.ServParams.OwnAddress = lis.Addr()
	pnode.RegisterNoderServer(s, n)
	go s.Serve(lis)
	go func() {
		<-stop
		s.Stop()
	}()
	// if err := s.Serve(lis); err != nil {
	// 	return err
	// }
	return nil
}

// func (n *Node) CreateBlock([]interfaces.BlockElement) interfaces.TransactionsContainer {
// 	lastBlock := db.GetLastBlock()
// 	block := block.NewBlock{}
// 	block.SetMiner(n.publicKey)
// 	return block
// }

func (n *Node) Connect(ctx context.Context, req *pnode.ConnectRequest) (*pnode.ConnectReply, error) {
	// TODO: finish
	fmt.Println("Connection request received!")
	p, _ := peer.FromContext(ctx)
	addr := p.Addr.String()
	fmt.Println("Address:", addr)
	n.NetParams.Nodes[addr] = struct{}{}
	return &pnode.ConnectReply{ReplyerAddress: n.NetParams.PubKey}, nil
}

func (n *Node) GetListOfNodes(ctx context.Context, req *pnode.ListOfNodesRequest) (*pnode.ListOfNodesReply, error) {
	nodeList := []string{}
	for key := range n.NetParams.Nodes {
		nodeList = append(nodeList, key)
	}
	return &pnode.ListOfNodesReply{Address: nodeList}, nil
}

func (n *Node) AddTransaction(ctx context.Context, req *pnode.AddTransactionRequest) (*pnode.AddTransactionReply, error) {
	trans, err := converter.TransactionProto2Local(req.Transaction)
	if err != nil {
		return &pnode.AddTransactionReply{Reply: pnode.AddTransactionReply_TR_Error, Message: err.Error(), Additional: ""}, err
	}

	// TODO: add validation
	trlists.AddToFirst([]transaction.Transaction{trans})

	//log.Printf("Node received transaction request with %f value and %f fee", in.Transction.Pay, in.Transction.Fee)

	return &pnode.AddTransactionReply{Reply: pnode.AddTransactionReply_TR_Ok, Message: "Ok!", Additional: "aga"}, nil
}

func (n *Node) PushBlock(ctx context.Context, req *pnode.PushBlockRequest) (*pnode.PushBlockReply, error) {
	// TODO: finish
	return &pnode.PushBlockReply{Reply: pnode.PushBlockReply_PBR_Ok}, nil
}

func (n *Node) GetBlocks(ctx context.Context, req *pnode.GetBlocksRequest) (*pnode.GetBlocksReply, error) {
	// TODO: finish
	return &pnode.GetBlocksReply{}, nil
}

func main() {
	err := flagParse()
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

	err = connectToNodes()
	if err != nil {
		fmt.Println("Error!", err)
		//return
	}

	go func() {
		for true {
			time.Sleep(5 * time.Second)
			func() {
				updateListOfNodes()
				fmt.Println("Cycle. Nodes:")
				node.NetParams.Mutex.Lock()
				defer node.NetParams.Mutex.Unlock()
				fmt.Println(node.NetParams.Nodes)
				//for key := range node.Nodes {
				//	fmt.Printf("'%s'\n", key)
				//}
				fmt.Printf("(%d elems)\n\n", len(node.NetParams.Nodes))
			}()

		}
	}()
	for true {
		time.Sleep(time.Second)
	}

}
