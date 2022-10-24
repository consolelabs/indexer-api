package model

import "time"

type NftMarketplaceCollectionSnapshot struct {
	CreatedTime       time.Time            `json:"created_time"`
	PlatformId        int64                `json:"platform_id"`
	CollectionAddress string               `json:"collection_address"`
	Url               string               `json:"url"`
	Description       string               `json:"description"`
	RoyaltyPercentage float64              `json:"royalty_percentage"`
	RoyalAddress      string               `json:"royalty_address"`
	TokenId           int64                `json:"token_id"`
	FloorPrice        string               `json:"floor_price"`
	TotalVolume       string               `json:"total_volume"`
	FloorPriceObj     *Price               `json:"floor_price_obj" gorm:"-"`
	TotalVolumeObj    *Price               `json:"total_volume_obj" gorm:"-"`
	Marketplace       *MarketplacePlatform `json:"marketplace" gorm:"-"`
}
