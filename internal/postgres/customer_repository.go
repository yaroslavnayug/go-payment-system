package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/yaroslavnayug/go-payment-system/internal/domain"
)

const tableName = "payment_system.customer"

type CustomerRepository struct {
	pgConn *pgxpool.Pool
}

func NewCustomerRepository(pgConn *pgxpool.Pool) *CustomerRepository {
	return &CustomerRepository{pgConn: pgConn}
}

func (a *CustomerRepository) Create(customer *domain.Customer) error {
	query := fmt.Sprintf(`
		INSERT INTO
			%s (
				generatedid, firstname, lastname, email, phone, country, region, city, street, building,
				passportnumber, passportissuedate, passportissuer, birthdate, birthplace
		)
        VALUES
			($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15);`, tableName)
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
	query := fmt.Sprintf(`
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
			%s
		WHERE
			generatedid=$1;`, tableName)

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
	if err == pgx.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (a *CustomerRepository) FindByPassportNumber(passportNumber string) (customer *domain.Customer, err error) {
	query := fmt.Sprintf(`
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
			%s
		WHERE
			passportnumber=$1;`, tableName)

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
	query := fmt.Sprintf(`
		UPDATE
			%s
		SET (
				generatedid, firstname, lastname, email, phone, country, region, city, street, building,
				passportnumber, passportissuedate, passportissuer, birthdate, birthplace
		) = ROW
			($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15);`, tableName)
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
	query := fmt.Sprintf(`
		DELETE FROM
			%s
		WHERE 
			generatedid = $1;`, tableName)
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
