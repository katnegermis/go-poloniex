package poloniex

import (
	"errors"
	"strconv"
	"time"
)

type PoloniexTimestamp struct {
	time.Time
}

func (pd *PoloniexTimestamp) UnmarshalJSON(data []byte) error {
	i, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return errors.New("Timestamp invalid (can't parse int64)")
	}
	pd.Time = time.Unix(i, 0)
	return nil
}
