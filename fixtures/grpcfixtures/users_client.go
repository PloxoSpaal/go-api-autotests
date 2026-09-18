package grpcfixtures

import (
	"github.com/Nikita-Filonov/axiom"
	"github.com/PloxoSpaal/go-api-autotests/clients/grpcclients"
)

const PublicUsersClientFixtureKey = "grpc-users-client"

func SetPublicUsersClientFixture(cfg *axiom.Config) (any, func(), error) {
	connection := GetPublicConnectionFixture(cfg)
	return grpcclients.NewUsersClient(connection), nil, nil
}

func GetPublicUsersClientFixture(cfg *axiom.Config) *grpcclients.UsersClient {
	return axiom.GetFixture[*grpcclients.UsersClient](cfg, PublicUsersClientFixtureKey)
}
