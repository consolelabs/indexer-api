package token

import "github.com/consolelabs/indexer-api/pkg/model"

type IToken interface {
	// token price
	GetTokenPriceDetail(symbol string) (tokenPriceDetail *model.TokenPriceDetail, err error)

	GetListToken() (tokens []model.Token, err error)
	GetLatestSnapshot(id int64, source string) (snapshot []model.TokenHistoryPriceSnapshot, err error)
	Save(model *model.TokenHistoryPriceSnapshot) error
	Upsert(model *model.TokenHistoryPriceSnapshot) error
}
