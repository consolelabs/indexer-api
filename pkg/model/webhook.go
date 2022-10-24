package model

type Webhook struct {
	Event             string `json:"event"`
	CollectionAddress string `json:"collection_address"`
	TokenId           string `json:"token_id"`
	Marketplace       string `json:"marketplace"`
	Transaction       string `json:"transaction"`
	From              string `json:"from"`
	To                string `json:"to"`
	Price             Price  `json:"price"`
	LastPrice         Price  `json:"last_price"`
	Hold              int64  `json:"hold"`
	ChainId           int64  `json:"chain_id"`
}

type NftSalesResponse struct {
	Message string `json:"message"`
}
