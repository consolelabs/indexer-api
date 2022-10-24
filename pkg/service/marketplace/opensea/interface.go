package opensea

import "github.com/consolelabs/indexer-api/pkg/model"

type IService interface {
	GetOpenseaCollection(collectionAddress string) (openseaCollection *model.OpenseaCollection, err error)
}
