package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wellingtonlope/money-exchange-api/wallet/app/repository"
	"github.com/wellingtonlope/money-exchange-api/wallet/domain"
)

func TestNewGetByIDOutput(t *testing.T) {
	expected := GetByIDOutput{ID: "1"}
	wallet := domain.Wallet{ID: "1"}

	got := NewGetByIDOutput(wallet)

	assert.Equal(t, expected, got)
}

func TestGetByIDHandle(t *testing.T) {
	t.Run("should get wallet by ID", func(t *testing.T) {
		walletRepository := new(repository.WalletMock)
		uc := NewGetByID(walletRepository)
		expectedWallet := domain.Wallet{ID: "1"}

		walletRepository.
			On("GetByID", expectedWallet.ID).
			Return(expectedWallet, nil)

		got, err := uc.Handle(expectedWallet.ID)

		assert.Nil(t, err)
		assert.Equal(t, NewGetByIDOutput(expectedWallet), got)
	})

	t.Run("shouldn't get wallet by ID when database return error", func(t *testing.T) {
		walletRepository := new(repository.WalletMock)
		uc := NewGetByID(walletRepository)
		expectedErr := repository.ErrRepositoryWalletNotFound

		walletRepository.
			On("GetByID", "1").
			Return(domain.Wallet{}, expectedErr)

		got, err := uc.Handle("1")

		assert.Equal(t, expectedErr, err)
		assert.Empty(t, got)
	})

}
