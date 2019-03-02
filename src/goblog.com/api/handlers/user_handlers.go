package handlers

import (
	. "goblog.com/api/utils"
	"regexp"
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

var generalPasswordErrorMessage = "Password must contain at least one special character, a number, and an uppercase character."

func ValidatePassword(password string) Validator {
	if isLongEnough(password) {
		return Validator{Ok: false, ErrMsg: NewValidationError("password cannot be empty")}
	}

	hasUpperCase, err := regexp.Compile(`.*[A-Z]+.*`)
	if err != nil {
		return Validator{Ok: false, ErrMsg: NewValidationErrorFromError(err)}
	}

	if !hasUpperCase.MatchString(password) {
		return Validator{Ok: false, ErrMsg: NewValidationError(generalPasswordErrorMessage)}
	}

	return Validator{Ok: true, ErrMsg: nil}
}

func isLongEnough(string string) bool {
	return len(string) <= 0
}

func containsDesiredCharacters(string string) bool {
	return !strings.Contains(string, "@")
}
