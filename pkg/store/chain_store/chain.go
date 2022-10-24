package chainstore

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/consolelabs/indexer-api/pkg/model"
	"github.com/consolelabs/indexer-api/pkg/request"
)

type store struct {
	db *gorm.DB
}

func New(db *gorm.DB) IChain {
	return &store{
		db: db,
	}
}

func (s *store) Get(query request.GetChainsQuery) ([]model.Chain, int64, error) {
	var chains []model.Chain
	db := s.db
	if query.ChainId != nil {
		db = db.Where("chain_id = ?", *query.ChainId)
	}
	if query.Name != nil {
		db = db.Where("lower(name) like lower(?)", fmt.Sprintf("%%%s%%", *query.Name))
	}
	if query.Symbol != nil {
		db = db.Where("lower(symbol) = lower(?)", *query.Symbol)
	}
	if query.IsEvm != nil {
		db = db.Where("is_evm = ?", *query.IsEvm)
	}

	var total int64
	if err := db.Model(&model.Chain{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if query.Sort != "" {
		db = db.Order(query.Sort)
	}
	if err := db.Offset(int(query.Page * query.Size)).Limit(int(query.Size)).Find(&chains).Error; err != nil {
		return nil, 0, err
	}
	return chains, total, nil
}
