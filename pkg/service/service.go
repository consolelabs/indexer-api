package service

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"

	"github.com/consolelabs/indexer-api/pkg/config"
	"github.com/consolelabs/indexer-api/pkg/logger"
	abi "github.com/consolelabs/indexer-api/pkg/service/abi"
	chaindata "github.com/consolelabs/indexer-api/pkg/service/chain_data"
	chainexplorer "github.com/consolelabs/indexer-api/pkg/service/chain_explorer"
	"github.com/consolelabs/indexer-api/pkg/service/gcs"
	opensea "github.com/consolelabs/indexer-api/pkg/service/marketplace/opensea"
	paintswap "github.com/consolelabs/indexer-api/pkg/service/marketplace/paintswap"
	quixotic "github.com/consolelabs/indexer-api/pkg/service/marketplace/quixotic"
)

type Service struct {
	ChainData     chaindata.IService
	ChainExplorer chainexplorer.IService
	Abi           abi.IService
	Opensea       opensea.IService
	Paintswap     paintswap.IService
	Quixotic      quixotic.IService
	Gcs           gcs.IService
}

func New(cfg *config.Config) *Service {
	db := connClickhouse(cfg)
	return &Service{
		ChainData:     chaindata.New(db),
		ChainExplorer: chainexplorer.New(cfg),
		Abi:           abi.New(cfg),
		Opensea:       opensea.New(cfg),
		Paintswap:     paintswap.New(cfg),
		Quixotic:      quixotic.New(cfg),
		Gcs:           gcs.New(cfg),
	}
}

// connClickhouse connect to clickhouse - where we store blockchain data, it should be private for this package only
func connClickhouse(cfg *config.Config) *sql.DB {
	defer logger.L.Info("clickhouse connected")

	conn := clickhouse.OpenDB(&clickhouse.Options{
		Addr: []string{fmt.Sprintf("%s:%s", cfg.Clickhouse.Host, cfg.Clickhouse.Port)},
		Auth: clickhouse.Auth{
			Database: cfg.Clickhouse.Name,
			Username: cfg.Clickhouse.User,
			Password: cfg.Clickhouse.Pass,
		},
		DialTimeout: 20 * time.Second,
		// Debug:       cfg.Debug,
	})
	conn.SetMaxIdleConns(5)
	conn.SetMaxOpenConns(10)
	conn.SetConnMaxLifetime(time.Hour)
	return conn
}
