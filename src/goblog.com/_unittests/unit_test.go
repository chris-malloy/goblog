package _unittests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"goblog.com/api/handlers"
	"goblog.com/api/utils"
	"testing"
)

func TestAPI(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Goblog API Unit Tests Suite")
}

func emailValidator(testCase string, shouldValidate bool) {
	newValidator := handlers.ValidateEmail(testCase)
	Expect(newValidator.Ok).To(Equal(shouldValidate))

	if !shouldValidate {
		Expect(newValidator.ErrMsg).ToNot(BeNil())
		Expect(newValidator.ErrMsg).To(BeAssignableToTypeOf(utils.ValidationError{}))
	}
}
