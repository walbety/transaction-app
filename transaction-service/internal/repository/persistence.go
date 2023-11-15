package repository

import (
	"context"
	"github.com/walbety/transaction-app/transaction-service/internal/canonical"
	"github.com/walbety/transaction-app/transaction-service/internal/repository/mongodb"
)

type Persistence interface {
	SaveTransaction(ctx context.Context, transaction *canonical.Transaction) (string, error)
	FindTransactionById(ctx context.Context, id string) (canonical.Transaction, error)
}

func NewClient(ctx context.Context) (Persistence, error) {
	return mongodb.New(ctx)
}
