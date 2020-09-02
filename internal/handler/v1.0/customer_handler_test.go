package v1

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/yaroslavnayug/go-payment-system/internal/domain"
	"github.com/yaroslavnayug/go-payment-system/internal/postgres/mocks"
	"github.com/yaroslavnayug/go-payment-system/internal/usecase"
)

func TestCustomerFromRequest_ValidationError(t *testing.T) {
	testCases := []struct {
		name   string
		input  *CreateCustomerRequest
		result string
	}{
		{
			"EmptyFirstName",
			&CreateCustomerRequest{},
			"first_name is mandatory field",
		},

		{
			"EmptyLastName",
			&CreateCustomerRequest{FirstName: "Bruce"},
			"last_name is mandatory field",
		},

		{
			"EmptyPhone",
			&CreateCustomerRequest{
				FirstName: "Bruce",
				LastName:  "Wayne",
			},
			"phone is mandatory field",
		},

		{
			"EmptyAddress",
			&CreateCustomerRequest{
				FirstName: "Bruce",
				LastName:  "Wayne",
				Phone:     "+1234566",
			},
			"address.country is mandatory field",
		},

		{
			"EmptyRegion",
			&CreateCustomerRequest{
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
			&CreateCustomerRequest{
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
			&CreateCustomerRequest{
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
			&CreateCustomerRequest{
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
			&CreateCustomerRequest{
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
			&CreateCustomerRequest{
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
			"EmptyPassportBirthPlace",
			&CreateCustomerRequest{
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
			&CreateCustomerRequest{
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
			&CreateCustomerRequest{
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
				}{Number: "1234567890", BirthDate: "01-01-2000", BirthPlace: "Nov", IssueDate: "01-01-2010"},
			},
			"passport.issuer is mandatory field",
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
	request := &CreateCustomerRequest{
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
	assert.Equal(t, "01-01-2000", customer.Passport.BirthDate)
	assert.Equal(t, "Nov", customer.Passport.BirthPlace)
	assert.Equal(t, "01-01-2010", customer.Passport.IssueDate)
	assert.Equal(t, "Gov", customer.Passport.Issuer)
}

func TestCreate_ValidationError(t *testing.T) {
	testCases := []struct {
		name             string
		input            []byte
		repositoryReturn interface{}
		expectedStatus   int
		expectedResult   string
	}{
		{
			"InvalidInput",
			[]byte(`{"first_name": "foo","last_name": "too"}`),
			nil,
			http.StatusBadRequest,
			`{"error":{"status":400,"message":"phone is mandatory field"}}`,
		},
		{
			"CustomerAlreadyExist",
			[]byte(`{
	"first_name": "foo",
	"last_name": "too",
	"phone": "+7993",
	"address": {
		"country": "R",
		"region": "R",
		"city": "R",
		"street": "R",
		"building": "105"
	},
	"passport": {
		"number": "1234567890",
		"birth_date": "01-01-2000",
		"birth_place": "R",
		"issuer": "MMM",
		"issue_date": "01-01-2000"
	}
}`),
			domain.NewValidationError("customer with such generated id or passport number already exist"),
			http.StatusBadRequest,
			`{"error":{"status":400,"message":"customer with such generated id or passport number already exist"}}`,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			// arrange
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repositoryMock := mocks.NewMockCustomerRepository(ctrl)
			repositoryMock.EXPECT().Create(gomock.Any()).AnyTimes().Return(test.repositoryReturn)

			customerUseCase := usecase.NewCustomerUseCase(
				repositoryMock,
			)
			handlerV1 := NewCustomerHandlerV1(logrus.New(), customerUseCase)
			handlerFunc := http.HandlerFunc(handlerV1.Create)

			// act
			request := httptest.NewRequest("POST", "/customer", bytes.NewBuffer(test.input))
			responseRecorder := httptest.NewRecorder()
			handlerFunc.ServeHTTP(responseRecorder, request)

			// assert
			assert.Equal(t, test.expectedStatus, responseRecorder.Code)
			assert.Equal(t, test.expectedResult, responseRecorder.Body.String())
		})
	}
}

func TestCreate_Success(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mocks.NewMockCustomerRepository(ctrl)
	repositoryMock.EXPECT().Create(gomock.Any()).Return(nil)

	useCase := usecase.NewCustomerUseCase(repositoryMock)

	server := NewCustomerHandlerV1(logrus.New(), useCase)
	handler := http.HandlerFunc(server.Create)

	// act
	var requestBody = []byte(`{
	"first_name": "foo",
	"last_name": "too",
	"phone": "+7993",
	"address": {
		"country": "R",
		"region": "R",
		"city": "R",
		"street": "R",
		"building": "105"
	},
	"passport": {
		"number": "1234567890",
		"birth_date": "01-01-2000",
		"birth_place": "R",
		"issuer": "MMM",
		"issue_date": "01-01-2000"
	}
}`)

	request := httptest.NewRequest("POST", "/customer", bytes.NewBuffer(requestBody))
	responseRecorder := httptest.NewRecorder()
	handler.ServeHTTP(responseRecorder, request)

	// assert
	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	//assert.Equal(t, `{"customer_id":"7530132b81209a9361a3e0786f1fe4d5"}`, responseRecorder.Body.String())
}
