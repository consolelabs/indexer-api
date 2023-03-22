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

func (s *store) GetTokenPriceDetail(symbol string) (tokenPriceDetail *model.TokenPriceDetail, err error) {
	return tokenPriceDetail, s.db.Table("token_price").Select("token_price.price, token.decimals, token.name").Joins("LEFT JOIN token ON token_price.token_id = token.id").Where("token.symbol = ?", symbol).First(&tokenPriceDetail).Error
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

func (s *store) Upsert(t *model.TokenHistoryPriceSnapshot) error {
	var snapshot *model.TokenHistoryPriceSnapshot
	err := s.db.Table("token_history_price_snapshot").Where("token_id = ? and source = ? and time = ?", t.TokenId, t.Source, t.Time).Find(&snapshot).Error
	if err != nil {
		return err
	}
	if snapshot.TokenId == t.TokenId {
		return nil
	}
	return s.db.Table("token_history_price_snapshot").Create(&t).Error
}
