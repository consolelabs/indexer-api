package main

import (
	"github.com/consolelabs/indexer-api/pkg/config"
	"github.com/consolelabs/indexer-api/pkg/entity"
	"github.com/consolelabs/indexer-api/pkg/logger"
	"github.com/consolelabs/indexer-api/pkg/service"
	"github.com/consolelabs/indexer-api/pkg/store"
)

func main() {
	cfg := config.LoadConfig(config.DefaultConfigLoaders())
	log := logger.NewLogrusLogger()

	log.Infof("Starting sync data nft collection")

	store := store.New(cfg)
	service := service.New(cfg)

	// sync created_time by collection_address
	e := entity.New(store, service)

	err := e.Nft.SyncDataNftCollection()
	if err != nil {
		log.Errorf(err, "UpdateCreatedTimeNftListing error")
	}
	log.Infof("Done update created time")
}
