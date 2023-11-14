package treasury

import (
	"context"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/walbety/transaction-app/exchange-service/internal/canonical"
	"io"
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
	RecordDateFormat                = "2006-01-02"
	TreasuryBaseUrl              = "https://api.fiscaldata.treasury.gov/services/api/fiscal_service"
	TreasuryExchangeRateEndpoint = "/v1/accounting/od/rates_of_exchange"
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

func (t Treasury) GetLatestRateGivenMaxDate(ctx context.Context, currency string, maxDate time.Time) (canonical.ExchangeRate, error) {
	var response ExchangeRateResponse

	maxDateFormated := maxDate.Format(RecordDateFormat)

	filterParam := fmt.Sprintf(asParam[filterCurrency], currency) + "," +
		fmt.Sprintf(asParam[filterRecordDate], maxDateFormated)

	params := url.Values{}
	params.Add(fields, asParam[fields])
	params.Add(pageNumber, asParam[pageNumber])
	params.Add(pageSize, asParam[pageSize])
	params.Add(sort, asParam[sort])
	params.Add(filter, filterParam)

	url, err := url.ParseRequestURI(TreasuryBaseUrl + TreasuryExchangeRateEndpoint)
	if err != nil {
		log.WithContext(ctx).
			WithField("treasury-url", TreasuryBaseUrl).
			Error("error at integration layer: error parsing treasury base url.")
		return canonical.ExchangeRate{}, err
	}
	//url.Path = TREASURY_EXCHANGE_RATE_ENDPOINT
	url.RawQuery = params.Encode()
	urlStr := fmt.Sprintf("%v", url)

	resp, err := http.Get(urlStr)
	defer resp.Body.Close()
	if err != nil {
		log.WithContext(ctx).
			WithField("url", urlStr).
			Error("error at integration layer: error calling treasury.")
		return canonical.ExchangeRate{}, canonical.ErrTreasuryServiceError
	}
	if resp.StatusCode != http.StatusOK {
		log.WithContext(ctx).
			WithFields(log.Fields{
				"currency":    currency,
				"maxDate":     maxDateFormated,
				"status-code": resp.Status,
			}).
			Warning("Treasury service returned with error.")
		return canonical.ExchangeRate{}, canonical.ErrInvalidCurrency
	} else {
		log.WithContext(ctx).
			WithFields(log.Fields{
				"currency": currency,
				"maxDate":  maxDateFormated,
			}).
			Info("Treasury service returned successfully.")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.WithError(err).Error("error at GetLatestRateGivenMaxDate - readAll")
		return canonical.ExchangeRate{}, err
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		log.WithError(err).Error("error at GetLatestRateGivenMaxDate - Unmarshal")
		return canonical.ExchangeRate{}, err
	}

	if len(response.Data) == 0 {
		return canonical.ExchangeRate{}, canonical.ErrInvalidCurrency
	}

	return MapToExchangeRate(response), nil
}

type ExchangeRateResponse struct {
	Data []struct {
		Currency     string `json:"currency"`
		ExchangeRate string `json:"exchange_rate"`
		RecordDate   string `json:"record_date"`
	}
}
