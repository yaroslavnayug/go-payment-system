package domain

import "time"

//go:generate mockgen -destination=../postgres/mocks/customer_repository_mock.go -package=mocks . CustomerRepository

type CustomerRepository interface {
	Create(customer *Customer) error
	FindByID(customerID string) (customer *Customer, err error)
	FindByPassportNumber(passportNumber string) (customer *Customer, err error)
	Update(customer *Customer) error
	Delete(customerID string) error
}

type Customer struct {
	GeneratedID string
	FirstName   string
	LastName    string
	Email       string
	Phone       string
	Address     Address
	Passport    Passport
}

type Address struct {
	Country  string
	Region   string
	City     string
	Street   string
	Building string
}

type Passport struct {
	Number     string
	IssueDate  time.Time
	Issuer     string
	BirthDate  time.Time
	BirthPlace string
}
