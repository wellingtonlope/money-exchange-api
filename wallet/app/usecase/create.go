package usecase

import (
	"github.com/wellingtonlope/money-exchange-api/wallet/app/repository"
	"github.com/wellingtonlope/money-exchange-api/wallet/domain"
)

type (
	CreateInput  struct{}
	CreateOutput struct {
		ID string
	}
)

func NewCreateOutput(wallet domain.Wallet) CreateOutput {
	return CreateOutput{wallet.ID}
}

type Create interface {
	Handle(CreateInput) (CreateOutput, error)
}

type create struct {
	walletRepository repository.Wallet
}

func NewCreate(walletRepository repository.Wallet) Create {
	return &create{walletRepository}
}

func (uc *create) Handle(input CreateInput) (CreateOutput, error) {
	wallet := domain.NewWallet()

	wallet, err := uc.walletRepository.Create(wallet)
	if err != nil {
		return CreateOutput{}, err
	}

	return NewCreateOutput(wallet), nil
}
