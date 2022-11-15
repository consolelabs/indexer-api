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

	v1.GET("/nft/ticker", h.GetNftDailyTradingVolume)
	v1.GET("/nft/ticker/:collection_address", h.GetNftCollectionTickers)
	v1.GET("/nft/ticker/:collection_address/:token_id", h.GetNftTokenTickers)
	v1.GET("/nft/metadata/attributes-icon", h.GetNftMetadataAttributesIcon)

	// Wallet
	v1.GET(":wallet_address/nft", h.GetNftTokensByWalletAddress)
}
