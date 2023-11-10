package service

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/walbety/transaction-app/transaction-service/internal/canonical"
)

func (svc *Service) ListProduct(ctx context.Context) (canonical.Product, error) {
	log.WithContext(ctx).Info("getting user (service)")
	return canonical.Product{
		MarketName: "Goiabada Premium",
		Code:       1337,
		Price:      89.99,
	}, nil
}
