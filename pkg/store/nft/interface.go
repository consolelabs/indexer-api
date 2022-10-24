package nft

import "github.com/consolelabs/indexer-api/pkg/model"

type INft interface {
	// Token
	UpsertToken(token *model.NftToken, upsertAttributes bool) error
	GetNftTokens(query NftTokenQuery) (tokens []model.NftToken, total int64, err error)
	GetTokensByWalletAddress(query WalletTokenQuery) (data []model.NftToken, total int64, err error)
	DeleteNftToken(query NftTokenQuery) error

	// ViewNftTokens
	RefreshViewNFTTokens() error

	// Owner
	SaveOwner(owner *model.NftOwner) error
	GetNftOwners(query NftOwnerQuery) (owners []model.NftOwner, total int64, err error)
	DeleteOwnerByCollectionAddressTokenId(collectionAddress, tokenId string) error
	UpdateOwnerByCollectionAddressTokenId(collectionAddress, tokenId string, ownerAddress string) error
	DeleteNftOwnerByCollectionAddress(collecionAddress string) error

	// Transfer
	SaveTransfer(transfer *model.NftTransfer) error
	DeleteNftTransferByContractAddress(collecionAddress string) error
	GetNftTransfers(query NftTransferQuery) (transfers []model.NftTransfer, total int64, err error)
	GetNftTokenTxHistory(collectionAddress, tokenId string) ([]model.NftTxHistory, error)

	// NftListing
	// GetListingByTokenIdAndContractAddress(tokenId, contractAddress string) (*model.NftListing, error)
	GetNftListing(query NftListingQuery) (listings []*model.NftListing, total int64, err error)
	GetNftListingByTokenID(collectionAddress, tokenID string) ([]model.NftListingMarketplace, error)
	GetNftListingSoldPrices(collectionAddress, tokenId, paymentToken, timeDuration, timeInterval, timeRoundUp string) (listingSoldPrices []*model.NftListingSoldPrice, err error)
	UpdateListing(listing *model.NftListing) error
	SaveListing(listing *model.NftListing) error
	GetListingAcrossPriceRanges(collectionAddress, paymentToken, priceRange string) (listingPriceRanges []*model.NftListingPriceRange, err error)
	GetListingAcrossRarityRankRanges(collectionAddress, paymentToken, rarityRankRange string) (listingRarityRankRanges []*model.NftListingRarityRankRange, err error)
	GetListingMarketCapAndVolumeByTimeRanges(collectionAddress, paymentToken, timeDuration, timeInterval, timeRoundUp string) (listingMarketCapAndVolumeTimeRanges []*model.NftListingMarketCapAndVolumeTimeRange, err error)
	GetListingUniqueOwnersByTimeRange(collectionAddress, timeDuration, timeInterval, timeRoundUp string) (listingUniqueOwnerTimeRanges []*model.NftListingUniqueOwnersTimeRange, err error)
	GetNftListingAvgPricesAndFloor(collectionAddress, tokenId, paymentToken, timeDuration, timeInterval, timeRoundUp string) (listingAvgPricesAndFloors []*model.NftListingAvgPriceAndFloor, err error)
	GetListingTotalVolumeByTimeRanges(collectionAddress, paymentToken, timeDuration, timeInterval, timeRoundUp string) (listingTotalVolumeTimeRanges []*model.NftListingTotalVolumeTimeRange, err error)

	// NftCollection
	GetCollectionByAddress(address string) (*model.NftCollection, error)
	UpdateCollection(collection *model.NftCollection) error
	GetCollections(query NftCollectionQuery) (data []model.NftCollection, total int64, err error)
	SaveNftCollection(nftCollection *model.NftCollection) error
	GetCollectionSaleVolumeAndFloorPrice(collectionAddress string, timeDuration string) (saleVolumeAndFloorPrices []*model.NftCollectionSaleVolumeAndFloorPrice, err error)
	GetCollectionsByWalletAddress(query WalletCollectionQuery) (data []model.NftCollection, total int64, err error)

	// Attribute
	GetAttributesByCollectionAddress(collectionAddress string) ([]model.NftTokenAttribute, error)
	AttributeUpsertOne(tokenAttribute model.NftTokenAttribute) error
	GetDistinctAttributesByCollectionAddress(collectionAddress string) ([]model.NftTokenAttribute, error)
	UpdateAttributeCount(count uint64, address, traitType, value string) error
	DeleteNftTokenAttributesByCollectionAddress(collecionAddress string) error
	GetCollectionTraitsCount(collectionAddress string) (attrCount []model.NftTokenAttributeCount, err error)
	GetAttributeByCollectionAddressTokenID(collectionAddress, tokenID string) ([]model.NftTokenAttribute, error)

	// Marketplace Platform
	GetPlatformByName(name string) (*model.MarketplacePlatform, error)
	GetPlatformsByCollectionAddress(collectionAddress string) ([]model.MarketplacePlatform, error)

	// Collection snapshot
	GetNftMarketplaceCollectionSnapshots(query NftTickerQuery) ([]model.NftMarketplaceCollectionSnapshot, error)
	GetMarketSnapshotFloorPriceByTimeRanges(collectionAddress, paymentToken, timeDuration, timeInterval, timeRoundUp string) (marketplaceCollectionSnapshotFloorPriceTimeRanges []*model.NftMarketplaceCollectionSnapshotFloorPriceTimeRange, err error)
	GetMarketSnapshotMarketCapByTimeRanges(collectionAddress, paymentToken, timeDuration, timeInterval, timeRoundUp string) (marketplaceCollectionSnapshotMarketCapTimeRanges []*model.NftMarketplaceCollectionSnapshotMarketCapTimeRange, err error)

	//Metadata
	GetNftMetadataAttributesIcon(query NftAttributeIconQuery) (icons []model.NftMetadataAttributesIcon, total int64, err error)

	// NftMarketplaceCollection
	UpsertNftMarketplaceCollection(model *model.NftMarketplaceCollection) error
	GetMarketplaceWithListingStats(collectionAddress string) (marketplacesWithListingStats []model.MarketplaceWithListingStats, err error)
	GetNftMarketplaceCollection(collectionAddress string) ([]model.NftListingMarketplace, error)

	// RefreshViewNFTCollections
	RefreshViewNFTCollections() error

	// ViewNftCollectionStats
	RefreshViewNFTCollectionStats() error

	// ViewNftCollectionAttributes
	RefreshViewNFTCollectionAttributes() error
}
