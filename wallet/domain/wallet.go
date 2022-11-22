package domain

type Wallet struct {
	ID string
}

func NewWallet() Wallet {
	return Wallet{}
}
