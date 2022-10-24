package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/consolelabs/indexer-api/pkg/config"
	"github.com/consolelabs/indexer-api/pkg/entity"
	mock_chain "github.com/consolelabs/indexer-api/pkg/entity/chain_entity/mocks"
	"github.com/consolelabs/indexer-api/pkg/logger"
	"github.com/consolelabs/indexer-api/pkg/model"
	"github.com/consolelabs/indexer-api/pkg/request"
	"github.com/consolelabs/indexer-api/pkg/response"
	"github.com/consolelabs/indexer-api/pkg/service"
	"github.com/consolelabs/indexer-api/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestHandler_GetChains(t *testing.T) {
	type fields struct {
		store    *store.Store
		service  *service.Service
		logger   logger.Logger
		entities *entity.Entity
	}

	type args struct {
		data string
	}

	type want struct {
		response   interface{}
		expectCode int
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cfg := config.LoadTestConfig()
	loggerMock := logger.NewLogrusLogger()
	serviceMock := service.New(&cfg)
	storeMock := store.New(&cfg)
	entityMock := entity.New(storeMock, serviceMock)
	chain := mock_chain.NewMockIChainEntity(ctrl)
	entityMock.Chain = chain

	tests := []struct {
		name             string
		fields           fields
		query            request.GetChainsQuery
		wantCode         int
		wantErr          error
		wantResponsePath string
	}{
		{
			name:     "ok_get_all",
			query:    request.GetChainsQuery{Pagination: &model.Pagination{}},
			wantCode: 200,
			fields: fields{
				store:    storeMock,
				service:  serviceMock,
				logger:   loggerMock,
				entities: entityMock,
			},
			wantErr:          nil,
			wantResponsePath: "testdata/get_chain/200.json",
		},
	}
	modelChains := response.GetChainsResponse{
		PaginationResponse: &response.PaginationResponse{
			Total: 3,
			Pagination: model.Pagination{
				Page: 0,
				Size: 50,
				Sort: "",
			},
		},
		Data: []model.Chain{
			{
				Id:      1,
				Name:    "Ethereum",
				Symbol:  "ETH",
				ChainId: 1,
				ScanUrl: "",
				IsEvm:   true,
			},
			{
				Id:      2,
				Name:    "Fantom Opera",
				Symbol:  "FTM",
				ChainId: 250,
				ScanUrl: "",
				IsEvm:   true,
			},
			{
				Id:      3,
				Name:    "Optimism",
				Symbol:  "OP",
				ChainId: 10,
				ScanUrl: "",
				IsEvm:   true,
			},
		},
	}
	chain.EXPECT().GetChains(request.GetChainsQuery{Pagination: &model.Pagination{}}).Return(&modelChains, nil).AnyTimes()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			router := gin.Default()
			var query []string
			if tt.query.ChainId != nil {
				query = append(query, fmt.Sprintf("chain_id=%v", *tt.query.ChainId))
			}
			if tt.query.Name != nil {
				query = append(query, fmt.Sprintf("name=%v", *tt.query.Name))
			}
			if tt.query.IsEvm != nil {
				query = append(query, fmt.Sprintf("is_evm=%v", *tt.query.IsEvm))
			}
			if tt.query.Symbol != nil {
				query = append(query, fmt.Sprintf("symbol=%v", *tt.query.Symbol))
			}
			if tt.query.Page != 0 {
				query = append(query, fmt.Sprintf("page=%v", tt.query.Page))
			}
			if tt.query.Size != 0 {
				query = append(query, fmt.Sprintf("size=%v", tt.query.Size))
			}
			if tt.query.Sort != "" {
				query = append(query, fmt.Sprintf("sort=%v", tt.query.Sort))
			}
			h := &Handler{
				store:    tt.fields.store,
				service:  tt.fields.service,
				logger:   tt.fields.logger,
				entities: tt.fields.entities,
			}
			router.GET("/api/v1/chain", h.GetChains)
			recorder := httptest.NewRecorder()

			request, _ := http.NewRequest(http.MethodGet, "/api/v1/chain", nil)
			router.ServeHTTP(recorder, request)

			require.Equal(t, tt.wantCode, recorder.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)

			require.JSONEq(t, string(expRespRaw), recorder.Body.String(), "[Handler.GetNFTDetail] response mismatched")
		})
	}
}
