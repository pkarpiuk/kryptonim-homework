package utils

import (
	"github.com/shopspring/decimal"
)

type ExchangeRate struct {
	Ticker    string
	Precision int32
	Rate      decimal.Decimal
}

func DecimalNewFromString(s string) decimal.Decimal {
	v, _ := decimal.NewFromString(s)
	return v
}

func init() {
	decimal.DivisionPrecision = 100
}
