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
