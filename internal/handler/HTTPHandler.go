package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/yaroslavnayug/go-payment-system/internal/api"
	"github.com/yaroslavnayug/go-payment-system/internal/domain/model"
	"github.com/yaroslavnayug/go-payment-system/internal/domain/service"
)

type HTTPHandler struct {
	logger         *logrus.Logger
	accountService *service.AccountService
}

func NewHTTPHandler(logger *logrus.Logger, accountService *service.AccountService) *HTTPHandler {
	return &HTTPHandler{logger: logger, accountService: accountService}
}

func (s *HTTPHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		s.logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	request := api.CreateAccountRequest{}
	err = json.Unmarshal(body, &request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	accountID, err := s.accountService.CreateAccount(request)
	if err != nil {
		if _, ok := err.(*model.ValidationError); !ok {
			s.logger.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// handle validation error
		validationErrorResponse, err := api.NewJSONResponse(err.Error(), nil)
		if err != nil {
			s.logger.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write(validationErrorResponse)
		if err != nil {
			s.logger.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		return
	}

	customResponse := make(map[string]uint64)
	customResponse["account_id"] = accountID
	response, err := api.NewJSONResponse(http.StatusText(http.StatusOK), customResponse)
	if err != nil {
		s.logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(response)
	if err != nil {
		s.logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
