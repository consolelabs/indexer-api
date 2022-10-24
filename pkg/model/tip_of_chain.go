package model

type TipOfChain struct {
	// Tip is a chainId and lastest fetched blocknumber mapping
	// e.g [1]200000, means chainId 1 (eth) has latest block number 200000
	BlockHeadNumber map[int64]int64
}
