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
	GetNftMetadataAttributesIcon(req request.GetNftAttributeIconQuery) (*response.NftMetadataAttributesIconResponse, error)
	GetCollectionMetadata(collectionAddress string) (*model.NftCollection, error)
	GetTokenActivities(collectionAddress, tokenId string, req request.GetNftTokenActivitiesRequest) (activities []*model.NftListing, total int64, err error)
	GetNftTokenTransactionHistory(collectionAddress, tokenId string) ([]model.NftTxHistory, error)
}
