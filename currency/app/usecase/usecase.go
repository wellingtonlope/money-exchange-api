package usecase

import "github.com/wellingtonlope/money-exchange-api/currency/domain"

type Output struct {
	ID          string
	Description string
	Country     string
}

func NewOutput(currency domain.Currency) Output {
	return Output{currency.ID, currency.Description, currency.Country}
}
