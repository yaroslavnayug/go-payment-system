package domain

//go:generate mockgen -destination=../postgres/mocks/customer_repository_mock.go -package=mocks . CustomerRepository
type CustomerRepository interface {
	Create(customer *Customer) error
	Find(customerID string) (customer *Customer, err error)
	Update(customer *Customer) error
	Delete(customerID string) error
}

type Customer struct {
	Uid         uint64
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
	IssueDate  string
	Issuer     string
	BirthDate  string
	BirthPlace string
}
