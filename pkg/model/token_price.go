package model

type TokenPriceDetail struct {
	TokenId  int64  `json:"token_id"`
	Price    string `json:"price"`
	Symbol   string `json:"symbol"`
	Name     string `json:"name"`
	Decimals int64  `json:"decimals"`
}

type ConvertTokenPrice struct {
	From struct {
		Amount string `json:"amount"`
		Symbol string `json:"symbol"`
		Name   string `json:"name"`
	} `json:"from"`
	To struct {
		Amount string `json:"amount"`
		Symbol string `json:"symbol"`
		Name   string `json:"name"`
	} `json:"to"`
}
