package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/consolelabs/indexer-api/pkg/logger"
	"github.com/consolelabs/indexer-api/pkg/request"
	"github.com/consolelabs/indexer-api/pkg/response"
)

func (h *Handler) GetConvertTokenPrice(c *gin.Context) {
	var body request.ConvertTokenPrice
	if err := c.ShouldBindJSON(&body); err != nil {
		h.logger.Error(err, "Can't bind request body to convert token price")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, body))
		return
	}
	convertTokenPrice, err := h.entities.Token.GetConvertTokenPrice(body.Amount, body.From, body.To)
	if err != nil {
		h.logger.Fields(logger.Fields{
			"amount": body.Amount,
			"from":   body.From,
			"To":     body.To,
		}).Error(err, "[GetConvertTokenPrice] get convert token price")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(convertTokenPrice, nil, nil, nil))
}
