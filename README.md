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
To build and run as a docker image:
```makefile
make run
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

## Deployment
Prerequisites:
* ECS Cluster
* Load Balancer
* VPC and subnets (can be default)
* Security groups (can be default)

1. To deploy this service navigate to **'terraform/dev'** folder and run **terraform init**.
2. Alter values in backend.tf to set up your **aws_profile**, to connect to aws 
3. Run **terraform plan** and review
4. Run **terraform apply** and check logs for service output.


## Improvements
* Making the client more interactive, it only runs and returns the output, it doesn't run whenever called. (As a lambda it would be able to be triggered on demand)
* RPC request code can be refactored to promote code reuse
* Extend the service to support more methods, as it will be more useful
* Improve the error handling, add error codes, so that logs are more descriptive when errors occur, for easier diagnosis


This program provides a simple example of how to use the client package to interact with a JSON-RPC API using Go code.
It demonstrates how to retrieve the latest block number from the Polygon blockchain and how to retrieve the block
information for a specific block. You can use this program as a starting point for building more complex applications
that interact with the Polygon JSON-RPC API.
