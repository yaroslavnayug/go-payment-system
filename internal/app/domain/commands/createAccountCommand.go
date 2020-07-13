package commands

import (
	"errors"
	"strings"

	"github.com/yaroslavnayug/go-payment-system/internal/app/api/requests"
	"github.com/yaroslavnayug/go-payment-system/internal/app/domain/model"
)

type CreateAccountCommand struct {
	account model.Account
}

func NewCreateAccountCommand(request requests.CreateAccountRequest) (CreateAccountCommand, error) {
	command := CreateAccountCommand{}
	if request.FirstName == "" {
		return command, errors.New("first_name is mandatory field")
	}
	if request.LastName == "" {
		return command, errors.New("last_name is mandatory field")
	}
	if request.PassportData == "" {
		return command, errors.New("passport_data is mandatory field")
	}
	if request.Phone == "" {
		return command, errors.New("phone is mandatory field")
	}

	passportData := strings.Replace(request.PassportData, "", "", -1)
	if len(passportData) < 10 {
		return command, errors.New("passport_data should be at least 10 characters long")
	}

	account := model.Account{
		FirstName:    request.FirstName,
		LastName:     request.LastName,
		PassportData: request.PassportData,
		Phone:        request.Phone,
		Address: model.Address{
			Country: request.Address.Country,
			State:   request.Address.State,
			City:    request.Address.City,
			Street:  request.Address.Street,
		},
	}
	return CreateAccountCommand{account: account}, nil
}
