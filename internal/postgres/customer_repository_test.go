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

func TestCreate_Find_Delete(t *testing.T) {
	t.Parallel()

	// arrange
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

	// act
	err := Repository.Create(customer)
	if err != nil {
		t.Error(err)
	}

	// assert
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

	// clean
	query := `DELETE FROM payment_system.customer WHERE generatedid = $1;`
	_, err = PostgresConnection.Exec(context.Background(), query, customer.GeneratedID)
	if err != nil {
		t.Error(err)
	}
}
