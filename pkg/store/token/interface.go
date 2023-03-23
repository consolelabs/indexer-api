package token

import "github.com/consolelabs/indexer-api/pkg/model"

type IToken interface {
	// token price
	GetTokenBySymbol(symbol string) (token *model.Token, err error)
	GetTokenPriceDetail(tokenId int64, time string) (tokenPriceDetail *model.TokenPrice, err error)
	CreateTokenPrice(model *model.TokenPrice) error

	GetListToken() (tokens []model.Token, err error)
	GetLatestSnapshot(id int64, source string) (snapshot []model.TokenHistoryPriceSnapshot, err error)
	Save(model *model.TokenHistoryPriceSnapshot) error
}
