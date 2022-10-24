package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/consolelabs/indexer-api/pkg/config"
	"github.com/consolelabs/indexer-api/pkg/handler"
)

func loadV1Routes(r *gin.Engine, h *handler.Handler, cfg *config.Config) {
	v1 := r.Group("/api/v1")

	//
	v1.POST("/contract/erc721", h.AddErc721ContractHandler)
	v1.GET("/contract/:address", h.GetContractStatusHandler)

	// NFT
	v1.GET("/nft", h.GetNftCollections)
	v1.GET("/nft/:collection_address", h.GetNftTokens)
	v1.GET("/nft/:collection_address/metadata", h.GetNftCollectionMetadata)
	v1.GET("/nft/:collection_address/:token_id", h.GetNftDetail)
	v1.GET("/nft/:collection_address/:token_id/activity", h.GetNftTokenActivities)
	v1.GET("/nft/:collection_address/:token_id/transaction-history", h.GetNftTokenTransactionHistory)
	v1.GET("/nft/:collection_address/:token_id/bid-history", h.GetNftTokenBidHistory)
	v1.GET("/nft/:collection_address/:token_id/metadata", h.GetNftTokenMetadata)
	v1.GET("/nft/:collection_address/:token_id/sold-prices", h.GetNftTokenListingSoldPrices)
	v1.GET("/nft/:collection_address/:token_id/avg-prices-and-floor", h.GetNftTokenListingAvgPricesAndFloor)
	v1.GET("/nft/:collection_address/query-options", h.GetNftQueryOptions)
	v1.GET("/nft/:collection_address/sale-volume-and-floor-prices", h.GetNftCollectionSaleVolumeAndFloorPrice)
	v1.GET("/nft/:collection_address/listing-unique-owners-time-ranges", h.GetListingUniqueOwnersByTimeRange)
	v1.GET("/nft/:collection_address/listing-across-price-ranges/:payment_token", h.GetNftListingAcrossPriceRanges)
	v1.GET("/nft/:collection_address/listing-across-rarity-rank-ranges/:payment_token", h.GetListingAcrossRarityRankRanges)
	v1.GET("/nft/:collection_address/listing-market-cap-and-volume-by-time-ranges", h.GetListingMarketCapAndVolumeByTimeRanges)
	v1.GET("/nft/:collection_address/market-snapshot-floor-price-by-time-ranges", h.GetMarketSnapshotFloorPriceByTimeRanges)
	v1.GET("/nft/:collection_address/market-snapshot-market-cap-by-time-ranges", h.GetMarketSnapshotMarketCapByTimeRanges)
	v1.GET("/nft/:collection_address/listing-total-volume-by-time-ranges", h.GetListingTotalVolumeByTimeRanges)
	v1.PUT("/nft/:collection_address/resync", h.ResyncNftCollection)
	v1.PUT("/nft/:collection_address/:token_id/resync", h.ResyncNftToken)
	v1.PUT("/nft/:collection_address/metadata/resync", h.ResyncNftCollectionMetadata)
	v1.PUT("/nft/:collection_address/rarity/resync", h.ResyncNftCollectionRarity)
	v1.GET("/nft/search", h.SearchNft)
	v1.POST("/nft/solana/:collection_address", h.ImportSolanaDataHandler)

	v1.GET("nft/ticker", h.GetNftDailyTradingVolume)
	v1.GET("/nft/ticker/:collection_address", h.GetNftCollectionTickers)
	v1.GET("/nft/ticker/:collection_address/:token_id", h.GetNftTokenTickers)
	v1.GET("/nft/metadata/attributes-icon", h.GetNftMetadataAttributesIcon)

	// Wallet
	v1.GET(":wallet_address/nft", h.GetNftTokensByWalletAddress)
	v1.GET(":wallet_address/collection", h.GetNftCollectionsByWalletAddress)

	// Chain
	v1.GET("/chain", h.GetChains)
}
