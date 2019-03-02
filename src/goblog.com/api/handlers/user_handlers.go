package handlers

import (
	. "goblog.com/api/utils"
	"strings"
)

func ValidateEmail(email string) Validator {
	if isLongEnough(email) {
		return Validator{Ok: false, ErrMsg: NewValidationError("email cannot be empty")}
	}

	if containsDesiredCharacters(email) {
		return Validator{Ok: false, ErrMsg: NewValidationError("email must contain `@` symbol")}
	}

	return Validator{Ok: true, ErrMsg: nil}
}

func isLongEnough(string string) bool {
	return len(string) <= 0
}

func containsDesiredCharacters(string string) bool {
	return !strings.Contains(string, "@")
}
