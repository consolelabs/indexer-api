package model

type NftMarketplaceCollection struct {
	PlatformID        int64                `json:"platform_id"`
	CollectionAddress string               `json:"collection_address"`
	URL               string               `json:"url"`
	Description       string               `json:"description"`
	RoyaltyPercentage float64              `json:"royalty_percentage"`
	RoyaltyAddress    string               `json:"royalty_address"`
	Platform          *MarketplacePlatform `json:"platform" gorm:"foreignKey:PlatformID"`
	TokenId           int64                `json:"token_id"`
	FloorPrice        float64              `json:"floor_price"`
	TotalVolume       float64              `json:"total_volume"`
	FloorPriceObj     *Price               `json:"floor_price_obj" gorm:"-"`
	TotalVolumeObj    *Price               `json:"total_volume_obj" gorm:"-"`
}
