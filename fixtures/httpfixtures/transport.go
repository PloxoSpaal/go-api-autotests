package httpfixtures

import (
	"github.com/Nikita-Filonov/axiom"
	"github.com/PloxoSpaal/go-api-autotests/resources"
	"github.com/PloxoSpaal/go-api-autotests/transport/httptransport"
)

const PublicTransportFixtureKey = "http-public-transport"
const PrivateTransportFixtureKey = "http-private-transport"

func SetPublicTransportFixture(cfg *axiom.Config) (any, func(), error) {
	settings := resources.GetConfigResource(cfg.Runner)
	return httptransport.NewPublic(cfg, settings.HTTP), nil, nil
}

func SetPrivateTransportFixture(cfg *axiom.Config) (any, func(), error) {
	session := GetSessionFixture(cfg)
	settings := resources.GetConfigResource(cfg.Runner)
	token := session.Response.Data.Token.AccessToken
	return httptransport.NewPrivate(cfg, settings.HTTP, token), nil, nil
}

func GetPublicTransportFixture(cfg *axiom.Config) *httptransport.Client {
	return axiom.GetFixture[*httptransport.Client](cfg, PublicTransportFixtureKey)
}

func GetPrivateTransportFixture(cfg *axiom.Config) *httptransport.Client {
	return axiom.GetFixture[*httptransport.Client](cfg, PrivateTransportFixtureKey)
}
