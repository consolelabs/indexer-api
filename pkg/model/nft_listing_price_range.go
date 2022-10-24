package model

type NftListingPriceRange struct {
	CollectionAddress           string                                      `json:"collection_address"`
	PriceRangeStart             *Price                                      `json:"price_range_start"`
	PriceRange                  string                                      `json:"price_range"`
	MarketplacePlatformListings []NftListingPriceRangeByMarketplacePlatform `json:"marketplace_platform_listings"`
}

type NftListingPriceRangeByMarketplacePlatform struct {
	MarketplacePlatform MarketplacePlatform `json:"marketplace_platform"`
	TotalListing        string              `json:"total_listing"`
}
