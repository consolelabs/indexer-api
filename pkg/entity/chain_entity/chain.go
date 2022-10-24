package chainentity

import (
	"fmt"

	"github.com/consolelabs/indexer-api/pkg/model"
	"github.com/consolelabs/indexer-api/pkg/request"
	"github.com/consolelabs/indexer-api/pkg/response"
	"github.com/consolelabs/indexer-api/pkg/service"
	"github.com/consolelabs/indexer-api/pkg/store"
)

type chainEntity struct {
	store   *store.Store
	service *service.Service
}

func New(store *store.Store, service *service.Service) IChainEntity {
	return &chainEntity{
		store:   store,
		service: service,
	}
}

func (e *chainEntity) GetChains(query request.GetChainsQuery) (*response.GetChainsResponse, error) {
	query.Standardize()
	data, total, err := e.store.Chain.Get(query)
	if err != nil {
		return nil, fmt.Errorf("[Entity][GetChains] failed while [Get] : %s", err.Error())
	}
	return &response.GetChainsResponse{
		Data: data,
		PaginationResponse: &response.PaginationResponse{
			Pagination: model.Pagination{
				Page: query.Page,
				Size: query.Size,
			},
			Total: total,
		},
	}, nil
}
