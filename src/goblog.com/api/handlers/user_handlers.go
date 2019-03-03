package handlers

import (
	. "goblog.com/api/utils"
	"regexp"
	"strings"
)

func ValidateEmail(email string) Validator {
	if isLongEnough(email, 0) {
		return Validator{Ok: false, ErrMsg: NewValidationError("email cannot be empty")}
	}

	if containsDesiredCharacters(email) {
		return Validator{Ok: false, ErrMsg: NewValidationError("email must contain `@` symbol")}
	}

	return Validator{Ok: true, ErrMsg: nil}
}

var generalPasswordErrorMessage = "Password must contain at least one special character, a number, a lowercase character, and an uppercase character."

func ValidatePassword(password string) Validator {
	if isLongEnough(password, 8) {
		return Validator{Ok: false, ErrMsg: NewValidationError("Password must be at least 8 characters long.")}
	}

	upperCaseRegex, err := compileRegex(`.*[A-Z]+.*`)
	if err != nil {
		return Validator{Ok: false, ErrMsg: NewValidationErrorFromError(err)}
	}

	lowerCaseRegex, err := compileRegex(`.*[a-z]+.*`)
	if err != nil {
		return Validator{Ok: false, ErrMsg: NewValidationErrorFromError(err)}
	}

	symbolRegex, err := compileRegex(`.*\W+.*`)
	if err != nil {
		return Validator{Ok: false, ErrMsg: NewValidationErrorFromError(err)}
	}

	numberRegex, err := compileRegex(`.*\d+.*`)
	if err != nil {
		return Validator{Ok: false, ErrMsg: NewValidationErrorFromError(err)}
	}

	if !hasRegex(upperCaseRegex, password) || !hasRegex(lowerCaseRegex, password) ||
		!hasRegex(symbolRegex, password) || !hasRegex(numberRegex, password) {
		return Validator{Ok: false, ErrMsg: NewValidationError(generalPasswordErrorMessage)}
	}

	return Validator{Ok: true, ErrMsg: nil}
}

func isLongEnough(string string, length int) bool {
	return len(string) <= length
}

func containsDesiredCharacters(string string) bool {
	return !strings.Contains(string, "@")
}

func compileRegex(regexMatcher string) (*regexp.Regexp, error) {
	return regexp.Compile(regexMatcher)
}

func hasRegex(compliedRegex *regexp.Regexp, string string) bool {
	return compliedRegex.MatchString(string)
}
