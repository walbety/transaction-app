//go:build test
// +build test

package mocks

import (
	"context"
	"github.com/stretchr/testify/mock"
	"github.com/walbety/transaction-app/transaction-service/internal/canonical"
)

type MockPersistence struct {
	mock.Mock
}

func (m *MockPersistence) SaveTransaction(_ context.Context, transaction *canonical.Transaction) (string, error) {
	args := m.Called(transaction)
	return args.Get(0).(string), args.Error(1)

}

func (m *MockPersistence) FindTransactionById(_ context.Context, id string) (canonical.Transaction, error) {
	args := m.Called(id)
	return args.Get(0).(canonical.Transaction), args.Error(1)
}
