package chainentity

import (
	"github.com/consolelabs/indexer-api/pkg/request"
	"github.com/consolelabs/indexer-api/pkg/response"
)

type IChainEntity interface {
	GetChains(query request.GetChainsQuery) (*response.GetChainsResponse, error)
}
