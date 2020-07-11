package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/yaroslavnayug/go-payment-system/internal/app/api/requests"
)

type Server struct{}

func (s *Server) CreateAccountRequest(w http.ResponseWriter, r *http.Request) {
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

	return
}
