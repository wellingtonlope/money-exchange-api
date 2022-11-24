package usecase

import (
	"github.com/stretchr/testify/assert"
	"github.com/wellingtonlope/money-exchange-api/currency/app/repository"
	"github.com/wellingtonlope/money-exchange-api/currency/domain"
	"testing"
)

func TestNewGetByIDOutput(t *testing.T) {
	t.Run("should create an output", func(t *testing.T) {
		currency := domain.Currency{ID: "BRL", Description: "Real", Country: "Brazil"}
		expectedOutput := Output{currency.ID, currency.Description, currency.Country}

		got := NewOutput(currency)

		assert.Equal(t, expectedOutput, got)
	})
}

func TestGetByID_Handle(t *testing.T) {
	t.Run("should get currency by ID", func(t *testing.T) {
		currencyRepository := new(repository.CurrencyMock)
		uc := NewGetByID(currencyRepository)
		expectedCurrency := domain.Currency{ID: "BRL", Description: "Real", Country: "Brazil"}

		currencyRepository.
			On("GetByID", expectedCurrency.ID).
			Return(expectedCurrency, nil)

		got, err := uc.Handle(expectedCurrency.ID)

		assert.Nil(t, err)
		assert.Equal(t, NewOutput(expectedCurrency), got)
	})

	t.Run("should fail when repository fail", func(t *testing.T) {
		currencyRepository := new(repository.CurrencyMock)
		uc := NewGetByID(currencyRepository)
		expectedID := "BRL"
		expectedError := repository.ErrRepositoryCurrencyNotFound

		currencyRepository.
			On("GetByID", expectedID).
			Return(domain.Currency{}, expectedError)

		got, err := uc.Handle(expectedID)

		assert.Equal(t, expectedError, err)
		assert.Empty(t, got)
	})
}
