package request

import (
	"strings"

	"github.com/consolelabs/indexer-api/pkg/model"
)

type GetNftCollectionsRequest struct {
	*model.Pagination
	Sort        string  `form:"sort" binding:"omitempty,oneof=one_day_volume -one_day_volume seven_day_volume -seven_day_volume thirty_day_volume -thirty_day_volume all_time_volume -all_time_volume floor_price -floor_price supply -supply name -name owners -owners" enums:"one_day_volume,-one_day_volume,seven_day_volume,-seven_day_volume,thirty_day_volume,-thirty_day_volume,all_time_volume,-all_time_volume,floor_price,-floor_price,supply,-supply,name,-name,owners,-owners"`
	Name        *string `json:"name" form:"name"`                                                                                                                        // search collection name
	Address     *string `json:"address" form:"address" binding:"omitempty"`                                                                                              // collection address
	Marketplace *string `json:"marketplace" form:"marketplace" binding:"omitempty,oneof=paintswap nftkey opensea looksrare"  enums:"paintswap,nftkey,opensea,looksrare"` // marketplace
}

type GetNftTokensRequest struct {
	*model.Pagination
	Sort          string   `json:"sort" form:"sort" binding:"omitempty,oneof=name -name token_id -token_id amount -amount rarity_rank -rarity_rank rarity_score -rarity_score rarity_tier -rarity_tier" enums:"name,-name,token_id,-token_id,amount,-amount,rarity_rank,-rarity_rank,rarity_score,-rarity_score,rarity_tier,-rarity_tier"` // sort by
	TokenId       *string  `json:"token_id" form:"token_id"`                                                                                                                                                                                                                                                                               // token ID
	Name          *string  `json:"name" form:"name"`                                                                                                                                                                                                                                                                                       // token name
	Currency      *string  `json:"currency" form:"currency" binding:"omitempty,oneof=eth ftm" enums:"eth,ftm"`                                                                                                                                                                                                                             // payment token/currency
	ListingType   *string  `json:"listing_type" form:"listing_type" binding:"omitempty,oneof=buy_now auction" enums:"buy_now,auction"`                                                                                                                                                                                                     // listing type
	ListingStatus *string  `json:"listing_status" form:"listing_status" binding:"omitempty,oneof=listed sold cancelled" enums:"listed,sold,cancelled"`                                                                                                                                                                                     // listing status
	PriceMin      *float64 `json:"price_min" form:"price_min" example:"0"`                                                                                                                                                                                                                                                                 // listing price min
	PriceMax      *float64 `json:"price_max" form:"price_max" example:"10" binding:"omitempty,gtefield=PriceMax"`                                                                                                                                                                                                                          // listing price max
	RarityRankMin *uint64  `json:"rarity_rank_min" form:"rarity_rank_min"`                                                                                                                                                                                                                                                                 // rarity rank min
	RarityRankMax *uint64  `json:"rarity_rank_max" form:"rarity_rank_max" binding:"omitempty,gtefield=RarityRankMin"`                                                                                                                                                                                                                      // rarity rank max
	Marketplaces  []string `json:"marketplaces" form:"marketplaces" swaggerignore:"true"`
	Traits        []string `json:"traits" form:"traits" swaggerignore:"true"` // Traits example mask=Television Head,Zorro Mask
	TraitsCount   []uint64 `json:"traits_count" form:"traits_count" swaggerignore:"true"`
}

// @Param       sort                    query     string                               false  "[field][option]<br>Available options: + (asc), - (desc)<br>Available fields: collection_address,token_id,name,amount,rarity_rank,rarity_score,rarity_tier"
// @Param       collection_addresses    query     []string                             false "collection addresses" collectionFormat(multi)
type GetNftTokensByAddressRequest struct {
	*model.Pagination
	Sort                string   `json:"sort" form:"sort" binding:"omitempty,oneof=collection_address token_id name amount rarity_rank rarity_score rarity_tier -collection_address -token_id -name -amount -rarity_rank -rarity_score -rarity_tier" enums:"collection_address,token_id,name,amount,rarity_rank,rarity_score,rarity_tier,-collection_address,-token_id,-name,-amount,-rarity_rank,-rarity_score,-rarity_tier"` // sort by
	CollectionAddresses []string `json:"collection_addresses" form:"collection_addresses" binding:"dive"`
	ChainId             *int64   `json:"chain_id" form:"chain_id"` // chain ID
}

type GetNftTickersRequest struct {
	From int64 `json:"from" form:"from" binding:"required"`                                                           // From date in UNIX millisecond (eg. 1392577232000)
	To   int64 `json:"to" form:"to" binding:"required,fieldgte=From 810000" msg:"to - from must >= 86400000 (1 day)"` // To date in UNIX millisecond (eg. 1422577232000)
}

func (req *GetNftTokensRequest) Standardize() {
	if req.Marketplaces != nil {
		var marketplaces []string
		for _, m := range req.Marketplaces {
			if m == "" {
				continue
			}
			marketplaces = append(marketplaces, strings.ToLower(m))
		}
		req.Marketplaces = marketplaces
	}
	if req.Traits != nil {
		var traits []string
		for _, t := range req.Traits {
			if !strings.Contains(t, "=") {
				continue
			}
			traits = append(traits, strings.ToLower(t))
		}
		req.Traits = traits
	}
	req.Pagination.Standardize()
}

type GetNftTokenActivitiesRequest struct {
	*model.Pagination
	Marketplace string `json:"marketplace"` // marketplace name
}

type SearchNftRequest struct {
	Text  string `json:"text" form:"text" binding:"required" msg:"text is required to search nft"`
	Limit int64  `json:"limit" form:"limit"`
}

func (req *SearchNftRequest) Standardize() {
	if req.Limit <= 0 || req.Limit >= 50 {
		req.Limit = 10
	}
}

type GetNftAttributeIconQuery struct {
	TraitTypes []string `json:"trait_type" form:"trait_type" msg:"invalid trait type"`
}

func (req *GetNftAttributeIconQuery) Standardize() {
	for idx, trait := range req.TraitTypes {
		req.TraitTypes[idx] = strings.ToLower(trait)
	}
}

type GetNftCollectionsByWalletAddressRequest struct {
	*model.Pagination
	Sort string `json:"sort" form:"sort" binding:"omitempty,oneof=name -name supply -supply owners -owners" enums:"name,-name,supply,-supply,owners,-owners"` // sort by
}
