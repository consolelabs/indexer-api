package main

import (
	"github.com/consolelabs/indexer-api/pkg/logger"
	"github.com/consolelabs/indexer-api/pkg/worker"
	snapshotsummarized "github.com/consolelabs/indexer-api/pkg/worker/marketplace-collection-snapshot-summarized"
)

func main() {
	log := logger.NewLogrusLogger()
	worker.Init(log)

	worker.AddCommand(worker.NewCmd(log).SetName("marketplace-collection-snapshot-summarized").SetAliases("snapsum").SetUsage("Summarize snapshot for collection").SetRun(snapshotsummarized.Run))

	if err := worker.Execute(); err != nil {
		log.Fatal(err, "failed to execute")
	}
}
