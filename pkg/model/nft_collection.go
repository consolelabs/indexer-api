package model

import (
	"fmt"
	"strings"
	"time"
)

type NftCollection struct {
	ID                     int64                      `json:"id"`
	Address                string                     `json:"collection_address"`
	Name                   string                     `json:"name"`
	Symbol                 string                     `json:"symbol"`
	ChainId                int64                      `json:"chain_id"`
	ERCFormat              string                     `json:"erc_format"`
	Supply                 uint64                     `json:"supply"`
	IsRarityCalculated     bool                       `json:"is_rarity_calculated"`
	Image                  string                     `json:"image"`
	Description            string                     `json:"description"`
	ContractScan           string                     `json:"contract_scan"`
	Discord                string                     `json:"discord"`
	Twitter                string                     `json:"twitter"`
	Website                string                     `json:"website"`
	Owners                 int64                      `json:"owners" gorm:"-"`
	MarketplaceCollections []NftMarketplaceCollection `json:"marketplace_collections" gorm:"-"`
	CreatedTime            time.Time                  `json:"created_time,omitempty"`
	LastUpdatedTime        time.Time                  `json:"last_updated_time,omitempty"`
	Chain                  *Chain                     `json:"chain" gorm:"foreignkey:ChainId;references:ChainId;<-:false"`
	Stats                  *ViewNftCollectionStats    `json:"stats" gorm:"-"`
	IsNotifySynced         bool                       `json:"is_notify_synced"`
	Marketplace            []NftListingMarketplace    `json:"marketplace" gorm:"-"`
}

func (n *NftCollection) MergeOpensea(c *OpenseaCollection) *NftCollection {
	if c == nil {
		return n
	}
	n.Name = c.Name
	n.Image = c.ImageUrl
	n.Description = c.Description
	n.Discord = c.DiscordUrl
	n.Twitter = "https://twitter.com/" + c.TwitterUsername
	n.Website = c.ExternalUrl

	return n
}

func (n *NftCollection) MergePaintswap(c *PaintswapCollection) *NftCollection {
	if c == nil {
		return n
	}
	n.Name = c.Name
	n.Description = c.Description
	n.Discord = c.Discord
	n.Twitter = "https://twitter.com/" + c.Twitter
	n.Website = c.Website
	n.Image = c.Poster

	return n
}

func (n *NftCollection) MergeNFTKey(c *NFTKeyCollection) *NftCollection {
	if c == nil {
		return n
	}
	n.Name = c.Name
	n.Description = c.Description
	n.Discord = c.Discord
	n.Twitter = "https://twitter.com/" + c.Twitter
	n.Website = c.Website
	n.Image = c.Poster

	return n
}

func (n *NftCollection) MergeQuixotic(c *QuixoticCollection) *NftCollection {
	if c == nil {
		return n
	}
	n.Name = c.Name
	n.Description = c.Description
	n.Image = c.ImageUrl
	n.Website = c.ExternalLink

	return n
}

type ViewNftCollectionStats struct {
	Address                  string  `json:"address"`
	PaymentToken             uint64  `json:"-"`
	FloorPrice               float64 `json:"-"`
	FloorPriceObj            *Price  `json:"floor_price"`
	OneDayVolume             float64 `json:"-"`
	OneDayVolumeObj          *Price  `json:"one_day_volume"`
	OneDayVolumeChange       float64 `json:"one_day_volume_change"`
	SevenDayVolume           float64 `json:"-"`
	SevenDayVolumeObj        *Price  `json:"seven_day_volume"`
	SevenDayVolumeChange     float64 `json:"seven_day_volume_change"`
	ThirtyDayVolume          float64 `json:"-"`
	ThirtyDayVolumeObj       *Price  `json:"thirty_day_volume"`
	ThirtyDayVolumeChange    float64 `json:"thirty_day_volume_change"`
	AllTimeVolume            float64 `json:"-"`
	AllTimeVolumeObj         *Price  `json:"all_time_volume"`
	AllTimeVolumeChange      float64 `json:"all_time_volume_change"`
	Token                    *Token  `json:"-"`
	OneDayVolumeChangeStr    string  `json:"one_day_volume_change_str"`
	SevenDayVolumeChangeStr  string  `json:"seven_day_volume_change_str"`
	ThirtyDayVolumeChangeStr string  `json:"thirty_day_volume_change_str"`
	AllTimeVolumeChangeStr   string  `json:"all_time_volume_change_str"`
}

func (stats *ViewNftCollectionStats) Standardize() {
	if stats == nil {
		return
	}
	stats.FloorPriceObj = &Price{Amount: fmt.Sprintf("%.0f", stats.FloorPrice)}
	stats.OneDayVolumeObj = &Price{Amount: fmt.Sprintf("%.0f", stats.OneDayVolume)}
	stats.SevenDayVolumeObj = &Price{Amount: fmt.Sprintf("%.0f", stats.SevenDayVolume)}
	stats.ThirtyDayVolumeObj = &Price{Amount: fmt.Sprintf("%.0f", stats.ThirtyDayVolume)}
	stats.AllTimeVolumeObj = &Price{Amount: fmt.Sprintf("%.0f", stats.AllTimeVolume)}
	// format volume change ratio
	stats.OneDayVolumeChangeStr = fmt.Sprintf("%.2f%%", stats.OneDayVolumeChange*100)
	stats.SevenDayVolumeChangeStr = fmt.Sprintf("%.2f%%", stats.SevenDayVolumeChange*100)
	stats.ThirtyDayVolumeChangeStr = fmt.Sprintf("%.2f%%", stats.ThirtyDayVolumeChange*100)
	stats.AllTimeVolumeChangeStr = fmt.Sprintf("%.2f%%", stats.AllTimeVolumeChange*100)
	if !strings.HasPrefix(stats.OneDayVolumeChangeStr, "-") {
		stats.OneDayVolumeChangeStr = fmt.Sprintf("+%s", stats.OneDayVolumeChangeStr)
	}
	if !strings.HasPrefix(stats.SevenDayVolumeChangeStr, "-") {
		stats.SevenDayVolumeChangeStr = fmt.Sprintf("+%s", stats.SevenDayVolumeChangeStr)
	}
	if !strings.HasPrefix(stats.ThirtyDayVolumeChangeStr, "-") {
		stats.ThirtyDayVolumeChangeStr = fmt.Sprintf("+%s", stats.ThirtyDayVolumeChangeStr)
	}
	if !strings.HasPrefix(stats.AllTimeVolumeChangeStr, "-") {
		stats.AllTimeVolumeChangeStr = fmt.Sprintf("+%s", stats.AllTimeVolumeChangeStr)
	}
	// assign model.Token object
	if stats.Token == nil {
		return
	}
	stats.FloorPriceObj.Token = *stats.Token
	stats.OneDayVolumeObj.Token = *stats.Token
	stats.SevenDayVolumeObj.Token = *stats.Token
	stats.ThirtyDayVolumeObj.Token = *stats.Token
	stats.AllTimeVolumeObj.Token = *stats.Token
}
