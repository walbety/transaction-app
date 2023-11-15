package service

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/walbety/transaction-app/transaction-service/internal/canonical"
	rate "github.com/walbety/transaction-app/transaction-service/internal/integration/mocks"
	repository "github.com/walbety/transaction-app/transaction-service/internal/repository/mocks"
	"math/big"
	"testing"
	"time"
)

var (
	repositoryMock *repository.MockPersistence
	rateMock       *rate.MockRate
	ctx context.Context
)

func initMocks() Service {
	rateMock = &rate.MockRate{}
	repositoryMock = &repository.MockPersistence{}

	return Service{
		exchange:    rateMock,
		persistence: repositoryMock,
	}
}

var transactionTestsCases = map[string]struct {
	id            string
	currency      string
	transactionIn canonical.Transaction

	repositoryReturn struct {
		transaction canonical.Transaction
		err         error
		assertCalls map[string]int
	}

	rateReturn struct {
		exchange    canonical.ExchangeRate
		err         error
		assertCalls map[string]int
	}
}{
	"GetLatestExchangeRateFromCurrencyAndDate_success": {
		id:       "test_case_id_success",
		currency: "Real",
	},
}

func TestGetLatestExchangeRateFromCurrencyAndDate(t *testing.T) {

	sampleDate := "20-07-1988"
	sampleTime,_ := time.Parse("02-01-2006", sampleDate)

	testCases := []struct {
		name                      string
		id                        string
		currency                  string
		date                      time.Time
		expectedReturn            canonical.ConvertedTransaction
		
		expectedReturnPersistence canonical.Transaction
		callPersist int
		
		expectedReturnExchange    canonical.ExchangeRate
		callExchange int

		errPersistence            error
		errExchange               error
		errFinal                  error
	}{
		{
			name:           "SUCCESS",
			id:             "123",
			currency: "Real",
			expectedReturn: canonical.ConvertedTransaction{
				Id:              "123",
				OriginalAmount:  *big.NewFloat(5.0),
				ConvertedAmount: *big.NewFloat(27.5),
				Currency:        "Real",
				TransactionDate: sampleTime,
				Description:     "test success desc",
				ExchangeRate:    "5.5",
			},
			expectedReturnPersistence : canonical.Transaction{
				Id:          "123",
				Amount:      "5",
				Date:        sampleTime,
				Description: "test success desc",
			},
			callPersist : 1,
			expectedReturnExchange: canonical.ExchangeRate{
				Currency:     "Real",
				ExchangeRate: "5.5",
				RecordDate:   sampleTime,
			},
			callExchange: 1,
		},
		{
			name:           "error at persistence.FindTransactionById",
			id:             "456",
			errPersistence: errors.New("find failed"),
			callPersist: 1,
			callExchange: 0,
			errFinal:       errors.New("find failed"),
		},
		{
			name:        "error at exchange.GetLatestRateGivenMaxDate",
			id:          "789",
			expectedReturnPersistence : canonical.Transaction{
				Id:          "789",
				Amount:      "5",
				Date:        sampleTime,
				Description: "test fail desc",
			},
			callPersist: 1,
			callExchange: 1,
			errExchange:  canonical.ErrDateMaxExceeded,
			errFinal:    canonical.ErrDateMaxExceeded,
		},
	}
	for _, tc := range testCases {
		svc := initMocks()
		ctx := context.Background()
		t.Run(tc.name, func(t *testing.T) {
			repositoryMock.On("FindTransactionById", mock.MatchedBy(func(transactionId string) bool {
				return transactionId == tc.id
			})).Return(tc.expectedReturnPersistence, tc.errPersistence)

			rateMock.On("GetLatestRateGivenMaxDate",
				mock.MatchedBy(func(currency string) bool {
					return currency == tc.currency
				}), mock.Anything).Return(tc.expectedReturnExchange, tc.errPersistence)

			returned, err := svc.GetLatestExchangeRateFromCurrencyAndDate(ctx, tc.id, tc.currency)


			assert.Equal(t, tc.errFinal, err)
			assert.EqualExportedValues(t, tc.expectedReturn, returned)
			repositoryMock.AssertNumberOfCalls(t, "FindTransactionById",tc.callPersist)
			rateMock.AssertNumberOfCalls(t, "GetLatestRateGivenMaxDate",tc.callExchange)
		})
	}
}
