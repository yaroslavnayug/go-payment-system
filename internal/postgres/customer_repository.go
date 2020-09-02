package postgres

import (
	"context"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/yaroslavnayug/go-payment-system/internal/domain"
)

type CustomerRepository struct {
	pgConn *pgxpool.Pool
}

func NewCustomerRepository(pgConn *pgxpool.Pool) *CustomerRepository {
	return &CustomerRepository{pgConn: pgConn}
}

func (a *CustomerRepository) Create(customer *domain.Customer) error {
	query := `
		INSERT INTO
			payment_system.customer (
				generatedid, firstname, lastname, email, phone, country, region, city, street, building,
				passportnumber, passportissuedate, passportissuer, birthdate, birthplace
		)
        VALUES
			($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15);`
	_, err := a.pgConn.Exec(
		context.Background(),
		query,
		customer.GeneratedID,
		customer.FirstName,
		customer.LastName,
		customer.Email,
		customer.Phone,
		customer.Address.Country,
		customer.Address.Region,
		customer.Address.City,
		customer.Address.Street,
		customer.Address.Building,
		customer.Passport.Number,
		customer.Passport.IssueDate,
		customer.Passport.Issuer,
		customer.Passport.BirthDate,
		customer.Passport.BirthPlace,
	)

	if pgErr, ok := err.(interface{ SQLState() string }); ok {
		if pgErr.SQLState() == pgerrcode.UniqueViolation {
			return domain.NewValidationError("customer with such generated id or passport number already exist")
		}
	}
	if err != nil {
		return err
	}
	return nil
}

func (a *CustomerRepository) Find(customerID string) (customer *domain.Customer, err error) {
	return nil, nil
}

func (a *CustomerRepository) Update(customer *domain.Customer) error {
	return nil
}

func (a *CustomerRepository) Delete(customerID string) error {
	return nil
}
