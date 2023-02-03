package main

import (
	"github.com/consolelabs/indexer-api/pkg/config"
	"github.com/consolelabs/indexer-api/pkg/handler"
	"github.com/consolelabs/indexer-api/pkg/logger"
	"github.com/consolelabs/indexer-api/pkg/queue"
	"github.com/consolelabs/indexer-api/pkg/service"
	"github.com/consolelabs/indexer-api/pkg/store"
	"os"
	"os/signal"
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

	address := []string{"0x5014d8C095A6861eb4962a87bdD38476d475Df0D", "0x51909465D96da5A81917E8c0629EeDeD8A0CbEB1"}
	chainIds := []int64{66, 66}
	for i := 0; i < len(address); i++ {
		h.AddErc721ContractScript(address[i], "", "", chainIds[i], false)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

}
