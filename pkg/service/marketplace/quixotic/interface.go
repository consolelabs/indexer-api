package quixotic

import "github.com/consolelabs/indexer-api/pkg/model"

type IService interface {
	GetQuixoticCollection(collectionAddress string) (*model.QuixoticCollection, error)
}
