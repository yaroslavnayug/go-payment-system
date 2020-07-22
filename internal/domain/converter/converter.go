package converter

import "github.com/yaroslavnayug/go-payment-system/internal/domain/model"

type Converter interface {
	ConvertToAccount(requestData interface{}) (*model.Account, error)
}
