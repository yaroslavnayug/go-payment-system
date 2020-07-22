package service

import (
	"github.com/yaroslavnayug/go-payment-system/internal/domain/converter"
	"github.com/yaroslavnayug/go-payment-system/internal/domain/model"
	"github.com/yaroslavnayug/go-payment-system/internal/domain/validator"
)

type AccountService struct {
	validator  validator.Validator
	converter  converter.Converter
	repository model.Repository
}

func NewAccountService(validator validator.Validator, converter converter.Converter, repository model.Repository) *AccountService {
	return &AccountService{validator: validator, converter: converter, repository: repository}
}

func (a *AccountService) CreateAccount(requestData interface{}) (accountID uint64, err error) {
	err = a.validator.ValidateCreateAccount(requestData)
	if err != nil {
		return 0, err
	}

	account, err := a.converter.ConvertToAccount(requestData)
	if err != nil {
		return 0, err
	}

	accountID, err = a.repository.CreateAccount(account)
	if err != nil {
		return 0, err
	}
	return accountID, nil
}
