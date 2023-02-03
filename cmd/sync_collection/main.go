package main

import (
	"github.com/consolelabs/indexer-api/pkg/config"
	handler "github.com/consolelabs/indexer-api/pkg/handler"
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

	address := []string{"0x280ea08C9408F30aAE0B08fA5bbc9638Bb213E69"}
	chainId := []int64{66}

	h, err := handler.New(store, service, queue.New(cfg, &queue.Config{
		Producer: true,
	}))
	if err != nil {
		logger.L.Fatalf(err, "Can't init handlers")
	}
	for i := 0; i < len(address); i++ {
		h.AddErc721ContractScript(address[i], "", "", chainId[i], false)
	}

}
