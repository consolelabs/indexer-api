package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gammazero/workerpool"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/consolelabs/indexer-api/pkg/config"
	"github.com/consolelabs/indexer-api/pkg/logger"
	"github.com/consolelabs/indexer-api/pkg/model"
)

var (
	SolScanBaseUrl = config.LoadConfig(config.DefaultConfigLoaders()).Solscan.ProApi
	SolScanToken   = config.LoadConfig(config.DefaultConfigLoaders()).Solscan.Token
)

func (h *Handler) MapSolanaData() error {

	fileContent, err := os.Open("solana_collection_top_10k.json")
	if err != nil {
		logger.L.Fatalf(err, "Can't open file")
		return err
	}

	defer fileContent.Close()

	byteResult, _ := ioutil.ReadAll(fileContent)
	var payload *model.SolscanCollection
	err = json.Unmarshal(byteResult, &payload)
	if err != nil {
		return nil
	}
	wp := workerpool.New(5)
	for _, ev := range payload.Data {
		ev := ev
		wp.Submit(func() {
			err := h.getSlug(ev.CollectionID)
			if err != nil {
				h.logger.Fields(logger.Fields{
					"collectionId": ev.CollectionID,
				}).Error(err, "failed to get data")
			}
			h.logger.Fields(logger.Fields{
				"collectionId": ev.CollectionID,
			}).Infof("get data")
		})
	}
	wp.StopWait()
	return nil
}

func (h *Handler) getOnChainCollectionAddress(id string) error {
	solanaMappingAddress := &model.SolanaMappingAddress{}
	nftCollection, err := h.getSolscanCollection(id)
	if err != nil {
		return err
	}
	if nftCollection != nil {
		if nftCollection.CollectionOnchainID != nil {
			solanaMappingAddress.OnchainAddress = *nftCollection.CollectionOnchainID
			solanaMappingAddress.SolscanId = fmt.Sprintf("solscan-%s", id)
			solanaMappingAddress.CreatedAt = time.Now()
			solanaMappingAddress.UpdatedAt = time.Now()
			err := h.store.Nft.SaveSolanaMapAddress(solanaMappingAddress)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (h *Handler) getSlug(id string) error {
	nftTokenFromSolScan, err := h.getNftTokenFromSolscan(id)
	if err != nil {
		return err
	}
	if len(nftTokenFromSolScan.Data.ListNfts) > 0 {
		firstNft := nftTokenFromSolScan.Data.ListNfts[0]
		nftTokenFromMagiceden, err := h.getNftTokenFromMagiceden(firstNft.NftAddress)
		if err != nil {
			return err
		}
		solanaMappingAddress, err := h.store.Nft.GetSolanaMapAddress(fmt.Sprintf("solscan-%s", id))
		if err != nil {
			return err
		}
		solanaMappingAddress.SolscanId = fmt.Sprintf("solscan-%s", id)
		solanaMappingAddress.Slug = nftTokenFromMagiceden.Collection
		err = h.store.Nft.SaveSolanaMapAddress(solanaMappingAddress)
		if err != nil {
			return err
		}

	}
	return nil
}

func (h *Handler) getSolscanCollection(id string) (*model.NftCollectionSolScan, error) {
	retry := 0
	var err error
	var nftCollection *model.NftCollectionSolScan
	for retry < 5 {
		nftCollectionResponse, err := h.fetchSolscanCollection(id)
		if err != nil {
			h.logger.Fields(logger.Fields{
				"collectionId": id,
				"retry":        retry,
			}).Error(err, "failed to get nft token metadata, retrying")
			retry++
			time.Sleep(7 * time.Duration(retry) * time.Second)
			continue
		}
		if nftCollectionResponse != nil {
			nftCollection = &nftCollectionResponse.Data.Data
		}
		break
	}

	if retry == 5 {
		return nil, err
	}

	return nftCollection, nil
}

func (h *Handler) getNftTokenFromSolscan(id string) (*model.NftTokenSolScanResponse, error) {
	retry := 0
	var err error
	var nftToken *model.NftTokenSolScanResponse
	for retry < 5 {
		nftToken, err = h.fetchNftTokenFromSolscan(id)
		if err != nil {
			h.logger.Fields(logger.Fields{
				"collectionId": id,
				"retry":        retry,
			}).Error(err, "failed to get nft token from collection, retrying")
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

func (h *Handler) getNftTokenFromMagiceden(tokenAddress string) (*model.NftTokenMagicedenReponse, error) {
	retry := 0
	var err error
	var nftToken *model.NftTokenMagicedenReponse
	for retry < 5 {
		nftToken, err = h.fetchNftTokenFromMagiceden(tokenAddress)
		if err != nil {
			h.logger.Fields(logger.Fields{
				"token-address": tokenAddress,
				"retry":         retry,
			}).Error(err, "failed to get nft token from collection, retrying")
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

func (h *Handler) fetchSolscanCollection(id string) (*model.NftCollectionResponse, error) {
	var client = &http.Client{}
	request, err := http.NewRequest("GET", fmt.Sprintf("https://api.solscan.io/collection/id?collectionId=%s&cluster=", id), nil)
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

	res := &model.NftCollectionResponse{}
	err = json.Unmarshal(resBody, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (h *Handler) fetchNftTokenFromSolscan(id string) (*model.NftTokenSolScanResponse, error) {
	var client = &http.Client{}
	request, err := http.NewRequest("GET", fmt.Sprintf("%s/nft/collection/list_nft/%s?page=%s", SolScanBaseUrl, id, "1"), nil)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("token", SolScanToken)

	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	resBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	res := &model.NftTokenSolScanResponse{}
	err = json.Unmarshal(resBody, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (h *Handler) fetchNftTokenFromMagiceden(tokenAddress string) (*model.NftTokenMagicedenReponse, error) {
	var client = &http.Client{}
	request, err := http.NewRequest("GET", fmt.Sprintf("%s/tokens/%s", MagicedenApi, tokenAddress), nil)
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

	res := &model.NftTokenMagicedenReponse{}
	err = json.Unmarshal(resBody, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
