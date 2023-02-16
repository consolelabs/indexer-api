package nftentity

import (
	"encoding/json"
	"fmt"
	"math/big"
	"time"

	"gorm.io/gorm"

	"github.com/consolelabs/indexer-api/pkg/logger"
	"github.com/consolelabs/indexer-api/pkg/model"
	"github.com/consolelabs/indexer-api/pkg/request"
	"github.com/consolelabs/indexer-api/pkg/response"
	"github.com/consolelabs/indexer-api/pkg/service"
	"github.com/consolelabs/indexer-api/pkg/store"
	"github.com/consolelabs/indexer-api/pkg/store/nft"
	"github.com/consolelabs/indexer-api/pkg/utils"
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
	priceChange1d, priceChange7d, priceChange30d, alltimeVolume := e.CalculateChangePercentageSnapshot(collectionAddress)

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
		TotalVolume:     alltimeVolume,
		FloorPrice:      e.CalculateNftCollectionFloorPrice(collection.Address),
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

func (e *entity) GetNftSoulBound(collectionAddress string) ([]model.NftTokenAttrSoulBound, error) {
	tokenAttrSoulBound, err := e.store.Nft.GetNftTokenAttrWithSoulBound(collectionAddress)
	if err != nil {
		logger.L.Fields(logger.Fields{
			"collection_address": collectionAddress,
		}).Errorf(err, "[Entity][GetNftSoulBound] store.GetNftTokenAttrWithSoulBound failed")
		return nil, err
	}
	return tokenAttrSoulBound, nil
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
			listings[i].EventType = "sold"
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

func (e *entity) CalculateChangePercentageSnapshot(collectionAddress string) (string, string, string, *model.Price) {
	priceChange1d, priceChange7d, priceChange30d, alltimeVolume := "0", "0", "0", model.Price{}

	stats, err := e.store.Nft.GetNftCollectionStats(collectionAddress)
	if err != nil {
		logger.L.Fields(logger.Fields{
			"collection_address": collectionAddress,
		}).Errorf(err, "[Entity][CalculateChangePercentageSnapshot] store.GetNftCollectionStats failed")
	}

	if len(stats) <= 0 {
		return priceChange1d, priceChange7d, priceChange30d, &alltimeVolume
	}

	priceChange1d = fmt.Sprintf("%.2f", stats[0].OneDayVolumeChange*100)
	priceChange7d = fmt.Sprintf("%.2f", stats[0].SevenDayVolumeChange*100)
	priceChange30d = fmt.Sprintf("%.2f", stats[0].ThirtyDayVolumeChange*100)

	alltimeVolume = model.Price{
		Token:       model.Token{},
		Amount:      fmt.Sprintf("%.f", stats[0].AllTimeVolume),
		AmountInUsd: 0,
	}

	return priceChange1d, priceChange7d, priceChange30d, &alltimeVolume
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
		}).Errorf(err, "[Entity][GetNftTokenTickers] store.GetNftListing failed")
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
	change30d, changePercentage30d := e.CalculatePriceChangePercentageToken(collectionAddress, tokenID, 30)
	change90d, changePercentage90d := e.CalculatePriceChangePercentageToken(collectionAddress, tokenID, 90)
	change365d, changePercentage365d := e.CalculatePriceChangePercentageToken(collectionAddress, tokenID, 365)

	// last sale
	lastSaleQ := nft.NftListingQuery{
		Sort:            "nl.created_time DESC",
		ContractAddress: collectionAddress,
		TokenId:         tokenID,
		ListingStatus:   "sold",
		Limit:           1,
	}
	sales, _, err := e.store.Nft.GetNftListing(lastSaleQ)
	if err != nil {
		logger.L.Fields(logger.Fields{
			"lastSaleQ": lastSaleQ,
		}).Errorf(err, "[Entity][GetNftTokenTickers] store.GetNftListing failed")
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

		FloorPrice: e.CalculateNftTokenFloorPrice(collectionAddress, tokenID),
		// PriceChange1d:  priceChange1d,
		// PriceChange7d:  priceChange7d,
		// PriceChange30d: priceChange30d,
		PriceChange30d:            change30d,
		PriceChange90d:            change90d,
		PriceChange365d:           change365d,
		PriceChangePercentage30d:  changePercentage30d,
		PriceChangePercentage90d:  changePercentage90d,
		PriceChangePercentage365d: changePercentage365d,
	}
	if len(sales) > 0 {
		res.LastSalePrice = sales[0].SoldPriceObj
		res.LastSaleAt = sales[0].CreatedTime
	}
	return res, nil
}

func (e *entity) CalculatePriceChangePercentageToken(collectionAddress, tokenID string, dayRange int) (string, string) {
	// priceChange1d, priceChange7d, priceChange30d := "0", "0", "0"
	var change, changePercentage string

	to := time.Now()
	query := nft.NftListingQuery{
		ContractAddress: collectionAddress,
		TokenId:         tokenID,
		ListingStatus:   "sold",
		From:            to.AddDate(0, 0, -dayRange).UnixMilli(),
		To:              to.UnixMilli(),
		Sort:            "nl.created_time DESC",
	}
	listings, _, err := e.store.Nft.GetNftListing(query)
	if err != nil {
		logger.L.Fields(logger.Fields{
			"query": query,
		}).Error(err, "[Entity.CalculatePriceChangePercentageToken][store.Nft.GetNftListing] cannot get listing in 30 days")
		return change, changePercentage
	}

	if len(listings) == 0 {
		return change, changePercentage
	}

	decimals := int(listings[0].SoldPriceObj.Token.Decimals)
	latestPriceInRange := utils.StringWeiToEther(listings[0].SoldPriceObj.Amount, decimals)
	oldestPriceInRange := utils.StringWeiToEther(listings[len(listings)-1].SoldPriceObj.Amount, decimals)
	pChange := new(big.Float).Sub(latestPriceInRange, oldestPriceInRange)
	if oldestPriceInRange.Cmp(big.NewFloat(0)) != 0 {
		pChangePercentage := new(big.Float).Quo(pChange, oldestPriceInRange)
		change = pChange.String()
		changePercentage = new(big.Float).Mul(pChangePercentage, big.NewFloat(100)).String() + "%"
	}
	return change, changePercentage
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

func (e *entity) CalculateNftCollectionFloorPrice(collectionAddress string) (price *model.Price) {
	query := nft.NftListingQuery{
		ContractAddress: collectionAddress,
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
