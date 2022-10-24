package model

type Transaction struct {
	Hash                 string `json:"hash"`
	Nonce                int64  `json:"nonce"`
	BlockHash            string `json:"block_hash"`
	BlockNumber          int64  `json:"block_number"`
	TransactionIndex     int64  `json:"transaction_index"`
	FromAddress          string `json:"from_address"`
	ToAddress            string `json:"to_address"`
	Value                int64  `json:"value"`
	Gas                  int64  `json:"gas"`
	GasPrice             int64  `json:"gas_price"`
	Input                []byte `json:"input"`
	BlockTimestamp       int64  `json:"block_timestamp"`
	MaxFeePerGas         int64  `json:"max_fee_per_gas"`
	MaxPriorityFeePerGas int64  `json:"max_priority_fee_per_gas"`
	TransactionType      int64  `json:"transaction_type"`
}
