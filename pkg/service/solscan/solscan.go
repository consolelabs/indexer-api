package solscan

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/consolelabs/indexer-api/pkg/config"
	"github.com/consolelabs/indexer-api/pkg/model"
)

type solscan struct {
	config *config.Config
}

func New(cfg *config.Config) IService {
	return &solscan{
		config: cfg,
	}
}

var solscanBaseURL = "https://api.solscan.io/collection/id"

func (s *solscan) GetSolanaCollection(collectionId string) (*model.SolanaCollectionMetadata, error) {
	var client = &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s?collectionId=%s", solscanBaseURL, collectionId), nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	resBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	res := &model.SolanaCollectionMetadata{}
	err = json.Unmarshal(resBody, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
