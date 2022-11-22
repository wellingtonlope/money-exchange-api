package memory

import (
	"github.com/google/uuid"
	"github.com/wellingtonlope/money-exchange-api/wallet/app/repository"
	"github.com/wellingtonlope/money-exchange-api/wallet/domain"
)

type wallet struct {
	wallets []domain.Wallet
}

func NewWallet() repository.Wallet {
	return &wallet{}
}

func (r *wallet) Create(wallet domain.Wallet) (domain.Wallet, error) {
	wallet.ID = uuid.NewString()
	r.wallets = append(r.wallets, wallet)
	return wallet, nil
}

func (r *wallet) GetByID(ID string) (domain.Wallet, error) {
	for _, wallet := range r.wallets {
		if ID == wallet.ID {
			return wallet, nil
		}
	}
	return domain.Wallet{}, repository.ErrRepositoryWalletNotFound
}
