package model

//go:generate mockgen -destination=../../persistence/mocks/repositoryMock.go -package=mocks . RepositoryInterface
type Repository interface {
	CreateAccount(*Account) (accountID uint64, err error)
	GetAccountByPassportData(passportData string) (*Account, error)
}
