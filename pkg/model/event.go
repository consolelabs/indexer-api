package model

type Event struct {
	ChainId          int64    `json:"chain_id"`
	LogIndex         int64    `json:"log_index"`
	TransactionHash  string   `json:"transaction_hash"`
	TransactionIndex int64    `json:"transaction_index"`
	BlockHash        string   `json:"block_hash"`
	BlockNumber      int64    `json:"block_number"`
	Address          string   `json:"address"`
	Data             []byte   `json:"data"`
	Topics           []string `json:"topics"`
}
