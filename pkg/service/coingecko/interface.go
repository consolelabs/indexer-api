package coingecko

import "github.com/consolelabs/indexer-api/pkg/model"

type IService interface {
	GetHistoryPrice(coingeckoId, date string) (*model.GetTokenHistoryPrice, error)
}
