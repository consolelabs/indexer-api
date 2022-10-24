package model

import (
	"time"
)

type NftCollectionSaleVolumeAndFloorPrice struct {
	CollectionAddress string    `json:"collection_address"`
	SaleVolume        *Price    `json:"sale_volume"`
	FloorPriceObj     *Price    `json:"floor_price"`
	CreatedTime       time.Time `json:"created_time"`
}
