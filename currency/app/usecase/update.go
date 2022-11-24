package usecase

import "github.com/wellingtonlope/money-exchange-api/currency/app/repository"

type UpdateInput struct {
	ID          string
	Description string
	Country     string
}

type Update interface {
	Handle(UpdateInput) (Output, error)
}

type update struct {
	currencyRepository repository.Currency
}

func NewUpdate(currencyRepository repository.Currency) Update {
	return &update{currencyRepository}
}

func (uc *update) Handle(input UpdateInput) (Output, error) {
	currency, err := uc.currencyRepository.GetByID(input.ID)
	if err != nil {
		return Output{}, err
	}

	currency.Update(input.Description, input.Country)

	err = uc.currencyRepository.Update(currency)
	if err != nil {
		return Output{}, err
	}

	return NewOutput(currency), nil
}
