package model

import "time"

type NftMarketplaceCollectionSnapshotFloorPriceTimeRange struct {
	TimeRangeStart time.Time `json:"time_range_start"`
	TimeInterval   string    `json:"time_interval"`
	FloorPriceObj  *Price    `json:"floor_price"`
}
