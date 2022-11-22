package memory

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wellingtonlope/money-exchange-api/wallet/app/repository"
	"github.com/wellingtonlope/money-exchange-api/wallet/domain"
)

func TestWalletCreate(t *testing.T) {
	t.Run("should create a wallet", func(t *testing.T) {
		repo := NewWallet()

		got, err := repo.Create(domain.Wallet{})

		assert.Nil(t, err)
		assert.NotEmpty(t, got)
	})
}

func TestGetByID(t *testing.T) {
	t.Run("should get wallet by ID", func(t *testing.T) {
		repo := NewWallet()
		inserted, _ := repo.Create(domain.Wallet{})

		got, err := repo.GetByID(inserted.ID)

		assert.Nil(t, err)
		assert.Equal(t, inserted, got)
	})

	t.Run("should not get wallet by ID when not exists", func(t *testing.T) {
		repo := NewWallet()
		_, _ = repo.Create(domain.Wallet{})

		got, err := repo.GetByID("1")

		assert.Equal(t, repository.ErrRepositoryWalletNotFound, err)
		assert.Empty(t, got)
	})
}
