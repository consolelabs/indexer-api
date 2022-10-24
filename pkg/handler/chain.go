package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/consolelabs/indexer-api/pkg/logger"
	"github.com/consolelabs/indexer-api/pkg/request"
	"github.com/consolelabs/indexer-api/pkg/response"
)

// GetChains godoc
// @Summary      Get supported chains
// @Description  Get supported chains
// @Tags         Chain
// @Accept       json
// @Produce      json
// @Param				 parameters	query request.GetChainsQuery false "get chains query"
// @Success      200  {object}  response.GetChainsResponse
// @Router       /chain [get]
func (h *Handler) GetChains(c *gin.Context) {
	var query request.GetChainsQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		h.logger.Error(err, "Can't bind query params to struct request.GetChainsQuery")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, query))
		return
	}

	collections, err := h.entities.Chain.GetChains(query)
	if err != nil {
		h.logger.Fields(logger.Fields{
			"query": query,
		}).Error(err, "Can't get chains data")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, collections)
}
