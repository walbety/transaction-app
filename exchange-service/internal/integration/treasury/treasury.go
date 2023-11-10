package treasury

import (
	"github.com/walbety/transaction-app/exchange-service/internal/integration"
	"google.golang.org/genproto/googleapis/type/date"
)

type Treasury struct{}

type (
	param string
)

const (
	RECORD_DATE_FORMAT = "2006-01-02"
)

var (
	fields         param = "fields"
	pageNumber     param = "page[number]"
	pageSize       param = "page[size]"
	sort           param = "sort"
	filterCurrency param = "filter"

	// note: simple way to isolate each combination of params of each useCase
	asParam = map[param]string{
		fields:         "currency,country_currency_desc,exchange_rate,record_date",
		pageNumber:     "1",
		pageSize:       "1",
		sort:           "-record_date",
		filterCurrency: "currency:eq:%s",
	}
)

func (t Treasury) GetLatestRateGivenMaxDate(currency string, maxDate date.Date) (integration.RateResponse, error) {
	//var response ExchangeRate


	return integration.RateResponse{}, nil
}

type ExchangeRate struct {
	Currency     string `json:"currency"`
	ExchangeRate string `json:"exchange_rate"`
	RecordDate   string `json:"record_date"`
}
