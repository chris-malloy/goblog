package _unittests

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"goblog.com/api/healthcheck"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("Health check handlers", func() {
	Context("When calleding the health check route with a valid response", func() {
		It("Should respond appropriately.", func() {
			request, err := http.NewRequest("GET", "/health-check", nil)
			Expect(err).To(BeNil())

			recorder := httptest.NewRecorder()
			handler := healthcheck.HealthCheck()
			handler.ServeHTTP(recorder, request)

			Expect(recorder.Code).To(Equal(http.StatusOK))
		})
	})

	Context("When calling the db ping route with a valid response", func() {
		It("Should respond appropriately.", func() {
			request, err := http.NewRequest("GET", "/dbaccess", nil)
			Expect(err).To(BeNil())

			recorder := httptest.NewRecorder()

			mockDB, _, err := sqlmock.New()
			Expect(err).To(BeNil(), fmt.Sprintf("an error '%s' was not expected when opening a stub database connection", err))
			defer mockDB.Close()

			Expect(err).To(BeNil())

			handler := healthcheck.DBAccess(mockDB)
			handler.ServeHTTP(recorder, request)

			Expect(recorder.Code).To(Equal(http.StatusOK))
		})
	})
})
