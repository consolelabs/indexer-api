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
		blockStats          *model.BlockStats
		err                 error
	)

	// handle for case evm, case solana assign chainId 0 for now
	// TODO(trkhoi): get creation block number for solana contract
	if contract.ChainId != 0 {
		// fetch creation block from chain explorer (etherscan, ftmscan, bscscan, ...)
		creationBlockNumber, err = e.service.ChainExplorer.GetCreationBlockNumber(contract)
		if err != nil {
			return err
		}

		// for case evm, check if creation block is synced
		// chainId = 99999999 means Aptos, no need
		if contract.ChainId != 99999999 {
			// fetch current synced block from clickhouse
			blockStats, err = e.service.ChainData.GetBlockStatsByChainId(contract.ChainId)
			if err != nil {
				return err
			}

			// check if we already synced all block in clickhouse
			if !isCreationBlockSynced(creationBlockNumber, blockStats, contract.ChainId) {
				return errors.New("block number not synced yet, TODO: add to queue and try later")
			}
		}

		// enrich data
		contract.CreationBlock = creationBlockNumber
	}

	ercFormat := "ERC_721"
	if contract.ChainId == 0 || contract.ChainId == 9999 || contract.ChainId == 9998 {
		ercFormat = ""
	}
	// store nft collection
	err = e.store.Nft.SaveNftCollection(&model.NftCollection{
		Address:         contract.Address,
		ChainId:         int64(contract.ChainId),
		CreatedTime:     time.Now(),
		LastUpdatedTime: time.Now(),
		Name:            name,
		Symbol:          symbol,
		ERCFormat:       ercFormat,
	})
	if err != nil {
		return err
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
