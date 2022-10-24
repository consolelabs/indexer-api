package chaindata

import "github.com/consolelabs/indexer-api/pkg/model"

type IService interface {
	GetTipOfChain() (*model.TipOfChain, error)
	GetEventsByAddressFromStartBlockToBlock(address string, start int64, to int64, chainId int64, out chan *model.Event) (err error)
	GetBlockStatsByChainId(chainId int64) (*model.BlockStats, error)
}
