package grpcfixtures

import (
	"github.com/Nikita-Filonov/axiom"
	"github.com/PloxoSpaal/go-api-autotests/clients/grpcclients"
)

const (
	PublicUsersClientFixtureKey  = "grpc-users-client"
	PrivateUsersClientFixtureKey = "grpc-private-users-client"
)

func SetPublicUsersClientFixture(cfg *axiom.Config) (any, func(), error) {
	return grpcclients.NewUsersClient(GetPublicConnectionFixture(cfg)), nil, nil
}

func SetPrivateUsersClientFixture(cfg *axiom.Config) (any, func(), error) {
	return grpcclients.NewUsersClient(GetPrivateConnectionFixture(cfg)), nil, nil
}

func GetPublicUsersClientFixture(cfg *axiom.Config) *grpcclients.UsersClient {
	return axiom.GetFixture[*grpcclients.UsersClient](cfg, PublicUsersClientFixtureKey)
}

func GetPrivateUsersClientFixture(cfg *axiom.Config) *grpcclients.UsersClient {
	return axiom.GetFixture[*grpcclients.UsersClient](cfg, PrivateUsersClientFixtureKey)
}
