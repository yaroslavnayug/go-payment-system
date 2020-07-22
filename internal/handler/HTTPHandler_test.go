package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/yaroslavnayug/go-payment-system/internal/domain/converter"
	"github.com/yaroslavnayug/go-payment-system/internal/domain/model"
	"github.com/yaroslavnayug/go-payment-system/internal/domain/service"
	"github.com/yaroslavnayug/go-payment-system/internal/domain/validator"
	"github.com/yaroslavnayug/go-payment-system/internal/persistence/mocks"
)

func TestCreateAccountRequest_InvalidInputValidation(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	accountService := service.NewAccountService(
		validator.NewJSONRequestValidator(),
		converter.NewJSONRequestConverter(),
		mocks.NewMockRepositoryInterface(ctrl),
	)
	server := NewHTTPHandler(logrus.New(), accountService)
	handler := http.HandlerFunc(server.CreateAccount)

	// act
	var requestBody = []byte(`{
	"first_name": "foo",
	"last_name": "too",
	"passport_data": "1234567890"
}`)

	request, err := http.NewRequest("POST", "/createAccount", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	responseRecorder := httptest.NewRecorder()
	handler.ServeHTTP(responseRecorder, request)

	// assert
	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	assert.Equal(t, `{"status_message":"phone is mandatory field","response":null}`, responseRecorder.Body.String())
}

func TestCreateAccountRequest_AccountAlreadyExist(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mocks.NewMockRepositoryInterface(ctrl)
	repositoryMock.EXPECT().CreateAccount(gomock.Any()).Return(uint64(0), model.NewValidationError("account with such passport_data already exist"))

	accountService := service.NewAccountService(
		validator.NewJSONRequestValidator(),
		converter.NewJSONRequestConverter(),
		repositoryMock,
	)

	server := NewHTTPHandler(logrus.New(), accountService)
	handler := http.HandlerFunc(server.CreateAccount)

	// act
	var requestBody = []byte(`{
	"first_name": "foo",
	"last_name": "too",
	"passport_data": "1234567890",
	"phone": "+7993"
}`)

	request, err := http.NewRequest("POST", "/createAccount", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	responseRecorder := httptest.NewRecorder()
	handler.ServeHTTP(responseRecorder, request)

	// assert
	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	assert.Equal(t, `{"status_message":"account with such passport_data already exist","response":null}`, responseRecorder.Body.String())
}

func TestCreateAccountRequest_Success(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mocks.NewMockRepositoryInterface(ctrl)
	repositoryMock.EXPECT().CreateAccount(gomock.Any()).Return(uint64(123), nil)

	accountService := service.NewAccountService(
		validator.NewJSONRequestValidator(),
		converter.NewJSONRequestConverter(),
		repositoryMock,
	)

	server := NewHTTPHandler(logrus.New(), accountService)
	handler := http.HandlerFunc(server.CreateAccount)

	// act
	var requestBody = []byte(`{
	"first_name": "foo",
	"last_name": "too",
	"passport_data": "1234567890",
	"phone": "+7993"
}`)

	request, err := http.NewRequest("POST", "/createAccount", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	responseRecorder := httptest.NewRecorder()
	handler.ServeHTTP(responseRecorder, request)

	// assert
	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Equal(t, `{"status_message":"OK","response":{"account_id":123}}`, responseRecorder.Body.String())
}
