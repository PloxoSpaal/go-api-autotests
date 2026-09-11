package tests

import (
	"github.com/Nikita-Filonov/axiom"
	"github.com/stretchr/testify/assert"
)

type createUserParams struct {
	Email string
}

func (s *AxiomSuite) TestUserCanBeCreatedWithDifferentEmails() {
	testCases := []axiom.Case{
		axiom.NewCase(
			axiom.WithCaseName("user can be created with student email"),
			axiom.WithCaseParams(createUserParams{
				Email: "student@example.com",
			}),
			axiom.WithCaseContext(
				axiom.WithContextRaw(s.T().Context()),
				axiom.WithContextData("operation", "create-user"),
			),
		),
		axiom.NewCase(
			axiom.WithCaseName("user can be created with teacher email"),
			axiom.WithCaseParams(createUserParams{
				Email: "teacher@example.com",
			}),
			axiom.WithCaseContext(
				axiom.WithContextRaw(s.T().Context()),
				axiom.WithContextData("operation", "create-user"),
			),
		),
	}

	for _, testCase := range testCases {
		s.RunCase(testCase, usersToolset.Action(func(cfg *axiom.Config, tools *axiomUsersTools) {
			params := axiom.GetParams[createUserParams](cfg)

			var user *axiomUser

			cfg.Step("create user", func() {
				user = tools.Client.Create(axiomUserData{
					Email: params.Email,
				})
			})

			cfg.Step("check created user", func() {
				assert.Equal(cfg.T(), params.Email, user.Email)
				assert.True(cfg.T(), user.Active)
			})

			cfg.Teardown("delete user", func() {
				tools.Client.Delete(user)
			})
		}))
	}

}
