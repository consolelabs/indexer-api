package model

import (
	"time"

	"gorm.io/gorm"
)

const (
	// listing status
	ListingStatusListed    = "listed"
	ListingStatusSold      = "sold"
	ListingStatusCancelled = "cancelled"
	ListingStatusDelisted  = "delisted"
	// listing type
	ListingTypeBuyNow  = "buy_now"
	ListingTypeAuction = "auction"
)

type NftListing struct {
	Id              uint64               `json:"id"`
	PlatformId      uint64               `json:"platform_id"`
	TokenId         string               `json:"token_id"`
	ContractAddress string               `json:"contract_address"`
	ChainId         int64                `json:"chain_id"`
	Quantity        string               `json:"quantity"`
	PaymentToken    int64                `json:"payment_token"`
	FromAddress     string               `json:"from_address"`
	ToAddress       string               `json:"to_address"`
	TransactionHash string               `json:"transaction_hash"`
	ListingType     string               `json:"listing_type"`
	ListingStatus   string               `json:"listing_status"`
	Marketplace     *MarketplacePlatform `json:"marketplace" gorm:"-"`
	CreatedTime     time.Time            `json:"created_time"`
	LastUpdatedTime time.Time            `json:"last_updated_time"`
	SoldPrice       string               `json:"sold_price"`
	ListingPrice    *string              `json:"listing_price"`
	SoldPriceObj    *Price               `json:"sold_price_obj" gorm:"-"`
	ListingPriceObj *Price               `json:"listing_price_obj" gorm:"-"`
	EventType       string               `json:"event_type" gorm:"-"`
}

type NftListingMarketplace struct {
	ContractAddress      string `json:"contract_address"`
	TokenId              string `json:"token_id"`
	PlatformId           uint64 `json:"platform_id"`
	PlatformName         string `json:"platform_name"`
	ListingPrice         string `json:"listing_price"`
	ListingStatus        string `json:"listing_status"`
	PaymentToken         string `json:"payment_token"`
	CreatedTime          string `json:"created_time"`
	PaymentTokenDecimals string `json:"payment_token_decimals"`
	FloorPrice           string `json:"floor_price"`
}

func (n *NftListing) BeforeCreate(tx *gorm.DB) (err error) {
	n.CreatedTime = time.Now()
	n.LastUpdatedTime = time.Now()
	return
}

func (n *NftListing) BeforeUpdate(tx *gorm.DB) (err error) {
	n.LastUpdatedTime = time.Now()
	return
}
