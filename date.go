package poloniex

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type PoloniexDate time.Time

func (pd *PoloniexDate) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), "\"")
	t, err := time.Parse("2006-01-02 15:04:05", s)
	if err != nil {
		return errors.New(fmt.Sprintf("DateTime invalid (can't parse %s)", s))
	}
	*pd = PoloniexDate(t)
	return nil
}
