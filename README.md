# Go Program for Interacting with Polygon JSON-RPC API

This Go program is a simple client made to interact with the Polygon JSON-RPC API. The program retrieves the latest 
block number from the Polygon blockchain and then retrieves the block information for that block. It then prints out the 
block number and block hash of the retrieved block.

## Usage

To use this program, you need to have Go installed on your system. Once complete, then run the program by executing the 
following command in the terminal:
```go
go run main.go
```

## The client Package

The client package contains two methods: **GetBlockNumber** and **GetBlockByNumber**. 
Here is a brief description of what each method does:

### GetBlockNumber

This method sends a POST request to the specified JSON-RPC endpoint with the **eth_blockNumber** method and returns the 
latest block number from the blockchain. It takes two parameters:

* **url**: The URL of the JSON-RPC endpoint.
* **method**: The name of the method to call (eth_blockNumber in this case).

### GetBlockByNumber

This method sends a POST request to the specified JSON-RPC endpoint with the **eth_getBlockByNumber** method and a block 
number parameter to retrieve block information. It takes three parameters:

* **url**: The URL of the JSON-RPC endpoint.
* **method**: The name of the method to call (eth_getBlockByNumber in this case).
* **blockNumber**: A slice of interface{} containing the block number and a boolean value.

## Conclusion

This program provides a simple example of how to use the client package to interact with a JSON-RPC API using Go code. 
It demonstrates how to retrieve the latest block number from the Polygon blockchain and how to retrieve the block 
information for a specific block. You can use this program as a starting point for building more complex applications 
that interact with the Polygon JSON-RPC API.
