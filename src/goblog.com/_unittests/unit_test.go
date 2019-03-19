package _unittests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"goblog.com/api/utils"
	"net/http"
	"net/http/httptest"
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

func mockRequestAndRecorder() (*http.Request, *httptest.ResponseRecorder) {
	request, _ := http.NewRequest("GET", "/thisdoesnotmatter", nil)
	recorder := httptest.NewRecorder()
	return request, recorder
}
