package handler

import (
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/consolelabs/indexer-api/pkg/config"
	"github.com/consolelabs/indexer-api/pkg/entity"
	"github.com/consolelabs/indexer-api/pkg/logger"
	"github.com/consolelabs/indexer-api/pkg/model"
	"github.com/consolelabs/indexer-api/pkg/request"
	"github.com/consolelabs/indexer-api/pkg/service"
	"github.com/consolelabs/indexer-api/pkg/store"
	chainstore "github.com/consolelabs/indexer-api/pkg/store/chain_store"
	contractstore "github.com/consolelabs/indexer-api/pkg/store/contract_store"
	nftstore "github.com/consolelabs/indexer-api/pkg/store/nft"
	table "github.com/consolelabs/indexer-api/pkg/store/table"
	"github.com/consolelabs/indexer-api/pkg/utils/testhelper"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestHandler_GetChains(t *testing.T) {
	// load env and test data
	cfg := config.LoadTestConfig()
	loggerMock := logger.NewLogrusLogger()
	serviceMock := service.New(&cfg)
	db := testhelper.LoadTestDB()

	storeMock := &store.Store{
		DB:       db,
		Contract: contractstore.New(db),
		Nft:      nftstore.New(db),
		Chain:    chainstore.New(db),
		Table:    *table.New(db),
	}

	entityMock := entity.New(storeMock, serviceMock)

	tests := []struct {
		name             string
		query            request.GetChainsQuery
		wantCode         int
		wantErr          error
		wantResponsePath string
	}{
		{
			name:             "ok_get_all",
			query:            request.GetChainsQuery{Pagination: &model.Pagination{}},
			wantCode:         200,
			wantErr:          nil,
			wantResponsePath: "testdata/get_chain/200.json",
		},
		{
			name:             "400_sort_error",
			query:            request.GetChainsQuery{Pagination: &model.Pagination{}, Sort: "invalid"},
			wantCode:         400,
			wantErr:          nil,
			wantResponsePath: "testdata/get_chain/400_invalid_sort.json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
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

			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = httptest.NewRequest("GET", fmt.Sprintf("/api/v1/chain?%s", strings.Join(query, "&")), nil)
			h := &Handler{
				store:    storeMock,
				service:  serviceMock,
				logger:   loggerMock,
				entities: entityMock,
			}

			h.GetChains(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)

			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.GetChains] response mismatched")
		})
	}
}
