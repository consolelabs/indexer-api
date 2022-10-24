package model

type Token struct {
	Symbol   string `json:"symbol"`
	IsNative bool   `json:"is_native"`
	Address  string `json:"address"`
	Decimals int64  `json:"decimals"`
	IconUrl  string `json:"icon_url"`
}
