package model

type Chain struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	Symbol     string `json:"symbol"`
	ChainId    int64  `json:"chain_id"`
	ScanUrl    string `json:"scan_url"`
	ScanApiUrl string `json:"-" gorm:"-"`
	ScanApiKey string `json:"-" gorm:"-"`
	IsEvm      bool   `json:"is_evm"`
}
