package request

import "github.com/consolelabs/indexer-api/pkg/model"

type GetChainsQuery struct {
	*model.Pagination
	// name,symbol,chain_id,is_evm
	Sort    string  `json:"sort" form:"sort" binding:"omitempty,oneof=name -name symbol -symbol chain_id -chain_id is_evm -is_evm" enums:"name,-name,symbol,-symbol,chain_id,-chain_id,is_evm,-is_evm"`
	Name    *string `json:"name" form:"name"`         // search chain name
	Symbol  *string `json:"symbol" form:"symbol"`     // chain symbol
	ChainId *uint64 `json:"chain_id" form:"chain_id"` // chain ID
	IsEvm   *bool   `json:"is_evm" form:"is_evm"`     // is EVM chain
}
