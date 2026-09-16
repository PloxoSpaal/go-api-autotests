package httpfixtures

import (
	"github.com/Nikita-Filonov/axiom"
	"github.com/PloxoSpaal/go-api-autotests/clients/httpclients"
)

const (
	PublicUsersClientFixtureKey  = "http-users-client"
	PrivateUsersClientFixtureKey = "http-private-users-client"
)

func SetPublicUsersClientFixture(cfg *axiom.Config) (any, func(), error) {
	transport := GetPublicTransportFixture(cfg)
	return httpclients.NewUsersClient(transport), nil, nil
}

func SetPrivateUsersClientFixture(cfg *axiom.Config) (any, func(), error) {
	transport := GetPrivateTransportFixture(cfg)
	return httpclients.NewUsersClient(transport), nil, nil
}

func GetPublicUsersClientFixture(cfg *axiom.Config) *httpclients.UsersClient {
	return axiom.GetFixture[*httpclients.UsersClient](cfg, PublicUsersClientFixtureKey)
}

func GetPrivateUsersClientFixture(cfg *axiom.Config) *httpclients.UsersClient {
	return axiom.GetFixture[*httpclients.UsersClient](cfg, PrivateUsersClientFixtureKey)
}
