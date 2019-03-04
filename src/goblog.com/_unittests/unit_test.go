package _unittests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"goblog.com/api/utils"
	"testing"
)

func TestAPI(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Goblog API Unit Tests Suite")
}

func checkValidator(validator utils.Validator, shouldValidate bool) {
	Expect(validator.Ok).To(Equal(shouldValidate))

	if !shouldValidate {
		Expect(validator.ErrMsg).ToNot(BeNil())
		Expect(validator.ErrMsg).To(BeAssignableToTypeOf(utils.ValidationError{}))
	}
}

func emailValidatorCallback(testCase string, shouldValidate bool) {
	newValidator := utils.ValidateEmail(testCase)
	checkValidator(newValidator, shouldValidate)
}

func passwordValidatorCallback(testCase string, shouldValidate bool) {
	newValidator := utils.ValidatePassword(testCase)
	checkValidator(newValidator, shouldValidate)
}
