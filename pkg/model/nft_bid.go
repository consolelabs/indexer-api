package model

import "time"

type NftBid struct {
	TokenId         string               `json:"token_id,omitempty"`
	ContractAddress string               `json:"contract_address,omitempty"`
	FromAddress     string               `json:"from_address,omitempty"`
	ToAddress       string               `json:"to_address,omitempty"`
	TransactionHash string               `json:"transaction_hash,omitempty"`
	Marketplace     *MarketplacePlatform `json:"marketplace,omitempty"`
	ExpiredTime     time.Time            `json:"expired_time,omitempty"`
	BidPrice        *Price               `json:"bid_price,omitempty"`
	UsdPrice        float64              `json:"usd_price"`
	FloorDifference float64              `json:"floor_difference"`
}
