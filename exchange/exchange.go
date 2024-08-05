package exchange

import (
	"example/kryptonim-homework/utils"

	"github.com/shopspring/decimal"
)

var exchangeRateTable = map[string]*utils.ExchangeRate{
	"BEER":  {Ticker: "BEER", Precision: 18, Rate: utils.DecimalNewFromString("0.00002461")},
	"FLOKI": {Ticker: "FLOKI", Precision: 18, Rate: utils.DecimalNewFromString("0.0001428")},
	"GATE":  {Ticker: "GATE", Precision: 18, Rate: utils.DecimalNewFromString("6.87")},
	"USDT":  {Ticker: "USDT", Precision: 6, Rate: utils.DecimalNewFromString("0.999")},
	"WBTC":  {Ticker: "WBTC", Precision: 8, Rate: utils.DecimalNewFromString("57037.22")}}

func DoExchange(from string, to string, amount decimal.Decimal) *decimal.Decimal {
	fromER, ok1 := exchangeRateTable[from]
	toER, ok2 := exchangeRateTable[to]
	if !ok1 || !ok2 {
		return nil
	}
	fromVal, toVal := fromER.Rate, toER.Rate
	result := fromVal.Mul(amount).Div(toVal).Round(toER.Precision).Truncate(toER.Precision)
	return &result
}
