package main

import (
	"github.com/consolelabs/indexer-api/pkg/logger"
	"github.com/consolelabs/indexer-api/pkg/worker"
	birdeyehistoryprice "github.com/consolelabs/indexer-api/pkg/worker/birdeye-history-price"
	coingeckohistoryprice "github.com/consolelabs/indexer-api/pkg/worker/coingecko-history-price"
	snapshotholder "github.com/consolelabs/indexer-api/pkg/worker/holder-snapshot-summarized"
	snapshotsummarized "github.com/consolelabs/indexer-api/pkg/worker/marketplace-collection-snapshot-summarized"
)

func main() {
	log := logger.NewLogrusLogger()
	worker.Init(log)

	worker.AddCommand(worker.NewCmd(log).SetName("marketplace-collection-snapshot-summarized").SetAliases("snapsum").SetUsage("Summarize snapshot for collection").SetRun(snapshotsummarized.Run))
	worker.AddCommand(worker.NewCmd(log).SetName("holder-snapshot-summarized").SetAliases("snapholdersum").SetUsage("Summarize snapshot for holder").SetRun(snapshotholder.Run))
	worker.AddCommand(worker.NewCmd(log).SetName("coingecko-history-price").SetAliases("coingeckohistoryprice").SetUsage("Crawl history price from Coingecko").SetRun(coingeckohistoryprice.Run))
	worker.AddCommand(worker.NewCmd(log).SetName("birdeye-history-price").SetAliases("birdeyehistoryprice").SetUsage("Crawl history price from Birdeye").SetRun(birdeyehistoryprice.Run))
	if err := worker.Execute(); err != nil {
		log.Fatal(err, "failed to execute")
	}
}
