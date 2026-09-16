package http

import (
	"net/http"

	"github.com/Nikita-Filonov/axiom"
	"github.com/PloxoSpaal/go-api-autotests/fixtures/httpfixtures"
	"github.com/PloxoSpaal/go-api-autotests/metadata"
	"github.com/PloxoSpaal/go-api-autotests/models"
	"github.com/PloxoSpaal/go-api-autotests/resources"
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

	s.RunCase(testCase, func(cfg *axiom.Config) {
		authenticationClient := httpfixtures.GetAuthenticationClientFixture(cfg)
		usersClient := httpfixtures.GetPublicUsersClientFixture(cfg)

		builder := resources.GetBuilderResource(cfg.Runner)
		user := builder.UserCreate()
		userRequest := user.HTTPRequest()

		createdUserResponse, err := usersClient.Create(cfg.Context.Raw, userRequest)

		require.NoError(cfg.T(), err)
		require.Equal(cfg.T(), http.StatusOK, createdUserResponse.StatusCode)
		require.NotEmpty(cfg.T(), createdUserResponse.Data.User.ID)

		request := models.LoginRequest{
			Email:    userRequest.Email,
			Password: userRequest.Password,
		}

		response, err := authenticationClient.Login(cfg.Context.Raw, request)

		require.NoError(cfg.T(), err)
		require.Equal(cfg.T(), http.StatusOK, response.StatusCode)

		assert.Equal(cfg.T(), "bearer", response.Data.Token.TokenType)
		assert.NotEmpty(cfg.T(), response.Data.Token.AccessToken)
		assert.NotEmpty(cfg.T(), response.Data.Token.RefreshToken)

		cfg.T().Logf("user %s successfully logged in", userRequest.Email)
	})
}
