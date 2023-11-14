package service

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/walbety/transaction-app/transaction-service/internal/canonical"
	"github.com/walbety/transaction-app/transaction-service/internal/integration"
	"github.com/walbety/transaction-app/transaction-service/internal/repository"
	"time"
)

type Transaction interface {
	GetLatestExchangeRateFromCurrencyAndDate(ctx context.Context, id, currency string, date time.Time) (canonical.ConvertedTransaction, error)
	SavePurchase(ctx context.Context, transaction canonical.Transaction) (string,error)
}

type Service struct {
	exchange integration.ExchangeService
	persistence repository.Persistence
}

func New(ctx context.Context) Transaction {
	persist, err := repository.NewClient(ctx)
	if err != nil {
		log.Panic("ERRROORRRORORO")
	}
	
	return Service{
		exchange: integration.NewExchangeService(),
		persistence: persist,
	}
}