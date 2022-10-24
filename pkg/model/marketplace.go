package model

import "math/big"

type PaintswapCollection struct {
	Id          string `json:"id"`
	Address     string `json:"address"`
	Name        string `json:"name"`
	Poster      string `json:"poster"`
	Description string `json:"description"`
	Discord     string `json:"discord"`
	Twitter     string `json:"twitter"`
	Website     string `json:"website"`
}

type OpenseaCollection struct {
	Name            string `json:"name"`
	ImageUrl        string `json:"image_url"`
	Description     string `json:"description"`
	DiscordUrl      string `json:"discord_url"`
	TwitterUsername string `json:"twitter_username"`
	ExternalUrl     string `json:"external_url"`
}

type QuixoticCollection struct {
	Address      string `json:"address"`
	Name         string `json:"name"`
	Symbol       string `json:"symbol"`
	ImageUrl     string `json:"image_url"`
	Description  string `json:"description"`
	ExternalLink string `json:"external_link"`
}

type OpenseaChainData struct {
	CollectionAddress string   `json:"collection_address"`
	TokenId           string   `json:"token_id"`
	Transaction       string   `json:"transaction"`
	Seller            string   `json:"seller"`
	Buyer             string   `json:"buyer"`
	Price             *big.Int `json:"price"`
}

type NFTKeyCollection struct {
	Id          string `json:"id"`
	Address     string `json:"address"`
	Name        string `json:"name"`
	Poster      string `json:"poster"`
	Description string `json:"description"`
	Discord     string `json:"discord"`
	Twitter     string `json:"twitter"`
	Website     string `json:"website"`
}