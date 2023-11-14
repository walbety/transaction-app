package integration

import (
	"context"
	"github.com/walbety/transaction-app/transaction-service/internal/canonical"
	"github.com/walbety/transaction-app/transaction-service/internal/integration/grpc/exchange"
	"time"
)

type ExchangeService interface {
	GetLatestRateGivenMaxDate(ctx context.Context, currency string, maxDate time.Time) (canonical.ExchangeRate, error)
}

const (

)

func NewExchangeService() ExchangeService{
	return exchange.Exchange{}
}

