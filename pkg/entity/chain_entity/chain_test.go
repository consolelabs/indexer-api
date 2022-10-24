package chainentity

import (
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/consolelabs/indexer-api/pkg/config"
	"github.com/consolelabs/indexer-api/pkg/logger"
	"github.com/consolelabs/indexer-api/pkg/model"
	"github.com/consolelabs/indexer-api/pkg/request"
	"github.com/consolelabs/indexer-api/pkg/response"
	"github.com/consolelabs/indexer-api/pkg/service"
	"github.com/consolelabs/indexer-api/pkg/store"
	mockchain "github.com/consolelabs/indexer-api/pkg/store/chain_store/mocks"
)

func Test_GetChains(t *testing.T) {
	type fields struct {
		store   *store.Store
		service *service.Service
	}
	type args struct {
		request.GetChainsQuery
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cfg := config.LoadTestConfig()
	logger.NewLogrusLogger()
	serviceMock := service.New(&cfg)
	storeMock := store.New(&cfg)
	chainStore := mockchain.NewMockIChain(ctrl)
	storeMock.Chain = chainStore

	e := &chainEntity{
		store:   storeMock,
		service: serviceMock,
	}

	queryAll := request.GetChainsQuery{
		Pagination: &model.Pagination{
			Page: 0,
			Size: 10,
		},
	}
	nameStr := "tom"
	searchQuery := request.GetChainsQuery{
		Name: &nameStr,
		Pagination: &model.Pagination{
			Page: 0,
			Size: 10,
		},
	}
	sortQuery := request.GetChainsQuery{
		Pagination: &model.Pagination{
			Sort: "chain_id",
			Page: 0,
			Size: 10,
		},
	}

	res := &response.GetChainsResponse{
		PaginationResponse: &response.PaginationResponse{Pagination: model.Pagination{Page: 0, Size: 10}, Total: 2},
		Data: []model.Chain{
			{
				Name:    "Fantom Opera",
				Symbol:  "FTM",
				ChainId: 250,
			},
			{
				Name:    "Ethereum Mainnet",
				Symbol:  "ETH",
				ChainId: 1,
			},
		},
	}

	tests := []struct {
		name    string
		args    args
		want    *response.GetChainsResponse
		wantErr bool
	}{
		{
			name:    "get all chains",
			args:    args{queryAll},
			want:    res,
			wantErr: false,
		},
		{
			name: "search chain by name",
			args: args{searchQuery},
			want: &response.GetChainsResponse{
				Data: res.Data[:1],
				PaginationResponse: &response.PaginationResponse{
					Total:      1,
					Pagination: model.Pagination{Page: 0, Size: 10},
				},
			},
			wantErr: false,
		},
		{
			name: "sort chains by chain_id in ascending order",
			args: args{sortQuery},
			want: &response.GetChainsResponse{
				Data:               []model.Chain{res.Data[1], res.Data[0]},
				PaginationResponse: res.PaginationResponse,
			},
			wantErr: false,
		},
	}

	chainStore.EXPECT().Get(queryAll).Return(res.Data, res.Total, nil).AnyTimes()
	chainStore.EXPECT().Get(searchQuery).Return(res.Data[:1], int64(1), nil).AnyTimes()
	chainStore.EXPECT().Get(sortQuery).Return([]model.Chain{res.Data[1], res.Data[0]}, int64(2), nil).AnyTimes()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := e.GetChains(tt.args.GetChainsQuery)
			if (err != nil) != tt.wantErr {
				t.Errorf("entity.GetChains() error = %v, wantErr %v", err, tt.wantErr)
			}

			if got != nil && tt.want != nil && !reflect.DeepEqual(*got, *tt.want) {
				t.Errorf("entity.GetChains() got = %v, want %v", *got, *tt.want)
			}
		})
	}
}
