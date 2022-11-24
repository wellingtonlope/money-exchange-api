package usecase

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/wellingtonlope/money-exchange-api/currency/app/repository"
	"github.com/wellingtonlope/money-exchange-api/currency/domain"
	"testing"
)

func TestNewCreateOutput(t *testing.T) {
	t.Run("should create an output", func(t *testing.T) {
		currency := domain.Currency{ID: "BRL", Description: "Real", Country: "Brazil"}
		expectedOutput := Output{currency.ID, currency.Description, currency.Country}

		got := NewOutput(currency)

		assert.Equal(t, expectedOutput, got)
	})
}

func TestCreate_Handle(t *testing.T) {
	t.Run("should create a currency", func(t *testing.T) {
		currencyRepository := new(repository.CurrencyMock)
		uc := NewCreate(currencyRepository)
		expectedCurrency := domain.Currency{ID: "BRL", Description: "Real", Country: "Brazil"}
		expectedOutput := Output{expectedCurrency.ID, expectedCurrency.Description, expectedCurrency.Country}

		currencyRepository.
			On("Create", expectedCurrency).
			Return(expectedCurrency, nil)

		input := CreateInput{expectedCurrency.ID, expectedCurrency.Description, expectedCurrency.Country}
		got, err := uc.Handle(input)

		assert.Nil(t, err)
		assert.Equal(t, expectedOutput, got)
	})

	t.Run("should fail when ID is invalid", func(t *testing.T) {
		currencyRepository := new(repository.CurrencyMock)
		uc := NewCreate(currencyRepository)

		input := CreateInput{"BR", "Real", "Brazil"}
		got, err := uc.Handle(input)

		assert.Equal(t, domain.ErrCurrencyIDInvalid, err)
		assert.Empty(t, got)
	})

	t.Run("should fail when repository fail", func(t *testing.T) {
		currencyRepository := new(repository.CurrencyMock)
		uc := NewCreate(currencyRepository)
		expectedCurrency := domain.Currency{ID: "BRL", Description: "Real", Country: "Brazil"}
		expectedErr := errors.New("i'm an error")

		currencyRepository.
			On("Create", expectedCurrency).
			Return(domain.Currency{}, expectedErr)

		input := CreateInput{expectedCurrency.ID, expectedCurrency.Description, expectedCurrency.Country}
		got, err := uc.Handle(input)

		assert.Empty(t, got)
		assert.Equal(t, expectedErr, err)
	})
}
