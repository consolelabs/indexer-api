package token

import "github.com/consolelabs/indexer-api/pkg/model"

type ITokenEntity interface {
	GetConvertTokenPrice(amount, from, to string) (*model.ConvertTokenPrice, error)
}
