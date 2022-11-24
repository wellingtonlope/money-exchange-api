package repository

import (
	"github.com/stretchr/testify/mock"
	"github.com/wellingtonlope/money-exchange-api/currency/domain"
)

type CurrencyMock struct {
	mock.Mock
}

func (m *CurrencyMock) Create(currency domain.Currency) (domain.Currency, error) {
	args := m.Called(currency)
	var result domain.Currency
	if args.Get(0) != nil {
		result = args.Get(0).(domain.Currency)
	}
	return result, args.Error(1)
}

func (m *CurrencyMock) Update(currency domain.Currency) error {
	args := m.Called(currency)
	return args.Error(0)
}

func (m *CurrencyMock) GetByID(ID string) (domain.Currency, error) {
	args := m.Called(ID)
	var result domain.Currency
	if args.Get(0) != nil {
		result = args.Get(0).(domain.Currency)
	}
	return result, args.Error(1)
}
