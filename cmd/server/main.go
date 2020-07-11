package main

import (
	"net/http"

	"github.com/yaroslavnayug/go-payment-system/internal/app/api"
)

func main() {
	server := api.Server{}
	http.HandleFunc("/createAccount", server.CreateAccountRequest)

	_ = http.ListenAndServe(":8090", nil)
}
