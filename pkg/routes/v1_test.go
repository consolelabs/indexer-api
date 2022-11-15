package routes

import (
	"fmt"
	"testing"

	"github.com/consolelabs/indexer-api/pkg/handler"
	"github.com/gin-gonic/gin"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"

	"github.com/stretchr/testify/require"
)

func Test_loadV1Routes(t *testing.T) {
	expectedRoutes := map[string]gin.RouteInfo{
		"/api/v1/contract/erc721": {
			Method:  "POST",
			Handler: "github.com/consolelabs/indexer-api/pkg/handler.(*Handler).AddErc721ContractHandler-fm",
		},
		"/api/v1/nft": {
			Method:  "GET",
			Handler: "github.com/consolelabs/indexer-api/pkg/handler.(*Handler).GetNftCollections-fm",
		},
		"/api/v1/nft/ticker": {
			Method:  "GET",
			Handler: "github.com/consolelabs/indexer-api/pkg/handler.(*Handler).GetNftDailyTradingVolume-fm",
		},
		"/api/v1/nft/ticker/:collection_address": {
			Method:  "GET",
			Handler: "github.com/consolelabs/indexer-api/pkg/handler.(*Handler).GetNftCollectionTickers-fm",
		},
		"/api/v1/nft/ticker/:collection_address/:token_id": {
			Method:  "GET",
			Handler: "github.com/consolelabs/indexer-api/pkg/handler.(*Handler).GetNftTokenTickers-fm",
		},
		"/api/v1/nft/metadata/attributes-icon": {
			Method:  "GET",
			Handler: "github.com/consolelabs/indexer-api/pkg/handler.(*Handler).GetNftMetadataAttributesIcon-fm",
		},
		"/api/v1/nft/:collection_address": {
			Method:  "GET",
			Handler: "github.com/consolelabs/indexer-api/pkg/handler.(*Handler).GetNftTokens-fm",
		},
		"/api/v1/nft/:collection_address/metadata": {
			Method:  "GET",
			Handler: "github.com/consolelabs/indexer-api/pkg/handler.(*Handler).GetNftCollectionMetadata-fm",
		},
		"/api/v1/nft/:collection_address/query-options": {
			Method:  "GET",
			Handler: "github.com/consolelabs/indexer-api/pkg/handler.(*Handler).GetNftQueryOptions-fm",
		},
		"/api/v1/nft/:collection_address/:token_id": {
			Method:  "GET",
			Handler: "github.com/consolelabs/indexer-api/pkg/handler.(*Handler).GetNftDetail-fm",
		},
		"/api/v1/nft/:collection_address/:token_id/activity": {
			Method:  "GET",
			Handler: "github.com/consolelabs/indexer-api/pkg/handler.(*Handler).GetNftTokenActivities-fm",
		},
		"/api/v1/contract/:address": {
			Method:  "GET",
			Handler: "github.com/consolelabs/indexer-api/pkg/handler.(*Handler).GetContractStatusHandler-fm",
		},
		"/api/v1/nft/:collection_address/:token_id/transaction-history": {
			Method:  "GET",
			Handler: "github.com/consolelabs/indexer-api/pkg/handler.(*Handler).GetNftTokenTransactionHistory-fm",
		},
	}
	h, err := handler.New(nil, nil, nil)
	require.NoError(t, err)

	router := gin.New()
	loadV1Routes(router, h, nil)

	routeInfo := router.Routes()

	for _, r := range routeInfo {
		require.NotNil(t, r.HandlerFunc)
		expected, ok := expectedRoutes[r.Path]
		require.True(t, ok, fmt.Sprintf("unexpected path: %s", r.Path))
		ignoreFields := cmpopts.IgnoreFields(gin.RouteInfo{}, "HandlerFunc", "Path")
		if !cmp.Equal(expected, r, ignoreFields) {
			t.Errorf("route mismatched. \n route path: %v \n diff: %+v", r.Path,
				cmp.Diff(expected, r, ignoreFields))
			t.FailNow()
		}
	}
}
