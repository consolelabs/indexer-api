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
		"/api/v1/nft/search": {
			Method:  "GET",
			Handler: "github.com/consolelabs/indexer-api/pkg/handler.(*Handler).SearchNft-fm",
		},
		"/api/v1/nft/metadata/attributes-icon": {
			Method:  "GET",
			Handler: "github.com/consolelabs/indexer-api/pkg/handler.(*Handler).GetNftMetadataAttributesIcon-fm",
		},
		"/api/v1/nft/:collection_address": {
			Method:  "GET",
			Handler: "github.com/consolelabs/indexer-api/pkg/handler.(*Handler).GetNftTokens-fm",
		},
		"/api/v1/nft/:collection_address/listing-across-price-ranges/:payment_token": {
			Method:  "GET",
			Handler: "github.com/consolelabs/indexer-api/pkg/handler.(*Handler).GetNftListingAcrossPriceRanges-fm",
		},
		"/api/v1/nft/:collection_address/listing-across-rarity-rank-ranges/:payment_token": {
			Method:  "GET",
			Handler: "github.com/consolelabs/indexer-api/pkg/handler.(*Handler).GetListingAcrossRarityRankRanges-fm",
		},
		"/api/v1/nft/:collection_address/metadata": {
			Method:  "GET",
			Handler: "github.com/consolelabs/indexer-api/pkg/handler.(*Handler).GetNftCollectionMetadata-fm",
		},
		"/api/v1/nft/:collection_address/query-options": {
			Method:  "GET",
			Handler: "github.com/consolelabs/indexer-api/pkg/handler.(*Handler).GetNftQueryOptions-fm",
		},
		"/api/v1/nft/:collection_address/sale-volume-and-floor-prices": {
			Method:  "GET",
			Handler: "github.com/consolelabs/indexer-api/pkg/handler.(*Handler).GetNftCollectionSaleVolumeAndFloorPrice-fm",
		},
		"/api/v1/nft/:collection_address/listing-unique-owners-time-ranges": {
			Method:  "GET",
			Handler: "github.com/consolelabs/indexer-api/pkg/handler.(*Handler).GetListingUniqueOwnersByTimeRange-fm",
		},
		"/api/v1/nft/:collection_address/:token_id": {
			Method:  "GET",
			Handler: "github.com/consolelabs/indexer-api/pkg/handler.(*Handler).GetNftDetail-fm",
		},
		"/api/v1/nft/:collection_address/:token_id/activity": {
			Method:  "GET",
			Handler: "github.com/consolelabs/indexer-api/pkg/handler.(*Handler).GetNftTokenActivities-fm",
		},
		"/api/v1/nft/:collection_address/:token_id/bid-history": {
			Method:  "GET",
			Handler: "github.com/consolelabs/indexer-api/pkg/handler.(*Handler).GetNftTokenBidHistory-fm",
		},
		"/api/v1/nft/:collection_address/:token_id/metadata": {
			Method:  "GET",
			Handler: "github.com/consolelabs/indexer-api/pkg/handler.(*Handler).GetNftTokenMetadata-fm",
		},
		"/api/v1/nft/:collection_address/:token_id/sold-prices": {
			Method:  "GET",
			Handler: "github.com/consolelabs/indexer-api/pkg/handler.(*Handler).GetNftTokenListingSoldPrices-fm",
		},
		"/api/v1/contract/:address": {
			Method:  "GET",
			Handler: "github.com/consolelabs/indexer-api/pkg/handler.(*Handler).GetContractStatusHandler-fm",
		},
		"/api/v1/chain": {
			Method:  "GET",
			Handler: "github.com/consolelabs/indexer-api/pkg/handler.(*Handler).GetChains-fm",
		},
		"/api/v1/:wallet_address/nft": {
			Method:  "GET",
			Handler: "github.com/consolelabs/indexer-api/pkg/handler.(*Handler).GetNftTokensByWalletAddress-fm",
		},
		"/api/v1/nft/:collection_address/resync": {
			Method:  "PUT",
			Handler: "github.com/consolelabs/indexer-api/pkg/handler.(*Handler).ResyncNftCollection-fm",
		},
		"/api/v1/nft/:collection_address/rarity/resync": {
			Method:  "PUT",
			Handler: "github.com/consolelabs/indexer-api/pkg/handler.(*Handler).ResyncNftCollectionRarity-fm",
		},
		"/api/v1/nft/:collection_address/metadata/resync": {
			Method:  "PUT",
			Handler: "github.com/consolelabs/indexer-api/pkg/handler.(*Handler).ResyncNftCollectionMetadata-fm",
		},
		"/api/v1/nft/:collection_address/:token_id/resync": {
			Method:  "PUT",
			Handler: "github.com/consolelabs/indexer-api/pkg/handler.(*Handler).ResyncNftToken-fm",
		},
		"/api/v1/nft/:collection_address/listing-market-cap-and-volume-by-time-ranges": {
			Method:  "GET",
			Handler: "github.com/consolelabs/indexer-api/pkg/handler.(*Handler).GetListingMarketCapAndVolumeByTimeRanges-fm",
		},
		"/api/v1/:wallet_address/collection": {
			Method:  "GET",
			Handler: "github.com/consolelabs/indexer-api/pkg/handler.(*Handler).GetNftCollectionsByWalletAddress-fm",
		},
		"/api/v1/nft/:collection_address/:token_id/avg-prices-and-floor": {
			Method:  "GET",
			Handler: "github.com/consolelabs/indexer-api/pkg/handler.(*Handler).GetNftTokenListingAvgPricesAndFloor-fm",
		},
		"/api/v1/nft/solana/:collection_address": {
			Method:  "POST",
			Handler: "github.com/consolelabs/indexer-api/pkg/handler.(*Handler).ImportSolanaDataHandler-fm",
		},
		"/api/v1/nft/:collection_address/:token_id/transaction-history": {
			Method:  "GET",
			Handler: "github.com/consolelabs/indexer-api/pkg/handler.(*Handler).GetNftTokenTransactionHistory-fm",
		},
		"/api/v1/nft/:collection_address/market-snapshot-floor-price-by-time-ranges": {
			Method:  "GET",
			Handler: "github.com/consolelabs/indexer-api/pkg/handler.(*Handler).GetMarketSnapshotFloorPriceByTimeRanges-fm",
		},
		"/api/v1/nft/:collection_address/market-snapshot-market-cap-by-time-ranges": {
			Method:  "GET",
			Handler: "github.com/consolelabs/indexer-api/pkg/handler.(*Handler).GetMarketSnapshotMarketCapByTimeRanges-fm",
		},
		"/api/v1/nft/:collection_address/listing-total-volume-by-time-ranges": {
			Method:  "GET",
			Handler: "github.com/consolelabs/indexer-api/pkg/handler.(*Handler).GetListingTotalVolumeByTimeRanges-fm",
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
