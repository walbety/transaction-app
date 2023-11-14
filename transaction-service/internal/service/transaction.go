package service

import (
	"context"
	"github.com/walbety/transaction-app/transaction-service/internal/canonical"
	"time"
)


func (s Service) GetLatestExchangeRateFromCurrencyAndDate(ctx context.Context, currency string, date time.Time) (canonical.ExchangeRate, error){


	return canonical.ExchangeRate{Currency: "TESTE CURRENCY"},nil
}

func (s Service) SavePurchase(ctx context.Context, transaction canonical.Transaction) (string,error){

	id, err := s.persistence.SaveTransaction(ctx,&transaction)

	return id, err
}