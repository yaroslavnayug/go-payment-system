package model

//go:generate mockgen -destination=../../persistence/mocks/accountRepositoryMock.go -package=mocks . AccountRepositoryInterface
type AccountRepositoryInterface interface {
	CreateAccount(Account) (accountID uint64, err error)
	GetAccountByPassportData(passportData string) (*Account, error)
}
