package node

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Overseven/blockchain/chain"
	"github.com/Overseven/blockchain/chain/block"
	chainimpl "github.com/Overseven/blockchain/chain/chain"
	tr "github.com/overseven/blockchain/transaction"
	//"github.com/davecgh/go-spew/spew"
)

var (
	localBlockchain chainimpl.Chain
)

func Run() {
	fmt.Println("Launching node.")
	http.HandleFunc("/transaction/new", receiveNewTransaction)
	http.ListenAndServe(":8090", nil)
	bl := block.Block{
		Id: 0,
	}
	blBase := chain.Block(bl)
	localBlockchain.Blocks = append(localBlockchain.Blocks, blBase)
	fmt.Println(localBlockchain)
}

func receiveNewTransaction(w http.ResponseWriter, req *http.Request) {
	var t tr.Transaction
	err := json.NewDecoder(req.Body).Decode(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//spew.Dump(t)

	valid := t.Verify()
	fmt.Println("Receive transaction, valid: ", valid)
}
