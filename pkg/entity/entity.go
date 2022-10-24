package entity

import (
	chainentity "github.com/consolelabs/indexer-api/pkg/entity/chain_entity"
	contract "github.com/consolelabs/indexer-api/pkg/entity/contract"
	nftentity "github.com/consolelabs/indexer-api/pkg/entity/nft_entity"
	nftsaleentity "github.com/consolelabs/indexer-api/pkg/entity/nft_sale_entity"
	"github.com/consolelabs/indexer-api/pkg/service"
	"github.com/consolelabs/indexer-api/pkg/store"
)

type Entity struct {
	Nft      nftentity.INftEntity
	Contract contract.IContractEntity
	NftSale  nftsaleentity.INftSaleEntity
	Chain    chainentity.IChainEntity
}

func New(store *store.Store, service *service.Service) *Entity {
	return &Entity{
		Nft:      nftentity.New(store, service),
		Contract: contract.New(store, service),
		Chain:    chainentity.New(store, service),
		NftSale:  nftsaleentity.New(store, service),
	}
}
