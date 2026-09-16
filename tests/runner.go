package tests

import (
	"github.com/Nikita-Filonov/axiom"
	"github.com/PloxoSpaal/go-api-autotests/metadata"
	"github.com/PloxoSpaal/go-api-autotests/resources"
)

var Runner = axiom.NewRunner(
	axiom.WithRunnerMeta(
		axiom.WithMetaTags(metadata.TagAPI, metadata.TagRegression),
		axiom.WithMetaEpic(metadata.EpicAPI),
		axiom.WithMetaLayer(metadata.LayerE2E),
	),
	axiom.WithRunnerResource(resources.ConfigResourceKey, resources.SetConfigResource),
	axiom.WithRunnerResource(resources.FakeResourceKey, resources.SetFakeResource),
	axiom.WithRunnerResource(resources.BuilderResourceKey, resources.SetBuilderResource),
)
