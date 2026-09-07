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
)
