package httpfixtures

import (
	"net/http"

	"github.com/Nikita-Filonov/axiom"
	"github.com/PloxoSpaal/go-api-autotests/clients/httpclients"
	"github.com/PloxoSpaal/go-api-autotests/models"
	"github.com/stretchr/testify/require"
)

const SessionFixtureKey = "http-session"

type SessionFixture struct {
	Request  models.LoginRequest
	Response httpclients.Response[models.LoginResponse]
}

func SetSessionFixture(cfg *axiom.Config) (any, func(), error) {
	userFixture := GetUserFixture(cfg)
	authenticationClient := GetAuthenticationClientFixture(cfg)

	fixture := SessionFixture{
		Request: models.LoginRequest{
			Email:    userFixture.Request.Email,
			Password: userFixture.Request.Password,
		},
	}

	cfg.Setup("Log in fixture user", func() {
		var err error
		fixture.Response, err = authenticationClient.Login(cfg.Context.Raw, fixture.Request)
		require.NoError(cfg.T(), err)
		require.Equal(cfg.T(), http.StatusOK, fixture.Response.StatusCode)
		require.NotEmpty(cfg.T(), fixture.Response.Data.Token.AccessToken)
		require.NotEmpty(cfg.T(), fixture.Response.Data.Token.RefreshToken)
	})

	return fixture, nil, nil
}

func GetSessionFixture(cfg *axiom.Config) SessionFixture {
	return axiom.GetFixture[SessionFixture](cfg, SessionFixtureKey)
}
