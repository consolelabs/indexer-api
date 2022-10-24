package handler

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/consolelabs/indexer-api/pkg/constant"
	"github.com/consolelabs/indexer-api/pkg/logger"
	"github.com/consolelabs/indexer-api/pkg/model"
	"github.com/consolelabs/indexer-api/pkg/queue/message"
	"github.com/consolelabs/indexer-api/pkg/request"
	"github.com/consolelabs/indexer-api/pkg/response"
	"github.com/consolelabs/indexer-api/pkg/utils"
)

// GetNftDetail godoc
// @Summary      Get nft by address
// @Description  Get nft details
// @Tags         NFT
// @Accept       json
// @Produce      json
// @Param        token_id           	path      int  		true  "token Id"
// @Param        collection_address   path      string  true  "collection address"
// @Success      200  {object}  response.NftTokenResponse
// @Router       /nft/{collection_address}/{token_id} [get]
func (h *Handler) GetNftDetail(c *gin.Context) {
	var params struct {
		Address string `uri:"collection_address" binding:"required"`
		TokenID string `uri:"token_id" binding:"required" msg:"token_id is required"`
	}

	if err := c.ShouldBindUri(&params); err != nil {
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, params))
		return
	}

	l := h.logger.Fields(logger.Fields{
		"address": params.Address,
		"id":      params.TokenID,
	})

	// address, err := utils.StringToHashAddress(params.Address)
	// if err != nil {
	// 	l.Error(err, "Can't convert input to hash address")
	// 	c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
	// 	return
	// }

	// entity
	nft, err := h.entities.Nft.GetTokenDetail(params.Address, params.TokenID)
	if utils.IsRecordNotFound(err) {
		l.Error(err, "NFT token not found")
		c.JSON(http.StatusNotFound, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	if err != nil {
		l.Error(err, "Can't get NFT detail data")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ToNftTokenResponse(*nft), nil, nil, nil))
}

// GetNftCollectionTickers godoc
// @Summary      Get floor prices snapshot by nft address (top 7d / 1m / ...)
// @Description  Get floor prices snapshot by nft address (top 7d / 1m / ...)
// @Tags         NFT
// @Accept       json
// @Produce      json
// @Param        collection_address   path      string                          true    "collection address"
// @Param        parameters	          query     request.GetNftTickersRequest    true    "get NFT tickers query"
// @Success      200  {object}  response.NftCollectionTickersResponse
// @Router       /nft/ticker/{collection_address} [get]
func (h *Handler) GetNftCollectionTickers(c *gin.Context) {
	var params struct {
		Address string `uri:"collection_address" binding:"address" msg:"invalid address"`
	}

	if err := c.ShouldBindUri(&params); err != nil {
		h.logger.Field("address", params.Address).Error(err, "Can't bind query params to struct request.GetNftTickersRequest")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, params))
		return
	}

	address, err := utils.StringToHashAddress(params.Address)
	if err != nil {
		h.logger.Field("address", address).Error(err, "Can't convert input to hash address")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	var req request.GetNftTickersRequest
	l := h.logger.Fields(logger.Fields{"address": address, "request": req})
	if err := c.ShouldBindQuery(&req); err != nil {
		l.Error(err, "Invalid from/to query")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, req))
		return
	}

	collection, err := h.entities.Nft.GetNftCollectionTickers(address, req)
	if utils.IsRecordNotFound(err) {
		l.Error(err, "NFT collection not found")
		c.JSON(http.StatusNotFound, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	if err != nil {
		l.Error(err, "Can't get NFT collection tickers")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(collection, nil, nil, nil))
}

// GetNftTokenTickers godoc
// @Summary      Get nft token ticker
// @Description  Get nft token ticker
// @Tags         NFT
// @Accept       json
// @Produce      json
// @Param        collection_address   path      string                          true    "collection address"
// @Param        token_id   path      string                          true    "token_id"
// @Param        parameters	          query     request.GetNftTickersRequest    true    "get NFT tickers query"
// @Success      200  {object}  response.NftTokenTickersDataResponse
// @Router       /nft/ticker/{collection_address}/{token_id} [get]
func (h *Handler) GetNftTokenTickers(c *gin.Context) {
	address := c.Param("collection_address")
	tokenID := c.Param("token_id")

	address, err := utils.StringToHashAddress(address)
	if err != nil {
		h.logger.Field("address", address).Error(err, "Can't convert input to hash address")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	l := h.logger.Fields(logger.Fields{"address": address})
	var req request.GetNftTickersRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		l.Error(err, "Invalid from/to query")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, req))
		return
	}

	collection, err := h.entities.Nft.GetNftTokenTickers(address, tokenID, req)
	if utils.IsRecordNotFound(err) {
		l.Error(err, "NFT collection not found")
		c.JSON(http.StatusNotFound, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	if err != nil {
		l.Error(err, "Can't get NFT collection tickers")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(collection, nil, nil, nil))
}

// GetNftDetail godoc
// @Summary      Get nft daily trading volume
// @Description  Get nft daily trading volume
// @Tags         NFT
// @Accept       json
// @Produce      json
// @Success      200  {object}  response.NftTradingVolumeResponse
// @Router       /nft/ticker [get]
func (h *Handler) GetNftDailyTradingVolume(c *gin.Context) {
	nftTradingVolume, err := h.entities.Nft.GetNftTradingVolume()
	if err != nil {
		h.logger.Error(err, "Can't get NFT trading volume")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(nftTradingVolume, nil, nil, nil))
}

// GetNftCollections godoc
// @Summary      Get nft collections
// @Description  Get nft collections
// @Tags         NFT
// @Accept       json
// @Produce      json
// @Param        query query    request.GetNftCollectionsRequest    false    "get NFT collections query"
// @Success      200   {object}      response.GetNftCollectionsResponse
// @Error        400   {object}      response.ErrorResponse
// @Error        404   {object}      response.ErrorResponse
// @Error        500   {object}      response.ErrorResponse
// @Router       /nft [get]
func (h *Handler) GetNftCollections(c *gin.Context) {
	var req request.GetNftCollectionsRequest
	var err error
	if err = c.ShouldBindQuery(&req); err != nil {
		h.logger.Error(err, "Can't bind query params to struct request.GetNftCollectionsQuery")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	if req.Address != nil {
		if *req.Address, err = utils.StringToHashAddress(*req.Address); err != nil {
			h.logger.Fields(logger.Fields{"request": req}).Error(err, "Can't convert input to hash address")
			c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
			return
		}
	}

	collections, total, err := h.entities.Nft.GetNftCollections(req)
	if err != nil {
		h.logger.Fields(logger.Fields{"request": req}).Error(err, "Can't get NFT collections")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(collections, &response.PaginationResponse{
		Pagination: model.Pagination{
			Page: req.Page,
			Size: req.Size,
		},
		Total: total,
	}, nil, nil))
}

// GetNftTokens godoc
// @Summary      Get tokens of a nft collection
// @Description  Get tokens of a nft collection
// @Tags         NFT
// @Accept       json
// @Produce      json
// @Param        collection_address   path    string                         true   "collection address"
// @Param 			 marketplaces 				query 	[]string 	false 	"marketplaces" Enums(paintswap,nftkey,opensea,looksrare) 						collectionFormat(multi)
// @Param 			 traits 							query 	[]string 	false 	"token traits/attributes" example(mask=Television Head,Zorro Mask) 	collectionFormat(multi)
// @Param 			 traits_count 				query 	[]int 		false 	"traits count" example(11) 																					collectionFormat(multi)
// @Param        parameters           query   request.GetNftTokensRequest    false  "get NFT tokens query"
// @Success      200  {object}  response.GetNftTokensResponse
// @Error        400  {object}  response.ErrorResponse
// @Error        404  {object}  response.ErrorResponse
// @Error        500  {object}  response.ErrorResponse
// @Router       /nft/{collection_address} [get]
func (h *Handler) GetNftTokens(c *gin.Context) {
	var params struct {
		Address string `uri:"collection_address" binding:"address" msg:"invalid address"`
	}

	if err := c.ShouldBindUri(&params); err != nil {
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, params))
		return
	}

	address, err := utils.StringToHashAddress(params.Address)
	if err != nil {
		h.logger.Fields(logger.Fields{
			"address": address,
		}).Error(err, "Can't convert input to hash address")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	var req request.GetNftTokensRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		h.logger.Fields(logger.Fields{
			"address": address,
			"req":     req,
		}).Error(err, "Can't bind query params to struct request.GetNftTokensQuery")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, req))
		return
	}

	tokens, total, err := h.entities.Nft.GetNftTokens(address, req)
	if err != nil {
		h.logger.Fields(logger.Fields{
			"address": address,
			"request": req,
		}).Error(err, "Can't get NFT tokens")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(tokens, &response.PaginationResponse{
		Pagination: model.Pagination{
			Page: req.Page,
			Size: req.Size,
		},
		Total: total,
	}, nil, nil))
}

// GetNftTokensByWalletAddress godoc
// @Summary     Get nft tokens by owner address
// @Description Get nft tokens by owner address
// @Tags        Wallet
// @Accept      json
// @Produce     json
// @Param       wallet_address path string true "wallet address"
// @Param       parameters              query     request.GetNftTokensByAddressRequest false "get tokens by wallet address query"
// @Success     200 {object} response.GetNftTokensResponse
// @Error       400 {object} response.ErrorResponse
// @Error       404 {object} response.ErrorResponse
// @Error       500 {object} response.ErrorResponse
// @Router      /{wallet_address}/nft [get]
func (h *Handler) GetNftTokensByWalletAddress(c *gin.Context) {
	var params struct {
		Address string `uri:"wallet_address" binding:"address" msg:"wallet_address is required"`
	}

	if err := c.ShouldBindUri(&params); err != nil {
		h.logger.Fields(logger.Fields{
			"address": params.Address,
		}).Error(err, "No wallet_address passed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, params))
		return
	}

	walletAddress, err := utils.StringToHashAddress(params.Address)
	if err != nil {
		h.logger.Fields(logger.Fields{
			"address": walletAddress,
		}).Error(err, "Can't convert input to hash address")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	var req request.GetNftTokensByAddressRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		h.logger.Fields(logger.Fields{
			"address": walletAddress,
		}).Error(err, "Can't bind query params to struct request.GetNftTokensByAddressQuery")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, req))
		return
	}

	tokens, total, err := h.entities.Nft.GetNftTokensByWalletAddress(walletAddress, req)
	if err != nil {
		h.logger.Fields(logger.Fields{
			"address": walletAddress,
			"request": req,
		}).Error(err, "Can't get NFT tokens by wallet address")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(tokens, &response.PaginationResponse{
		Pagination: model.Pagination{
			Page: req.Page,
			Size: req.Size,
		},
		Total: total,
	}, nil, nil))
}

// GetNftQueryOptions godoc
// @Summary			Get NFT avalable query options (attributes,marketplaces,...)
// @Description	Get NFT avalable query options (attributes,marketplaces,...)
// @Tags    		NFT
// @Accept  		json
// @Produce 		json
// @Param   		collection_address path string true "collection address"
// @Success			200 {object} response.GetNftQueryOptionsResponse
// @Error       400 {object} response.ErrorResponse
// @Error       500 {object} response.ErrorResponse
// @Router			/nft/{collection_address}/query-options [get]
func (h *Handler) GetNftQueryOptions(c *gin.Context) {
	var params struct {
		Address string `uri:"collection_address" binding:"address" msg:"invalid address"`
	}

	if err := c.ShouldBindUri(&params); err != nil {
		h.logger.Fields(logger.Fields{
			"address": params.Address,
		}).Error(err, "[GetNftQueryOptions] unable to parse url")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, params))
		return
	}

	logger := h.logger.Fields(logger.Fields{
		"address": params.Address,
	})

	collectionAddress, err := utils.StringToHashAddress(params.Address)
	if err != nil {
		logger.Error(err, "Can't convert input to hash address")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	marketplaces, attrs, traitCounts, err := h.entities.Nft.GetQueryOptions(collectionAddress)
	if utils.IsRecordNotFound(err) {
		logger.Error(err, "NFT collection not found")
		c.JSON(http.StatusNotFound, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	if err != nil {
		logger.Error(err, "Can't get NFT query options")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	data := response.GetNftQueryOptionsData{
		Marketplaces: marketplaces,
		Attributes:   attrs,
		TraitCounts:  traitCounts,
	}

	c.JSON(http.StatusOK, response.CreateResponse(data, nil, nil, nil))
}

// ResyncCollection   godoc
// @Summary      ResyncCollection
// @Description  ResyncCollection
// @Tags         NFT
// @Accept       json
// @Produce      json
// @Param        collection_address   path      string  true  "collection address"
// @Success      200 {object} response.ResponseString
// @Error        400 {object} response.ErrorResponse
// @Error        500 {object} response.ErrorResponse
// @Router       /nft/{collection_address}/resync [put]
func (h *Handler) ResyncNftCollection(c *gin.Context) {
	var params struct {
		Address string `uri:"collection_address" binding:"address" msg:"invalid address"`
	}

	if err := c.ShouldBindUri(&params); err != nil {
		h.logger.Fields(logger.Fields{
			"address": params.Address,
		}).Error(err, "[ResyncNftCollection] unable to parse url")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, params))
		return
	}

	logger := h.logger.Fields(logger.Fields{
		"address": params.Address,
	})

	err := h.entities.Nft.ResyncNftCollection(params.Address)
	if err != nil {
		logger.Error(err, "Can't resyn nft collection")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
	}
	c.JSON(http.StatusOK, response.CreateResponse("ok", nil, nil, nil))
}

// ResyncCollectionToken   godoc
// @Summary      ResyncCollectionToken
// @Description  ResyncCollectionToken
// @Tags         NFT
// @Accept       json
// @Produce      json
// @Param        collection_address   path   string  true  "collection address"
// @Param        token_id             path   string  true  "token id"
// @Success      200 {object} response.ResponseString
// @Error        400 {object} response.ErrorResponse
// @Error        500 {object} response.ErrorResponse
// @Router       /nft/{collection_address}/{token_id}/resync [put]
func (h *Handler) ResyncNftToken(c *gin.Context) {
	var params struct {
		Address string `uri:"collection_address" binding:"address" msg:"invalid address"`
		TokenID string `uri:"token_id" binding:"required" msg:"token_id is required"`
	}

	if err := c.ShouldBindUri(&params); err != nil {
		h.logger.Fields(logger.Fields{
			"address": params.Address,
			"id":      params.TokenID,
		}).Error(err, "[ResyncNftToken] unable to parse url")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, params))
		return
	}

	logger := h.logger.Fields(logger.Fields{
		"address": params.Address,
		"id":      params.TokenID,
	})

	collection, err := h.entities.Nft.GetCollectionByAddress(params.Address)
	if err != nil {
		logger.Error(err, "Can't get nft collection")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
	}

	msg := &message.EnrichNftTokenMessage{
		NftId:             params.TokenID,
		CollectionAddress: params.Address,
		ChainId:           collection.ChainId,
	}

	err = h.queue.Enqueue(message.KafkaMessage{
		EnrichNftToken: msg,
	})
	if err != nil {
		logger.Error(err, "Can't add nft collection to queue")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
	}

	c.JSON(http.StatusOK, response.CreateResponse("ok", nil, nil, nil))
}

// GetNftMetadataAttributesIcon godoc
// @Summary      Get nft trait icons
// @Description  Get nft trait icons
// @Tags         NFT
// @Accept       json
// @Produce      json
// @Param        parameters query request.GetNftAttributeIconQuery false "filter icon by traits"
// @Success      200 {object} response.NftMetadataAttributesIconResponse
// @Error        400 {object} response.ErrorResponse
// @Error        500 {object} response.ErrorResponse
// @Router       /nft/metadata/attributes-icon [get]
func (h *Handler) GetNftMetadataAttributesIcon(c *gin.Context) {
	var req request.GetNftAttributeIconQuery
	if err := c.ShouldBindQuery(&req); err != nil {
		h.logger.Fields(logger.Fields{
			"query": req,
		}).Error(err, "Can't bind query params to struct request.GetNftMetadataAttributesIcon")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, req))
		return
	}

	resp, err := h.entities.Nft.GetNftMetadataAttributesIcon(req)
	if err != nil {
		h.logger.Fields(logger.Fields{
			"query": req,
		}).Error(err, "Can't get NFT metadata attributes icon")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, resp)
}

// ResyncNftCollectionMetadata godoc
// @Summary      resync collection metadata
// @Description  resync collection metadata
// @Tags         NFT
// @Accept       json
// @Produce      json
// @Param        collection_address   path   string  true  "collection address"
// @Param        chain-id	query		string		true		"1 / 10 / 250"
// @Success      200 {object} response.ResponseString
// @Error        400 {object} response.ErrorResponse
// @Error        500 {object} response.ErrorResponse
// @Router       /nft/{collection_address}/metadata/resync [put]
func (h *Handler) ResyncNftCollectionMetadata(c *gin.Context) {
	var params struct {
		Address string `uri:"collection_address" binding:"address" msg:"invalid address"`
	}
	log := h.logger
	if err := c.ShouldBindUri(&params); err != nil {
		log.AddField("address", params.Address).Error(err, "Invalid collection address")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, params))
		return
	}
	log.AddField("address", params.Address)

	var query struct {
		ChainID int64 `form:"chain-id" binding:"required"`
	}
	if err := c.ShouldBindQuery(&query); err != nil {
		log.AddField("query", query).Error(err, "[ResyncNftCollectionMetadata] - invalid request")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, query))
		return
	}
	log.AddField("query", query)

	err := h.entities.Nft.ResyncNftCollectionMetadata(params.Address, query.ChainID)
	if err != nil {
		log.Error(err, "Can't enrich old collection")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse("ok", nil, nil, nil))
}

// ResyncNftCollectionRarity godoc
// @Summary      resync collection rarity
// @Description  resync collection rarity
// @Tags         NFT
// @Accept       json
// @Produce      json
// @Param        collection_address   path   string  true  "collection address"
// @Success      200 {object} response.ResponseString
// @Error        400 {object} response.ErrorResponse
// @Error        500 {object} response.ErrorResponse
// @Router       /nft/{collection_address}/rarity/resync [put]
func (h *Handler) ResyncNftCollectionRarity(c *gin.Context) {
	var params struct {
		Address string `uri:"collection_address" binding:"address" msg:"invalid address"`
	}

	if err := c.ShouldBindUri(&params); err != nil {
		h.logger.AddField("address", params.Address).Error(err, "[ResyncNftCollectionRarity] unable to parse url")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, params))
		return
	}

	err := h.entities.Nft.ResyncNftCollectionRarity(params.Address)
	if err != nil {
		h.logger.Fields(logger.Fields{
			"address": params,
		}).Error(err, "Failed to resync rarity collection")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse("ok", nil, nil, nil))
}

// GetNftCollectionMetadata godoc
// @Summary      Get nft collection metadata
// @Description  Get nft collection metadata
// @Tags         NFT
// @Accept       json
// @Produce      json
// @Param        collection_address   path   string  true  "collection address"
// @Success      200 {object} response.GetNftCollectionMetadataResponse
// @Error        400 {object} response.ErrorResponse
// @Error        500 {object} response.ErrorResponse
// @Router       /nft/{collection_address}/metadata [get]
func (h *Handler) GetNftCollectionMetadata(c *gin.Context) {
	var params struct {
		Address string `uri:"collection_address" binding:"address" msg:"invalid address"`
	}

	log := h.logger.AddField("address", params.Address)
	if err := c.ShouldBindUri(&params); err != nil {
		log.Error(err, "Time duration is not valid")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, params))
		return
	}

	collectionAddress, err := utils.StringToHashAddress(params.Address)
	if err != nil {
		log.Error(err, "Can't convert input to hash address")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	res, err := h.entities.Nft.GetCollectionMetadata(collectionAddress)
	if utils.IsRecordNotFound(err) {
		log.Error(err, "NFT collection not found")
		c.JSON(http.StatusNotFound, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	if err != nil {
		log.Error(err, "[GetNftCollectionMetadata] GetCollectionMetadata failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(res, nil, nil, nil))
}

// GetNftTokenActivities godoc
// @Summary      Get nft token activities
// @Description  Get nft token activities
// @Tags         NFT
// @Accept       json
// @Produce      json
// @Param        token_id           	path      int                                  true  "token Id"
// @Param        collection_address   path      string                               true  "collection address"
// @Param        parameters           query     request.GetNftTokenActivitiesRequest false "query"
// @Success      200 {object} response.GetNftTokenActivitiesResponse
// @Error        400 {object} response.ErrorResponse
// @Error        500 {object} response.ErrorResponse
// @Router       /nft/{collection_address}/{token_id}/activity [get]
func (h *Handler) GetNftTokenActivities(c *gin.Context) {
	var params struct {
		Address string `uri:"collection_address" binding:"address" msg:"invalid address"`
		TokenID string `uri:"token_id" binding:"required" msg:"token_id is required"`
	}

	if err := c.ShouldBindUri(&params); err != nil {
		h.logger.Fields(logger.Fields{
			"address": params.Address,
			"id":      params.TokenID,
		}).Error(err, "[GetNftTokenActivities] unable to parse url")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, params))
		return
	}

	log := h.logger.Fields(logger.Fields{
		"address": params.Address,
		"id":      params.TokenID,
	})

	address, err := utils.StringToHashAddress(params.Address)
	if err != nil {
		log.Error(err, "Can't convert input to hash address")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	// query params
	var req request.GetNftTokenActivitiesRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		h.logger.Fields(logger.Fields{
			"pagination": req,
		}).Error(err, "Can't bind query params to request.GetNftTokenActivitiesRequest")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, req))
		return
	}

	activities, total, err := h.entities.Nft.GetTokenActivities(address, params.TokenID, req)
	if err != nil {
		log.Error(err, "Can't get token activities")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(activities, &response.PaginationResponse{
		Pagination: model.Pagination{
			Page: req.Page,
			Size: req.Size,
		},
		Total: total,
	}, nil, nil))
}

// GetNftTokenTransactionHistory godoc
// @Summary      Get nft token tx history
// @Description  Get nft token tx history
// @Tags         NFT
// @Accept       json
// @Produce      json
// @Param        token_id           	path      int                                  true  "token Id"
// @Param        collection_address   path      string                               true  "collection address"
// @Success      200 {object} response.GetNftTokenTransactionHistory
// @Error        400 {object} response.ErrorResponse
// @Error        500 {object} response.ErrorResponse
// @Router       /nft/{collection_address}/{token_id}/transaction-history [get]
func (h *Handler) GetNftTokenTransactionHistory(c *gin.Context) {
	var params struct {
		Address string `uri:"collection_address" binding:"required"`
		TokenID string `uri:"token_id" binding:"required" msg:"token_id is required"`
	}

	if err := c.ShouldBindUri(&params); err != nil {
		h.logger.Fields(logger.Fields{
			"address": params.Address,
			"id":      params.TokenID,
		}).Error(err, "[GetNftTokenActivities] unable to parse url")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, params))
		return
	}

	txHistory, err := h.entities.Nft.GetNftTokenTransactionHistory(params.Address, params.TokenID)
	if err != nil {
		h.logger.Fields(logger.Fields{
			"address": params.Address,
			"id":      params.TokenID,
		}).Error(err, "[GetNftTokenTransactionHistory] cannot get tx history")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(txHistory, nil, nil, nil))
}

// GetNftTokenBidHistory godoc
// @Summary      Get nft token bid history
// @Description  Get nft token bid history
// @Tags         NFT
// @Accept       json
// @Produce      json
// @Param        token_id           	path      int              true  "token Id"
// @Param        collection_address   path      string           true  "collection address"
// @Param        parameters           query     model.Pagination false "pagination"
// @Success      200 {object} response.GetNftTokenBidHistoryResponse
// @Error        400 {object} response.ErrorResponse
// @Error        500 {object} response.ErrorResponse
// @Router       /nft/{collection_address}/{token_id}/bid-history [get]
func (h *Handler) GetNftTokenBidHistory(c *gin.Context) {
	var params struct {
		Address string `uri:"collection_address" binding:"address" msg:"invalid address"`
		TokenID string `uri:"token_id" binding:"required" msg:"token_id is required"`
	}

	if err := c.ShouldBindUri(&params); err != nil {
		h.logger.Fields(logger.Fields{
			"address": params.Address,
			"id":      params.TokenID,
		}).Error(err, "[GetNftTokenBidHistory] unable to parse url")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, params))
		return
	}

	log := h.logger.Fields(logger.Fields{
		"address": params.Address,
		"id":      params.TokenID,
	})

	address, err := utils.StringToHashAddress(params.Address)
	if err != nil {
		log.Error(err, "Can't convert input to hash address")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	// query params
	var req model.Pagination
	if err := c.ShouldBindQuery(&req); err != nil {
		log.AddField("pagination", req).Error(err, "Can't bind query params to request.GetNftTokenActivitiesRequest")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, req))
		return
	}

	bids, total, err := h.entities.Nft.GetTokenBidHistory(address, params.TokenID, &req)
	if err != nil {
		log.Error(err, "Can't get token activities")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(bids, &response.PaginationResponse{
		Pagination: model.Pagination{
			Page: req.Page,
			Size: req.Size,
		},
		Total: total,
	}, nil, nil))
}

// GetNftTokenMetadata godoc
// @Summary      Get nft token metadata
// @Description  Get nft token metadata
// @Tags         NFT
// @Accept       json
// @Produce      json
// @Param        token_id           	path      int  		true  "token Id"
// @Param        collection_address   path      string  true  "collection address"
// @Success      200 {object} response.GetNftTokenMetadataResponse
// @Error        400 {object} response.ErrorResponse
// @Error        500 {object} response.ErrorResponse
// @Router       /nft/{collection_address}/{token_id}/metadata [get]
func (h *Handler) GetNftTokenMetadata(c *gin.Context) {
	var params struct {
		Address string `uri:"collection_address" binding:"address" msg:"invalid address"`
		TokenID string `uri:"token_id" binding:"required" msg:"token_id is required"`
	}

	if err := c.ShouldBindUri(&params); err != nil {
		h.logger.Fields(logger.Fields{
			"address": params.Address,
			"id":      params.TokenID,
		}).Error(err, "[GetNftTokenListingSoldPrices] unable to parse url")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, params))
		return
	}

	l := h.logger.Fields(logger.Fields{
		"address": params.Address,
		"id":      params.TokenID,
	})

	// validate request
	address, err := utils.StringToHashAddress(params.Address)
	if err != nil {
		l.Error(err, "Can't convert input to hash address")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	// entity
	res, err := h.entities.Nft.GetTokenMetadata(address, params.TokenID)
	if utils.IsRecordNotFound(err) {
		l.Error(err, "NFT token not found")
		c.JSON(http.StatusNotFound, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	if err != nil {
		l.Error(err, "Can't get NFT token metadata")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetNftTokenListingSoldPrices godoc
// @Summary      Get nft prices
// @Description  Get nft prices for chart
// @Tags         NFT
// @Accept       json
// @Produce      json
// @Param        token_id           	path      int  		true  "token Id"
// @Param        collection_address   path      string  true  "collection address"
// @Param        currency			  query      string  false  "payment token/currency" Enums(eth,ftm)
// @Param        time_duration        query			string	false		"time duration one of 24h (default) 7d 30d all"
// @Param        time_interval        query			string	false		"time interval one of 1h (default) 1d 1M"
// @Param        time_round_up        query			string	false		"time round up one of h (default) d M"
// @Success      200 {object} response.GetNftTokenListingSoldPricesResponse
// @Error        400 {object} response.ErrorResponse
// @Error        500 {object} response.ErrorResponse
// @Router       /nft/{collection_address}/{token_id}/sold-prices [get]
func (h *Handler) GetNftTokenListingSoldPrices(c *gin.Context) {
	var params struct {
		Address string `uri:"collection_address" binding:"address" msg:"invalid address"`
		TokenID string `uri:"token_id" binding:"required" msg:"token_id is required"`
	}

	if err := c.ShouldBindUri(&params); err != nil {
		h.logger.Fields(logger.Fields{
			"address": params.Address,
			"id":      params.TokenID,
		}).Error(err, "[GetNftTokenListingSoldPrices] unable to parse url")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, params))
		return
	}

	logger := h.logger.Fields(logger.Fields{
		"address": params.Address,
		"id":      params.TokenID,
	})

	var query struct {
		TimeDuration string `form:"time_duration" binding:"omitempty,oneof=24h 7d 30d all" enums:"24h 7d 30d all"`
		TimeInterval string `form:"time_interval" binding:"omitempty,oneof=1h 1d 1M" enums:"1h 1d 1M"`
		TimeRoundUp  string `form:"time_round_up" binding:"omitempty,oneof=d h M" enums:"d h M"`
		Currency     string `form:"currency" binding:"omitempty,oneof=eth ftm" enums:"eth,ftm"` // payment token/currency
	}

	if err := c.ShouldBindQuery(&query); err != nil {
		logger.AddField("query", query).Error(err, "Unable to parse query")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, query))
		return
	}

	var timeDuration constant.TimeDuration = constant.TimeDuration24Hours
	if duration, ok := constant.StringToTimeDuration(query.TimeDuration); ok {
		timeDuration = duration
	}

	var timeInterval constant.TimeInterval = constant.TimeInterval1Hour
	if interval, ok := constant.StringToTimeInterval(query.TimeInterval); ok {
		timeInterval = interval
	}

	var timeRoundUp constant.TimeRoundUp = constant.TimeRoundUpHours
	if roundUp, ok := constant.StringToTimeRoundUp(query.TimeRoundUp); ok {
		timeRoundUp = roundUp
	}

	var paymentCurrency constant.Currency = constant.ETH
	if currency, ok := constant.StringToCurrency(query.Currency); ok {
		paymentCurrency = currency
	}

	address, err := utils.StringToHashAddress(params.Address)
	if err != nil {
		logger.Error(err, "Can't convert input to hash address")
	}

	prices, err := h.entities.Nft.GetNftListingSoldPrices(address, params.TokenID, paymentCurrency.String(), timeDuration.String(), timeInterval.String(), timeRoundUp.String())
	if err != nil {
		logger.Error(err, "Can't get token prices")
		code := http.StatusInternalServerError
		if err == gorm.ErrRecordNotFound {
			code = http.StatusNotFound
		}
		c.JSON(code, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(prices, nil, nil, nil))
}

// GetNftTokenListingAvgPricesAndFloor godoc
// @Summary      Get nft avg prices
// @Description  Get nft avg prices and floor price for chart
// @Tags         NFT
// @Accept       json
// @Produce      json
// @Param        token_id           	path      int  		true  "token Id"
// @Param        collection_address   path      string  true  "collection address"
// @Param        currency			  query      string  false  "payment token/currency" Enums(eth,ftm)
// @Param        time_duration        query			string	false		"time duration one of 24h (default) 7d 30d all"
// @Param        time_interval        query			string	false		"time interval one of 1h (default) 1d 1M"
// @Param        time_round_up        query			string	false		"time round up one of h (default) d M"
// @Success      200 {object} response.GetNftTokenListingAvgPricesAndFloorResponse
// @Error        400 {object} response.ErrorResponse
// @Error        500 {object} response.ErrorResponse
// @Router       /nft/{collection_address}/{token_id}/avg-prices-and-floor [get]
func (h *Handler) GetNftTokenListingAvgPricesAndFloor(c *gin.Context) {
	var params struct {
		Address string `uri:"collection_address" binding:"address" msg:"invalid address"`
		TokenID string `uri:"token_id" binding:"required" msg:"token_id is required"`
	}

	if err := c.ShouldBindUri(&params); err != nil {
		h.logger.Fields(logger.Fields{
			"address": params.Address,
			"id":      params.TokenID,
		}).Error(err, "[GetNftTokenListingSoldPrices] unable to parse url")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, params))
		return
	}

	logger := h.logger.Fields(logger.Fields{
		"address": params.Address,
		"id":      params.TokenID,
	})

	var query struct {
		TimeDuration string `form:"time_duration" binding:"omitempty,oneof=24h 7d 30d all" enums:"24h 7d 30d all"`
		TimeInterval string `form:"time_interval" binding:"omitempty,oneof=1h 1d 1M" enums:"1h 1d 1M"`
		TimeRoundUp  string `form:"time_round_up" binding:"omitempty,oneof=d h M" enums:"d h M"`
		Currency     string `form:"currency" binding:"omitempty,oneof=eth ftm" enums:"eth,ftm"` // payment token/currency
	}

	if err := c.ShouldBindQuery(&query); err != nil {
		logger.AddField("query", query).Error(err, "Unable to parse query")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, query))
		return
	}

	var timeDuration constant.TimeDuration = constant.TimeDuration24Hours
	if duration, ok := constant.StringToTimeDuration(query.TimeDuration); ok {
		timeDuration = duration
	}

	var timeInterval constant.TimeInterval = constant.TimeInterval1Hour
	if interval, ok := constant.StringToTimeInterval(query.TimeInterval); ok {
		timeInterval = interval
	}

	var timeRoundUp constant.TimeRoundUp = constant.TimeRoundUpHours
	if roundUp, ok := constant.StringToTimeRoundUp(query.TimeRoundUp); ok {
		timeRoundUp = roundUp
	}

	var paymentCurrency constant.Currency = constant.ETH
	if currency, ok := constant.StringToCurrency(query.Currency); ok {
		paymentCurrency = currency
	}

	address, err := utils.StringToHashAddress(params.Address)
	if err != nil {
		logger.Error(err, "Can't convert input to hash address")
	}

	prices, err := h.entities.Nft.GetNftListingAvgPricesAndFloor(address, params.TokenID, paymentCurrency.String(), timeDuration.String(), timeInterval.String(), timeRoundUp.String())
	if err != nil {
		logger.Error(err, "Can't get token prices")
		code := http.StatusInternalServerError
		if err == gorm.ErrRecordNotFound {
			code = http.StatusNotFound
		}
		c.JSON(code, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(prices, nil, nil, nil))
}

// SearchNft godoc
// @Summary      Search nft globally
// @Description  Search nft globally
// @Tags         NFT
// @Accept       json
// @Produce      json
// @Param				 parameters	query		request.SearchNftRequest	true	"filter"
// @Success      200 {object} response.SearchNftResponse
// @Error        400 {object} response.ErrorResponse
// @Error        500 {object} response.ErrorResponse
// @Router       /nft/search [get]
func (h *Handler) SearchNft(c *gin.Context) {
	var req request.SearchNftRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		h.logger.Fields(logger.Fields{"body": req}).Error(err, "[SearchNft] BindJSON failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, req))
		return
	}

	collections, assets, err := h.entities.Nft.SearchNft(&req)
	if err != nil {
		h.logger.Error(err, "entity.SearchNft failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.SearchNftResponseData{
		Collections: collections,
		Assets:      assets,
	}, nil, nil, nil))
}

// GetNftCollectionSaleVolumeAndFloorPrice godoc
// @Summary      Get nft collection sale volume and floor price
// @Description  Get nft collection sale volume and floor price
// @Tags         NFT
// @Accept       json
// @Produce      json
// @Param        collection_address   path      string  true  "collection address"
// @Param        time_duration        query			string	false		"time duration one of 7d (default) 30d 3M 12M"
// @Success      200 {object} response.GetNftCollectionSaleVolumeAndFloorPrice
// @Error        400 {object} response.ErrorResponse
// @Error        500 {object} response.ErrorResponse
// @Router       /nft/{collection_address}/sale-volume-and-floor-prices [get]
func (h *Handler) GetNftCollectionSaleVolumeAndFloorPrice(c *gin.Context) {
	var params struct {
		Address string `uri:"collection_address" binding:"address" msg:"invalid address"`
	}

	log := h.logger
	if err := c.ShouldBindUri(&params); err != nil {
		log.AddField("address", params.Address).Error(err, "[GetNftTokenListingSoldPrices] unable to parse url")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, params))
		return
	}
	log.AddField("address", params.Address)

	var query struct {
		TimeDuration string `form:"time_duration" binding:"omitempty,oneof=7d 30d 3M 12M" enums:"7d,30d,3M,12M"`
	}

	if err := c.ShouldBindQuery(&query); err != nil {
		log.AddField("query", query).Error(err, "Time duration is not valid")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, query))
		return
	}
	log.AddField("query", query)

	var timeDuration constant.SaleVolumeAndFloorPriceTimeDuration = constant.SaleVolumeAndFloorPriceTimeDuration7Days
	if duration, ok := constant.StringToSaleVolumeAndFloorPriceTimeDuration(query.TimeDuration); ok {
		timeDuration = duration
	}

	address, err := utils.StringToHashAddress(params.Address)
	if err != nil {
		log.Error(err, "Can't convert input to hash address")
	}

	saleVolumeAndFloorPrice, err := h.entities.Nft.GetCollectionSaleVolumeAndFloorPrice(address, timeDuration.String())
	if err != nil {
		log.Error(err, "Can't get collection sale volume and floor price")
		code := http.StatusInternalServerError
		if err == gorm.ErrRecordNotFound {
			code = http.StatusNotFound
		}
		c.JSON(code, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	res := &response.GetNftCollectionSaleVolumeAndFloorPrice{
		Data: saleVolumeAndFloorPrice,
	}

	c.JSON(http.StatusOK, res)
}

// GetNftListingAcrossPriceRanges godoc
// @Summary      Get nft listing across price ranges
// @Description  Get nft listing across price ranges
// @Tags         NFT
// @Accept       json
// @Produce      json
// @Param        collection_address   path      string  true   "collection address"
// @Param        payment_token        path      string  true   "payment token"
// @Param        price_range	        query     string	false  "price range, regex \d{12,20} 1000000000000000000(10^18) as default"
// @Success      200 {object} response.GetNftListingAcrossPriceRangesResponse
// @Error        400 {object} response.ErrorResponse
// @Error        500 {object} response.ErrorResponse
// @Router       /nft/{collection_address}/listing-across-price-ranges/payment_token [get]
func (h *Handler) GetNftListingAcrossPriceRanges(c *gin.Context) {
	var params struct {
		Address      string `uri:"collection_address" binding:"address" msg:"invalid address"`
		PaymentToken string `uri:"payment_token" binding:"required,min=3" msg:"payment token is required"`
	}

	log := h.logger
	if err := c.ShouldBindUri(&params); err != nil {
		log.AddField("address", params.Address).Error(err, "[GetNftListingAcrossPriceRanges] unable to parse url")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, params))
		return
	}
	log.AddField("address", params.Address)

	var query struct {
		PriceRange string `form:"price_range"`
	}

	if err := c.ShouldBindQuery(&query); err != nil {
		log.AddField("query", query).Error(err, "Unable to parse query")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, query))
		return
	}

	priceRange := strings.TrimSpace(query.PriceRange)
	if priceRange == "" {
		priceRange = "1000000000000000000"
	} else if match, _ := regexp.MatchString("\\d{12,20}", priceRange); !match {
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("[GetNftListingAcrossPriceRanges] wrong format for price-range"), query))
		return
	}

	address, err := utils.StringToHashAddress(params.Address)
	if err != nil {
		log.Error(err, "Can't convert input to hash address")
	}

	listingAcrossPriceRanges, err := h.entities.Nft.GetListingAcrossPriceRanges(address, params.PaymentToken, priceRange)
	if err != nil {
		log.Error(err, "Can't get collection sale volume and floor price")
		code := http.StatusInternalServerError
		if err == gorm.ErrRecordNotFound {
			code = http.StatusNotFound
		}
		c.JSON(code, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(listingAcrossPriceRanges, nil, nil, nil))
}

// GetListingAcrossRarityRankRanges godoc
// @Summary      Get nft listing across rarity rank ranges
// @Description  Get nft listing across rarity rank ranges
// @Tags         NFT
// @Accept       json
// @Produce      json
// @Param        collection_address   path      string  true  "collection address"
// @Param        payment_token   path      string  true  "payment token"
// @Param		 rarity_rank_range	query		string		false		"rarity rank range, regex \d{1,10} with 200 as default"
// @Success      200  {object}  response.GetListingAcrossRarityRankRangesResponse
// @Router       /nft/{collection_address}/listing-across-rarity-rank-ranges/payment_token [get]
func (h *Handler) GetListingAcrossRarityRankRanges(c *gin.Context) {
	var params struct {
		Address      string `uri:"collection_address" binding:"address" msg:"invalid address"`
		PaymentToken string `uri:"payment_token" binding:"required,min=3" msg:"payment token is required"`
	}

	log := h.logger
	if err := c.ShouldBindUri(&params); err != nil {
		log.AddField("address", params.Address).Error(err, "[GetListingAcrossRarityRankRanges] unable to parse url")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, params))
		return
	}
	log.AddField("address", params.Address)

	var query struct {
		RarityRankRange string `form:"rarity_rank_range"`
	}

	if err := c.ShouldBindQuery(&query); err != nil {
		log.AddField("query", query).Error(err, "Unable to parse query")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, query))
		return
	}

	rarityRankRange := strings.TrimSpace(query.RarityRankRange)
	if rarityRankRange == "" {
		rarityRankRange = "200"
	} else if match, _ := regexp.MatchString("\\d{1,10}", rarityRankRange); !match {
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("[GetListingAcrossRarityRankRanges] wrong format for rarity-rank-range"), query))
		return
	}

	address, err := utils.StringToHashAddress(params.Address)
	if err != nil {
		log.Error(err, "Can't convert input to hash address")
	}

	listingAcrossRarityRankRanges, err := h.entities.Nft.GetListingAcrossRarityRankRanges(address, params.PaymentToken, rarityRankRange)
	if err != nil {
		log.Error(err, "Can't get collection sale volume and floor price")
		code := http.StatusInternalServerError
		if err == gorm.ErrRecordNotFound {
			code = http.StatusNotFound
		}
		c.JSON(code, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	res := &response.GetListingAcrossRarityRankRangesResponse{
		Data: listingAcrossRarityRankRanges,
	}

	c.JSON(http.StatusOK, res)
}

// GetListingMarketCapAndVolumeByTimeRanges godoc
// @Summary      Get nft listing market cap and volume by time ranges
// @Description  Get nft listing market cap and volume by time ranges
// @Tags         NFT
// @Accept       json
// @Produce      json
// @Param        collection_address   path      string  true  "collection address"
// @Param        currency			  query      string  false  "payment token/currency" Enums(eth,ftm)
// @Param        time_duration        query			string	false		"time duration one of 24h (default) 7d 30d all"
// @Param        time_interval        query			string	false		"time interval one of 1h (default) 1d 1M"
// @Param        time_round_up        query			string	false		"time round up one of h (default) d M"
// @Success      200  {object}  response.GetListingMarketCapAndVolumeTimeRangesResponse
// @Router       /nft/{collection_address}/listing-market-cap-and-volume-by-time-ranges [get]
func (h *Handler) GetListingMarketCapAndVolumeByTimeRanges(c *gin.Context) {
	var params struct {
		Address string `uri:"collection_address" binding:"address" msg:"invalid address"`
	}

	log := h.logger
	if err := c.ShouldBindUri(&params); err != nil {
		log.AddField("address", params.Address).Error(err, "[GetListingAcrossRarityRankRanges] unable to parse url")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, params))
		return
	}
	log.AddField("address", params.Address)

	var query struct {
		TimeDuration string `form:"time_duration" binding:"omitempty,oneof=24h 7d 30d all" enums:"24h 7d 30d all"`
		TimeInterval string `form:"time_interval" binding:"omitempty,oneof=1h 1d 1M" enums:"1h 1d 1M"`
		TimeRoundUp  string `form:"time_round_up" binding:"omitempty,oneof=d h M" enums:"d h M"`
		Currency     string `form:"currency" binding:"omitempty,oneof=eth ftm" enums:"eth,ftm"` // payment token/currency
	}

	if err := c.ShouldBindQuery(&query); err != nil {
		log.AddField("query", query).Error(err, "Unable to parse query")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, query))
		return
	}

	var timeDuration constant.TimeDuration = constant.TimeDuration24Hours
	if duration, ok := constant.StringToTimeDuration(query.TimeDuration); ok {
		timeDuration = duration
	}

	var timeInterval constant.TimeInterval = constant.TimeInterval1Hour
	if interval, ok := constant.StringToTimeInterval(query.TimeInterval); ok {
		timeInterval = interval
	}

	var timeRoundUp constant.TimeRoundUp = constant.TimeRoundUpHours
	if roundUp, ok := constant.StringToTimeRoundUp(query.TimeRoundUp); ok {
		timeRoundUp = roundUp
	}

	var paymentCurrency constant.Currency = constant.ETH
	if currency, ok := constant.StringToCurrency(query.Currency); ok {
		paymentCurrency = currency
	}

	address, err := utils.StringToHashAddress(params.Address)
	if err != nil {
		log.Error(err, "Can't convert input to hash address")
	}

	listingMartketCapAndVolumeTimeRanges, err := h.entities.Nft.GetListingMarketCapAndVolumeByTimeRanges(address, paymentCurrency.String(), timeDuration.String(), timeInterval.String(), timeRoundUp.String())
	if err != nil {
		log.Error(err, "Can't get collection market cap and volume")
		code := http.StatusInternalServerError
		if err == gorm.ErrRecordNotFound {
			code = http.StatusNotFound
		}
		c.JSON(code, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	res := &response.GetListingMarketCapAndVolumeTimeRangesResponse{
		Data: listingMartketCapAndVolumeTimeRanges,
	}

	c.JSON(http.StatusOK, res)
}

// GetMarketSnapshotFloorPriceByTimeRanges godoc
// @Summary      Get nft market snapshot floor price by time ranges
// @Description  Get nft market snapshot floor price by time ranges
// @Tags         NFT
// @Accept       json
// @Produce      json
// @Param        collection_address   path      string  true  "collection address"
// @Param        currency			  query      string  false  "payment token/currency" Enums(eth,ftm)
// @Param        time_duration        query			string	false		"time duration one of 24h (default) 7d 30d all"
// @Param        time_interval        query			string	false		"time interval one of 1h (default) 1d 1M"
// @Param        time_round_up        query			string	false		"time round up one of h (default) d M"
// @Success      200  {object}  response.GetMarketSnapshotFloorPriceByTimeRangesResponse
// @Router       /nft/{collection_address}/market-snapshot-floor-price-by-time-ranges [get]
func (h *Handler) GetMarketSnapshotFloorPriceByTimeRanges(c *gin.Context) {
	var params struct {
		Address string `uri:"collection_address" binding:"address" msg:"invalid address"`
	}

	log := h.logger
	if err := c.ShouldBindUri(&params); err != nil {
		log.AddField("address", params.Address).Error(err, "[GetMarketSnapshotFloorPriceByTimeRanges] unable to parse url")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, params))
		return
	}
	log.AddField("address", params.Address)

	var query struct {
		TimeDuration string `form:"time_duration" binding:"omitempty,oneof=24h 7d 30d all" enums:"24h 7d 30d all"`
		TimeInterval string `form:"time_interval" binding:"omitempty,oneof=1h 1d 1M" enums:"1h 1d 1M"`
		TimeRoundUp  string `form:"time_round_up" binding:"omitempty,oneof=d h M" enums:"d h M"`
		Currency     string `form:"currency" binding:"omitempty,oneof=eth ftm" enums:"eth,ftm"` // payment token/currency
	}

	if err := c.ShouldBindQuery(&query); err != nil {
		log.AddField("query", query).Error(err, "Unable to parse query")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, query))
		return
	}

	var timeDuration constant.TimeDuration = constant.TimeDuration24Hours
	if duration, ok := constant.StringToTimeDuration(query.TimeDuration); ok {
		timeDuration = duration
	}

	var timeInterval constant.TimeInterval = constant.TimeInterval1Hour
	if interval, ok := constant.StringToTimeInterval(query.TimeInterval); ok {
		timeInterval = interval
	}

	var timeRoundUp constant.TimeRoundUp = constant.TimeRoundUpHours
	if roundUp, ok := constant.StringToTimeRoundUp(query.TimeRoundUp); ok {
		timeRoundUp = roundUp
	}

	var paymentCurrency constant.Currency = constant.ETH
	if currency, ok := constant.StringToCurrency(query.Currency); ok {
		paymentCurrency = currency
	}

	address, err := utils.StringToHashAddress(params.Address)
	if err != nil {
		log.Error(err, "Can't convert input to hash address")
	}

	marketplaceCollectionSnapshotFloorPriceTimeRanges, err := h.entities.Nft.GetMarketSnapshotFloorPriceByTimeRanges(address, paymentCurrency.String(), timeDuration.String(), timeInterval.String(), timeRoundUp.String())
	if err != nil {
		log.Error(err, "Can't get market collection snapshot floor price")
		code := http.StatusInternalServerError
		if err == gorm.ErrRecordNotFound {
			code = http.StatusNotFound
		}
		c.JSON(code, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	res := &response.GetMarketSnapshotFloorPriceByTimeRangesResponse{
		Data: marketplaceCollectionSnapshotFloorPriceTimeRanges,
	}

	c.JSON(http.StatusOK, res)
}

// GetMarketSnapshotMarketCapByTimeRanges godoc
// @Summary      Get nft market snapshot market cap by time ranges
// @Description  Get nft market snapshot market cap by time ranges
// @Tags         NFT
// @Accept       json
// @Produce      json
// @Param        collection_address   path      string  true  "collection address"
// @Param        currency			  query      string  false  "payment token/currency" Enums(eth,ftm)
// @Param        time_duration        query			string	false		"time duration one of 24h (default) 7d 30d all"
// @Param        time_interval        query			string	false		"time interval one of 1h (default) 1d 1M"
// @Param        time_round_up        query			string	false		"time round up one of h (default) d M"
// @Success      200  {object}  response.GetMarketSnapshotMarketCapByTimeRangesResponse
// @Router       /nft/{collection_address}/market-snapshot-market-cap-by-time-ranges [get]
func (h *Handler) GetMarketSnapshotMarketCapByTimeRanges(c *gin.Context) {
	var params struct {
		Address string `uri:"collection_address" binding:"address" msg:"invalid address"`
	}

	log := h.logger
	if err := c.ShouldBindUri(&params); err != nil {
		log.AddField("address", params.Address).Error(err, "[GetMarketSnapshotMarketCapByTimeRanges] unable to parse url")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, params))
		return
	}
	log.AddField("address", params.Address)

	var query struct {
		TimeDuration string `form:"time_duration" binding:"omitempty,oneof=24h 7d 30d all" enums:"24h 7d 30d all"`
		TimeInterval string `form:"time_interval" binding:"omitempty,oneof=1h 1d 1M" enums:"1h 1d 1M"`
		TimeRoundUp  string `form:"time_round_up" binding:"omitempty,oneof=d h M" enums:"d h M"`
		Currency     string `form:"currency" binding:"omitempty,oneof=eth ftm" enums:"eth,ftm"` // payment token/currency
	}

	if err := c.ShouldBindQuery(&query); err != nil {
		log.AddField("query", query).Error(err, "Unable to parse query")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, query))
		return
	}

	var timeDuration constant.TimeDuration = constant.TimeDuration24Hours
	if duration, ok := constant.StringToTimeDuration(query.TimeDuration); ok {
		timeDuration = duration
	}

	var timeInterval constant.TimeInterval = constant.TimeInterval1Hour
	if interval, ok := constant.StringToTimeInterval(query.TimeInterval); ok {
		timeInterval = interval
	}

	var timeRoundUp constant.TimeRoundUp = constant.TimeRoundUpHours
	if roundUp, ok := constant.StringToTimeRoundUp(query.TimeRoundUp); ok {
		timeRoundUp = roundUp
	}

	var paymentCurrency constant.Currency = constant.ETH
	if currency, ok := constant.StringToCurrency(query.Currency); ok {
		paymentCurrency = currency
	}

	address, err := utils.StringToHashAddress(params.Address)
	if err != nil {
		log.Error(err, "Can't convert input to hash address")
	}

	marketplaceCollectionSnapshotMarketCapTimeRanges, err := h.entities.Nft.GetMarketSnapshotMarketCapByTimeRanges(address, paymentCurrency.String(), timeDuration.String(), timeInterval.String(), timeRoundUp.String())
	if err != nil {
		log.Error(err, "Can't get market collection snapshot market cap")
		code := http.StatusInternalServerError
		if err == gorm.ErrRecordNotFound {
			code = http.StatusNotFound
		}
		c.JSON(code, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	res := &response.GetMarketSnapshotMarketCapByTimeRangesResponse{
		Data: marketplaceCollectionSnapshotMarketCapTimeRanges,
	}

	c.JSON(http.StatusOK, res)
}

// GetListingTotalVolumeByTimeRanges godoc
// @Summary      Get nft listing total volume by time ranges
// @Description  Get nft listing total volume by time ranges
// @Tags         NFT
// @Accept       json
// @Produce      json
// @Param        collection_address   path      string  true  "collection address"
// @Param        currency			  query      string  false  "payment token/currency" Enums(eth,ftm)
// @Param        time_duration        query			string	false		"time duration one of 24h (default) 7d 30d all"
// @Param        time_interval        query			string	false		"time interval one of 1h (default) 1d 1M"
// @Param        time_round_up        query			string	false		"time round up one of h (default) d M"
// @Success      200  {object}  response.GetListingTotalVolumeByTimeRangesResponse
// @Router       /nft/{collection_address}/listing-total-volume-by-time-ranges [get]
func (h *Handler) GetListingTotalVolumeByTimeRanges(c *gin.Context) {
	var params struct {
		Address string `uri:"collection_address" binding:"address" msg:"invalid address"`
	}

	log := h.logger
	if err := c.ShouldBindUri(&params); err != nil {
		log.AddField("address", params.Address).Error(err, "[GetListingTotalVolumeByTimeRanges] unable to parse url")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, params))
		return
	}
	log.AddField("address", params.Address)

	var query struct {
		TimeDuration string `form:"time_duration" binding:"omitempty,oneof=24h 7d 30d all" enums:"24h 7d 30d all"`
		TimeInterval string `form:"time_interval" binding:"omitempty,oneof=1h 1d 1M" enums:"1h 1d 1M"`
		TimeRoundUp  string `form:"time_round_up" binding:"omitempty,oneof=d h M" enums:"d h M"`
		Currency     string `form:"currency" binding:"omitempty,oneof=eth ftm" enums:"eth,ftm"` // payment token/currency
	}

	if err := c.ShouldBindQuery(&query); err != nil {
		log.AddField("query", query).Error(err, "Unable to parse query")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, query))
		return
	}

	var timeDuration constant.TimeDuration = constant.TimeDuration24Hours
	if duration, ok := constant.StringToTimeDuration(query.TimeDuration); ok {
		timeDuration = duration
	}

	var timeInterval constant.TimeInterval = constant.TimeInterval1Hour
	if interval, ok := constant.StringToTimeInterval(query.TimeInterval); ok {
		timeInterval = interval
	}

	var timeRoundUp constant.TimeRoundUp = constant.TimeRoundUpHours
	if roundUp, ok := constant.StringToTimeRoundUp(query.TimeRoundUp); ok {
		timeRoundUp = roundUp
	}

	var paymentCurrency constant.Currency = constant.ETH
	if currency, ok := constant.StringToCurrency(query.Currency); ok {
		paymentCurrency = currency
	}

	address, err := utils.StringToHashAddress(params.Address)
	if err != nil {
		log.Error(err, "Can't convert input to hash address")
	}

	listingTotalVolumeTimeRanges, err := h.entities.Nft.GetListingTotalVolumeByTimeRanges(address, paymentCurrency.String(), timeDuration.String(), timeInterval.String(), timeRoundUp.String())
	if err != nil {
		log.Error(err, "Can't get listing total volume")
		code := http.StatusInternalServerError
		if err == gorm.ErrRecordNotFound {
			code = http.StatusNotFound
		}
		c.JSON(code, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	res := &response.GetListingTotalVolumeByTimeRangesResponse{
		Data: listingTotalVolumeTimeRanges,
	}

	c.JSON(http.StatusOK, res)
}

// GetNftCollectionsByWalletAddress godoc
// @Summary     Get nft collections by owner address
// @Description Get nft collections by owner address
// @Tags        Wallet
// @Accept      json
// @Produce     json
// @Param       wallet_address path string true "wallet address"
// @Param       parameters   query   request.GetNftCollectionsByWalletAddressRequest false "query"
// @Success     200 {object} response.GetNftCollectionsResponse
// @Error       400 {object} response.ErrorResponse
// @Error       404 {object} response.ErrorResponse
// @Error       500 {object} response.ErrorResponse
// @Router      /{wallet_address}/collection [get]
func (h *Handler) GetNftCollectionsByWalletAddress(c *gin.Context) {
	var params struct {
		Address string `uri:"wallet_address" binding:"address" msg:"wallet_address is required"`
	}

	if err := c.ShouldBindUri(&params); err != nil {
		h.logger.Fields(logger.Fields{
			"address": params.Address,
		}).Error(err, "No wallet_address passed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, params))
		return
	}

	walletAddress, err := utils.StringToHashAddress(params.Address)
	if err != nil {
		h.logger.Fields(logger.Fields{
			"address": walletAddress,
		}).Error(err, "Can't convert input to hash address")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	var req request.GetNftCollectionsByWalletAddressRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		h.logger.Fields(logger.Fields{
			"address": walletAddress,
		}).Error(err, "Can't bind query params to struct request.GetNftCollectionsByWalletAddressRequest")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, req))
		return
	}

	collections, total, err := h.entities.Nft.GetNftCollectionsByWalletAddress(walletAddress, req)
	if err != nil {
		h.logger.Fields(logger.Fields{
			"address": walletAddress,
			"request": req,
		}).Error(err, "entity.GetNftCollectionsByWalletAddress failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(collections, &response.PaginationResponse{
		Pagination: model.Pagination{
			Page: req.Page,
			Size: req.Size,
		},
		Total: total,
	}, nil, nil))
}

// GetListingUniqueOwnersByTimeRange godoc
// @Summary      Get nft listing unique owners by time ranges
// @Description  Get nft listing unique owners by time ranges
// @Tags         NFT
// @Accept       json
// @Produce      json
// @Param        collection_address   path      string  true  "collection address"
// @Param        time_duration        query			string	false		"time duration one of 24h (default) 7d 30d all"
// @Param        time_interval        query			string	false		"time interval one of 1h (default) 1d 1M"
// @Param        time_round_up        query			string	false		"time round up one of h (default) d M"
// @Success      200  {object}  response.GetListingUniqueOwnersByTimeRangeResponse
// @Router       /nft/{collection_address}/listing-unique-owners-time-ranges [get]
func (h *Handler) GetListingUniqueOwnersByTimeRange(c *gin.Context) {
	var params struct {
		Address string `uri:"collection_address" binding:"address" msg:"invalid address"`
	}

	log := h.logger
	if err := c.ShouldBindUri(&params); err != nil {
		log.AddField("address", params.Address).Error(err, "[GetListingAcrossRarityRankRanges] unable to parse url")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, params))
		return
	}
	log.AddField("address", params.Address)

	var query struct {
		TimeDuration string `form:"time_duration" binding:"omitempty,oneof=24h 7d 30d all" enums:"24h 7d 30d all"`
		TimeInterval string `form:"time_interval" binding:"omitempty,oneof=1h 1d 1M" enums:"1h 1d 1M"`
		TimeRoundUp  string `form:"time_round_up" binding:"omitempty,oneof=d h M" enums:"d h M"`
	}

	if err := c.ShouldBindQuery(&query); err != nil {
		log.AddField("query", query).Error(err, "Unable to parse query")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, query))
		return
	}

	var timeDuration constant.TimeDuration = constant.TimeDuration24Hours
	if duration, ok := constant.StringToTimeDuration(query.TimeDuration); ok {
		timeDuration = duration
	}

	var timeInterval constant.TimeInterval = constant.TimeInterval1Hour
	if interval, ok := constant.StringToTimeInterval(query.TimeInterval); ok {
		timeInterval = interval
	}

	var timeRoundUp constant.TimeRoundUp = constant.TimeRoundUpHours
	if roundUp, ok := constant.StringToTimeRoundUp(query.TimeRoundUp); ok {
		timeRoundUp = roundUp
	}

	address, err := utils.StringToHashAddress(params.Address)
	if err != nil {
		log.Error(err, "Can't convert input to hash address")
	}

	listingUniqueOwnersRanges, err := h.entities.Nft.GetListingUniqueOwnersByTimeRange(address, timeDuration.String(), timeInterval.String(), timeRoundUp.String())
	if err != nil {
		log.Error(err, "Can't get collection sale volume and floor price")
		code := http.StatusInternalServerError
		if err == gorm.ErrRecordNotFound {
			code = http.StatusNotFound
		}
		c.JSON(code, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	res := &response.GetListingUniqueOwnersByTimeRangeResponse{
		Data: listingUniqueOwnersRanges,
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) ImportSolanaDataHandler(c *gin.Context) {
	collectionAddress, _ := c.Params.Get("collection_address")
	file, _, _ := c.Request.FormFile("file")
	contents, _ := ioutil.ReadAll(file)

	data := string(contents)
	result := []model.Element{}
	json.Unmarshal([]byte(data), &result)

	for _, itm := range result {
		tokenName := itm.Data.Name
		tokenId := strings.Split(tokenName, "#")[1]
		tokenDescription := itm.MetadataFromUri.Description
		tokenImage := itm.MetadataFromUri.Image
		tokenAttributes := itm.MetadataFromUri.Attributes

		nftToken := &model.NftToken{
			TokenId:           tokenId,
			CreatedTime:       time.Now(),
			LastUpdatedTime:   time.Now(),
			CollectionAddress: collectionAddress,
			Name:              tokenName,
			Description:       tokenDescription,
			Amount:            "1",
			Image:             tokenImage,
		}
		_ = h.store.Nft.UpsertToken(nftToken, false)

		for _, attrItm := range tokenAttributes {
			_ = h.store.Nft.AttributeUpsertOne(model.NftTokenAttribute{
				CollectionAddress: collectionAddress,
				TokenId:           tokenId,
				TraitType:         attrItm.TraitType,
				Value:             attrItm.Value,
			})
		}
	}
	c.JSON(http.StatusOK, "OK")
}
