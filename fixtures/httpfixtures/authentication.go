package httpfixtures

import (
	"github.com/Nikita-Filonov/axiom"
	"github.com/PloxoSpaal/go-api-autotests/clients/httpclients"
	"github.com/PloxoSpaal/go-api-autotests/fixtures"
	"github.com/PloxoSpaal/go-api-autotests/models"
)

const SessionFixtureKey = "http-session"

type SessionFixture struct {
	Request  models.LoginRequest
	Response httpclients.Response[models.LoginResponse]
}

func SetSessionFixture(cfg *axiom.Config) (any, func(), error) {
	assertion := fixtures.GetAssertionFixture(cfg)
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
		assertion.HTTPOK(fixture.Response.StatusCode, err)
		assertion.HTTPLoginResponse(fixture.Response.Data)
	})

	return fixture, nil, nil
}

func GetSessionFixture(cfg *axiom.Config) SessionFixture {
	return axiom.GetFixture[SessionFixture](cfg, SessionFixtureKey)
}
