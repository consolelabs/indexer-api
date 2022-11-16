package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/consolelabs/indexer-api/pkg/logger"
	"github.com/consolelabs/indexer-api/pkg/model"
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
