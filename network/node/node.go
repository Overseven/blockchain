package node

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"net"
	"strconv"

	"github.com/overseven/blockchain/network"
	"github.com/overseven/blockchain/network/node/trlists"
	"github.com/overseven/blockchain/network/protocol/converter"
	pnode "github.com/overseven/blockchain/network/protocol/node"
	"github.com/overseven/blockchain/transaction"
	"github.com/overseven/blockchain/utility"
	"google.golang.org/grpc"
)

type Node struct {
	pnode.UnimplementedNoderServer
	ServParams  network.ServerParams
	ActiveNodes network.NodesContainer
	Wallet      network.Wallet
}

func NewNode() *Node {
	node := new(Node)
	node.ActiveNodes.Nodes = map[string]interface{}{}
	return node
}

func (n *Node) SetPrivateKey(key *ecdsa.PrivateKey) {
	n.Wallet.PrivKey = key
	n.Wallet.PubKey = utility.PrivToPubKey(key)
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
	go func() {
		err := s.Serve(lis)
		if err != nil {
			fmt.Println("Serve err!", err.Error())
		}
	}()
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
	//fmt.Println("Connection request received!")
	addr := req.RequesterAddress

	//fmt.Println("Address:", addr)
	n.ActiveNodes.Mutex.Lock()
	defer n.ActiveNodes.Mutex.Unlock()
	n.ActiveNodes.Nodes[addr] = struct{}{}
	return &pnode.ConnectReply{ReplyerAddress: n.Wallet.PubKey}, nil
}

func (n *Node) GetListOfNodes(ctx context.Context, req *pnode.ListOfNodesRequest) (*pnode.ListOfNodesReply, error) {
	var nodeList []string
	n.ActiveNodes.Mutex.Lock()
	defer n.ActiveNodes.Mutex.Unlock()
	for key := range n.ActiveNodes.Nodes {
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

func (n *Node) GetBlocks(ctx context.Context, req *pnode.BlocksRequest) (*pnode.BlocksReply, error) {
	// TODO: finish
	return &pnode.BlocksReply{}, nil
}
