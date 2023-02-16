package contract

import (
	"errors"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/consolelabs/indexer-api/pkg/model"
	"github.com/consolelabs/indexer-api/pkg/service"
	"github.com/consolelabs/indexer-api/pkg/store"
)

type contractEntity struct {
	store   *store.Store
	service *service.Service
}

func New(store *store.Store, service *service.Service) IContractEntity {
	return &contractEntity{
		store:   store,
		service: service,
	}
}

func (e *contractEntity) AddContract(contract model.Contract, name string, symbol string) error {
	oldContract, _ := e.store.Contract.GetByAddress(contract.Address)
	if oldContract != nil {
		return errors.New("address has been added")
	}

	var (
		creationBlockNumber int64
		err                 error
	)

	if contract.ChainId != 999 && contract.ChainId != 66 && contract.ChainId != 9999 && contract.ChainId != 9998 && contract.ChainId != 9997 {
		// fetch creation block from chain explorer (etherscan, ftmscan, bscscan, ...)
		creationBlockNumber, err = e.service.ChainExplorer.GetCreationBlockNumber(contract)
		if err != nil {
			return err
		}

		// enrich data
		contract.CreationBlock = creationBlockNumber
	}

	ercFormat := "ERC_721"
	if contract.ChainId == 999 || contract.ChainId == 9999 || contract.ChainId == 9998 {
		ercFormat = ""
	}
	// store nft collection
	if contract.ChainId != 9999 {
		err = e.store.Nft.SaveNftCollection(&model.NftCollection{
			Address:         contract.Address,
			ChainId:         int64(contract.ChainId),
			CreatedTime:     time.Now(),
			LastUpdatedTime: time.Now(),
			ERCFormat:       ercFormat,
			Name:            name,
			Symbol:          symbol,
		})
		if err != nil {
			return err
		}
	}

	return e.store.Contract.Save(&contract)
}

func isCreationBlockSynced(creationBlockNumber int64, blockStats *model.BlockStats, chainId int64) bool {
	return (creationBlockNumber > blockStats.FirstBlock && creationBlockNumber < blockStats.LastBlock)
}

func (e *contractEntity) GetContract(address string) (*model.Contract, error) {
	contract, err := e.store.Contract.GetByAddress(address)
	if err != nil {
		log.Errorf("[entity.Contract] Fail to get contract %s", address)
		return nil, err
	}
	return contract, nil
}
