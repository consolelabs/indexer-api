package model

import (
	"time"
)

type NftListingSoldPrice struct {
	TimeRangeStart time.Time `json:"time_range_start"`
	TimeInterval   string    `json:"time_interval"`
	SoldPriceObj   *Price    `json:"sold_price_obj"`
}
