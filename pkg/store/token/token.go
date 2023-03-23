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

func (s *store) GetTokenPriceDetail(tokenId int64, time string) (tokenPriceDetail *model.TokenPrice, err error) {
	return tokenPriceDetail, s.db.Table("token_price").Where("token_id = ? and time >= ?", tokenId, time).First(&tokenPriceDetail).Error
}

func (s *store) GetTokenBySymbol(symbol string) (token *model.Token, err error) {
	return token, s.db.Table("token").Where("symbol = ?", symbol).First(&token).Error
}
func (s *store) CreateTokenPrice(model *model.TokenPrice) error {
	return s.db.Table("token_price").Create(&model).Error
}

func (s *store) GetListToken() (tokens []model.Token, err error) {
	return tokens, s.db.Table("token").Find(&tokens).Error
}

func (s *store) GetLatestSnapshot(id int64, source string) (snapshot []model.TokenHistoryPriceSnapshot, err error) {
	return snapshot, s.db.Table("token_history_price_snapshot").Where("token_id = ? AND source = ?", id, source).Order("time DESC").Limit(1).Find(&snapshot).Error
}

func (s *store) Save(model *model.TokenHistoryPriceSnapshot) error {
	return s.db.Table("token_history_price_snapshot").Create(&model).Error
}
