package v1

import (
	"encoding/json"
	"net/http"

	"github.com/valyala/fasthttp"
	"github.com/yaroslavnayug/go-payment-system/internal/domain"
	"github.com/yaroslavnayug/go-payment-system/internal/usecase"
	"go.uber.org/zap"
)

const CustomerIdUrlPath = "id"

type CustomerHandlerV1 struct {
	logger          *zap.Logger
	customerUseCase *usecase.CustomerUseCase
}

func NewCustomerHandlerV1(logger *zap.Logger, customerService *usecase.CustomerUseCase) *CustomerHandlerV1 {
	return &CustomerHandlerV1{logger: logger, customerUseCase: customerService}
}

type CustomerRequest struct {
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

type CustomerResponse struct {
	CustomerID string `json:"customer_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Address    struct {
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

func (s *CustomerHandlerV1) Create(ctx *fasthttp.RequestCtx) {
	request := &CustomerRequest{}
	err := json.Unmarshal(ctx.PostBody(), request)
	if err != nil {
		WriteError(ctx, http.StatusText(fasthttp.StatusBadRequest), fasthttp.StatusBadRequest)
		return
	}

	customer, err := customerFromRequest(request)
	if err != nil {
		WriteError(ctx, err.Error(), fasthttp.StatusBadRequest)
		return
	}

	err = s.customerUseCase.Create(customer)
	if err != nil {
		if _, isValidationError := err.(*domain.ValidationError); isValidationError {
			WriteError(ctx, err.Error(), fasthttp.StatusConflict)
			return
		} else {
			s.logger.Error(err)
			WriteError(ctx, http.StatusText(fasthttp.StatusInternalServerError), fasthttp.StatusInternalServerError)
			return
		}
	}
	WriteSuccessPOST(ctx, responseFromCustomer(customer))
}

func (s *CustomerHandlerV1) Find(ctx *fasthttp.RequestCtx) {
	customerID := ctx.UserValue(CustomerIdUrlPath)
	if _, ok := customerID.(string); !ok {
		WriteError(ctx, fasthttp.StatusMessage(fasthttp.StatusNotFound), fasthttp.StatusNotFound)
		return
	}
	customer, err := s.customerUseCase.Find(customerID.(string))
	if err != nil {
		s.logger.Error(err)
		WriteError(ctx, fasthttp.StatusMessage(fasthttp.StatusInternalServerError), fasthttp.StatusInternalServerError)
		return
	}
	if customer == nil {
		WriteError(ctx, fasthttp.StatusMessage(fasthttp.StatusNotFound), fasthttp.StatusNotFound)
		return
	}
	WriteSuccessGET(ctx, responseFromCustomer(customer))
}

func (s *CustomerHandlerV1) Update(ctx *fasthttp.RequestCtx) {
	customerID := ctx.UserValue(CustomerIdUrlPath)
	if _, ok := customerID.(string); !ok {
		WriteError(ctx, fasthttp.StatusMessage(fasthttp.StatusNotFound), fasthttp.StatusNotFound)
		return
	}

	request := &CustomerRequest{}
	err := json.Unmarshal(ctx.PostBody(), request)
	if err != nil {
		WriteError(ctx, http.StatusText(fasthttp.StatusBadRequest), fasthttp.StatusBadRequest)
		return
	}

	customer, err := customerFromRequest(request)
	if err != nil {
		WriteError(ctx, err.Error(), fasthttp.StatusBadRequest)
		return
	}

	err = s.customerUseCase.Update(customer, customerID.(string))
	if err != nil {
		if _, isValidationError := err.(*domain.ValidationError); isValidationError {
			WriteError(ctx, err.Error(), fasthttp.StatusConflict)
			return
		} else {
			s.logger.Error(err)
			WriteError(ctx, http.StatusText(fasthttp.StatusInternalServerError), fasthttp.StatusInternalServerError)
			return
		}
	}
	WriteSuccessPUT(ctx)
}

func (s *CustomerHandlerV1) Delete(ctx *fasthttp.RequestCtx) {
	customerID := ctx.UserValue(CustomerIdUrlPath)
	if _, ok := customerID.(string); !ok {
		WriteError(ctx, fasthttp.StatusMessage(fasthttp.StatusNotFound), fasthttp.StatusNotFound)
		return
	}
	err := s.customerUseCase.Delete(customerID.(string))
	if err != nil {
		s.logger.Error(err)
		WriteError(ctx, fasthttp.StatusMessage(fasthttp.StatusInternalServerError), fasthttp.StatusInternalServerError)
		return
	}
	WriteSuccessDELETE(ctx)
}
