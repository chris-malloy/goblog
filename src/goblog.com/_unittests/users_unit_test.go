package _unittests

import (
	"database/sql"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"goblog.com/api/models"
)

var _ = Describe("User Functions", func() {
	Context("When creating a new user manager", func() {
		It("Returns the user manager when given a valid db handle.", func() {
			db, _ := sql.Open("postgres", "postgres://nowhere:nowhere@localhost:5432/nothing")
			userManager, err := models.NewUserManager(db)
			Expect(err).To(BeNil())
			Expect(userManager).ToNot(BeNil())
		})
	})
})

