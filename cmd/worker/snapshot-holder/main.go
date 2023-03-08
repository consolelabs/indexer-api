package main

import (
	"github.com/consolelabs/indexer-api/pkg/logger"
	"github.com/consolelabs/indexer-api/pkg/worker"
	snapshotholder "github.com/consolelabs/indexer-api/pkg/worker/holder-snapshot-summarized"
)

func main() {
	log := logger.NewLogrusLogger()
	worker.Init(log)

	worker.AddCommand(worker.NewCmd(log).SetName("holder-snapshot-summarized").SetAliases("snapholdersum").SetUsage("Summarize snapshot for holder").SetRun(snapshotholder.Run))

	if err := worker.Execute(); err != nil {
		log.Fatal(err, "failed to execute")
	}
}
