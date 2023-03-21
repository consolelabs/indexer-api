package request

type ConvertTokenPrice struct {
	Amount    string `json:"amount"`
	FromToken string `json:"from_token"`
	ToToken   string `json:"to_token"`
}
