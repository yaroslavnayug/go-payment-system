package handlers

import (
	"github.com/yaroslavnayug/go-payment-system/internal/app/domain/commands"
	"github.com/yaroslavnayug/go-payment-system/internal/app/domain/model"
)

type CreateAccountHandler struct {
	repository model.AccountRepositoryInterface
}

func NewCreateAccountHandler(repository model.AccountRepositoryInterface) *CreateAccountHandler {
	return &CreateAccountHandler{repository: repository}
}
func (h *CreateAccountHandler) Handle(command commands.CreateAccountCommand) error {
	return nil
}
