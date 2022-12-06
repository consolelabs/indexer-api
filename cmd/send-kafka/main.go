package main

import (
	"github.com/consolelabs/indexer-api/pkg/config"
	"github.com/consolelabs/indexer-api/pkg/entity"
	"github.com/consolelabs/indexer-api/pkg/logger"
	"github.com/consolelabs/indexer-api/pkg/queue"
	"github.com/consolelabs/indexer-api/pkg/service"
	"github.com/consolelabs/indexer-api/pkg/store"
)

func main() {
	cfg := config.LoadConfig(config.DefaultConfigLoaders())
	log := logger.NewLogrusLogger()

	log.Infof("Server starting")

	store := store.New(cfg)
	service := service.New(cfg)

	// sync created_time by collection_address
	e := entity.New(store, service)
	queue := queue.New(cfg, &queue.Config{
		Producer: true,
	})

	err := e.Nft.SendKafkaMessage(cfg, queue)
	if err != nil {
		log.Errorf(err, "SendKafkaMessage error")
	}
	log.Infof("Done send kafka message")

}
