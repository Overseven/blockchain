package node

import (
	"context"
	"crypto/ecdsa"
	"net"
	"strconv"
	"sync"

	"github.com/overseven/blockchain/balance"
	"github.com/overseven/blockchain/node/trlists"
	"github.com/overseven/blockchain/protocol/converter"
	pnode "github.com/overseven/blockchain/protocol/node"
	"github.com/overseven/blockchain/transaction"
	"github.com/overseven/blockchain/utility"
	"google.golang.org/grpc"
)

type Node struct {
	pnode.UnimplementedNoderServer
	ListeningPort uint32
	UsersBalance  balance.Balance
	PrivKey       *ecdsa.PrivateKey
	PubKey        []byte
	mutex         sync.Mutex
	Connected     []Connection
}

type Connection struct {
	ipAddress string
	pubKey    []byte
}

func (n *Node) SetPrivateKey(key *ecdsa.PrivateKey) {
	n.PrivKey = key
	n.PubKey = utility.PrivToPubKey(key)
}

func (n *Node) StartListening(stop chan bool) error {
	// TODO: stop signal handling
	// TODO: goroutine ?
	lis, err := net.Listen("tcp", "localhost:"+strconv.FormatUint(uint64(n.ListeningPort), 10))
	if err != nil {
		return err
	}
	s := grpc.NewServer()
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

	return &pnode.ConnectReply{ReplyerAddress: n.PubKey}, nil
}

func (n *Node) GetListOfNodes(ctx context.Context, req *pnode.ListOfNodesRequest) (*pnode.ListOfNodesReply, error) {
	nodeList := []string{}
	for _, c := range n.Connected {
		nodeList = append(nodeList, c.ipAddress)
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
