package model

const (
	PlatformPaintswap = "paintswap"
	PlatformNftkey    = "nftkey"
	PlatformX2Y2    = "x2y2"
)

type MarketplacePlatform struct {
	ID      uint64 `json:"id"`
	Name    string `json:"name"`
	URL     string `json:"url"`
	IconUrl string `json:"icon_url"`
}

type MarketplaceWithListingStats struct {
	MarketplacePlatform
	ListingNumber     int64   `json:"listing_number"`
	ListingPercentage float64 `json:"listing_percentage"`
}
