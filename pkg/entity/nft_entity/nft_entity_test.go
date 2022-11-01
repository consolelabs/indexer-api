package nftentity

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/consolelabs/indexer-api/pkg/config"
	"github.com/consolelabs/indexer-api/pkg/logger"
	"github.com/consolelabs/indexer-api/pkg/model"
	"github.com/consolelabs/indexer-api/pkg/request"
	"github.com/consolelabs/indexer-api/pkg/response"
	"github.com/consolelabs/indexer-api/pkg/service"
	mock_chaindata "github.com/consolelabs/indexer-api/pkg/service/chain_data/mocks"
	mock_chainexplorer "github.com/consolelabs/indexer-api/pkg/service/chain_explorer/mocks"
	"github.com/consolelabs/indexer-api/pkg/store"
	"github.com/consolelabs/indexer-api/pkg/store/nft"
	mock_nft "github.com/consolelabs/indexer-api/pkg/store/nft/mocks"
)

func Test_GetNftCollections(t *testing.T) {
	type fields struct {
		store   *store.Store
		service *service.Service
	}
	type args struct {
		request.GetNftCollectionsRequest
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cfg := config.LoadTestConfig()
	logger.NewLogrusLogger()
	serviceMock := service.New(&cfg)
	storeMock := store.New(&cfg)
	chainData := mock_chaindata.NewMockIService(ctrl)
	chainExplorer := mock_chainexplorer.NewMockIService(ctrl)
	nftStore := mock_nft.NewMockINft(ctrl)
	serviceMock.ChainData = chainData
	serviceMock.ChainExplorer = chainExplorer
	storeMock.Nft = nftStore

	e := &entity{
		store:   storeMock,
		service: serviceMock,
	}

	synced := true
	queryAll := nft.NftCollectionQuery{
		Synced: &synced,
		Offset: 0,
		Limit:  10,
		Sort:   "all_time_volume DESC",
	}
	nameStr := "rab"
	searchQuery := nft.NftCollectionQuery{
		Name:   &nameStr,
		Synced: &synced,
		Offset: 0,
		Limit:  10,
		Sort:   "thirty_day_volume DESC",
	}
	sortQuery := nft.NftCollectionQuery{
		Synced: &synced,
		Offset: 0,
		Limit:  10,
		Sort:   "thirty_day_volume DESC",
	}

	collections := []model.NftCollection{
		{
			ID:                 1,
			Address:            "0x7acee5d0acc520fab33b3ea25d4feef1ffebde73",
			Name:               "Cyber Neko",
			Symbol:             "neko",
			ChainId:            250,
			ERCFormat:          "721",
			Supply:             6666,
			IsRarityCalculated: false,
		},
		{
			ID:                 2,
			Address:            "0x7D1070fdbF0eF8752a9627a79b00221b53F231fA",
			Name:               "Cyber Rabby",
			Symbol:             "rabby",
			ChainId:            1,
			ERCFormat:          "721",
			Supply:             2222,
			IsRarityCalculated: false,
		},
	}

	res := &response.GetNftCollectionsResponse{
		PaginationResponse: &response.PaginationResponse{Pagination: model.Pagination{Page: 0, Size: 10}, Total: 2},
		Data:               collections,
	}

	tests := []struct {
		name    string
		args    args
		want    *response.GetNftCollectionsResponse
		wantErr bool
	}{
		{
			name: "successfully get all nft collections",
			args: args{
				request.GetNftCollectionsRequest{
					Pagination: &model.Pagination{
						Page: int64(queryAll.Offset / queryAll.Limit),
						Size: int64(queryAll.Limit),
						Sort: "-all_time_volume",
					},
				},
			},
			want:    res,
			wantErr: false,
		},
		{
			name: "successfully search nft collections by name",
			args: args{request.GetNftCollectionsRequest{
				Name: searchQuery.Name,
				Pagination: &model.Pagination{
					Page: int64(searchQuery.Offset / searchQuery.Limit),
					Size: int64(searchQuery.Limit),
					Sort: "-thirty_day_volume",
				},
			},
			},
			want: &response.GetNftCollectionsResponse{
				Data: res.Data[1:],
				PaginationResponse: &response.PaginationResponse{
					Total:      1,
					Pagination: model.Pagination{Page: 0, Size: 10},
				},
			},
			wantErr: false,
		},
		{
			name: "successfully sort collections by chain_id in ascending order",
			args: args{request.GetNftCollectionsRequest{
				Pagination: &model.Pagination{
					Page: int64(sortQuery.Offset / sortQuery.Limit),
					Size: int64(sortQuery.Limit),
					Sort: "-thirty_day_volume",
				},
			},
			},
			want: &response.GetNftCollectionsResponse{
				Data:               []model.NftCollection{collections[1], collections[0]},
				PaginationResponse: res.PaginationResponse,
			},
			wantErr: false,
		},
	}

	nftStore.EXPECT().GetCollections(queryAll).Return(collections, res.Total, nil).AnyTimes()
	nftStore.EXPECT().GetCollections(searchQuery).Return(collections[1:], int64(1), nil).AnyTimes()
	nftStore.EXPECT().GetCollections(sortQuery).Return([]model.NftCollection{collections[1], collections[0]}, int64(2), nil).AnyTimes()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotData, gotTotal, err := e.GetNftCollections(tt.args.GetNftCollectionsRequest)
			if (err != nil) != tt.wantErr {
				t.Errorf("entity.GetNftCollections() error = %v, wantErr %v", err, tt.wantErr)
			}
			if gotData != nil && tt.want != nil && tt.want.Data != nil && !reflect.DeepEqual(gotData, tt.want.Data) {
				t.Errorf("entity.GetNftCollections() gotData = %v, wantData = %v", gotData, tt.want.Data)
			}
			if tt.want != nil && gotTotal != tt.want.Total {
				t.Errorf("entity.GetNftCollections() gotTotal = %v, wantTotal = %v", gotTotal, tt.want.Total)
			}
		})
	}
}

func Test_GetNftTokens(t *testing.T) {
	type fields struct {
		store   *store.Store
		service *service.Service
	}
	type args struct {
		address string
		request.GetNftTokensRequest
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cfg := config.LoadTestConfig()
	logger.NewLogrusLogger()
	serviceMock := service.New(&cfg)
	storeMock := store.New(&cfg)
	chainData := mock_chaindata.NewMockIService(ctrl)
	chainExplorer := mock_chainexplorer.NewMockIService(ctrl)
	nftStore := mock_nft.NewMockINft(ctrl)
	serviceMock.ChainData = chainData
	serviceMock.ChainExplorer = chainExplorer
	storeMock.Nft = nftStore

	e := &entity{
		store:   storeMock,
		service: serviceMock,
	}

	emptyAddress := ""
	rabbyAddress := "0x7D1070fdbF0eF8752a9627a79b00221b53F231fA"
	invalidQuery := nft.NftTokenQuery{
		CollectionAddress: &emptyAddress,
		Offset:            0,
		Limit:             10,
	}
	queryAll := nft.NftTokenQuery{
		CollectionAddress: &rabbyAddress,
		Offset:            0,
		Limit:             10,
	}
	traits := []string{"mask=red soul fire skull", "background=gray"}
	searchQuery := nft.NftTokenQuery{
		CollectionAddress: &rabbyAddress,
		Offset:            0,
		Limit:             10,
		Traits:            traits,
	}

	res := &response.GetNftTokensResponse{
		PaginationResponse: &response.PaginationResponse{Pagination: model.Pagination{Page: 0, Size: 10}, Total: 2},
		Data: []model.NftToken{
			{
				TokenId:           "1821",
				CollectionAddress: "0x7D1070fdbF0eF8752a9627a79b00221b53F231fA",
				Name:              "Cyber Rabby #1821",
				Description:       "The true pioneer Omnichain NFT to be minted on Ethereum and transferred across chains",
				Amount:            "1",
				Image:             "https://google.com",
				Attributes: []model.NftTokenAttribute{
					{
						TraitType: "Mask",
						Value:     "Red Soul Fire Skull",
						Count:     2,
						Rarity:    "Legendary",
						Frequency: "0.01%",
						ChainId:   1,
					},
					{
						TraitType: "Aspect",
						Value:     "Yang",
						Count:     1133,
						Rarity:    "Common",
						Frequency: "50.99%",
						ChainId:   1,
					},
				},
			},
			{
				TokenId:           "1859",
				CollectionAddress: "0x7D1070fdbF0eF8752a9627a79b00221b53F231fA",
				Name:              "Cyber Rabby #1859",
				Description:       "The true pioneer Omnichain NFT to be minted on Ethereum and transferred across chains",
				Amount:            "1",
				Image:             "https://google.com",
				Attributes: []model.NftTokenAttribute{
					{
						TraitType: "Mask",
						Value:     "Red Soul Fire Skull",
						Count:     2,
						Rarity:    "Legendary",
						Frequency: "0.01%",
						ChainId:   1,
					},
					{
						TraitType: "Aspect",
						Value:     "Yin",
						Count:     1089,
						Rarity:    "Common",
						Frequency: "49.01%",
						ChainId:   1,
					},
				},
			},
		},
	}

	tests := []struct {
		name    string
		args    args
		want    *response.GetNftTokensResponse
		wantErr bool
	}{
		{
			name:    "bad request - no address passed",
			args:    args{emptyAddress, request.GetNftTokensRequest{Pagination: &model.Pagination{Page: 0, Size: 10}}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "successfully get rabby tokens",
			args: args{rabbyAddress, request.GetNftTokensRequest{
				Pagination: &model.Pagination{
					Page: int64(queryAll.Offset / queryAll.Limit),
					Size: int64(queryAll.Limit),
				},
			}},
			want:    res,
			wantErr: false,
		},
		{
			name: "successfully search nft collections by name",
			args: args{rabbyAddress, request.GetNftTokensRequest{
				Pagination: &model.Pagination{
					Page: int64(searchQuery.Offset / searchQuery.Limit),
					Size: int64(searchQuery.Limit),
				},
				Traits: traits,
			}},
			want: &response.GetNftTokensResponse{
				Data: res.Data[:1],
				PaginationResponse: &response.PaginationResponse{
					Total:      1,
					Pagination: model.Pagination{Page: 0, Size: 10},
				},
			},
			wantErr: false,
		},
	}

	nftStore.EXPECT().GetNftTokens(invalidQuery).Return(nil, int64(0), errors.New("[Entity][GetNftTokens] failed while [GetNftTokens]")).AnyTimes()
	nftStore.EXPECT().GetNftTokens(queryAll).Return(res.Data, res.Total, nil).AnyTimes()
	nftStore.EXPECT().GetNftTokens(searchQuery).Return(res.Data[:1], int64(1), nil).AnyTimes()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotData, gotTotal, err := e.GetNftTokens(tt.args.address, tt.args.GetNftTokensRequest)
			if (err != nil) != tt.wantErr {
				t.Errorf("entity.GetNftTokens() error = %v, wantErr %v", err, tt.wantErr)
			}
			if gotData != nil && tt.want != nil && tt.want.Data != nil && !reflect.DeepEqual(gotData, tt.want.Data) {
				t.Errorf("entity.GetNftTokens() gotData = %v, wantData = %v", gotData, tt.want.Data)
			}
			if tt.want != nil && gotTotal != tt.want.Total {
				t.Errorf("entity.GetNftTokens() gotTotal = %v, wantTotal = %v", gotTotal, tt.want.Total)
			}
		})
	}
}

// func Test_GetNftCollectionTickers(t *testing.T) {
// 	type fields struct {
// 		store   *store.Store
// 		service *service.Service
// 	}
// 	type args struct {
// 		address string
// 		query   request.GetNftTickersRequest
// 	}

// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()
// 	cfg := config.LoadTestConfig()
// 	logger.NewLogrusLogger()
// 	serviceMock := service.New(&cfg)
// 	storeMock := store.New(&cfg)
// 	chainData := mock_chaindata.NewMockIService(ctrl)
// 	chainExplorer := mock_chainexplorer.NewMockIService(ctrl)
// 	nftStore := mock_nft.NewMockINft(ctrl)
// 	serviceMock.ChainData = chainData
// 	serviceMock.ChainExplorer = chainExplorer
// 	storeMock.Nft = nftStore

// 	e := &entity{
// 		store:   storeMock,
// 		service: serviceMock,
// 	}

// 	invalidAddress := "asd"
// 	collectionAddress := "0x7D1070fdbF0eF8752a9627a79b00221b53F231fA"

// 	token := model.Token{Symbol: "eth", IsNative: true, Address: "asd", Decimals: 18}
// 	snapshots := []model.NftMarketplaceCollectionSnapshot{
// 		{
// 			CreatedTime:   time.UnixMilli(1654926207000),
// 			FloorPriceObj: &model.Price{Amount: "12.12", Token: token},
// 		},
// 		{
// 			CreatedTime:   time.UnixMilli(1655012607000),
// 			FloorPriceObj: &model.Price{Amount: "11", Token: token},
// 		},
// 		{
// 			CreatedTime:   time.UnixMilli(1655099007000),
// 			FloorPriceObj: &model.Price{Amount: "30", Token: token},
// 		},
// 	}

// 	wantRes := response.NftCollectionTickersResponse{
// 		Tickers: response.TokenTickers{
// 			Timestamps: []int64{1654926207000, 1655012607000, 1655099007000},
// 			Prices: []model.Price{
// 				{
// 					Amount: "12.12",
// 					Token:  token,
// 				},
// 				{
// 					Amount: "11",
// 					Token:  token,
// 				},
// 				{
// 					Amount: "30",
// 					Token:  token,
// 				},
// 			},
// 		},
// 		FloorPrice: &model.Price{
// 			Amount: "14.14",
// 			Token:  token,
// 		},
// 		Name:         "Cyber Rabby",
// 		Address:      collectionAddress,
// 		Chain:        "Ethereum Mainnet",
// 		Marketplaces: []string{"paintswap"},
// 		TotalVolume: &model.Price{
// 			Amount: "70.43",
// 			Token:  token,
// 		},
// 	}
// 	tickerQuery := nft.NftTickerQuery{From: 1654926207000, To: 1655099007000, CollectionAddress: collectionAddress}
// 	tokenQuery := nft.NftTokenQuery{CollectionAddress: &collectionAddress, Offset: 0, Limit: 10}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    *response.NftCollectionTickersResponse
// 		wantErr bool
// 	}{
// 		{
// 			name:    "invalid collection address",
// 			args:    args{invalidAddress, request.GetNftTickersRequest{From: tickerQuery.From, To: tickerQuery.To}},
// 			want:    nil,
// 			wantErr: true,
// 		},
// 		{
// 			name:    "get successfully",
// 			args:    args{collectionAddress, request.GetNftTickersRequest{From: tickerQuery.From, To: tickerQuery.To}},
// 			want:    &wantRes,
// 			wantErr: false,
// 		},
// 	}
// 	collection := model.NftCollection{
// 		Name:    "Cyber Rabby",
// 		Address: collectionAddress,
// 		Chain: &model.Chain{
// 			Name: "Ethereum Mainnet",
// 		},
// 	}
// 	latestSnapshot := model.NftMarketplaceCollectionSnapshot{
// 		FloorPriceObj:  wantRes.FloorPrice,
// 		TotalVolumeObj: wantRes.TotalVolume,
// 	}

// 	collectionQuery := nft.NftCollectionQuery{
// 		Address: &invalidAddress,
// 		Limit:   1,
// 	}
// 	// nftStore.EXPECT().GetCollectionByAddress(invalidAddress).Return(nil, fmt.Errorf("record not found")).AnyTimes()
// 	// nftStore.EXPECT().GetCollectionByAddress(collectionAddress).Return(&collection, nil).AnyTimes()
// 	nftStore.EXPECT().GetCollections(collectionQuery).Return(nil, int64(0), fmt.Errorf("record not found"))
// 	collectionQuery.Address = &collectionAddress
// 	nftStore.EXPECT().GetCollections(collectionQuery).Return([]model.NftCollection{collection}, int64(1), nil)
// 	nftStore.EXPECT().GetLatestNftMarketplaceCollectionSnapshot(collectionAddress).Return(&latestSnapshot, nil).AnyTimes()
// 	nftStore.EXPECT().GetPlatformsByCollectionAddress(collectionAddress).Return([]model.MarketplacePlatform{{Name: "paintswap"}}, nil).AnyTimes()
// 	nftStore.EXPECT().GetNftMarketplaceCollectionSnapshots(tickerQuery).Return(snapshots, nil).AnyTimes()
// 	nftStore.EXPECT().GetNftTokens(tokenQuery).Return(nil, int64(3), nil).AnyTimes()
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := e.GetNftCollectionTickers(tt.args.address, tt.args.query)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("entity.GetNftCollectionTickers() error = %v, wantErr = %v", err, tt.wantErr)
// 			}

// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("entity.GetNftCollectionTickers() got = %v, want = %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func Test_GetCollectionMetadata(t *testing.T) {
	type fields struct {
		store   *store.Store
		service *service.Service
	}
	type args struct {
		collectionAddress string
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cfg := config.LoadTestConfig()
	logger.NewLogrusLogger()
	serviceMock := service.New(&cfg)
	storeMock := store.New(&cfg)
	chainData := mock_chaindata.NewMockIService(ctrl)
	chainExplorer := mock_chainexplorer.NewMockIService(ctrl)
	nftStore := mock_nft.NewMockINft(ctrl)
	serviceMock.ChainData = chainData
	serviceMock.ChainExplorer = chainExplorer
	storeMock.Nft = nftStore

	e := &entity{
		store:   storeMock,
		service: serviceMock,
	}

	invalidAddress := "asd"
	validAddress := "0x7aCeE5D0acC520faB33b3Ea25D4FEEF1FfebDE73"
	token := model.Token{Symbol: "eth", IsNative: true, Address: "asd", Decimals: 18}
	res := &model.NftCollection{
		Address: validAddress,
		Name:    "Cyber Neko",
		Symbol:  "NEKO",
		ChainId: 250,
		Stats: &model.ViewNftCollectionStats{
			FloorPriceObj: &model.Price{
				Token:  token,
				Amount: "",
			},
			OneDayVolumeObj: &model.Price{
				Token:  token,
				Amount: "0",
			},
			OneDayVolumeChange: 0,
			SevenDayVolumeObj: &model.Price{
				Token:  token,
				Amount: "226102129999999999983000",
			},
			SevenDayVolumeChange: 385,
			ThirtyDayVolumeObj: &model.Price{
				Token:  token,
				Amount: "229540129999999999983000",
			},
			ThirtyDayVolumeChange: 54,
			AllTimeVolumeObj: &model.Price{
				Token:  token,
				Amount: "336767499999999999993000",
			},
			AllTimeVolumeChange: 336767499999999999993000,
		},
		Marketplace: []model.NftListingMarketplace{{
			ContractAddress: "0x7aCeE5D0acC520faB33b3Ea25D4FEEF1FfebDE73",
			PlatformName:    "paintswap",
		}},
	}

	tests := []struct {
		name    string
		args    args
		want    *model.NftCollection
		wantErr bool
	}{
		{
			name:    "error due to invalid address",
			args:    args{invalidAddress},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "successfully get nft collection metadata",
			args:    args{validAddress},
			want:    res,
			wantErr: false,
		},
	}

	query := nft.NftCollectionQuery{
		Limit:   1,
		Address: &invalidAddress,
	}
	nftStore.EXPECT().GetCollections(query).Return(nil, int64(0), fmt.Errorf("record not found"))
	nftStore.EXPECT().GetNftMarketplaceCollection(invalidAddress).Return(nil, fmt.Errorf("record not found")).AnyTimes()
	//
	query.Address = &validAddress
	nftStore.EXPECT().GetCollections(query).Return([]model.NftCollection{*res}, int64(1), nil)
	nftStore.EXPECT().GetNftMarketplaceCollection(validAddress).Return([]model.NftListingMarketplace{}, nil).AnyTimes()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := e.GetCollectionMetadata(tt.args.collectionAddress)
			if (err != nil) != tt.wantErr {
				t.Errorf("entity.GetCollectionMetadata() error = %v, wantErr %v", err, tt.wantErr)
			}

			if got != nil && tt.want != nil && !reflect.DeepEqual(*got, *tt.want) {
				t.Errorf("entity.GetCollectionMetadata() got = %v, want = %v", *got, *tt.want)
			}
		})
	}
}

func Test_GetTokenActivities(t *testing.T) {
	type fields struct {
		store   *store.Store
		service *service.Service
	}
	type args struct {
		collectionAddress string
		tokenId           string
		*request.GetNftTokenActivitiesRequest
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cfg := config.LoadTestConfig()
	logger.NewLogrusLogger()
	serviceMock := service.New(&cfg)
	storeMock := store.New(&cfg)
	chainData := mock_chaindata.NewMockIService(ctrl)
	chainExplorer := mock_chainexplorer.NewMockIService(ctrl)
	nftStore := mock_nft.NewMockINft(ctrl)
	serviceMock.ChainData = chainData
	serviceMock.ChainExplorer = chainExplorer
	storeMock.Nft = nftStore

	e := &entity{
		store:   storeMock,
		service: serviceMock,
	}

	token := model.Token{Symbol: "FTM", IsNative: true, Address: "0x0000000000000000000000000000000000000000", Decimals: 18}
	tokenId := "1821"
	invalidAddress := "asd"
	validAddress := "0x7aCeE5D0acC520faB33b3Ea25D4FEEF1FfebDE73"
	listings := []*model.NftListing{
		{
			TokenId:         tokenId,
			ContractAddress: validAddress,
			ChainId:         1,
			Quantity:        "1",
			FromAddress:     "0x4b5772dcbE389C438A0134531d7B14069C72937D",
			ToAddress:       "0xB6b648C695baF16093Ad36D262923873CE48a9b1",
			TransactionHash: "0x057fa9cc84bf42b112qwb43a32ae29f1f01a766a202f8c3c60eecf9c4b1755e6",
			ListingStatus:   "sold",
			ListingType:     "buy_now",
			Marketplace: &model.MarketplacePlatform{
				ID:   1,
				Name: "paintswap",
				URL:  "https://paintswap.finance/",
			},
			SoldPrice:    "15000000000000000000",
			ListingPrice: "34000000000000000000",
			SoldPriceObj: &model.Price{
				Token:  token,
				Amount: "15000000000000000000",
			},
			ListingPriceObj: &model.Price{
				Token:  token,
				Amount: "34000000000000000000",
			},
		},
	}

	tests := []struct {
		name    string
		args    args
		want    *response.GetNftTokenActivitiesResponse
		wantErr bool
	}{
		{
			name:    "no data due to invalid address",
			args:    args{invalidAddress, tokenId, &request.GetNftTokenActivitiesRequest{Pagination: &model.Pagination{}}},
			want:    nil,
			wantErr: false,
		},
		{
			name: "successfully get nft token activities",
			args: args{validAddress, tokenId, &request.GetNftTokenActivitiesRequest{Pagination: &model.Pagination{}}},
			want: &response.GetNftTokenActivitiesResponse{
				Data: listings,
				PaginationResponse: &response.PaginationResponse{
					Total: int64(1),
				},
			},
			wantErr: false,
		},
	}

	query := nft.NftListingQuery{
		ContractAddress: invalidAddress,
		TokenId:         tokenId,
		Sort:            "created_time DESC",
		Limit:           50,
	}
	nftStore.EXPECT().GetNftListing(query).Return(nil, int64(0), nil)
	//
	query.ContractAddress = validAddress
	nftStore.EXPECT().GetNftListing(query).Return(listings, int64(1), nil)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotData, gotTotal, err := e.GetTokenActivities(tt.args.collectionAddress, tt.args.tokenId, *tt.args.GetNftTokenActivitiesRequest)
			if (err != nil) != tt.wantErr {
				t.Errorf("entity.GetTokenActivities() error = %v, wantErr %v", err, tt.wantErr)
			}
			if gotData != nil && tt.want != nil && tt.want.Data != nil && !reflect.DeepEqual(gotData, tt.want.Data) {
				t.Errorf("entity.GetTokenActivities() gotData = %v, wantData = %v", gotData, tt.want.Data)
			}
			if tt.want != nil && gotTotal != tt.want.Total {
				t.Errorf("entity.GetTokenActivities() gotTotal = %v, wantTotal = %v", gotTotal, tt.want.Total)
			}
		})
	}
}
