package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/valyala/fasthttp"
	"github.com/yaroslavnayug/go-payment-system/internal/domain"
	handler "github.com/yaroslavnayug/go-payment-system/internal/handler/common"
	"github.com/yaroslavnayug/go-payment-system/internal/usecase"
	"go.uber.org/zap"
)

const CustomerIdUrlPath = "id"

type CustomerHandlerV1 struct {
	logger         *zap.Logger
	useCase        *usecase.CustomerUseCase
	responseWriter handler.ResponseWriterInterface
}

func NewCustomerHandlerV1(
	logger *zap.Logger,
	customerService *usecase.CustomerUseCase,
	responseWriter handler.ResponseWriterInterface,
) *CustomerHandlerV1 {
	return &CustomerHandlerV1{logger: logger, useCase: customerService, responseWriter: responseWriter}
}

// swagger:parameters CreateCustomer UpdateCustomer
type CustomerBody struct {
	// in:body
	CustomerID string `json:"customer_id"`
	// in:body
	FirstName string `json:"first_name"`
	// in:body
	LastName string `json:"last_name"`
	// in:body
	Email string `json:"email"`
	// in:body
	Phone string `json:"phone"`
	// in:body
	Address struct {
		Country  string `json:"country"`
		Region   string `json:"region"`
		City     string `json:"city"`
		Street   string `json:"street"`
		Building string `json:"building"`
	} `json:"address"`
	// in:body
	Passport struct {
		Number     string `json:"number"`
		IssueDate  string `json:"issue_date"`
		Issuer     string `json:"issuer"`
		BirthDate  string `json:"birth_date"`
		BirthPlace string `json:"birth_place"`
	} `json:"passport"`
}

// swagger:route POST /customer customers CreateCustomer
// Creates a new customer.
// responses:
//  200:
//  400: ErrorResponse
//  409: ErrorResponse
//  500: ErrorResponse
func (h *CustomerHandlerV1) Create(ctx *fasthttp.RequestCtx) {
	request := &CustomerBody{}
	err := json.Unmarshal(ctx.PostBody(), request)
	if err != nil {
		h.responseWriter.WriteError(ctx, http.StatusText(fasthttp.StatusBadRequest), fasthttp.StatusBadRequest)
		return
	}

	customer, err := customerFromRequest(request)
	if err != nil {
		h.responseWriter.WriteError(ctx, err.Error(), fasthttp.StatusBadRequest)
		return
	}

	err = h.useCase.Create(customer)
	if err != nil {
		if err, isValidationError := err.(*domain.ValidationError); isValidationError {
			h.responseWriter.WriteError(ctx, err.Error(), fasthttp.StatusConflict)
			return
		} else {
			h.logger.Error(fmt.Sprintf("error while create customer. request: %s, error: %s", ctx.PostBody(), err.Error()))
			h.responseWriter.WriteError(
				ctx,
				http.StatusText(fasthttp.StatusInternalServerError),
				fasthttp.StatusInternalServerError,
			)
			return
		}
	}
	h.responseWriter.WriteSuccessPOST(ctx, responseFromCustomer(customer))
}

// swagger:route GET /customer/{id} customers FindCustomer
// Finds existing customer by ID.
// responses:
//  200:
//  400: ErrorResponse
//  404: ErrorResponse
//  500: ErrorResponse
func (h *CustomerHandlerV1) Find(ctx *fasthttp.RequestCtx) {
	customerID := ctx.UserValue(CustomerIdUrlPath)
	if _, ok := customerID.(string); !ok {
		h.responseWriter.WriteError(ctx, fasthttp.StatusMessage(fasthttp.StatusBadRequest), fasthttp.StatusBadRequest)
		return
	}
	customer, err := h.useCase.Find(customerID.(string))
	if err != nil {
		h.logger.Error(fmt.Sprintf("error while find customer. customerID: %s, error: %s", customerID, err.Error()))
		h.responseWriter.WriteError(
			ctx,
			fasthttp.StatusMessage(fasthttp.StatusInternalServerError),
			fasthttp.StatusInternalServerError,
		)
		return
	}
	if customer == nil {
		h.responseWriter.WriteError(ctx, fasthttp.StatusMessage(fasthttp.StatusNotFound), fasthttp.StatusNotFound)
		return
	}
	h.responseWriter.WriteSuccessGET(ctx, responseFromCustomer(customer))
}

// swagger:route PUT /customer/{id} customers UpdateCustomer
// Updates existing customer.
// responses:
//  200:
//  400: ErrorResponse
//  404: ErrorResponse
//  500: ErrorResponse
func (h *CustomerHandlerV1) Update(ctx *fasthttp.RequestCtx) {
	customerID := ctx.UserValue(CustomerIdUrlPath)
	if _, ok := customerID.(string); !ok {
		h.responseWriter.WriteError(ctx, fasthttp.StatusMessage(fasthttp.StatusBadRequest), fasthttp.StatusBadRequest)
		return
	}

	request := &CustomerBody{}
	err := json.Unmarshal(ctx.PostBody(), request)
	if err != nil {
		h.responseWriter.WriteError(ctx, http.StatusText(fasthttp.StatusBadRequest), fasthttp.StatusBadRequest)
		return
	}

	customer, err := customerFromRequest(request)
	if err != nil {
		h.responseWriter.WriteError(ctx, err.Error(), fasthttp.StatusBadRequest)
		return
	}

	err = h.useCase.Update(customer, customerID.(string))
	if err != nil {
		if err, isValidationError := err.(*domain.ValidationError); isValidationError {
			h.responseWriter.WriteError(ctx, err.Error(), fasthttp.StatusNotFound)
			return
		} else {
			h.logger.Error(
				fmt.Sprintf("error while update customer. request: %s, error: %s", ctx.PostBody(), err.Error()),
			)
			h.responseWriter.WriteError(
				ctx,
				http.StatusText(fasthttp.StatusInternalServerError),
				fasthttp.StatusInternalServerError,
			)
			return
		}
	}
	h.responseWriter.WriteSuccessPUT(ctx)
}

// swagger:route DELETE /customer/{id} customers DeleteCustomer
// Deletes existing customer.
// responses:
//  200:
//  204:
//  400: ErrorResponse
//  500: ErrorResponse
func (h *CustomerHandlerV1) Delete(ctx *fasthttp.RequestCtx) {
	customerID := ctx.UserValue(CustomerIdUrlPath)
	if _, ok := customerID.(string); !ok {
		h.responseWriter.WriteError(ctx, fasthttp.StatusMessage(fasthttp.StatusBadRequest), fasthttp.StatusBadRequest)
		return
	}
	err := h.useCase.Delete(customerID.(string))
	if err != nil {
		h.logger.Error(fmt.Sprintf("error while delete customer. customerID: %s, error: %s", customerID, err.Error()))
		h.responseWriter.WriteError(
			ctx,
			fasthttp.StatusMessage(fasthttp.StatusInternalServerError),
			fasthttp.StatusInternalServerError,
		)
		return
	}
	h.responseWriter.WriteSuccessDELETE(ctx)
}
