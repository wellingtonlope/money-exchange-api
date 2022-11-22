package usecase

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wellingtonlope/money-exchange-api/wallet/app/repository"
	"github.com/wellingtonlope/money-exchange-api/wallet/domain"
)

func TestNewCreateOutput(t *testing.T) {
	expected := CreateOutput{ID: "1"}
	wallet := domain.Wallet{ID: "1"}

	got := NewCreateOutput(wallet)

	assert.Equal(t, expected, got)
}

func TestCreateHandle(t *testing.T) {
	t.Run("should create an wallet", func(t *testing.T) {
		walletRepository := new(repository.WalletMock)
		uc := NewCreate(walletRepository)
		expectedWallet := domain.Wallet{ID: "1"}

		walletRepository.
			On("Create", domain.Wallet{}).
			Return(expectedWallet, nil)

		got, err := uc.Handle(CreateInput{})

		assert.Nil(t, err)
		assert.Equal(t, NewCreateOutput(expectedWallet), got)
	})

	t.Run("shouldn't create an wallet when database fail", func(t *testing.T) {
		walletRepository := new(repository.WalletMock)
		uc := NewCreate(walletRepository)
		expectedErr := errors.New("i'm an error")

		walletRepository.
			On("Create", domain.Wallet{}).
			Return(domain.Wallet{}, expectedErr)

		got, err := uc.Handle(CreateInput{})

		assert.Empty(t, got)
		assert.Equal(t, expectedErr, err)
	})
}
