package birdeye

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/consolelabs/indexer-api/pkg/config"
	"github.com/consolelabs/indexer-api/pkg/logger"
	"github.com/consolelabs/indexer-api/pkg/model"
	"github.com/consolelabs/indexer-api/pkg/store"
)

type Birdeye struct {
	store  *store.Store
	config *config.Config
	logger logger.Logger
}

func New(store *store.Store, config *config.Config) IService {
	return &Birdeye{
		store:  store,
		config: config,
		logger: logger.L,
	}
}

func (b *Birdeye) GetTokenPriceByAddress(token *model.Token, date string) (*model.BirdeyeHistoryPrice, error) {
	retry := 0
	var err error
	var tokenPriceHistory *model.BirdeyeHistoryPrice
	for retry < 3 {
		tokenPriceHistory, err = b.fetchTokenPriceByAddress(token, date)
		if err != nil {
			b.logger.Fields(logger.Fields{
				"token": fmt.Sprintf("Token %s-%s", token.Symbol, date),
				"retry": retry,
			}).Error(err, "failed to get token price, retrying")
			retry++
			time.Sleep(60 * time.Second)
			continue
		}
		break
	}

	if retry == 3 {
		return nil, err
	}

	return tokenPriceHistory, nil
}

func (b *Birdeye) fetchTokenPriceByAddress(token *model.Token, date string) (*model.BirdeyeHistoryPrice, error) {
	var client = &http.Client{}

	layout := "02-01-2006"
	fromTime, err := time.Parse(layout, date)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/public/history_price?address=%s&address_type=token&time_from=%d&time_to=%d", b.config.Birdeye.Api, token.Address, fromTime.Unix(), fromTime.Unix())
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")

	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	if response.Status == "429 Too Many Requests" {
		return nil, errors.New("rate limit")
	}

	defer response.Body.Close()
	resBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	res := &model.BirdeyeHistoryPrice{}
	err = json.Unmarshal(resBody, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
