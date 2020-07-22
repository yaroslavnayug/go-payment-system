package converter

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yaroslavnayug/go-payment-system/internal/api"
)

func TestJSONRequestConverter_WrongRequestType(t *testing.T) {
	converter := NewJSONRequestConverter()

	_, err := converter.ConvertToAccount("foo")

	assert.EqualError(t, err, "invalid request type sent to JSONRequestConverter: expected CreateAccountRequest, got foo")
}

func TestJSONRequestConverter_AddressCanBeEmpty(t *testing.T) {
	request := api.CreateAccountRequest{
		FirstName:    "TestFirstName",
		LastName:     "TestLastName",
		PassportData: "1 2 3 4 5 6 7 8 9 0",
		Phone:        "+543232",
	}
	converter := NewJSONRequestConverter()

	account, err := converter.ConvertToAccount(request)

	assert.Nil(t, err)
	assert.Equal(t, "TestFirstName", account.FirstName)
	assert.Equal(t, "TestLastName", account.LastName)
	assert.Equal(t, "1234567890", account.PassportData)
	assert.Equal(t, "+543232", account.Phone)
	assert.Equal(t, "", account.Country)
	assert.Equal(t, "", account.Region)
	assert.Equal(t, "", account.City)
	assert.Equal(t, "", account.Street)
}

func TestJSONRequestConverter_AllDataFilled(t *testing.T) {
	request := api.CreateAccountRequest{
		FirstName:    "TestFirstName",
		LastName:     "TestLastName",
		PassportData: "1 2 3 4 5 6 7 8 9 0",
		Phone:        "+543232",
		Address:      api.Address{Country: "Russia", Region: "NSO", City: "NSK", Street: "Red Square"},
	}
	converter := NewJSONRequestConverter()

	account, err := converter.ConvertToAccount(request)

	assert.Nil(t, err)
	assert.Equal(t, "TestFirstName", account.FirstName)
	assert.Equal(t, "TestLastName", account.LastName)
	assert.Equal(t, "1234567890", account.PassportData)
	assert.Equal(t, "+543232", account.Phone)
	assert.Equal(t, "Russia", account.Country)
	assert.Equal(t, "NSO", account.Region)
	assert.Equal(t, "NSK", account.City)
	assert.Equal(t, "Red Square", account.Street)
}
