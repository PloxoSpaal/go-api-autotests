package tests

import "github.com/Nikita-Filonov/axiom"

var testRunner = axiom.NewRunner(
	axiom.WithRunnerMeta(
		axiom.WithMetaEpic("Learning platform"),
		axiom.WithMetaParentSuite("Axiom examples"),
		axiom.WithMetaLayer("unit"),
		axiom.WithMetaPlatform("go"),
		axiom.WithMetaSeverity(axiom.SeverityNormal),
		axiom.WithMetaTag("axiom"),
		axiom.WithMetaLabel("owner", "qa-platform"),
	),
	axiom.WithRunnerContext(
		axiom.WithContextData("environment", "local"),
	),
	axiom.WithRunnerHooks(
		axiom.WithBeforeAll(logBeforeAll),
		axiom.WithAfterAll(logAfterAll),
		axiom.WithBeforeTest(logBeforeTest),
		axiom.WithAfterTest(logAfterTest),
		axiom.WithBeforeStep(logBeforeStep),
		axiom.WithAfterStep(logAfterStep),
	),
	axiom.WithRunnerPlugins(durationPlugin()),
)
