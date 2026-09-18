package http

import (
	"github.com/Nikita-Filonov/axiom"
	"github.com/PloxoSpaal/go-api-autotests/builders"
	"github.com/PloxoSpaal/go-api-autotests/clients/httpclients"
	"github.com/PloxoSpaal/go-api-autotests/fake"
	"github.com/PloxoSpaal/go-api-autotests/fixtures/httpfixtures"
	"github.com/PloxoSpaal/go-api-autotests/resources"
)

type suiteTools struct {
	cfg *axiom.Config

	Fake    *fake.Fake
	Builder *builders.Builder
}

func newSuiteTools(cfg *axiom.Config) *suiteTools {
	return &suiteTools{
		cfg:     cfg,
		Fake:    resources.GetFakeResource(cfg.Runner),
		Builder: resources.GetBuilderResource(cfg.Runner),
	}
}

func (t *suiteTools) PublicUsersClient() *httpclients.UsersClient {
	return httpfixtures.GetPublicUsersClientFixture(t.cfg)
}

func (t *suiteTools) PrivateUsersClient() *httpclients.UsersClient {
	return httpfixtures.GetPrivateUsersClientFixture(t.cfg)
}

func (t *suiteTools) AuthenticationClient() *httpclients.AuthenticationClient {
	return httpfixtures.GetAuthenticationClientFixture(t.cfg)
}

func (t *suiteTools) User() httpfixtures.UserFixture {
	return httpfixtures.GetUserFixture(t.cfg)
}

var suiteToolset = axiom.NewToolset("http.tools", newSuiteTools)
