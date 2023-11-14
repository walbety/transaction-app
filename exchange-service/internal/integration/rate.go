package integration

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/walbety/transaction-app/exchange-service/internal/canonical"
	"github.com/walbety/transaction-app/exchange-service/internal/integration/treasury"
	"time"
)

type ExchangeService interface {
	GetLatestRateGivenMaxDate(ctx context.Context, currency string, maxDate time.Time) (canonical.ExchangeRate, error)
}

const (
	TREASURY = "treasury"
)

// add a layer to enable to switching sources through configuration
func New(provider string) ExchangeService {
	switch provider {
	case TREASURY, "": // default service
		return treasury.Treasury{}
	default:
		log.Fatal("No source defined for exchange integration")
		return nil
	}
}
