package grpc

import (
	"github.com/Nikita-Filonov/axiom"
	"github.com/PloxoSpaal/go-api-autotests/builders"
	"github.com/PloxoSpaal/go-api-autotests/clients/grpcclients"
	"github.com/PloxoSpaal/go-api-autotests/fake"
	"github.com/PloxoSpaal/go-api-autotests/fixtures/grpcfixtures"
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

func (t *suiteTools) PublicUsersClient() *grpcclients.UsersClient {
	return grpcfixtures.GetPublicUsersClientFixture(t.cfg)
}

func (t *suiteTools) PrivateUsersClient() *grpcclients.UsersClient {
	return grpcfixtures.GetPrivateUsersClientFixture(t.cfg)
}

func (t *suiteTools) AuthenticationClient() *grpcclients.AuthenticationClient {
	return grpcfixtures.GetAuthenticationClientFixture(t.cfg)
}

func (t *suiteTools) User() grpcfixtures.UserFixture {
	return grpcfixtures.GetUserFixture(t.cfg)
}

var suiteToolset = axiom.NewToolset("grpc.tools", newSuiteTools)
