package client

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func GetBlockNumber(url string, method string) (interface{}, error) {
	// Prepare the request body
	var result = &GetBlockNumberResponse{}
	requestData := GetBlockNumberRequest{
		JsonRPC: "2.0",
		Method:  method,
		ID:      2,
	}
	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}
	// Send the request
	resp, err := http.Post(url, "application/json", bytes.NewReader(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// Parse the response
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&result); err != nil {
		return nil, err
	}
	return result.Result, nil
}

func GetBlockByNumber(url string, method string, blockNumber []interface{}) (*GetBlockByNumberResponse, error) {
	// Prepare the request body
	var result GetBlockByNumberResponse
	requestData := GetBlockByNumberRequest{
		JsonRPC: "2.0",
		Method:  method,
		Params:  blockNumber,
		ID:      2,
	}
	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}
	// Send the request
	resp, err := http.Post(url, "application/json", bytes.NewReader(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// Parse the response
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
