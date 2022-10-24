package paintswap

import "github.com/consolelabs/indexer-api/pkg/model"

type IService interface {
	GetPaintswapCollection(collectionAddress string) (paintswapCollection *model.PaintswapCollection, err error)
}
