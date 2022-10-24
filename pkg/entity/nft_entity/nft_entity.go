package nftentity

import (
	"encoding/json"
	"fmt"
	"math/big"
	"sort"
	"sync"
	"time"

	"github.com/consolelabs/indexer-api/pkg/constant"
	"github.com/consolelabs/indexer-api/pkg/logger"
	"github.com/consolelabs/indexer-api/pkg/model"
	"github.com/consolelabs/indexer-api/pkg/request"
	"github.com/consolelabs/indexer-api/pkg/response"
	"github.com/consolelabs/indexer-api/pkg/service"
	"github.com/consolelabs/indexer-api/pkg/store"
	"github.com/consolelabs/indexer-api/pkg/store/nft"
	"github.com/consolelabs/indexer-api/pkg/utils"
	"gorm.io/gorm"
)

type entity struct {
	store   *store.Store
	service *service.Service
}

const NFT_TRADING_VOLUME = "[{\"collection_address\": \"0x7D1070fdbF0eF8752a9627a79b00221b53F231fA\", \"collection_name\": \"Cyber Rabby\", \"collection_symbol\": \"rabby\", \"collection_chain_id\": 1, \"trading_volume\": 123456.19, \"token\": \"eth\"}, {\"collection_address\": \"0x7acee5d0acc520fab33b3ea25d4feef1ffebde73\", \"collection_name\": \"Cyber Neko\", \"collection_symbol\": \"neko\", \"collection_chain_id\": 250, \"trading_volume\": 45681.24, \"token\": \"ftm\"}, {\"collection_address\": \"0x7acee5d0acc520fab33b3ea25d4feef1ffebde72\", \"collection_name\": \"Cyber Neko 1\", \"collection_symbol\": \"neko_1\", \"collection_chain_id\": 250, \"trading_volume\": 67853.12, \"token\": \"ftm\"}, {\"collection_address\": \"0x7acee5d0acc520fab33b3ea25d4feef1ffebde71\", \"collection_name\": \"Cyber Neko 2\", \"collection_symbol\": \"neko_2\", \"collection_chain_id\": 250, \"trading_volume\": 981245.54, \"token\": \"ftm\"}, {\"collection_address\": \"0x7acee5d0acc520fab33b3ea25d4feef1ffebde70\", \"collection_name\": \"Cyber Neko 3\", \"collection_symbol\": \"neko_3\", \"collection_chain_id\": 250, \"trading_volume\": 67432.24, \"token\": \"ftm\"}, {\"collection_address\": \"0x7acee5d0acc520fab33b3ea25d4feef1ffebde60\", \"collection_name\": \"Cyber Neko 4\", \"collection_symbol\": \"neko_4\", \"collection_chain_id\": 250, \"trading_volume\": 45681.24, \"token\": \"ftm\"}, {\"collection_address\": \"0x7acee5d0acc520fab33b3ea25d4feef1ffebde61\", \"collection_name\": \"Cyber Neko 5\", \"collection_symbol\": \"neko_5\", \"collection_chain_id\": 250, \"trading_volume\": 74317.43, \"token\": \"ftm\"}, {\"collection_address\": \"0x7acee5d0acc520fab33b3ea25d4feef1ffebde62\", \"collection_name\": \"Cyber Neko 6\", \"collection_symbol\": \"neko_6\", \"collection_chain_id\": 250, \"trading_volume\": 84931.25, \"token\": \"ftm\"},{\"collection_address\": \"0x7acee5d0acc520fab33b3ea25d4feef1ffebde63\", \"collection_name\": \"Cyber Neko 7\", \"collection_symbol\": \"neko_7\", \"collection_chain_id\": 250, \"trading_volume\": 19421.09, \"token\": \"ftm\"},{\"collection_address\": \"0x7acee5d0acc520fab33b3ea25d4feef1ffebde64\", \"collection_name\": \"Cyber Neko 8\", \"collection_symbol\": \"neko_8\", \"collection_chain_id\": 250, \"trading_volume\": 35762.57, \"token\": \"ftm\"},{\"collection_address\": \"0x7acee5d0acc520fab33b3ea25d4feef1ffebde65\", \"collection_name\": \"Cyber Neko 9\", \"collection_symbol\": \"neko_9\", \"collection_chain_id\": 250, \"trading_volume\": 84071.54, \"token\": \"ftm\"},{\"collection_address\": \"0x7acee5d0acc520fab33b3ea25d4feef1ffebde66\", \"collection_name\": \"Cyber Neko 10\", \"collection_symbol\": \"neko_10\", \"collection_chain_id\": 250, \"trading_volume\": 23953.75, \"token\": \"ftm\"},{\"collection_address\": \"0x7acee5d0acc520fab33b3ea25d4feef1ffebde67\", \"collection_name\": \"Cyber Neko 11\", \"collection_symbol\": \"neko_11\", \"collection_chain_id\": 250, \"trading_volume\": 16471.32, \"token\": \"ftm\"},{\"collection_address\": \"0x7acee5d0acc520fab33b3ea25d4feef1ffebde68\", \"collection_name\": \"Cyber Neko 12\", \"collection_symbol\": \"neko_12\", \"collection_chain_id\": 250, \"trading_volume\": 94682.35, \"token\": \"ftm\"},{\"collection_address\": \"0x7acee5d0acc520fab33b3ea25d4feef1ffebde69\", \"collection_name\": \"Cyber Neko 13\", \"collection_symbol\": \"neko_13\", \"collection_chain_id\": 250, \"trading_volume\": 23875.24, \"token\": \"ftm\"}]"

func New(store *store.Store, service *service.Service) INftEntity {
	return &entity{
		store:   store,
		service: service,
	}
}

func (e *entity) GetTokenDetail(address string, tokenId string) (*model.NftToken, error) {
	l := logger.L.Fields(logger.Fields{
		"token_id":           tokenId,
		"collection_address": address,
	})

	collectionQ := nft.NftCollectionQuery{
		Address: &address,
		Limit:   1,
	}
	collections, _, err := e.store.Nft.GetCollections(collectionQ)
	if err != nil {
		logger.L.Fields(logger.Fields{
			"address": address,
			"query":   collectionQ,
		}).Error(err, "[Entity][GetTokenDetail] store.GetCollections failed")
		return nil, err
	}
	if len(collections) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	collection := collections[0]
	collection.Stats.Standardize()

	query := nft.NftTokenQuery{
		TokenId:           &tokenId,
		CollectionAddress: &address,
		Offset:            0,
		Limit:             1,
	}
	tokens, tokensCount, err := e.store.Nft.GetNftTokens(query)
	if err != nil {
		l.Errorf(err, "[Entity][GetTokenDetail] GetNftTokens failed")
		return nil, err
	}

	listing, err := e.store.Nft.GetNftListingByTokenID(address, tokenId)
	if err != nil {
		l.Infof("[Entity][GetTokenDetail] GetNftListingByTokenID dont have data")
	}

	if tokensCount == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	token := &tokens[0]
	token.Rarity = &model.NftTokenRarity{
		Total:  collection.Supply,
		Rarity: token.RarityTier,
		Rank:   token.RarityRank,
		Score:  token.RarityScore,
	}

	token.ListingMarketplace = e.RemoveDuplicateListingMarketplace(collection, listing)

	return token, nil
}

func (e *entity) RemoveDuplicateListingMarketplace(collection model.NftCollection, listingMarketplace []model.NftListingMarketplace) []model.NftListingMarketplace {
	keys := make(map[int]bool)
	list := []model.NftListingMarketplace{}

	for _, entry := range listingMarketplace {
		entry.FloorPrice = collection.Stats.FloorPriceObj.Amount
		if _, value := keys[int(entry.PlatformId)]; !value {
			keys[int(entry.PlatformId)] = true
			list = append(list, entry)
		}
	}
	return list
}

func (e *entity) GetNftCollectionTickers(collectionAddress string, req request.GetNftTickersRequest) (*response.NftCollectionTickersData, error) {
	collectionQ := nft.NftCollectionQuery{
		Address: &collectionAddress,
		Limit:   1,
	}
	collections, _, err := e.store.Nft.GetCollections(collectionQ)
	if err != nil {
		logger.L.Fields(logger.Fields{
			"address": collectionAddress,
			"query":   collectionQ,
		}).Error(err, "[Entity][GetNftCollectionTickers] store.GetCollections failed")
		return nil, err
	}
	if len(collections) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	collection := collections[0]
	collection.Stats.Standardize()

	tickerQuery := nft.NftTickerQuery{
		From:              req.From,
		To:                req.To,
		CollectionAddress: collectionAddress,
	}
	snapshots, err := e.store.Nft.GetNftMarketplaceCollectionSnapshots(tickerQuery)
	if err != nil {
		logger.L.Fields(logger.Fields{
			"query": tickerQuery,
		}).Error(err, "[Entity][GetNftCollectionTickers] store.GetNftMarketplaceCollectionSnapshots failed")
		return nil, err
	}

	// temp calculate like this becasue missing too much data in snapshot
	// TODO(trkhoi): fix when have enough data
	priceChange1d, priceChange7d, priceChange30d := e.CalculateChangePercentageSnapshot(collectionAddress)

	var (
		timestamps   []int64
		prices       []model.Price
		marketplaces []string
	)
	mkMap := make(map[string]bool)
	for _, snapshot := range snapshots {
		timestamps = append(timestamps, snapshot.CreatedTime.UnixMilli())
		prices = append(prices, *snapshot.FloorPriceObj)
		mkMap[snapshot.Marketplace.Name] = true
	}
	for m := range mkMap {
		marketplaces = append(marketplaces, m)
	}

	// last sale
	lastSaleQ := nft.NftListingQuery{
		Sort:            "created_time DESC",
		ContractAddress: collectionAddress,
		ListingStatus:   "sold",
		Limit:           1,
	}
	sales, _, err := e.store.Nft.GetNftListing(lastSaleQ)
	if err != nil {
		logger.L.Fields(logger.Fields{
			"lastSaleQ": lastSaleQ,
		}).Errorf(err, "[Entity][GetNftCollectionTickers] store.GetNftListing failed")
		return nil, err
	}

	res := &response.NftCollectionTickersData{
		Tickers: response.TokenTickers{
			Prices:     prices,
			Timestamps: timestamps,
		},
		Marketplaces:    marketplaces,
		Name:            collection.Name,
		Address:         collection.Address,
		CollectionImage: collection.Image,
		Items:           int64(collection.Supply),
		Owners:          collection.Owners,
		Chain:           collection.Chain,
		TotalVolume:     collection.Stats.AllTimeVolumeObj,
		FloorPrice:      collection.Stats.FloorPriceObj,
		PriceChange1d:   priceChange1d,
		PriceChange7d:   priceChange7d,
		PriceChange30d:  priceChange30d,
	}
	if len(sales) > 0 {
		res.LastSalePrice = sales[0].SoldPriceObj
	}
	return res, nil
}

func (e *entity) GetNftTradingVolume() ([]response.NftTradingVolume, error) {
	res := []response.NftTradingVolume{}
	json.Unmarshal([]byte(NFT_TRADING_VOLUME), &res)
	return res, nil
}

func (e *entity) GetNftCollections(req request.GetNftCollectionsRequest) ([]model.NftCollection, int64, error) {
	req.Standardize()
	synced := true
	query := nft.NftCollectionQuery{
		Name:        req.Name,
		Address:     req.Address,
		Marketplace: req.Marketplace,
		Synced:      &synced,
		Offset:      int(req.Page * req.Size),
		Limit:       int(req.Size),
		Sort:        req.Pagination.Sort,
	}
	collections, total, err := e.store.Nft.GetCollections(query)
	if err != nil {
		logger.L.Fields(logger.Fields{
			"query": query,
		}).Error(err, "[Entity][GetNftCollections] store.GetCollections failed")
	}

	for i := range collections {
		coll := collections[i]
		coll.Stats.Standardize()
	}
	return collections, total, err
}

func (e *entity) GetNftTokens(address string, req request.GetNftTokensRequest) ([]model.NftToken, int64, error) {
	req.Standardize()
	query := nft.NftTokenQuery{
		CollectionAddress: &address,
		TokenId:           req.TokenId,
		Name:              req.Name,
		Currency:          req.Currency,
		ListingType:       req.ListingType,
		ListingStatus:     req.ListingStatus,
		PriceMin:          req.PriceMin,
		PriceMax:          req.PriceMax,
		RarityRankMin:     req.RarityRankMin,
		RarityRankMax:     req.RarityRankMax,
		Marketplaces:      req.Marketplaces,
		Traits:            req.Traits,
		TraitsCount:       req.TraitsCount,
		Offset:            int(req.Size * req.Page),
		Limit:             int(req.Size),
		Sort:              req.Pagination.Sort,
	}
	fmt.Println(query)
	tokens, total, err := e.store.Nft.GetNftTokens(query)
	if err != nil {
		logger.L.Fields(logger.Fields{
			"query": query,
		}).Error(err, "[Entity][GetNftTokens] store.GetNftTokens failed")
	}
	return tokens, total, err
}

func (e *entity) GetNftTokensByWalletAddress(walletAddress string, req request.GetNftTokensByAddressRequest) ([]model.NftToken, int64, error) {
	req.Standardize()
	query := nft.WalletTokenQuery{
		WalletAddress:       walletAddress,
		CollectionAddresses: req.CollectionAddresses,
		ChainId:             req.ChainId,
		Offset:              int(req.Page * req.Size),
		Limit:               int(req.Size),
	}
	tokens, total, err := e.store.Nft.GetTokensByWalletAddress(query)
	if err != nil {
		logger.L.Fields(logger.Fields{
			"query": query,
		}).Error(err, "[Entity][GetNftTokensByWalletAddress] store.GetTokensByWalletAddress failed")
	}
	// tokensAtributes
	for i, token := range tokens {
		att, _ := e.store.Nft.GetAttributeByCollectionAddressTokenID(token.CollectionAddress, token.TokenId)
		tokens[i].Attributes = att
	}
	return tokens, total, err
}

func (e *entity) GetQueryOptions(collectionAddress string) ([]model.MarketplaceWithListingStats, map[string][]model.NftTokenAttribute, []model.NftTokenAttributeCount, error) {
	log := logger.L.Field("collection_address", collectionAddress)
	_, err := e.store.Nft.GetCollectionByAddress(collectionAddress)
	if err != nil {
		log.Error(err, "[Entity.GetQueryOptions] GetCollectionByAddress failed")
		return nil, nil, nil, err
	}

	mCollectionsMetadata, err := e.store.Nft.GetMarketplaceWithListingStats(collectionAddress)
	if err != nil {
		log.Error(err, "[Entity.GetQueryOptions] GetMarketplaceWithListingStats failed")
		return nil, nil, nil, err
	}

	attrs, err := e.store.Nft.GetDistinctAttributesByCollectionAddress(collectionAddress)
	if err != nil {
		log.Error(err, "[Entity.GetQueryOptions] GetDistinctAttributesByCollectionAddress failed")
		return nil, nil, nil, err
	}

	attributesMetadata := make(map[string][]model.NftTokenAttribute)
	for _, attr := range attrs {
		attributesMetadata[attr.TraitType] = append(attributesMetadata[attr.TraitType], attr)
	}

	traitCounts, err := e.store.Nft.GetCollectionTraitsCount(collectionAddress)
	if err != nil {
		log.Error(err, "[Entity.GetQueryOptions] GetCollectionTraitsCount failed")
		return nil, nil, nil, err
	}

	return mCollectionsMetadata, attributesMetadata, traitCounts, nil
}

// used to get collection from handler layer
func (e *entity) GetCollectionByAddress(collectionAddress string) (*model.NftCollection, error) {
	return e.store.Nft.GetCollectionByAddress(collectionAddress)
}

func (e *entity) ResyncNftCollection(collectionAddress string) error {
	deleteQuery := nft.NftTokenQuery{
		CollectionAddress: &collectionAddress,
	}
	err := e.store.Nft.DeleteNftToken(deleteQuery)
	if err != nil {
		logger.L.Fields(logger.Fields{
			"address": collectionAddress,
		}).Error(err, "[nft_entity.DeleteNftTokenByCollectionAddress] Fail to delete nft tokens")
		return err
	}

	err = e.store.Nft.DeleteNftTokenAttributesByCollectionAddress(collectionAddress)
	if err != nil {
		logger.L.Fields(logger.Fields{
			"address": collectionAddress,
		}).Error(err, "[nft_entity.DeleteNftTokenAttributesByCollectionAddress] Fail to delete nft token attributes")
		return err
	}

	err = e.store.Nft.DeleteNftOwnerByCollectionAddress(collectionAddress)
	if err != nil {
		logger.L.Fields(logger.Fields{
			"address": collectionAddress,
		}).Error(err, "[nft_entity.DeleteNftOwnerByCollectionAddress] Fail to delete nft owner")
		return err
	}

	// TODO: remove contract address -> collection address, in db too
	err = e.store.Nft.DeleteNftTransferByContractAddress(collectionAddress)
	if err != nil {
		logger.L.Fields(logger.Fields{
			"address": collectionAddress,
		}).Error(err, "[nft_entity.DeleteNftTransferByContractAddress] Fail to delete nft transfer")
		return err
	}

	// TODO: change to e.store.Contract.ResyncContract(collectionAddress)
	contract, err := e.store.Contract.GetByAddress(collectionAddress)
	if err != nil {
		logger.L.Fields(logger.Fields{
			"address": collectionAddress,
		}).Error(err, "[nft_entity.GetByAddress] Cannot get contract")
		return err
	}
	contract.LastUpdatedBlock = 0

	return e.store.Contract.Update(contract)
}

func (e *entity) GetNftMetadataAttributesIcon(req request.GetNftAttributeIconQuery) (*response.NftMetadataAttributesIconResponse, error) {
	req.Standardize()
	query := nft.NftAttributeIconQuery{
		TraitTypes: req.TraitTypes,
	}

	icons, total, err := e.store.Nft.GetNftMetadataAttributesIcon(query)
	if err != nil {
		logger.L.Error(err, "[nft_entity.GetNftMetadataAttributesIcon] Fail to get nft metadata attributes icon")
		return nil, err
	}
	return &response.NftMetadataAttributesIconResponse{
		Data: icons,
		PaginationResponse: &response.PaginationResponse{
			Pagination: model.Pagination{
				Page: 0,
				Size: total,
			},
			Total: total,
		},
	}, nil
}

func (e *entity) ResyncNftCollectionMetadata(collectionAddress string, chainId int64) error {
	collection, err := e.store.Nft.GetCollectionByAddress(collectionAddress)
	if err != nil {
		logger.L.Field("address", collectionAddress).Error(err, "Failed to get nft collection")
		return err
	}

	switch chainId {
	// if eth, get from opensea
	case constant.C.ChainId.Eth:
		metadata, err := e.service.Opensea.GetOpenseaCollection(collectionAddress)
		if err != nil {
			return err
		}
		collection = collection.MergeOpensea(metadata)

	// if optimism, get from quixotic
	case constant.C.ChainId.Opt:
		metadata, err := e.service.Quixotic.GetQuixoticCollection(collectionAddress)
		if err != nil {
			return err
		}
		collection = collection.MergeQuixotic(metadata)

	// if ftm, get from paintswap
	case constant.C.ChainId.Ftm:
		metadata, err := e.service.Paintswap.GetPaintswapCollection(collectionAddress)
		if err != nil {
			return err
		}
		collection = collection.MergePaintswap(metadata)
	}

	// enrich name & symbol
	name, symbol, err := e.service.Abi.GetNameAndSymbol(collectionAddress, chainId)
	if err != nil {
		logger.L.Error(err, "failed to get name and symbol")
		return err
	}
	collection.Name = name
	collection.Symbol = symbol

	err = e.store.Nft.UpdateCollection(collection)
	if err != nil {
		logger.L.Field("address", collectionAddress).Error(err, "Failed to update image for nft collection")
		return err
	}

	platform, err := e.store.Nft.GetPlatformByName(utils.ConvertChainToMarketplace(collection.ChainId))
	if err != nil {
		logger.L.Field("address", collectionAddress).Error(err, "Failed to get platform by chain")
		return err
	}

	return e.store.Nft.UpsertNftMarketplaceCollection(&model.NftMarketplaceCollection{
		CollectionAddress: collection.Address,
		PlatformID:        int64(platform.ID),
	})
}

type traitTypeValue struct {
	TraitType  string
	TraitValue string
}

func (e *entity) ResyncNftCollectionRarity(collectionAddress string) error {
	collection, err := e.store.Nft.GetCollectionByAddress(collectionAddress)
	if err != nil {
		logger.L.Field("address", collectionAddress).Error(err, "Failed to get nft collection")
		return err
	}
	logger.L.Field("collection", collectionAddress).Info("Start rarity calculate")
	queryAll := nft.NftTokenQuery{
		CollectionAddress: &collection.Address,
		Offset:            0,
		Limit:             10,
	}
	// calculate total token
	_, count, err := e.store.Nft.GetNftTokens(queryAll)
	if err != nil {
		logger.L.Field("collection", collectionAddress).Error(err, "failed GetNftTokens db")
		return err
	}
	if uint64(count) != collection.Supply {
		collection.Supply = uint64(count)
		err = e.store.Nft.UpdateCollection(collection)
		if err != nil {
			logger.L.Field("collection", collectionAddress).Error(err, "failed UpdateCollection db")
			return err
		}
	}

	// get tkAttrs from db
	logger.L.Field("collection", collectionAddress).Info("Get TokenAttributes from db by collectionAddress")
	tkAttrs, err := e.store.Nft.GetAttributesByCollectionAddress(collection.Address)
	if err != nil {
		if err != nil {
			logger.L.Field("collection", collectionAddress).Error(err, "failed to GetAttributesByCollectionAddress from db")
			return err
		}
	}
	mapValueCount := make(map[traitTypeValue]uint64, len(tkAttrs))
	for _, tkAttr := range tkAttrs {
		mapValueCount[traitTypeValue{
			TraitType:  tkAttr.TraitType,
			TraitValue: tkAttr.Value,
		}]++
	}

	mapTokenIdAttributes := map[string][]model.NftTokenAttribute{}
	var chanBuf = make(chan bool, 10)
	wg := sync.WaitGroup{}
	wg.Add(len(tkAttrs))
	mt := sync.Mutex{}
	for _, tkAttr := range tkAttrs {
		go func(tkAttr model.NftTokenAttribute) {
			defer wg.Done()
			chanBuf <- true
			if k, exist := mapValueCount[traitTypeValue{TraitType: tkAttr.TraitType, TraitValue: tkAttr.Value}]; exist && tkAttr.Count != k {
				tkAttr.Count = k
				err = e.store.Nft.AttributeUpsertOne(tkAttr)
				if err != nil {
					logger.L.Fields(logger.Fields{
						"collection": collection.Address,
						"trait_type": tkAttr.TraitType,
						"value":      tkAttr.Value,
					}).Error(err, "failed to UpdateAttributeCount db")
					return
				}
			}
			mt.Lock()
			mapTokenIdAttributes[tkAttr.TokenId] = append(mapTokenIdAttributes[tkAttr.TokenId], tkAttr)
			mt.Unlock()
			<-chanBuf
		}(tkAttr)
	}
	wg.Wait()
	var scores []float64
	for _, attrs := range mapTokenIdAttributes {
		rarityScore := float64(1)
		for _, attr := range attrs {
			rarityScore *= float64(int(attr.Count)) / float64(int(collection.Supply)) * 10
		}
		scores = append(scores, rarityScore)
	}
	sort.Float64s(scores)
	wg.Add(len(mapTokenIdAttributes))
	for tokenId, attrs := range mapTokenIdAttributes {
		go func(tokenId string, attrs []model.NftTokenAttribute) {
			defer wg.Done()
			chanBuf <- true
			// get nft
			tokenQuery := nft.NftTokenQuery{
				CollectionAddress: &collection.Address,
				TokenId:           &tokenId,
				Limit:             1,
			}
			nftTokens, tokensCount, err := e.store.Nft.GetNftTokens(tokenQuery)
			if err != nil {
				logger.L.Fields(logger.Fields{
					"collection": collection.Address,
					"token_id":   tokenId,
				}).Error(err, "fail to calculate rarity")
				return
			}
			if tokensCount == 0 {
				logger.L.Fields(logger.Fields{
					"collection": collection.Address,
					"token_id":   tokenId,
				}).Error(err, "NFT token not found")
				return
			}
			nftToken := &nftTokens[0]
			rarityScore := float64(1)
			for _, attr := range attrs {
				rarityScore *= float64(int(attr.Count)) / float64(int(collection.Supply)) * 10
			}
			nftToken.RarityScore = fmt.Sprintf("%f", rarityScore)
			nftToken.RarityTier = utils.GetRarityTraitFrequency(rarityScore, scores)
			nftToken.RarityRank = uint64(sort.SearchFloat64s(scores, rarityScore) + 1)
			err = e.store.Nft.UpsertToken(nftToken, false)
			if err != nil {
				logger.L.Fields(logger.Fields{
					"collection": nftToken.CollectionAddress,
					"token_id":   nftToken.TokenId,
				}).Error(err, "fail to upsert nft token")
				return
			}
			<-chanBuf
		}(tokenId, attrs)
	}
	wg.Wait()

	return nil
}

func (e *entity) GetCollectionMetadata(collectionAddress string) (*model.NftCollection, error) {
	query := nft.NftCollectionQuery{
		Address: &collectionAddress,
		Limit:   1,
	}
	collections, collectionsCount, err := e.store.Nft.GetCollections(query)
	if err != nil {
		logger.L.Fields(logger.Fields{
			"query": query,
		}).Errorf(err, "[Entity][GetCollectionMetadata] store.GetCollections failed")
		return nil, err
	}
	if collectionsCount == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	coll := collections[0]
	coll.Stats.Standardize()

	// temp do like this to get marketplace for collection
	// TODO(trkhoi): get data from nft_marketplace_collection once it done
	marketplaceCollection, err := e.store.Nft.GetNftMarketplaceCollection(collectionAddress)
	if err != nil {
		logger.L.Fields(logger.Fields{
			"collection_address": collectionAddress,
		}).Errorf(err, "[Entity][GetCollectionMetadata] store.GetNftMarketplaceCollection failed")
		return nil, err
	}

	if len(marketplaceCollection) > 0 {
		coll.Marketplace = marketplaceCollection
	}

	return &coll, err
}

func (e *entity) GetTokenActivities(collectionAddress, tokenId string, req request.GetNftTokenActivitiesRequest) ([]*model.NftListing, int64, error) {
	req.Pagination.Standardize()
	query := nft.NftListingQuery{
		ContractAddress: collectionAddress,
		TokenId:         tokenId,
		Marketplace:     req.Marketplace,
		Offset:          int(req.Pagination.Page * req.Pagination.Size),
		Limit:           int(req.Pagination.Size),
		Sort:            "created_time DESC",
	}
	listings, total, err := e.store.Nft.GetNftListing(query)
	if err != nil {
		logger.L.Fields(logger.Fields{
			"query": query,
		}).Errorf(err, "[Entity][GetTokenActivities] store.GetNftListing failed")
	}
	for i, listing := range listings {
		switch listing.ListingStatus {
		case model.ListingStatusSold:
			listings[i].EventType = "sale"
		case model.ListingStatusCancelled:
			listings[i].EventType = "cancelled"
		case model.ListingStatusListed:
			listings[i].EventType = "listing"
		default:
			listings[i].EventType = "transfer"
		}
	}
	return listings, total, err
}

func (e *entity) GetNftTokenTransactionHistory(collectionAddress, tokenId string) ([]model.NftTxHistory, error) {
	txHistory, err := e.store.Nft.GetNftTokenTxHistory(collectionAddress, tokenId)
	if err != nil {
		logger.L.Fields(logger.Fields{
			"collection_address": collectionAddress,
			"token_id":           tokenId,
		}).Errorf(err, "[Entity][GetTokenMetadata] store.GetNftTokenTxHistory failed")
		return nil, err
	}

	txHistoryFormatted := make([]model.NftTxHistory, 0)
	for _, tx := range txHistory {
		switch tx.ListingStatus {
		case "listed":
			tx.EventType = "listing"
		case "sold":
			tx.EventType = "sale"
		case "cancelled":
			tx.EventType = "cancelled"
		default:
			tx.EventType = "transfer"
		}
		txHistoryFormatted = append(txHistoryFormatted, tx)
	}
	return txHistoryFormatted, nil
}

func (e *entity) GetNftListingSoldPrices(collectionAddress, tokenId, paymentToken, timeDuration, timeInterval, timeRoundUp string) (listingSoldPrices []*model.NftListingSoldPrice, err error) {
	listingSoldPrices, err = e.store.Nft.GetNftListingSoldPrices(collectionAddress, tokenId, paymentToken, timeDuration, timeInterval, timeRoundUp)
	if err != nil {
		logger.L.Fields(logger.Fields{
			"collectionAddress": collectionAddress,
			"tokenId":           tokenId,
			"paymentToken":      paymentToken,
			"timeDuration":      timeDuration,
			"timeInterval":      timeInterval,
			"timeRoundUp":       timeRoundUp,
		}).Errorf(err, "[Entity][GetPrices] store.GetNftListingSoldPrices failed")
		return nil, err
	}
	return listingSoldPrices, err
}

func (e *entity) GetNftListingAvgPricesAndFloor(collectionAddress, tokenId, paymentToken, timeDuration, timeInterval, timeRoundUp string) (listingAvgPricesAndFloors []*model.NftListingAvgPriceAndFloor, err error) {
	listingAvgPricesAndFloors, err = e.store.Nft.GetNftListingAvgPricesAndFloor(collectionAddress, tokenId, paymentToken, timeDuration, timeInterval, timeRoundUp)
	if err != nil {
		logger.L.Fields(logger.Fields{
			"collectionAddress": collectionAddress,
			"tokenId":           tokenId,
			"paymentToken":      paymentToken,
			"timeDuration":      timeDuration,
			"timeInterval":      timeInterval,
			"timeRoundUp":       timeRoundUp,
		}).Errorf(err, "[Entity][GetPrices] store.GetNftListingAvgPricesAndFloor failed")
		return nil, err
	}
	return listingAvgPricesAndFloors, err
}

func (e *entity) GetListingAcrossPriceRanges(collectionAddress, paymentToken, priceRange string) (listingPriceRanges []*model.NftListingPriceRange, err error) {
	listingAcrossPriceRanges, err := e.store.Nft.GetListingAcrossPriceRanges(collectionAddress, paymentToken, priceRange)
	if err != nil {
		logger.L.Fields(logger.Fields{
			"collectionAddress": collectionAddress,
			"paymentToken":      paymentToken,
			"priceRange":        priceRange,
		}).Errorf(err, "[Entity][GetListingAcrossPriceRanges] store.GetListingAcrossPriceRanges failed")
		return nil, err
	}
	return listingAcrossPriceRanges, err
}

func (e *entity) GetListingAcrossRarityRankRanges(collectionAddress, paymentToken, rarityRankRange string) (listingPriceRanges []*model.NftListingRarityRankRange, err error) {
	listingAcrossRarityRankRanges, err := e.store.Nft.GetListingAcrossRarityRankRanges(collectionAddress, paymentToken, rarityRankRange)
	if err != nil {
		logger.L.Fields(logger.Fields{
			"collectionAddress": collectionAddress,
			"paymentToken":      paymentToken,
			"rarityRankRange":   rarityRankRange,
		}).Errorf(err, "[Entity][GetListingAcrossRarityRankRanges] store.GetListingAcrossRarityRankRanges failed")
		return nil, err
	}
	return listingAcrossRarityRankRanges, err
}

func (e *entity) GetCollectionSaleVolumeAndFloorPrice(collectionAddress, timeDuration string) ([]*model.NftCollectionSaleVolumeAndFloorPrice, error) {
	saleVolumeAndFloorPrices, err := e.store.Nft.GetCollectionSaleVolumeAndFloorPrice(collectionAddress, timeDuration)
	if err != nil {
		logger.L.Fields(logger.Fields{
			"collectionAddress": collectionAddress,
			"timeDuration":      timeDuration,
		}).Errorf(err, "[Entity][GetCollectionSaleVolumeAndFloorPrice] store.GetNftListingSoldPrices failed")
		return nil, err
	}
	return saleVolumeAndFloorPrices, err
}

func (e *entity) GetListingMarketCapAndVolumeByTimeRanges(collectionAddress, paymentToken, timeDuration, timeInterval, timeRoundUp string) ([]*model.NftListingMarketCapAndVolumeTimeRange, error) {
	listingMarketCapAndVolumeTimeRanges, err := e.store.Nft.GetListingMarketCapAndVolumeByTimeRanges(collectionAddress, paymentToken, timeDuration, timeInterval, timeRoundUp)
	if err != nil {
		logger.L.Fields(logger.Fields{
			"collectionAddress": collectionAddress,
			"paymentToken":      paymentToken,
			"timeDuration":      timeDuration,
			"timeInterval":      timeInterval,
			"timeRoundUp":       timeRoundUp,
		}).Errorf(err, "[Entity][GetListingMarketCapAndVolumeByTimeRanges] store.GetListingMarketCapAndVolumeByTimeRanges failed")
		return nil, err
	}
	return listingMarketCapAndVolumeTimeRanges, err
}

func (e *entity) GetListingUniqueOwnersByTimeRange(collectionAddress, timeDuration, timeInterval, timeRoundUp string) (listingUniqueOwnerTimeRanges []*model.NftListingUniqueOwnersTimeRange, err error) {
	listingUniqueOwnersTimeRanges, err := e.store.Nft.GetListingUniqueOwnersByTimeRange(collectionAddress, timeDuration, timeInterval, timeRoundUp)
	if err != nil {
		logger.L.Fields(logger.Fields{
			"collectionAddress": collectionAddress,
			"timeDuration":      timeDuration,
			"timeInterval":      timeInterval,
			"timeRoundUp":       timeRoundUp,
		}).Errorf(err, "[Entity][GetListingUniqueOwnersByTimeRange] store.GetListingUniqueOwnersByTimeRange failed")
		return nil, err
	}
	return listingUniqueOwnersTimeRanges, err
}

// TODO: query for real data, remove mock
func (e *entity) GetTokenBidHistory(collectionAddress, tokenId string, req *model.Pagination) ([]*model.NftBid, int64, error) {
	req.Standardize()
	query := nft.NftListingQuery{
		ContractAddress: collectionAddress,
		TokenId:         tokenId,
		Offset:          int(req.Page * req.Size),
		Limit:           int(req.Size),
		Sort:            "created_time DESC",
	}
	listings, total, err := e.store.Nft.GetNftListing(query)
	if err != nil {
		logger.L.Fields(logger.Fields{
			"query": query,
		}).Errorf(err, "[Entity][GetTokenBidsHistory] store.GetNftListing failed")
		return nil, 0, err
	}

	var bids []*model.NftBid
	for _, listing := range listings {
		bid := &model.NftBid{
			TokenId:         listing.TokenId,
			ContractAddress: listing.ContractAddress,
			Marketplace:     listing.Marketplace,
			FromAddress:     listing.FromAddress,
			ToAddress:       listing.ToAddress,
			TransactionHash: listing.TransactionHash,
			BidPrice:        listing.ListingPriceObj,
			ExpiredTime:     listing.CreatedTime,
			UsdPrice:        189.72,
			FloorDifference: -49.17,
		}
		bids = append(bids, bid)
	}

	return bids, total, err
}

func (e *entity) GetTokenMetadata(collectionAddress, tokenId string) (*response.GetNftTokenMetadataResponse, error) {
	res := response.GetNftTokenMetadataData{}
	// latest listing
	tokenQ := nft.NftTokenQuery{
		Limit:             1,
		CollectionAddress: &collectionAddress,
		TokenId:           &tokenId,
	}
	tokens, tokensCount, err := e.store.Nft.GetNftTokens(tokenQ)
	if err != nil {
		logger.L.Fields(logger.Fields{
			"tokenQ": tokenQ,
		}).Errorf(err, "[Entity][GetTokenMetadata] store.GetNftTokens failed")
		return nil, err
	}
	if tokensCount == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	res.LatestListing = tokens[0].Listing

	// token min/max sold price
	priceListingQ := nft.NftListingQuery{
		Sort:            "sold_price DESC",
		ContractAddress: collectionAddress,
		TokenId:         tokenId,
		ListingStatus:   "sold",
	}
	listings, total, err := e.store.Nft.GetNftListing(priceListingQ)
	if err != nil {
		logger.L.Fields(logger.Fields{
			"priceListingQ": priceListingQ,
		}).Errorf(err, "[Entity][GetTokenMetadata] store.GetNftListing2 failed")
		return nil, err
	}
	if total > 0 {
		res.MaxPrice = listings[0].SoldPriceObj
		res.MinPrice = listings[len(listings)-1].SoldPriceObj
	}

	// token sales
	soldListingQ := nft.NftListingQuery{
		Sort:            "created_time DESC",
		ContractAddress: collectionAddress,
		TokenId:         tokenId,
		ListingStatus:   "sold",
	}
	listingsSortByPrice, totalSales, err := e.store.Nft.GetNftListing(soldListingQ)
	if err != nil {
		logger.L.Fields(logger.Fields{
			"soldListingQ": soldListingQ,
		}).Errorf(err, "[Entity][GetTokenMetadata] store.GetNftListing3 failed")
		return nil, err
	}
	if totalSales > 0 {
		res.LastSale = listingsSortByPrice[0].SoldPriceObj
		res.TotalSales = totalSales
	}

	// token total owners, mint date, current/longest hold time
	transferQ := nft.NftTransferQuery{
		CollectionAddress: &collectionAddress,
		TokenId:           &tokenId,
		Sort:              "created_time DESC",
	}
	transfers, totalTransfers, err := e.store.Nft.GetNftTransfers(transferQ)
	if err != nil {
		logger.L.Fields(logger.Fields{
			"transferQ": transferQ,
		}).Errorf(err, "[Entity][GetTokenMetadata] store.GetNftTransfers failed")
		return nil, err
	}
	if totalTransfers > 0 {
		res.TotalOwners = totalTransfers
		firstTransfer := transfers[len(transfers)-1]
		res.Dob = firstTransfer.CreatedTime
		res.Creator = firstTransfer.To
		now := time.Now()
		res.CurrentHoldTimeInSecs = now.Sub(transfers[0].CreatedTime).Seconds()
		// calculate longest hold time in secs
		res.LongestHoldTimeInSecs = now.Sub(transfers[0].CreatedTime).Seconds()
		for i := range transfers {
			if i == len(transfers)-1 {
				continue
			}
			holdTimeInSecs := transfers[i].CreatedTime.Sub(transfers[i+1].CreatedTime).Seconds()
			if holdTimeInSecs > res.LongestHoldTimeInSecs {
				res.LongestHoldTimeInSecs = holdTimeInSecs
			}
		}
	}

	// owner
	ownerQ := nft.NftOwnerQuery{
		CollectionAddress: &collectionAddress,
		TokenId:           &tokenId,
	}
	owners, ownersCount, err := e.store.Nft.GetNftOwners(ownerQ)
	if err != nil {
		logger.L.Fields(logger.Fields{
			"ownerQ": ownerQ,
		}).Errorf(err, "[Entity][GetTokenMetadata] store.GetNftOwners failed")
		return nil, err
	}

	if ownersCount > 0 {
		res.Owner = &owners[0]
	}

	return &response.GetNftTokenMetadataResponse{Data: res}, nil
}

func (e *entity) SearchNft(req *request.SearchNftRequest) ([]model.NftCollection, []model.NftToken, error) {
	req.Standardize()
	collQ := nft.NftCollectionQuery{Limit: int(req.Limit)}
	tokenQ := nft.NftTokenQuery{Limit: int(req.Limit)}
	if isAddr := utils.IsEvmAddress(req.Text); isAddr {
		collQ.Address = &req.Text
		tokenQ.CollectionAddress = &req.Text
	} else {
		collQ.Name = &req.Text
		tokenQ.Name = &req.Text
	}

	collections, _, err := e.store.Nft.GetCollections(collQ)
	if err != nil {
		logger.L.Fields(logger.Fields{
			"collQ": collQ,
		}).Errorf(err, "[Entity][SearchNft] store.GetCollections failed")
		return nil, nil, err
	}
	tokens, _, err := e.store.Nft.GetNftTokens(tokenQ)
	if err != nil {
		logger.L.Fields(logger.Fields{
			"tokenQ": tokenQ,
		}).Errorf(err, "[Entity][SearchNft] store.GetNftTokens failed")
		return nil, nil, err
	}
	return collections, tokens, nil
}

func (e *entity) GetNftCollectionsByWalletAddress(walletAddress string, req request.GetNftCollectionsByWalletAddressRequest) ([]model.NftCollection, int64, error) {
	req.Standardize()
	query := nft.WalletCollectionQuery{
		WalletAddress: walletAddress,
		Offset:        int(req.Page * req.Size),
		Limit:         int(req.Size),
	}
	collections, total, err := e.store.Nft.GetCollectionsByWalletAddress(query)
	if err != nil {
		logger.L.Fields(logger.Fields{
			"query": query,
		}).Error(err, "[Entity][GetNftCollectionsByWalletAddress] store.GetCollectionsByWalletAddress failed")
	}
	return collections, total, err
}

func (e *entity) GetMarketSnapshotFloorPriceByTimeRanges(collectionAddress, paymentToken, timeDuration, timeInterval, timeRoundUp string) ([]*model.NftMarketplaceCollectionSnapshotFloorPriceTimeRange, error) {
	marketplaceCollectionSnapshotFloorPriceTimeRanges, err := e.store.Nft.GetMarketSnapshotFloorPriceByTimeRanges(collectionAddress, paymentToken, timeDuration, timeInterval, timeRoundUp)
	if err != nil {
		logger.L.Fields(logger.Fields{
			"collectionAddress": collectionAddress,
			"paymentToken":      paymentToken,
			"timeDuration":      timeDuration,
			"timeInterval":      timeInterval,
			"timeRoundUp":       timeRoundUp,
		}).Errorf(err, "[Entity][GetMarketSnapshotFloorPriceByTimeRanges] store.GetMarketSnapshotFloorPriceByTimeRanges failed")
		return nil, err
	}
	return marketplaceCollectionSnapshotFloorPriceTimeRanges, err
}

func (e *entity) GetMarketSnapshotMarketCapByTimeRanges(collectionAddress, paymentToken, timeDuration, timeInterval, timeRoundUp string) ([]*model.NftMarketplaceCollectionSnapshotMarketCapTimeRange, error) {
	marketplaceCollectionSnapshotMarketCapTimeRanges, err := e.store.Nft.GetMarketSnapshotMarketCapByTimeRanges(collectionAddress, paymentToken, timeDuration, timeInterval, timeRoundUp)
	if err != nil {
		logger.L.Fields(logger.Fields{
			"collectionAddress": collectionAddress,
			"paymentToken":      paymentToken,
			"timeDuration":      timeDuration,
			"timeInterval":      timeInterval,
			"timeRoundUp":       timeRoundUp,
		}).Errorf(err, "[Entity][GetMarketSnapshotMarketCapByTimeRanges] store.GetMarketSnapshotMarketCapByTimeRanges failed")
		return nil, err
	}
	return marketplaceCollectionSnapshotMarketCapTimeRanges, err
}

func (e *entity) GetListingTotalVolumeByTimeRanges(collectionAddress, paymentToken, timeDuration, timeInterval, timeRoundUp string) ([]*model.NftListingTotalVolumeTimeRange, error) {
	listingTotalVolumeTimeRanges, err := e.store.Nft.GetListingTotalVolumeByTimeRanges(collectionAddress, paymentToken, timeDuration, timeInterval, timeRoundUp)
	if err != nil {
		logger.L.Fields(logger.Fields{
			"collectionAddress": collectionAddress,
			"paymentToken":      paymentToken,
			"timeDuration":      timeDuration,
			"timeInterval":      timeInterval,
			"timeRoundUp":       timeRoundUp,
		}).Errorf(err, "[Entity][GetListingTotalVolumeByTimeRanges] store.GetListingTotalVolumeByTimeRanges failed")
		return nil, err
	}
	return listingTotalVolumeTimeRanges, err
}

func (e *entity) CalculateChangePercentageSnapshot(collectionAddress string) (string, string, string) {
	priceChange1d, priceChange7d, priceChange30d := "0", "0", "0"

	now := time.Now()
	before30Days := now.AddDate(0, 0, -30)

	tickerQuery := nft.NftTickerQuery{
		From:              before30Days.UnixMilli(),
		To:                now.UnixMilli(),
		CollectionAddress: collectionAddress,
	}
	snapshots, err := e.store.Nft.GetNftMarketplaceCollectionSnapshots(tickerQuery)
	if err != nil {
		logger.L.Fields(logger.Fields{
			"query": tickerQuery,
		}).Error(err, "[Entity][GetNftCollectionTickers] store.GetNftMarketplaceCollectionSnapshots failed")
		// return nil, err
		return "", "", ""
	}

	if len(snapshots) > 0 {
		decimals := int(snapshots[len(snapshots)-1].FloorPriceObj.Token.Decimals)
		priceNow := utils.StringWeiToEther(snapshots[len(snapshots)-1].FloorPriceObj.Amount, decimals)
		price1d := utils.StringWeiToEther(snapshots[len(snapshots)-2].FloorPriceObj.Amount, decimals)
		price30d := utils.StringWeiToEther(snapshots[0].FloorPriceObj.Amount, decimals)

		priceChange1dBigFloat := new(big.Float)
		priceChange1dBigFloat = priceChange1dBigFloat.Sub(priceNow, price1d)
		priceChange1dBigFloat = new(big.Float).Quo(priceChange1dBigFloat, price1d)
		priceChange1dBigFloat = priceChange1dBigFloat.Mul(priceChange1dBigFloat, big.NewFloat(100))
		priceChange1d = fmt.Sprintf("%.2f", priceChange1dBigFloat.Abs(priceChange1dBigFloat))

		priceChange30dBigFloat := new(big.Float)
		priceChange30dBigFloat = priceChange30dBigFloat.Sub(priceNow, price30d)
		priceChange30dBigFloat = new(big.Float).Quo(priceChange30dBigFloat, price30d)
		priceChange30dBigFloat = priceChange30dBigFloat.Mul(priceChange30dBigFloat, big.NewFloat(100))
		priceChange30d = fmt.Sprintf("%.2f", priceChange30dBigFloat.Abs(priceChange30dBigFloat))
	}
	return priceChange1d, priceChange7d, priceChange30d
}

func (e *entity) GetNftTokenTickers(collectionAddress, tokenID string, req request.GetNftTickersRequest) (*response.NftTokenTickersData, error) {
	// get data nft token
	tokenQ := nft.NftTokenQuery{
		TokenId:           &tokenID,
		CollectionAddress: &collectionAddress,
		Offset:            0,
		Limit:             1,
	}

	tokens, tokensCount, err := e.store.Nft.GetNftTokens(tokenQ)
	if err != nil {
		logger.L.Fields(logger.Fields{
			"address": collectionAddress,
			"tokenID": tokenID,
			"query":   tokenQ,
		}).Errorf(err, "[Entity.GetNftTokenTickers][store.Nft.GetNftTokens] cannot get token")
		return nil, err
	}

	if tokensCount == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	token := &tokens[0]

	query := nft.NftListingQuery{
		ContractAddress: collectionAddress,
		TokenId:         tokenID,
		From:            req.From,
		To:              req.To,
		ListingStatus:   "sold",
		Sort:            "created_time DESC",
	}

	// get data listing for token tickers
	listings, _, err := e.store.Nft.GetNftListing(query)
	if err != nil {
		logger.L.Fields(logger.Fields{
			"query": query,
		}).Errorf(err, "[Entity][GetTokenActivities] store.GetNftListing failed")
	}

	var (
		timestamps []int64
		prices     []model.Price
	)

	for i, listing := range listings {
		listings[i].EventType = "sale"
		prices = append(prices, *listing.SoldPriceObj)
		timestamps = append(timestamps, listing.CreatedTime.UnixMilli())
	}

	// temp calculate like this becasue missing too much data in snapshot
	// TODO(trkhoi): fix when have enough data
	priceChange1d, priceChange7d, priceChange30d := e.CalculatePriceChangePercentageToken(collectionAddress, tokenID)

	// last sale
	lastSaleQ := nft.NftListingQuery{
		Sort:            "created_time DESC",
		ContractAddress: collectionAddress,
		TokenId:         tokenID,
		ListingStatus:   "sold",
		Limit:           1,
	}
	sales, _, err := e.store.Nft.GetNftListing(lastSaleQ)
	if err != nil {
		logger.L.Fields(logger.Fields{
			"lastSaleQ": lastSaleQ,
		}).Errorf(err, "[Entity][GetNftCollectionTickers] store.GetNftListing failed")
		return nil, err
	}

	res := &response.NftTokenTickersData{
		Tickers: response.TokenTickers{
			Prices:     prices,
			Timestamps: timestamps,
		},

		Name:              token.Name,
		TokenId:           token.TokenId,
		CollectionAddress: token.CollectionAddress,
		Description:       token.Description,
		Image:             token.Image,
		ImageCDN:          token.ImageCDN,
		RarityRank:        token.RarityRank,
		RarityScore:       token.RarityScore,
		RarityTier:        token.RarityTier,

		FloorPrice:     e.CalculateNftTokenFloorPrice(collectionAddress, tokenID),
		PriceChange1d:  priceChange1d,
		PriceChange7d:  priceChange7d,
		PriceChange30d: priceChange30d,
	}
	if len(sales) > 0 {
		res.LastSalePrice = sales[0].SoldPriceObj
	}
	return res, nil
}

func (e *entity) CalculatePriceChangePercentageToken(collectionAddress, tokenID string) (string, string, string) {
	priceChange1d, priceChange7d, priceChange30d := "0", "0", "0"

	query := nft.NftListingQuery{
		ContractAddress: collectionAddress,
		TokenId:         tokenID,
		ListingStatus:   "sold",
		From:            time.Now().AddDate(0, 0, -30).UnixMilli(),
		To:              time.Now().UnixMilli(),
		Sort:            "created_time DESC",
	}
	listings, _, err := e.store.Nft.GetNftListing(query)
	if err != nil {
		logger.L.Fields(logger.Fields{
			"query": query,
		}).Error(err, "[Entity.CalculatePriceChangePercentageToken][store.Nft.GetNftListing] cannot get listing in 30 days")
		// return nil, err
		return "", "", ""
	}

	if len(listings) > 0 {
		decimals := int(listings[len(listings)-1].SoldPriceObj.Token.Decimals)
		priceNow := utils.StringWeiToEther(listings[len(listings)-1].SoldPriceObj.Amount, decimals)
		price1d := utils.StringWeiToEther(listings[len(listings)-2].SoldPriceObj.Amount, decimals)
		price30d := utils.StringWeiToEther(listings[0].SoldPriceObj.Amount, decimals)

		priceChange1dBigFloat := new(big.Float)
		priceChange1dBigFloat = priceChange1dBigFloat.Sub(priceNow, price1d)
		priceChange1dBigFloat = new(big.Float).Quo(priceChange1dBigFloat, price1d)
		priceChange1dBigFloat = priceChange1dBigFloat.Mul(priceChange1dBigFloat, big.NewFloat(100))
		priceChange1d = fmt.Sprintf("%.2f", priceChange1dBigFloat.Abs(priceChange1dBigFloat))

		priceChange30dBigFloat := new(big.Float)
		priceChange30dBigFloat = priceChange30dBigFloat.Sub(priceNow, price30d)
		priceChange30dBigFloat = new(big.Float).Quo(priceChange30dBigFloat, price30d)
		priceChange30dBigFloat = priceChange30dBigFloat.Mul(priceChange30dBigFloat, big.NewFloat(100))
		priceChange30d = fmt.Sprintf("%.2f", priceChange30dBigFloat.Abs(priceChange30dBigFloat))
	}
	return priceChange1d, priceChange7d, priceChange30d
}

func (e *entity) CalculateNftTokenFloorPrice(collectionAddress, tokenID string) (price *model.Price) {
	query := nft.NftListingQuery{
		ContractAddress: collectionAddress,
		TokenId:         tokenID,
		ListingStatus:   "sold",
		Sort:            "CAST(sold_price AS DECIMAL) ASC",
		Limit:           1,
	}

	// get data listing for token tickers
	listings, _, err := e.store.Nft.GetNftListing(query)
	if err != nil {
		logger.L.Fields(logger.Fields{
			"query": query,
		}).Errorf(err, "[Entity][GetTokenActivities] store.GetNftListing failed")
		return price
	}
	if len(listings) > 0 {
		price = listings[0].SoldPriceObj
	}
	return price
}
