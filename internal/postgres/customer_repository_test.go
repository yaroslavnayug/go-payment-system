// +build integration

package postgres

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/log/zapadapter"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/assert"
	"github.com/yaroslavnayug/go-payment-system/internal/domain"
	"go.uber.org/zap"
)

var (
	Repository         *CustomerRepository
	PostgresConnection *pgxpool.Pool
)

const PgxLogLevel = pgx.LogLevelError

func TestMain(m *testing.M) {
	dsn := os.Getenv("POSTGRESQL_URL")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()

	logger, _ := zap.NewProduction()
	defer func() {
		_ = logger.Sync()
	}()

	pgxCfg, _ := pgx.ParseConfig(dsn)
	pgxCfg.Logger = zapadapter.NewLogger(logger)
	pgxCfg.LogLevel = PgxLogLevel
	pgxCfg.PreferSimpleProtocol = true

	pgxPoolCfg, _ := pgxpool.ParseConfig("")
	pgxPoolCfg.ConnConfig = pgxCfg
	pgxPoolCfg.MaxConns = 10
	pgxPoolCfg.MinConns = 1

	connection, err := pgxpool.ConnectConfig(ctx, pgxPoolCfg)
	if err != nil {
		panic(fmt.Errorf("unable to connect to database: %v", err))
	}
	PostgresConnection = connection
	Repository = NewCustomerRepository(connection)

	code := m.Run()
	os.Exit(code)
}

func TestCreate_Find_Update_Delete(t *testing.T) {
	t.Parallel()

	// clean
	query := `DELETE FROM payment_system.customer WHERE generatedid = $1;`
	_, err := PostgresConnection.Exec(context.Background(), query, "foobar123")
	if err != nil {
		t.Error(err)
	}

	// arrange Create
	issueDate, _ := time.Parse(domain.DateFormat, "01-01-2000")
	birthDate, _ := time.Parse(domain.DateFormat, "01-01-2020")
	customer := &domain.Customer{
		GeneratedID: "foobar123",
		FirstName:   "Bruce",
		LastName:    "Wayne",
		Email:       "goo@gmail.com",
		Phone:       "+7123",
		Address: domain.Address{
			Country:  "Russia",
			Region:   "Msk",
			City:     "Moscow",
			Street:   "Marks",
			Building: "105",
		},
		Passport: domain.Passport{
			Number:     "1234567890",
			IssueDate:  issueDate,
			Issuer:     "Foo",
			BirthDate:  birthDate,
			BirthPlace: "Nsk",
		},
	}

	// act Create
	err = Repository.Create(customer)
	if err != nil {
		t.Error(err)
	}

	// assert Create via FindByID
	dbCustomer, err := Repository.FindByID(customer.GeneratedID)
	if err != nil {
		t.Error(err)
	}
	assert.Greater(t, dbCustomer.Uid, uint64(0))
	assert.Equal(t, customer.GeneratedID, dbCustomer.GeneratedID)
	assert.Equal(t, customer.FirstName, dbCustomer.FirstName)
	assert.Equal(t, customer.LastName, dbCustomer.LastName)
	assert.Equal(t, customer.Email, dbCustomer.Email)
	assert.Equal(t, customer.Phone, dbCustomer.Phone)
	assert.Equal(t, customer.Address.Country, dbCustomer.Address.Country)
	assert.Equal(t, customer.Address.Region, dbCustomer.Address.Region)
	assert.Equal(t, customer.Address.City, dbCustomer.Address.City)
	assert.Equal(t, customer.Address.Street, dbCustomer.Address.Street)
	assert.Equal(t, customer.Address.Building, dbCustomer.Address.Building)
	assert.Equal(t, customer.Passport.Number, dbCustomer.Passport.Number)
	assert.Equal(t, customer.Passport.BirthDate, dbCustomer.Passport.BirthDate)
	assert.Equal(t, customer.Passport.BirthPlace, dbCustomer.Passport.BirthPlace)
	assert.Equal(t, customer.Passport.IssueDate, dbCustomer.Passport.IssueDate)
	assert.Equal(t, customer.Passport.Issuer, dbCustomer.Passport.Issuer)

	// arrange Update
	issueDate, _ = time.Parse(domain.DateFormat, "01-01-2020")
	birthDate, _ = time.Parse(domain.DateFormat, "01-01-2021")
	customer = &domain.Customer{
		GeneratedID: "foobar123",
		FirstName:   "Bruce_new",
		LastName:    "Wayne_new",
		Email:       "goo_new@gmail.com",
		Phone:       "+7123_new",
		Address: domain.Address{
			Country:  "Russia_new",
			Region:   "Msk_new",
			City:     "Moscow_new",
			Street:   "Marks_new",
			Building: "105_new",
		},
		Passport: domain.Passport{
			Number:     "1234567890",
			IssueDate:  issueDate,
			Issuer:     "Foo_new",
			BirthDate:  birthDate,
			BirthPlace: "Nsk_new",
		},
	}

	// act Update
	err = Repository.Update(customer)
	if err != nil {
		t.Error(err)
	}

	// assert Update via FindByPassportNumber
	updatedCustomer, err := Repository.FindByPassportNumber(customer.Passport.Number)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, customer.FirstName, updatedCustomer.FirstName)
	assert.Equal(t, customer.LastName, updatedCustomer.LastName)
	assert.Equal(t, customer.Email, updatedCustomer.Email)
	assert.Equal(t, customer.Phone, updatedCustomer.Phone)
	assert.Equal(t, customer.Address.Country, updatedCustomer.Address.Country)
	assert.Equal(t, customer.Address.Region, updatedCustomer.Address.Region)
	assert.Equal(t, customer.Address.City, updatedCustomer.Address.City)
	assert.Equal(t, customer.Address.Street, updatedCustomer.Address.Street)
	assert.Equal(t, customer.Address.Building, updatedCustomer.Address.Building)
	assert.Equal(t, customer.Passport.Number, updatedCustomer.Passport.Number)
	assert.Equal(t, customer.Passport.BirthDate, updatedCustomer.Passport.BirthDate)
	assert.Equal(t, customer.Passport.BirthPlace, updatedCustomer.Passport.BirthPlace)
	assert.Equal(t, customer.Passport.IssueDate, updatedCustomer.Passport.IssueDate)
	assert.Equal(t, customer.Passport.Issuer, updatedCustomer.Passport.Issuer)

	// act Delete
	err = Repository.Delete(updatedCustomer.GeneratedID)
	if err != nil {
		t.Error(err)
	}

	// assert Delete via FindByID
	customer, err = Repository.FindByID(updatedCustomer.GeneratedID)
	if err != nil {
		t.Error(err)
	}
	assert.Nil(t, customer)
}
