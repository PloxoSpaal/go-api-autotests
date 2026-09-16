package httpfixtures

import (
	"github.com/Nikita-Filonov/axiom"
	"github.com/PloxoSpaal/go-api-autotests/resources"
	"github.com/PloxoSpaal/go-api-autotests/transport/httptransport"
)

const PublicTransportFixtureKey = "http-public-transport"

func SetPublicTransportFixture(cfg *axiom.Config) (any, func(), error) {
	settings := resources.GetConfigResource(cfg.Runner)
	return httptransport.NewPublic(cfg, settings.HTTP), nil, nil
}

func GetPublicTransportFixture(cfg *axiom.Config) *httptransport.Client {
	return axiom.GetFixture[*httptransport.Client](cfg, PublicTransportFixtureKey)
}
