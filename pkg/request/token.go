package request

type ConvertTokenPrice struct {
	Amount string `json:"amount"`
	From   string `json:"from"`
	To     string `json:"to"`
}
