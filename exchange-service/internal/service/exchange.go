package service

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/walbety/transaction-app/exchange-service/internal/canonical"
	"github.com/walbety/transaction-app/exchange-service/internal/config"
	"time"
)

func (svc *Service) GetLatestExchangeRateFromCurrencyAndDate(ctx context.Context, currency string, date time.Time) (canonical.ExchangeRate, error){
	exchange, err := svc.exchange.GetLatestRateGivenMaxDate(ctx, currency, date)
	if err != nil {
		log.WithError(err).
			Errorf("error retrieving exchange from provider: %s ", config.Env.Integration.Exchange.Provider)
		return canonical.ExchangeRate{}, err
	}

	return exchange,nil
}