package grpcfixtures

import (
	"github.com/Nikita-Filonov/axiom"
	"github.com/PloxoSpaal/go-api-autotests/clients/grpcclients"
)

const AuthenticationClientFixtureKey = "grpc-authentication-client"

func SetAuthenticationClientFixture(cfg *axiom.Config) (any, func(), error) {
	connection := GetPublicConnectionFixture(cfg)
	return grpcclients.NewAuthenticationClient(connection), nil, nil
}

func GetAuthenticationClientFixture(cfg *axiom.Config) *grpcclients.AuthenticationClient {
	return axiom.GetFixture[*grpcclients.AuthenticationClient](cfg, AuthenticationClientFixtureKey)
}
