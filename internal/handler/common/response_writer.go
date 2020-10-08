package handler

import "github.com/valyala/fasthttp"

type ResponseWriterInterface interface {
	WriteSuccessPOST(ctx *fasthttp.RequestCtx, responseBody interface{})
	WriteSuccessGET(ctx *fasthttp.RequestCtx, responseBody interface{})
	WriteSuccessPUT(ctx *fasthttp.RequestCtx)
	WriteSuccessDELETE(ctx *fasthttp.RequestCtx)
	WriteError(ctx *fasthttp.RequestCtx, message string, code int)
}
