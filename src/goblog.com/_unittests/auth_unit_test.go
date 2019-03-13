package _unittests

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"goblog.com/api/handlers"
	"net/http"
	"net/http/httptest"
	"os"
)

var _ = Describe("JWT Authentication middleware functions", func() {
	Context("When accessing restricted APIs", func() {
		BeforeEach(func() {
			err := os.Setenv("JWT_SECRET", "test-secret-yay")
			_ = fmt.Errorf("error: %v setting JWT_SECRET", err)
		})

		It("Doesn't allow access when called with an invalid authorization header.", func() {
			handler := handlers.RequireLogin()
			server, called := mockCall(handler)

			request, _ := http.NewRequest("GET", "/thisdoesnotmatter", nil)
			recorder := httptest.NewRecorder()
			server.ServeHTTP(recorder, request)

			request.Header.Set("Authorization", "thisisinvalid")

			Expect(called).To(BeFalse())
			Expect(recorder.Code).To(Equal(http.StatusUnauthorized))
		})
	})
})
