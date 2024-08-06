package main

import (
	"net/http"
	"strings"

	"example/kryptonim-homework/exchange"
	"example/kryptonim-homework/rates"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

func getRates(c *gin.Context) {
	currenciesStr, _ := c.GetQuery("currencies")
	currenciesVec := strings.Split(currenciesStr, ",")
	if len(currenciesVec) <= 1 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid input parameters"})
	} else {
		result := rates.DoRates(currenciesVec)
		if result != nil {
			c.IndentedJSON(http.StatusOK, result)
		} else {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Error from openexchangerates"})
		}
	}
}

func getExchange(c *gin.Context) {
	from, okF := c.GetQuery("from")
	to, okT := c.GetQuery("to")
	amountStr, okA := c.GetQuery("amount")
	amount, okA2 := decimal.NewFromString(amountStr)
	if !okF || !okT || !okA || (okA2 != nil) {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid input parameters"})
	}
	result := exchange.DoExchange(from, to, amount)
	if result != nil {
		resultF64, _ := result.Float64()
		c.IndentedJSON(http.StatusOK, gin.H{"from": from, "to": to, "amount": resultF64})
	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid input parameters"})
	}
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/rates", getRates)
	r.GET("/exchange", getExchange)
	return r
}

func main() {
	router := setupRouter()
	router.Run(":8080")
}
