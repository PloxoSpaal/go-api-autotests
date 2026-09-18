package grpcfixtures

import (
	"github.com/Nikita-Filonov/axiom"
	"github.com/PloxoSpaal/go-api-autotests/clients/grpcclients"
)

const ExercisesClientFixtureKey = "grpc-exercises-client"

func SetExercisesClientFixture(cfg *axiom.Config) (any, func(), error) {
	connection := GetPrivateConnectionFixture(cfg)
	return grpcclients.NewExercisesClient(connection), nil, nil
}

func GetExercisesClientFixture(cfg *axiom.Config) *grpcclients.ExercisesClient {
	return axiom.GetFixture[*grpcclients.ExercisesClient](cfg, ExercisesClientFixtureKey)
}
