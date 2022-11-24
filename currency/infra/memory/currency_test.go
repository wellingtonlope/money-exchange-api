package memory

import (
	"github.com/stretchr/testify/assert"
	"github.com/wellingtonlope/money-exchange-api/currency/app/repository"
	"github.com/wellingtonlope/money-exchange-api/currency/domain"
	"testing"
)

func TestCurrency_Create(t *testing.T) {
	t.Run("should create a currency", func(t *testing.T) {
		repo := &currency{}
		expectedInserted := domain.Currency{ID: "BRL", Description: "Real", Country: "Brazil"}

		got, err := repo.Create(expectedInserted)

		assert.Nil(t, err)
		assert.Equal(t, expectedInserted, got)
		assert.Len(t, repo.currencies, 1)
		assert.Equal(t, expectedInserted, repo.currencies[0])
	})

	t.Run("should fail create a duplicated currency", func(t *testing.T) {
		expectedInserted := domain.Currency{ID: "BRL", Description: "Real", Country: "Brazil"}
		repo := &currency{[]domain.Currency{expectedInserted}}
		expectedErr := repository.ErrRepositoryCurrencyDuplicated

		insertDuplicated := domain.Currency{ID: expectedInserted.ID}
		got, err := repo.Create(insertDuplicated)

		assert.Equal(t, expectedErr, err)
		assert.Empty(t, got)
		assert.Len(t, repo.currencies, 1)
		assert.Equal(t, expectedInserted, repo.currencies[0])
	})
}

func TestCurrency_Update(t *testing.T) {
	t.Run("should update a currency", func(t *testing.T) {
		repo := &currency{[]domain.Currency{{ID: "BRL"}}}
		expectedUpdated := domain.Currency{ID: "BRL", Description: "Real", Country: "Brazil"}

		err := repo.Update(expectedUpdated)

		assert.Nil(t, err)
		assert.Len(t, repo.currencies, 1)
		assert.Equal(t, expectedUpdated, repo.currencies[0])
	})

	t.Run("should not update when currency not exists", func(t *testing.T) {
		repo := &currency{}
		expectedUpdated := domain.Currency{ID: "BRL", Description: "Real", Country: "Brazil"}

		err := repo.Update(expectedUpdated)

		assert.Equal(t, repository.ErrRepositoryCurrencyNotFound, err)
	})
}

func TestCurrency_GetByID(t *testing.T) {
	t.Run("should get currency by ID", func(t *testing.T) {
		expected := domain.Currency{ID: "BRL", Description: "Real", Country: "Brazil"}
		repo := &currency{[]domain.Currency{expected}}

		got, err := repo.GetByID(expected.ID)

		assert.Nil(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("should not get by ID when currency not exists", func(t *testing.T) {
		repo := &currency{}

		got, err := repo.GetByID("BRL")

		assert.Equal(t, repository.ErrRepositoryCurrencyNotFound, err)
		assert.Empty(t, got)
	})
}
