package token

import (
	"github.com/consolelabs/indexer-api/pkg/model"
	"gorm.io/gorm"
)

type store struct {
	db *gorm.DB
}

func New(db *gorm.DB) IToken {
	return &store{
		db: db,
	}
}

func (s *store) GetTokenPriceDetail(symbol string) (*model.TokenPriceDetail, error) {
	db := s.db.Table("token AS t").Select(
		"t.id",
		"t.name",
		"t.symbol",
		"t.decimals",
		"tp.price",
	).Joins("LEFT JOIN token_price tp ON t.id = tp.token_id")

	rows, err := db.Where("t.symbol = ?", symbol).Rows()
	if err != nil {
		return nil, err
	}
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()
	tokenPriceDetail := model.TokenPriceDetail{}
	for rows.Next() {
		err := rows.Scan(
			&tokenPriceDetail.TokenId,
			&tokenPriceDetail.Name,
			&tokenPriceDetail.Symbol,
			&tokenPriceDetail.Decimals,
			&tokenPriceDetail.Price,
		)
		if err != nil {
			return nil, err

		}
	}
	return &tokenPriceDetail, nil
}
