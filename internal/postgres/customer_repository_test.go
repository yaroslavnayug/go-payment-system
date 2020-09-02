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

func TestCreate(t *testing.T) {
	t.Parallel()

	// arrange
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()

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
			IssueDate:  "01-01-2000",
			Issuer:     "Foo",
			BirthDate:  "01-01-2000",
			BirthPlace: "Nsk",
		},
	}

	// act
	err := Repository.Create(customer)
	if err != nil {
		t.Error(err)
	}

	// assert
	query := `
		SELECT COUNT(*) FROM
			payment_system.customer
	   WHERE
			generatedid=$1
		AND
			firstname=$2
		AND
			lastname=$3
		AND
			email=$4
		AND
			phone=$5
		AND
			country=$6
		AND
			region=$7
		AND
			city=$8
		AND
			street=$9
		AND
			building=$10
		AND
			passportnumber=$11
		AND
			passportissuer=$12
		AND
			passportissuedate=$13
		AND
			birthdate=$14
		AND
			birthplace=$15;`
	var countRows uint64
	result := PostgresConnection.QueryRow(
		ctx,
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
		customer.Passport.Issuer,
		customer.Passport.IssueDate,
		customer.Passport.BirthDate,
		customer.Passport.BirthPlace,
	)
	err = result.Scan(&countRows)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, uint64(1), countRows)

	// clean
	query = `DELETE FROM payment_system.customer WHERE generatedid = $1;`
	_, err = PostgresConnection.Exec(context.Background(), query, customer.GeneratedID)
	if err != nil {
		t.Error(err)
	}
}

func TestCreateAccount_AccountAlreadyExist(t *testing.T) {
	t.Parallel()

	// arrange
	customer := &domain.Customer{
		GeneratedID: "foobar123456",
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
			Number:     "1234567899",
			IssueDate:  "01-01-2000",
			Issuer:     "Foo",
			BirthDate:  "01-01-2000",
			BirthPlace: "Nsk",
		},
	}

	query := `
		INSERT INTO
			payment_system.customer
			(generatedid, firstname, lastname, phone, country, region, city, street, building, passportnumber, passportissuedate, passportissuer, birthdate, birthplace)
     VALUES
			($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14);`
	_, err := PostgresConnection.Exec(
		context.Background(),
		query,
		customer.GeneratedID,
		customer.FirstName,
		customer.LastName,
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
		t.Error(err)
	}

	// act
	err = Repository.Create(customer)

	// assert
	assert.IsType(t, new(domain.ValidationError), err)
	assert.Equal(t, "customer with such generated id or passport number already exist", err.Error())

	// clean
	query = `DELETE FROM payment_system.customer WHERE generatedid = $1;`
	_, _ = PostgresConnection.Exec(context.Background(), query, customer.GeneratedID)
}
