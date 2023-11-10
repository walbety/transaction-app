package integration

import (
	log "github.com/sirupsen/logrus"
	"github.com/walbety/transaction-app/exchange-service/internal/integration/treasury"
	"time"
)

func mapToExchangeRate(in treasury.ExchangeRateResponse) ExchangeRate {
	recordDate, err := time.Parse(treasury.RECORD_DATE_FORMAT, in.RecordDate)
	if err != nil {
		log.WithField("record-date", recordDate).
			Errorf("error at integration mapper: invalid recordDate format returned")
	}
	return ExchangeRate{
		ExchangeRate: in.ExchangeRate,
		Currency: in.Currency,
		RecordDate: recordDate,
	}
}