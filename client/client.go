package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	tr "github.com/overseven/blockchain/transaction"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	fmt.Println("Launching server...")
	http.HandleFunc("/test", test)
	http.ListenAndServe(":8090", nil)
}

func test(w http.ResponseWriter, req *http.Request) {
	var t tr.Transaction
	err := json.NewDecoder(req.Body).Decode(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	spew.Dump(t)

	valid := tr.Verify(&t)
	fmt.Println("valid: ", valid)
}
