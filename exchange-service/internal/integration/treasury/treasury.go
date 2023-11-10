package treasury

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/walbety/transaction-app/exchange-service/internal/integration"
	"net/http"

	"net/url"
	"time"
)

type Treasury struct {
}

type (
	param string
)

const (
	RECORD_DATE_FORMAT              = "2006-01-02"
	TREASURY_BASE_URL               = "https://api.fiscaldata.treasury.gov/services/api/fiscal_service/"
	TREASURY_EXCHANGE_RATE_ENDPOINT = "v1/accounting/od/rates_of_exchange"
)

var (
	fields           string = "fields"
	pageNumber              = "page[number]"
	pageSize                = "page[size]"
	sort                    = "sort"
	filter                  = "filter"
	filterCurrency          = "filterCurrency"
	filterRecordDate        = "filterRecordDate"

	// note: simple way to isolate each combination of params of each useCase
	asParam = map[string]string{
		fields:           "currency,country_currency_desc,exchange_rate,record_date",
		pageNumber:       "1",
		pageSize:         "1",
		sort:             "-record_date",
		filterCurrency:   "currency:eq:%s",
		filterRecordDate: "record_date:lte:%s",
	}
)

func (t Treasury) GetLatestRateGivenMaxDate(ctx context.Context, currency string, maxDate time.Time) (integration.ExchangeRate, error) {
	var response ExchangeRateResponse

	maxDateFormated := maxDate.Format(RECORD_DATE_FORMAT)

	filterParam := fmt.Sprintf(asParam[filterCurrency], currency) + "," +
		fmt.Sprintf(asParam[filterRecordDate], maxDateFormated)

	params := url.Values{}
	params.Add(fields, asParam[fields])
	params.Add(pageNumber, asParam[pageNumber])
	params.Add(pageSize, asParam[pageSize])
	params.Add(sort, asParam[sort])
	params.Add(filter, filterParam)

	url, err := url.ParseRequestURI(TREASURY_BASE_URL)
	if err != nil {
		log.WithContext(ctx).
			WithField("treasury-url", TREASURY_BASE_URL).
			Error("error at integration layer: error parsing treasury base url.")
		return integration.ExchangeRate{}, err
	}
	url.Path = TREASURY_EXCHANGE_RATE_ENDPOINT
	url.RawQuery = params.Encode()
	urlStr := fmt.Sprintf("%v", url)

	resp, err := http.Get(urlStr)
	if err != nil {
		log.WithContext(ctx).
			WithField("url", urlStr).
			Error("error at integration layer: error calling treasury.")
		return integration.ExchangeRate{}, err
	}
	if resp.StatusCode != http.StatusOK {
		log.WithContext(ctx).
			WithFields(log.Fields{
				"currency": currency,
				"maxDate": maxDateFormated,
				"status-code": resp.Status,
			}).
			Warning("Treasury service returned with error.")
	} else {
		log.WithContext(ctx).
			WithFields(log.Fields{
				"currency": currency,
				"maxDate": maxDateFormated,
			}).
			Info("Treasury service returned successfully.")
	}


	return integration.ExchangeRate{}, nil
}

type ExchangeRateResponse struct {
	Currency     string `json:"currency"`
	ExchangeRate string `json:"exchange_rate"`
	RecordDate   string `json:"record_date"`
}
