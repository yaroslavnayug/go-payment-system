package persistence

import (
	"github.com/jackc/pgx/pgxpool"
	"github.com/yaroslavnayug/go-payment-system/internal/app/domain/model"
)

type AccountRepository struct {
	Postgres *pgxpool.Pool
}

func (a *AccountRepository) Persist(account model.Account) error {
	return nil
}

func (a *AccountRepository) Delete(account model.Account) error {
	return nil
}
