package assertions

import (
	"github.com/Nikita-Filonov/api-go-autotests-server/gen/go/v1"
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

func (a *Assertion) GRPCCreateUserResponse(actual *v1.GetUserResponse, expected builders.UserCreate) {
	a.cfg.Step("Check create gRPC user response", func() {
		a.NotNil(actual, "create user response")
		a.GRPCUserCreated(actual.GetUser(), expected)
	})
}

func (a *Assertion) GRPCUserCreated(actual *v1.User, expected builders.UserCreate) {
	a.cfg.Step("Check created gRPC user", func() {
		a.NotNil(actual, "user")
		a.NotEmpty(actual.GetId(), "user ID")
		a.Equal(actual.GetEmail(), expected.Email, "user email")
		a.Equal(actual.GetLastName(), expected.LastName, "user last name")
		a.Equal(actual.GetFirstName(), expected.FirstName, "user first name")
		a.Equal(actual.GetMiddleName(), expected.MiddleName, "user middle name")
	})
}

func (a *Assertion) GRPCGetUserResponse(actual *v1.GetUserResponse, expected *v1.GetUserResponse) {
	a.cfg.Step("Check get gRPC user response", func() {
		a.NotNil(actual, "actual get user response")
		a.NotNil(expected, "expected get user response")
		a.GRPCUser(actual.GetUser(), expected.GetUser())
	})
}

func (a *Assertion) GRPCUser(actual, expected *v1.User) {
	a.cfg.Step("Check gRPC user", func() {
		a.NotNil(actual, "actual user")
		a.NotNil(expected, "expected user")
		a.Equal(actual.GetId(), expected.GetId(), "user ID")
		a.Equal(actual.GetEmail(), expected.GetEmail(), "user email")
		a.Equal(actual.GetLastName(), expected.GetLastName(), "user last name")
		a.Equal(actual.GetFirstName(), expected.GetFirstName(), "user first name")
		a.Equal(actual.GetMiddleName(), expected.GetMiddleName(), "user middle name")
	})
}

func (a *Assertion) GRPCUpdateUserResponse(actual *v1.GetUserResponse, expected builders.UserUpdate) {
	a.cfg.Step("Check update gRPC user response", func() {
		a.NotNil(actual, "update user response")
		a.GRPCUserUpdated(actual.GetUser(), expected)
	})
}

func (a *Assertion) GRPCUserUpdated(actual *v1.User, expected builders.UserUpdate) {
	a.cfg.Step("Check updated gRPC user", func() {
		a.NotNil(actual, "user")

		if &expected.Email != nil {
			a.Equal(actual.GetEmail(), &expected.Email, "user email")
		}
		if &expected.LastName != nil {
			a.Equal(actual.GetLastName(), &expected.LastName, "user last name")
		}
		if &expected.FirstName != nil {
			a.Equal(actual.GetFirstName(), &expected.FirstName, "user first name")
		}
		if &expected.MiddleName != nil {
			a.Equal(actual.GetMiddleName(), &expected.MiddleName, "user middle name")
		}
	})
}
