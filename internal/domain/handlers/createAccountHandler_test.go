package handlers

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/yaroslavnayug/go-payment-system/internal/domain/commands"
	"github.com/yaroslavnayug/go-payment-system/internal/domain/model"
	"github.com/yaroslavnayug/go-payment-system/internal/persistence/mocks"
)

func TestCreateAccountHandler_AccountAlreadyExist(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mocks.NewMockAccountRepositoryInterface(ctrl)
	mock.EXPECT().CreateAccount(gomock.Any()).Return(uint64(0), model.NewValidationError("account with such passport_data already exist"))
	handler := NewCreateAccountHandler(mock)

	command := &commands.CreateAccountCommand{
		Account: model.Account{
			Id:           0,
			FirstName:    "AccountAlreadyExist",
			LastName:     "AccountAlreadyExist",
			PassportData: "1234567890",
			Phone:        "+7",
			Country:      "",
			Region:       "",
			City:         "",
			Street:       "",
		},
	}

	// act
	accountID, err := handler.Handle(command)

	// assert
	assert.Equal(t, "account with such passport_data already exist", err.Error())
	assert.Equal(t, uint64(0), accountID)
}

func TestCreateAccountHandler_AccountCreated(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mocks.NewMockAccountRepositoryInterface(ctrl)
	mock.EXPECT().CreateAccount(gomock.Any()).Return(uint64(666), nil)
	handler := NewCreateAccountHandler(mock)

	command := &commands.CreateAccountCommand{
		Account: model.Account{
			Id:           0,
			FirstName:    "AccountAlreadyExist",
			LastName:     "AccountAlreadyExist",
			PassportData: "1234567890",
			Phone:        "+7",
			Country:      "",
			Region:       "",
			City:         "",
			Street:       "",
		},
	}

	// act
	accountID, err := handler.Handle(command)

	// assert
	assert.Nil(t, err)
	assert.Equal(t, uint64(666), accountID)
}
