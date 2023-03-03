package main

import (
	"github.com/consolelabs/indexer-api/pkg/logger"
	"github.com/consolelabs/indexer-api/pkg/worker"
	syncsolana "github.com/consolelabs/indexer-api/pkg/worker/sync-solana"
)

func main() {
	log := logger.NewLogrusLogger()
	worker.Init(log)

	worker.AddCommand(worker.NewCmd(log).SetName("sync-holder-collection").SetAliases("sync-holder").SetUsage("Collect holder for solana collection").SetRun(syncsolana.Run))

	if err := worker.Execute(); err != nil {
		log.Fatal(err, "failed to execute")
	}
}
