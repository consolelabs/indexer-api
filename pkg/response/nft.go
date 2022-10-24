package response

import (
	"sort"
	"time"

	"github.com/consolelabs/indexer-api/pkg/model"
)

type NftTokenData struct {
	TokenId           string                         `json:"token_id"`
	CollectionAddress string                         `json:"collection_address"`
	Name              string                         `json:"name"`
	Description       string                         `json:"description"`
	Amount            string                         `json:"amount"`
	Image             string                         `json:"image"`
	ImageCDN          string                         `json:"image_cdn"`
	ThumbnailCDN      string                         `json:"thumbnail_cdn"`
	ImageContentType  string                         `json:"image_content_type"`
	Attributes        NftTokenAttributeResponseSlice `json:"attributes"`
	Rarity            *NftTokenRarityResponse        `json:"rarity"`
	Owner             *model.NftOwner                `json:"owner"`
	Marketplace       []model.NftListingMarketplace  `json:"marketplace"`
}

type NftTokenRarityResponse struct {
	Rank   uint64 `json:"rank"`
	Score  string `json:"score"`
	Total  uint64 `json:"total"`
	Rarity string `json:"rarity,omitempty"`
}

type NftTokenAttributeResponse struct {
	TraitType string `json:"trait_type"`
	Value     string `json:"value"`
	Count     uint64 `json:"count"`
	Rarity    string `json:"rarity"`
	Frequency string `json:"frequency"`
}

type NftTokenAttributeResponseSlice []NftTokenAttributeResponse

func (v NftTokenAttributeResponseSlice) Less(i, j int) bool {
	return v[i].Count > v[j].Count
}
func (v NftTokenAttributeResponseSlice) Len() int {
	return len(v)
}
func (v NftTokenAttributeResponseSlice) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

func ToNftTokenResponse(token model.NftToken) (resp NftTokenData) {
	resp = NftTokenData{
		TokenId:           token.TokenId,
		CollectionAddress: token.CollectionAddress,
		Name:              token.Name,
		Description:       token.Description,
		Amount:            token.Amount,
		Image:             token.Image,
		ImageCDN:          token.ImageCDN,
		ThumbnailCDN:      token.ThumbnailCDN,
		ImageContentType:  token.ImageContentType,
		Owner:             token.Owner,
		Rarity: &NftTokenRarityResponse{
			Rank:   token.Rarity.Rank,
			Score:  token.Rarity.Score,
			Total:  token.Rarity.Total,
			Rarity: token.Rarity.Rarity,
		},
		Marketplace: token.ListingMarketplace,
	}
	resp.Attributes = make(NftTokenAttributeResponseSlice, len(token.Attributes))
	for idx, att := range token.Attributes {
		resp.Attributes[idx] = NftTokenAttributeResponse{
			TraitType: att.TraitType,
			Value:     att.Value,
			Count:     att.Count,
			Rarity:    att.Rarity,
			Frequency: att.Frequency,
		}
	}
	sort.Sort(resp.Attributes)

	return resp
}

type TokenTickers struct {
	Timestamps []int64       `json:"timestamps"`
	Prices     []model.Price `json:"prices"`
}

type NftCollectionTickersData struct {
	Tickers         TokenTickers `json:"tickers"`
	Name            string       `json:"name"`
	Address         string       `json:"address"`
	Chain           *model.Chain `json:"chain"`
	Marketplaces    []string     `json:"marketplaces"`
	Owners          int64        `json:"owners"`
	Items           int64        `json:"items"`
	CollectionImage string       `json:"collection_image"`
	FloorPrice      *model.Price `json:"floor_price"`
	TotalVolume     *model.Price `json:"total_volume"`
	LastSalePrice   *model.Price `json:"last_sale_price"`
	PriceChange1d   string       `json:"price_change_1d"`
	PriceChange7d   string       `json:"price_change_7d"`
	PriceChange30d  string       `json:"price_change_30d"`
}

type NftTokenTickersData struct {
	Tickers           TokenTickers `json:"tickers"`
	Name              string       `json:"name"`
	TokenId           string       `json:"token_id"`
	CollectionAddress string       `json:"collection_address"`
	Description       string       `json:"description"`
	Image             string       `json:"image"`
	ImageCDN          string       `json:"image_cdn"`
	RarityRank        uint64       `json:"rarity_rank"`
	RarityScore       string       `json:"rarity_score"`
	RarityTier        string       `json:"rarity_tier"`
	FloorPrice        *model.Price `json:"floor_price"`
	LastSalePrice     *model.Price `json:"last_sale_price"`
	PriceChange1d     string       `json:"price_change_1d"`
	PriceChange7d     string       `json:"price_change_7d"`
	PriceChange30d    string       `json:"price_change_30d"`
}

type NftTokenTickersDataResponse struct {
	Data NftTokenTickersData `json:"data"`
}

type GetNftQueryOptionsData struct {
	Marketplaces []model.MarketplaceWithListingStats  `json:"marketplaces"`
	Attributes   map[string][]model.NftTokenAttribute `json:"attributes"`
	TraitCounts  []model.NftTokenAttributeCount       `json:"trait_counts"`
}

type NftMetadataAttributesIcon struct {
	Id          int    `json:"id,omitempty"`
	TraitType   string `json:"trait_type,omitempty"`
	DiscordIcon string `json:"discord_icon,omitempty"`
	UnicodeIcon string `json:"unicode_icon,omitempty"`
}
type GetNftTokenMetadataData struct {
	LatestListing         *model.NftListing `json:"latest_listing"`
	LastSale              *model.Price      `json:"last_sale"`
	MaxPrice              *model.Price      `json:"max_price"`
	MinPrice              *model.Price      `json:"min_price"`
	TotalSales            int64             `json:"total_sales"`
	TotalOwners           int64             `json:"total_owners"`
	Dob                   time.Time         `json:"dob"`
	CurrentHoldTimeInSecs float64           `json:"current_hold_time_in_secs"`
	LongestHoldTimeInSecs float64           `json:"longest_hold_time_in_secs"`
	Creator               string            `json:"creator"`
	Owner                 *model.NftOwner   `json:"owner"`
}

// Data type for swagger
type SearchNftResponseData struct {
	Collections []model.NftCollection `json:"collections"`
	Assets      []model.NftToken      `json:"assets"`
}

type SearchNftResponse struct {
	Data SearchNftResponseData `json:"data,omitempty"`
}
type GetNftCollectionSaleVolumeAndFloorPrice struct {
	*PaginationResponse `json:",omitempty"`
	Data                []*model.NftCollectionSaleVolumeAndFloorPrice `json:"data,omitempty"`
}
type GetNftListingAcrossPriceRangesResponse struct {
	*PaginationResponse `json:",omitempty"`
	Data                []*model.NftListingPriceRange `json:"data,omitempty"`
}
type NftCollectionTickersResponse struct {
	Data NftCollectionTickersData `json:"data,omitempty"`
}
type GetNftCollectionsResponse struct {
	*PaginationResponse `json:",omitempty"`
	Data                []model.NftCollection `json:"data,omitempty"`
}

type GetNftCollectionMetadataResponse struct {
	Data *model.NftCollection `json:"data,omitempty"`
}

type GetNftTokensResponse struct {
	*PaginationResponse `json:",omitempty"`
	Data                []model.NftToken `json:"data,omitempty"`
}
type GetNftQueryOptionsResponse struct {
	Data GetNftQueryOptionsData `json:"data,omitempty"`
}
type NftMetadataAttributesIconResponse struct {
	*PaginationResponse `json:",omitempty"`
	Data                []model.NftMetadataAttributesIcon `json:"data,omitempty"`
}
type GetNftTokenActivitiesResponse struct {
	*PaginationResponse `json:",omitempty"`
	Data                []*model.NftListing `json:"data,omitempty"`
}
type GetNftTokenBidHistoryResponse struct {
	*PaginationResponse `json:",omitempty"`
	Data                []*model.NftBid `json:"data,omitempty"`
}
type GetNftTokenMetadataResponse struct {
	Data GetNftTokenMetadataData `json:"data,omitempty"`
}
type GetNftTokenListingSoldPricesResponse struct {
	*PaginationResponse `json:",omitempty"`
	Data                []*model.NftListingSoldPrice `json:"data,omitempty"`
}

type GetNftTokenListingAvgPricesAndFloorResponse struct {
	*PaginationResponse `json:",omitempty"`
	Data                []*model.NftListingAvgPriceAndFloor `json:"data,omitempty"`
}
type NftTokenResponse struct {
	Data NftTokenData `json:"data,omitempty"`
}

type GetListingAcrossRarityRankRangesResponse struct {
	Data []*model.NftListingRarityRankRange `json:"data,omitempty"`
}

type GetListingMarketCapAndVolumeTimeRangesResponse struct {
	Data []*model.NftListingMarketCapAndVolumeTimeRange `json:"data,omitempty"`
}

type GetMarketSnapshotFloorPriceByTimeRangesResponse struct {
	Data []*model.NftMarketplaceCollectionSnapshotFloorPriceTimeRange `json:"data,omitempty"`
}

type GetMarketSnapshotMarketCapByTimeRangesResponse struct {
	Data []*model.NftMarketplaceCollectionSnapshotMarketCapTimeRange `json:"data,omitempty"`
}

type GetListingTotalVolumeByTimeRangesResponse struct {
	Data []*model.NftListingTotalVolumeTimeRange `json:"data,omitempty"`
}

type GetListingUniqueOwnersByTimeRangeResponse struct {
	Data []*model.NftListingUniqueOwnersTimeRange `json:"data,omitempty"`
}

type GetNftTokenTransactionHistory struct {
	Data []model.NftTxHistory `json:"data,omitempty"`
}
