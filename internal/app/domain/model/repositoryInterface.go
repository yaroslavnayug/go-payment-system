package model

type AccountRepositoryInterface interface {
	Persist(account Account) error
	Delete(account Account) error
}
