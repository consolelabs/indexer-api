package paintswap

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/consolelabs/indexer-api/pkg/config"
	"github.com/consolelabs/indexer-api/pkg/model"
)

type paintswap struct {
	config *config.Config
}

func New(cfg *config.Config) IService {
	return &paintswap{
		config: cfg,
	}
}

var (
	paintswapBaseURL = "https://api.paintswap.finance"
)

func (m *paintswap) GetPaintswapCollection(collectionAddress string) (*model.PaintswapCollection, error) {
	// TODO: centralize
	url := fmt.Sprintf("%s/v2/collections/%s", paintswapBaseURL, collectionAddress)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

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
		Collection *model.PaintswapCollection `json:"collection"`
	}

	data := &respstruct{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}
	return data.Collection, nil
}
