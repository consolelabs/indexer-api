package birdeye_history_price

import (
	"fmt"
	"math"
	"time"

	"github.com/consolelabs/indexer-api/pkg/config"
	"github.com/consolelabs/indexer-api/pkg/logger"
	"github.com/consolelabs/indexer-api/pkg/model"
	"github.com/consolelabs/indexer-api/pkg/service"
	"github.com/consolelabs/indexer-api/pkg/store"
)

func Run(args ...string) error {
	l := logger.L
	cfg := config.LoadConfig(config.DefaultConfigLoaders())

	store := store.New(cfg)
	service := service.New(cfg)

	l.Info("Starting worker")

	for {
		listTokens, err := store.Token.GetListToken()
		if err != nil {
			l.Error(err, "[GetListToken] Fail to get list token from database")
		}

		for i, t := range listTokens {
			t := t
			l.Fields(logger.Fields{"token": t.Symbol}).Infof("Processing token %d/%d", i+1, len(listTokens))

			executeIndexingBirdeyeHistoryPrice(store, service, t)
		}
	}
}

func executeIndexingBirdeyeHistoryPrice(store *store.Store, service *service.Service, token model.Token) {
	latestSnapshot, err := store.Token.GetLatestSnapshot(token.Id, "birdeye")
	if err != nil {
		logger.L.Error(err, "[GetLatestSnapshot] Fail to get latest snapshot from database")
		return
	}

	if len(latestSnapshot) > 0 {
		if latestSnapshot[0].Time.Day() == time.Now().Day() && latestSnapshot[0].Time.Month() == time.Now().Month() && latestSnapshot[0].Time.Year() == time.Now().Year() {
			logger.L.Fields(logger.Fields{"token": token.Symbol}).Info("Token already indexed today")
			// sleep 30 minute and then recheck to avoid spam server
			time.Sleep(30 * time.Minute)
			return
		}
	}

	// TODO: make this can be config
	startDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	now := time.Now()
	listDate := genrateListDate(startDate, now)
	for _, date := range listDate {
		if token.Address != "0x0000000000000000000000000000000000000000" {
			history, err := service.Birdeye.GetTokenPriceByAddress(&token, date)
			if err != nil {
				logger.L.Fields(logger.Fields{"token": token.Symbol, "date": date, "source": "birdeye"}).Error(err, "[GetHistoryPrice] Fail to get history price from birdeye")
				return
			}
			if history != nil {
				if history.Success == true && len(history.Data.Items) > 0 {
					price := math.Floor(history.Data.Items[0].Value * math.Pow(10, float64(token.Decimals)))
					stringPrice := fmt.Sprintf("%d", int64(price))
					dateTimestamp, _ := time.Parse("02-01-2006", date)

					snapshot := model.TokenHistoryPriceSnapshot{
						TokenId: token.Id,
						Price:   stringPrice,
						Source:  "birdeye",
						Time:    dateTimestamp,
					}

					err = store.Token.Save(&snapshot)
					if err != nil {
						logger.L.Fields(logger.Fields{"token": token.Symbol, "date": date, "source": "birdeye"}).Error(err, "[Save] Fail to save snapshot to database")
						return
					}

					logger.L.Fields(logger.Fields{"token": token.Symbol, "date": date, "source": "birdeye"}).Info("save success")
				}
			}
		}
	}

}

func genrateListDate(startDate, endDate time.Time) (listDate []string) {
	for {
		if startDate.After(endDate) {
			break
		}

		listDate = append(listDate, startDate.Format("02-01-2006"))
		startDate = startDate.AddDate(0, 0, 1)
	}
	return listDate
}
