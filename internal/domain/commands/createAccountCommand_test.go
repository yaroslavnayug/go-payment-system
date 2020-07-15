package commands

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yaroslavnayug/go-payment-system/internal/api"
)

func TestNewCreateAccountCommand_EmptyFirstName(t *testing.T) {
	request := api.CreateAccountRequest{}

	_, err := NewCreateAccountCommand(request)
	assert.Equal(t, "first_name is mandatory field", err.Error())
}

func TestNewCreateAccountCommand_EmptyLastName(t *testing.T) {
	request := api.CreateAccountRequest{
		FirstName: "TestName",
	}

	_, err := NewCreateAccountCommand(request)
	assert.Equal(t, "last_name is mandatory field", err.Error())
}

func TestNewCreateAccountCommand_EmptyPassportData(t *testing.T) {
	request := api.CreateAccountRequest{
		FirstName: "TestName",
		LastName:  "TestName",
	}

	_, err := NewCreateAccountCommand(request)
	assert.Equal(t, "passport_data is mandatory field", err.Error())
}

func TestNewCreateAccountCommand_EmptyPhone(t *testing.T) {
	request := api.CreateAccountRequest{
		FirstName:    "TestName",
		LastName:     "TestName",
		PassportData: "12345",
	}

	_, err := NewCreateAccountCommand(request)
	assert.Equal(t, "phone is mandatory field", err.Error())
}

func TestNewCreateAccountCommand_PassportLengthIsLessThan10(t *testing.T) {
	request := api.CreateAccountRequest{
		FirstName:    "TestName",
		LastName:     "TestName",
		PassportData: "1 2 3 4 5 6 7 8 9",
		Phone:        "+543232",
	}

	_, err := NewCreateAccountCommand(request)
	assert.Equal(t, "passport_data should be at least 10 characters long", err.Error())
}

func TestNewCreateAccountCommand_AddressCanBeEmpty(t *testing.T) {
	request := api.CreateAccountRequest{
		FirstName:    "TestFirstName",
		LastName:     "TestLastName",
		PassportData: "1 2 3 4 5 6 7 8 9 0",
		Phone:        "+543232",
	}

	command, err := NewCreateAccountCommand(request)

	assert.Nil(t, err)
	assert.Equal(t, "TestFirstName", command.Account.FirstName)
	assert.Equal(t, "TestLastName", command.Account.LastName)
	assert.Equal(t, "1234567890", command.Account.PassportData)
	assert.Equal(t, "+543232", command.Account.Phone)
	assert.Equal(t, "", command.Account.Country)
	assert.Equal(t, "", command.Account.Region)
	assert.Equal(t, "", command.Account.City)
	assert.Equal(t, "", command.Account.Street)
}

func TestNewCreateAccountCommand_AllDataFilled(t *testing.T) {
	request := api.CreateAccountRequest{
		FirstName:    "TestFirstName",
		LastName:     "TestLastName",
		PassportData: "1 2 3 4 5 6 7 8 9 0",
		Phone:        "+543232",
		Address:      api.Address{Country: "Russia", Region: "NSO", City: "NSK", Street: "Red Square"},
	}

	command, err := NewCreateAccountCommand(request)

	assert.Nil(t, err)
	assert.Equal(t, "TestFirstName", command.Account.FirstName)
	assert.Equal(t, "TestLastName", command.Account.LastName)
	assert.Equal(t, "1234567890", command.Account.PassportData)
	assert.Equal(t, "+543232", command.Account.Phone)
	assert.Equal(t, "Russia", command.Account.Country)
	assert.Equal(t, "NSO", command.Account.Region)
	assert.Equal(t, "NSK", command.Account.City)
	assert.Equal(t, "Red Square", command.Account.Street)
}
