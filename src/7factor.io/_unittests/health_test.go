package _unittests

import (
	"7factor.io/api"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("Health check handler", func() {
	Context("When called with a valid response", func() {
		It("Should respond appropriately", func() {
			request, err := http.NewRequest("GET", "/health-check", nil)
			Expect(err).To(BeNil())

			recorder := httptest.NewRecorder()
			handler := api.HealthCheck()
			handler.ServeHTTP(recorder, request)

			Expect(recorder.Code).To(Equal(http.StatusOK))
		})
	})
})