package usecase

import (
	"github.com/wellingtonlope/money-exchange-api/currency/app/repository"
)

type GetByID interface {
	Handle(string) (Output, error)
}

type getByID struct {
	currencyRepository repository.Currency
}

func NewGetByID(currencyRepository repository.Currency) GetByID {
	return &getByID{currencyRepository}
}

func (uc *getByID) Handle(ID string) (Output, error) {
	currency, err := uc.currencyRepository.GetByID(ID)
	if err != nil {
		return Output{}, err
	}

	return NewOutput(currency), nil
}
