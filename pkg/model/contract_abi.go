package model

type ContractAbi struct {
	ID int64 `json:"id"`

	Abi     string `json:"abi"`
	Address string `json:"address"`
}
