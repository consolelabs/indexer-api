package nft

import "github.com/consolelabs/indexer-api/pkg/model"

type ITableNft interface {
	GetCollections() ([]model.NftCollection, error)
	GetTokensByCollectionAddress(collectionAddress string) ([]model.NftToken, error)
	UpdateCollection(collection *model.NftCollection) error
	GetNftTokens(collectionAddress string, tokenId string) (model.NftToken, error)
}
