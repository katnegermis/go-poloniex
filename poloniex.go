// Package Poloniex is an implementation of the Poloniex API in Golang.
package poloniex

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

const (
	API_BASE = "https://poloniex.com/" // Poloniex API endpoint
)

// New returns an instantiated poloniex struct
func New(apiKey, apiSecret string) *Poloniex {
	client := NewClient(apiKey, apiSecret)
	return &Poloniex{client}
}

// New returns an instantiated poloniex struct with custom timeout
func NewWithCustomTimeout(apiKey, apiSecret string, timeout time.Duration) *Poloniex {
	client := NewClientWithCustomTimeout(apiKey, apiSecret, timeout)
	return &Poloniex{client}
}

// poloniex represent a poloniex client
type Poloniex struct {
	client *client
}

// GetTickers is used to get the ticker for all markets
func (b *Poloniex) GetTickers() (tickers map[string]Ticker, err error) {
	r, err := b.client.do("GET", "public?command=returnTicker", nil, false)
	if err != nil {
		return
	}
	if err = json.Unmarshal(r, &tickers); err != nil {
		return
	}
	return
}

// GetVolumes is used to get the volume for all markets
func (b *Poloniex) GetVolumes() (vc VolumeCollection, err error) {
	r, err := b.client.do("GET", "public?command=return24hVolume", nil, false)
	if err != nil {
		return
	}
	if err = json.Unmarshal(r, &vc); err != nil {
		return
	}
	return
}

func (b *Poloniex) GetCurrencies() (currencies Currencies, err error) {
	r, err := b.client.do("GET", "public?command=returnCurrencies", nil, false)
	if err != nil {
		return
	}
	if err = json.Unmarshal(r, &currencies.Pair); err != nil {
		return
	}
	return
}

// GetOrderBook is used to get retrieve the orderbook for a given market
// market: a string literal for the market (ex: BTC_NXT). 'all' not implemented.
// cat: bid, ask or both to identify the type of orderbook to return.
// depth: how deep of an order book to retrieve
func (b *Poloniex) GetOrderBook(market, cat string, depth int) (orderBook OrderBook, err error) {
	// not implemented
	if cat != "bid" && cat != "ask" && cat != "both" {
		cat = "both"
	}
	if depth > 100 {
		depth = 100
	}
	if depth < 1 {
		depth = 1
	}

	r, err := b.client.do("GET", fmt.Sprintf("public?command=returnOrderBook&currencyPair=%s&depth=%d", strings.ToUpper(market), depth), nil, false)
	if err != nil {
		return
	}
	if err = json.Unmarshal(r, &orderBook); err != nil {
		return
	}
	if orderBook.Error != "" {
		err = errors.New(orderBook.Error)
		return
	}
	return
}

// Returns candlestick chart data. Required GET parameters are "currencyPair",
// "period" (candlestick period in seconds; valid values are 300, 900, 1800,
// 7200, 14400, and 86400), "start", and "end". "Start" and "end" are given in
// UNIX timestamp format and used to specify the date range for the data
// returned.
func (b *Poloniex) ChartData(currencyPair string, period int, start, end time.Time) (candles []*CandleStick, err error) {
	r, err := b.client.do("GET", fmt.Sprintf(
		"/public?command=returnChartData&currencyPair=%s&period=%d&start=%d&end=%d",
		strings.ToUpper(currencyPair),
		period,
		start.Unix(),
		end.Unix(),
	), nil, false)
	if err != nil {
		return
	}

	if err = json.Unmarshal(r, &candles); err != nil {
		return
	}

	return
}

func (b *Poloniex) GetBalances() (balances map[string]Balance, err error) {
	balances = make(map[string]Balance)
	r, err := b.client.doCommand("returnCompleteBalances", nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(r, &balances); err != nil {
		return
	}

	return
}

func (b *Poloniex) GetAllTradeHistory(since time.Time) (trades map[string][]Trade, err error) {
	trades = make(map[string][]Trade)
	payload := map[string]string{
		"currencyPair": "all",
		"start":        fmt.Sprintf("%d", since.Unix()),
	}
	r, err := b.client.doCommand("returnTradeHistory", payload)
	if err != nil {
		return
	}

	if err = json.Unmarshal(r, &trades); err != nil {
		return
	}

	return
}
