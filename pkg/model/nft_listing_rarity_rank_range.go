package model

type NftListingRarityRankRange struct {
	CollectionAddress           string                                           `json:"collection_address"`
	RarityRankRangeStart        string                                           `json:"rarity_rank_range_start"`
	RarityRankRange             string                                           `json:"rarity_rank_range"`
	MarketplacePlatformListings []NftListingRarityRankRangeByMarketplacePlatform `json:"marketplace_platform_listings"`
}

type NftListingRarityRankRangeByMarketplacePlatform struct {
	MarketplacePlatform MarketplacePlatform `json:"marketplace_platform"`
	AveragePrice        Price               `json:"average_price"`
}
