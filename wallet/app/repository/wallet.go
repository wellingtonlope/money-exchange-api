package repository

import (
	"errors"

	"github.com/wellingtonlope/money-exchange-api/wallet/domain"
)

var (
	ErrRepositoryWalletNotFound = errors.New("wallet not found")
)

type Wallet interface {
	Create(domain.Wallet) (domain.Wallet, error)
	GetByID(ID string) (domain.Wallet, error)
}
