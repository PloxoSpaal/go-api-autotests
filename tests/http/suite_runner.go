package http

import (
	"github.com/Nikita-Filonov/axiom"
	"github.com/PloxoSpaal/go-api-autotests/fixtures/httpfixtures"
	"github.com/PloxoSpaal/go-api-autotests/metadata"
	"github.com/PloxoSpaal/go-api-autotests/tests"
)

var suiteRunner = tests.Runner.Join(
	axiom.NewRunner(
		axiom.WithRunnerMeta(
			axiom.WithMetaTag(metadata.TagHTTP),
			axiom.WithMetaSuite(metadata.SuiteHTTP),
		),
		axiom.WithRunnerFixture(httpfixtures.PublicTransportFixtureKey, httpfixtures.SetPublicTransportFixture),
		axiom.WithRunnerFixture(httpfixtures.PrivateTransportFixtureKey, httpfixtures.SetPrivateTransportFixture),

		axiom.WithRunnerFixture(httpfixtures.AuthenticationClientFixtureKey, httpfixtures.SetAuthenticationClientFixture),
		axiom.WithRunnerFixture(httpfixtures.PublicUsersClientFixtureKey, httpfixtures.SetPublicUsersClientFixture),
		axiom.WithRunnerFixture(httpfixtures.PrivateUsersClientFixtureKey, httpfixtures.SetPrivateUsersClientFixture),
		axiom.WithRunnerFixture(httpfixtures.FilesClientFixtureKey, httpfixtures.SetFilesClientFixture),
		axiom.WithRunnerFixture(httpfixtures.CoursesClientFixtureKey, httpfixtures.SetCoursesClientFixture),
		axiom.WithRunnerFixture(httpfixtures.ExercisesClientFixtureKey, httpfixtures.SetExercisesClientFixture),

		axiom.WithRunnerFixture(httpfixtures.UserFixtureKey, httpfixtures.SetUserFixture),
		axiom.WithRunnerFixture(httpfixtures.SessionFixtureKey, httpfixtures.SetSessionFixture),
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
