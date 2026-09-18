package grpc

import (
	"github.com/Nikita-Filonov/api-go-autotests-server/gen/go/v1"
	"github.com/Nikita-Filonov/axiom"
	"github.com/PloxoSpaal/go-api-autotests/metadata"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func (s *suite) TestLogin() {
	testCase := axiom.NewCase(
		axiom.WithCaseName("login with correct email and password"),
		axiom.WithCaseMeta(
			axiom.WithMetaTag(metadata.TagSmoke),
			axiom.WithMetaStory(metadata.StoryLogin),
			axiom.WithMetaSeverity(axiom.SeverityBlocker),
		),
	)

	s.RunCase(testCase, suiteToolset.Action(func(cfg *axiom.Config, tools *suiteTools) {
		userFixture := tools.User()

		request := &v1.LoginRequest{
			Email:    userFixture.Request.GetEmail(),
			Password: userFixture.Request.GetPassword(),
		}
		response, err := tools.AuthenticationClient().Login(cfg.Context.Raw, request)

		require.NoError(cfg.T(), err)
		require.NotNil(cfg.T(), response)
		require.NotNil(cfg.T(), response.GetToken())

		assert.Equal(cfg.T(), "bearer", response.GetToken().GetTokenType())
		assert.NotEmpty(cfg.T(), response.GetToken().GetAccessToken())
		assert.NotEmpty(cfg.T(), response.GetToken().GetRefreshToken())

		cfg.T().Logf("user %s successfully logged in", request.GetEmail())
	}))
}
