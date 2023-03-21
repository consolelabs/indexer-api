package holder_snapshot_summarized

import (
	"github.com/consolelabs/indexer-api/pkg/config"
	"github.com/consolelabs/indexer-api/pkg/logger"
	"github.com/consolelabs/indexer-api/pkg/store"
)

func Run(args ...string) error {
	l := logger.L
	cfg := config.LoadConfig(config.DefaultConfigLoaders())

	store := store.New(cfg)

	// snapshot for platform paintswap and nftkey first
	err := store.Nft.SummarizeSnapshotHolder()
	if err != nil {
		l.Fields(logger.Fields{"task": "snapshot holder by date"}).Error(err, "failed to summarize snapshot holder")
		return err
	}

	l.Info("Starting worker")
	defer l.Info("Finished worker")
	return nil
}
