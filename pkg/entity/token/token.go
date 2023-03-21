package token

import (
	"errors"
	"fmt"
	"math"
	"math/big"
	"strconv"

	"github.com/consolelabs/indexer-api/pkg/model"
	"github.com/consolelabs/indexer-api/pkg/service"
	"github.com/consolelabs/indexer-api/pkg/store"
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

	priceFromToken := new(big.Int)
	priceFromToken, ok := priceFromToken.SetString(fromTokenPrice.Price, 10)
	if !ok {
		return nil, errors.New("failed to convert big int")
	}
	priceToToken := new(big.Int)
	priceToToken, ok = priceToToken.SetString(toTokenPrice.Price, 10)
	if !ok {
		return nil, errors.New("failed to convert big int")
	}

	amountFromToken, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return nil, err
	}

	decimalFromToken := new(big.Int)
	decimalFromToken = decimalFromToken.SetInt64(int64(math.Pow10(fromTokenPrice.Decimals)))
	if !ok {
		return nil, errors.New("failed to convert big int")
	}

	decimalToToken := new(big.Int)
	decimalToToken = decimalToToken.SetInt64(int64(math.Pow10(toTokenPrice.Decimals)))
	if !ok {
		return nil, errors.New("failed to convert big int")
	}

	divisionFrom := new(big.Int).Div(priceFromToken, decimalFromToken)
	fmt.Println(divisionFrom.Int64())
	divisionTo := new(big.Int).Div(priceToToken, decimalToToken)
	fmt.Println(divisionTo.Int64())

	amountToToken := amountFromToken * float64(divisionFrom.Int64()) / float64(divisionTo.Int64())

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
