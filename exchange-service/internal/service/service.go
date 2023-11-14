package service

import (
	"context"
	"github.com/walbety/transaction-app/exchange-service/internal/canonical"
	"github.com/walbety/transaction-app/exchange-service/internal/config"
	"github.com/walbety/transaction-app/exchange-service/internal/integration"
	"time"
)

type Exchange interface {
	GetLatestExchangeRateFromCurrencyAndDate(ctx context.Context, currency string, date time.Time) (canonical.ExchangeRate, error)
}

type Service struct {
	exchange integration.ExchangeService
}

func New() Service {
	return Service{
		exchange: integration.New(config.Env.Integration.Exchange.Provider),
	}
}

func (svc *Service) welcome() {
	return
}
