package usecase

import (
	"github.com/wellingtonlope/money-exchange-api/currency/app/repository"
	"github.com/wellingtonlope/money-exchange-api/currency/domain"
)

type CreateInput struct {
	ID          string
	Description string
	Country     string
}

type Create interface {
	Handle(CreateInput) (Output, error)
}

type create struct {
	currencyRepository repository.Currency
}

func NewCreate(currencyRepository repository.Currency) Create {
	return &create{currencyRepository}
}

func (uc *create) Handle(input CreateInput) (Output, error) {
	currency, err := domain.NewCurrency(input.ID, input.Description, input.Country)
	if err != nil {
		return Output{}, err
	}

	currency, err = uc.currencyRepository.Create(currency)
	if err != nil {
		return Output{}, err
	}

	return NewOutput(currency), nil
}
