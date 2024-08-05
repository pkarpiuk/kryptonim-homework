package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"example/kryptonim-homework/exchange"
	"example/kryptonim-homework/rates"
	"example/kryptonim-homework/utils"
)

/*
 TODO:
   - obsługa błędów
*/

func getRates(c *gin.Context) {
	currenciesStr, _ := c.GetQuery("currencies")
	currenciesVec := strings.Split(currenciesStr, ",")
	fmt.Println(currenciesStr)
	result := rates.DoRates(currenciesVec)
	c.IndentedJSON(http.StatusOK, result)
}

func getExchange(c *gin.Context) {
	from, _ := c.GetQuery("from")
	to, _ := c.GetQuery("to")
	amountStr, _ := c.GetQuery("amount")
	amount := utils.DecimalNewFromString(amountStr)
	result := exchange.DoExchange(from, to, amount)
	resultF64, _ := result.Float64()
	c.IndentedJSON(http.StatusOK, gin.H{"from": from, "to": to, "amount": resultF64})
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
