package model

import (
	"time"
)

type NftToken struct {
	TokenId           string `json:"token_id"`
	CollectionAddress string `json:"collection_address,omitempty"`
	Name              string `json:"name,omitempty"`
	Description       string `json:"description,omitempty"`
	Amount            string `json:"amount,omitempty"`
	Image             string `json:"image,omitempty"`
	ImageCDN          string `json:"image_cdn,omitempty"`
	ThumbnailCDN      string `json:"thumbnail_cdn,omitempty"`
	ImageContentType  string `json:"image_content_type,omitempty"`
	RarityRank        uint64 `json:"rarity_rank,omitempty"`
	RarityScore       string `json:"rarity_score,omitempty"`
	RarityTier        string `json:"rarity_tier"`
	IsSelfHosted      bool   `json:"is_self_hosted"`

	Attributes []NftTokenAttribute `json:"attributes" gorm:"-"` // disable write
	// Attributes []NftTokenAttribute `json:"attributes" gorm:"foreignkey:TokenId,CollectionAddress;references:TokenId,CollectionAddress;<-:false"` // disable write
	Rarity *NftTokenRarity `json:"rarity" gorm:"-"`
	// Listing    *NftListing         `json:"listing" gorm:"foreignKey:TokenId,ContractAddress;references:TokenId,CollectionAddress;<-:false"` // disable write
	Listing            *NftListing             `json:"listing" gorm:"-"`
	CreatedTime        time.Time               `json:"created_time,omitempty"`
	LastUpdatedTime    time.Time               `json:"last_updated_time,omitempty"`
	Owner              *NftOwner               `json:"owner" gorm:"-"`
	ListingMarketplace []NftListingMarketplace `json:"listing_marketplace" gorm:"-"`
}
