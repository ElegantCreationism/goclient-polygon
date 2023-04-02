package client

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// GetBlockNumber sends an HTTP POST request to the server at the specified URL with the given method to retrieve the current block number.
func GetBlockNumber(url string, method string) (interface{}, error) {
	// Create a new pointer to a GetBlockNumberResponse struct to hold the response from the server.
	var result = &GetBlockNumberResponse{}

	// Create a new GetBlockNumberRequest struct with the given method and request ID.
	requestData := GetBlockNumberRequest{
		JsonRPC: "2.0",
		Method:  method,
		ID:      2,
	}

	// Marshal the request data to JSON.
	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}

	// Send the HTTP POST request with the JSON request data in the request body.
	resp, err := http.Post(url, "application/json", bytes.NewReader(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Decode the response body JSON data into the result struct.
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&result); err != nil {
		return nil, err
	}

	// Return the block number from the response.
	return result.Result, nil
}

// GetBlockByNumber sends an HTTP POST request to the server at the specified URL with the given method and block number parameters to retrieve the specified block.
func GetBlockByNumber(url string, method string, blockNumber []interface{}) (*GetBlockByNumberResponse, error) {
	// Create a new GetBlockByNumberResponse struct to hold the response from the server.
	var result GetBlockByNumberResponse

	// Create a new GetBlockByNumberRequest struct with the given method, block number parameters, and request ID.
	requestData := GetBlockByNumberRequest{
		JsonRPC: "2.0",
		Method:  method,
		Params:  blockNumber,
		ID:      2,
	}

	// Marshal the request data to JSON.
	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}

	// Send the HTTP POST request with the JSON request data in the request body.
	resp, err := http.Post(url, "application/json", bytes.NewReader(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Decode the response body JSON data into the result struct.
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&result); err != nil {
		return nil, err
	}

	// Return the result struct.
	return &result, nil
}
