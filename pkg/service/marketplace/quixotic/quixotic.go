package quixotic

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/consolelabs/indexer-api/pkg/config"
	"github.com/consolelabs/indexer-api/pkg/model"
)

type quixotic struct {
	config *config.Config
}

func New(cfg *config.Config) IService {
	return &quixotic{
		config: cfg,
	}
}

var (
	quixoticBaseUrl = "https://api.quixotic.io"
)

func (e *quixotic) GetQuixoticCollection(collectionAddress string) (*model.QuixoticCollection, error) {
	url := fmt.Sprintf("%s/api/v1/collection/%s", quixoticBaseUrl, collectionAddress)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("X-API-KEY", e.config.MarketplaceApiKey.Quixotic)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	data := &model.QuixoticCollection{}
	return data, json.Unmarshal(body, data)
}
