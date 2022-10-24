package model

import (
	"time"
)

type NftTransfer struct {
	From            string    `json:"from"`
	To              string    `json:"to"`
	TokenId         string    `json:"token_id"`
	EventTime       int64     `json:"event_time"`
	TransactionHash string    `json:"transaction_hash"`
	BlockNumber     int64     `json:"block_number"`
	ContractAddress string    `json:"contract_address"`
	CreatedTime     time.Time `json:"created_time"`
	LastUpdatedTime time.Time `json:"last_updated_time"`
}

type NftTxHistory struct {
	From            string    `json:"from"`
	To              string    `json:"to"`
	TokenId         string    `json:"token_id"`
	TransactionHash string    `json:"transaction_hash"`
	ContractAddress string    `json:"contract_address"`
	CreatedTime     time.Time `json:"created_time"`
	ListingStatus   string    `json:"listing_status"`
	EventType       string    `json:"event_type"`
}
