package coingeckohistoryprice

import (
	"fmt"
	"math"
	"math/big"
	"sync"
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

		coingeckoTokens := filterTokenWithCoingecko(listTokens)

		var wg sync.WaitGroup
		for i, t := range coingeckoTokens {
			wg.Add(1)
			l.Fields(logger.Fields{"token": t.Symbol}).Infof("Processing token %d/%d", i+1, len(coingeckoTokens))

			go executeIndexingCoingeckoHistoryPrice(&wg, cfg, store, service, t)
		}
		wg.Wait()
	}
}

func executeIndexingCoingeckoHistoryPrice(wg *sync.WaitGroup, cfg *config.Config, store *store.Store, service *service.Service, token model.Token) {
	defer wg.Done()
	latestSnapshot, err := store.Token.GetLatestSnapshot(token.Id, "coingecko")
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
		history, err := processGetTokenPrice(service, &token, date)
		if err != nil {
			logger.L.Fields(logger.Fields{"token": token.Symbol, "date": date, "coingecko": token.CoingeckoId}).Error(err, "[GetHistoryPrice] Fail to get history price from coingecko")
			return
		}

		price := big.NewFloat(math.Floor(history.MarketData.CurrentPrice.Usd * math.Pow(10, float64(token.Decimals))))

		stringPrice := fmt.Sprintf("%.0f", price)
		dateTimestamp, _ := time.Parse("02-01-2006", date)

		snapshot := model.TokenHistoryPriceSnapshot{
			TokenId: token.Id,
			Price:   stringPrice,
			Source:  "coingecko",
			Time:    dateTimestamp,
		}
		err = store.Token.Save(&snapshot)
		if err != nil {
			logger.L.Fields(logger.Fields{"token": token.Symbol, "date": date, "coingecko": token.CoingeckoId}).Error(err, "[Save] Fail to save snapshot to database")
			return
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

func filterTokenWithCoingecko(tokens []model.Token) (listTokens []model.Token) {
	for _, t := range tokens {
		if t.CoingeckoId != "" {
			listTokens = append(listTokens, t)
		}
	}
	return listTokens
}

func processGetTokenPrice(service *service.Service, token *model.Token, date string) (history *model.GetTokenHistoryPrice, err error) {
	retry := 0
	for retry < 5 {
		history, err = service.Coingecko.GetHistoryPrice(token.CoingeckoId, date)
		if err != nil {
			logger.L.Fields(logger.Fields{
				"token": token.Symbol,
				"retry": retry,
				"date":  date,
			}).Error(err, "failed to get token history price, retrying")
			retry++
			time.Sleep(60 * time.Second)
			continue
		}
		break
	}

	if retry == 5 {
		return nil, err
	}

	return history, nil
}
