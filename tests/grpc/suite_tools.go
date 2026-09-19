package grpc

import (
	"github.com/Nikita-Filonov/axiom"
	"github.com/PloxoSpaal/go-api-autotests/assertions"
	"github.com/PloxoSpaal/go-api-autotests/builders"
	"github.com/PloxoSpaal/go-api-autotests/clients/grpcclients"
	"github.com/PloxoSpaal/go-api-autotests/fake"
	"github.com/PloxoSpaal/go-api-autotests/fixtures"
	"github.com/PloxoSpaal/go-api-autotests/fixtures/grpcfixtures"
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

func (t *suiteTools) FilesClient() *grpcclients.FilesClient {
	return grpcfixtures.GetFilesClientFixture(t.cfg)
}

func (t *suiteTools) CoursesClient() *grpcclients.CoursesClient {
	return grpcfixtures.GetCoursesClientFixture(t.cfg)
}

func (t *suiteTools) ExercisesClient() *grpcclients.ExercisesClient {
	return grpcfixtures.GetExercisesClientFixture(t.cfg)
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

func (t *suiteTools) File() grpcfixtures.FileFixture {
	return grpcfixtures.GetFileFixture(t.cfg)
}

func (t *suiteTools) Course() grpcfixtures.CourseFixture {
	return grpcfixtures.GetCourseFixture(t.cfg)
}

func (t *suiteTools) Exercise() grpcfixtures.ExerciseFixture {
	return grpcfixtures.GetExerciseFixture(t.cfg)
}

var suiteToolset = axiom.NewToolset("grpc.tools", newSuiteTools)
