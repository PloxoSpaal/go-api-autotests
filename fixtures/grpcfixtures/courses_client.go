package grpcfixtures

import (
	"github.com/Nikita-Filonov/axiom"
	"github.com/PloxoSpaal/go-api-autotests/clients/grpcclients"
)

const CoursesClientFixtureKey = "grpc-courses-client"

func SetCoursesClientFixture(cfg *axiom.Config) (any, func(), error) {
	connection := GetPrivateConnectionFixture(cfg)
	return grpcclients.NewCoursesClient(connection), nil, nil
}

func GetCoursesClientFixture(cfg *axiom.Config) *grpcclients.CoursesClient {
	return axiom.GetFixture[*grpcclients.CoursesClient](cfg, CoursesClientFixtureKey)
}
