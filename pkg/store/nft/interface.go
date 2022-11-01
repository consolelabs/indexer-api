package nft

import "github.com/consolelabs/indexer-api/pkg/model"

type INft interface {
	// Token
	GetNftTokens(query NftTokenQuery) (tokens []model.NftToken, total int64, err error)

	// ViewNftTokens
	RefreshViewNFTTokens() error

	// Owner
	SaveOwner(owner *model.NftOwner) error
	DeleteOwnerByCollectionAddressTokenId(collectionAddress, tokenId string) error
	UpdateOwnerByCollectionAddressTokenId(collectionAddress, tokenId string, ownerAddress string) error

	// Transfer
	SaveTransfer(transfer *model.NftTransfer) error
	GetNftTokenTxHistory(collectionAddress, tokenId string) ([]model.NftTxHistory, error)

	// NftListing
	GetNftListing(query NftListingQuery) (listings []*model.NftListing, total int64, err error)
	GetNftListingByTokenID(collectionAddress, tokenID string) ([]model.NftListingMarketplace, error)
	UpdateListing(listing *model.NftListing) error
	SaveListing(listing *model.NftListing) error

	// NftCollection
	GetCollections(query NftCollectionQuery) (data []model.NftCollection, total int64, err error)
	SaveNftCollection(nftCollection *model.NftCollection) error
	GetCollectionsByWalletAddress(query WalletCollectionQuery) (data []model.NftCollection, total int64, err error)

	// Attribute
	GetAttributesByCollectionAddress(collectionAddress string) ([]model.NftTokenAttribute, error)
	UpdateAttributeCount(count uint64, address, traitType, value string) error

	// Marketplace Platform
	GetPlatformsByCollectionAddress(collectionAddress string) ([]model.MarketplacePlatform, error)

	// Collection snapshot
	GetNftMarketplaceCollectionSnapshots(query NftTickerQuery) ([]model.NftMarketplaceCollectionSnapshot, error)

	//Metadata
	GetNftMetadataAttributesIcon(query NftAttributeIconQuery) (icons []model.NftMetadataAttributesIcon, total int64, err error)

	// NftMarketplaceCollection
	GetNftMarketplaceCollection(collectionAddress string) ([]model.NftListingMarketplace, error)

	// RefreshViewNFTCollections
	RefreshViewNFTCollections() error

	// ViewNftCollectionStats
	RefreshViewNFTCollectionStats() error

	// ViewNftCollectionAttributes
	RefreshViewNFTCollectionAttributes() error
}
