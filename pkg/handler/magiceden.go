package handler

import (
	"encoding/json"
	"fmt"
	"github.com/consolelabs/indexer-api/pkg/logger"
	"github.com/gammazero/workerpool"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/consolelabs/indexer-api/pkg/model"
)

var MagicedenApi = "https://api-mainnet.magiceden.dev/v2"

func (h *Handler) SyncMagicedenData() error {
	offset := 0
	limit := 200
	chainId := 999
	for {
		h.logger.Fields(logger.Fields{
			"offset": offset,
			"limit":  limit,
		}).Info("start fetching ")
		collections, err := h.getCollection(offset, limit)
		if err != nil {
			h.logger.Fields(logger.Fields{"offset": offset}).Error(err, "failed to get magiceden collection data")
			return err
		}
		if len(collections) == 0 {
			break
		}
		// 3. loop through each events, get token data
		wp := workerpool.New(5)
		for _, ev := range collections {
			ev := ev
			wp.Submit(func() {
				if ev != nil {
					collection, err := h.store.Nft.GetCollectionByNameAndChainID(ev.Name, chainId)
					if err != nil {
						h.logger.Fields(logger.Fields{"collection name": ev.Name}).Error(err, "not found record in nft_collection")
					} else {
						if collection != nil {
							collection.LastUpdatedTime = time.Now()
							if ev.Twitter != "" {
								collection.Twitter = ev.Twitter
							}
							if ev.Discord != "" {
								collection.Discord = ev.Discord
							}
							if ev.Website != "" {
								collection.Website = ev.Website
							}
							if ev.Image != "" {
								collection.Image = ev.Image
							}
							h.store.Nft.UpsertCollection(collection)
						}
					}

				}
			})
		}
		wp.StopWait()

		time.Sleep(1 * time.Second)
		offset += 200
	}
	return nil
}

func (h *Handler) getCollection(offset, limit int) (collection []*model.MagicedenCollection, err error) {
	retry := 0
	for retry < 5 {
		collection, err = h.getMagicedenCollection(offset, limit)
		if err != nil {
			h.logger.Fields(logger.Fields{
				"offset": offset,
				"retry":  retry,
			}).Error(err, "failed to get collection metadata, retrying")
			retry++
			time.Sleep(7 * time.Duration(retry) * time.Second)
			continue
		}
		break
	}

	if retry == 5 {
		return nil, err
	}

	return collection, nil
}

func (h *Handler) getMagicedenCollection(offset, limit int) ([]*model.MagicedenCollection, error) {
	var client = &http.Client{}
	url := fmt.Sprintf("%s/collections?offset=%d&limit=%d", MagicedenApi, offset, limit)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	//request.Header.Add("Content-Type", "application/json")
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	resBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var res []*model.MagicedenCollection
	err = json.Unmarshal(resBody, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
