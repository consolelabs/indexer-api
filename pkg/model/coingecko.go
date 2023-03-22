package model

type GetTokenHistoryPrice struct {
	Id         string      `json:"id"`
	Symbol     string      `json:"symbol"`
	Name       string      `json:"name"`
	Image      interface{} `json:"image"`
	MarketData MarketData  `json:"market_data"`
}

type MarketData struct {
	CurrentPrice CurrentPrice `json:"current_price"`
}

type CurrentPrice struct {
	Usd float64 `json:"usd"`
}
