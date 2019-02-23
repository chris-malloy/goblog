package _unittests

import (
	"errors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	log "github.com/sirupsen/logrus"
	"goblog.com/api/utils"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("Render HTTP Errors functions", func() {
	BeforeEach(func() {
		log.SetOutput(ioutil.Discard)
	})

	Context("When given complete inputs", func() {
		It("Should render the error correctly", func() {

			expectedStatus := http.StatusBadRequest
			expectedBody := `{ "status": 400, "message": "test error" }`

			recorder := httptest.NewRecorder()
			utils.RenderErrorWithStatus(recorder, errors.New("test error"), expectedStatus)

			Expect(recorder.Code).To(Equal(expectedStatus))
			Expect(recorder.Body.String()).To(Equal(expectedBody))
		})
	})

	Context("When given incomplete inputs", func() {
		It("Should render a default error message", func() {
			expectedStatus := http.StatusBadRequest
			expectedBody := `{ "status": 400, "message": "no error specified" }`

			recorder := httptest.NewRecorder()
			utils.RenderErrorWithStatus(recorder, nil, expectedStatus)

			Expect(recorder.Code).To(Equal(expectedStatus))
			Expect(recorder.Body.String()).To(Equal(expectedBody))
		})
	})

	Context("When not passed a request writer", func() {
		It("Panics", func() {
			Expect(func() {
				utils.RenderErrorWithStatus(nil, errors.New("test error"), 400)
			}).To(Panic())
		})
	})

	Context("When passed a validation error", func() {
		It("Renders a bad request message", func() {
			expectedStatus := http.StatusBadRequest
			expectedBody := `{ "status": 400, "message": "test error" }`

			recorder := httptest.NewRecorder()
			utils.RenderErrorAndDeriveStatus(recorder, utils.NewValidationError("test error"))

			Expect(recorder.Code).To(Equal(expectedStatus))
			Expect(recorder.Body.String()).To(Equal(expectedBody))
		})
	})

	Context("When passed any other error", func() {
		It("Renders a server error", func() {
			expectedStatus := http.StatusInternalServerError
			expectedBody := `{ "status": 500, "message": "test error" }`

			recorder := httptest.NewRecorder()
			utils.RenderErrorAndDeriveStatus(recorder, errors.New("test error"))

			Expect(recorder.Code).To(Equal(expectedStatus))
			Expect(recorder.Body.String()).To(Equal(expectedBody))
		})
	})
})

