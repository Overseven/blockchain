package node

import (
	"encoding/json"
	"fmt"
	chain2 "github.com/overseven/blockchain/chain"
	"github.com/overseven/blockchain/transaction/itransaction"
	"net/http"

	"github.com/overseven/blockchain/chain"

	//"github.com/davecgh/go-spew/spew"
)

var (
	localBlockchain chain.Chain
)

func Run() {
	fmt.Println("Launching node.")
	http.HandleFunc("/transaction/new", receiveNewTransaction)
	http.ListenAndServe(":8090", nil)
	bl := chain2.Block{
		Id: 0,
	}
	//blBase := chain.Block(bl)
	localBlockchain.Blocks = append(localBlockchain.Blocks, &bl)
	fmt.Println(localBlockchain)
}

func receiveNewTransaction(w http.ResponseWriter, req *http.Request) {
	var t itransaction.ITransaction
	err := json.NewDecoder(req.Body).Decode(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//spew.Dump(t)

	valid := t.Verify()
	fmt.Println("Receive transaction, valid: ", valid)
}
