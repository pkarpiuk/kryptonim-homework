package rates

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/shopspring/decimal"
	"example/kryptonim-homework/utils"
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
	}
	result := make(map[string]*utils.ExchangeRate)
	resp, err := http.Get(fmt.Sprintf("https://openexchangerates.org/api/latest.json?app_id=%s&base=%s", APP_ID, BASE_CURRENCY))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&resultRaw)
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

	for i := 0; i < len(currencies)-1; i++ {
		for j := i + 1; j < len(currencies); j++ {
			result = append(result, doRate(currencies[i], currencies[j]), doRate(currencies[j], currencies[i]))
		}
	}

	data, _ := json.Marshal(result)
	fmt.Printf("%s\n", data)

	return result
}
