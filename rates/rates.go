package rates

import (
	"encoding/json"
	"fmt"
	"net/http"

	"example/kryptonim-homework/utils"

	"github.com/shopspring/decimal"
)

type RateRecord struct {
	From string  `json:"from"`
	To   string  `json:"to"`
	Rate float64 `json:"rate"`
}

const APP_ID = "e3955e944b054959a0ffff2551c97c66"
const BASE_CURRENCY = "USD"

func DownloadRates(currencies []string) map[string]*utils.ExchangeRate {
	var resultRaw struct {
		Rates map[string]float64
		Error bool
	}
	result := make(map[string]*utils.ExchangeRate)
	resp, err := http.Get(fmt.Sprintf("https://openexchangerates.org/api/latest.json?app_id=%s&base=%s", APP_ID, BASE_CURRENCY))
	if err != nil {
		return nil
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&resultRaw)
	if (err != nil) || resultRaw.Error {
		return nil
	}
	for ticker, val := range resultRaw.Rates {
		result[ticker] = &utils.ExchangeRate{Ticker: ticker, Precision: 18, Rate: decimal.NewFromFloat(val)}
	}
	return result
}

func DoRates(currencies []string) []*RateRecord {
	result := make([]*RateRecord, 0)
	exchangeRateTable := DownloadRates(currencies)
	doRate := func(fromCur, toCur string) *RateRecord {
		record := RateRecord{fromCur, toCur, 0.0}
		f, _ := decimal.NewFromString("1.0")
		t, _ := decimal.NewFromString("1.0")
		if fromCur != "USD" {
			f = exchangeRateTable[fromCur].Rate
		}
		if toCur != "USD" {
			t = exchangeRateTable[toCur].Rate
		}
		toPrecision := exchangeRateTable[toCur].Precision
		record.Rate, _ = t.Div(f).Round(toPrecision).Truncate(toPrecision).Float64()

		return &record
	}

	if exchangeRateTable != nil {
		return nil
	}

	for i := 0; i < len(currencies)-1; i++ {
		for j := i + 1; j < len(currencies); j++ {
			_, existsI := exchangeRateTable[currencies[i]]
			_, existsJ := exchangeRateTable[currencies[j]]
			if existsI && existsJ {
				result = append(result, doRate(currencies[i], currencies[j]), doRate(currencies[j], currencies[i]))
			}
		}
	}

	return result
}
