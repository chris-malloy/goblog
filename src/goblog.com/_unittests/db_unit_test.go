package _unittests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"goblog.com/api/db"
	"os"
)

var _ = Describe("The Creds Function", func() {
	Context("When not passed all necessary creds", func() {
		It("Should fail with an error.", func() {
			_, err := db.GetCredsFromEnv()
			Expect(err).ToNot(BeNil())
		})
	})

	Context("When passed all necessary creds", func() {
		BeforeEach(func() {
			os.Setenv("DB_USER", "samus")
			os.Setenv("DB_PASS", "aran")
			os.Setenv("DB_HOST", "localhost")
			os.Setenv("DB_NAME", "metroid")
		})

		It("Should convert to a valid connection string", func() {
			creds, err := db.GetCredsFromEnv()
			Expect(err).To(BeNil(), "creds error should be nil")

			connStr := creds.ToConnectionString()
			Expect(connStr).To(Equal("postgres://samus:aran@localhost:5432/metroid"))
		})

		It("Should convert a valid connection string with options", func() {
			os.Setenv("DB_OPTN", "sslmode=disable")
			creds, err := db.GetCredsFromEnv()
			Expect(err).To(BeNil(), "creds error should be nil")

			connStr := creds.ToConnectionString()
			Expect(connStr).To(Equal("postgres://samus:aran@localhost:5432/metroid?sslmode=disable"))
		})

		AfterEach(func() {
			os.Unsetenv("DB_HOST")
			os.Unsetenv("DB_USER")
			os.Unsetenv("DB_PASS")
			os.Unsetenv("DB_NAME")
		})
	})
})
