package handler

import (
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/consolelabs/indexer-api/pkg/config"
	"github.com/consolelabs/indexer-api/pkg/entity"
	"github.com/consolelabs/indexer-api/pkg/logger"
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

func TestHandler_GetNftDetail(t *testing.T) {

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
		givenAddress     string
		givenToken       string
		wantCode         int
		wantErr          error
		wantResponsePath string
	}{
		{
			name:             "ok_get_all",
			wantCode:         200,
			givenAddress:     "0x7aCeE5D0acC520faB33b3Ea25D4FEEF1FfebDE73",
			givenToken:       "1",
			wantErr:          nil,
			wantResponsePath: "testdata/get_nft_detail/200.json",
		},
		// {
		// 	name:             "400_params_error",
		// 	givenAddress:     "invalid",
		// 	givenToken:       "1",
		// 	wantCode:         400,
		// 	wantErr:          nil,
		// 	wantResponsePath: "testdata/get_nft_detail/400_invalid_address.json",
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = httptest.NewRequest("GET", fmt.Sprintf("/api/v1/nft/%s/%s", tt.givenAddress, tt.givenToken), nil)
			ctx.AddParam("collection_address", tt.givenAddress)
			ctx.AddParam("token_id", tt.givenToken)

			h := &Handler{
				store:    storeMock,
				service:  serviceMock,
				logger:   loggerMock,
				entities: entityMock,
			}

			h.GetNftDetail(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)

			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.GetNFTDetail] response mismatched")
		})
	}
}
