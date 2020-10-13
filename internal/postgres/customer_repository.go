package postgres

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/yaroslavnayug/go-payment-system/internal/domain"
)

const tableName = "customer"

var customerColumns = []string{
	"uid",
	"firstname",
	"lastname",
	"email",
	"phone",
	"country",
	"region",
	"city",
	"street",
	"building",
	"passportnumber",
	"passportissuedate",
	"passportissuer",
	"birthdate",
	"birthplace",
}

var preparedCustomerColumns = strings.Join(customerColumns, ", ")

type CustomerRepository struct {
	pgConn *pgxpool.Pool
}

func NewCustomerRepository(pgConn *pgxpool.Pool) *CustomerRepository {
	return &CustomerRepository{pgConn: pgConn}
}

func (a *CustomerRepository) Create(customer *domain.Customer) error {
	query := fmt.Sprintf(
		`INSERT INTO %s (%s) VALUES (%s);`,
		tableName,
		preparedCustomerColumns,
		getSubstitutionVerbsForColumns(customerColumns),
	)
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

	if err != nil {
		return err
	}
	return nil
}

func (a *CustomerRepository) FindByID(customerID string) (customer *domain.Customer, err error) {
	query := fmt.Sprintf(
		`SELECT %s FROM %s WHERE uid=$1;`,
		preparedCustomerColumns,
		tableName,
	)

	customer = &domain.Customer{}
	queryRow := a.pgConn.QueryRow(
		context.Background(),
		query,
		customerID,
	)
	err = queryRow.Scan(
		&customer.GeneratedID,
		&customer.FirstName,
		&customer.LastName,
		&customer.Email,
		&customer.Phone,
		&customer.Address.Country,
		&customer.Address.Region,
		&customer.Address.City,
		&customer.Address.Street,
		&customer.Address.Building,
		&customer.Passport.Number,
		&customer.Passport.IssueDate,
		&customer.Passport.Issuer,
		&customer.Passport.BirthDate,
		&customer.Passport.BirthPlace,
	)
	if err == pgx.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (a *CustomerRepository) FindByPassportNumber(passportNumber string) (customer *domain.Customer, err error) {
	query := fmt.Sprintf(
		`SELECT %s FROM %s WHERE	passportnumber=$1;`,
		preparedCustomerColumns,
		tableName,
	)

	customer = &domain.Customer{}
	queryRow := a.pgConn.QueryRow(
		context.Background(),
		query,
		passportNumber,
	)
	err = queryRow.Scan(
		&customer.GeneratedID,
		&customer.FirstName,
		&customer.LastName,
		&customer.Email,
		&customer.Phone,
		&customer.Address.Country,
		&customer.Address.Region,
		&customer.Address.City,
		&customer.Address.Street,
		&customer.Address.Building,
		&customer.Passport.Number,
		&customer.Passport.IssueDate,
		&customer.Passport.Issuer,
		&customer.Passport.BirthDate,
		&customer.Passport.BirthPlace,
	)

	if err == pgx.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (a *CustomerRepository) Update(customer *domain.Customer) error {
	query := fmt.Sprintf(
		`UPDATE %s SET (%s) = ROW (%s) WHERE uid='%s';`,
		tableName,
		preparedCustomerColumns,
		getSubstitutionVerbsForColumns(customerColumns),
		customer.GeneratedID,
	)
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

	if err != nil {
		return err
	}
	return nil
}

func (a *CustomerRepository) Delete(customerID string) error {
	query := fmt.Sprintf(
		`DELETE FROM	%s WHERE uid = $1;`,
		tableName,
	)
	_, err := a.pgConn.Exec(
		context.Background(),
		query,
		customerID,
	)

	if err != nil {
		return err
	}
	return nil
}
