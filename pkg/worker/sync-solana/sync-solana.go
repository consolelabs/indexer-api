package sync_solana

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/gammazero/workerpool"

	"github.com/consolelabs/indexer-api/pkg/config"
	"github.com/consolelabs/indexer-api/pkg/entity"
	"github.com/consolelabs/indexer-api/pkg/logger"
	"github.com/consolelabs/indexer-api/pkg/model"
	"github.com/consolelabs/indexer-api/pkg/service"
	"github.com/consolelabs/indexer-api/pkg/store"
)

type SyncSolana struct {
	store    *store.Store
	config   *config.Config
	service  *service.Service
	logger   logger.Logger
	entities *entity.Entity
}

func New(store *store.Store, service *service.Service, config *config.Config) (*SyncSolana, error) {
	return &SyncSolana{
		store:   store,
		service: service,
		config:  config,
		logger:  logger.L,
	}, nil
}

func (s *SyncSolana) MapSolanaData() error {
	collections, err := s.store.Nft.GetSolScanCollections()
	if err != nil {
		s.logger.Fields(logger.Fields{
			"worker": "sync holder collection",
		}).Error(err, "failed query all nft_collection")
	}
	wp := workerpool.New(5)
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

func (s *SyncSolana) mapTokenWithHolder(id string) error {
	holders, err := s.getAllHolder(id)
	if err != nil {
		s.logger.Fields(logger.Fields{
			"collection id": id,
		}).Error(err, "failed to get all holder")
		return err
	}
	wp1 := workerpool.New(5)

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
					"holder": holder,
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

func (s *SyncSolana) getAllHolder(id string) ([]string, error) {
	offset := 0
	limit := 50
	var holders []string
	for {
		nftHolder, err := s.getCollectionHolder(id, offset, limit)
		if err != nil {
			return holders, err
		}
		if len(nftHolder.Data.Data) == 0 {
			break
		}
		for _, v := range nftHolder.Data.Data {
			holders = append(holders, v.WalletAddress)
		}
		offset += limit
	}

	return holders, nil
}

func (s *SyncSolana) getAllNftFromHolder(id string) ([]model.NftSolScanData, error) {
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

func (s *SyncSolana) getCollectionHolder(id string, offset, limit int) (*model.NftHolder, error) {
	retry := 0
	var err error
	var nftToken *model.NftHolder
	for retry < 5 {
		nftToken, err = s.fetchCollectionHolder(id, offset, limit)
		s.logger.Fields(logger.Fields{
			"collectionId": id,
			"offset":       offset,
		}).Info("get holder")
		if err != nil {
			s.logger.Fields(logger.Fields{
				"collectionId": id,
				"offset":       offset,
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

func (s *SyncSolana) getNftTokenFromHolder(holderAddress string, offset, limit int) (*model.NftFromHolder, error) {
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

func (s *SyncSolana) fetchCollectionHolder(id string, offset, limit int) (*model.NftHolder, error) {
	var client = &http.Client{}
	request, err := http.NewRequest("GET", fmt.Sprintf("%s/public/nft/collection/stats?collectionId=%s&filter=holders&offset=%d&limit=%d", s.config.SolScan.BaseUrl, id, offset, limit), nil)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")
	//request.Header.Add("token", SolScanToken)

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

func (s *SyncSolana) fetchNftTokenFromHolder(holderAddress string, offset, limit int) (*model.NftFromHolder, error) {
	var client = &http.Client{}
	request, err := http.NewRequest("GET", fmt.Sprintf("%s/wallet/nft/%s?offset=%d&limit=%d", s.config.SolScan.BaseUrl, holderAddress, offset, limit), nil)
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

	res := &model.NftFromHolder{}
	err = json.Unmarshal(resBody, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
