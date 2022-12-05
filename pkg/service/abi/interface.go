package abi

type IService interface {
	CheckERC721(address string, chainId int64) (bool, error)
	GetNameAndSymbol(address string, chainId int64) (name string, symbol string, err error)
}
