package poloniex

import (
	"strings"
)

type Trade struct {
	GlobalTradeId int64        `json:"globalTradeId"`
	TradeId       int64        `json:"tradeId,string"`
	Date          PoloniexDate `json:"date"`
	Rate          string       `json:"rate"`
	Amount        string       `json:"amount"`
	Total         string       `json:"total"`
	Fee           string       `json:"fee"`
	OrderNumber   string       `json:"orderNumber"`
	Type          TradeType    `json:"type"`
	Category      string       `json:"category"`
}

type TradeType string

func (t *TradeType) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	*t = TradeType(s)
	return nil
}

const (
	TradeTypeBuy  = "buy"
	TradeTypeSell = "sell"
)
