package exchange

import (
	"context"
	"github.com/walbety/transaction-app/transaction-service/internal/canonical"
	"time"
)

type Exchange struct {

}

func (s Exchange)GetLatestRateGivenMaxDate(ctx context.Context, currency string, maxDate time.Time) (canonical.ExchangeRate, error) {

	// bater no grpc
	return canonical.ExchangeRate{
		Currency: "aaaaaaaaaaaaaa",
	},nil
}

