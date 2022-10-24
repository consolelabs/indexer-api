package model

type Block struct {
	Number           int64  `json:"number"`
	Hash             string `json:"hash"`
	ParentHash       string `json:"parent_hash"`
	Nonce            string `json:"nonce"`
	Sha3Uncles       string `json:"sha3_uncles"`
	LogsBloom        string `json:"logs_bloom"`
	TransactionsRoot string `json:"transactions_root"`
	StateRoot        string `json:"state_root"`
	ReceiptsRoot     string `json:"receipts_root"`
	Miner            string `json:"miner"`
	Difficulty       int64  `json:"difficulty"`
	TotalDifficulty  int64  `json:"total_difficulty"`
	Size             int64  `json:"size"`
	ExtraData        []byte `json:"extra_data"`
	GasLimit         int64  `json:"gas_limit"`
	GasUsed          int64  `json:"gas_used"`
	Timestamp        int64  `json:"timestamp"`
	TransactionCount int64  `json:"transaction_count"`
	BaseFeePerGas    int64  `json:"base_fee_per_gas"`
}
