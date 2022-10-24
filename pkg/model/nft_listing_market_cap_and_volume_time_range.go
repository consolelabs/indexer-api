package model

type NftListingMarketCapAndVolumeTimeRange struct {
	TimeRangeStart string `json:"time_range_start"`
	TimeInterval   string `json:"time_interval"`
	MarketCap      *Price `json:"market_cap"`
	Volume         *Price `json:"volume"`
}
