package model

import (
	"time"
)

type NftOwner struct {
	OwnerAddress      string `json:"owner_address,omitempty"`
	CollectionAddress string `json:"collection_address,omitempty"`
	TokenId           string `json:"token_id,omitempty"`
	ChainId           int64  `json:"chain_id,omitempty"`
	// specify references to join/preload nft_token
	// <-:false    Disable write permission
	Token           *NftToken `json:"token" gorm:"foreignkey:TokenId,CollectionAddress;references:TokenId,CollectionAddress;<-:false"`
	CreatedTime     time.Time `json:"created_time,omitempty"`
	LastUpdatedTime time.Time `json:"last_updated_time,omitempty"`
}
