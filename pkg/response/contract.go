package response

import "time"

type ContractResponseData struct {
	ID               int64     `json:"id"`
	LastUpdatedTime  time.Time `json:"last_updated_time"`
	LastUpdatedBlock int64     `json:"last_updated_block"`
	CreationBlock    int64     `json:"creation_block"`
	CreatedTime      time.Time `json:"created_time"`
	Address          string    `json:"address"`
	ChainId          int64     `json:"chain_id"`
	Type             string
	IsProxy          *bool  `json:"is_proxy"`
	LogicAddress     string `json:"logic_address"`
	Protocol         string
	GrpcAddress      string
	IsSynced         bool `json:"is_synced"`
}

// for swagger
type ContractResponse struct {
	Data ContractResponseData `json:"data"`
}
