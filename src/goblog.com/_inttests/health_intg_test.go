package _inttests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"net/http"
	"strings"
)

var _ bool = Describe("The API server", func() {
	Context("While running and reachable", func() {
		It("Responds to the health check route", func() {
			response := get("/status")

			Expect(response.StatusCode).To(Equal(http.StatusOK))

			expected := `{"gtg":true,"err":"none"}`
			body, err := ioutil.ReadAll(response.Body)
			defer response.Body.Close()

			Expect(err).To(BeNil(), "Reading body from response failed")
			Expect(strings.TrimSpace(string(body))).To(Equal(expected))
		})
	})
})
