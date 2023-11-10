package service

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/walbety/transaction-app/exchange-service/internal/canonical"
)

func (svc *Service) ListUsers(ctx context.Context) (canonical.User, error) {
	log.WithContext(ctx).Info("getting user (service)")
	return canonical.User{
		Name: "jao",
		Code: 123,
	}, nil
}
