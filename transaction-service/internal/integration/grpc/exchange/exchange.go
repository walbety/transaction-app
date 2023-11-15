package exchange

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/walbety/transaction-app/transaction-service/internal/canonical"
	"github.com/walbety/transaction-app/transaction-service/internal/config"
	"github.com/walbety/transaction-app/transaction-service/internal/integration/grpc"
	"github.com/walbety/transaction-app/transaction-service/internal/integration/grpc/exchange/impl"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type grpcImpl struct {
	client impl.ExchangeServiceClient
}
type Service interface {
	GetLatestRateGivenMaxDate(ctx context.Context, currency string, date time.Time) (canonical.ExchangeRate, error)
}

func New() Service {
	return &grpcImpl{
		client: impl.NewExchangeServiceClient(grpc.ConnectGprcInsecure(
			config.Env.Services.Exchange.Address,
		)),
	}
}

func (s grpcImpl) GetLatestRateGivenMaxDate(ctx context.Context, currency string, maxDate time.Time) (canonical.ExchangeRate, error) {

	req := impl.ExchangeRateRequest{
		Currency: currency,
		Date:     timestamppb.New(maxDate),
	}
	resp, err := s.client.GetExchangeRateByCurrencyAndDate(ctx, &req)
	if err != nil {
		log.WithError(err).Error("error at grpc.Exchange.GetExchangeRateByCurrencyAndDate")
		return canonical.ExchangeRate{}, nil
	}
	log.WithContext(ctx).WithField("exchange-service response:", resp).Info("Exchange-service grpc call returned successfully.")

	return canonical.ExchangeRate{
		Currency:     resp.Currency,
		RecordDate:   resp.Date.AsTime(),
		ExchangeRate: resp.ExchangeRate,
	}, nil
}
