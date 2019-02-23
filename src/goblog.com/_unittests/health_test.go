package _unittests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"goblog.com/api/handlers"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("Health check handler", func() {
	Context("When called with a valid response", func() {
		It("Should respond appropriately", func() {
			request, err := http.NewRequest("GET", "/health-check", nil)
			Expect(err).To(BeNil())

			recorder := httptest.NewRecorder()
			handler := handlers.HealthCheck()
			handler.ServeHTTP(recorder, request)

			Expect(recorder.Code).To(Equal(http.StatusOK))
		})
	})
})