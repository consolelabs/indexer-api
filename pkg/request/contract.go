package request

type AddContractRequest struct {
	Address      string `json:"address"`
	ChainId      int64  `json:"chain_id" binding:"omitempty,gte=0" msg:"invalid chain id"`
	Name         string `json:"name"`
	Symbol       string `json:"symbol"`
	PriorityFlag bool   `json:"priority_flag"`
}
