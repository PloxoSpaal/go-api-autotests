package httpfixtures

import (
	"github.com/Nikita-Filonov/axiom"
	"github.com/PloxoSpaal/go-api-autotests/clients/httpclients"
)

const PublicUsersClientFixtureKey = "http-users-client"

func SetPublicUsersClientFixture(cfg *axiom.Config) (any, func(), error) {
	transport := GetPublicTransportFixture(cfg)
	return httpclients.NewUsersClient(transport), nil, nil
}

func GetPublicUsersClientFixture(cfg *axiom.Config) *httpclients.UsersClient {
	return axiom.GetFixture[*httpclients.UsersClient](cfg, PublicUsersClientFixtureKey)
}
