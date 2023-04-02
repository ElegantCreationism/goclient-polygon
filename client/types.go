package client

type GetBlockNumberRequest struct {
	JsonRPC string `json:"jsonrpc"`
	Method  string `json:"method"`
	ID      int    `json:"id"`
}

type GetBlockNumberResponse struct {
	JsonRPC string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  string `json:"result"`
}

type GetBlockByNumberRequest struct {
	JsonRPC string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	ID      int           `json:"id"`
}

type GetBlockByNumberResponse struct {
	JsonRPC string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  struct {
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
	} `json:"result"`
}
