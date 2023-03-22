package model

type Token struct {
	Id          int64   `json:"id"`
	Symbol      string  `json:"symbol"`
	IsNative    bool    `json:"is_native"`
	Address     string  `json:"address"`
	Decimals    int64   `json:"decimals"`
	IconUrl     *string `json:"icon_url"`
	CoingeckoId string  `json:"coingecko_id"`
}
