package httpfixtures

import (
	"github.com/Nikita-Filonov/axiom"
	"github.com/PloxoSpaal/go-api-autotests/clients/httpclients"
)

const ExercisesClientFixtureKey = "http-exercises-client"

func SetExercisesClientFixture(cfg *axiom.Config) (any, func(), error) {
	transport := GetPrivateTransportFixture(cfg)
	return httpclients.NewExercisesClient(transport), nil, nil
}

func GetExercisesClientFixture(cfg *axiom.Config) *httpclients.ExercisesClient {
	return axiom.GetFixture[*httpclients.ExercisesClient](cfg, ExercisesClientFixtureKey)
}
