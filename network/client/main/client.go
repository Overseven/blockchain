package main

import (
	"fmt"

	"github.com/overseven/try-network/network/client"
)

func main() {
	cl := client.NewClient()
	err := cl.FlagParse()
	if err != nil {
		fmt.Println("Error!", err.Error())
		return
	}
}
