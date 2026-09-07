package tests

import "github.com/Nikita-Filonov/axiom"

var testRunner = axiom.NewRunner(
	axiom.WithRunnerMeta(
		axiom.WithMetaTag("axiom"),
	),
)
