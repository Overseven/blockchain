package main

import (
	"context"
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"

	pcoord "github.com/overseven/blockchain/protocol/coordinator"
	"google.golang.org/grpc"
)

var coord Coordinator

type Coordinator struct {
	ListeningPort uint64
	nodes         map[string]NodeInfo
	pcoord.UnimplementedCoordinatorServer
	lock sync.Mutex
}

type NodeInfo struct {
	ip     net.IP
	port   uint64
	pubKey []byte
}

func init() {
	fmt.Println("init!")
	coord.nodes = map[string]NodeInfo{}
}

func fillRandomData() {
	coord.nodes = map[string]NodeInfo{}
	coord.nodes["124.52.23.24:9000"] = NodeInfo{}
	coord.nodes["14.35.90.15:9001"] = NodeInfo{}
	coord.nodes["234.73.5.50:9002"] = NodeInfo{}
	coord.nodes["117.11.26.31:9003"] = NodeInfo{}
}

func (c *Coordinator) StartListening(stop chan interface{}) error {
	// TODO: stop signal handling
	// TODO: goroutine ?
	lis, err := net.Listen("tcp", "localhost:"+strconv.FormatUint(c.ListeningPort, 10))
	if err != nil {
		return err
	}
	s := grpc.NewServer()

	pcoord.RegisterCoordinatorServer(s, c)
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

func (c *Coordinator) Connect(ctx context.Context, req *pcoord.ConnectRequest) (*pcoord.ConnectReply, error) {
	fmt.Println("Get request to connect")

	// md, ok := metadata.FromIncomingContext(ctx)
	str := req.RequesterAddress
	pair := strings.Split(str, ":")
	ip := net.ParseIP(pair[0])
	if ip == nil {
		return &pcoord.ConnectReply{Ok: false}, nil
	}
	_, err := strconv.ParseUint(pair[1], 10, 64)
	if err != nil {
		return &pcoord.ConnectReply{Ok: false}, nil
	}
	fmt.Println("Get request to connect: ", str)
	if len(str) == 0 {
		return &pcoord.ConnectReply{Ok: false}, nil
	}
	if str == "" {
		return &pcoord.ConnectReply{Ok: false}, nil
	}
	c.nodes[str] = NodeInfo{}
	fmt.Println("Connect done")
	return &pcoord.ConnectReply{Ok: true}, nil
}
func (c *Coordinator) GetListOfNodes(ctx context.Context, req *pcoord.ListOfNodesRequest) (*pcoord.ListOfNodesReply, error) {
	nodes := []string{}

	coord.lock.Lock()
	defer coord.lock.Unlock()

	for key := range c.nodes {
		nodes = append(nodes, key)
	}
	fmt.Println("Get request to list of nodes")
	return &pcoord.ListOfNodesReply{Address: nodes}, nil
}

func infinityPing() {

	ping := func() {
		coord.lock.Lock()
		defer coord.lock.Unlock()

		for addr := range coord.nodes {
			a := addr
			go func() {
				timeout := time.Duration(2 * time.Second)
				_, err := net.DialTimeout("tcp", a, timeout)
				if err != nil {
					fmt.Printf("%s %s\n", a, "not responding")
					coord.lock.Lock()
					defer coord.lock.Unlock()
					//fmt.Println("remove node address:", a)
					_, ok := coord.nodes[a]
					if ok {
						delete(coord.nodes, a)
					}

				} else {
					//fmt.Printf("%s %s\n", a, "responding")
				}
			}()
		}
	}
	go func() {
		for true {
			time.Sleep(10 * time.Second)
			ping()
		}
	}()

	go func() {
		for true {
			time.Sleep(15 * time.Second)
			fmt.Printf("\n\nNodes after resync:\n")
			for addr := range coord.nodes {
				fmt.Println(addr)
			}

		}
	}()

	for true {

	}
}

func main() {
	fillRandomData()
	coord.ListeningPort = 9004
	stopCh := make(chan interface{})
	err := coord.StartListening(stopCh)
	if err != nil {
		fmt.Println("Error!", err)
		return
	}
	infinityPing()
}
