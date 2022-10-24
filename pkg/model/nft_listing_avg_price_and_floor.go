package model

import (
	"time"
)

type NftListingAvgPriceAndFloor struct {
	TimeRangeStart time.Time `json:"time_range_start"`
	TimeInterval   string    `json:"time_interval"`
	AvgPriceObj    *Price    `json:"avg_price_obj"`
	FloorPriceObj  *Price    `json:"floor_price_obj"`
}
