package chain_explorer

import "github.com/consolelabs/indexer-api/pkg/model"

type IService interface {
	GetCreationBlockNumber(model.Contract) (int64, error)
}
