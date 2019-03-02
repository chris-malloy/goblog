package utils

type Validator struct {
	Ok     bool
	ErrMsg error
}

type ValidationError struct {
	message string
}

func (e ValidationError) Error() string {
	return e.message
}

func NewValidationError(text string) error {
	return ValidationError{text}
}

func NewValidationErrorFromError(err error) error {
	return ValidationError{err.Error()}
}
