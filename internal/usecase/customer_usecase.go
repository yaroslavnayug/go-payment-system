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

func (c *CustomerUseCase) Create(customer *domain.Customer) error {
	customerExist, err := c.repo.FindByPassportNumber(customer.Passport.Number)
	if err != nil {
		return err
	}
	if customerExist != nil {
		return domain.NewValidationError("customer with such passport number already exist")
	}

	uniqueCustomerID, err := hash.GenerateUniqueCustomerID(
		customer.FirstName,
		customer.Passport.Number,
		time.Now().Unix(),
	)
	if err != nil {
		return err
	}

	customer.GeneratedID = uniqueCustomerID
	err = c.repo.Create(customer)
	if err != nil {
		return err
	}
	return nil
}

func (c *CustomerUseCase) Find(customerID string) (*domain.Customer, error) {
	customer, err := c.repo.FindByID(customerID)
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (c *CustomerUseCase) Update(customer *domain.Customer, customerID string) error {
	existingCustomer, err := c.repo.FindByID(customerID)
	if err != nil {
		return err
	}
	if existingCustomer == nil {
		return domain.NewValidationError("customer with such id not found")
	}
	err = c.repo.Update(customer)
	if err != nil {
		return err
	}
	return nil
}

func (c *CustomerUseCase) Delete(customerID string) error {
	err := c.repo.Delete(customerID)
	if err != nil {
		return err
	}
	return nil
}
