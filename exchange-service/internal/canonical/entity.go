package canonical

import "time"

type ExchangeRate struct {
	Currency     string
	ExchangeRate string
	RecordDate   time.Time // note: using date.Date to encapsulate any date format adaptability
}

