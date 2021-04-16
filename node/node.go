package node

import (
	"context"
	"crypto/ecdsa"
	"net"
	"strconv"
	"sync"

	"github.com/overseven/blockchain/balance"
	"github.com/overseven/blockchain/db"
	"github.com/overseven/blockchain/interfaces"
	"github.com/overseven/blockchain/node/trlists"
	"github.com/overseven/blockchain/protocol/converter"
	pb "github.com/overseven/blockchain/protocol/node"
	"github.com/overseven/blockchain/transaction"
	"github.com/overseven/blockchain/utility"
	"google.golang.org/grpc"
	"grpc-go-1.35.0/codes"
	"grpc-go-1.35.0/status"
)

type Node struct {
	pb.UnimplementedNoderServer
	Mode          interfaces.ClientMode
	ListeningPort uint32
	usersBalance  balance.Balance
	privateKey    *ecdsa.PrivateKey
	publicKey     []byte
	mutex         sync.Mutex
}

func (n *Node) SetPrivateKey(key *ecdsa.PrivateKey) {
	n.privateKey = key
	n.publicKey = utility.PrivToPubKey(key)
}

func (n *Node) GetPrivateKey() *ecdsa.PrivateKey {
	return n.privateKey
}

func (n *Node) SetPublicKey(key []byte) {
	n.publicKey = key
}

func (n *Node) GetPublicKey() []byte {
	return n.publicKey
}

func (n *Node) SetPort(port uint32) {
	n.ListeningPort = port
}

func (n *Node) GetPort() uint32 {
	return n.ListeningPort
}

func (n *Node) StartListening(stop chan bool) error {
	// TODO: stop signal handling
	// TODO: goroutine ?
	lis, err := net.Listen("tcp", "localhost:"+strconv.FormatUint(uint64(n.ListeningPort), 10))
	if err != nil {
		return err
	}
	s := grpc.NewServer()
	pb.RegisterNoderServer(s, n)
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

func (n *Node) CreateBlock([]interfaces.BlockElement) interfaces.TransactionsContainer {
	lastBlock := db.GetLastBlock()
	block := block.NewBlock{}
	block.SetMiner(n.publicKey)
	return block
}

func (n *Node) Connect(ctx context.Context, req *pb.ConnectRequest) (*pb.ConnectReply, error) {
	// TODO: finish
	return nil, status.Errorf(codes.Unimplemented, "method Connect not implemented")
}

func (n *Node) GetListOfNodes(ctx context.Context, req *pb.ListOfNodesRequest) (*pb.ListOfNodesReply, error) {
	// TODO: finish
	return nil, status.Errorf(codes.Unimplemented, "method GetListOfNodes not implemented")
}

func (n *Node) AddTransaction(ctx context.Context, req *pb.AddTransactionRequest) (*pb.AddTransactionReply, error) {
	trans, err := converter.TransactionProto2Local(req.Transaction)
	if err != nil {
		return &pb.AddTransactionReply{Reply: pb.AddTransactionReply_TR_Error, Message: err.Error(), Additional: ""}, err
	}

	// TODO: add validation
	trlists.AddToFirst([]transaction.Transaction{trans})

	//log.Printf("Node received transaction request with %f value and %f fee", in.Transction.Pay, in.Transction.Fee)

	return &pb.AddTransactionReply{Reply: pb.AddTransactionReply_TR_Ok, Message: "Ok!", Additional: "aga"}, nil
}

func (n *Node) PushBlock(ctx context.Context, req *pb.PushBlockRequest) (*pb.PushBlockReply, error) {
	// TODO: finish
	return &pb.PushBlockReply{Cod: pb.AddTransactionReply_TR_Ok, Message: "Ok!", Additional: "aga"}, nil
}

func (n *Node) GetBlocks(ctx context.Context, req *pb.GetBlocksRequest) (*pb.GetBlocksReply, error) {
	// TODO: finish
	return nil, status.Errorf(codes.Unimplemented, "method GetBlocks not implemented")
}
