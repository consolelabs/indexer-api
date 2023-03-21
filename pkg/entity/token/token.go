package token

import (
	"fmt"
	"github.com/consolelabs/indexer-api/pkg/model"
	"github.com/consolelabs/indexer-api/pkg/service"
	"github.com/consolelabs/indexer-api/pkg/store"
	"strconv"
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

func (t *tokenEntity) GetConvertTokenPrice(amount, from, to string) (*model.ConvertTokenPrice, error) {
	convertTokenPrice := model.ConvertTokenPrice{}

	// calculate amount to
	fromTokenPrice, err := t.store.Token.GetTokenPriceDetail(from)
	if err != nil {
		return nil, err
	}
	toTokenPrice, err := t.store.Token.GetTokenPriceDetail(to)
	if err != nil {
		return nil, err
	}
	priceFromToken, err := strconv.ParseFloat(fromTokenPrice.Price, 64)
	if err != nil {
		return nil, err
	}
	priceToToken, err := strconv.ParseFloat(toTokenPrice.Price, 64)
	if err != nil {
		return nil, err
	}
	amountFromToken, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return nil, err
	}
	amountToToken := (amountFromToken * priceFromToken) / priceToToken

	// fetch data
	amountTo := fmt.Sprintf("%f", amountToToken)
	convertTokenPrice.From.Amount = amount
	convertTokenPrice.From.Symbol = from
	convertTokenPrice.From.Name = fromTokenPrice.Name

	convertTokenPrice.To.Amount = amountTo
	convertTokenPrice.To.Symbol = to
	convertTokenPrice.To.Name = toTokenPrice.Name

	return &convertTokenPrice, nil
}
