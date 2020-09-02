package v1

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Error Error `json:"error"`
}

type Error struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func WriteSuccess(w http.ResponseWriter, responseBody interface{}) {
	response, err := json.Marshal(responseBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = w.Write(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func WriteError(w http.ResponseWriter, message string, code int) {
	customError := Error{
		Status:  code,
		Message: message,
	}
	response, err := json.Marshal(&ErrorResponse{Error: customError})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(code)
	_, _ = w.Write(response)
}
