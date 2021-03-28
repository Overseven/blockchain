package node

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/overseven/blockchain/balance"
	"github.com/overseven/blockchain/block"
	chain "github.com/overseven/blockchain/chain"
	"github.com/overseven/blockchain/interfaces"
	//"github.com/davecgh/go-spew/spew"
)

var (
	localBlockchain chain.Chain
	usersBalance    balance.Balance
)

func Run() {
	fmt.Println("Launching node.")
	http.HandleFunc("/transaction/new", receiveNewTransaction)
	http.ListenAndServe(":8090", nil)
	bl := block.Block{
		Id: 0,
	}
	//blBase := chain.Block(bl)
	localBlockchain.Blocks = append(localBlockchain.Blocks, &bl)
	fmt.Println(&localBlockchain)
}

func receiveNewTransaction(w http.ResponseWriter, req *http.Request) {
	var t interfaces.BlockElement
	err := json.NewDecoder(req.Body).Decode(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//spew.Dump(t)

	valid := t.Verify(&usersBalance)
	fmt.Println("Receive transaction, valid: ", valid)
}
