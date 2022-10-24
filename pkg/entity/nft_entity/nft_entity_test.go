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
	mock_contract "github.com/consolelabs/indexer-api/pkg/store/contract_store/mocks"
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

func Test_GetNftTokensByAddress(t *testing.T) {
	type fields struct {
		store   *store.Store
		service *service.Service
	}
	type args struct {
		address string
		request.GetNftTokensByAddressRequest
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

	ethChainID := int64(1)
	userAddress := "0x52e108267a0E2b46bd5C15f9500aF584fA744fEB"
	rabbyAddress := "0x7D1070fdbF0eF8752a9627a79b00221b53F231fA"
	nekoAddress := "0x7acee5d0acc520fab33b3ea25d4feef1ffebde73"
	invalidQuery := nft.WalletTokenQuery{
		Offset: 0,
		Limit:  10,
	}
	validQuery := nft.WalletTokenQuery{
		WalletAddress: userAddress,
		Offset:        0,
		Limit:         10,
	}
	chainFilterQuery := nft.WalletTokenQuery{
		WalletAddress: userAddress,
		ChainId:       &ethChainID,
		Offset:        0,
		Limit:         10,
	}

	res := &response.Response[[]model.NftToken]{
		PaginationResponse: &response.PaginationResponse{Pagination: model.Pagination{Page: 0, Size: 10}, Total: 4},
		Data: []model.NftToken{
			{
				TokenId:           "16",
				CollectionAddress: rabbyAddress,
				Name:              "Cyber Rabby #16",
				Description:       "The true pioneer Omnichain NFT to be minted on Ethereum and transferred across chains",
				Amount:            "1",
				Image:             "https://google.com",
				Attributes:        []model.NftTokenAttribute{},
			},
			{
				TokenId:           "1821",
				CollectionAddress: rabbyAddress,
				Name:              "Cyber Rabby #1821",
				Description:       "The true pioneer Omnichain NFT to be minted on Ethereum and transferred across chains",
				Amount:            "1",
				Image:             "https://google.com",
				Attributes:        []model.NftTokenAttribute{},
			},
			{
				TokenId:           "5180",
				CollectionAddress: nekoAddress,
				Name:              "Cyber Neko 5180",
				Description:       "Cyber Neko NFT",
				Amount:            "1",
				Image:             "https://google.com",
				Attributes:        []model.NftTokenAttribute{},
			},
			{
				TokenId:           "1780",
				CollectionAddress: rabbyAddress,
				Name:              "Cyber Rabby #1780",
				Description:       "The true pioneer Omnichain NFT to be minted on Ethereum and transferred across chains",
				Amount:            "1",
				Image:             "https://google.com",
				Attributes:        []model.NftTokenAttribute{},
			},
		},
	}

	tests := []struct {
		name    string
		args    args
		want    *response.Response[[]model.NftToken]
		wantErr bool
	}{
		{
			name: "bad request - no address passed",
			args: args{"", request.GetNftTokensByAddressRequest{
				Pagination: &model.Pagination{Page: int64(invalidQuery.Offset / invalidQuery.Limit), Size: 10},
			}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "query all",
			args: args{userAddress, request.GetNftTokensByAddressRequest{
				Pagination: &model.Pagination{Page: int64(validQuery.Offset / validQuery.Limit), Size: int64(validQuery.Limit)},
			}},
			want:    res,
			wantErr: false,
		},
		{
			name: "filter by chain",
			args: args{userAddress, request.GetNftTokensByAddressRequest{
				ChainId:    chainFilterQuery.ChainId,
				Pagination: &model.Pagination{Page: int64(chainFilterQuery.Offset / chainFilterQuery.Limit), Size: int64(chainFilterQuery.Limit)},
			}},
			want: &response.Response[[]model.NftToken]{
				Data: []model.NftToken{res.Data[0], res.Data[1], res.Data[3]},
				PaginationResponse: &response.PaginationResponse{
					Total:      3,
					Pagination: model.Pagination{Page: 0, Size: 10},
				},
			},
			wantErr: false,
		},
	}

	nftStore.EXPECT().GetTokensByWalletAddress(invalidQuery).Return(nil, int64(0), errors.New("[Entity][GetNftTokensByWalletAddress] failed while [GetTokensByWalletAddress]"))
	nftStore.EXPECT().GetTokensByWalletAddress(validQuery).Return(res.Data, res.Total, nil)
	for _, r := range res.Data {
		nftStore.EXPECT().GetAttributeByCollectionAddressTokenID(r.CollectionAddress, r.TokenId).Return([]model.NftTokenAttribute{}, nil).AnyTimes()

	}
	nftStore.EXPECT().GetTokensByWalletAddress(chainFilterQuery).Return([]model.NftToken{res.Data[0], res.Data[1], res.Data[3]}, int64(3), nil)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotData, gotTotal, err := e.GetNftTokensByWalletAddress(tt.args.address, tt.args.GetNftTokensByAddressRequest)
			if (err != nil) != tt.wantErr {
				t.Errorf("entity.GetNftTokensByWalletAddress() error = %v, wantErr %v", err, tt.wantErr)
			}
			if gotData != nil && tt.want != nil && tt.want.Data != nil && !reflect.DeepEqual(gotData, tt.want.Data) {
				t.Errorf("entity.GetNftTokensByWalletAddress() gotData = %v, wantData %v", gotData, tt.want.Data)
			}
			if tt.want != nil && gotTotal != tt.want.Total {
				t.Errorf("entity.GetNftTokensByWalletAddress() gotTotal = %v, wantTotal %v", gotTotal, tt.want.Total)
			}
		})
	}
}

func Test_GetSearchOptions(t *testing.T) {
	type fields struct {
		store   *store.Store
		service *service.Service
	}
	type args struct {
		address string
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
	collectionAddress := "0x7D1070fdbF0eF8752a9627a79b00221b53F231fA"

	marketplaces := []model.MarketplaceWithListingStats{
		{
			MarketplacePlatform: model.MarketplacePlatform{
				ID:   1,
				Name: "paintswap",
				URL:  "https://paintswap.finance/",
			},
			ListingNumber:     2,
			ListingPercentage: 1,
		},
		{
			MarketplacePlatform: model.MarketplacePlatform{
				ID:   2,
				Name: "nftkey",
				URL:  "https://nftkey.app/",
			},
			ListingNumber:     2,
			ListingPercentage: 1,
		},
	}
	attrs := map[string][]model.NftTokenAttribute{
		"Aspect": {
			{
				Count:     1133,
				Frequency: "50.99%",
				TraitType: "Aspect",
				Value:     "Yang",
				Rarity:    "Common",
			},
			{
				Count:     1089,
				Frequency: "49.01%",
				TraitType: "Aspect",
				Value:     "Yin",
				Rarity:    "Common",
			},
		},
	}
	traitCounts := []model.NftTokenAttributeCount{
		{
			Id:    10,
			Count: 1,
			Ratio: 0.5,
		},
		{
			Id:    11,
			Count: 1,
			Ratio: 0.5,
		},
	}

	type wantRes struct {
		marketplaces []model.MarketplaceWithListingStats
		attrs        map[string][]model.NftTokenAttribute
		traitCounts  []model.NftTokenAttributeCount
	}
	tests := []struct {
		name    string
		args    args
		want    wantRes
		wantErr bool
	}{
		{
			name: "return empty",
			args: args{invalidAddress},
			want: wantRes{
				marketplaces: nil,
				attrs:        nil,
				traitCounts:  nil,
			},
			wantErr: true,
		},
		{
			name: "get successfully",
			args: args{collectionAddress},
			want: wantRes{
				marketplaces: marketplaces,
				attrs:        attrs,
				traitCounts:  traitCounts,
			},
			wantErr: false,
		},
	}

	nftStore.EXPECT().GetCollectionByAddress(invalidAddress).Return(nil, fmt.Errorf("record not found")).AnyTimes()
	//
	nftStore.EXPECT().GetCollectionByAddress(collectionAddress).Return(&model.NftCollection{}, nil).AnyTimes()
	nftStore.EXPECT().GetMarketplaceWithListingStats(collectionAddress).Return(marketplaces, nil).AnyTimes()
	nftStore.EXPECT().GetDistinctAttributesByCollectionAddress(collectionAddress).Return([]model.NftTokenAttribute{
		{
			Count:     1133,
			Frequency: "50.99%",
			TraitType: "Aspect",
			Value:     "Yang",
			Rarity:    "Common",
		},
		{
			Count:     1089,
			Frequency: "49.01%",
			TraitType: "Aspect",
			Value:     "Yin",
			Rarity:    "Common",
		},
	}, nil).AnyTimes()
	nftStore.EXPECT().GetCollectionTraitsCount(collectionAddress).Return(traitCounts, nil)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMarketplaces, gotAttrs, gotTraitCounts, err := e.GetQueryOptions(tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("entity.GetQueryOptions() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(gotMarketplaces, tt.want.marketplaces) {
				t.Errorf("entity.GetQueryOptions() got = %v, want %v", gotMarketplaces, tt.want.marketplaces)
			}

			if !reflect.DeepEqual(gotAttrs, tt.want.attrs) {
				t.Errorf("entity.GetQueryOptions() got = %v, want %v", gotAttrs, tt.want.attrs)
			}

			if !reflect.DeepEqual(gotAttrs, tt.want.attrs) {
				t.Errorf("entity.GetQueryOptions() got = %v, want %v", gotTraitCounts, tt.want.traitCounts)
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

func Test_entity_ResyncNftCollection(t *testing.T) {
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
	nftStore := mock_nft.NewMockINft(ctrl)
	contractStore := mock_contract.NewMockIContract(ctrl)
	storeMock.Nft = nftStore
	storeMock.Contract = contractStore

	addr1 := "0x6139b9C548FBd1C50d2768f3464D89c8744aB5f2"
	addr2 := "0x5dB2394A5Abcbb7eE33d09D1d027D0215A76aFCe"
	addr3 := "0xa3AEe8BcE55BEeA1951EF834b99f3Ac60d1ABeeB"
	addr4 := "0xBD4455dA5929D5639EE098ABFaa3241e9ae111Af"
	addr5 := "0x600ea64f353085a0Df6fba7335086951ee90f36f"
	addr6 := "0x7D8820FA92EB1584636f4F5b8515B5476B75171a"
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			fields: fields{
				store:   storeMock,
				service: serviceMock,
			},
			args: args{
				collectionAddress: addr1,
			},
			wantErr: false,
		},
		{
			name: "delete nft token fail",
			fields: fields{
				store:   storeMock,
				service: serviceMock,
			},
			args: args{
				collectionAddress: addr2,
			},
			wantErr: true,
		},
		{
			name: "delete nft token attrs fail",
			fields: fields{
				store:   storeMock,
				service: serviceMock,
			},
			args: args{
				collectionAddress: addr3,
			},
			wantErr: true,
		},
		{
			name: "delete nft owner fail",
			fields: fields{
				store:   storeMock,
				service: serviceMock,
			},
			args: args{
				collectionAddress: addr4,
			},
			wantErr: true,
		},
		{
			name: "delete nft transfer fail",
			fields: fields{
				store:   storeMock,
				service: serviceMock,
			},
			args: args{
				collectionAddress: addr5,
			},
			wantErr: true,
		},
		{
			name: "update last_updated_block = 0 fail",
			fields: fields{
				store:   storeMock,
				service: serviceMock,
			},
			args: args{
				collectionAddress: addr6,
			},
			wantErr: true,
		},
	}

	// success
	contractModel := model.Contract{
		CreationBlock: 0,
		Address:       addr1,
	}
	contractUpdateModel := model.Contract{
		CreationBlock:    0,
		Address:          addr1,
		LastUpdatedBlock: 0,
	}
	nftStore.EXPECT().DeleteNftToken(nft.NftTokenQuery{
		CollectionAddress: &addr1,
	}).Return(nil).AnyTimes()
	nftStore.EXPECT().DeleteNftTokenAttributesByCollectionAddress("0x6139b9C548FBd1C50d2768f3464D89c8744aB5f2").Return(nil).AnyTimes()
	nftStore.EXPECT().DeleteNftOwnerByCollectionAddress("0x6139b9C548FBd1C50d2768f3464D89c8744aB5f2").Return(nil).AnyTimes()
	nftStore.EXPECT().DeleteNftTransferByContractAddress("0x6139b9C548FBd1C50d2768f3464D89c8744aB5f2").Return(nil).AnyTimes()
	contractStore.EXPECT().GetByAddress("0x6139b9C548FBd1C50d2768f3464D89c8744aB5f2").Return(&contractModel, nil).AnyTimes()
	contractStore.EXPECT().Update(&contractUpdateModel).Return(nil).AnyTimes()
	// delete nft token fail
	nftStore.EXPECT().DeleteNftToken(nft.NftTokenQuery{
		CollectionAddress: &addr2,
	}).Return(errors.New("delete nft token fail")).AnyTimes()
	// delete nft token attr fail
	nftStore.EXPECT().DeleteNftToken(nft.NftTokenQuery{
		CollectionAddress: &addr3,
	}).Return(nil).AnyTimes()
	nftStore.EXPECT().DeleteNftTokenAttributesByCollectionAddress("0xa3AEe8BcE55BEeA1951EF834b99f3Ac60d1ABeeB").Return(errors.New("delete nft token attrs fail")).AnyTimes()
	// delete nft owner fail
	nftStore.EXPECT().DeleteNftToken(nft.NftTokenQuery{
		CollectionAddress: &addr4,
	}).Return(nil).AnyTimes()
	nftStore.EXPECT().DeleteNftTokenAttributesByCollectionAddress("0xBD4455dA5929D5639EE098ABFaa3241e9ae111Af").Return(nil).AnyTimes()
	nftStore.EXPECT().DeleteNftOwnerByCollectionAddress("0xBD4455dA5929D5639EE098ABFaa3241e9ae111Af").Return(errors.New("delete nft owner fail")).AnyTimes()
	// delete nft transfer fail
	nftStore.EXPECT().DeleteNftToken(nft.NftTokenQuery{
		CollectionAddress: &addr5,
	}).Return(nil).AnyTimes()
	nftStore.EXPECT().DeleteNftTokenAttributesByCollectionAddress("0x600ea64f353085a0Df6fba7335086951ee90f36f").Return(nil).AnyTimes()
	nftStore.EXPECT().DeleteNftOwnerByCollectionAddress("0x600ea64f353085a0Df6fba7335086951ee90f36f").Return(nil).AnyTimes()
	nftStore.EXPECT().DeleteNftTransferByContractAddress("0x600ea64f353085a0Df6fba7335086951ee90f36f").Return(errors.New("delete nft transfer fail")).AnyTimes()
	// update last_updated_block = 0 fail
	contractModel1 := model.Contract{
		CreationBlock: 0,
		Address:       "0x7D8820FA92EB1584636f4F5b8515B5476B75171a",
	}
	contractUpdateModel1 := model.Contract{
		CreationBlock:    0,
		Address:          "0x7D8820FA92EB1584636f4F5b8515B5476B75171a",
		LastUpdatedBlock: 0,
	}
	nftStore.EXPECT().DeleteNftToken(nft.NftTokenQuery{
		CollectionAddress: &addr6,
	}).Return(nil).AnyTimes()
	nftStore.EXPECT().DeleteNftTokenAttributesByCollectionAddress("0x7D8820FA92EB1584636f4F5b8515B5476B75171a").Return(nil).AnyTimes()
	nftStore.EXPECT().DeleteNftOwnerByCollectionAddress("0x7D8820FA92EB1584636f4F5b8515B5476B75171a").Return(nil).AnyTimes()
	nftStore.EXPECT().DeleteNftTransferByContractAddress("0x7D8820FA92EB1584636f4F5b8515B5476B75171a").Return(nil).AnyTimes()
	contractStore.EXPECT().GetByAddress("0x7D8820FA92EB1584636f4F5b8515B5476B75171a").Return(&contractModel1, nil).AnyTimes()
	contractStore.EXPECT().Update(&contractUpdateModel1).Return(errors.New("delete nft transfer fail")).AnyTimes()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &entity{
				store:   tt.fields.store,
				service: tt.fields.service,
			}
			if err := e.ResyncNftCollection(tt.args.collectionAddress); (err != nil) != tt.wantErr {
				t.Errorf("entity.ResyncNftCollection() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

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
