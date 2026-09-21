package http

import (
	"github.com/Nikita-Filonov/axiom"
	"github.com/PloxoSpaal/go-api-autotests/assertions"
	"github.com/PloxoSpaal/go-api-autotests/builders"
	"github.com/PloxoSpaal/go-api-autotests/clients/httpclients"
	"github.com/PloxoSpaal/go-api-autotests/fake"
	"github.com/PloxoSpaal/go-api-autotests/fixtures"
	"github.com/PloxoSpaal/go-api-autotests/fixtures/httpfixtures"
	"github.com/PloxoSpaal/go-api-autotests/resources"
)

type suiteTools struct {
	cfg *axiom.Config

	Fake      *fake.Fake
	Builder   *builders.Builder
	Assertion *assertions.Assertion
}

func newSuiteTools(cfg *axiom.Config) *suiteTools {
	return &suiteTools{
		cfg:       cfg,
		Fake:      resources.GetFakeResource(cfg.Runner),
		Builder:   resources.GetBuilderResource(cfg.Runner),
		Assertion: fixtures.GetAssertionFixture(cfg),
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

func (t *suiteTools) File() httpfixtures.FileFixture {
	return httpfixtures.GetFileFixture(t.cfg)
}

func (t *suiteTools) FilesClient() *httpclients.FilesClient {
	return httpfixtures.GetFilesClientFixture(t.cfg)
}
func (t *suiteTools) CoursesClient() *httpclients.CoursesClient {
	return httpfixtures.GetCoursesClientFixture(t.cfg)
}
func (t *suiteTools) ExercisesClient() *httpclients.ExercisesClient {
	return httpfixtures.GetExercisesClientFixture(t.cfg)
}

var suiteToolset = axiom.NewToolset("http.tools", newSuiteTools)
