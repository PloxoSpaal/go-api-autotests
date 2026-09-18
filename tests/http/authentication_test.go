package http

import (
	"net/http"

	"github.com/Nikita-Filonov/axiom"
	"github.com/PloxoSpaal/go-api-autotests/metadata"
	"github.com/PloxoSpaal/go-api-autotests/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func (s *suite) TestLoginUser() {
	testCase := axiom.NewCase(
		axiom.WithCaseName("Login user"),
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
		require.NoError(cfg.T(), err)
		require.Equal(cfg.T(), http.StatusOK, response.StatusCode)

		assert.Equal(cfg.T(), "bearer", response.Data.Token.TokenType)
		assert.NotEmpty(cfg.T(), response.Data.Token.AccessToken)
		assert.NotEmpty(cfg.T(), response.Data.Token.RefreshToken)

		cfg.T().Logf("user %s successfully logged in", request.Email)
	}))
}
