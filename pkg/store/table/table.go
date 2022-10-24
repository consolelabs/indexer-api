package table

import (
	nfttablestore "github.com/consolelabs/indexer-api/pkg/store/table/nft"
	"gorm.io/gorm"
)

type Table struct {
	db  *gorm.DB
	Nft nfttablestore.ITableNft
}

func New(db *gorm.DB) *Table {
	return &Table{
		db:  db,
		Nft: nfttablestore.New(db),
	}
}
