package main

import (
	"github.com/consolelabs/indexer-api/pkg/config"
	"github.com/consolelabs/indexer-api/pkg/logger"
	"github.com/consolelabs/indexer-api/pkg/service"
)

func main() {
	cfg := config.LoadConfig(config.DefaultConfigLoaders())
	log := logger.NewLogrusLogger()

	log.Info("Starting worker")
	defer log.Info("Finished worker")

	service := service.New(cfg)

	err := service.SolScan.MapSolanaData()
	if err != nil {
		logger.L.Fatalf(err, "Can't init sync holder")
	}
}
