package interfaces

//pb "github.com/overseven/blockchain/protocol"

type NetworkNode interface {
	Networker
	WalletOwner
	ChainHolder
	Init()
	StartListening() error
	GetWaitingTrans() []BlockElement
	CreateBlock([]BlockElement) TransactionsContainer
	// SendTransactionToAllNodes(element BlockElement) ([]pb.AddTransactionReply_Code, error)
	// SendTransaction(element BlockElement, nodeAddress string) (pb.AddTransactionReply_Code, error)
}
