package validator

type Validator interface {
	ValidateCreateAccount(requestData interface{}) error
}
