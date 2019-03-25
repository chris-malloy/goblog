package _inttests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"goblog.com/api/models"
)

var _ = Describe("The Authorization Module", func() {
	Context("When given a handle to a seeded database", func() {
		var userId int64
		db := getAndPingDB()
		authorizer, _ := models.NewAuthorizer(db)

		BeforeEach(func() {
			userId = insertTestUser(db)
		})

		It("Should log users in with a valid password.", func() {
			ok, err := authorizer.Authenticate(models.LoginRequest{
				Email: "christopher.malloy@7factor.io",
				Password: "abc123",
			})

			Expect(err).To(BeNil())
			Expect(ok).To(BeTrue())
		})

		It("Should not log users in and fail when given an invalid password.", func() {
			ok, err := authorizer.Authenticate(models.LoginRequest{
				Email: "christopher.malloy@7factor.io",
				Password: "badpassword",
			})

			Expect(err).ToNot(BeNil())
			Expect(ok).To(BeFalse())
		})

		It("Should not log users in and fail when given and unknown user email.", func() {
			ok, err := authorizer.Authenticate(models.LoginRequest{
				Email: "thisdoesnotexist@nowhere.com",
				Password: "notthepassword",
			})

			Expect(err).ToNot(BeNil())
			Expect(ok).To(BeFalse())
		})

		AfterEach(func() {
			clearTable("users", db)
		})
	})
})
