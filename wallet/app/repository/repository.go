package repository

type AllRepositories struct {
	Wallet Wallet
}

type Repositories interface {
	GetAll() (*AllRepositories, error)
}
