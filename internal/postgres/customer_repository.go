package postgres

import (
	"context"

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

	if err != nil {
		return err
	}
	return nil
}

func (a *CustomerRepository) FindByID(customerID string) (customer *domain.Customer, err error) {
	query := `
		SELECT
			uid,
			generatedid,
			firstname,
			lastname,
			email,
			phone,
			country,
			region,
			city,
			street,
			building,
			passportnumber,
			passportissuer,
			passportissuedate,
			birthdate,
			birthplace
		FROM
			payment_system.customer
		WHERE
			generatedid=$1;`

	customer = &domain.Customer{}
	queryRow := a.pgConn.QueryRow(
		context.Background(),
		query,
		customerID,
	)
	err = queryRow.Scan(
		&customer.Uid,
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
		&customer.Passport.Issuer,
		&customer.Passport.IssueDate,
		&customer.Passport.BirthDate,
		&customer.Passport.BirthPlace,
	)

	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (a *CustomerRepository) FindByPassportNumber(passportNumber string) (customer *domain.Customer, err error) {
	query := `
		SELECT
			uid,
			generatedid,
			firstname,
			lastname,
			email,
			phone,
			country,
			region,
			city,
			street,
			building,
			passportnumber,
			passportissuer,
			passportissuedate,
			birthdate,
			birthplace
		FROM
			payment_system.customer
		WHERE
			passportnumber=$1;`

	customer = &domain.Customer{}
	queryRow := a.pgConn.QueryRow(
		context.Background(),
		query,
		passportNumber,
	)
	err = queryRow.Scan(
		&customer.Uid,
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
		&customer.Passport.Issuer,
		&customer.Passport.IssueDate,
		&customer.Passport.BirthDate,
		&customer.Passport.BirthPlace,
	)

	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (a *CustomerRepository) Update(customer *domain.Customer) error {
	return nil
}

func (a *CustomerRepository) Delete(customerID string) error {
	return nil
}
