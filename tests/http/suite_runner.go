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
		axiom.WithRunnerFixture(httpfixtures.AuthenticationClientFixtureKey, httpfixtures.SetAuthenticationClientFixture),
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
