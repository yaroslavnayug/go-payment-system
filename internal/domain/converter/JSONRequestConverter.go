package converter

import (
	"fmt"
	"strings"

	"github.com/yaroslavnayug/go-payment-system/internal/api"
	"github.com/yaroslavnayug/go-payment-system/internal/domain/model"
)

type JSONRequestConverter struct{}

func NewJSONRequestConverter() *JSONRequestConverter {
	return &JSONRequestConverter{}
}

func (b *JSONRequestConverter) ConvertToAccount(requestData interface{}) (*model.Account, error) {
	request, jsonRequest := requestData.(api.CreateAccountRequest)
	if !jsonRequest {
		return nil, fmt.Errorf("invalid request type sent to JSONRequestConverter: expected CreateAccountRequest, got %+v", requestData)
	}

	passportData := strings.Replace(request.PassportData, " ", "", -1)
	account := &model.Account{
		FirstName:    request.FirstName,
		LastName:     request.LastName,
		PassportData: passportData,
		Phone:        request.Phone,
		Country:      request.Address.Country,
		Region:       request.Address.Region,
		City:         request.Address.City,
		Street:       request.Address.Street,
	}
	return account, nil
}
