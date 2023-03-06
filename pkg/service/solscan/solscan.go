package solscan

import (
	"encoding/json"
	"fmt"
	"github.com/Danny-Dasilva/CycleTLS/cycletls"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/gammazero/workerpool"

	"github.com/consolelabs/indexer-api/pkg/config"
	"github.com/consolelabs/indexer-api/pkg/logger"
	"github.com/consolelabs/indexer-api/pkg/model"
	"github.com/consolelabs/indexer-api/pkg/store"
)

type SolScan struct {
	store  *store.Store
	config *config.Config
	logger logger.Logger
}

func New(store *store.Store, config *config.Config) IService {
	return &SolScan{
		store:  store,
		config: config,
		logger: logger.L,
	}
}

func (s *SolScan) MapSolanaData() error {
	collections, err := s.store.Nft.GetSolScanCollections()
	if err != nil {
		s.logger.Fields(logger.Fields{
			"worker": "sync holder collection",
		}).Error(err, "failed query all nft_collection")
	}
	wp := workerpool.New(3)
	for _, ev := range collections {
		ev := ev
		wp.Submit(func() {
			s.logger.Fields(logger.Fields{
				"collectionId": ev.Address,
			}).Infof("map collection with holder data")
			address := ""
			if strings.Contains(ev.Address, "-") {
				address = strings.Split(ev.Address, "-")[1]
			}
			err := s.mapTokenWithHolder(address)
			if err != nil {
				s.logger.Fields(logger.Fields{
					"collectionId": ev.Address,
				}).Error(err, "failed to get data")
			}

		})
	}

	wp.StopWait()
	return nil
}

func (s *SolScan) mapTokenWithHolder(id string) error {
	holders, err := s.getAllHolder(id)
	if err != nil {
		s.logger.Fields(logger.Fields{
			"collection id": id,
		}).Error(err, "failed to get all holder")
		return err
	}
	wp1 := workerpool.New(3)

	for _, holder := range holders {
		holder := holder
		id := id
		wp1.Submit(func() {
			s.logger.Fields(logger.Fields{
				"collection": id,
				"holder":     holder,
			}).Infof("get nft of data")

			nftFromHolder, err := s.getAllNftFromHolder(holder)
			if err != nil {
				s.logger.Fields(logger.Fields{
					"collection": id,
					"holder":     holder,
				}).Error(err, "failed to get all nft from holder after retrying")
			}
			for _, nft := range nftFromHolder {
				if nft.NftCollectionID == id {
					tokenId := nft.TokenAddress
					if strings.Contains(nft.NftName, "#") {
						tokenId = strings.Split(nft.NftName, "#")[1]
					}
					owner := model.NftOwner{}
					owner.OwnerAddress = holder
					owner.CollectionAddress = fmt.Sprintf("solscan-%s", id)
					owner.TokenId = tokenId
					owner.ChainId = 999
					owner.CreatedTime = time.Now()
					owner.LastUpdatedTime = time.Now()
					s.store.Nft.UpsertOwner(&owner)
				}
			}
		})
	}
	wp1.StopWait()
	return nil
}

func (s *SolScan) getAllHolder(id string) ([]string, error) {
	page := 1
	var holders []string
	for {
		nftHolder, err := s.getCollectionHolder(id, page)
		if err != nil {
			return holders, err
		}
		if len(nftHolder.Data.Holders) == 0 {
			break
		}
		for _, v := range nftHolder.Data.Holders {
			holders = append(holders, v.WalletAddress)
		}
		page += 1
	}

	return holders, nil
}

func (s *SolScan) getAllNftFromHolder(id string) ([]model.NftSolScanData, error) {
	offset := 0
	limit := 100
	var nfts []model.NftSolScanData
	for {
		nft, err := s.getNftTokenFromHolder(id, offset, limit)
		if err != nil {
			return nfts, err
		}
		if len(nft.Data.ListNfts) == 0 {
			break
		}
		for _, v := range nft.Data.ListNfts {
			nfts = append(nfts, v)
		}
		offset += limit
	}

	return nfts, nil
}

func (s *SolScan) getCollectionHolder(id string, page int) (*model.NftHolder, error) {
	retry := 0
	var err error
	var nftToken *model.NftHolder
	for retry < 5 {
		nftToken, err = s.fetchCollectionHolder(id, page)
		s.logger.Fields(logger.Fields{
			"collectionId": id,
			"page":         page,
		}).Info("get holder")
		if err != nil {
			s.logger.Fields(logger.Fields{
				"collectionId": id,
				"page":         page,
				"retry":        retry,
			}).Error(err, "failed to get nft holder, retrying")
			retry++
			time.Sleep(7 * time.Duration(retry) * time.Second)
			continue
		}
		break
	}

	if retry == 5 {
		return nil, err
	}

	return nftToken, nil
}

func (s *SolScan) getNftTokenFromHolder(holderAddress string, offset, limit int) (*model.NftFromHolder, error) {
	retry := 0
	var err error
	var nftToken *model.NftFromHolder
	for retry < 5 {
		nftToken, err = s.fetchNftTokenFromHolder(holderAddress, offset, limit)
		if err != nil {
			s.logger.Fields(logger.Fields{
				"holder": holderAddress,
				"retry":  retry,
			}).Error(err, "failed to get nft token from holder, retrying")
			retry++
			time.Sleep(7 * time.Duration(retry) * time.Second)
			continue
		}
		break
	}

	if retry == 5 {
		return nil, err
	}

	return nftToken, nil
}

func (s *SolScan) fetchCollectionHolder(id string, page int) (*model.NftHolder, error) {
	var client = &http.Client{}
	request, err := http.NewRequest("GET", fmt.Sprintf("%s/nft/collection/holders/%s?page=%d", s.config.SolScan.BaseUrl, id, page), nil)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("token", s.config.SolScan.SolScanToken)

	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	resBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	res := &model.NftHolder{}
	err = json.Unmarshal(resBody, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *SolScan) fetchNftTokenFromHolder(holderAddress string, offset, limit int) (*model.NftFromHolder, error) {
	url := fmt.Sprintf("%s/wallet/nft/%s?offset=%d&limit=%d", s.config.SolScan.BaseUrl, holderAddress, offset, limit)
	//var client = &http.Client{}
	//request, err := http.NewRequest("GET", url , nil)
	//if err != nil {
	//	return nil, err
	//}
	//
	//request.Header.Add("Content-Type", "application/json")
	//
	//response, err := client.Do(request)

	//defer response.Body.Close()
	//resBody, err := ioutil.ReadAll(response.Body)
	//if err != nil {
	//	return nil, err
	//}

	response, err := s.requestByCyclets(url)

	if err != nil {
		return nil, err
	}

	res := &model.NftFromHolder{}
	err = json.Unmarshal([]byte(response.Body), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *SolScan) requestByCyclets(url string) (*cycletls.Response, error) {
	client := cycletls.Init()
	response, err := client.Do(url, cycletls.Options{
		Body:      "",
		Ja3:       "771,4865-4867-4866-49195-49199-52393-52392-49196-49200-49162-49161-49171-49172-51-57-47-53-10,0-23-65281-10-11-35-16-5-51-43-13-45-28-21,29-23-24-25-256-257,0",
		UserAgent: "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:87.0) Gecko/20100101 Firefox/87.0",
	}, "GET")
	if err != nil {
		return nil, err
	}
	return &response, nil
}
