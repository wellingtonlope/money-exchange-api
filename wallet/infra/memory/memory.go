package memory

import "github.com/wellingtonlope/money-exchange-api/wallet/app/repository"

type repositories struct{}

func NewRepositories() repository.Repositories {
	return &repositories{}
}

func (r *repositories) GetAll() (*repository.AllRepositories, error) {
	return &repository.AllRepositories{
		Wallet: NewWallet(),
	}, nil
}
