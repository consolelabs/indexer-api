package chain_explorer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/consolelabs/indexer-api/pkg/config"
	"github.com/consolelabs/indexer-api/pkg/constant"
	"github.com/consolelabs/indexer-api/pkg/model"
)

const (
	moduleAccount = "account"
	actionTxlist  = "txlist"
	startBlock    = "0"
	endBlock      = "99999999"
)

type chainExplorerEntity struct {
	config *config.Config
}

func New(cfg *config.Config) IService {
	return &chainExplorerEntity{
		config: cfg,
	}
}

// TODO: multichain
func (e *chainExplorerEntity) GetCreationBlockNumber(contract model.Contract) (int64, error) {
	chainApiInfo := e.selectChain(contract.ChainId)
	uri := chainApiInfo.ScanApiUrl + "?module=%s&action=%s&address=%s&startblock=%s&endblock=%s&apikey=%s"

	client := &http.Client{
		Timeout: time.Second * 60,
	}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(uri, moduleAccount, actionTxlist, contract.Address, startBlock, endBlock, chainApiInfo.ScanApiKey), nil)
	if err != nil {
		return 0, err
	}

	q := req.URL.Query()
	req.URL.RawQuery = q.Encode()
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	type response struct {
		Result []model.ChainExplorerTransaction `json:"result"`
	}

	var r response
	err = json.Unmarshal(body, &r)
	if err != nil {
		return 0, err
	}

	creationBlock, err := strconv.Atoi(r.Result[0].BlockNumber)
	if err != nil {
		return 0, err
	}

	return int64(creationBlock), err
}

func (e *chainExplorerEntity) selectChain(chainId int64) model.Chain {
	switch chainId {
	case constant.C.ChainId.Eth:
		return model.Chain{
			ScanApiUrl: "https://api.etherscan.io/api",
			ScanApiKey: e.config.ChainExplorerApiKey.Eth,
		}
	case constant.C.ChainId.Bsc:
		return model.Chain{
			ScanApiUrl: "https://api.bscscan.com/api",
			ScanApiKey: e.config.ChainExplorerApiKey.Bsc,
		}
	case constant.C.ChainId.Ftm:
		return model.Chain{
			ScanApiUrl: "https://api.ftmscan.com/api",
			ScanApiKey: e.config.ChainExplorerApiKey.Ftm,
		}
	case constant.C.ChainId.Opt:
		return model.Chain{
			ScanApiUrl: "https://api-optimistic.etherscan.io/api",
			ScanApiKey: e.config.ChainExplorerApiKey.Optimism,
		}
	default:
		return model.Chain{
			ScanApiUrl: "https://api.etherscan.io/api",
			ScanApiKey: e.config.ChainExplorerApiKey.Eth,
		}
	}
}
