package _unittests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"goblog.com/api/handlers"
	. "goblog.com/api/utils"
	"testing"
)

func TestAPI(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Goblog API Unit Tests Suite")
}

func checkValidator(validator Validator, shouldValidate bool) {
	Expect(validator.Ok).To(Equal(shouldValidate))

	if !shouldValidate {
		Expect(validator.ErrMsg).ToNot(BeNil())
		Expect(validator.ErrMsg).To(BeAssignableToTypeOf(ValidationError{}))
	}
}

func emailValidatorCallback(testCase string, shouldValidate bool) {
	newValidator := handlers.ValidateEmail(testCase)
	checkValidator(newValidator, shouldValidate)
}

func passwordValidatorCallback(testCase string, shouldValidate bool) {
	newValidator := handlers.ValidatePassword(testCase)
	checkValidator(newValidator, shouldValidate)
}
