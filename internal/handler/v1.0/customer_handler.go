package v1

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/yaroslavnayug/go-payment-system/internal/domain"
	"github.com/yaroslavnayug/go-payment-system/internal/usecase"
)

type CustomerHandlerV1 struct {
	logger          *logrus.Logger
	customerUseCase *usecase.CustomerUseCase
}

func NewCustomerHandlerV1(logger *logrus.Logger, customerService *usecase.CustomerUseCase) *CustomerHandlerV1 {
	return &CustomerHandlerV1{logger: logger, customerUseCase: customerService}
}

type CreateCustomerRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Address   struct {
		Country  string `json:"country"`
		Region   string `json:"region"`
		City     string `json:"city"`
		Street   string `json:"street"`
		Building string `json:"building"`
	} `json:"address"`
	Passport struct {
		Number     string `json:"number"`
		IssueDate  string `json:"issue_date"`
		Issuer     string `json:"issuer"`
		BirthDate  string `json:"birth_date"`
		BirthPlace string `json:"birth_place"`
	} `json:"passport"`
}

type CreateCustomerResponse struct {
	CustomerID string `json:"customer_id"`
}

func customerFromRequest(request *CreateCustomerRequest) (*domain.Customer, error) {
	if request.FirstName == "" {
		return nil, domain.NewValidationError("first_name is mandatory field")
	}
	if request.LastName == "" {
		return nil, domain.NewValidationError("last_name is mandatory field")
	}
	if request.Phone == "" {
		return nil, domain.NewValidationError("phone is mandatory field")
	}
	if request.Address.Country == "" {
		return nil, domain.NewValidationError("address.country is mandatory field")
	}
	if request.Address.Region == "" {
		return nil, domain.NewValidationError("address.region is mandatory field")
	}
	if request.Address.City == "" {
		return nil, domain.NewValidationError("address.city is mandatory field")
	}
	if request.Address.Street == "" {
		return nil, domain.NewValidationError("address.street is mandatory field")
	}
	if request.Address.Building == "" {
		return nil, domain.NewValidationError("address.building is mandatory field")
	}
	passportNumber := strings.Replace(request.Passport.Number, " ", "", -1)
	if len(passportNumber) != 10 {
		return nil, domain.NewValidationError("passport.number should be 10 characters long")
	}
	if request.Passport.BirthDate == "" {
		return nil, domain.NewValidationError("passport.birth_date is mandatory field")
	}
	if request.Passport.BirthPlace == "" {
		return nil, domain.NewValidationError("passport.birth_place is mandatory field")
	}
	if request.Passport.Issuer == "" {
		return nil, domain.NewValidationError("passport.issuer is mandatory field")
	}
	if request.Passport.IssueDate == "" {
		return nil, domain.NewValidationError("passport.issue_date is mandatory field")
	}

	customer := &domain.Customer{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		Phone:     request.Phone,
		Address: domain.Address{
			Country:  request.Address.Country,
			Region:   request.Address.Region,
			City:     request.Address.City,
			Street:   request.Address.Street,
			Building: request.Address.Building,
		},
		Passport: domain.Passport{
			Number:     request.Passport.Number,
			IssueDate:  request.Passport.IssueDate,
			Issuer:     request.Passport.Issuer,
			BirthDate:  request.Passport.BirthDate,
			BirthPlace: request.Passport.BirthPlace,
		},
	}
	return customer, nil
}

func (s *CustomerHandlerV1) Create(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		s.logger.Error(err)
		WriteError(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	request := &CreateCustomerRequest{}
	err = json.Unmarshal(body, request)
	if err != nil {
		WriteError(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	customer, err := customerFromRequest(request)
	if err != nil {
		WriteError(w, err.Error(), http.StatusBadRequest)
		return
	}

	customerID, err := s.customerUseCase.Create(customer)
	if err != nil {
		if _, isValidationError := err.(*domain.ValidationError); isValidationError {
			WriteError(w, err.Error(), http.StatusBadRequest)
			return
		} else {
			s.logger.Error(err)
			WriteError(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}

	WriteSuccess(w, CreateCustomerResponse{CustomerID: customerID})
}
