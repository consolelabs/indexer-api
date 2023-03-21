package token

import "github.com/consolelabs/indexer-api/pkg/model"

type IToken interface {
	GetTokenPriceDetail(symbol string) (tokenPriceDetail *model.TokenPriceDetail, err error)
}
