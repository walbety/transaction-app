package treasury

import (
	log "github.com/sirupsen/logrus"
	"github.com/walbety/transaction-app/exchange-service/internal/canonical"
	"time"
)

func MapToExchangeRate(in ExchangeRateResponse) canonical.ExchangeRate {
	recordDate, err := time.Parse(RecordDateFormat, in.Data[0].RecordDate)
	if err != nil {
		log.WithField("record-date", recordDate).
			Errorf("error at integration mapper: invalid recordDate format returned")
	}
	return canonical.ExchangeRate{
		ExchangeRate: in.Data[0].ExchangeRate,
		Currency:     in.Data[0].Currency,
		RecordDate:   recordDate,
	}
}
