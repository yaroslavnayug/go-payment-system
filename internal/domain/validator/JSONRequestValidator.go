package validator

import (
	"fmt"
	"strings"

	"github.com/yaroslavnayug/go-payment-system/internal/api"
	"github.com/yaroslavnayug/go-payment-system/internal/domain/model"
)

type JSONRequestValidator struct{}

func NewJSONRequestValidator() *JSONRequestValidator {
	return &JSONRequestValidator{}
}

func (v *JSONRequestValidator) ValidateCreateAccount(requestData interface{}) error {
	request, json := requestData.(api.CreateAccountRequest)
	if !json {
		return fmt.Errorf("invalid request type sent to JSONRequestValidator: expected CreateAccountRequest, got %+v", requestData)
	}

	if request.FirstName == "" {
		return model.NewValidationError("first_name is mandatory field")
	}
	if request.LastName == "" {
		return model.NewValidationError("last_name is mandatory field")
	}
	if request.PassportData == "" {
		return model.NewValidationError("passport_data is mandatory field")
	}
	if request.Phone == "" {
		return model.NewValidationError("phone is mandatory field")
	}
	passportData := strings.Replace(request.PassportData, " ", "", -1)
	if len(passportData) < 10 {
		return model.NewValidationError("passport_data should be at least 10 characters long")
	}
	return nil
}
