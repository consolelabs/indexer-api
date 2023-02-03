package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/consolelabs/indexer-api/pkg/logger"
	"github.com/consolelabs/indexer-api/pkg/model"
	"github.com/consolelabs/indexer-api/pkg/queue/message"
	"github.com/consolelabs/indexer-api/pkg/request"
	"github.com/consolelabs/indexer-api/pkg/response"
	"github.com/consolelabs/indexer-api/pkg/utils"
)

// AddContract   godoc
// @Summary      Add contract
// @Description  Add contract
// @Tags         Contract
// @Accept       json
// @Produce      json
// @Param Request body request.AddContractRequest true "Add contract request"
// @Success      200  {object}  response.ResponseString
// @Router       /contract/erc721 [post]
func (h *Handler) AddErc721ContractHandler(c *gin.Context) {
	var body request.AddContractRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		h.logger.Error(err, "Can't bind request body to addContractRequest")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, body))
		return
	}

	// Validate correct format of contract address, not allow lowercase, currently set chainId = 0 for solana
	// TODO(trkhoi): convert to correct format for solana
	var checksumAddress string
	var err error
	if body.ChainId != 9999 && body.ChainId != 9997 {
		checksumAddress, err = utils.StringToHashAddress(body.Address)
		if err != nil {
			h.logger.Fields(logger.Fields{
				"address": body.Address,
				"chainId": body.ChainId,
			}).Error(err, "Address does not have correct format")
			c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("Address does not have correct format"), nil))
			return
		}
		if body.ChainId != 0 {
			body.Address = checksumAddress
		}
	}

	err = h.entities.Contract.AddContract(model.Contract{
		Address:     body.Address,
		ChainId:     int64(body.ChainId),
		GrpcAddress: "indexer-grpc:80",
		Type:        "ERC721",
	}, body.Name, body.Symbol)
	if err != nil {
		h.logger.Fields(logger.Fields{
			"address": body.Address,
			"chainId": body.ChainId,
		}).Error(err, "Cannot add contract")
		code := http.StatusInternalServerError
		if err.Error() == "address has been added" {
			code = http.StatusConflict
		}
		c.JSON(code, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	syncFullMsg, _ := json.Marshal(message.KafkaMessage{
		Topic:   "sync_full_collection",
		Address: body.Address,
		ChainId: body.ChainId,
	})

	if body.PriorityFlag {
		h.logger.Fields(logger.Fields{"address": body.Address, "chainId": body.ChainId}).Info("enqueue with priority")
		h.queue.PriorityEnqueue(body.ChainId, syncFullMsg)
	} else {
		h.logger.Fields(logger.Fields{"address": body.Address, "chainId": body.ChainId}).Info("enqueue with non priority")
		h.queue.Enqueue(body.ChainId, syncFullMsg)
	}

	// integratedEvMsg, _ := json.Marshal(message.NftEventKafkaMessage{
	// 	Event: "collection_integrated",
	// 	Data: &message.NftEventKafkaMessageData{
	// 		Address: body.Address,
	// 		ChainId: body.ChainId,
	// 	},
	// })
	// h.queue.Enqueue(body.ChainId, integratedEvMsg)

	c.JSON(http.StatusOK, response.CreateResponse("ok", nil, nil, nil))
}

// AddContract   godoc
// @Summary      Get contract
// @Description  Get contract
// @Tags         Contract
// @Accept       json
// @Produce      json
// @Param        address   path      string  true  "contract address"
// @Success      200  {object}  response.ContractResponse
// @Router       /contract/{address} [get]
func (h *Handler) GetContractStatusHandler(c *gin.Context) {
	address := c.Param("address")
	contract, err := h.entities.Contract.GetContract(address)
	if err != nil {
		h.logger.Fields(logger.Fields{
			"address": address,
		}).Error(err, "Cannot get contract")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot get contract"})
		return
	}
	// contract status is always "syncing", to check if contract is synced:
	// - LastUpdatedBlock > 0 -> it has already done with first round sync
	c.JSON(http.StatusOK, response.CreateResponse(response.ContractResponseData{
		ID:               contract.ID,
		LastUpdatedTime:  contract.LastUpdatedTime,
		LastUpdatedBlock: contract.LastUpdatedBlock,
		CreationBlock:    contract.CreationBlock,
		CreatedTime:      contract.CreatedTime,
		Address:          contract.Address,
		ChainId:          contract.ChainId,
		Type:             contract.Type,
		IsProxy:          contract.IsProxy,
		LogicAddress:     contract.LogicAddress,
		Protocol:         contract.Protocol,
		GrpcAddress:      contract.GrpcAddress,
		IsSynced:         contract.LastUpdatedBlock > 0,
	}, nil, nil, nil))
}

func (h *Handler) AddErc721ContractScript(address, name, symbol string, chainId int64, priorityFlag bool) {
	// Validate correct format of contract address, not allow lowercase, currently set chainId = 0 for solana
	// TODO(trkhoi): convert to correct format for solana
	var checksumAddress string
	var err error
	if chainId != 9999 && chainId != 9997 {
		checksumAddress, err = utils.StringToHashAddress(address)
		if err != nil {
			h.logger.Fields(logger.Fields{
				"address": address,
				"chainId": chainId,
			}).Error(err, "Address does not have correct format")
			return
		}
		if chainId != 0 {
			address = checksumAddress
		}
	}

	err = h.entities.Contract.AddContract(model.Contract{
		Address:     address,
		ChainId:     int64(chainId),
		GrpcAddress: "indexer-grpc:80",
		Type:        "ERC721",
	}, name, symbol)
	if err != nil {
		h.logger.Fields(logger.Fields{
			"address": address,
			"chainId": chainId,
		}).Error(err, "Cannot add contract")
		return
	}

	syncFullMsg, _ := json.Marshal(message.KafkaMessage{
		Topic:   "sync_full_collection",
		Address: address,
		ChainId: chainId,
	})

	if priorityFlag {
		h.logger.Fields(logger.Fields{"address": address, "chainId": chainId}).Info("enqueue with priority")
		h.queue.PriorityEnqueue(chainId, syncFullMsg)
	} else {
		h.logger.Fields(logger.Fields{"address": address, "chainId": chainId}).Info("enqueue with non priority")
		h.queue.Enqueue(chainId, syncFullMsg)
	}

}
