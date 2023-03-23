package birdeye

import "github.com/consolelabs/indexer-api/pkg/model"

type IService interface {
	GetTokenPriceByAddress(token *model.Token, date string) (*model.BirdeyeHistoryPrice, error)
	GetCurrentPrice(tokenAddress string) (*model.BirdeyeCurrentPriceData, error)
}
