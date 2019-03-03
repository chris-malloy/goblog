package _unittests

import (
	"database/sql"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	"goblog.com/api/models"
)

var _ = Describe("User Functions", func() {
	DescribeTable("When validating an email address", emailValidatorCallback,
		Entry("it validates a good email.", "christopher.malloy@7factor.io", true),
		Entry("it validates a weird email.", "thisisweird@ something.com", true),
		Entry("it errors with a nil email.", nil, false),
		Entry("it errors if it does not contain the `@` symbol.", "this is not an email", false),
	)

	DescribeTable("When validating a password", passwordValidatorCallback,
		Entry("it validates a good password.", "abCde1234@", true),
		Entry("it errors with a nil password.", nil, false),
		Entry("it errors with no uppercase characters.", "abcd1234$", false),
		Entry("it errors if it has no lowercase characters.", "ABCD1234$", false),
		Entry("it errors if it's shorter than required.", "Ab1@", false),
		Entry("it errors if it has no symbols.", "AbCd12345", false),
		Entry("it errors if it contains no numbers.", "AbCdefgh%", false),
	)

	Context("When creating a new user manager", func() {
		It("Returns the user manager when given a valid database handle.", func() {
			db, _ := sql.Open("postgres", "postgres://nowhere:nowhere@localhost:5432/nothing")
			userManager, err := models.NewUserManager(db)
			Expect(err).To(BeNil())
			Expect(userManager).ToNot(BeNil())
		})

		It("Errors with a nil database handle", func() {
			users, err := models.NewUserManager(nil)
			Expect(err).ToNot(BeNil())
			Expect(users).To(BeNil())
		})
	})
})
