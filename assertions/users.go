package assertions

import (
	"github.com/Nikita-Filonov/api-go-autotests-server/gen/go/v1"
	"github.com/Nikita-Filonov/axiom"
	"github.com/PloxoSpaal/go-api-autotests/builders"
	"github.com/PloxoSpaal/go-api-autotests/models"
)

func (a *Assertion) HTTPCreateUserResponse(actual models.UserResponse, expected builders.UserCreate) {
	message := "Check create HTTP user response"
	a.cfg.Step(message, func() {
		a.cfg.Log(axiom.NewInfoLog(message))
		a.HTTPUserCreated(actual.User, expected)
	})
}

func (a *Assertion) HTTPUserCreated(actual models.User, expected builders.UserCreate) {
	message := "Check created HTTP user"
	a.cfg.Step(message, func() {
		a.cfg.Log(axiom.NewInfoLog(message))
		a.NotEmpty(actual.ID, "user ID")
		a.Equal(actual.Email, expected.Email, "user email")
		a.Equal(actual.LastName, expected.LastName, "user last name")
		a.Equal(actual.FirstName, expected.FirstName, "user first name")
		a.Equal(actual.MiddleName, expected.MiddleName, "user middle name")
	})
}

func (a *Assertion) HTTPGetUserResponse(actual, expected models.UserResponse) {
	message := "Check get HTTP user response"
	a.cfg.Step(message, func() {
		a.cfg.Log(axiom.NewInfoLog(message))
		a.HTTPUser(actual.User, expected.User)
	})
}

func (a *Assertion) HTTPUser(actual, expected models.User) {
	message := "Check HTTP user"
	a.cfg.Step(message, func() {
		a.cfg.Log(axiom.NewInfoLog(message))
		a.Equal(actual.ID, expected.ID, "user ID")
		a.Equal(actual.Email, expected.Email, "user email")
		a.Equal(actual.LastName, expected.LastName, "user last name")
		a.Equal(actual.FirstName, expected.FirstName, "user first name")
		a.Equal(actual.MiddleName, expected.MiddleName, "user middle name")
	})
}

func (a *Assertion) HTTPUpdateUserResponse(actual models.UserResponse, expected builders.UserUpdate) {
	message := "Check update HTTP user response"
	a.cfg.Step(message, func() {
		a.cfg.Log(axiom.NewInfoLog(message))
		a.HTTPUserUpdated(actual.User, expected)
	})
}

func (a *Assertion) HTTPUserUpdated(actual models.User, expected builders.UserUpdate) {
	message := "Check updated HTTP user"
	a.cfg.Step(message, func() {
		a.cfg.Log(axiom.NewInfoLog(message))
		if &expected.Email != nil {
			a.Equal(actual.Email, expected.Email, "user email")
		}
		if &expected.LastName != nil {
			a.Equal(actual.LastName, expected.LastName, "user last name")
		}
		if &expected.FirstName != nil {
			a.Equal(actual.FirstName, expected.FirstName, "user first name")
		}
		if &expected.MiddleName != nil {
			a.Equal(actual.MiddleName, expected.MiddleName, "user middle name")
		}
	})
}

func (a *Assertion) GRPCCreateUserResponse(actual *v1.GetUserResponse, expected builders.UserCreate) {
	message := "Check create gRPC user response"
	a.cfg.Step(message, func() {
		a.cfg.Log(axiom.NewInfoLog(message))
		a.NotNil(actual, "create user response")
		a.GRPCUserCreated(actual.GetUser(), expected)
	})
}

func (a *Assertion) GRPCUserCreated(actual *v1.User, expected builders.UserCreate) {
	message := "Check created gRPC user"
	a.cfg.Step(message, func() {
		a.cfg.Log(axiom.NewInfoLog(message))
		a.NotNil(actual, "user")
		a.NotEmpty(actual.GetId(), "user ID")
		a.Equal(actual.GetEmail(), expected.Email, "user email")
		a.Equal(actual.GetLastName(), expected.LastName, "user last name")
		a.Equal(actual.GetFirstName(), expected.FirstName, "user first name")
		a.Equal(actual.GetMiddleName(), expected.MiddleName, "user middle name")
	})
}

func (a *Assertion) GRPCGetUserResponse(actual, expected *v1.GetUserResponse) {
	message := "Check get gRPC user response"
	a.cfg.Step(message, func() {
		a.cfg.Log(axiom.NewInfoLog(message))
		a.NotNil(actual, "actual get user response")
		a.NotNil(expected, "expected get user response")
		a.GRPCUser(actual.GetUser(), expected.GetUser())
	})
}

func (a *Assertion) GRPCUser(actual, expected *v1.User) {
	message := "Check gRPC user"
	a.cfg.Step(message, func() {
		a.cfg.Log(axiom.NewInfoLog(message))
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
	message := "Check update gRPC user response"
	a.cfg.Step(message, func() {
		a.cfg.Log(axiom.NewInfoLog(message))
		a.NotNil(actual, "update user response")
		a.GRPCUserUpdated(actual.GetUser(), expected)
	})
}

func (a *Assertion) GRPCUserUpdated(actual *v1.User, expected builders.UserUpdate) {
	message := "Check updated gRPC user"
	a.cfg.Step(message, func() {
		a.cfg.Log(axiom.NewInfoLog(message))
		a.NotNil(actual, "user")

		if &expected.Email != nil {
			a.Equal(actual.GetEmail(), expected.Email, "user email")
		}
		if &expected.LastName != nil {
			a.Equal(actual.GetLastName(), expected.LastName, "user last name")
		}
		if &expected.FirstName != nil {
			a.Equal(actual.GetFirstName(), expected.FirstName, "user first name")
		}
		if &expected.MiddleName != nil {
			a.Equal(actual.GetMiddleName(), expected.MiddleName, "user middle name")
		}
	})
}
