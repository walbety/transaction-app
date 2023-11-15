package grpc

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/walbety/transaction-app/exchange-service/internal/channels/grpc/impl"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s grpcServer) GetExchangeRateByCurrencyAndDate(context context.Context, request *impl.ExchangeRateRequest) (*impl.ExchangeRateResponse, error){

	log.WithFields(log.Fields{"req.currency": request.Currency})

	resp, err := s.svc.GetLatestExchangeRateFromCurrencyAndDate(context, request.Currency, request.Date.AsTime())
	return &impl.ExchangeRateResponse{
		Currency: resp.Currency,
		Date: timestamppb.New(resp.RecordDate),
		ExchangeRate: resp.ExchangeRate,
	}, err
}