package httpfixtures

import (
	"github.com/Nikita-Filonov/axiom"
	"github.com/PloxoSpaal/go-api-autotests/clients/httpclients"
)

const AuthenticationClientFixtureKey = "http-authentication-client"

func SetAuthenticationClientFixture(cfg *axiom.Config) (any, func(), error) {
	transport := GetPublicTransportFixture(cfg)
	return httpclients.NewAuthenticationClient(transport), nil, nil
}

func GetAuthenticationClientFixture(cfg *axiom.Config) *httpclients.AuthenticationClient {
	return axiom.GetFixture[*httpclients.AuthenticationClient](cfg, AuthenticationClientFixtureKey)
}
