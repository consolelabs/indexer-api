package model

import (
	"time"

	"gorm.io/gorm"
)

type Contract struct {
	ID               int64     `json:"id"`
	LastUpdatedTime  time.Time `json:"last_updated_time"`
	LastUpdatedBlock int64     `json:"last_updated_block"`
	CreationBlock    int64     `json:"creation_block"`
	CreatedTime      time.Time `json:"created_time"`

	Address string `json:"address"`
	ChainId int64  `json:"chain_id"`
	Type    string `json:"type"`

	IsProxy      *bool  `json:"is_proxy"`
	LogicAddress string `json:"logic_address"`

	Protocol    string `json:"protocol"`
	GrpcAddress string `json:"grpc_address"`
}

func (c *Contract) BeforeCreate(tx *gorm.DB) (err error) {
	c.CreatedTime = time.Now()
	c.LastUpdatedTime = time.Now()
	return
}
