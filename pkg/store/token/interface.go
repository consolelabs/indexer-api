package token

import "github.com/consolelabs/indexer-api/pkg/model"

type IToken interface {
	GetTokenPriceDetail(symbol string) (*model.TokenPriceDetail, error)
}
