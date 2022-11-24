package memory

import (
	"github.com/wellingtonlope/money-exchange-api/currency/app/repository"
	"github.com/wellingtonlope/money-exchange-api/currency/domain"
)

type currency struct {
	currencies []domain.Currency
}

func NewCurrency() repository.Currency {
	return &currency{}
}

func (r *currency) Create(currency domain.Currency) (domain.Currency, error) {
	_, err := r.GetByID(currency.ID)
	if err != nil && err != repository.ErrRepositoryCurrencyNotFound {
		return domain.Currency{}, err
	}
	if err == nil {
		return domain.Currency{}, repository.ErrRepositoryCurrencyDuplicated
	}
	r.currencies = append(r.currencies, currency)
	return currency, nil
}

func (r *currency) Update(currency domain.Currency) error {
	for i, cur := range r.currencies {
		if currency.ID == cur.ID {
			r.currencies[i] = currency
			return nil
		}
	}
	return repository.ErrRepositoryCurrencyNotFound
}

func (r *currency) GetByID(ID string) (domain.Currency, error) {
	for _, cur := range r.currencies {
		if ID == cur.ID {
			return cur, nil
		}
	}
	return domain.Currency{}, repository.ErrRepositoryCurrencyNotFound
}
