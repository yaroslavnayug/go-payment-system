package v1

import (
	"strings"
	"time"

	"github.com/yaroslavnayug/go-payment-system/internal/domain"
)

func customerFromRequest(request *CustomerBody) (*domain.Customer, error) {
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
	birthDate, err := time.Parse(domain.DateFormat, request.Passport.BirthDate)
	if err != nil {
		return nil, domain.NewValidationError("wrong format for passport.birth_date. DD-MM-YYYY expected")
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
	issueDate, err := time.Parse(domain.DateFormat, request.Passport.IssueDate)
	if err != nil {
		return nil, domain.NewValidationError("wrong format for passport.issue_date. DD-MM-YYYY expected")
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
			IssueDate:  issueDate,
			Issuer:     request.Passport.Issuer,
			BirthDate:  birthDate,
			BirthPlace: request.Passport.BirthPlace,
		},
	}
	return customer, nil
}

func responseFromCustomer(customer *domain.Customer) *CustomerBody {
	issueDate := customer.Passport.IssueDate.Format(domain.DateFormat)
	birthDate := customer.Passport.BirthDate.Format(domain.DateFormat)
	return &CustomerBody{
		CustomerID: customer.GeneratedID,
		FirstName:  customer.FirstName,
		LastName:   customer.LastName,
		Email:      customer.Email,
		Phone:      customer.Phone,
		Address: struct {
			Country  string `json:"country"`
			Region   string `json:"region"`
			City     string `json:"city"`
			Street   string `json:"street"`
			Building string `json:"building"`
		}{
			Country:  customer.Address.Country,
			Region:   customer.Address.Region,
			City:     customer.Address.City,
			Street:   customer.Address.Street,
			Building: customer.Address.Building,
		},
		Passport: struct {
			Number     string `json:"number"`
			IssueDate  string `json:"issue_date"`
			Issuer     string `json:"issuer"`
			BirthDate  string `json:"birth_date"`
			BirthPlace string `json:"birth_place"`
		}{
			Number:     customer.Passport.Number,
			IssueDate:  issueDate,
			Issuer:     customer.Passport.Issuer,
			BirthDate:  birthDate,
			BirthPlace: customer.Passport.BirthPlace,
		},
	}
}
