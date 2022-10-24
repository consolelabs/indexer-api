package ethclient

import (
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/consolelabs/indexer-api/pkg/config"
	"github.com/consolelabs/indexer-api/pkg/constant"
	"github.com/consolelabs/indexer-api/pkg/logger"
)

type EthClient struct {
	Eth      *ethclient.Client
	Ftm      *ethclient.Client
	Optimism *ethclient.Client
}

func New(cfg *config.Config) *EthClient {
	ethClient, err := ethclient.Dial(cfg.RpcUrl.Eth)
	if err != nil {
		logger.L.Fatal(err, "Can't connect to Ethereum node")
	}

	ftmClient, err := ethclient.Dial(cfg.RpcUrl.Ftm)
	if err != nil {
		logger.L.Fatal(err, "Can't connect to Ftm node")
	}

	optClient, err := ethclient.Dial(cfg.RpcUrl.Optimism)
	if err != nil {
		logger.L.Fatal(err, "Can't connect to Optimism node")
	}

	return &EthClient{
		Eth:      ethClient,
		Ftm:      ftmClient,
		Optimism: optClient,
	}
}

func (e *EthClient) GetByChainId(chainId int64) *ethclient.Client {
	if chainId == constant.C.ChainId.Eth {
		return e.Eth
	}
	if chainId == constant.C.ChainId.Ftm {
		return e.Ftm
	}
	return e.Optimism
}
