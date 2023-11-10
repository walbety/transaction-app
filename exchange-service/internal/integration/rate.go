package integration

import (
	log "github.com/sirupsen/logrus"
	"github.com/walbety/transaction-app/exchange-service/internal/integration/treasury"
	"google.golang.org/genproto/googleapis/type/date"
)

type ExchangeService interface {
	GetLatestRateGivenMaxDate(currency string, maxDate date.Date) (RateResponse, error)
}

const (
	TREASURY = "treasury"
)

type RateResponse struct {
	Currency     string
	ExchangeRate string
	RecordDate   date.Date // note: using date.Date to encapsulate any date format adaptability
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
