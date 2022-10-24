package chainstore

import (
	"github.com/consolelabs/indexer-api/pkg/model"
	"github.com/consolelabs/indexer-api/pkg/request"
)

type IChain interface {
	Get(query request.GetChainsQuery) (data []model.Chain, total int64, err error)
}
