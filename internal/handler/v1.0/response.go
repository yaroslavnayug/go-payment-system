package v1

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

type ErrorResponse struct {
	Error Error `json:"error"`
}

type Error struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func WriteSuccessPOST(ctx *fasthttp.RequestCtx, responseBody interface{}) {
	response, err := json.Marshal(responseBody)
	if err != nil {
		ctx.Response.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	ctx.Response.SetStatusCode(fasthttp.StatusCreated)
	ctx.Response.SetBody(response)
}

func WriteSuccessGET(ctx *fasthttp.RequestCtx, responseBody interface{}) {
	response, err := json.Marshal(responseBody)
	if err != nil {
		ctx.Response.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	ctx.Response.SetStatusCode(fasthttp.StatusOK)
	ctx.Response.SetBody(response)
}

func WriteSuccessPUT(ctx *fasthttp.RequestCtx) {
	ctx.Response.SetStatusCode(fasthttp.StatusOK)
}

func WriteSuccessDELETE(ctx *fasthttp.RequestCtx) {
	ctx.Response.SetStatusCode(fasthttp.StatusNoContent)
}

func WriteError(ctx *fasthttp.RequestCtx, message string, code int) {
	customError := Error{
		Status:  code,
		Message: message,
	}
	response, err := json.Marshal(&ErrorResponse{Error: customError})
	if err != nil {
		ctx.Response.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	ctx.Response.SetStatusCode(code)
	ctx.Response.SetBody(response)
}
