package handlers

import (
	"github.com/yaroslavnayug/go-payment-system/internal/domain/commands"
	"github.com/yaroslavnayug/go-payment-system/internal/domain/model"
)

type CreateAccountHandler struct {
	repository model.AccountRepositoryInterface
}

func NewCreateAccountHandler(repository model.AccountRepositoryInterface) *CreateAccountHandler {
	return &CreateAccountHandler{repository: repository}
}
func (h *CreateAccountHandler) Handle(command *commands.CreateAccountCommand) (accountID uint64, err error) {
	accountID, err = h.repository.CreateAccount(command.Account)
	if err != nil {
		return 0, err
	}
	return accountID, nil
}
