package requests

import "encoding/json"

type Response struct {
	StatusCode    int         `json:"status_code"`
	StatusMessage string      `json:"status_message"`
	Response      interface{} `json:"response"`
}

func NewJSONResponse(statusCode int, statusMessage string, customResponse interface{}) ([]byte, error) {
	response := &Response{
		StatusCode:    statusCode,
		StatusMessage: statusMessage,
		Response:      customResponse,
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}
	return jsonResponse, nil
}
