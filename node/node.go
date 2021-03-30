package node

import (
	"context"
	"crypto/ecdsa"
	"net"
	"strconv"
	"sync"

	"github.com/overseven/blockchain/balance"
	"github.com/overseven/blockchain/chain"
	"github.com/overseven/blockchain/interfaces"
	pb "github.com/overseven/blockchain/protocol"
	"github.com/overseven/blockchain/protocol/converter"
	"github.com/overseven/blockchain/utility"
	"google.golang.org/grpc"
)

type Node struct {
	pb.UnimplementedNoderServer
	Mode          interfaces.ClientMode
	ListeningPort uint32
	usersBalance  balance.Balance
	localChain    interfaces.BlockConnecter
	privateKey    *ecdsa.PrivateKey
	publicKey     []byte
	waitingTrans  []interfaces.BlockElement
	mutex         sync.Mutex
}

func (n *Node) Init() {
	n.localChain = new(chain.Chain)
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

// SayHello implements helloworld.GreeterServer
func (n *Node) AddTransaction(ctx context.Context, in *pb.AddTransactionRequest) (*pb.AddTransactionReply, error) {
	trans, err := converter.TransactionProto2Local(in.Transction)
	if err != nil {
		return &pb.AddTransactionReply{Reply: pb.AddTransactionReply_TR_Error, Message: err.Error(), Additional: ""}, err
	}
	n.mutex.Lock()
	defer n.mutex.Unlock()
	n.waitingTrans = append(n.waitingTrans, trans)

	//log.Printf("Node received transaction request with %f value and %f fee", in.Transction.Pay, in.Transction.Fee)

	return &pb.AddTransactionReply{Reply: pb.AddTransactionReply_TR_Ok, Message: "Ok!", Additional: "aga"}, nil
}

func (n *Node) GetWaitingTrans() []interfaces.BlockElement {
	return n.waitingTrans
}

func (n *Node) CreateBlock([]interfaces.BlockElement) interfaces.TransactionsContainer {
	block := n.localChain.NewBlock()
	block.SetMiner(n.publicKey)
	return block
}

func (n *Node) GetChain() interfaces.BlockConnecter {
	return n.localChain
}

func (n *Node) SetChain(b interfaces.BlockConnecter) {
	n.localChain = b
}
