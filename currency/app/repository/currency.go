package repository

import (
	"errors"
	"github.com/wellingtonlope/money-exchange-api/currency/domain"
)

var (
	ErrRepositoryCurrencyNotFound   = errors.New("currency not found")
	ErrRepositoryCurrencyDuplicated = errors.New("currency already exists")
)

type Currency interface {
	Create(domain.Currency) (domain.Currency, error)
	Update(domain.Currency) error
	GetByID(string) (domain.Currency, error)
}
