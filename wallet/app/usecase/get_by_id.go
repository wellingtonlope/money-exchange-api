package usecase

import (
	"github.com/wellingtonlope/money-exchange-api/wallet/app/repository"
	"github.com/wellingtonlope/money-exchange-api/wallet/domain"
)

type (
	GetByIDOutput struct {
		ID string
	}
)

func NewGetByIDOutput(wallet domain.Wallet) GetByIDOutput {
	return GetByIDOutput{ID: wallet.ID}
}

type GetByID interface {
	Handle(string) (GetByIDOutput, error)
}

type getByID struct {
	walletRepository repository.Wallet
}

func NewGetByID(walletRepository repository.Wallet) GetByID {
	return &getByID{walletRepository}
}

func (uc *getByID) Handle(ID string) (GetByIDOutput, error) {
	wallet, err := uc.walletRepository.GetByID(ID)
	if err != nil {
		return GetByIDOutput{}, err
	}

	return NewGetByIDOutput(wallet), nil
}
