package persistence

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/yaroslavnayug/go-payment-system/internal/app/domain/model"
)

type AccountRepository struct {
	pgConn *pgxpool.Pool
}

func NewAccountRepository(pgConn *pgxpool.Pool) *AccountRepository {
	return &AccountRepository{pgConn: pgConn}
}

func (a *AccountRepository) Persist(account model.Account) error {
	return nil
}

func (a *AccountRepository) Delete(account model.Account) error {
	return nil
}
