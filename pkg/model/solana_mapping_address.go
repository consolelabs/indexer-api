package model

import "time"

type SolanaMappingAddress struct {
	Id             int64     `json:"id"`
	OnchainAddress string    `json:"onchain_address"`
	SolscanId      string    `json:"solscan_id"`
	Slug           string    `json:"slug"`
	CreatedTime    time.Time `json:"created_time"`
	UpdatedTime    time.Time `json:"updated_time"`
}
