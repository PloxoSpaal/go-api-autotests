package grpcfixtures

import (
	"github.com/Nikita-Filonov/api-go-autotests-server/gen/go/v1"
	"github.com/Nikita-Filonov/axiom"
	"github.com/stretchr/testify/require"
)

const SessionFixtureKey = "grpc-session"

type SessionFixture struct {
	Request  *v1.LoginRequest
	Response *v1.LoginResponse
}

func SetSessionFixture(cfg *axiom.Config) (any, func(), error) {
	userFixture := GetUserFixture(cfg)
	authenticationClient := GetAuthenticationClientFixture(cfg)

	fixture := SessionFixture{
		Request: &v1.LoginRequest{
			Email:    userFixture.Request.GetEmail(),
			Password: userFixture.Request.GetPassword(),
		},
	}

	cfg.Setup("Log in fixture user", func() {
		var err error
		fixture.Response, err = authenticationClient.Login(cfg.Context.Raw, fixture.Request)
		require.NoError(cfg.T(), err)
		require.NotNil(cfg.T(), fixture.Response)
		require.NotNil(cfg.T(), fixture.Response.GetToken())
		require.NotEmpty(cfg.T(), fixture.Response.GetToken().GetAccessToken())
		require.NotEmpty(cfg.T(), fixture.Response.GetToken().GetRefreshToken())
	})

	return fixture, nil, nil
}

func GetSessionFixture(cfg *axiom.Config) SessionFixture {
	return axiom.GetFixture[SessionFixture](cfg, SessionFixtureKey)
}
