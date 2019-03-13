package _unittests

import (
	"github.com/husobee/vestigo"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"goblog.com/api/utils"
	"net/http"
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

func mockCall(handler vestigo.Middleware) (http.HandlerFunc, bool) {
	called := false
	mockCall := func(writer http.ResponseWriter, request *http.Request) {
		called = true
	}
	return handler(mockCall), called
}

