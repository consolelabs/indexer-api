package marketplacecollectionsnapshotsummarized

import (
	"github.com/consolelabs/indexer-api/pkg/config"
	"github.com/consolelabs/indexer-api/pkg/logger"
	"github.com/consolelabs/indexer-api/pkg/store"
)

func Run(args ...string) error {
	l := logger.L
	cfg := config.LoadConfig(config.DefaultConfigLoaders())

	store := store.New(cfg)

	platforms, err := store.Nft.GetAllMarketplacePlatform()
	if err != nil {
		l.Error(err, "failed to get all marketplace platform")
		return err
	}

	// snapshot for platform paintswap and nftkey first
	// TODO(trkhoi): add other platform snapshot
	for _, platform := range platforms {
		if platform.ID == 1 || platform.ID == 2 || platform.ID == 3 {
			err := store.Nft.SummarizeSnapshotCollection(int64(platform.ID))
			if err != nil {
				l.Fields(logger.Fields{"platformId": platform.ID, "platformName": platform.Name}).Error(err, "failed to summarize snapshot collection")
				return err
			}
		}
	}

	l.Info("Starting worker")
	defer l.Info("Finished worker")
	return nil
}
