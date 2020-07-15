package usecase

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/yaroslavnayug/go-payment-system/internal/api"
	"github.com/yaroslavnayug/go-payment-system/internal/domain/commands"
	"github.com/yaroslavnayug/go-payment-system/internal/domain/handlers"
	"github.com/yaroslavnayug/go-payment-system/internal/domain/model"
)

type PaymentSystemAPI struct {
	logger            *logrus.Logger
	accountRepository model.AccountRepositoryInterface
}

func NewPaymentSystemAPI(logger *logrus.Logger, accountRepository model.AccountRepositoryInterface) *PaymentSystemAPI {
	return &PaymentSystemAPI{logger: logger, accountRepository: accountRepository}
}

func (s *PaymentSystemAPI) CreateAccountRequest(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	request := api.CreateAccountRequest{}
	err = json.Unmarshal(body, &request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	command, validationError := commands.NewCreateAccountCommand(request)
	if validationError != nil {
		response, err := api.NewJSONResponse(validationError.Error(), nil)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write(response)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		return
	}

	handler := handlers.NewCreateAccountHandler(s.accountRepository)
	accountID, err := handler.Handle(command)
	if err != nil {
		if _, ok := err.(*model.ValidationError); !ok {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// handle validation error
		validationErrorResponse, err := api.NewJSONResponse(err.Error(), nil)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write(validationErrorResponse)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		return
	}

	customResponse := make(map[string]uint64)
	customResponse["account_id"] = accountID
	response, err := api.NewJSONResponse(http.StatusText(http.StatusOK), customResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
