package _inttests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"goblog.com/api/models"
)

var newUserRequest = models.NewUserRequest{Email: "chris@testuser.com", FirstName: "Chris", LastName: "Malloy", Password: "Abcde123@"}

var _ = Describe("The user module", func() {
	var userId int64
	db := getAndPingDB()

	BeforeEach(func() {
		userId = insertTestUser(db)
	})

	Context("When given a valid db handle", func() {
		userManger, _ := models.NewUserManager(db)

		It("Should add a new user.", func() {
			newUser, err := userManger.CreateUser(newUserRequest)
			Expect(err).To(BeNil())
			Expect(newUser.ID).ToNot(BeNil())
			Expect(newUser.Email).To(Equal("chris@testuser.com"))
		})
	})

	AfterEach(func() {
		clearTable("users", db)
	})
})
