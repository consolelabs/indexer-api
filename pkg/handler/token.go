package handler

import (
	"github.com/consolelabs/indexer-api/pkg/logger"
	"github.com/consolelabs/indexer-api/pkg/request"
	"github.com/consolelabs/indexer-api/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetConvertTokenPrice(c *gin.Context) {
	var body request.ConvertTokenPrice
	if err := c.ShouldBindJSON(&body); err != nil {
		h.logger.Error(err, "Can't bind request body to convert token price")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, body))
		return
	}
	convertTokenPrice, err := h.entities.Token.GetConvertTokenPrice(body.Amount, body.FromToken, body.ToToken)
	if err != nil {
		h.logger.Fields(logger.Fields{
			"amount": body.Amount,
			"from":   body.FromToken,
			"To":     body.ToToken,
		}).Error(err, "[GetConvertTokenPrice] get convert token price")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(convertTokenPrice, nil, nil, nil))
}
