package model

import "time"

type TokenPrice struct {
	TokenId   int64     `json:"token_id"`
	Price     string    `json:"price"`
	Time      time.Time `json:"time"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ConvertTokenPrice struct {
	From struct {
		Amount string `json:"amount"`
		Symbol string `json:"symbol"`
		Name   string `json:"name"`
	} `json:"from"`
	To struct {
		Amount string `json:"amount"`
		Symbol string `json:"symbol"`
		Name   string `json:"name"`
	} `json:"to"`
}
