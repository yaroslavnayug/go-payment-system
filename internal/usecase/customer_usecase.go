package usecase

import (
	"time"

	"github.com/yaroslavnayug/go-payment-system/internal/domain"
	"github.com/yaroslavnayug/go-payment-system/internal/hash"
)

type CustomerUseCase struct {
	repo domain.CustomerRepository
}

func NewCustomerUseCase(repo domain.CustomerRepository) *CustomerUseCase {
	return &CustomerUseCase{repo: repo}
}

func (a *CustomerUseCase) Create(customer *domain.Customer) (customerID string, err error) {
	uniqueCustomerID, err := hash.GenerateUniqueCustomerID(
		customer.FirstName,
		customer.Passport.Number,
		time.Now().Unix(),
	)
	if err != nil {
		return "", err
	}

	customer.GeneratedID = uniqueCustomerID
	err = a.repo.Create(customer)
	if err != nil {
		return "", err
	}
	return uniqueCustomerID, nil
}
