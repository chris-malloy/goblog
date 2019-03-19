package _unittests

import (
	"database/sql"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"goblog.com/api/handlers"
	"goblog.com/api/models"
	"net/http"
	"net/http/httptest"
	"os"
	"time"
)

var _ = Describe("JWT Authentication middleware functions", func() {
	Context("When accessing restricted APIs", func() {
		BeforeEach(func() {
			err := os.Setenv("JWT_SECRET", "test-secret-yay")
			_ = fmt.Errorf("error: %v setting JWT_SECRET", err)
		})

		It("Doesn't allow access when called with an invalid authorization header.", func() {
			called := false
			handler := handlers.RequireLogin()
			server := handler(func(writer http.ResponseWriter, request *http.Request) {
				called = true
				writer.WriteHeader(http.StatusOK)
			})

			request, _ := http.NewRequest("GET", "/thisdoesnotmatter", nil)
			request.Header.Set("Authorization", "thisisinvalid")

			recorder := httptest.NewRecorder()
			server.ServeHTTP(recorder, request)

			Expect(called).To(BeFalse())
			Expect(recorder.Code).To(Equal(http.StatusUnauthorized))
		})

		It("Doesn't allow access when called without a JWT.", func() {
			called := false
			handler := handlers.RequireLogin()
			server := handler(func(writer http.ResponseWriter, request *http.Request) {
				called = true
				writer.WriteHeader(http.StatusOK)
			})

			request, _ := http.NewRequest("GET", "/thisdoesnotmattter", nil)
			recorder := httptest.NewRecorder()
			server.ServeHTTP(recorder, request)

			Expect(called).To(BeFalse())
			Expect(recorder.Code).To(Equal(http.StatusUnauthorized))
		})

		It("Allows access when called with a valid JWT.", func() {
			called := false
			handler := handlers.RequireLogin()
			server := handler(func(writer http.ResponseWriter, request *http.Request) {
				called = true
				writer.WriteHeader(http.StatusOK)
			})

			secret, _ := os.LookupEnv("JWT_SECRET")
			mockClaims := jwt.MapClaims{
				"sub": 1,
				"iat": time.Now().Unix(),
				"exp": time.Now().Add(time.Hour * 4).Unix(),
			}
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, mockClaims)
			tokenString, _ := token.SignedString([]byte(secret))

			request, _ := http.NewRequest("GET", "/thisdoesnotmatter", nil)
			request.Header.Set("Authorization", "Bearer "+tokenString)

			recorder := httptest.NewRecorder()
			server.ServeHTTP(recorder, request)

			Expect(called).To(BeTrue())
			Expect(recorder.Code).To(Equal(http.StatusOK))
		})
	})
})

var _ = Describe("The Authorizer", func() {
	Context("When creating a new auth manager", func() {
		It("Works with a valid database handle.", func() {
			db, _ := sql.Open("postgres", "postgres://nowhere:nowhere@localhost:5432/nothing")
			authorizer, err := models.NewAuthManager(db)
			Expect(err).To(BeNil())
			Expect(authorizer).ToNot(BeNil())
		})

		It("Fails with an invalid database handle.", func() {
			authorizer, err := models.NewAuthManager(nil)
			Expect(err).ToNot(BeNil())
			Expect(authorizer).To(BeNil())
		})
	})
})
