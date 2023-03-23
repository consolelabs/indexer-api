package token

import (
	"github.com/consolelabs/indexer-api/pkg/model"
	"github.com/consolelabs/indexer-api/pkg/request"
)

type ITokenEntity interface {
	GetConvertTokenPrice(req request.ConvertTokenPrice) (*model.ConvertTokenPrice, error)
}
