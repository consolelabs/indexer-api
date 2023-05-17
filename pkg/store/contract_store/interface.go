package contractstore

import "github.com/consolelabs/indexer-api/pkg/model"

type IContract interface {
	Get() (contracts []*model.Contract, err error)
	Save(new *model.Contract) (err error)
	// Update(new *model.Contract) (err error)
	GetByAddress(address string) (contract *model.Contract, err error)
	GetByAddressAndChainId(address string, chainId int64) (contract *model.Contract, err error)
}
