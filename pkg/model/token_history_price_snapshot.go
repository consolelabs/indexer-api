package model

import "time"

type TokenHistoryPriceSnapshot struct {
	Id        int64     `json:"id"`
	TokenId   int64     `json:"token_id"`
	Price     string    `json:"price"`
	Source    string    `json:"source"`
	Time      time.Time `json:"time"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
