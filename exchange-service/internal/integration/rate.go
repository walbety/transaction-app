package integration

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/walbety/transaction-app/exchange-service/internal/integration/treasury"
	"time"
)

type ExchangeService interface {
	GetLatestRateGivenMaxDate(ctx context.Context, currency string, maxDate time.Time) (ExchangeRate, error)
}

const (
	TREASURY = "treasury"
)

type ExchangeRate struct {
	Currency     string
	ExchangeRate string
	RecordDate   time.Time // note: using date.Date to encapsulate any date format adaptability
}

// add a layer to enable to switching sources through configuration
func New(source string) ExchangeService {
	switch source {
	case TREASURY, "": // default service
		return treasury.Treasury{}
	default:
		log.Fatal("No source defined for exchange integration")
		return nil
	}
}
