package utils

import (
	"fmt"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/params"
)

func WeiToEther(wei *big.Int, decimals ...int) *big.Float {
	f := new(big.Float)
	f.SetPrec(236) //  IEEE 754 octuple-precision binary floating-point format: binary256
	f.SetMode(big.ToNearestEven)
	fWei := new(big.Float)
	fWei.SetPrec(236) //  IEEE 754 octuple-precision binary floating-point format: binary256
	fWei.SetMode(big.ToNearestEven)

	var e *big.Float
	if len(decimals) == 0 {
		e = big.NewFloat(params.Ether)
	} else {
		e = big.NewFloat(math.Pow(10, float64(decimals[0])))
	}
	return f.Quo(fWei.SetInt(wei), e)
}

func StringWeiToEther(stringWei string, decimals int) *big.Float {
	if decimals == 0 {
		decimals = 18
	}
	wei := new(big.Int)
	wei.SetString(stringWei, 10)
	return WeiToEther(wei, decimals)
}

func ConvertFloatToStringBigInt(value float64, decimals int64) string {
	bigFloatValue := big.NewFloat(math.Floor(value * math.Pow(10, float64(decimals))))
	return fmt.Sprintf("%.0f", bigFloatValue)
}

func ConvertStringBigIntToFloat(value string, decimals int64) float64 {
	bigIntValue := new(big.Int)
	bigIntValue.SetString(value, 10)
	bigFloatValue := new(big.Float)
	bigFloatValue.SetInt(bigIntValue)
	res, _ := bigFloatValue.Quo(bigFloatValue, big.NewFloat(math.Pow(10, float64(decimals)))).Float64()
	return res
}
