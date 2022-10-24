package model

import "time"

type NftListingTotalVolumeTimeRange struct {
	TimeRangeStart time.Time `json:"time_range_start"`
	TimeInterval   string    `json:"time_interval"`
	TotalVolumeObj *Price    `json:"total_volume"`
}
