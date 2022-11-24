package usecase

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/wellingtonlope/money-exchange-api/currency/app/repository"
	"github.com/wellingtonlope/money-exchange-api/currency/domain"
	"testing"
)

func TestUpdate_Handle(t *testing.T) {
	t.Run("should update a currency", func(t *testing.T) {
		currencyRepository := new(repository.CurrencyMock)
		uc := NewUpdate(currencyRepository)
		expectedBeforeUpdate := domain.Currency{ID: "BRL"}
		expectedUpdated := domain.Currency{ID: "BRL", Description: "Real", Country: "Brazil"}
		input := UpdateInput{expectedUpdated.ID, expectedUpdated.Description, expectedUpdated.Country}

		currencyRepository.
			On("GetByID", expectedUpdated.ID).
			Return(expectedBeforeUpdate, nil)
		currencyRepository.
			On("Update", expectedUpdated).
			Return(nil)

		got, err := uc.Handle(input)

		assert.Nil(t, err)
		assert.Equal(t, NewOutput(expectedUpdated), got)
	})

	t.Run("should fail when not found currency", func(t *testing.T) {
		currencyRepository := new(repository.CurrencyMock)
		uc := NewUpdate(currencyRepository)
		input := UpdateInput{ID: "BRL"}
		expectedError := repository.ErrRepositoryCurrencyNotFound

		currencyRepository.
			On("GetByID", input.ID).
			Return(domain.Currency{}, expectedError)

		got, err := uc.Handle(input)

		assert.Equal(t, expectedError, err)
		assert.Empty(t, got)
	})

	t.Run("should fail to updated currency", func(t *testing.T) {
		currencyRepository := new(repository.CurrencyMock)
		uc := NewUpdate(currencyRepository)
		expectedBeforeUpdate := domain.Currency{ID: "BRL"}
		expectedUpdated := domain.Currency{ID: "BRL", Description: "Real", Country: "Brazil"}
		input := UpdateInput{expectedUpdated.ID, expectedUpdated.Description, expectedUpdated.Country}
		expectedError := errors.New("i'm an error")

		currencyRepository.
			On("GetByID", expectedUpdated.ID).
			Return(expectedBeforeUpdate, nil)
		currencyRepository.
			On("Update", expectedUpdated).
			Return(expectedError)

		got, err := uc.Handle(input)

		assert.Equal(t, expectedError, err)
		assert.Empty(t, got)
	})
}
