package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCurrency(t *testing.T) {
	t.Run("should create a currency", func(t *testing.T) {
		expectedCurrency := Currency{ID: "BRL", Description: "Real", Country: "Brazil"}

		got, err := NewCurrency(expectedCurrency.ID, expectedCurrency.Description, expectedCurrency.Country)

		assert.Nil(t, err)
		assert.Equal(t, expectedCurrency, got)
	})

	t.Run("should fail when ID is empty", func(t *testing.T) {
		expectedError := ErrCurrencyIDInvalid

		got, err := NewCurrency("", "Real", "Brazil")

		assert.Empty(t, got)
		assert.Equal(t, expectedError, err)
	})

	t.Run("should fail when ID is invalid", func(t *testing.T) {
		expectedError := ErrCurrencyIDInvalid

		got, err := NewCurrency("BR", "Real", "Brazil")

		assert.Empty(t, got)
		assert.Equal(t, expectedError, err)
	})
}

func TestCurrency_Update(t *testing.T) {
	expectedCurrency := Currency{ID: "BRL", Description: "Real", Country: "Brazil"}
	currency := Currency{ID: expectedCurrency.ID}

	currency.Update(expectedCurrency.Description, expectedCurrency.Country)

	assert.Equal(t, expectedCurrency, currency)
}
