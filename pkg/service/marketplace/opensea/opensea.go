package opensea

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/consolelabs/indexer-api/pkg/config"
	"github.com/consolelabs/indexer-api/pkg/model"
)

type opensea struct {
	config *config.Config
}

func New(cfg *config.Config) IService {
	return &opensea{
		config: cfg,
	}
}

var (
	openseaBaseURL = "https://api.opensea.io"
)

func (m *opensea) GetOpenseaCollection(collectionAddress string) (*model.OpenseaCollection, error) {
	// TODO: centralize
	url := fmt.Sprintf("%s/api/v1/asset_contract/%s", openseaBaseURL, collectionAddress)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("X-API-KEY", m.config.MarketplaceApiKey.Opensea)

	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	type respstruct struct {
		Collection *model.OpenseaCollection `json:"collection"`
	}

	resData := &respstruct{}
	err = json.Unmarshal(body, resData)
	if err != nil {
		return nil, err
	}

	return resData.Collection, nil
}
