package http

import (
	"net/http"

	"github.com/Nikita-Filonov/axiom"
	"github.com/PloxoSpaal/go-api-autotests/metadata"
	"github.com/PloxoSpaal/go-api-autotests/models"
)

func (s *suite) TestLoginUser() {
	testCase := axiom.NewCase(
		axiom.WithCaseID("HTTP-AUTHENTICATION-001"),
		axiom.WithCaseName("login with correct email and password"),
		axiom.WithCaseMeta(
			axiom.WithMetaTag(metadata.TagSmoke),
			axiom.WithMetaStory(metadata.StoryLogin),
			axiom.WithMetaSeverity(axiom.SeverityBlocker),
		),
	)

	s.RunCase(testCase, suiteToolset.Action(func(cfg *axiom.Config, tools *suiteTools) {
		userFixture := tools.User()
		request := models.LoginRequest{
			Email:    userFixture.Request.Email,
			Password: userFixture.Request.Password,
		}
		response, err := tools.AuthenticationClient().Login(cfg.Context.Raw, request)

		tools.Assertion.HTTPOK(response.StatusCode, err)
		tools.Assertion.HTTPLoginResponse(response.Data)
	}))
}

func (s *suite) TestLoginWithInvalidCredentials() {
	testCase := axiom.NewCase(
		axiom.WithCaseID("HTTP-AUTHENTICATION-002"),
		axiom.WithCaseName("login with incorrect email and password"),
		axiom.WithCaseMeta(
			axiom.WithMetaTag(metadata.TagNegative),
			axiom.WithMetaStory(metadata.StoryLogin),
			axiom.WithMetaSeverity(axiom.SeverityCritical),
		),
	)

	s.RunCase(testCase, suiteToolset.Action(func(cfg *axiom.Config, tools *suiteTools) {
		request := models.LoginRequest{
			Email:    tools.Fake.Email(),
			Password: tools.Fake.Password(),
		}
		response, err := tools.AuthenticationClient().Login(cfg.Context.Raw, request)

		tools.Assertion.NoError(err)
		tools.Assertion.HTTPStatus(response.StatusCode, http.StatusUnauthorized)
		tools.Assertion.HTTPError(response.APIError, "Invalid credentials")
	}))
}
