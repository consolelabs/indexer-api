package model

import "time"

type NftMarketplaceCollectionSnapshotMarketCapTimeRange struct {
	TimeRangeStart time.Time `json:"time_range_start"`
	TimeInterval   string    `json:"time_interval"`
	MarketCapObj   *Price    `json:"market_cap"`
}
