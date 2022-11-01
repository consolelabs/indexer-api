package abi

type IService interface {
	CheckERC721(address string, chainId int64) (bool, error)
}
