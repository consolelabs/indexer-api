package model

type Price struct {
	Token  Token  `json:"token"`
	Amount string `json:"amount"`
	// TODO: get exchange rate and calculate
	AmountInUsd float64 `json:"amount_in_usd"`
}
