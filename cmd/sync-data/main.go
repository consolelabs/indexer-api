package main

import (
	"github.com/consolelabs/indexer-api/pkg/config"
	"github.com/consolelabs/indexer-api/pkg/handler"
	"github.com/consolelabs/indexer-api/pkg/logger"
	"github.com/consolelabs/indexer-api/pkg/queue"
	"github.com/consolelabs/indexer-api/pkg/service"
	"github.com/consolelabs/indexer-api/pkg/store"
)

func main() {
	cfg := config.LoadConfig(config.DefaultConfigLoaders())
	log := logger.NewLogrusLogger()

	log.Infof("Server starting")

	service := service.New(cfg)
	store := store.New(cfg)
	h, err := handler.New(store, service, queue.New(cfg, &queue.Config{
		Producer: true,
	}))
	if err != nil {
		logger.L.Fatalf(err, "Can't init handlers")
	}
	err = h.SyncMagicedenData()
	if err != nil {
		logger.L.Fatalf(err, "Can't init handlers")
	}
}
