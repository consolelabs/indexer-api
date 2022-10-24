package abi

import (
	"context"

	"github.com/consolelabs/indexer-api/pkg/config"
	"github.com/consolelabs/indexer-api/pkg/constant"
	abi "github.com/consolelabs/indexer-api/pkg/grpc/abi_parser/erc721/gen"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type abiEntity struct {
	config *config.Config
}

func New(cfg *config.Config) IService {
	return &abiEntity{
		config: cfg,
	}
}
func (e *abiEntity) CheckERC721(address string, chainId int64) (bool, error) {
	rpcUrl := e.selectRpcUrl(chainId)
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		return false, err
	}

	addressHash := common.HexToAddress(address)
	instance, err := abi.NewERC721(addressHash, client)
	if err != nil {
		return false, err
	}

	isERC721, err := instance.SupportsInterface(&bind.CallOpts{
		Context: context.Background(),
	}, constant.C.ERC721Topic["Transfer"].InterfaceId)
	if err != nil {
		return false, err
	}

	return isERC721, nil
}

func (e *abiEntity) GetNameAndSymbol(address string, chainId int64) (name string, symbol string, err error) {
	rpcUrl := e.selectRpcUrl(chainId)
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		return "", "", err
	}

	addressHash := common.HexToAddress(address)
	instance, err := abi.NewERC721(addressHash, client)
	if err != nil {
		return "", "", err
	}
	name, err = instance.Name(&bind.CallOpts{})
	if err != nil {
		return "", "", err
	}
	symbol, err = instance.Symbol(&bind.CallOpts{})
	if err != nil {
		return "", "", err
	}
	return name, symbol, nil
}

func (e *abiEntity) selectRpcUrl(chainId int64) string {
	switch chainId {
	case 1:
		return e.config.RpcUrl.Eth
	case 250:
		return e.config.RpcUrl.Ftm
	default:
		return e.config.RpcUrl.Eth
	}
}
