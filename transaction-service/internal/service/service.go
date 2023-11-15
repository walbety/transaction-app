package service

import (
	"context"
	"github.com/walbety/transaction-app/transaction-service/internal/canonical"
	"github.com/walbety/transaction-app/transaction-service/internal/integration"
	"github.com/walbety/transaction-app/transaction-service/internal/repository"
)

type Transaction interface {
	GetLatestExchangeRateFromCurrencyAndDate(ctx context.Context, id, currency string) (canonical.ConvertedTransaction, error)
	SavePurchase(ctx context.Context, transaction canonical.Transaction) (string, error)
}

type Service struct {
	exchange    integration.ExchangeService
	persistence repository.Persistence
}

func New(integ integration.ExchangeService, repo repository.Persistence) Transaction {
	return Service{
		exchange:    integ,
		persistence: repo,
	}
}
