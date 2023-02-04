package main

import (
	"encoding/json"
	"github.com/consolelabs/indexer-api/pkg/config"
	"github.com/consolelabs/indexer-api/pkg/handler"
	"github.com/consolelabs/indexer-api/pkg/logger"
	"github.com/consolelabs/indexer-api/pkg/queue"
	"github.com/consolelabs/indexer-api/pkg/service"
	"github.com/consolelabs/indexer-api/pkg/store"
	"io/ioutil"
	"os"
	"os/signal"
	"strconv"
)

type Data struct {
	Address []string `json:"address"`
	ChainId []string `json:"chain_id"`
}

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

	fileContent, err := os.Open("data.json")

	if err != nil {
		logger.L.Fatalf(err, "Can't open file")
		return
	}

	defer fileContent.Close()

	byteResult, _ := ioutil.ReadAll(fileContent)

	var payload Data
	err = json.Unmarshal(byteResult, &payload)
	if err != nil {
		log.Fatal(err, "Error during Unmarshal(): ")
	}

	for i := 0; i < len(payload.Address); i++ {
		chainId, _ := strconv.Atoi(payload.ChainId[i])
		address := payload.Address[i]
		h.AddErc721ContractScript(address, "", "", int64(chainId), false)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

}
