package utils

import "github.com/consolelabs/indexer-api/pkg/constant"

// TODO: move to db
func ConvertChainToMarketplace(chainId int64) string {
	switch chainId {
	case constant.C.ChainId.Eth:
		return "opensea"
	case constant.C.ChainId.Ftm:
		return "paintswap"
	case constant.C.ChainId.Opt:
		return "quixotic"
	default:
		return ""
	}
}
