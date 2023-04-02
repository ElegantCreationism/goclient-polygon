package client_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/ElegantCreationism/goclient-polygon/client"
)

func TestMakeGetBlockNumberRequest(t *testing.T) {
	expectedResult := "0x272919c"

	// Create a mock server to handle the request
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Check the request method
		if req.Method != "POST" {
			t.Errorf("expected method 'POST', but got '%s'", req.Method)
		}
		// Check the request URL
		if req.URL.String() != "/api" {
			t.Errorf("expected URL '/api', but got '%s'", req.URL.String())
		}
		// Check the request body
		var requestData client.GetBlockNumberRequest
		if err := json.NewDecoder(req.Body).Decode(&requestData); err != nil {
			t.Fatal(err)
		}
		expectedData := client.GetBlockNumberRequest{
			JsonRPC: "2.0",
			Method:  "eth_blockNumber",
			ID:      2,
		}
		if requestData != expectedData {
			t.Errorf("expected request body '%v', but got '%v'", expectedData, requestData)
		}
		// Send the response
		responseData := client.GetBlockNumberResponse{
			JsonRPC: "2.0",
			ID:      2,
			Result:  expectedResult,
		}
		jsonData, err := json.Marshal(responseData)
		if err != nil {
			t.Fatal(err)
		}
		rw.Header().Set("Content-Type", "application/json")
		rw.Write(jsonData)
	}))
	defer server.Close()

	// Make the request to the mock server
	result, err := client.GetBlockNumber(server.URL+"/api", "eth_blockNumber")
	if err != nil {
		t.Fatal(err)
	}
	// Check the response
	if result != expectedResult {
		t.Errorf("expected result '%s', but got '%v'", expectedResult, result)
	}
}

func TestMakeGetBlockByNumberRequest(t *testing.T) {
	// Prepare the test server to mock the API endpoint
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check the request method and URL path
		if r.Method != http.MethodPost {
			t.Errorf("Expected request method to be POST; got %s", r.Method)
		}
		if r.URL.Path != "/some-url" {
			t.Errorf("Expected request URL path to be /some-url; got %s", r.URL.Path)
		}

		// Check the request body
		reqBody, err := io.ReadAll(r.Body)
		if err != nil {
			t.Errorf("Error reading request body: %s", err.Error())
		}
		var reqData client.GetBlockByNumberRequest
		err = json.Unmarshal(reqBody, &reqData)
		if err != nil {
			t.Errorf("Error decoding request body: %s", err.Error())
		}
		if reqData.JsonRPC != "2.0" {
			t.Errorf("Expected JSON-RPC version to be 2.0; got %s", reqData.JsonRPC)
		}
		if reqData.Method != "eth_getBlockByNumber" {
			t.Errorf("Expected method to be eth_getBlockByNumber; got %s", reqData.Method)
		}
		if len(reqData.Params) != 2 || reqData.Params[0] != "0x272919c" || reqData.Params[1] != true {
			t.Errorf("Expected params to be [\"0x272919c\", true]; got %v", reqData.Params)
		}

		// Send the response
		respData := client.GetBlockByNumberResponse{
			JsonRPC: "2.0",
			ID:      2,
			Result: struct {
				BaseFeePerGas   string `json:"baseFeePerGas"`
				Difficulty      string `json:"difficulty"`
				ExtraData       string `json:"extraData"`
				GasLimit        string `json:"gasLimit"`
				GasUsed         string `json:"gasUsed"`
				Hash            string `json:"hash"`
				LogsBloom       string `json:"logsBloom"`
				Miner           string `json:"miner"`
				MixHash         string `json:"mixHash"`
				Nonce           string `json:"nonce"`
				Number          string `json:"number"`
				ParentHash      string `json:"parentHash"`
				ReceiptsRoot    string `json:"receiptsRoot"`
				Sha3Uncles      string `json:"sha3Uncles"`
				Size            string `json:"size"`
				StateRoot       string `json:"stateRoot"`
				Timestamp       string `json:"timestamp"`
				TotalDifficulty string `json:"totalDifficulty"`
				Transactions    []struct {
					BlockHash            string        `json:"blockHash"`
					BlockNumber          string        `json:"blockNumber"`
					From                 string        `json:"from"`
					Gas                  string        `json:"gas"`
					GasPrice             string        `json:"gasPrice"`
					MaxFeePerGas         string        `json:"maxFeePerGas,omitempty"`
					MaxPriorityFeePerGas string        `json:"maxPriorityFeePerGas,omitempty"`
					Hash                 string        `json:"hash"`
					Input                string        `json:"input"`
					Nonce                string        `json:"nonce"`
					To                   string        `json:"to"`
					TransactionIndex     string        `json:"transactionIndex"`
					Value                string        `json:"value"`
					Type                 string        `json:"type"`
					AccessList           []interface{} `json:"accessList,omitempty"`
					ChainId              string        `json:"chainId,omitempty"`
					V                    string        `json:"v"`
					R                    string        `json:"r"`
					S                    string        `json:"s"`
				} `json:"transactions"`
				TransactionsRoot string        `json:"transactionsRoot"`
				Uncles           []interface{} `json:"uncles"`
			}(struct {
				BaseFeePerGas   string
				Difficulty      string
				ExtraData       string
				GasLimit        string
				GasUsed         string
				Hash            string
				LogsBloom       string
				Miner           string
				MixHash         string
				Nonce           string
				Number          string
				ParentHash      string
				ReceiptsRoot    string
				Sha3Uncles      string
				Size            string
				StateRoot       string
				Timestamp       string
				TotalDifficulty string
				Transactions    []struct {
					BlockHash            string        `json:"blockHash"`
					BlockNumber          string        `json:"blockNumber"`
					From                 string        `json:"from"`
					Gas                  string        `json:"gas"`
					GasPrice             string        `json:"gasPrice"`
					MaxFeePerGas         string        `json:"maxFeePerGas,omitempty"`
					MaxPriorityFeePerGas string        `json:"maxPriorityFeePerGas,omitempty"`
					Hash                 string        `json:"hash"`
					Input                string        `json:"input"`
					Nonce                string        `json:"nonce"`
					To                   string        `json:"to"`
					TransactionIndex     string        `json:"transactionIndex"`
					Value                string        `json:"value"`
					Type                 string        `json:"type"`
					AccessList           []interface{} `json:"accessList,omitempty"`
					ChainId              string        `json:"chainId,omitempty"`
					V                    string        `json:"v"`
					R                    string        `json:"r"`
					S                    string        `json:"s"`
				}
				TransactionsRoot string
				Uncles           []interface{}
			}{BaseFeePerGas: "", Difficulty: "", ExtraData: "", GasLimit: "", GasUsed: "", Hash: "0x123456", LogsBloom: "", Miner: "", MixHash: "", Nonce: "", Number: "0x1", ParentHash: "", ReceiptsRoot: "", Sha3Uncles: "", Size: "", StateRoot: "", Timestamp: "", TotalDifficulty: "", Transactions: nil, TransactionsRoot: "", Uncles: nil}),
		}

		respBody, err := json.Marshal(respData)
		if err != nil {
			t.Errorf("Error encoding response body: %s", err.Error())
		}
		_, _ = w.Write(respBody)
	}))
	defer testServer.Close()

	// Make the request
	blockNumber := []interface{}{"0x272919c", true}
	response, err := client.GetBlockByNumber(testServer.URL+"/some-url", "eth_getBlockByNumber", blockNumber)
	if err != nil {
		t.Errorf("Error making request: %s", err.Error())
	}

	// Check the response
	expectedResponse := &client.GetBlockByNumberResponse{
		JsonRPC: "2.0",
		ID:      2,
		Result: struct {
			BaseFeePerGas   string `json:"baseFeePerGas"`
			Difficulty      string `json:"difficulty"`
			ExtraData       string `json:"extraData"`
			GasLimit        string `json:"gasLimit"`
			GasUsed         string `json:"gasUsed"`
			Hash            string `json:"hash"`
			LogsBloom       string `json:"logsBloom"`
			Miner           string `json:"miner"`
			MixHash         string `json:"mixHash"`
			Nonce           string `json:"nonce"`
			Number          string `json:"number"`
			ParentHash      string `json:"parentHash"`
			ReceiptsRoot    string `json:"receiptsRoot"`
			Sha3Uncles      string `json:"sha3Uncles"`
			Size            string `json:"size"`
			StateRoot       string `json:"stateRoot"`
			Timestamp       string `json:"timestamp"`
			TotalDifficulty string `json:"totalDifficulty"`
			Transactions    []struct {
				BlockHash            string        `json:"blockHash"`
				BlockNumber          string        `json:"blockNumber"`
				From                 string        `json:"from"`
				Gas                  string        `json:"gas"`
				GasPrice             string        `json:"gasPrice"`
				MaxFeePerGas         string        `json:"maxFeePerGas,omitempty"`
				MaxPriorityFeePerGas string        `json:"maxPriorityFeePerGas,omitempty"`
				Hash                 string        `json:"hash"`
				Input                string        `json:"input"`
				Nonce                string        `json:"nonce"`
				To                   string        `json:"to"`
				TransactionIndex     string        `json:"transactionIndex"`
				Value                string        `json:"value"`
				Type                 string        `json:"type"`
				AccessList           []interface{} `json:"accessList,omitempty"`
				ChainId              string        `json:"chainId,omitempty"`
				V                    string        `json:"v"`
				R                    string        `json:"r"`
				S                    string        `json:"s"`
			} `json:"transactions"`
			TransactionsRoot string        `json:"transactionsRoot"`
			Uncles           []interface{} `json:"uncles"`
		}(struct {
			BaseFeePerGas   string
			Difficulty      string
			ExtraData       string
			GasLimit        string
			GasUsed         string
			Hash            string
			LogsBloom       string
			Miner           string
			MixHash         string
			Nonce           string
			Number          string
			ParentHash      string
			ReceiptsRoot    string
			Sha3Uncles      string
			Size            string
			StateRoot       string
			Timestamp       string
			TotalDifficulty string
			Transactions    []struct {
				BlockHash            string        `json:"blockHash"`
				BlockNumber          string        `json:"blockNumber"`
				From                 string        `json:"from"`
				Gas                  string        `json:"gas"`
				GasPrice             string        `json:"gasPrice"`
				MaxFeePerGas         string        `json:"maxFeePerGas,omitempty"`
				MaxPriorityFeePerGas string        `json:"maxPriorityFeePerGas,omitempty"`
				Hash                 string        `json:"hash"`
				Input                string        `json:"input"`
				Nonce                string        `json:"nonce"`
				To                   string        `json:"to"`
				TransactionIndex     string        `json:"transactionIndex"`
				Value                string        `json:"value"`
				Type                 string        `json:"type"`
				AccessList           []interface{} `json:"accessList,omitempty"`
				ChainId              string        `json:"chainId,omitempty"`
				V                    string        `json:"v"`
				R                    string        `json:"r"`
				S                    string        `json:"s"`
			}
			TransactionsRoot string
			Uncles           []interface{}
		}{BaseFeePerGas: "", Difficulty: "", ExtraData: "", GasLimit: "", GasUsed: "", Hash: "0x123456", LogsBloom: "", Miner: "", MixHash: "", Nonce: "", Number: "0x1", ParentHash: "", ReceiptsRoot: "", Sha3Uncles: "", Size: "", StateRoot: "", Timestamp: "", TotalDifficulty: "", Transactions: nil, TransactionsRoot: "", Uncles: nil}),
	}

	if !reflect.DeepEqual(response, expectedResponse) {
		t.Errorf("Response not as expected. Expected %+v; got %+v", expectedResponse, response)
	}
}
