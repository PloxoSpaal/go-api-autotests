package grpc

import (
	"github.com/Nikita-Filonov/axiom"
	"github.com/PloxoSpaal/go-api-autotests/fixtures/grpcfixtures"
	"github.com/PloxoSpaal/go-api-autotests/metadata"
	"github.com/PloxoSpaal/go-api-autotests/tests"
)

var suiteRunner = tests.Runner.Join(
	axiom.NewRunner(
		axiom.WithRunnerMeta(
			axiom.WithMetaTag(metadata.TagGRPC),
			axiom.WithMetaSuite(metadata.SuiteGRPC),
		),
		axiom.WithRunnerFixture(grpcfixtures.PublicConnectionFixtureKey, grpcfixtures.SetPublicConnectionFixture),
		axiom.WithRunnerFixture(grpcfixtures.PrivateConnectionFixtureKey, grpcfixtures.SetPrivateConnectionFixture),

		axiom.WithRunnerFixture(grpcfixtures.PublicUsersClientFixtureKey, grpcfixtures.SetPublicUsersClientFixture),
		axiom.WithRunnerFixture(grpcfixtures.PrivateUsersClientFixtureKey, grpcfixtures.SetPrivateUsersClientFixture),
		axiom.WithRunnerFixture(grpcfixtures.AuthenticationClientFixtureKey, grpcfixtures.SetAuthenticationClientFixture),
		axiom.WithRunnerFixture(grpcfixtures.FilesClientFixtureKey, grpcfixtures.SetFilesClientFixture),
		axiom.WithRunnerFixture(grpcfixtures.CoursesClientFixtureKey, grpcfixtures.SetCoursesClientFixture),
		axiom.WithRunnerFixture(grpcfixtures.ExercisesClientFixtureKey, grpcfixtures.SetExercisesClientFixture),

		axiom.WithRunnerFixture(grpcfixtures.UserFixtureKey, grpcfixtures.SetUserFixture),
		axiom.WithRunnerFixture(grpcfixtures.FileFixtureKey, grpcfixtures.SetFileFixture),
		axiom.WithRunnerFixture(grpcfixtures.CourseFixtureKey, grpcfixtures.SetCourseFixture),
		axiom.WithRunnerFixture(grpcfixtures.SessionFixtureKey, grpcfixtures.SetSessionFixture),
	),
)

var usersRunner = suiteRunner.Join(
	axiom.NewRunner(
		axiom.WithRunnerMeta(
			axiom.WithMetaTag(metadata.TagUsers),
			axiom.WithMetaFeature(metadata.FeatureUsers),
		),
	),
)

var authenticationRunner = suiteRunner.Join(
	axiom.NewRunner(
		axiom.WithRunnerMeta(
			axiom.WithMetaTag(metadata.TagAuthentication),
			axiom.WithMetaFeature(metadata.FeatureAuthentication),
		),
	),
)
