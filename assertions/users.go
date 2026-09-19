package assertions

import (
	"github.com/PloxoSpaal/go-api-autotests/builders"
	"github.com/PloxoSpaal/go-api-autotests/models"
)

func (a *Assertion) HTTPCreateUserResponse(actual models.UserResponse, expected builders.UserCreate) {
	a.cfg.Step("Check create HTTP user response", func() {
		a.HTTPUserCreated(actual.User, expected)
	})
}

func (a *Assertion) HTTPUserCreated(actual models.User, expected builders.UserCreate) {
	a.cfg.Step("Check created HTTP user", func() {
		a.NotEmpty(actual.ID, "user ID")
		a.Equal(actual.Email, expected.Email, "user email")
		a.Equal(actual.LastName, expected.LastName, "user last name")
		a.Equal(actual.FirstName, expected.FirstName, "user first name")
		a.Equal(actual.MiddleName, expected.MiddleName, "user middle name")
	})
}

func (a *Assertion) HTTPGetUserResponse(actual models.UserResponse, expected models.UserResponse) {
	a.cfg.Step("Check get HTTP user response", func() {
		a.HTTPUser(actual.User, expected.User)
	})
}

func (a *Assertion) HTTPUser(actual, expected models.User) {
	a.cfg.Step("Check HTTP user", func() {
		a.Equal(actual.ID, expected.ID, "user ID")
		a.Equal(actual.Email, expected.Email, "user email")
		a.Equal(actual.LastName, expected.LastName, "user last name")
		a.Equal(actual.FirstName, expected.FirstName, "user first name")
		a.Equal(actual.MiddleName, expected.MiddleName, "user middle name")
	})
}

func (a *Assertion) HTTPUpdateUserResponse(actual models.UserResponse, expected builders.UserUpdate) {
	a.cfg.Step("Check update HTTP user response", func() {
		a.HTTPUserUpdated(actual.User, expected)
	})
}

func (a *Assertion) HTTPUserUpdated(actual models.User, expected builders.UserUpdate) {
	a.cfg.Step("Check updated HTTP user", func() {
		if &expected.Email != nil {
			a.Equal(actual.Email, &expected.Email, "user email")
		}
		if &expected.LastName != nil {
			a.Equal(actual.LastName, &expected.LastName, "user last name")
		}
		if &expected.FirstName != nil {
			a.Equal(actual.FirstName, &expected.FirstName, "user first name")
		}
		if &expected.MiddleName != nil {
			a.Equal(actual.MiddleName, &expected.MiddleName, "user middle name")
		}
	})
}
