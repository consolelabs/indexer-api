package contractstore

import (
	"gorm.io/gorm"

	"github.com/consolelabs/indexer-api/pkg/model"
)

type store struct {
	db *gorm.DB
}

func New(db *gorm.DB) IContract {
	return &store{
		db: db,
	}
}

func (s *store) Get() ([]*model.Contract, error) {
	var contracts []*model.Contract
	// Note: at first, we want to serve ftm for lightweightness
	var ss string = "ERC721"
	return contracts, s.db.Where("type = ?", ss).Order("chain_id desc").Find(&contracts).Error
}

func (s *store) Save(new *model.Contract) error {
	return s.db.Save(new).Error
}

// func (s *store) Update(new *model.Contract) error {
// 	return s.db.Save(new).Error
// }

func (s *store) GetByAddress(address string) (*model.Contract, error) {
	var contract model.Contract
	err := s.db.Table("contract").Where("address = ?", address).First(&contract).Error
	if err != nil {
		return nil, err
	}
	return &contract, nil
}

func (s *store) GetByAddressAndChainId(address string, chainId int64) (contract *model.Contract, err error) {
	return contract, s.db.Table("contract").Where("address = ? and chain_id = ?", address, chainId).First(&contract).Error
}
