package contract

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"

	"github.com/consolelabs/indexer-api/pkg/config"
	"github.com/consolelabs/indexer-api/pkg/logger"
	"github.com/consolelabs/indexer-api/pkg/model"
	"github.com/consolelabs/indexer-api/pkg/service"
	mock_chaindata "github.com/consolelabs/indexer-api/pkg/service/chain_data/mocks"
	mock_chainexplorer "github.com/consolelabs/indexer-api/pkg/service/chain_explorer/mocks"
	"github.com/consolelabs/indexer-api/pkg/store"
	mock_contractstore "github.com/consolelabs/indexer-api/pkg/store/contract_store/mocks"
	mock_nft "github.com/consolelabs/indexer-api/pkg/store/nft/mocks"
)

func Test_isCreationBlockSynced(t *testing.T) {
	type args struct {
		creationBlockNumber int64
		blockStats          *model.BlockStats
		chainId             int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Synced creation block",
			args: args{
				creationBlockNumber: 1000,
				blockStats: &model.BlockStats{
					FirstBlock: 900,
					LastBlock:  1200,
				},
				chainId: 1,
			},
			want: true,
		},
		{
			name: "Not synced creation block",
			args: args{
				creationBlockNumber: 1300,
				blockStats: &model.BlockStats{
					FirstBlock: 900,
					LastBlock:  1200,
				},
				chainId: 1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isCreationBlockSynced(tt.args.creationBlockNumber, tt.args.blockStats, tt.args.chainId); got != tt.want {
				t.Errorf("isCreationBlockSynced() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_contractEntity_AddContract(t *testing.T) {
	type fields struct {
		store   *store.Store
		service *service.Service
	}
	type args struct {
		contract model.Contract
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cfg := config.LoadTestConfig()
	_ = logger.NewLogrusLogger()
	serviceMock := service.New(&cfg)
	storeMock := store.New(&cfg)
	chainData := mock_chaindata.NewMockIService(ctrl)
	chainExplorer := mock_chainexplorer.NewMockIService(ctrl)
	contractStore := mock_contractstore.NewMockIContract(ctrl)
	nftStore := mock_nft.NewMockINft(ctrl)
	serviceMock.ChainData = chainData
	serviceMock.ChainExplorer = chainExplorer
	storeMock.Contract = contractStore
	storeMock.Nft = nftStore

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Add contract fail due to chain explorer",
			fields: fields{
				store:   storeMock,
				service: serviceMock,
			},
			args: args{
				contract: model.Contract{
					Address:     "0x97cEDf6e9116f3422b52Ac1ae339685D527721e7",
					ChainId:     1,
					GrpcAddress: "0.0.0.0:8081",
				},
			},
			wantErr: true,
		},
		{
			name: "Add contract fail due to not synced creation blocked",
			fields: fields{
				store:   storeMock,
				service: serviceMock,
			},
			args: args{
				contract: model.Contract{
					Address:     "0x97cEDf6e9116f3422b52Ac1ae339685D527721e7",
					ChainId:     1,
					GrpcAddress: "0.0.0.0:8081",
				},
			},
			wantErr: true,
		},
		{
			name: "Add contract fail due to chain data",
			fields: fields{
				store:   storeMock,
				service: serviceMock,
			},
			args: args{
				contract: model.Contract{
					Address:     "0x97cEDf6e9116f3422b52Ac1ae339685D527721e7",
					ChainId:     1,
					GrpcAddress: "0.0.0.0:8081",
				},
			},
			wantErr: true,
		},
	}

	contractStore.EXPECT().GetByAddress("0x97cEDf6e9116f3422b52Ac1ae339685D527721e7").Return(nil, nil).AnyTimes()
	blockStatsModel := model.BlockStats{
		FirstBlock: 11900000,
		LastBlock:  14936154,
	}
	chainExplorer.EXPECT().GetCreationBlockNumber(model.Contract{Address: "0x97cEDf6e9116f3422b52Ac1ae339685D527721e7", ChainId: 1, GrpcAddress: "0.0.0.0:8081"}).Return(int64(1), nil).AnyTimes()
	chainData.EXPECT().GetBlockStatsByChainId(int64(1)).Return(&blockStatsModel, nil).AnyTimes()
	nftStore.EXPECT().SaveNftCollection(model.NftCollection{
		Address: "0x97cEDf6e9116f3422b52Ac1ae339685D527721e7",
		ChainId: 1,
		Name:    "Test",
		Symbol:  "Test",
	}).Return(nil).AnyTimes()
	contractStore.EXPECT().Save(model.Contract{
		Address:     "0x97cEDf6e9116f3422b52Ac1ae339685D527721e7",
		ChainId:     1,
		GrpcAddress: "0.0.0.0:8081",
	}).Return(nil).AnyTimes()
	chainExplorer.EXPECT().GetCreationBlockNumber(model.Contract{Address: "0x97cEDf6e9116f3422b52Ac1ae339685D527721e7", ChainId: 1, GrpcAddress: "0.0.0.0:8081"}).Return(int64(0), errors.New("Fail in chain explorer")).AnyTimes()
	chainData.EXPECT().GetBlockStatsByChainId(int64(1)).Return(nil, errors.New(("Fail in chain data"))).AnyTimes()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &contractEntity{
				store:   tt.fields.store,
				service: tt.fields.service,
			}
			if err := e.AddContract(tt.args.contract); (err != nil) != tt.wantErr {
				t.Errorf("contractEntity.AddContract() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_contractEntity_GetContract(t *testing.T) {
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
	_ = logger.NewLogrusLogger()
	serviceMock := service.New(&cfg)
	storeMock := store.New(&cfg)
	contractStore := mock_contractstore.NewMockIContract(ctrl)
	storeMock.Contract = contractStore
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Contract
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Success",
			fields: fields{
				store:   storeMock,
				service: serviceMock,
			},
			args: args{
				address: "success_address",
			},
			want: &model.Contract{
				ID:               1,
				LastUpdatedTime:  time.Date(2022, 06, 8, 13, 51, 38, 6763, time.UTC),
				LastUpdatedBlock: 1,
				CreationBlock:    1000,
				CreatedTime:      time.Date(2022, 06, 8, 13, 51, 38, 6763, time.UTC),
				Address:          "success_address",
				ChainId:          1,
				Type:             "test",
				IsProxy:          nil,
				LogicAddress:     "test",
				Protocol:         "test",
				GrpcAddress:      "test",
			},
			wantErr: false,
		},
		{
			name: "Fail",
			fields: fields{
				store:   storeMock,
				service: serviceMock,
			},
			args: args{
				address: "fail_address",
			},
			wantErr: true,
		},
	}

	contractSuccess := model.Contract{
		ID:               1,
		LastUpdatedTime:  time.Date(2022, 06, 8, 13, 51, 38, 6763, time.UTC),
		LastUpdatedBlock: 1,
		CreationBlock:    1000,
		CreatedTime:      time.Date(2022, 06, 8, 13, 51, 38, 6763, time.UTC),
		Address:          "success_address",
		ChainId:          1,
		Type:             "test",
		IsProxy:          nil,
		LogicAddress:     "test",
		Protocol:         "test",
		GrpcAddress:      "test",
	}
	contractStore.EXPECT().GetByAddress("success_address").Return(&contractSuccess, nil).AnyTimes()
	contractStore.EXPECT().GetByAddress("fail_address").Return(nil, errors.New("Fail address")).AnyTimes()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &contractEntity{
				store:   tt.fields.store,
				service: tt.fields.service,
			}
			got, err := e.GetContract(tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("contractEntity.GetContract() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("contractEntity.GetContract() = %v, want %v", got, tt.want)
			}
		})
	}
}
