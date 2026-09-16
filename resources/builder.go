package resources

import (
	"github.com/Nikita-Filonov/axiom"
	"github.com/PloxoSpaal/go-api-autotests/builders"
)

const BuilderResourceKey = "builder"

func SetBuilderResource(r *axiom.Runner) (any, func(), error) {
	generator := GetFakeResource(r)
	return builders.New(generator), nil, nil
}

func GetBuilderResource(runner *axiom.Runner) *builders.Builder {
	return axiom.MustResource[*builders.Builder](runner, BuilderResourceKey)
}
