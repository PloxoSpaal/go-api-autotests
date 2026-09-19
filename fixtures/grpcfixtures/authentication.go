package grpcfixtures

import (
	"github.com/Nikita-Filonov/api-go-autotests-server/gen/go/v1"
	"github.com/Nikita-Filonov/axiom"
	"github.com/PloxoSpaal/go-api-autotests/fixtures"
)

const SessionFixtureKey = "grpc-session"

type SessionFixture struct {
	Request  *v1.LoginRequest
	Response *v1.LoginResponse
}

func SetSessionFixture(cfg *axiom.Config) (any, func(), error) {
	assertion := fixtures.GetAssertionFixture(cfg)
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
		assertion.NoError(err)
		assertion.GRPCLoginResponse(fixture.Response)

	})

	return fixture, nil, nil
}

func GetSessionFixture(cfg *axiom.Config) SessionFixture {
	return axiom.GetFixture[SessionFixture](cfg, SessionFixtureKey)
}
