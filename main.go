package main

import (
	"fmt"
	"github.com/ElegantCreationism/goclient-polygon/client"
	"log"
)

func main() {
	url := "https://polygon-rpc.com/"
	blockNumber, err := client.GetBlockNumber(url, "eth_blockNumber")
	if err != nil {
		log.Fatal(err)
	}

	params := []interface{}{fmt.Sprintf("%v", blockNumber), true}
	blockResponse, err := client.GetBlockByNumber(url, "eth_getBlockByNumber", params)
	if err != nil {
		// Handle error
		log.Fatal(err)
	}

	fmt.Printf("Block number: %v, block hash: %v\n", blockResponse.Result.Number, blockResponse.Result.Hash)
}
