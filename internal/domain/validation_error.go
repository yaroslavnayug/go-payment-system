package domain

type ValidationError struct {
	errStr string
}

func NewValidationError(text string) error {
	return &ValidationError{text}
}

func (e *ValidationError) Error() string {
	return e.errStr
}
