package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yaroslavnayug/go-payment-system/internal/api"
)

func TestJSONRequestValidator_WrongRequestType(t *testing.T) {
	validator := NewJSONRequestValidator()

	err := validator.ValidateCreateAccount("foo")

	assert.EqualError(t, err, "invalid request type sent to JSONRequestValidator: expected CreateAccountRequest, got foo")
}

func TestJSONRequestValidator_EmptyFirstName(t *testing.T) {
	request := api.CreateAccountRequest{}
	validator := NewJSONRequestValidator()

	err := validator.ValidateCreateAccount(request)

	assert.Equal(t, "first_name is mandatory field", err.Error())
}

func TestJSONRequestValidator_EmptyLastName(t *testing.T) {
	request := api.CreateAccountRequest{
		FirstName: "TestName",
	}
	validator := NewJSONRequestValidator()

	err := validator.ValidateCreateAccount(request)

	assert.Equal(t, "last_name is mandatory field", err.Error())
}

func TestJSONRequestValidator_EmptyPassportData(t *testing.T) {
	request := api.CreateAccountRequest{
		FirstName: "TestName",
		LastName:  "TestName",
	}
	validator := NewJSONRequestValidator()

	err := validator.ValidateCreateAccount(request)

	assert.Equal(t, "passport_data is mandatory field", err.Error())
}

func TestJSONRequestValidator_EmptyPhone(t *testing.T) {
	request := api.CreateAccountRequest{
		FirstName:    "TestName",
		LastName:     "TestName",
		PassportData: "12345",
	}
	validator := NewJSONRequestValidator()

	err := validator.ValidateCreateAccount(request)

	assert.Equal(t, "phone is mandatory field", err.Error())
}

func TestJSONRequestValidator_PassportLengthIsLessThan10(t *testing.T) {
	request := api.CreateAccountRequest{
		FirstName:    "TestName",
		LastName:     "TestName",
		PassportData: "1 2 3 4 5 6 7 8 9",
		Phone:        "+543232",
	}
	validator := NewJSONRequestValidator()

	err := validator.ValidateCreateAccount(request)

	assert.Equal(t, "passport_data should be at least 10 characters long", err.Error())
}

func TestJSONRequestValidator_AddressCanBeEmpty(t *testing.T) {
	request := api.CreateAccountRequest{
		FirstName:    "TestFirstName",
		LastName:     "TestLastName",
		PassportData: "1 2 3 4 5 6 7 8 9 0",
		Phone:        "+543232",
	}
	validator := NewJSONRequestValidator()

	err := validator.ValidateCreateAccount(request)

	assert.Nil(t, err)
}
