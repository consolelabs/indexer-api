package coingecko

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/consolelabs/indexer-api/pkg/config"
	"github.com/consolelabs/indexer-api/pkg/model"
)

type Service struct {
	getHistoryPrice string
	getPriceURL     string
}

func New(cfg *config.Config) IService {
	apiKey := cfg.CoingeckoApiKey
	return &Service{
		getHistoryPrice: "https://pro-api.coingecko.com/api/v3/coins/%s/history?date=%s&localization=false&x_cg_pro_api_key=" + apiKey,
		getPriceURL:     "https://pro-api.coingecko.com/api/v3/simple/price?ids=%s&vs_currencies=%s&x_cg_pro_api_key=" + apiKey,
	}
}

func (s Service) GetHistoryPrice(coingeckoId, date string) (*model.GetTokenHistoryPrice, error) {
	var client = &http.Client{}
	request, err := http.NewRequest("GET", fmt.Sprintf(s.getHistoryPrice, coingeckoId, date), nil)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")

	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	resBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	res := &model.GetTokenHistoryPrice{}
	err = json.Unmarshal(resBody, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type CoinPriceResponse map[string]map[string]float64

func (s Service) GetCurrentPrice(coingeckoId string) (float64, error) {
	var client = &http.Client{}
	request, err := http.NewRequest("GET", fmt.Sprintf(s.getPriceURL, coingeckoId, "usd"), nil)
	if err != nil {
		return 0, err
	}

	request.Header.Add("Content-Type", "application/json")

	response, err := client.Do(request)

	if err != nil {
		return 0, err
	}

	defer response.Body.Close()
	resBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0, err
	}

	res := &CoinPriceResponse{}
	err = json.Unmarshal(resBody, res)
	if err != nil {
		return 0, err
	}

	prices := make(map[string]float64)
	for k, v := range *res {
		prices[k] = v["usd"]
	}

	return prices[coingeckoId], nil
}
