package service

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/walbety/transaction-app/transaction-service/internal/canonical"
	"github.com/walbety/transaction-app/transaction-service/internal/config"
	"math/big"
	"time"
)


func (s Service) GetLatestExchangeRateFromCurrencyAndDate(ctx context.Context, id, currency string, date time.Time) (canonical.ConvertedTransaction, error){

	transaction, err := s.persistence.FindTransactionById(ctx,id)
	if err != nil {
		log.WithError(err).Error("error at FindTransactionById")
		return canonical.ConvertedTransaction{}, err
	}

	exchangeRate , err := s.exchange.GetLatestRateGivenMaxDate(ctx,currency,date)
	if err != nil {
		log.WithError(err).Error("error at GetLatestRateGivenMaxDate")
		return canonical.ConvertedTransaction{}, err
	}

	if err := validateDates(transaction.Date, exchangeRate.RecordDate); err != nil {
		return canonical.ConvertedTransaction{}, err
	}

	return calculateConvertedAmount(transaction,exchangeRate),nil
}

func (s Service) SavePurchase(ctx context.Context, transaction canonical.Transaction) (string,error){

	id, err := s.persistence.SaveTransaction(ctx,&transaction)

	return id, err
}

func validateDates(transactionDate ,exchangeRateDate time.Time) error {

	timeDelta := transactionDate.Sub(exchangeRateDate)
	if int64(timeDelta.Hours()/24/30) > config.Env.Validations.Exchange.DateMonthsMax {
		return canonical.ErrDateMaxExceeded
	}

	return nil
}

func calculateConvertedAmount(purchase canonical.Transaction, exchangeRate canonical.ExchangeRate) canonical.ConvertedTransaction {
	result := canonical.ConvertedTransaction{
		Id: purchase.Id,
		OriginalAmount: purchase.Amount,
		Description: purchase.Description,
		ExchangeRate: exchangeRate.ExchangeRate,
		TransactionDate:  purchase.Date,
		Currency: exchangeRate.Currency,
	}

	exchangeRat := big.NewRat(1,1)
	exchangeRat, _  = exchangeRat.SetString(exchangeRate.ExchangeRate)
	convertedAmount := exchangeRat.Mul(exchangeRat, &purchase.Amount)
	result.ConvertedAmount = *convertedAmount

	return result
}