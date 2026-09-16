package httpfixtures

import (
	"github.com/Nikita-Filonov/axiom"
	"github.com/PloxoSpaal/go-api-autotests/clients/httpclients"
)

const CoursesClientFixtureKey = "http-courses-client"

func SetCoursesClientFixture(cfg *axiom.Config) (any, func(), error) {
	transport := GetPrivateTransportFixture(cfg)
	return httpclients.NewCoursesClient(transport), nil, nil
}

func GetCoursesClientFixture(cfg *axiom.Config) *httpclients.CoursesClient {
	return axiom.GetFixture[*httpclients.CoursesClient](cfg, CoursesClientFixtureKey)
}
