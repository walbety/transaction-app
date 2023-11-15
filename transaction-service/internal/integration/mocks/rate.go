//go:build test
// +build test

package mocks

import (
	"context"
	"github.com/stretchr/testify/mock"
	"github.com/walbety/transaction-app/transaction-service/internal/canonical"
	"time"
)

type MockRate struct {
	mock.Mock
}

func (m *MockRate) GetLatestRateGivenMaxDate(_ context.Context, currency string, maxDate time.Time) (canonical.ExchangeRate, error) {
	args := m.Called(currency, maxDate)
	return args.Get(0).(canonical.ExchangeRate), args.Error(1)
}
