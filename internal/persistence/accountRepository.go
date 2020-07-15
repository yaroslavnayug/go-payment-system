package persistence

import (
	"context"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/yaroslavnayug/go-payment-system/internal/domain/model"
)

type AccountRepository struct {
	pgConn *pgxpool.Pool
}

func NewAccountRepository(pgConn *pgxpool.Pool) *AccountRepository {
	return &AccountRepository{pgConn: pgConn}
}

func (a *AccountRepository) CreateAccount(account model.Account) (accountID uint64, err error) {
	query := `
		INSERT INTO
			payment_system.account (firstname, lastname, passportdata, phone, country, region, city, street)
        VALUES
			($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id;`
	err = a.pgConn.QueryRow(
		context.Background(),
		query,
		account.FirstName,
		account.LastName,
		account.PassportData,
		account.Phone,
		account.Country,
		account.Region,
		account.City,
		account.Street,
	).Scan(&accountID)

	if pgErr, ok := err.(interface{ SQLState() string }); ok {
		if pgErr.SQLState() == pgerrcode.UniqueViolation {
			return 0, model.NewValidationError("account with such passport_data already exist")
		}
	}
	if err != nil {
		return 0, err
	}
	return accountID, nil
}

func (a *AccountRepository) GetAccountByPassportData(passportData string) (*model.Account, error) {
	query :=
		`SELECT
			id,
			firstname, 
			lastname, 
			passportdata, 
			phone, 
			country,
			region,
			city,
			street
		FROM
			payment_system.account
		WHERE
			passportdata=$1;`
	account := &model.Account{}
	err := a.pgConn.QueryRow(context.Background(), query, passportData).
		Scan(
			&account.Id,
			&account.FirstName,
			&account.LastName,
			&account.PassportData,
			&account.Phone,
			&account.Country,
			&account.Region,
			&account.City,
			&account.Street,
		)
	if err == pgx.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return account, nil
}
