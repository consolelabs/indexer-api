package token

import (
	"gorm.io/gorm"

	"github.com/consolelabs/indexer-api/pkg/model"
)

type store struct {
	db *gorm.DB
}

func New(db *gorm.DB) IToken {
	return &store{
		db: db,
	}
}

func (s *store) GetTokenPriceDetail(symbol string) (tokenPriceDetail *model.TokenPriceDetail, err error) {
	return tokenPriceDetail, s.db.Table("token_price").Select("token_price.price, token.decimals, token.name").Joins("LEFT JOIN token ON token_price.token_id = token.id").Where("token.symbol = ?", symbol).First(&tokenPriceDetail).Error
}