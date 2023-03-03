package sync_solana

import (
	"github.com/consolelabs/indexer-api/pkg/config"
	"github.com/consolelabs/indexer-api/pkg/logger"
	"github.com/consolelabs/indexer-api/pkg/service"
	"github.com/consolelabs/indexer-api/pkg/store"
)

func Run(args ...string) error {
	cfg := config.LoadConfig(config.DefaultConfigLoaders())
	log := logger.NewLogrusLogger()

	log.Info("Starting worker")
	defer log.Info("Finished worker")

	service := service.New(cfg)
	store := store.New(cfg)
	h, err := New(store, service, cfg)
	if err != nil {
		logger.L.Fatalf(err, "Can't init sync holder")
	}
	err = h.MapSolanaData()
	if err != nil {
		logger.L.Fatalf(err, "Can't init sync holder")
	}
	return nil
}
