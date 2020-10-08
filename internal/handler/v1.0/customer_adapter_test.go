package v1

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yaroslavnayug/go-payment-system/internal/domain"
)

func TestCustomerFromRequest_ValidationError(t *testing.T) {
	testCases := []struct {
		name   string
		input  *CustomerBody
		result string
	}{
		{
			"EmptyFirstName",
			&CustomerBody{},
			"first_name is mandatory field",
		},

		{
			"EmptyLastName",
			&CustomerBody{FirstName: "Bruce"},
			"last_name is mandatory field",
		},

		{
			"EmptyPhone",
			&CustomerBody{
				FirstName: "Bruce",
				LastName:  "Wayne",
			},
			"phone is mandatory field",
		},

		{
			"EmptyAddress",
			&CustomerBody{
				FirstName: "Bruce",
				LastName:  "Wayne",
				Phone:     "+1234566",
			},
			"address.country is mandatory field",
		},

		{
			"EmptyRegion",
			&CustomerBody{
				FirstName: "Bruce",
				LastName:  "Wayne",
				Phone:     "+1234566",
				Address: struct {
					Country  string `json:"country"`
					Region   string `json:"region"`
					City     string `json:"city"`
					Street   string `json:"street"`
					Building string `json:"building"`
				}{Country: "Russia"},
			},
			"address.region is mandatory field",
		},

		{
			"EmptyCity",
			&CustomerBody{
				FirstName: "Bruce",
				LastName:  "Wayne",
				Phone:     "+1234566",
				Address: struct {
					Country  string `json:"country"`
					Region   string `json:"region"`
					City     string `json:"city"`
					Street   string `json:"street"`
					Building string `json:"building"`
				}{Country: "Russia", Region: "Sakha"},
			},
			"address.city is mandatory field",
		},

		{
			"EmptyStreet",
			&CustomerBody{
				FirstName: "Bruce",
				LastName:  "Wayne",
				Phone:     "+1234566",
				Address: struct {
					Country  string `json:"country"`
					Region   string `json:"region"`
					City     string `json:"city"`
					Street   string `json:"street"`
					Building string `json:"building"`
				}{Country: "Russia", Region: "Sakha", City: "Yakutsk"},
			},
			"address.street is mandatory field",
		},

		{
			"EmptyBuilding",
			&CustomerBody{
				FirstName: "Bruce",
				LastName:  "Wayne",
				Phone:     "+1234566",
				Address: struct {
					Country  string `json:"country"`
					Region   string `json:"region"`
					City     string `json:"city"`
					Street   string `json:"street"`
					Building string `json:"building"`
				}{Country: "Russia", Region: "Sakha", City: "Yakutsk", Street: "Marks"},
			},
			"address.building is mandatory field",
		},

		{
			"EmptyPassportNumber",
			&CustomerBody{
				FirstName: "Bruce",
				LastName:  "Wayne",
				Phone:     "+1234566",
				Address: struct {
					Country  string `json:"country"`
					Region   string `json:"region"`
					City     string `json:"city"`
					Street   string `json:"street"`
					Building string `json:"building"`
				}{Country: "Russia", Region: "Sakha", City: "Yakutsk", Street: "Marks", Building: "105"},
				Passport: struct {
					Number     string `json:"number"`
					IssueDate  string `json:"issue_date"`
					Issuer     string `json:"issuer"`
					BirthDate  string `json:"birth_date"`
					BirthPlace string `json:"birth_place"`
				}{},
			},
			"passport.number should be 10 characters long",
		},

		{
			"EmptyPassportBirthDate",
			&CustomerBody{
				FirstName: "Bruce",
				LastName:  "Wayne",
				Phone:     "+1234566",
				Address: struct {
					Country  string `json:"country"`
					Region   string `json:"region"`
					City     string `json:"city"`
					Street   string `json:"street"`
					Building string `json:"building"`
				}{Country: "Russia", Region: "Sakha", City: "Yakutsk", Street: "Marks", Building: "105"},
				Passport: struct {
					Number     string `json:"number"`
					IssueDate  string `json:"issue_date"`
					Issuer     string `json:"issuer"`
					BirthDate  string `json:"birth_date"`
					BirthPlace string `json:"birth_place"`
				}{Number: "1234567890"},
			},
			"passport.birth_date is mandatory field",
		},
		{
			"InvalidBirthDate",
			&CustomerBody{
				FirstName: "Bruce",
				LastName:  "Wayne",
				Phone:     "+1234566",
				Address: struct {
					Country  string `json:"country"`
					Region   string `json:"region"`
					City     string `json:"city"`
					Street   string `json:"street"`
					Building string `json:"building"`
				}{Country: "Russia", Region: "Sakha", City: "Yakutsk", Street: "Marks", Building: "105"},
				Passport: struct {
					Number     string `json:"number"`
					IssueDate  string `json:"issue_date"`
					Issuer     string `json:"issuer"`
					BirthDate  string `json:"birth_date"`
					BirthPlace string `json:"birth_place"`
				}{Number: "1234567890", BirthDate: "01-01"},
			},
			"wrong format for passport.birth_date. DD-MM-YYYY expected",
		},

		{
			"EmptyPassportBirthPlace",
			&CustomerBody{
				FirstName: "Bruce",
				LastName:  "Wayne",
				Phone:     "+1234566",
				Address: struct {
					Country  string `json:"country"`
					Region   string `json:"region"`
					City     string `json:"city"`
					Street   string `json:"street"`
					Building string `json:"building"`
				}{Country: "Russia", Region: "Sakha", City: "Yakutsk", Street: "Marks", Building: "105"},
				Passport: struct {
					Number     string `json:"number"`
					IssueDate  string `json:"issue_date"`
					Issuer     string `json:"issuer"`
					BirthDate  string `json:"birth_date"`
					BirthPlace string `json:"birth_place"`
				}{Number: "1234567890", BirthDate: "01-01-2000"},
			},
			"passport.birth_place is mandatory field",
		},

		{
			"EmptyPassportIssuer",
			&CustomerBody{
				FirstName: "Bruce",
				LastName:  "Wayne",
				Phone:     "+1234566",
				Address: struct {
					Country  string `json:"country"`
					Region   string `json:"region"`
					City     string `json:"city"`
					Street   string `json:"street"`
					Building string `json:"building"`
				}{Country: "Russia", Region: "Sakha", City: "Yakutsk", Street: "Marks", Building: "105"},
				Passport: struct {
					Number     string `json:"number"`
					IssueDate  string `json:"issue_date"`
					Issuer     string `json:"issuer"`
					BirthDate  string `json:"birth_date"`
					BirthPlace string `json:"birth_place"`
				}{Number: "1234567890", BirthDate: "01-01-2000", BirthPlace: "Nov"},
			},
			"passport.issuer is mandatory field",
		},

		{
			"EmptyPassportIssueDate",
			&CustomerBody{
				FirstName: "Bruce",
				LastName:  "Wayne",
				Phone:     "+1234566",
				Address: struct {
					Country  string `json:"country"`
					Region   string `json:"region"`
					City     string `json:"city"`
					Street   string `json:"street"`
					Building string `json:"building"`
				}{Country: "Russia", Region: "Sakha", City: "Yakutsk", Street: "Marks", Building: "105"},
				Passport: struct {
					Number     string `json:"number"`
					IssueDate  string `json:"issue_date"`
					Issuer     string `json:"issuer"`
					BirthDate  string `json:"birth_date"`
					BirthPlace string `json:"birth_place"`
				}{Number: "1234567890", BirthDate: "01-01-2010", BirthPlace: "Nov", Issuer: "Test", IssueDate: ""},
			},
			"passport.issue_date is mandatory field",
		},
		{
			"InvalidIssueDate",
			&CustomerBody{
				FirstName: "Bruce",
				LastName:  "Wayne",
				Phone:     "+1234566",
				Address: struct {
					Country  string `json:"country"`
					Region   string `json:"region"`
					City     string `json:"city"`
					Street   string `json:"street"`
					Building string `json:"building"`
				}{Country: "Russia", Region: "Sakha", City: "Yakutsk", Street: "Marks", Building: "105"},
				Passport: struct {
					Number     string `json:"number"`
					IssueDate  string `json:"issue_date"`
					Issuer     string `json:"issuer"`
					BirthDate  string `json:"birth_date"`
					BirthPlace string `json:"birth_place"`
				}{Number: "1234567890", BirthDate: "01-01-2000", BirthPlace: "Nov", Issuer: "Test", IssueDate: "01-01"},
			},
			"wrong format for passport.issue_date. DD-MM-YYYY expected",
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			_, err := customerFromRequest(test.input)
			assert.Equal(t, test.result, err.Error())
		})
	}
}

func TestCustomerFromRequest_AllMandatoryFields(t *testing.T) {
	request := &CustomerBody{
		FirstName: "Bruce",
		LastName:  "Wayne",
		Email:     "goo@gmail.com",
		Phone:     "+1234566",
		Address: struct {
			Country  string `json:"country"`
			Region   string `json:"region"`
			City     string `json:"city"`
			Street   string `json:"street"`
			Building string `json:"building"`
		}{Country: "Russia", Region: "Sakha", City: "Yakutsk", Street: "Marks", Building: "105"},
		Passport: struct {
			Number     string `json:"number"`
			IssueDate  string `json:"issue_date"`
			Issuer     string `json:"issuer"`
			BirthDate  string `json:"birth_date"`
			BirthPlace string `json:"birth_place"`
		}{Number: "1234567890", BirthDate: "01-01-2000", BirthPlace: "Nov", IssueDate: "01-01-2010", Issuer: "Gov"},
	}

	customer, err := customerFromRequest(request)
	assert.Nil(t, err)
	assert.Equal(t, "Bruce", customer.FirstName)
	assert.Equal(t, "Wayne", customer.LastName)
	assert.Equal(t, "goo@gmail.com", customer.Email)
	assert.Equal(t, "+1234566", customer.Phone)
	assert.Equal(t, "Russia", customer.Address.Country)
	assert.Equal(t, "Sakha", customer.Address.Region)
	assert.Equal(t, "Yakutsk", customer.Address.City)
	assert.Equal(t, "Marks", customer.Address.Street)
	assert.Equal(t, "105", customer.Address.Building)
	assert.Equal(t, "1234567890", customer.Passport.Number)
	assert.Equal(t, "01-01-2000", customer.Passport.BirthDate.Format(domain.DateFormat))
	assert.Equal(t, "Nov", customer.Passport.BirthPlace)
	assert.Equal(t, "01-01-2010", customer.Passport.IssueDate.Format(domain.DateFormat))
	assert.Equal(t, "Gov", customer.Passport.Issuer)
}
