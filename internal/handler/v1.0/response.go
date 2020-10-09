package v1

import (
	"encoding/json"
	"fmt"

	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

const ContentTypeJSON = "application/json"

// swagger:response ErrorResponse
type ErrorResponse struct {
	Error Error `json:"error"`
}

type Error struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type JSONResponseWriter struct {
	logger *zap.Logger
}

func NewJSONResponseWriter(logger *zap.Logger) *JSONResponseWriter {
	return &JSONResponseWriter{logger: logger}
}

func (w *JSONResponseWriter) WriteSuccessPOST(ctx *fasthttp.RequestCtx, responseBody interface{}) {
	response, err := json.Marshal(responseBody)
	if err != nil {
		w.logger.Error(fmt.Sprintf("error while marshal response body: %+v, error %s", responseBody, err.Error()))
		ctx.Response.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	ctx.Response.Header.SetContentType(ContentTypeJSON)
	ctx.Response.SetStatusCode(fasthttp.StatusCreated)
	ctx.Response.SetBody(response)
}

func (w *JSONResponseWriter) WriteSuccessGET(ctx *fasthttp.RequestCtx, responseBody interface{}) {
	response, err := json.Marshal(responseBody)
	if err != nil {
		w.logger.Error(fmt.Sprintf("error while marshal response body: %+v, error %s", responseBody, err.Error()))
		ctx.Response.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	ctx.Response.Header.SetContentType(ContentTypeJSON)
	ctx.Response.SetStatusCode(fasthttp.StatusOK)
	ctx.Response.SetBody(response)
}

func (w *JSONResponseWriter) WriteSuccessPUT(ctx *fasthttp.RequestCtx) {
	ctx.Response.SetStatusCode(fasthttp.StatusOK)
}

func (w *JSONResponseWriter) WriteSuccessDELETE(ctx *fasthttp.RequestCtx) {
	ctx.Response.SetStatusCode(fasthttp.StatusNoContent)
}

func (w *JSONResponseWriter) WriteError(ctx *fasthttp.RequestCtx, message string, code int) {
	customError := Error{
		Status:  code,
		Message: message,
	}
	responseBody := &ErrorResponse{Error: customError}

	response, err := json.Marshal(responseBody)
	if err != nil {
		w.logger.Error(fmt.Sprintf("error while marshal responseBody: %+v, error %s", responseBody, err.Error()))
		ctx.Response.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	ctx.Response.Header.SetContentType(ContentTypeJSON)
	ctx.Response.SetStatusCode(code)
	ctx.Response.SetBody(response)
}
