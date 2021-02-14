package node

import (
	"encoding/json"
	"fmt"
	"net/http"
	tr "github.com/overseven/blockchain/transaction"
	blockchain "github.com/overseven/blockchain/blockchain"
	//"github.com/davecgh/go-spew/spew"
)

var (
	localBlockchain blockchain.Blockchain
)
func Run(){
	fmt.Println("Launching node.")
	http.HandleFunc("/transaction/new", receiveNewTransaction)
	http.ListenAndServe(":8090", nil)
	bl := blockchain.Block{
		Id: 0,
	}
	localBlockchain.Blocks = append(localBlockchain.Blocks, bl)
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