package store

import (
	"database/sql"
	"fmt"

	"github.com/consolelabs/indexer-api/pkg/store/token"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/consolelabs/indexer-api/pkg/config"
	"github.com/consolelabs/indexer-api/pkg/logger"
	contractstore "github.com/consolelabs/indexer-api/pkg/store/contract_store"
	nftstore "github.com/consolelabs/indexer-api/pkg/store/nft"
	table "github.com/consolelabs/indexer-api/pkg/store/table"
)

type Store struct {
	DB       *gorm.DB
	Contract contractstore.IContract
	Nft      nftstore.INft
	Table    table.Table
	Token    token.IToken
}

func New(cfg *config.Config) *Store {
	db := connDb(cfg)
	return &Store{
		DB:       db,
		Contract: contractstore.New(db),
		Nft:      nftstore.New(db),
		Table:    *table.New(db),
		Token:    token.New(db),
	}
}

func connDb(cfg *config.Config) *gorm.DB {
	ds := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.Postgres.User, cfg.Postgres.Pass,
		cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.Name,
	)

	conn, err := sql.Open("postgres", ds)
	if err != nil {
		logger.L.Fatalf(err, "failed to open database connection")
	}

	db, err := gorm.Open(postgres.New(
		postgres.Config{Conn: conn}),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		})
	if err != nil {
		logger.L.Fatalf(err, "failed to open database connection")
	}

	logger.L.Info("database connected")

	return db
}
