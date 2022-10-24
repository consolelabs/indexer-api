package model

type BlockStats struct {
	ChainId         int64   `json:"chain_id"`
	FirstBlock      int64   `json:"first_block"`
	LastBlock       int64   `json:"last_block"`
	TotalBlocks     int64   `json:"total_blocks"`
	ActualBlocks    uint64  `json:"actual_blocks"`
	RemainingBlocks int64   `json:"remaining_blocks"`
	Percentage      float64 `json:"percentage"`
}
