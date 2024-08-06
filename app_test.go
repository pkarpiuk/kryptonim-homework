package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test200(t *testing.T) {
	router := setupRouter()
	var resultRaw []*struct {
		From string
		To   string
		Rate float64
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/rates?currencies=USD,GBP,EUR", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	err := json.NewDecoder(w.Body).Decode(&resultRaw)
	assert.Equal(t, nil, err)
	arr := make([]string, 0)
	for _, rec := range resultRaw {
		arr = append(arr, fmt.Sprint(rec.From, rec.To))
	}
	sort.Strings(arr)
	res := strings.Join(arr, ",")
	assert.Equal(t, "EURGBP,EURUSD,GBPEUR,GBPUSD,USDEUR,USDGBP", res)
}

func Test400NoParameter(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/rates", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}

func Test400EmptyParameter(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/rates?currencies=", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}

func Test400OneTicker(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/rates?currencies=PLN", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}

func Test200TwoTickersButOneBad(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/rates?currencies=PLN,USDQ", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "[]", w.Body.String())
}
