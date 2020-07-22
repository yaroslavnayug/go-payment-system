package model

type ValidationError struct {
	s string
}

func NewValidationError(text string) error {
	return &ValidationError{text}
}

func (e *ValidationError) Error() string {
	return e.s
}
