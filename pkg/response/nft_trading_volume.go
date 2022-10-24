package response

type NftTradingVolume struct {
	CollectionAddress string  `json:"collection_address"`
	CollectionName    string  `json:"collection_name"`
	CollectionSymbol  string  `json:"collection_symbol"`
	CollectionChainId int     `json:"collection_chain_id"`
	TradingVolume     float64 `json:"trading_volume"`
	Token             string  `json:"token"`
}

type NftTradingVolumeResponse struct {
	Data NftTradingVolume `json:"data"`
}
