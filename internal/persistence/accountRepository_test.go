// +build integration

package persistence

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/assert"
	"github.com/yaroslavnayug/go-payment-system/internal/domain/model"
)

var (
	Repository         *AccountRepository
	PostgresConnection *pgxpool.Pool
)

func TestMain(m *testing.M) {
	dsn := os.Getenv("POSTGRESQL_URL")
	connection, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		panic(fmt.Errorf("unable to connect to database: %v", err))
	}
	PostgresConnection = connection
	Repository = NewAccountRepository(connection)

	code := m.Run()
	os.Exit(code)
}

func TestGetAccountByPassportData_NonExistentAccount(t *testing.T) {
	t.Parallel()

	// act
	account, _ := Repository.GetAccountByPassportData("1234567890")

	// assert
	assert.Nil(t, account)
}

func TestGetAccountByPassportData_ExistentAccount(t *testing.T) {
	t.Parallel()

	// arrange
	query := `
		INSERT INTO
			payment_system.account
			(firstname, lastname, passportdata, phone, country, region, city, street)
      VALUES
			('eafirstname', 'ealastname', '1234567890', '+7953499', 'USA', 'California', 'New-York', 'Wall');`
	_, _ = PostgresConnection.Exec(context.Background(), query)

	// act
	account, _ := Repository.GetAccountByPassportData("1234567890")

	// assert
	assert.NotNil(t, account)
	assert.Greater(t, account.Id, uint64(0))
	assert.Equal(t, "eafirstname", account.FirstName)
	assert.Equal(t, "ealastname", account.LastName)
	assert.Equal(t, "+7953499", account.Phone)
	assert.Equal(t, "USA", account.Country)
	assert.Equal(t, "California", account.Region)
	assert.Equal(t, "New-York", account.City)
	assert.Equal(t, "Wall", account.Street)

	// clean
	query = `DELETE FROM payment_system.account WHERE id = $1;`
	_, _ = PostgresConnection.Exec(context.Background(), query, account.Id)
}

func TestCreateAccount(t *testing.T) {
	t.Parallel()

	// arrange
	account := model.Account{
		FirstName:    "TestCreate",
		LastName:     "TestCreate",
		PassportData: "9876543210",
		Phone:        "+35555",
		Country:      "Spain",
		Region:       "Madrid",
		City:         "Porto",
		Street:       "Jurutuba",
	}

	// act
	accountID, _ := Repository.CreateAccount(account)

	// assert
	query := `
		SELECT COUNT(*) FROM
			payment_system.account
	    WHERE
			id=$1
		AND
			firstname=$2
		AND
			lastname=$3
		AND
			passportdata=$4
		AND
			phone=$5
		AND
			country=$6
		AND
			region=$7
		AND
			city=$8
		AND
			street=$9;`
	var countRows uint64
	_ = PostgresConnection.QueryRow(
		context.Background(),
		query,
		accountID,
		account.FirstName,
		account.LastName,
		account.PassportData,
		account.Phone,
		account.Country,
		account.Region,
		account.City,
		account.Street,
	).Scan(&countRows)

	assert.Equal(t, uint64(1), countRows)

	// clean
	query = `DELETE FROM payment_system.account WHERE passportdata = $1;`
	_, _ = PostgresConnection.Exec(context.Background(), query, account.PassportData)
}

func TestCreateAccount_AccountAlreadyExist(t *testing.T) {
	t.Parallel()

	// arrange
	account := model.Account{
		FirstName:    "AccountAlreadyExist",
		LastName:     "AccountAlreadyExist",
		PassportData: "6666666666",
		Phone:        "+35555",
		Country:      "Spain",
	}

	query := `
		INSERT INTO
			payment_system.account
			(firstname, lastname, passportdata, phone, country)
      VALUES
			($1, $2, $3, $4, $5);`
	_, _ = PostgresConnection.Exec(context.Background(), query, account.FirstName, account.LastName, account.PassportData, account.Phone, account.Country)

	// act
	_, err := Repository.CreateAccount(account)

	// assert
	assert.IsType(t, new(model.ValidationError), err)
	assert.Equal(t, "account with such passport_data already exist", err.Error())

	// clean
	query = `DELETE FROM payment_system.account WHERE passportdata = $1;`
	_, _ = PostgresConnection.Exec(context.Background(), query, "6666666666")
}
