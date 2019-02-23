package _unittests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"goblog.com/api/db"
)

var _ = Describe("The Creds Function", func() {
	Context("When not passed all necessary creds", func() {
		It("Should fail with an error.", func() {
			_, err := db.GetCredsFromEnv()
			Expect(err).ToNot(BeNil())
		})
	})
})
