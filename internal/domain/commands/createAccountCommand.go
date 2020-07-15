package commands

import (
	"strings"

	"github.com/yaroslavnayug/go-payment-system/internal/api"
	"github.com/yaroslavnayug/go-payment-system/internal/domain/model"
)

type CreateAccountCommand struct {
	Account model.Account
}

func NewCreateAccountCommand(request api.CreateAccountRequest) (*CreateAccountCommand, error) {
	if request.FirstName == "" {
		return nil, model.NewValidationError("first_name is mandatory field")
	}
	if request.LastName == "" {
		return nil, model.NewValidationError("last_name is mandatory field")
	}
	if request.PassportData == "" {
		return nil, model.NewValidationError("passport_data is mandatory field")
	}
	if request.Phone == "" {
		return nil, model.NewValidationError("phone is mandatory field")
	}

	passportData := strings.Replace(request.PassportData, " ", "", -1)
	if len(passportData) < 10 {
		return nil, model.NewValidationError("passport_data should be at least 10 characters long")
	}

	account := model.Account{
		FirstName:    request.FirstName,
		LastName:     request.LastName,
		PassportData: passportData,
		Phone:        request.Phone,
		Country:      request.Address.Country,
		Region:       request.Address.Region,
		City:         request.Address.City,
		Street:       request.Address.Street,
	}
	return &CreateAccountCommand{Account: account}, nil
}
