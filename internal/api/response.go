package api

import "encoding/json"

type Response struct {
	StatusMessage string      `json:"status_message"`
	Response      interface{} `json:"response"`
}

func NewJSONResponse(statusMessage string, customResponse interface{}) ([]byte, error) {
	response := &Response{
		StatusMessage: statusMessage,
		Response:      customResponse,
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}
	return jsonResponse, nil
}
