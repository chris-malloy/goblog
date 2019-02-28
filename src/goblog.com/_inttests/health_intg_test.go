package _inttests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"net/http"
	"strings"
)

var _ = Describe("The API server", func() {
	Context("While running and reachable", func() {
		It("Responds to the health check route", func() {
			response := get("/status")

			Expect(response.StatusCode).To(Equal(http.StatusOK))

			expected := `{"up":true,"err":"none"}`
			body, err := ioutil.ReadAll(response.Body)
			defer response.Body.Close()

			Expect(err).To(BeNil(), "reading body from response failed")
			Expect(strings.TrimSpace(string(body))).To(Equal(expected))
		})

		It("Responds to the DBAccess route", func() {
			response := get("/dbaccess")

			Expect(response.StatusCode).To(Equal(http.StatusOK))

			expected := `{"up":true,"err":"none"}`
			body, err := ioutil.ReadAll(response.Body)
			defer response.Body.Close()

			Expect(err).To(BeNil(), "reading body from response failed")
			Expect(strings.TrimSpace(string(body))).To(Equal(expected))
		})
	})
})
