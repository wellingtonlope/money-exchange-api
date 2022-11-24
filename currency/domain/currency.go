package domain

import "errors"

var (
	ErrCurrencyIDInvalid = errors.New("currency ID is invalid")
)

type Currency struct {
	ID          string
	Description string
	Country     string
}

func NewCurrency(ID, description, country string) (Currency, error) {
	if len(ID) < 3 {
		return Currency{}, ErrCurrencyIDInvalid
	}
	return Currency{ID, description, country}, nil
}

func (c *Currency) Update(description, country string) {
	c.Description = description
	c.Country = country
}
