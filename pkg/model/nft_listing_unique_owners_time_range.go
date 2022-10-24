package model

import "time"

type NftListingUniqueOwnersTimeRange struct {
	TimeRangeStart time.Time `json:"time_range_start"`
	TimeInterval   string    `json:"time_interval"`
	UniqueOwners   int       `json:"unique_owners"`
}
