package token

import (
	"fmt"
	"time"

	"gorm.io/gorm"

	"github.com/consolelabs/indexer-api/pkg/logger"
	"github.com/consolelabs/indexer-api/pkg/model"
	"github.com/consolelabs/indexer-api/pkg/request"
	"github.com/consolelabs/indexer-api/pkg/service"
	"github.com/consolelabs/indexer-api/pkg/store"
	"github.com/consolelabs/indexer-api/pkg/utils"
)

type tokenEntity struct {
	store   *store.Store
	service *service.Service
}

func New(store *store.Store, service *service.Service) ITokenEntity {
	return &tokenEntity{
		store:   store,
		service: service,
	}
}

func (t *tokenEntity) GetConvertTokenPrice(req request.ConvertTokenPrice) (*model.ConvertTokenPrice, error) {
	fromToken, err := t.store.Token.GetTokenBySymbol(req.From)
	if err != nil {
		logger.L.Fields(logger.Fields{"fromToken": req.From}).Error(err, "failed to get from token")
		return nil, err
	}
	toToken, err := t.store.Token.GetTokenBySymbol(req.To)
	if err != nil {
		logger.L.Fields(logger.Fields{"toToken": req.To}).Error(err, "failed to get to token")
		return nil, err
	}

	fromTokenPrice, toTokenPrice, err := t.GetAndCacheTokenPrice(req, fromToken, toToken)
	if err != nil {
		logger.L.Fields(logger.Fields{"fromToken": req.From, "toToken": req.To}).Error(err, "failed to get token price")
		return nil, err
	}

	amountFrom := utils.ConvertStringBigIntToFloat(fromTokenPrice.Price, fromToken.Decimals)
	amountTo := utils.ConvertStringBigIntToFloat(toTokenPrice.Price, toToken.Decimals)

	convertTokenPrice := model.ConvertTokenPrice{}
	// fetch data
	amountToToken := fmt.Sprintf("%f", amountFrom/amountTo)
	convertTokenPrice.From.Amount = req.Amount
	convertTokenPrice.From.Symbol = req.From

	convertTokenPrice.To.Amount = amountToToken
	convertTokenPrice.To.Symbol = req.To

	return &convertTokenPrice, nil
}

func (t *tokenEntity) GetAndCacheTokenPrice(req request.ConvertTokenPrice, fromToken *model.Token, toToken *model.Token) (*model.TokenPrice, *model.TokenPrice, error) {
	now := time.Now().UTC()
	cacheTime := now.Add(-time.Minute * 30).Format("2006-01-02T15:04:05Z")
	// calculate amount to
	fromTokenPrice, err := t.store.Token.GetTokenPriceDetail(fromToken.Id, cacheTime)
	if err != nil && err != gorm.ErrRecordNotFound {
		logger.L.Fields(logger.Fields{"fromToken": req.From}).Error(err, "failed to get from token price detail")
		return nil, nil, err
	}

	if err == gorm.ErrRecordNotFound {
		fromTokenBirdeyePrice, err := t.service.Birdeye.GetCurrentPrice(fromToken.Address)
		if err != nil {
			logger.L.Fields(logger.Fields{"fromToken": req.From}).Error(err, "failed to get from token birdeye price")
			return nil, nil, err
		}

		fromTokenPrice = &model.TokenPrice{
			TokenId: fromToken.Id,
			Price:   utils.ConvertFloatToStringBigInt(fromTokenBirdeyePrice.Data.Value, fromToken.Decimals),
			Time:    time.Unix(fromTokenBirdeyePrice.Data.UpdateUnixTime, 0).UTC(),
		}

		err = t.store.Token.CreateTokenPrice(fromTokenPrice)
		if err != nil {
			logger.L.Fields(logger.Fields{"fromToken": req.From}).Error(err, "failed to create from token price")
			return nil, nil, err
		}
	}

	toTokenPrice, err := t.store.Token.GetTokenPriceDetail(toToken.Id, cacheTime)
	if err != nil && err != gorm.ErrRecordNotFound {
		logger.L.Fields(logger.Fields{"toToken": req.To}).Error(err, "failed to get to token price detail")
		return nil, nil, err
	}

	if err == gorm.ErrRecordNotFound {
		toTokenBirdeyePrice, err := t.service.Birdeye.GetCurrentPrice(toToken.Address)
		if err != nil {
			logger.L.Fields(logger.Fields{"toToken": req.To}).Error(err, "failed to get to token birdeye price")
			return nil, nil, err
		}
		toTokenPrice = &model.TokenPrice{
			TokenId: toToken.Id,
			Price:   utils.ConvertFloatToStringBigInt(toTokenBirdeyePrice.Data.Value, toToken.Decimals),
			Time:    time.Unix(toTokenBirdeyePrice.Data.UpdateUnixTime, 0).UTC(),
		}

		err = t.store.Token.CreateTokenPrice(toTokenPrice)
		if err != nil {
			logger.L.Fields(logger.Fields{"toToken": req.To}).Error(err, "failed to create frtoom token price")
			return nil, nil, err
		}
	}

	return fromTokenPrice, toTokenPrice, nil
}
