package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/yaroslavnayug/go-payment-system/internal/app/api/requests"
	"github.com/yaroslavnayug/go-payment-system/internal/app/domain/commands"
	"github.com/yaroslavnayug/go-payment-system/internal/app/domain/handlers"
	"github.com/yaroslavnayug/go-payment-system/internal/app/domain/model"
)

type ServerAPI struct {
	accountRepository model.AccountRepositoryInterface
}

func (s *ServerAPI) CreateAccountRequest(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	request := requests.CreateAccountRequest{}
	err = json.Unmarshal(body, &request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	command, err := commands.NewCreateAccountCommand(request)
	if err != nil {
		response, err := requests.NewJSONResponse(http.StatusBadRequest, err.Error(), nil)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(response)
		return
	}

	handler := handlers.NewCreateAccountHandler(s.accountRepository)
	err = handler.Handle(command)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	response, err := requests.NewJSONResponse(http.StatusOK, http.StatusText(http.StatusOK), nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(response)
	return
}
