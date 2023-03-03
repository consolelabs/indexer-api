package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gammazero/workerpool"

	"github.com/consolelabs/indexer-api/pkg/logger"
	"github.com/consolelabs/indexer-api/pkg/model"
)

var (
	SolScanBaseUrl = "https://pro-api.solscan.io/v1.0"
	SolScanToken   = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkQXQiOjE2NzU3NzcyODYxMjgsImVtYWlsIjoibmdvdHJvbmdraG9pMTEyQGdtYWlsLmNvbSIsImFjdGlvbiI6InRva2VuLWFwaSIsImlhdCI6MTY3NTc3NzI4Nn0.DCT8Fh8j9uWVpnQSMnq0uuzqeBngNLxc4r8a1Aa2C4Q"
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
			h.logger.Fields(logger.Fields{
				"collectionId": ev.CollectionID,
			}).Infof("get data")
			err := h.mapTokenWithHolder(ev.CollectionID)
			if err != nil {
				h.logger.Fields(logger.Fields{
					"collectionId": ev.CollectionID,
				}).Error(err, "failed to get data")
			}

		})
	}
	wp.StopWait()
	return nil
}

func (h *Handler) mapTokenWithHolder(id string) error {
	holders, err := h.getAllHolder(id)
	if err != nil {
		return err
	}

	for _, holder := range holders {
		nftFromHolder, err := h.getAllNftFromHolder(holder)
		if err != nil {
			return err
		}
		for _, nft := range nftFromHolder {
			if nft.NftCollectionID == id {
				tokenId := ""
				if strings.Contains(nft.NftName, "#") {
					tokenId = strings.Split(nft.NftName, "#")[1]
				}
				owner := model.NftOwner{}
				owner.OwnerAddress = holder
				owner.CollectionAddress = fmt.Sprintf("solscan-%s", id)
				if tokenId == "" {
					owner.TokenId = nft.NftName
				}
				owner.TokenId = tokenId
				owner.ChainId = 999
				owner.CreatedTime = time.Now()
				owner.LastUpdatedTime = time.Now()
				err := h.store.Nft.UpsertOwner(&owner)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (h *Handler) getAllHolder(id string) ([]string, error) {
	page := 1
	var holders []string
	for {
		nftHolder, err := h.getCollectionHolder(id, page)
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

func (h *Handler) getAllNftFromHolder(id string) ([]model.NftSolScanData, error) {
	page := 1
	var nfts []model.NftSolScanData
	for {
		nft, err := h.getNftTokenFromHolder(id, page)
		if err != nil {
			return nfts, err
		}
		if len(nft.Data.ListNfts) == 0 {
			break
		}
		for _, v := range nft.Data.ListNfts {
			nfts = append(nfts, v)
		}
		page += 1
	}

	return nfts, nil
}

func (h *Handler) getCollectionHolder(id string, page int) (*model.NftHolder, error) {
	retry := 0
	var err error
	var nftToken *model.NftHolder
	for retry < 5 {
		nftToken, err = h.fetchCollectionHolder(id, page)
		if err != nil {
			h.logger.Fields(logger.Fields{
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

func (h *Handler) getNftTokenFromHolder(tokenAddress string, page int) (*model.NftFromHolder, error) {
	retry := 0
	var err error
	var nftToken *model.NftFromHolder
	for retry < 5 {
		nftToken, err = h.fetchNftTokenFromHolder(tokenAddress, page)
		if err != nil {
			h.logger.Fields(logger.Fields{
				"token-address": tokenAddress,
				"retry":         retry,
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

func (h *Handler) fetchCollectionHolder(id string, page int) (*model.NftHolder, error) {
	var client = &http.Client{}
	request, err := http.NewRequest("GET", fmt.Sprintf("%s/nft/collection/holders/%s?page=%d", SolScanBaseUrl, id, page), nil)
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

	res := &model.NftHolder{}
	err = json.Unmarshal(resBody, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (h *Handler) fetchNftTokenFromHolder(holderAddress string, page int) (*model.NftFromHolder, error) {
	var client = &http.Client{}
	request, err := http.NewRequest("GET", fmt.Sprintf("%s/nft/wallet/list_nft/%s?page=%d&historry=false", SolScanBaseUrl, holderAddress, page), nil)
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

	res := &model.NftFromHolder{}
	err = json.Unmarshal(resBody, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
