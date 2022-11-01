package nft

import (
	"github.com/consolelabs/indexer-api/pkg/model"
	"gorm.io/gorm"
)

type store struct {
	db *gorm.DB
}

func New(db *gorm.DB) ITableNft {
	return &store{
		db: db,
	}
}

// TODO: fixme
func (s *store) GetCollections() ([]model.NftCollection, error) {
	var collections []model.NftCollection
	return collections, s.db.Table("nft_collection").Where("is_notify_synced = false").Find(&collections).Error
}

func (s *store) GetTokensByCollectionAddress(collectionAddress string) ([]model.NftToken, error) {
	var tokens []model.NftToken
	return tokens, s.db.Table("nft_token").Where("collection_address = ?", collectionAddress).Find(&tokens).Error
}

// func (s *store) UpdateCollection(collection *model.NftCollection) error {
// 	return s.db.Table("nft_collection").Where("address = ?", collection.Address).Save(&collection).Error
// }

func (s *store) GetNftTokens(collectionAddress string, tokenId string) (model.NftToken, error) {
	var token model.NftToken
	return token, s.db.Table("nft_token").Where("collection_address = ? AND token_id = ?", collectionAddress, tokenId).Find(&token).Error
}
