package contract

import (
	"github.com/consolelabs/indexer-api/pkg/model"
)

type IContractEntity interface {
	AddContract(contract model.Contract, name string, symbol string) error
	GetContract(address string) (*model.Contract, error)
}
