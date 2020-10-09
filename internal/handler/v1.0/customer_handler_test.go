package v1

import (
	"net"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/buaazp/fasthttprouter"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttputil"
	"github.com/yaroslavnayug/go-payment-system/internal/domain"
	"go.uber.org/zap"

	"github.com/yaroslavnayug/go-payment-system/internal/postgres/mocks"
	"github.com/yaroslavnayug/go-payment-system/internal/usecase"
)

func TestCreate_Success(t *testing.T) {
	t.Parallel()

	// arrange deps
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mocks.NewMockCustomerRepository(ctrl)
	repositoryMock.EXPECT().FindByPassportNumber(gomock.Any()).Return(nil, nil)
	repositoryMock.EXPECT().Create(gomock.Any()).Return(nil)
	useCase := usecase.NewCustomerUseCase(repositoryMock)
	logger, _ := zap.NewDevelopment()
	writer := NewJSONResponseWriter(logger)
	handlerV1 := NewCustomerHandlerV1(logger, useCase, writer)

	// arrange fake server
	router := fasthttprouter.New()
	router.POST("/customer", handlerV1.Create)

	ln := fasthttputil.NewInmemoryListener()

	s := &fasthttp.Server{
		Handler: router.Handler,
	}
	go func() {
		_ = s.Serve(ln)
	}()

	client := fasthttp.Client{
		Dial: func(addr string) (net.Conn, error) {
			return ln.Dial()
		},
	}
	request, response := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(request)
		fasthttp.ReleaseResponse(response)
	}()

	// act
	var requestBody = []byte(`{
		"first_name": "foo",
		"last_name": "too",
		"phone": "+7993",
		"address": {
			"country": "R",
			"region": "R",
			"city": "R",
			"street": "R",
			"building": "105"
		},
		"passport": {
			"number": "1234567890",
			"birth_date": "01-01-2000",
			"birth_place": "R",
			"issuer": "MMM",
			"issue_date": "01-01-2000"
		}
	}`)

	request.Header.SetMethod(fasthttp.MethodPost)
	request.SetBody(requestBody)
	request.SetRequestURI("/customer")
	request.SetHost("localhost")

	_ = client.Do(request, response)

	// assert
	assert.Equal(t, fasthttp.StatusCreated, response.Header.StatusCode())
}

func TestCreate_ValidationError(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name             string
		input            []byte
		repositoryReturn interface{}
		expectedStatus   int
		expectedResult   string
	}{
		{
			"InvalidInput",
			[]byte(`{"first_name": "foo","last_name": "too"}`),
			nil,
			fasthttp.StatusBadRequest,
			`{"error":{"status":400,"message":"phone is mandatory field"}}`,
		},
		{
			"CustomerAlreadyExist",
			[]byte(`{
		"first_name": "foo",
		"last_name": "too",
		"phone": "+7993",
		"address": {
			"country": "R",
			"region": "R",
			"city": "R",
			"street": "R",
			"building": "105"
		},
		"passport": {
			"number": "1234567890",
			"birth_date": "01-01-2000",
			"birth_place": "R",
			"issuer": "MMM",
			"issue_date": "01-01-2000"
		}
	}`),
			domain.NewValidationError("customer with such generated id or passport number already exist"),
			fasthttp.StatusConflict,
			`{"error":{"status":409,"message":"customer with such passport number already exist"}}`,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			// arrange deps
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repositoryMock := mocks.NewMockCustomerRepository(ctrl)
			repositoryMock.EXPECT().FindByPassportNumber(gomock.Any()).AnyTimes().Return(&domain.Customer{}, nil)
			useCase := usecase.NewCustomerUseCase(repositoryMock)
			logger, _ := zap.NewDevelopment()
			writer := NewJSONResponseWriter(logger)
			handlerV1 := NewCustomerHandlerV1(logger, useCase, writer)

			// arrange fake server
			router := fasthttprouter.New()
			router.POST("/customer", handlerV1.Create)

			ln := fasthttputil.NewInmemoryListener()

			s := &fasthttp.Server{
				Handler: router.Handler,
			}
			go func() {
				_ = s.Serve(ln)
			}()

			client := fasthttp.Client{
				Dial: func(addr string) (net.Conn, error) {
					return ln.Dial()
				},
			}
			request, response := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
			defer func() {
				fasthttp.ReleaseRequest(request)
				fasthttp.ReleaseResponse(response)
			}()

			// act
			request.Header.SetMethod(fasthttp.MethodPost)
			request.SetBody(test.input)
			request.SetRequestURI("/customer")
			request.SetHost("localhost")

			_ = client.Do(request, response)

			// assert
			assert.Equal(t, test.expectedStatus, response.Header.StatusCode())
			assert.Equal(t, test.expectedResult, string(response.Body()))
		})
	}
}

func TestFind_CustomerFound(t *testing.T) {
	// arrange deps
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	birthDate, _ := time.Parse(domain.DateFormat, "10-10-2020")
	issueDate, _ := time.Parse(domain.DateFormat, "20-10-2020")
	customer := domain.Customer{
		GeneratedID: "foobar",
		FirstName:   "Bruce",
		LastName:    "Wayne",
		Email:       "batman@gmail.com",
		Phone:       "+789",
		Address:     domain.Address{},
		Passport: domain.Passport{
			BirthDate: birthDate,
			IssueDate: issueDate,
		},
	}

	repositoryMock := mocks.NewMockCustomerRepository(ctrl)
	repositoryMock.EXPECT().FindByID(gomock.Any()).Return(&customer, nil)

	useCase := usecase.NewCustomerUseCase(repositoryMock)
	logger, _ := zap.NewDevelopment()
	writer := NewJSONResponseWriter(logger)
	handlerV1 := NewCustomerHandlerV1(logger, useCase, writer)

	// arrange fake server
	router := fasthttprouter.New()
	router.GET("/customer/:id", handlerV1.Find)

	listener := fasthttputil.NewInmemoryListener()

	server := &fasthttp.Server{
		Handler: router.Handler,
	}
	go func() {
		_ = server.Serve(listener)
	}()

	client := fasthttp.Client{
		Dial: func(addr string) (net.Conn, error) {
			return listener.Dial()
		},
	}
	request, response := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(request)
		fasthttp.ReleaseResponse(response)
	}()

	// act
	request.SetRequestURI("/customer/foobar")
	request.Header.SetMethod(fasthttp.MethodGet)
	request.SetHost("localhost")

	_ = client.Do(request, response)

	// assert
	assert.Equal(t, http.StatusOK, response.Header.StatusCode())

	expectedBody := `{
	"customer_id":"foobar",
	"first_name":"Bruce",
	"last_name":"Wayne",
	"email":"batman@gmail.com",
	"phone":"+789",
	"address":{
		"country":"",
		"region":"",
		"city":"",
		"street":"",
		"building":""
	},
	"passport":{
		"number":"",
		"issue_date":"20-10-2020",
		"issuer":"",
		"birth_date":"10-10-2020",
		"birth_place":""
	}
}`
	assert.Equal(t, strings.Replace(strings.Replace(expectedBody, "\t", "", -1), "\n", "", -1), string(response.Body()))
}

func TestFind_CustomerNotFound(t *testing.T) {
	// arrange deps
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mocks.NewMockCustomerRepository(ctrl)
	repositoryMock.EXPECT().FindByID(gomock.Any()).Return(nil, nil)

	useCase := usecase.NewCustomerUseCase(repositoryMock)
	logger, _ := zap.NewDevelopment()
	writer := NewJSONResponseWriter(logger)
	handlerV1 := NewCustomerHandlerV1(logger, useCase, writer)

	// arrange fake server
	router := fasthttprouter.New()
	router.GET("/customer/:id", handlerV1.Find)

	listener := fasthttputil.NewInmemoryListener()

	server := &fasthttp.Server{
		Handler: router.Handler,
	}
	go func() {
		_ = server.Serve(listener)
	}()

	client := fasthttp.Client{
		Dial: func(addr string) (net.Conn, error) {
			return listener.Dial()
		},
	}
	request, response := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(request)
		fasthttp.ReleaseResponse(response)
	}()

	// act
	request.SetRequestURI("/customer/foobar")
	request.Header.SetMethod(fasthttp.MethodGet)
	request.SetHost("localhost")

	_ = client.Do(request, response)

	// assert
	assert.Equal(t, http.StatusNotFound, response.Header.StatusCode())
	assert.Equal(t, `{"error":{"status":404,"message":"Not Found"}}`, string(response.Body()))
}
