package model

type BirdeyeHistoryPrice struct {
	Data struct {
		Items []struct {
			Address  string  `json:"address"`
			UnixTime int     `json:"unixTime"`
			Value    float64 `json:"value"`
		} `json:"items"`
	} `json:"data"`
	Success bool `json:"success"`
}

type BirdeyeCurrentPriceData struct {
	Data    BirdeyeCurrentPrice `json:"data"`
	Success bool                `json:"success"`
}

type BirdeyeCurrentPrice struct {
	Value           float64 `json:"value"`
	UpdateUnixTime  int64   `json:"updateUnixTime"`
	UpdateHumanTime string  `json:"updateHumanTime"`
}
