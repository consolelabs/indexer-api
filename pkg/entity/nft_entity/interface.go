package nftentity

import (
	"github.com/consolelabs/indexer-api/pkg/model"
	"github.com/consolelabs/indexer-api/pkg/request"
	"github.com/consolelabs/indexer-api/pkg/response"
)

type INftEntity interface {
	GetTokenDetail(address, tokenId string) (token *model.NftToken, err error)
	GetNftCollectionTickers(address string, query request.GetNftTickersRequest) (*response.NftCollectionTickersData, error)
	GetNftTokenTickers(collectionAddress, tokenID string, req request.GetNftTickersRequest) (*response.NftTokenTickersData, error)
	GetNftTradingVolume() ([]response.NftTradingVolume, error)
	GetNftCollections(query request.GetNftCollectionsRequest) (collections []model.NftCollection, total int64, err error)
	GetNftTokens(address string, query request.GetNftTokensRequest) (tokens []model.NftToken, total int64, err error)
	GetNftTokensByWalletAddress(walletAddress string, query request.GetNftTokensByAddressRequest) (tokens []model.NftToken, total int64, err error)
	GetQueryOptions(collectionAddress string) (marketplaces []model.MarketplaceWithListingStats, attributesMetadata map[string][]model.NftTokenAttribute, traitCounts []model.NftTokenAttributeCount, err error)
	GetCollectionByAddress(collectionAddress string) (*model.NftCollection, error)
	ResyncNftCollection(collectionAddress string) error
	GetNftMetadataAttributesIcon(req request.GetNftAttributeIconQuery) (*response.NftMetadataAttributesIconResponse, error)
	ResyncNftCollectionMetadata(collectionAddress string, chainId int64) error
	ResyncNftCollectionRarity(collectionAddress string) error
	GetCollectionMetadata(collectionAddress string) (*model.NftCollection, error)
	GetTokenActivities(collectionAddress, tokenId string, req request.GetNftTokenActivitiesRequest) (activities []*model.NftListing, total int64, err error)
	GetTokenBidHistory(collectionAddress, tokenId string, req *model.Pagination) (bids []*model.NftBid, total int64, err error)
	GetTokenMetadata(collectionAddress, tokenId string) (token *response.GetNftTokenMetadataResponse, err error)
	GetNftListingSoldPrices(collectionAddress, tokenId, paymentToken, timeDuration, timeInterval, timeRoundUp string) (listingSoldPrices []*model.NftListingSoldPrice, err error)
	SearchNft(*request.SearchNftRequest) (collections []model.NftCollection, assets []model.NftToken, err error)
	GetCollectionSaleVolumeAndFloorPrice(collectionAddress, timeDuration string) ([]*model.NftCollectionSaleVolumeAndFloorPrice, error)
	GetListingAcrossPriceRanges(collectionAddress, paymentToken, priceRange string) (listingPriceRanges []*model.NftListingPriceRange, err error)
	GetListingAcrossRarityRankRanges(collectionAddress, paymentToken, rarityRankRange string) (listingPriceRanges []*model.NftListingRarityRankRange, err error)
	GetListingMarketCapAndVolumeByTimeRanges(collectionAddress, paymentToken, timeDuration, timeInterval, timeRoundUp string) (listingMarketCapAndVolumeTimeRanges []*model.NftListingMarketCapAndVolumeTimeRange, err error)
	GetNftCollectionsByWalletAddress(walletAddress string, query request.GetNftCollectionsByWalletAddressRequest) (collections []model.NftCollection, total int64, err error)
	GetListingUniqueOwnersByTimeRange(collectionAddress, timeDuration, timeInterval, timeRoundUp string) (listingUniqueOwnerTimeRanges []*model.NftListingUniqueOwnersTimeRange, err error)
	GetNftListingAvgPricesAndFloor(collectionAddress, tokenId, paymentToken, timeDuration, timeInterval, timeRoundUp string) (listingAvgPricesAndFloors []*model.NftListingAvgPriceAndFloor, err error)
	GetMarketSnapshotFloorPriceByTimeRanges(collectionAddress, paymentToken, timeDuration, timeInterval, timeRoundUp string) (marketplaceCollectionSnapshotFloorPriceTimeRanges []*model.NftMarketplaceCollectionSnapshotFloorPriceTimeRange, err error)
	GetMarketSnapshotMarketCapByTimeRanges(collectionAddress, paymentToken, timeDuration, timeInterval, timeRoundUp string) ([]*model.NftMarketplaceCollectionSnapshotMarketCapTimeRange, error)
	GetListingTotalVolumeByTimeRanges(collectionAddress, paymentToken, timeDuration, timeInterval, timeRoundUp string) ([]*model.NftListingTotalVolumeTimeRange, error)
	GetNftTokenTransactionHistory(collectionAddress, tokenId string) ([]model.NftTxHistory, error)
}
